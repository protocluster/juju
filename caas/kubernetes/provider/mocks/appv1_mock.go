// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/apps/v1 (interfaces: AppsV1Interface,DeploymentInterface,StatefulSetInterface,DaemonSetInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/apps/v1"
	v10 "k8s.io/api/autoscaling/v1"
	v11 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v12 "k8s.io/client-go/kubernetes/typed/apps/v1"
	rest "k8s.io/client-go/rest"
	reflect "reflect"
)

// MockAppsV1Interface is a mock of AppsV1Interface interface
type MockAppsV1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockAppsV1InterfaceMockRecorder
}

// MockAppsV1InterfaceMockRecorder is the mock recorder for MockAppsV1Interface
type MockAppsV1InterfaceMockRecorder struct {
	mock *MockAppsV1Interface
}

// NewMockAppsV1Interface creates a new mock instance
func NewMockAppsV1Interface(ctrl *gomock.Controller) *MockAppsV1Interface {
	mock := &MockAppsV1Interface{ctrl: ctrl}
	mock.recorder = &MockAppsV1InterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAppsV1Interface) EXPECT() *MockAppsV1InterfaceMockRecorder {
	return m.recorder
}

// ControllerRevisions mocks base method
func (m *MockAppsV1Interface) ControllerRevisions(arg0 string) v12.ControllerRevisionInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerRevisions", arg0)
	ret0, _ := ret[0].(v12.ControllerRevisionInterface)
	return ret0
}

// ControllerRevisions indicates an expected call of ControllerRevisions
func (mr *MockAppsV1InterfaceMockRecorder) ControllerRevisions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerRevisions", reflect.TypeOf((*MockAppsV1Interface)(nil).ControllerRevisions), arg0)
}

// DaemonSets mocks base method
func (m *MockAppsV1Interface) DaemonSets(arg0 string) v12.DaemonSetInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DaemonSets", arg0)
	ret0, _ := ret[0].(v12.DaemonSetInterface)
	return ret0
}

// DaemonSets indicates an expected call of DaemonSets
func (mr *MockAppsV1InterfaceMockRecorder) DaemonSets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DaemonSets", reflect.TypeOf((*MockAppsV1Interface)(nil).DaemonSets), arg0)
}

// Deployments mocks base method
func (m *MockAppsV1Interface) Deployments(arg0 string) v12.DeploymentInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deployments", arg0)
	ret0, _ := ret[0].(v12.DeploymentInterface)
	return ret0
}

// Deployments indicates an expected call of Deployments
func (mr *MockAppsV1InterfaceMockRecorder) Deployments(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deployments", reflect.TypeOf((*MockAppsV1Interface)(nil).Deployments), arg0)
}

// RESTClient mocks base method
func (m *MockAppsV1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTClient")
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

// RESTClient indicates an expected call of RESTClient
func (mr *MockAppsV1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTClient", reflect.TypeOf((*MockAppsV1Interface)(nil).RESTClient))
}

// ReplicaSets mocks base method
func (m *MockAppsV1Interface) ReplicaSets(arg0 string) v12.ReplicaSetInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplicaSets", arg0)
	ret0, _ := ret[0].(v12.ReplicaSetInterface)
	return ret0
}

// ReplicaSets indicates an expected call of ReplicaSets
func (mr *MockAppsV1InterfaceMockRecorder) ReplicaSets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplicaSets", reflect.TypeOf((*MockAppsV1Interface)(nil).ReplicaSets), arg0)
}

// StatefulSets mocks base method
func (m *MockAppsV1Interface) StatefulSets(arg0 string) v12.StatefulSetInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatefulSets", arg0)
	ret0, _ := ret[0].(v12.StatefulSetInterface)
	return ret0
}

// StatefulSets indicates an expected call of StatefulSets
func (mr *MockAppsV1InterfaceMockRecorder) StatefulSets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatefulSets", reflect.TypeOf((*MockAppsV1Interface)(nil).StatefulSets), arg0)
}

