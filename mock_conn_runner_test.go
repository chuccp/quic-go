// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/quic-go/quic-go (interfaces: ConnRunner)

// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	protocol "github.com/quic-go/quic-go/internal/protocol"
	gomock "go.uber.org/mock/gomock"
)

// MockConnRunner is a mock of ConnRunner interface.
type MockConnRunner struct {
	ctrl     *gomock.Controller
	recorder *MockConnRunnerMockRecorder
}

// MockConnRunnerMockRecorder is the mock recorder for MockConnRunner.
type MockConnRunnerMockRecorder struct {
	mock *MockConnRunner
}

// NewMockConnRunner creates a new mock instance.
func NewMockConnRunner(ctrl *gomock.Controller) *MockConnRunner {
	mock := &MockConnRunner{ctrl: ctrl}
	mock.recorder = &MockConnRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnRunner) EXPECT() *MockConnRunnerMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockConnRunner) Add(arg0 protocol.ConnectionID, arg1 packetHandler) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockConnRunnerMockRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockConnRunner)(nil).Add), arg0, arg1)
}

// AddResetToken mocks base method.
func (m *MockConnRunner) AddResetToken(arg0 protocol.StatelessResetToken, arg1 packetHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddResetToken", arg0, arg1)
}

// AddResetToken indicates an expected call of AddResetToken.
func (mr *MockConnRunnerMockRecorder) AddResetToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddResetToken", reflect.TypeOf((*MockConnRunner)(nil).AddResetToken), arg0, arg1)
}

// GetStatelessResetToken mocks base method.
func (m *MockConnRunner) GetStatelessResetToken(arg0 protocol.ConnectionID) protocol.StatelessResetToken {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatelessResetToken", arg0)
	ret0, _ := ret[0].(protocol.StatelessResetToken)
	return ret0
}

// GetStatelessResetToken indicates an expected call of GetStatelessResetToken.
func (mr *MockConnRunnerMockRecorder) GetStatelessResetToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatelessResetToken", reflect.TypeOf((*MockConnRunner)(nil).GetStatelessResetToken), arg0)
}

// Remove mocks base method.
func (m *MockConnRunner) Remove(arg0 protocol.ConnectionID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Remove", arg0)
}

// Remove indicates an expected call of Remove.
func (mr *MockConnRunnerMockRecorder) Remove(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockConnRunner)(nil).Remove), arg0)
}

// RemoveResetToken mocks base method.
func (m *MockConnRunner) RemoveResetToken(arg0 protocol.StatelessResetToken) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveResetToken", arg0)
}

// RemoveResetToken indicates an expected call of RemoveResetToken.
func (mr *MockConnRunnerMockRecorder) RemoveResetToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveResetToken", reflect.TypeOf((*MockConnRunner)(nil).RemoveResetToken), arg0)
}

// ReplaceWithClosed mocks base method.
func (m *MockConnRunner) ReplaceWithClosed(arg0 []protocol.ConnectionID, arg1 protocol.Perspective, arg2 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReplaceWithClosed", arg0, arg1, arg2)
}

// ReplaceWithClosed indicates an expected call of ReplaceWithClosed.
func (mr *MockConnRunnerMockRecorder) ReplaceWithClosed(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplaceWithClosed", reflect.TypeOf((*MockConnRunner)(nil).ReplaceWithClosed), arg0, arg1, arg2)
}

// Retire mocks base method.
func (m *MockConnRunner) Retire(arg0 protocol.ConnectionID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Retire", arg0)
}

// Retire indicates an expected call of Retire.
func (mr *MockConnRunnerMockRecorder) Retire(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Retire", reflect.TypeOf((*MockConnRunner)(nil).Retire), arg0)
}
