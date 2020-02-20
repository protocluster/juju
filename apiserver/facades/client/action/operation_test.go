// Copyright 2020 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package action_test

import (
	"strconv"

	jc "github.com/juju/testing/checkers"
	"github.com/kr/pretty"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/apiserver/params"
	"github.com/juju/juju/state"
)

type operationSuite struct {
	baseSuite
}

var _ = gc.Suite(&operationSuite{})

func (s *operationSuite) setupOperations(c *gc.C) {
	s.toSupportNewActionID(c)

	arg := params.Actions{
		Actions: []params.Action{
			{Receiver: s.wordpressUnit.Tag().String(), Name: "fakeaction", Parameters: map[string]interface{}{}},
			{Receiver: s.mysqlUnit.Tag().String(), Name: "fakeaction", Parameters: map[string]interface{}{}},
			{Receiver: s.wordpressUnit.Tag().String(), Name: "fakeaction", Parameters: map[string]interface{}{}},
			{Receiver: s.mysqlUnit.Tag().String(), Name: "anotherfakeaction", Parameters: map[string]interface{}{}},
		}}

	r, err := s.action.Enqueue(arg)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(r.Results, gc.HasLen, len(arg.Actions))

	// There's only one operation created.
	ops, err := s.Model.AllOperations()
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(ops, gc.HasLen, 1)
	operationID, err := strconv.Atoi(ops[0].Id())
	c.Assert(err, jc.ErrorIsNil)

	a, err := s.Model.Action(strconv.Itoa(operationID + 1))
	c.Assert(err, jc.ErrorIsNil)
	_, err = a.Begin()
	c.Assert(err, jc.ErrorIsNil)
	a, err = s.Model.Action(strconv.Itoa(operationID + 2))
	c.Assert(err, jc.ErrorIsNil)
	_, err = a.Finish(state.ActionResults{Status: state.ActionCompleted})
	c.Assert(err, jc.ErrorIsNil)
}

func (s *operationSuite) TestOperationsStatusFilter(c *gc.C) {
	s.setupOperations(c)
	// Set up a non running operation.
	arg := params.Actions{
		Actions: []params.Action{
			{Receiver: s.wordpressUnit.Tag().String(), Name: "fakeaction", Parameters: map[string]interface{}{}},
		}}
	_, err := s.action.Enqueue(arg)
	c.Assert(err, jc.ErrorIsNil)

	operations, err := s.action.Operations(params.OperationQueryArgs{
		Status: []string{"running"},
	})
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(operations.Truncated, jc.IsFalse)
	c.Assert(operations.Results, gc.HasLen, 1)
	result := operations.Results[0]
	c.Assert(result.Actions, gc.HasLen, 4)
	c.Assert(result.Actions[0].Action, gc.NotNil)
	if result.Enqueued.IsZero() {
		c.Fatal("enqueued time not set")
	}
	if result.Started.IsZero() {
		c.Fatal("started time not set")
	}
	c.Assert(result.Status, gc.Equals, "running")

	action := result.Actions[0].Action
	c.Assert(action.Name, gc.Equals, "fakeaction")
	c.Assert(action.Receiver, gc.Equals, "unit-wordpress-0")
	c.Assert(action.Tag, gc.Equals, "action-2")
	c.Assert(result.Actions[0].Status, gc.Equals, "running")
	action = result.Actions[1].Action
	c.Assert(action.Name, gc.Equals, "fakeaction")
	c.Assert(action.Receiver, gc.Equals, "unit-mysql-0")
	c.Assert(action.Tag, gc.Equals, "action-3")
	c.Assert(result.Actions[1].Status, gc.Equals, "completed")
	action = result.Actions[2].Action
	c.Assert(action.Name, gc.Equals, "fakeaction")
	c.Assert(action.Receiver, gc.Equals, "unit-wordpress-0")
	c.Assert(action.Tag, gc.Equals, "action-4")
	c.Assert(result.Actions[2].Status, gc.Equals, "pending")
	action = result.Actions[3].Action
	c.Assert(action.Name, gc.Equals, "anotherfakeaction")
	c.Assert(action.Receiver, gc.Equals, "unit-mysql-0")
	c.Assert(action.Tag, gc.Equals, "action-5")
	c.Assert(result.Actions[3].Status, gc.Equals, "pending")
}

func (s *operationSuite) TestOperationsNameFilter(c *gc.C) {
	s.setupOperations(c)
	// Set up a second operation.
	arg := params.Actions{
		Actions: []params.Action{
			{Receiver: s.wordpressUnit.Tag().String(), Name: "fakeaction", Parameters: map[string]interface{}{}},
		}}
	_, err := s.action.Enqueue(arg)
	c.Assert(err, jc.ErrorIsNil)

	operations, err := s.action.Operations(params.OperationQueryArgs{
		ActionNames: []string{"anotherfakeaction"},
	})
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(operations.Results, gc.HasLen, 1)
	result := operations.Results[0]
	c.Assert(result.Actions, gc.HasLen, 1)
	c.Assert(result.Actions[0].Action, gc.NotNil)
	if result.Enqueued.IsZero() {
		c.Fatal("enqueued time not set")
	}
	if result.Started.IsZero() {
		c.Fatal("started time not set")
	}
	c.Assert(result.Status, gc.Equals, "running")
	action := result.Actions[0].Action
	c.Assert(action.Name, gc.Equals, "anotherfakeaction")
	c.Assert(action.Receiver, gc.Equals, "unit-mysql-0")
	c.Assert(action.Tag, gc.Equals, "action-5")
	c.Assert(result.Actions[0].Status, gc.Equals, "pending")
}