// MockDeploymentInterface is a mock of DeploymentInterface interface
type MockDeploymentInterface struct {
	ctrl     *gomock.Controller
	recorder *MockDeploymentInterfaceMockRecorder
}

// MockDeploymentInterfaceMockRecorder is the mock recorder for MockDeploymentInterface
type MockDeploymentInterfaceMockRecorder struct {
	mock *MockDeploymentInterface
}

// NewMockDeploymentInterface creates a new mock instance
func NewMockDeploymentInterface(ctrl *gomock.Controller) *MockDeploymentInterface {
	mock := &MockDeploymentInterface{ctrl: ctrl}
	mock.recorder = &MockDeploymentInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDeploymentInterface) EXPECT() *MockDeploymentInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockDeploymentInterface) Create(arg0 *v1.Deployment) (*v1.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*v1.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockDeploymentInterfaceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDeploymentInterface)(nil).Create), arg0)
}

// Delete mocks base method
func (m *MockDeploymentInterface) Delete(arg0 string, arg1 *v11.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockDeploymentInterfaceMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDeploymentInterface)(nil).Delete), arg0, arg1)
}

// DeleteCollection mocks base method
func (m *MockDeploymentInterface) DeleteCollection(arg0 *v11.DeleteOptions, arg1 v11.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection
func (mr *MockDeploymentInterfaceMockRecorder) DeleteCollection(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockDeploymentInterface)(nil).DeleteCollection), arg0, arg1)
}

// Get mocks base method
func (m *MockDeploymentInterface) Get(arg0 string, arg1 v11.GetOptions) (*v1.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*v1.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockDeploymentInterfaceMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDeploymentInterface)(nil).Get), arg0, arg1)
}

// GetScale mocks base method
func (m *MockDeploymentInterface) GetScale(arg0 string, arg1 v11.GetOptions) (*v10.Scale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScale", arg0, arg1)
	ret0, _ := ret[0].(*v10.Scale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScale indicates an expected call of GetScale
func (mr *MockDeploymentInterfaceMockRecorder) GetScale(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScale", reflect.TypeOf((*MockDeploymentInterface)(nil).GetScale), arg0, arg1)
}

// List mocks base method
func (m *MockDeploymentInterface) List(arg0 v11.ListOptions) (*v1.DeploymentList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(*v1.DeploymentList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockDeploymentInterfaceMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockDeploymentInterface)(nil).List), arg0)
}

// Patch mocks base method
func (m *MockDeploymentInterface) Patch(arg0 string, arg1 types.PatchType, arg2 []byte, arg3 ...string) (*v1.Deployment, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch
func (mr *MockDeploymentInterfaceMockRecorder) Patch(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockDeploymentInterface)(nil).Patch), varargs...)
}

// Update mocks base method
func (m *MockDeploymentInterface) Update(arg0 *v1.Deployment) (*v1.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(*v1.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockDeploymentInterfaceMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDeploymentInterface)(nil).Update), arg0)
}

// UpdateScale mocks base method
func (m *MockDeploymentInterface) UpdateScale(arg0 string, arg1 *v10.Scale) (*v10.Scale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateScale", arg0, arg1)
	ret0, _ := ret[0].(*v10.Scale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateScale indicates an expected call of UpdateScale
func (mr *MockDeploymentInterfaceMockRecorder) UpdateScale(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateScale", reflect.TypeOf((*MockDeploymentInterface)(nil).UpdateScale), arg0, arg1)
}

// UpdateStatus mocks base method
func (m *MockDeploymentInterface) UpdateStatus(arg0 *v1.Deployment) (*v1.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", arg0)
	ret0, _ := ret[0].(*v1.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStatus indicates an expected call of UpdateStatus
func (mr *MockDeploymentInterfaceMockRecorder) UpdateStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockDeploymentInterface)(nil).UpdateStatus), arg0)
}

