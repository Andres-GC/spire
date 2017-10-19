// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/spiffe/spire/proto/server/upstreamca (interfaces: UpstreamCa)

// Package mock_upstreamca is a generated GoMock package.
package mock_upstreamca

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	plugin "github.com/spiffe/spire/proto/common/plugin"
	upstreamca "github.com/spiffe/spire/proto/server/upstreamca"
)

// MockUpstreamCa is a mock of UpstreamCa interface
type MockUpstreamCa struct {
	ctrl     *gomock.Controller
	recorder *MockUpstreamCaMockRecorder
}

// MockUpstreamCaMockRecorder is the mock recorder for MockUpstreamCa
type MockUpstreamCaMockRecorder struct {
	mock *MockUpstreamCa
}

// NewMockUpstreamCa creates a new mock instance
func NewMockUpstreamCa(ctrl *gomock.Controller) *MockUpstreamCa {
	mock := &MockUpstreamCa{ctrl: ctrl}
	mock.recorder = &MockUpstreamCaMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUpstreamCa) EXPECT() *MockUpstreamCaMockRecorder {
	return m.recorder
}

// Configure mocks base method
func (m *MockUpstreamCa) Configure(arg0 *plugin.ConfigureRequest) (*plugin.ConfigureResponse, error) {
	ret := m.ctrl.Call(m, "Configure", arg0)
	ret0, _ := ret[0].(*plugin.ConfigureResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Configure indicates an expected call of Configure
func (mr *MockUpstreamCaMockRecorder) Configure(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Configure", reflect.TypeOf((*MockUpstreamCa)(nil).Configure), arg0)
}

// GetPluginInfo mocks base method
func (m *MockUpstreamCa) GetPluginInfo(arg0 *plugin.GetPluginInfoRequest) (*plugin.GetPluginInfoResponse, error) {
	ret := m.ctrl.Call(m, "GetPluginInfo", arg0)
	ret0, _ := ret[0].(*plugin.GetPluginInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPluginInfo indicates an expected call of GetPluginInfo
func (mr *MockUpstreamCaMockRecorder) GetPluginInfo(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPluginInfo", reflect.TypeOf((*MockUpstreamCa)(nil).GetPluginInfo), arg0)
}

// SubmitCSR mocks base method
func (m *MockUpstreamCa) SubmitCSR(arg0 *upstreamca.SubmitCSRRequest) (*upstreamca.SubmitCSRResponse, error) {
	ret := m.ctrl.Call(m, "SubmitCSR", arg0)
	ret0, _ := ret[0].(*upstreamca.SubmitCSRResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitCSR indicates an expected call of SubmitCSR
func (mr *MockUpstreamCaMockRecorder) SubmitCSR(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitCSR", reflect.TypeOf((*MockUpstreamCa)(nil).SubmitCSR), arg0)
}
