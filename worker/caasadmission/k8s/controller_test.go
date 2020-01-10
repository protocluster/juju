package k8s

import (
	"testing"

	gc "gopkg.in/check.v1"
)

type K8sSuite struct {
}

var _ = gc.Suite(&K8sSuite{})

func TestK8sSuite(t *testing.T) { gc.TestingT(t) }