func (s *operationSuite) TestOperationsAppFilter(c *gc.C) {
	s.setupOperations(c)
	// Set up a second operation for a different app.
	arg := params.Actions{
		Actions: []params.Action{
			{Receiver: s.mysqlUnit.Tag().String(), Name: "fakeaction", Parameters: map[string]interface{}{}},
		}}
	_, err := s.action.Enqueue(arg)
	c.Assert(err, jc.ErrorIsNil)

	operations, err := s.action.Operations(params.OperationQueryArgs{
		Applications: []string{"wordpress"},
	})
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(operations.Results, gc.HasLen, 1)
	result := operations.Results[0]

	c.Assert(result.Actions, gc.HasLen, 2)
	c.Assert(result.Actions[0].Action, gc.NotNil)
	if result.Enqueued.IsZero() {
		c.Fatal("enqueued time not set")
	}
	if result.Started.IsZero() {
		c.Fatal("started time not set")
	}
	c.Assert(result.Status, gc.Equals, "running")
	action := result.Actions[0].Action
	c.Assert(action.Name, gc.Equals, "fakeaction")
	c.Assert(action.Receiver, gc.Equals, "unit-wordpress-0")
	c.Assert(action.Tag, gc.Equals, "action-2")
	c.Assert(result.Actions[0].Status, gc.Equals, "running")
	action = result.Actions[1].Action
	c.Assert(action.Name, gc.Equals, "fakeaction")
	c.Assert(action.Receiver, gc.Equals, "unit-wordpress-0")
	c.Assert(action.Tag, gc.Equals, "action-4")
	c.Assert(result.Actions[1].Status, gc.Equals, "pending")
}

func (s *operationSuite) TestOperationsUnitFilter(c *gc.C) {
	s.setupOperations(c)
	// Set up an operation with a pending action.
	arg := params.Actions{
		Actions: []params.Action{
			{Receiver: s.wordpressUnit.Tag().String(), Name: "fakeaction", Parameters: map[string]interface{}{}},
		}}
	_, err := s.action.Enqueue(arg)
	c.Assert(err, jc.ErrorIsNil)

	operations, err := s.action.Operations(params.OperationQueryArgs{
		Units:  []string{"wordpress/0"},
		Status: []string{"pending"},
	})
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(operations.Results, gc.HasLen, 1)
	result := operations.Results[0]

	c.Assert(result.Actions, gc.HasLen, 1)
	c.Assert(result.Actions[0].Action, gc.NotNil)
	if result.Enqueued.IsZero() {
		c.Fatal("enqueued time not set")
	}
	c.Assert(result.Status, gc.Equals, "pending")
	action := result.Actions[0].Action
	c.Assert(action.Name, gc.Equals, "fakeaction")
	c.Assert(action.Receiver, gc.Equals, "unit-wordpress-0")
	c.Assert(action.Tag, gc.Equals, "action-7")
	c.Assert(result.Actions[0].Status, gc.Equals, "pending")
}

func (s *operationSuite) TestOperationsAppAndUnitFilter(c *gc.C) {
	s.setupOperations(c)
	// Set up an operation with a pending action.
	arg := params.Actions{
		Actions: []params.Action{
			{Receiver: s.wordpressUnit.Tag().String(), Name: "fakeaction", Parameters: map[string]interface{}{}},
		}}
	_, err := s.action.Enqueue(arg)
	c.Assert(err, jc.ErrorIsNil)

	operations, err := s.action.Operations(params.OperationQueryArgs{
		Applications: []string{"mysql"},
		Units:        []string{"wordpress/0"},
		Status:       []string{"running"},
	})
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(operations.Results, gc.HasLen, 1)
	c.Log(pretty.Sprint(operations.Results))
	result := operations.Results[0]

	c.Assert(result.Actions, gc.HasLen, 4)
	c.Assert(result.Actions[0].Action, gc.NotNil)
	if result.Enqueued.IsZero() {
		c.Fatal("enqueued time not set")
	}
	if result.Started.IsZero() {
		c.Fatal("started time not set")
	}

	action := result.Actions[0].Action
	c.Assert(action.Name, gc.Equals, "fakeaction")
	c.Assert(action.Receiver, gc.Equals, "unit-wordpress-0")
	c.Assert(action.Tag, gc.Equals, "action-2")
	c.Assert(result.Actions[0].Status, gc.Equals, "running")
	action = result.Actions[1].Action
	c.Assert(action.Name, gc.Equals, "fakeaction")
	c.Assert(action.Receiver, gc.Equals, "unit-mysql-0")
	c.Assert(action.Tag, gc.Equals, "action-3")
	c.Assert(result.Actions[1].Status, gc.Equals, "completed")
	action = result.Actions[2].Action
	c.Assert(action.Name, gc.Equals, "fakeaction")
	c.Assert(action.Receiver, gc.Equals, "unit-wordpress-0")
	c.Assert(action.Tag, gc.Equals, "action-4")
	c.Assert(result.Actions[2].Status, gc.Equals, "pending")
	action = result.Actions[3].Action
	c.Assert(action.Name, gc.Equals, "anotherfakeaction")
	c.Assert(action.Receiver, gc.Equals, "unit-mysql-0")
	c.Assert(action.Tag, gc.Equals, "action-5")
	c.Assert(result.Actions[3].Status, gc.Equals, "pending")
}
