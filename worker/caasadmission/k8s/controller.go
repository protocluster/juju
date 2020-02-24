package k8s

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/juju/errors"
	"gopkg.in/juju/worker.v1/catacomb"

	"github.com/juju/juju/caas/kubernetes/provider"
)

type Mux interface {
	AddHandler(string, string, http.Handler) error
	RemoveHandler(string, string)
}

type Logger interface {
	Debugf(string, ...interface{})
	Errorf(string, ...interface{})
	Infof(string, ...interface{})
}

// Kubernetes controller responsible
type Controller struct {
	admissionCreator AdmissionCreator
	catacomb         catacomb.Catacomb
	logger           Logger
	mux              Mux
	path             string
	rbacMapper       provider.RBACMapper
}

func AdmissionPathForModel(modelUUID string) string {
	return fmt.Sprintf("/k8s/admission/%s", url.PathEscape(modelUUID))
}

func NewController(
	logger Logger,
	mux Mux,
	path string,
	admissionCreator AdmissionCreator,
	rbacMapper provider.RBACMapper) (*Controller, error) {

	c := &Controller{
		admissionCreator: admissionCreator,
		logger:           logger,
		mux:              mux,
		path:             path,
		rbacMapper:       rbacMapper,
	}

	if err := catacomb.Invoke(catacomb.Plan{
		Site: &c.catacomb,
		Work: c.loop,
	}); err != nil {
		return c, errors.Trace(err)
	}

	if err := c.catacomb.Add(c.rbacMapper); err != nil {
		return c, errors.Trace(err)
	}

	return c, nil
}

func (c *Controller) Kill() {
	c.catacomb.Kill(nil)
}

func (c *Controller) loop() error {
	if err := c.mux.AddHandler(http.MethodPost, c.path,
		admissionHandler(c.logger, c.rbacMapper)); err != nil {
		return errors.Trace(err)
	}
	defer c.mux.RemoveHandler(http.MethodPost, c.path)

	admissionCleanup, err := c.admissionCreator.EnsureMutatingWebhookConfiguration()
	if err != nil {
		return errors.Trace(err)
	}
	defer admissionCleanup()

	select {
	case <-c.catacomb.Dying():
		return c.catacomb.ErrDying()
	}
	return nil
}

func (c *Controller) Wait() error {
	return c.catacomb.Wait()
}
