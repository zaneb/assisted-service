// Code generated by MockGen. DO NOT EDIT.
// Source: installcfg.go

// Package installcfg is a generated GoMock package.
package installcfg

import (
	gomock "github.com/golang/mock/gomock"
	common "github.com/openshift/assisted-service/internal/common"
	reflect "reflect"
)

// MockInstallConfigBuilder is a mock of InstallConfigBuilder interface
type MockInstallConfigBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockInstallConfigBuilderMockRecorder
}

// MockInstallConfigBuilderMockRecorder is the mock recorder for MockInstallConfigBuilder
type MockInstallConfigBuilderMockRecorder struct {
	mock *MockInstallConfigBuilder
}

// NewMockInstallConfigBuilder creates a new mock instance
func NewMockInstallConfigBuilder(ctrl *gomock.Controller) *MockInstallConfigBuilder {
	mock := &MockInstallConfigBuilder{ctrl: ctrl}
	mock.recorder = &MockInstallConfigBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInstallConfigBuilder) EXPECT() *MockInstallConfigBuilderMockRecorder {
	return m.recorder
}

// GetInstallConfig mocks base method
func (m *MockInstallConfigBuilder) GetInstallConfig(cluster *common.Cluster, addRhCa bool, ca string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstallConfig", cluster, addRhCa, ca)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstallConfig indicates an expected call of GetInstallConfig
func (mr *MockInstallConfigBuilderMockRecorder) GetInstallConfig(cluster, addRhCa, ca interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstallConfig", reflect.TypeOf((*MockInstallConfigBuilder)(nil).GetInstallConfig), cluster, addRhCa, ca)
}

// ValidateInstallConfigPatch mocks base method
func (m *MockInstallConfigBuilder) ValidateInstallConfigPatch(cluster *common.Cluster, patch string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateInstallConfigPatch", cluster, patch)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateInstallConfigPatch indicates an expected call of ValidateInstallConfigPatch
func (mr *MockInstallConfigBuilderMockRecorder) ValidateInstallConfigPatch(cluster, patch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateInstallConfigPatch", reflect.TypeOf((*MockInstallConfigBuilder)(nil).ValidateInstallConfigPatch), cluster, patch)
}
