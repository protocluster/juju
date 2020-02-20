// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package space_test

import (
	"bytes"
	"fmt"

	"github.com/golang/mock/gomock"
	"github.com/juju/cmd"
	"github.com/juju/cmd/cmdtesting"
	"github.com/juju/errors"
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/apiserver/params"
	"github.com/juju/juju/cmd/juju/space"
	"github.com/juju/juju/cmd/juju/space/mocks"
	"github.com/juju/juju/core/network"
)

type ShowSuite struct {
	BaseSpaceSuite
}

var _ = gc.Suite(&ShowSuite{})

func (s *ShowSuite) SetUpTest(c *gc.C) {
	s.BaseSpaceSuite.SetUpTest(c)
	s.newCommand = space.NewShowSpaceCommand
}

func setUpMocks(c *gc.C) (*gomock.Controller, *mocks.MockSpaceAPI) {
	ctrl := gomock.NewController(c)
	api := mocks.NewMockSpaceAPI(ctrl)
	api.EXPECT().Close()
	return ctrl, api
}

func (s *ShowSuite) getDefaultSpace() network.ShowSpace {
	result := network.ShowSpace{
		Space: network.SpaceInfo{
			ID:   "1",
			Name: "default",
			Subnets: []network.SubnetInfo{{
				CIDR:       "4.3.2.0/28",
				ProviderId: "abc",
				VLANTag:    0,
			}},
		},
		Applications: []string{"ubuntu,mysql"},
		MachineCount: 4,
	}
	return result
}

func (s *ShowSuite) TestRunShowSpaceSucceedsMock(c *gc.C) {
	ctrl, api := setUpMocks(c)
	defer ctrl.Finish()
	spaceName := "default"
	expectedStdout := `space:
  id: "1"
  name: default
  subnets:
  - cidr: 4.3.2.0/28
    provider-id: abc
    vlan-tag: 0
applications:
- ubuntu,mysql
machine-count: 4
`
	api.EXPECT().ShowSpace(spaceName).Return(s.getDefaultSpace(), nil)

	ctx, err := s.runCommand(c, api, spaceName)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(cmdtesting.Stdout(ctx), gc.Equals, expectedStdout)

}
func (s *ShowSuite) runCommand(c *gc.C, api space.SpaceAPI, name string) (*cmd.Context, error) {
	base := space.NewSpaceCommandBase(api)
	command := space.ShowSpaceCommand{
		SpaceCommandBase: base,
		Name:             "",
	}
	return cmdtesting.RunCommand(c, &command, name)
}

func (s *ShowSuite) TestRunWhenShowSpacesNotSupported(c *gc.C) {
	ctrl, api := setUpMocks(c)
	defer ctrl.Finish()
	spaceName := "default"

	expectedErr := errors.NewNotSupported(nil, "spaces not supported")
	api.EXPECT().ShowSpace(spaceName).Return(network.ShowSpace{}, expectedErr)

	_, err := s.runCommand(c, api, spaceName)

	c.Assert(err, jc.Satisfies, errors.IsNotSupported)
}

func (s *ShowSuite) TestRunWhenShowSpacesAPIFails(c *gc.C) {
	ctrl, api := setUpMocks(c)
	defer ctrl.Finish()
	spaceName := "default"

	apiErr := errors.New("API error")
	api.EXPECT().ShowSpace(spaceName).Return(network.ShowSpace{}, apiErr)

	_, err := s.runCommand(c, api, spaceName)
	expectedMsg := fmt.Sprintf("cannot retrieve space %q: API error", spaceName)
	c.Assert(err, gc.ErrorMatches, expectedMsg)
}

func (s *ShowSuite) TestRunUnauthorizedMentionsJujuGrant(c *gc.C) {
	apiErr := &params.Error{
		Message: "permission denied",
		Code:    params.CodeUnauthorized,
	}

	ctrl, api := setUpMocks(c)
	defer ctrl.Finish()
	spaceName := "default"
	api.EXPECT().ShowSpace(spaceName).Return(network.ShowSpace{}, apiErr)

	_, err := s.runCommand(c, api, spaceName)
	expectedErrMsg := fmt.Sprintf("cannot retrieve space %q: permission denied", spaceName)
	c.Assert(err, gc.ErrorMatches, expectedErrMsg)

}

func (s *ShowSuite) TestRunWhenSpacesBlocked(c *gc.C) {
	apiErr := &params.Error{Code: params.CodeOperationBlocked, Message: "nope"}
	ctrl, api := setUpMocks(c)
	defer ctrl.Finish()

	spaceName := "default"
	api.EXPECT().ShowSpace(spaceName).Return(network.ShowSpace{}, apiErr)
	ctx, err := s.runCommand(c, api, spaceName)

	c.Assert(err, gc.ErrorMatches, `
cannot retrieve space "default": nope

All operations that change model have been disabled for the current model.
To enable changes, run

    juju enable-command all

`[1:])
	c.Assert(ctx.Stderr.(*bytes.Buffer).String(), gc.Equals, "")
	c.Assert(ctx.Stdout.(*bytes.Buffer).String(), gc.Equals, "")
}
