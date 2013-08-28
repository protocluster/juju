// Copyright 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package agent

import (
	"fmt"

	"launchpad.net/juju-core/state"
	"launchpad.net/juju-core/state/api"
)

// formatter114 is the formatter for the 1.14 format.
type formatter114 struct {
}

// conf holds information for a given agent.
type xconf struct {
	// DataDir specifies the path of the data directory used by all
	// agents
	dataDir string

	// StateServerCert and StateServerKey hold the state server
	// certificate and private key in PEM format.
	StateServerCert []byte `yaml:",omitempty"`
	StateServerKey  []byte `yaml:",omitempty"`

	StatePort int `yaml:",omitempty"`
	APIPort   int `yaml:",omitempty"`

	// OldPassword specifies a password that should be
	// used to connect to the state if StateInfo.Password
	// is blank or invalid.
	OldPassword string

	// MachineNonce is set at provisioning/bootstrap time and used to
	// ensure the agent is running on the correct instance.
	MachineNonce string

	// StateInfo specifies how the agent should connect to the
	// state.  The password may be empty if an old password is
	// specified, or when bootstrapping.
	StateInfo *state.Info `yaml:",omitempty"`

	// OldAPIPassword specifies a password that should
	// be used to connect to the API if APIInfo.Password
	// is blank or invalid.
	OldAPIPassword string

	// APIInfo specifies how the agent should connect to the
	// state through the API.
	APIInfo *api.Info `yaml:",omitempty"`
}

// Ensure that the formatter114 struct implements the formatter interface.
var _ formatter = (*formatter114)(nil)

func (*formatter114) read(dirName string) (*configInternal, error) {
	return nil, fmt.Errorf("not implemented")
}

func (*formatter114) write(dirName string, config *configInternal) error {
	return fmt.Errorf("not implemented")
}

func (*formatter114) writeCommands(dirName string, config *configInternal) ([]string, error) {
	return nil, fmt.Errorf("not implemented")
}

// ReadConf reads configuration data for the given
// entity from the given data directory.
// func xReadConf(dataDir, tag string) (Config, error) {
// 	dir := tools.Dir(dataDir, tag)
// 	data, err := ioutil.ReadFile(path.Join(dir, "agent.conf"))
// 	if err != nil {
// 		return nil, err
// 	}
// 	var c conf
// 	if err := goyaml.Unmarshal(data, &c); err != nil {
// 		return nil, err
// 	}
// 	c.dataDir = dataDir
// 	if err := c.check(); err != nil {
// 		return nil, err
// 	}
// 	if c.StateInfo != nil {
// 		c.StateInfo.Tag = tag
// 	}
// 	if c.APIInfo != nil {
// 		c.APIInfo.Tag = tag
// 	}
// 	return &c, nil
// }
