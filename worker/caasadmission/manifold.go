// Copyright 2020 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package caasadmission

import (
	"crypto/tls"

	"github.com/juju/errors"
	"gopkg.in/juju/worker.v1"
	"gopkg.in/juju/worker.v1/dependency"
	admission "k8s.io/api/admissionregistration/v1beta1"

	"github.com/juju/juju/agent"
	"github.com/juju/juju/apiserver/apiserverhttp"
	"github.com/juju/juju/worker/caasrbacmapper"
)

// K8sBroker describes a Kubernetes broker interface this worker needs to
// function.
type K8sBroker interface {
	// CurrentModel returns the current model the broker is targetting
	CurrentModel() string

	// GetCurrentNamespace returns the current namespace being targeted on the
	// broker
	GetCurrentNamespace() string

	// EnsureMutatingWebhookConfiguration make the supplied webhook config exist
	// inside the k8s cluster if it currently does not. Return values is a
	// cleanup function that will destroy the webhook configuration from k8s
	// when called and a subsequent error if there was a problem. If error is
	// not nil then no other return values should be considered valid.
	EnsureMutatingWebhookConfiguration(*admission.MutatingWebhookConfiguration) (func(), error)
}

// Logger represents the methods used by the worker to log details
type Logger interface {
	Debugf(string, ...interface{})
	Errorf(string, ...interface{})
	Infof(string, ...interface{})
}

// ManifoldConfig describes the resources used by the admission worker
type ManifoldConfig struct {
	AgentName  string
	BrokerName string
	CertGetter func() *tls.Certificate
	CertificateBroker
	Logger         Logger
	Mux            *apiserverhttp.Mux
	RBACMapperName string
}

// Manifold returns a Manifold that encapsulates a Kubernetes mutating admission
// controller. Manifold has no outputs.
func Manifold(config ManifoldConfig) dependency.Manifold {
	return dependency.Manifold{
		Inputs: []string{
			config.AgentName,
			config.BrokerName,
			config.RBACMapperName,
		},
		Output: nil,
		Start:  config.Start,
	}
}

// Start is used to start the manifold an extract a worker from the supplied
// configuration.
func (c ManifoldConfig) Start(context dependency.Context) (worker.Worker, error) {
	if err := c.Validate(); err != nil {
		return nil, errors.Trace(err)
	}

	var agent agent.Agent
	if err := context.Get(c.AgentName, &agent); err != nil {
		return nil, errors.Trace(err)
	}

	var broker K8sBroker
	if err := context.Get(c.BrokerName, &broker); err != nil {
		return nil, errors.Trace(err)
	}

	var rbacMapper caasrbacmapper.Mapper
	if err := context.Get(c.RBACMapperName, &rbacMapper); err != nil {
		return nil, errors.Trace(err)
	}

	currentConfig := agent.CurrentConfig()
	admissionPath := AdmissionPathForModel(currentConfig.Model().Id())
	port := int32(17070)
	certBroker := &CertWatcherBroker{c.CertGetter}
	admissionCreator, err := NewAdmissionCreator(certBroker,
		broker.GetCurrentNamespace(), broker.CurrentModel(),
		broker.EnsureMutatingWebhookConfiguration,
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

	return NewController(
		c.Logger,
		c.Mux,
		AdmissionPathForModel(currentConfig.Model().Id()),
		admissionCreator,
		rbacMapper)
}

// Validate is used to to establish if the configuration is valid for use when
// creating new workers.
func (c ManifoldConfig) Validate() error {
	if c.AgentName == "" {
		return errors.NotValidf("empty AgentName")
	}
	if c.BrokerName == "" {
		return errors.NotValidf("empty BrokerName")
	}
	//if c.CertificateBroker == nil {
	//	return errors.NotValidf("nil certificate broker")
	//}
	if c.CertGetter == nil {
		return errors.NotValidf("nil CertGetter")
	}
	if c.Logger == nil {
		return errors.NotValidf("nil Logger")
	}
	if c.Mux == nil {
		return errors.NotValidf("nil apiserverhttp.Mux reference")
	}
	if c.RBACMapperName == "" {
		return errors.NotValidf("empty RBACMapperName")
	}
	return nil
}
