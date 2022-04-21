// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rudderlabs/rudder-server/config/backend-config (interfaces: BackendConfig)

// Package mock_backendconfig is a generated GoMock package.
package mock_backendconfig

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	backendconfig "github.com/rudderlabs/rudder-server/config/backend-config"
	pubsub "github.com/rudderlabs/rudder-server/utils/pubsub"
)

// MockBackendConfig is a mock of BackendConfig interface.
type MockBackendConfig struct {
	ctrl     *gomock.Controller
	recorder *MockBackendConfigMockRecorder
}

// MockBackendConfigMockRecorder is the mock recorder for MockBackendConfig.
type MockBackendConfigMockRecorder struct {
	mock *MockBackendConfig
}

// NewMockBackendConfig creates a new mock instance.
func NewMockBackendConfig(ctrl *gomock.Controller) *MockBackendConfig {
	mock := &MockBackendConfig{ctrl: ctrl}
	mock.recorder = &MockBackendConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackendConfig) EXPECT() *MockBackendConfigMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockBackendConfig) Get(arg0 string) (backendconfig.ConfigT, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(backendconfig.ConfigT)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockBackendConfigMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockBackendConfig)(nil).Get), arg0)
}

// GetWorkspaceIDForSourceID mocks base method.
func (m *MockBackendConfig) GetWorkspaceIDForSourceID(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkspaceIDForSourceID", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetWorkspaceIDForSourceID indicates an expected call of GetWorkspaceIDForSourceID.
func (mr *MockBackendConfigMockRecorder) GetWorkspaceIDForSourceID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkspaceIDForSourceID", reflect.TypeOf((*MockBackendConfig)(nil).GetWorkspaceIDForSourceID), arg0)
}

// GetWorkspaceIDForWriteKey mocks base method.
func (m *MockBackendConfig) GetWorkspaceIDForWriteKey(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkspaceIDForWriteKey", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetWorkspaceIDForWriteKey indicates an expected call of GetWorkspaceIDForWriteKey.
func (mr *MockBackendConfigMockRecorder) GetWorkspaceIDForWriteKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkspaceIDForWriteKey", reflect.TypeOf((*MockBackendConfig)(nil).GetWorkspaceIDForWriteKey), arg0)
}

// GetWorkspaceLibrariesForWorkspaceID mocks base method.
func (m *MockBackendConfig) GetWorkspaceLibrariesForWorkspaceID(arg0 string) backendconfig.LibrariesT {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkspaceLibrariesForWorkspaceID", arg0)
	ret0, _ := ret[0].(backendconfig.LibrariesT)
	return ret0
}

// GetWorkspaceLibrariesForWorkspaceID indicates an expected call of GetWorkspaceLibrariesForWorkspaceID.
func (mr *MockBackendConfigMockRecorder) GetWorkspaceLibrariesForWorkspaceID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkspaceLibrariesForWorkspaceID", reflect.TypeOf((*MockBackendConfig)(nil).GetWorkspaceLibrariesForWorkspaceID), arg0)
}

// SetUp mocks base method.
func (m *MockBackendConfig) SetUp() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetUp")
}

// SetUp indicates an expected call of SetUp.
func (mr *MockBackendConfigMockRecorder) SetUp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUp", reflect.TypeOf((*MockBackendConfig)(nil).SetUp))
}

// StartWithIDs mocks base method.
func (m *MockBackendConfig) StartWithIDs(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartWithIDs", arg0)
}

// StartWithIDs indicates an expected call of StartWithIDs.
func (mr *MockBackendConfigMockRecorder) StartWithIDs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartWithIDs", reflect.TypeOf((*MockBackendConfig)(nil).StartWithIDs), arg0)
}

// Stop mocks base method.
func (m *MockBackendConfig) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockBackendConfigMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockBackendConfig)(nil).Stop))
}

// Subscribe mocks base method.
func (m *MockBackendConfig) Subscribe(arg0 chan pubsub.DataEvent, arg1 backendconfig.Topic) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Subscribe", arg0, arg1)
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockBackendConfigMockRecorder) Subscribe(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockBackendConfig)(nil).Subscribe), arg0, arg1)
}

// WaitForConfig mocks base method.
func (m *MockBackendConfig) WaitForConfig(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForConfig", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForConfig indicates an expected call of WaitForConfig.
func (mr *MockBackendConfigMockRecorder) WaitForConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForConfig", reflect.TypeOf((*MockBackendConfig)(nil).WaitForConfig), arg0)
}
