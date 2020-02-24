package caasadmission

import (
	"crypto/tls"

	"github.com/juju/juju/agent"
	"github.com/juju/juju/apiserver/apiserverhttp"
	"github.com/juju/juju/caas"
	"github.com/juju/juju/caas/kubernetes/provider"
	"github.com/juju/juju/worker/caasadmission/k8s"

	"github.com/juju/errors"
	"gopkg.in/juju/worker.v1"
	"gopkg.in/juju/worker.v1/dependency"
	admission "k8s.io/api/admissionregistration/v1beta1"
)

type Logger interface {
	Debugf(string, ...interface{})
	Errorf(string, ...interface{})
	Infof(string, ...interface{})
}

type ManifoldConfig struct {
	AgentName   string
	BrokerName  string
	CertWatcher func() *tls.Certificate
	Logger      Logger
	Mux         *apiserverhttp.Mux
}

func Manifold(config ManifoldConfig) dependency.Manifold {
	return dependency.Manifold{
		Inputs: []string{
			config.AgentName,
			config.BrokerName,
		},
		Output: nil,
		Start:  config.Start,
	}
}

func (c ManifoldConfig) Start(context dependency.Context) (worker.Worker, error) {
	if err := c.Validate(); err != nil {
		return nil, errors.Trace(err)
	}

	var agent agent.Agent
	if err := context.Get(c.AgentName, &agent); err != nil {
		return nil, errors.Trace(err)
	}

	var broker caas.Broker
	if err := context.Get(c.BrokerName, &broker); err != nil {
		return nil, errors.Trace(err)
	}

	k8sClient, ok := broker.(*provider.KubernetesClient)
	if !ok {
		return nil, errors.NewNotValid(nil, "cannot cast caas broker to KubernetesClient")
	}

	currentConfig := agent.CurrentConfig()
	admissionPath := k8s.AdmissionPathForModel(currentConfig.Model().Id())
	port := int32(17070)
	admissionCreator, err := k8s.NewAdmissionCreator(c.CertWatcher,
		k8sClient.GetCurrentNamespace(), k8sClient.GetCurrentModel(),
		k8sClient.EnsureMutatingWebhookConfiguration,
		&admission.ServiceReference{
			Name:      "controller-service",
			Namespace: "controller-microk8s-localhost",
			Path:      &admissionPath,
			Port:      &port,
		},
	)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return k8s.NewController(
		c.Logger,
		c.Mux,
		k8s.AdmissionPathForModel(currentConfig.Model().Id()),
		admissionCreator,
		k8sClient.RBACMapper())
}

func (c ManifoldConfig) Validate() error {
	if c.AgentName == "" {
		return errors.NotValidf("empty AgentName")
	}
	if c.BrokerName == "" {
		return errors.NotValidf("empty BrokerName")
	}
	if c.CertWatcher == nil {
		return errors.NotValidf("nil CertWatcher")
	}
	if c.Mux == nil {
		return errors.NotValidf("nil apiserverhttp.Mux reference")
	}
	return nil
}
