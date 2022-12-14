// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"
	dto "team-task/internal/dto"

	gomock "github.com/golang/mock/gomock"
)

// MockUserGrade is a mock of UserGrade interface.
type MockUserGrade struct {
	ctrl     *gomock.Controller
	recorder *MockUserGradeMockRecorder
}

// MockUserGradeMockRecorder is the mock recorder for MockUserGrade.
type MockUserGradeMockRecorder struct {
	mock *MockUserGrade
}

// NewMockUserGrade creates a new mock instance.
func NewMockUserGrade(ctrl *gomock.Controller) *MockUserGrade {
	mock := &MockUserGrade{ctrl: ctrl}
	mock.recorder = &MockUserGradeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserGrade) EXPECT() *MockUserGradeMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockUserGrade) Get(userID string) (dto.UserGrade, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", userID)
	ret0, _ := ret[0].(dto.UserGrade)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockUserGradeMockRecorder) Get(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserGrade)(nil).Get), userID)
}

// Set mocks base method.
func (m *MockUserGrade) Set(userGrade dto.UserGrade) (dto.UserGrade, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", userGrade)
	ret0, _ := ret[0].(dto.UserGrade)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Set indicates an expected call of Set.
func (mr *MockUserGradeMockRecorder) Set(userGrade interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockUserGrade)(nil).Set), userGrade)
}

// MockBackupper is a mock of Backupper interface.
type MockBackupper struct {
	ctrl     *gomock.Controller
	recorder *MockBackupperMockRecorder
}

// MockBackupperMockRecorder is the mock recorder for MockBackupper.
type MockBackupperMockRecorder struct {
	mock *MockBackupper
}

// NewMockBackupper creates a new mock instance.
func NewMockBackupper(ctrl *gomock.Controller) *MockBackupper {
	mock := &MockBackupper{ctrl: ctrl}
	mock.recorder = &MockBackupperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackupper) EXPECT() *MockBackupperMockRecorder {
	return m.recorder
}

// Backup mocks base method.
func (m *MockBackupper) Backup() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Backup")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Backup indicates an expected call of Backup.
func (mr *MockBackupperMockRecorder) Backup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Backup", reflect.TypeOf((*MockBackupper)(nil).Backup))
}
