// Copyright 2020 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package test

import (
	"github.com/juju/errors"
	"k8s.io/apimachinery/pkg/types"
)

type RBACMapper struct {
	AppNameForServiceAccountFunc func(types.UID) (string, error)
}

func (r *RBACMapper) AppNameForServiceAccount(t types.UID) (string, error) {
	if r.AppNameForServiceAccountFunc == nil {
		return "", errors.NotFoundf("no service account for app found with id %v", t)
	}
	return r.AppNameForServiceAccountFunc(t)
}

func (r *RBACMapper) Kill() {
}

func (r *RBACMapper) Wait() error {
	return nil
}
