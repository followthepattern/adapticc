// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cerbos/cerbos/client (interfaces: Client)

// Package mock_client is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	client "github.com/cerbos/cerbos/client"
	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// CheckResourceBatch mocks base method.
func (m *MockClient) CheckResourceBatch(arg0 context.Context, arg1 *client.Principal, arg2 *client.ResourceBatch) (*client.CheckResourceBatchResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckResourceBatch", arg0, arg1, arg2)
	ret0, _ := ret[0].(*client.CheckResourceBatchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckResourceBatch indicates an expected call of CheckResourceBatch.
func (mr *MockClientMockRecorder) CheckResourceBatch(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckResourceBatch", reflect.TypeOf((*MockClient)(nil).CheckResourceBatch), arg0, arg1, arg2)
}

// CheckResourceSet mocks base method.
func (m *MockClient) CheckResourceSet(arg0 context.Context, arg1 *client.Principal, arg2 *client.ResourceSet, arg3 ...string) (*client.CheckResourceSetResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckResourceSet", varargs...)
	ret0, _ := ret[0].(*client.CheckResourceSetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckResourceSet indicates an expected call of CheckResourceSet.
func (mr *MockClientMockRecorder) CheckResourceSet(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckResourceSet", reflect.TypeOf((*MockClient)(nil).CheckResourceSet), varargs...)
}

// CheckResources mocks base method.
func (m *MockClient) CheckResources(arg0 context.Context, arg1 *client.Principal, arg2 *client.ResourceBatch) (*client.CheckResourcesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckResources", arg0, arg1, arg2)
	ret0, _ := ret[0].(*client.CheckResourcesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckResources indicates an expected call of CheckResources.
func (mr *MockClientMockRecorder) CheckResources(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckResources", reflect.TypeOf((*MockClient)(nil).CheckResources), arg0, arg1, arg2)
}

// IsAllowed mocks base method.
func (m *MockClient) IsAllowed(arg0 context.Context, arg1 *client.Principal, arg2 *client.Resource, arg3 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAllowed", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsAllowed indicates an expected call of IsAllowed.
func (mr *MockClientMockRecorder) IsAllowed(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAllowed", reflect.TypeOf((*MockClient)(nil).IsAllowed), arg0, arg1, arg2, arg3)
}

// PlanResources mocks base method.
func (m *MockClient) PlanResources(arg0 context.Context, arg1 *client.Principal, arg2 *client.Resource, arg3 string) (*client.PlanResourcesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PlanResources", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*client.PlanResourcesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PlanResources indicates an expected call of PlanResources.
func (mr *MockClientMockRecorder) PlanResources(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlanResources", reflect.TypeOf((*MockClient)(nil).PlanResources), arg0, arg1, arg2, arg3)
}

// ServerInfo mocks base method.
func (m *MockClient) ServerInfo(arg0 context.Context) (*client.ServerInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerInfo", arg0)
	ret0, _ := ret[0].(*client.ServerInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerInfo indicates an expected call of ServerInfo.
func (mr *MockClientMockRecorder) ServerInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerInfo", reflect.TypeOf((*MockClient)(nil).ServerInfo), arg0)
}

// With mocks base method.
func (m *MockClient) With(arg0 ...client.RequestOpt) client.Client {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "With", varargs...)
	ret0, _ := ret[0].(client.Client)
	return ret0
}

// With indicates an expected call of With.
func (mr *MockClientMockRecorder) With(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "With", reflect.TypeOf((*MockClient)(nil).With), arg0...)
}

// WithPrincipal mocks base method.
func (m *MockClient) WithPrincipal(arg0 *client.Principal) client.PrincipalContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithPrincipal", arg0)
	ret0, _ := ret[0].(client.PrincipalContext)
	return ret0
}

// WithPrincipal indicates an expected call of WithPrincipal.
func (mr *MockClientMockRecorder) WithPrincipal(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithPrincipal", reflect.TypeOf((*MockClient)(nil).WithPrincipal), arg0)
}