// Watch mocks base method
func (m *MockDeploymentInterface) Watch(arg0 v11.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (mr *MockDeploymentInterfaceMockRecorder) Watch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockDeploymentInterface)(nil).Watch), arg0)
}

// MockStatefulSetInterface is a mock of StatefulSetInterface interface
type MockStatefulSetInterface struct {
	ctrl     *gomock.Controller
	recorder *MockStatefulSetInterfaceMockRecorder
}

// MockStatefulSetInterfaceMockRecorder is the mock recorder for MockStatefulSetInterface
type MockStatefulSetInterfaceMockRecorder struct {
	mock *MockStatefulSetInterface
}

// NewMockStatefulSetInterface creates a new mock instance
func NewMockStatefulSetInterface(ctrl *gomock.Controller) *MockStatefulSetInterface {
	mock := &MockStatefulSetInterface{ctrl: ctrl}
	mock.recorder = &MockStatefulSetInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStatefulSetInterface) EXPECT() *MockStatefulSetInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockStatefulSetInterface) Create(arg0 *v1.StatefulSet) (*v1.StatefulSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*v1.StatefulSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockStatefulSetInterfaceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStatefulSetInterface)(nil).Create), arg0)
}

// Delete mocks base method
func (m *MockStatefulSetInterface) Delete(arg0 string, arg1 *v11.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockStatefulSetInterfaceMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStatefulSetInterface)(nil).Delete), arg0, arg1)
}

// DeleteCollection mocks base method
func (m *MockStatefulSetInterface) DeleteCollection(arg0 *v11.DeleteOptions, arg1 v11.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection
func (mr *MockStatefulSetInterfaceMockRecorder) DeleteCollection(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockStatefulSetInterface)(nil).DeleteCollection), arg0, arg1)
}

// Get mocks base method
func (m *MockStatefulSetInterface) Get(arg0 string, arg1 v11.GetOptions) (*v1.StatefulSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*v1.StatefulSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockStatefulSetInterfaceMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStatefulSetInterface)(nil).Get), arg0, arg1)
}

// GetScale mocks base method
func (m *MockStatefulSetInterface) GetScale(arg0 string, arg1 v11.GetOptions) (*v10.Scale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScale", arg0, arg1)
	ret0, _ := ret[0].(*v10.Scale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScale indicates an expected call of GetScale
func (mr *MockStatefulSetInterfaceMockRecorder) GetScale(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScale", reflect.TypeOf((*MockStatefulSetInterface)(nil).GetScale), arg0, arg1)
}

// List mocks base method
func (m *MockStatefulSetInterface) List(arg0 v11.ListOptions) (*v1.StatefulSetList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(*v1.StatefulSetList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockStatefulSetInterfaceMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockStatefulSetInterface)(nil).List), arg0)
}

// Patch mocks base method
func (m *MockStatefulSetInterface) Patch(arg0 string, arg1 types.PatchType, arg2 []byte, arg3 ...string) (*v1.StatefulSet, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1.StatefulSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch
func (mr *MockStatefulSetInterfaceMockRecorder) Patch(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockStatefulSetInterface)(nil).Patch), varargs...)
}

// Update mocks base method
func (m *MockStatefulSetInterface) Update(arg0 *v1.StatefulSet) (*v1.StatefulSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(*v1.StatefulSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockStatefulSetInterfaceMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockStatefulSetInterface)(nil).Update), arg0)
}

// UpdateScale mocks base method
func (m *MockStatefulSetInterface) UpdateScale(arg0 string, arg1 *v10.Scale) (*v10.Scale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateScale", arg0, arg1)
	ret0, _ := ret[0].(*v10.Scale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateScale indicates an expected call of UpdateScale
func (mr *MockStatefulSetInterfaceMockRecorder) UpdateScale(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateScale", reflect.TypeOf((*MockStatefulSetInterface)(nil).UpdateScale), arg0, arg1)
}

// UpdateStatus mocks base method
func (m *MockStatefulSetInterface) UpdateStatus(arg0 *v1.StatefulSet) (*v1.StatefulSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", arg0)
	ret0, _ := ret[0].(*v1.StatefulSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStatus indicates an expected call of UpdateStatus
func (mr *MockStatefulSetInterfaceMockRecorder) UpdateStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockStatefulSetInterface)(nil).UpdateStatus), arg0)
}

