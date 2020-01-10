package caasadmission

import (
	"github.com/juju/juju/worker/caasadmission/k8s"

	"gopkg.in/juju/worker.v1"
	"gopkg.in/juju/worker.v1/dependency"
)

type Logger interface {
	Debugf(string, ...interface{})
	Errorf(string, ...interface{})
	Infof(string, ...interface{})
}

type ManifoldConfig struct {
	Logger Logger
}

func Manifold(config ManifoldConfig) dependency.Manifold {
	return dependency.Manifold{
		Inputs: []string{},
		Output: nil,
		Start: func(ctx dependency.Context) (worker.Worker, error) {
			return k8s.NewController(config.Logger), nil
		},
	}
}
