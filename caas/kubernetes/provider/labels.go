// Copyright 2020 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package provider

func LabelsForModel(model string) map[string]string {
	return map[string]string{
		labelModel: model,
	}
}
