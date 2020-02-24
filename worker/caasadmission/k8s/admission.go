package k8s

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"github.com/juju/errors"
	admission "k8s.io/api/admissionregistration/v1beta1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/juju/juju/caas/kubernetes/provider"
)

// Represents a creator of mutating webhooks that is context aware of the
// current controller.
type AdmissionCreator interface {
	EnsureMutatingWebhookConfiguration() (func(), error)
}

// Func type of AdmissionCreator
type AdmissionCreatorFunc func() (func(), error)

const (
	Component = "admission"
)

var (
	anyMatch = []string{"*"}
)

// Implements AdmissionCreator interface for func type
func (a AdmissionCreatorFunc) EnsureMutatingWebhookConfiguration() (func(), error) {
	return a()
}

// Instantiates a new AdmissionCreator for the supplied context arguments.
func NewAdmissionCreator(
	certWatcher func() *tls.Certificate,
	namespace, model string,
	ensureConfig func(*admission.MutatingWebhookConfiguration) (func(), error),
	service *admission.ServiceReference) (AdmissionCreator, error) {

	caPems := &bytes.Buffer{}
	for _, derCert := range certWatcher().Certificate {
		cert, err := x509.ParseCertificate(derCert)
		if err != nil {
			return nil, errors.Trace(err)
		}
		if !cert.IsCA {
			continue
		}
		if err := pem.Encode(caPems, &pem.Block{
			Type:  "CERTIFICATE",
			Bytes: derCert}); err != nil {
			return nil, errors.Trace(err)
		}
	}
	if caPems.Len() == 0 {
		return nil, errors.NewNotValid(nil, "no CA certificates available for controller")
	}

	failurePolicy := admission.Ignore
	matchPolicy := admission.Equivalent
	ruleScope := admission.NamespacedScope

	// MutatingWebjook Obj
	obj := admission.MutatingWebhookConfiguration{
		ObjectMeta: meta.ObjectMeta{
			Labels:    provider.LabelsForModel(model),
			Name:      fmt.Sprintf("%s-model-admission", model),
			Namespace: namespace,
		},
		Webhooks: []admission.MutatingWebhook{
			admission.MutatingWebhook{
				ClientConfig: admission.WebhookClientConfig{
					CABundle: caPems.Bytes(),
					Service:  service,
				},
				FailurePolicy: &failurePolicy,
				MatchPolicy:   &matchPolicy,
				Name:          provider.MakeK8sDomain(Component),
				NamespaceSelector: &meta.LabelSelector{
					MatchLabels: provider.LabelsForModel(model),
				},
				Rules: []admission.RuleWithOperations{
					admission.RuleWithOperations{
						Operations: []admission.OperationType{
							admission.Create,
							admission.Update,
						},
						Rule: admission.Rule{
							APIGroups:   anyMatch,
							APIVersions: anyMatch,
							Resources:   anyMatch,
							Scope:       &ruleScope,
						},
					},
				},
			},
		},
	}
	return AdmissionCreatorFunc(func() (func(), error) {
		return ensureConfig(&obj)
	}), nil
}
