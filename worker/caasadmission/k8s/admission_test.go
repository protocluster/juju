package k8s_test

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"time"

	"github.com/juju/juju/worker/caasadmission/k8s"

	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"
	admission "k8s.io/api/admissionregistration/v1beta1"
)

type AdmissionSuite struct {
}

type dummyAdmissionCreator struct {
	EnsureMutatingWebhookConfigurationFunc func() (func(), error)
}

type dummyCertWatcher struct {
	certificate *tls.Certificate
}

var _ = gc.Suite(&AdmissionSuite{})

func (d *dummyCertWatcher) GetCertificate(c *gc.C) func() *tls.Certificate {
	return func() *tls.Certificate {
		if d.certificate != nil {
			return d.certificate
		}
		priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		c.Assert(err, jc.ErrorIsNil)

		template := x509.Certificate{
			BasicConstraintsValid: true,
			IsCA:                  true,
			KeyUsage:              x509.KeyUsageCertSign,
			NotAfter:              time.Now().Add(time.Minute * 5),
			NotBefore:             time.Now(),
			SerialNumber:          big.NewInt(10),
			Subject: pkix.Name{
				Organization: []string{"JuJu"},
			},
		}

		derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template,
			&priv.PublicKey, priv)
		c.Assert(err, jc.ErrorIsNil)

		d.certificate = &tls.Certificate{
			Certificate: [][]byte{derBytes},
			PrivateKey:  priv,
		}

		return d.certificate
	}
}

func (d *dummyAdmissionCreator) EnsureMutatingWebhookConfiguration() (func(), error) {
	if d.EnsureMutatingWebhookConfigurationFunc == nil {
		return func() {}, nil
	}
	return d.EnsureMutatingWebhookConfigurationFunc()
}

func int32Ptr(i int32) *int32 {
	return &i
}

func strPtr(s string) *string {
	return &s
}

func (a *AdmissionSuite) TestAdmissionCreatorObject(c *gc.C) {
	certWatcher := dummyCertWatcher{}
	serviceRef := &admission.ServiceReference{
		Namespace: "testns",
		Name:      "test",
		Path:      strPtr("/test"),
		Port:      int32Ptr(1111),
	}

	admissionCreator, err := k8s.NewAdmissionCreator(
		certWatcher.GetCertificate(c), "testns", "testmodel",
		func(obj *admission.MutatingWebhookConfiguration) (func(), error) {
			c.Assert(obj.Namespace, gc.Equals, "testns")
			return func() {}, nil
		}, serviceRef)
	c.Assert(err, jc.ErrorIsNil)

	_, err = admissionCreator.EnsureMutatingWebhookConfiguration()
	c.Assert(err, jc.ErrorIsNil)
}
