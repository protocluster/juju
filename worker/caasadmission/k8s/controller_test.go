package k8s_test

import (
	"net/http"
	"sync"
	"testing"

	k8stest "github.com/juju/juju/caas/kubernetes/provider/test"
	"github.com/juju/juju/worker/caasadmission/k8s"

	"github.com/juju/errors"
	"github.com/juju/loggo"
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"
)

type ControllerSuite struct {
}

type dummyMux struct {
	AddHandlerFunc    func(string, string, http.Handler) error
	RemoveHandlerFunc func(string, string)
}

var _ = gc.Suite(&ControllerSuite{})

func TestControllerSuite(t *testing.T) { gc.TestingT(t) }

func (d *dummyMux) AddHandler(i, j string, h http.Handler) error {
	if d.AddHandlerFunc == nil {
		return nil
	}
	return d.AddHandlerFunc(i, j, h)
}

func (d *dummyMux) RemoveHandler(i, j string) {
	if d.RemoveHandlerFunc != nil {
		d.RemoveHandlerFunc(i, j)
	}
}

func (s *ControllerSuite) TestControllerStartup(c *gc.C) {
	var (
		logger     = loggo.Logger{}
		rbacMapper = &k8stest.RBACMapper{}
		waitGroup  = sync.WaitGroup{}
		path       = "/test"
	)
	// Setup function counter
	waitGroup.Add(2)
	mux := &dummyMux{
		AddHandlerFunc: func(m, p string, _ http.Handler) error {
			c.Assert(m, jc.DeepEquals, http.MethodPost)
			c.Assert(p, jc.DeepEquals, path)
			waitGroup.Done()
			return nil
		},
		RemoveHandlerFunc: func(_, _ string) {
			waitGroup.Done()
		},
	}
	creator := &dummyAdmissionCreator{
		EnsureMutatingWebhookConfigurationFunc: func() (func(), error) {
			waitGroup.Done()
			return func() { waitGroup.Done() }, nil
		},
	}

	ctrl, err := k8s.NewController(logger, mux, path, creator, rbacMapper)
	c.Assert(err, jc.ErrorIsNil)

	waitGroup.Wait()
	waitGroup.Add(2)
	ctrl.Kill()

	// Cleanup function counter
	waitGroup.Wait()
	err = ctrl.Wait()
	c.Assert(err, jc.ErrorIsNil)
}

func (s *ControllerSuite) TestControllerStartupMuxError(c *gc.C) {
	var (
		logger     = loggo.Logger{}
		rbacMapper = &k8stest.RBACMapper{}
		waitGroup  = sync.WaitGroup{}
		path       = "/test"
	)
	// Setup function counter
	waitGroup.Add(1)
	mux := &dummyMux{
		AddHandlerFunc: func(m, p string, _ http.Handler) error {
			waitGroup.Done()
			c.Assert(m, jc.DeepEquals, http.MethodPost)
			c.Assert(p, jc.DeepEquals, path)
			return errors.NewNotValid(nil, "not valid")
		},
	}
	creator := &dummyAdmissionCreator{}

	ctrl, err := k8s.NewController(logger, mux, path, creator, rbacMapper)
	c.Assert(err, jc.ErrorIsNil)

	waitGroup.Wait()
	ctrl.Kill()
	err = ctrl.Wait()
	c.Assert(errors.IsNotValid(err), jc.IsTrue)
}

func (s *ControllerSuite) TestControllerStartupAdmissionError(c *gc.C) {
	var (
		logger     = loggo.Logger{}
		rbacMapper = &k8stest.RBACMapper{}
		waitGroup  = sync.WaitGroup{}
		path       = "/test"
	)
	// Setup function counter
	waitGroup.Add(1)
	mux := &dummyMux{}
	creator := &dummyAdmissionCreator{
		EnsureMutatingWebhookConfigurationFunc: func() (func(), error) {
			waitGroup.Done()
			return func() {}, errors.NewNotValid(nil, "not valid")
		},
	}

	ctrl, err := k8s.NewController(logger, mux, path, creator, rbacMapper)
	c.Assert(err, jc.ErrorIsNil)

	waitGroup.Wait()
	ctrl.Kill()
	err = ctrl.Wait()
	c.Assert(errors.IsNotValid(err), jc.IsTrue)
}
