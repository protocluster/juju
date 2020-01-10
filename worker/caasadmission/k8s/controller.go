package k8s

import (
	"context"
	_ "encoding/json"
	"net/http"
	"time"

	"github.com/juju/errors"
)

type Controller struct {
	server *http.Server
}

type Logger interface {
	Debugf(string, ...interface{})
	Errorf(string, ...interface{})
	Infof(string, ...interface{})
}

const (
	DefaultPort = "4433"
	StopTimeout = time.Second * 120
)

func NewController(logger Logger) *Controller {
	c := &Controller{
		server: &http.Server{
			Addr:    ":" + DefaultPort,
			Handler: admissionHandler(logger),
		},
	}
	return c
}

func (c *Controller) Wait() error {
	if c.server == nil {
		return errors.NewNotValid(nil, "controller server is nil")
	}

	if err := c.server.ListenAndServe(); err != http.ErrServerClosed {
		return errors.Wrapf(err, nil, "starting k8s admission webhook server")
	}
	return nil
}

func (c *Controller) Kill() {
	if c.server == nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), StopTimeout)
	defer cancel()
	c.server.Shutdown(ctx)
}
