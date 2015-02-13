// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package upgrades_test

import (
	jc "github.com/juju/testing/checkers"
	"github.com/juju/utils/featureflag"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/juju/osenv"
	jujutesting "github.com/juju/juju/juju/testing"
	"github.com/juju/juju/provider/ec2"
	"github.com/juju/juju/state"
	"github.com/juju/juju/storage/pool"
	"github.com/juju/juju/upgrades"
)

type defaultStoragePoolsSuite struct {
	jujutesting.JujuConnSuite
}

var _ = gc.Suite(&defaultStoragePoolsSuite{})

func (s *defaultStoragePoolsSuite) TestDefaultStoragePools(c *gc.C) {
	s.PatchEnvironment(osenv.JujuFeatureFlagEnvKey, "storage")
	featureflag.SetFlagsFromEnvironment(osenv.JujuFeatureFlagEnvKey)

	err := upgrades.AddDefaultStoragePools(s.State, &mockAgentConfig{dataDir: s.DataDir()})
	settings := state.NewStateSettings(s.State)
	err = pool.AddDefaultStoragePools(settings, &mockAgentConfig{dataDir: s.DataDir()})
	c.Assert(err, jc.ErrorIsNil)
	pm := pool.NewPoolManager(settings)
	for _, pName := range []string{"ebs", "ebs-ssd"} {
		p, err := pm.Get(pName)
		c.Assert(err, jc.ErrorIsNil)
		c.Assert(p.Type(), gc.Equals, ec2.EBS_ProviderType)
	}
}
