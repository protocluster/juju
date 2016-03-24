// Copyright 2016 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package proxyupdater

import (
	"strings"

	"github.com/juju/juju/apiserver/common"
	"github.com/juju/juju/apiserver/params"
	"github.com/juju/juju/environs/config"
	"github.com/juju/juju/network"
	"github.com/juju/juju/state"
)

func init() {
	common.RegisterStandardFacade("ProxyUpdater", 1, NewAPI)
}

// API defines the methods the ProxyUpdater API facade implements.
type API interface {
	WatchForProxyConfigAndAPIHostPortChanges() state.NotifyWatcher
	ProxyConfig() (params.ProxyConfigResult, error)
}

// Backing defines the state methods this facede needs, so they can be
// mocked for testing.
type State interface {
	EnvironConfig() (*config.Config, error)
	APIHostPorts() ([][]network.HostPort, error)
	WatchAPIHostPorts() state.NotifyWatcher
	WatchForEnvironConfigChanges() state.NotifyWatcher
}

type proxyUpdaterAPI struct {
	st         State
	resources  *common.Resources
	authorizer common.Authorizer
}

// NewAPI creates a new API server-side facade with a state.State backing.
func NewAPI(st *state.State, res *common.Resources, auth common.Authorizer) (API, error) {
	return newAPIWithBacking(st, res, auth)
}

// newAPIWithBacking creates a new server-side API facade with the given Backing.
func newAPIWithBacking(st State, resources *common.Resources, authorizer common.Authorizer) (API, error) {
	if !authorizer.AuthClient() {
		return nil, common.ErrPerm
	}
	return &proxyUpdaterAPI{
		st:         st,
		resources:  resources,
		authorizer: authorizer,
	}, nil
}

// WatchForProxyConfigAndAPIHostPortChanges watches both the environment config and the API host
// ports for changes.
func (api *proxyUpdaterAPI) WatchForProxyConfigAndAPIHostPortChanges() state.NotifyWatcher {
	return common.NewMultiNotifyWatcher(
		api.st.WatchForEnvironConfigChanges(),
		api.st.WatchAPIHostPorts())
}

func (api *proxyUpdaterAPI) ProxyConfig() (params.ProxyConfigResult, error) {
	var cfg params.ProxyConfigResult
	env, err := api.st.EnvironConfig()
	if err != nil {
		return cfg, err
	}
	apiHostPorts, err := api.st.APIHostPorts()

	cfg.ProxySettings = env.ProxySettings()
	cfg.APTProxySettings = env.AptProxySettings()
	noProxy := strings.Split(cfg.ProxySettings.NoProxy, ",")
	noAptProxy := strings.Split(cfg.APTProxySettings.NoProxy, ",")

	for _, host := range apiHostPorts {
		for _, hp := range host {
			noProxy = append(noProxy, hp.Address.Value)
			noAptProxy = append(noAptProxy, hp.Address.Value)
		}
	}

	cfg.ProxySettings.NoProxy = strings.Join(noProxy, ",")
	cfg.APTProxySettings.NoProxy = strings.Join(noAptProxy, ",")

	return cfg, nil
}