// Watch mocks base method
func (m *MockStatefulSetInterface) Watch(arg0 v11.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (mr *MockStatefulSetInterfaceMockRecorder) Watch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockStatefulSetInterface)(nil).Watch), arg0)
}

// MockDaemonSetInterface is a mock of DaemonSetInterface interface
type MockDaemonSetInterface struct {
	ctrl     *gomock.Controller
	recorder *MockDaemonSetInterfaceMockRecorder
}

// MockDaemonSetInterfaceMockRecorder is the mock recorder for MockDaemonSetInterface
type MockDaemonSetInterfaceMockRecorder struct {
	mock *MockDaemonSetInterface
}

// NewMockDaemonSetInterface creates a new mock instance
func NewMockDaemonSetInterface(ctrl *gomock.Controller) *MockDaemonSetInterface {
	mock := &MockDaemonSetInterface{ctrl: ctrl}
	mock.recorder = &MockDaemonSetInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDaemonSetInterface) EXPECT() *MockDaemonSetInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockDaemonSetInterface) Create(arg0 *v1.DaemonSet) (*v1.DaemonSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*v1.DaemonSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockDaemonSetInterfaceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDaemonSetInterface)(nil).Create), arg0)
}

// Delete mocks base method
func (m *MockDaemonSetInterface) Delete(arg0 string, arg1 *v11.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockDaemonSetInterfaceMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDaemonSetInterface)(nil).Delete), arg0, arg1)
}

// DeleteCollection mocks base method
func (m *MockDaemonSetInterface) DeleteCollection(arg0 *v11.DeleteOptions, arg1 v11.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection
func (mr *MockDaemonSetInterfaceMockRecorder) DeleteCollection(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockDaemonSetInterface)(nil).DeleteCollection), arg0, arg1)
}

// Get mocks base method
func (m *MockDaemonSetInterface) Get(arg0 string, arg1 v11.GetOptions) (*v1.DaemonSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*v1.DaemonSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockDaemonSetInterfaceMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDaemonSetInterface)(nil).Get), arg0, arg1)
}

// List mocks base method
func (m *MockDaemonSetInterface) List(arg0 v11.ListOptions) (*v1.DaemonSetList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(*v1.DaemonSetList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockDaemonSetInterfaceMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockDaemonSetInterface)(nil).List), arg0)
}

// Patch mocks base method
func (m *MockDaemonSetInterface) Patch(arg0 string, arg1 types.PatchType, arg2 []byte, arg3 ...string) (*v1.DaemonSet, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1.DaemonSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch
func (mr *MockDaemonSetInterfaceMockRecorder) Patch(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockDaemonSetInterface)(nil).Patch), varargs...)
}

// Update mocks base method
func (m *MockDaemonSetInterface) Update(arg0 *v1.DaemonSet) (*v1.DaemonSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(*v1.DaemonSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockDaemonSetInterfaceMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDaemonSetInterface)(nil).Update), arg0)
}

// UpdateStatus mocks base method
func (m *MockDaemonSetInterface) UpdateStatus(arg0 *v1.DaemonSet) (*v1.DaemonSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", arg0)
	ret0, _ := ret[0].(*v1.DaemonSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStatus indicates an expected call of UpdateStatus
func (mr *MockDaemonSetInterfaceMockRecorder) UpdateStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockDaemonSetInterface)(nil).UpdateStatus), arg0)
}

// Watch mocks base method
func (m *MockDaemonSetInterface) Watch(arg0 v11.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (mr *MockDaemonSetInterfaceMockRecorder) Watch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockDaemonSetInterface)(nil).Watch), arg0)
}
