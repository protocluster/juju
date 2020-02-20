// Copyright 2020 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package spaces_test

import (
	"fmt"

	"github.com/golang/mock/gomock"
	"github.com/juju/errors"
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/apiserver/facades/client/spaces"
	"github.com/juju/juju/apiserver/facades/client/spaces/mocks"
	"github.com/juju/juju/controller"
	"github.com/juju/juju/core/settings"
	"github.com/juju/juju/state"
	"github.com/juju/juju/testing"
)

type SpaceRenameSuite struct {
	state    *mocks.MockRenameSpaceState
	space    *mocks.MockRenameSpace
	settings *mocks.MockSettings
}

var _ = gc.Suite(&SpaceRenameSuite{})

func (s *SpaceRenameSuite) TearDownTest(c *gc.C) {
}

func (s *SpaceRenameSuite) TestSuccess(c *gc.C) {
	toName := "blub"
	fromName := "db"

	ctrl := s.setupMocks(c)
	defer ctrl.Finish()

	s.space.EXPECT().Name().Return(fromName).Times(2)
	s.space.EXPECT().RenameSpaceOps(toName).Return(nil)
	currentConfig := s.getDefaultControllerConfig(c, map[string]interface{}{controller.JujuHASpace: fromName, controller.JujuManagementSpace: "nochange"})
	s.state.EXPECT().ControllerConfig().Return(currentConfig, nil)
	s.state.EXPECT().ConstraintsOpsForSpaceNameChange(fromName, toName).Return(nil, nil)

	expectedConfigDelta := settings.ItemChanges{{
		Type:     1,
		Key:      controller.JujuHASpace,
		OldValue: fromName,
		NewValue: toName,
	}}
	s.settings.EXPECT().DeltaOps(state.ControllerSettingsGlobalKey, expectedConfigDelta).Return(nil, nil)

	op := spaces.NewRenameSpaceModelOp(true, s.settings, s.state, s.space, toName)
	ops, err := op.Build(0)
	c.Assert(err, jc.ErrorIsNil)

	// this is because the code itself does not test for the ops but for expected constraints and delta,
	// which are used to create the ops.
	c.Assert(ops, gc.HasLen, 0)
}

func (s *SpaceRenameSuite) TestNotControllerModelSuccess(c *gc.C) {
	toName := "blub"
	fromName := "db"

	ctrl := s.setupMocks(c)
	defer ctrl.Finish()

	s.space.EXPECT().Name().Return(fromName).Times(1)
	s.state.EXPECT().ConstraintsOpsForSpaceNameChange(fromName, toName).Return(nil, nil)
	s.space.EXPECT().RenameSpaceOps(toName).Return(nil)

	op := spaces.NewRenameSpaceModelOp(false, s.settings, s.state, s.space, toName)
	ops, err := op.Build(0)
	c.Assert(err, jc.ErrorIsNil)

	// this is because the code itself does not test for the ops but for expected constraints and delta,
	// which are used to create the ops.
	c.Assert(ops, gc.HasLen, 0)
}

func (s *SpaceRenameSuite) TestErrorSettingsChanges(c *gc.C) {
	toName := "blub"
	fromName := "db"

	ctrl := s.setupMocks(c)
	defer ctrl.Finish()

	s.space.EXPECT().Name().Return(fromName).Times(2)
	s.state.EXPECT().ConstraintsOpsForSpaceNameChange(fromName, toName).Return(nil, nil)
	s.space.EXPECT().RenameSpaceOps(toName).Return(nil)

	bamErr := errors.New("bam")
	s.state.EXPECT().ControllerConfig().Return(nil, bamErr)

	op := spaces.NewRenameSpaceModelOp(true, s.settings, s.state, s.space, toName)
	ops, err := op.Build(0)
	c.Assert(err, gc.ErrorMatches, fmt.Sprintf("retrieving setting changes: %v", bamErr.Error()))

	c.Assert(ops, gc.HasLen, 0)
}

func (s *SpaceRenameSuite) TestErrorConstraintsChanges(c *gc.C) {
	toName := "blub"
	fromName := "db"

	ctrl := s.setupMocks(c)
	defer ctrl.Finish()

	s.space.EXPECT().Name().Return(fromName).Times(1)

	bamErr := errors.New("bam")
	s.state.EXPECT().ConstraintsOpsForSpaceNameChange(fromName, toName).Return(nil, bamErr)

	op := spaces.NewRenameSpaceModelOp(true, s.settings, s.state, s.space, toName)
	ops, err := op.Build(0)
	c.Assert(err, gc.ErrorMatches, fmt.Sprintf("retrieving constraint changes: %v", bamErr.Error()))

	c.Assert(ops, gc.HasLen, 0)
}

func (s *SpaceRenameSuite) getDefaultControllerConfig(c *gc.C, attr map[string]interface{}) controller.Config {
	cfg, err := controller.NewConfig(testing.ControllerTag.Id(), testing.CACert, attr)
	c.Assert(err, jc.ErrorIsNil)
	return cfg
}

func (s *SpaceRenameSuite) setupMocks(c *gc.C) *gomock.Controller {
	ctrl := gomock.NewController(c)

	s.space = mocks.NewMockRenameSpace(ctrl)
	s.state = mocks.NewMockRenameSpaceState(ctrl)
	s.settings = mocks.NewMockSettings(ctrl)

	return ctrl
}
