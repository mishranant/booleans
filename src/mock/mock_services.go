// Code generated by MockGen. DO NOT EDIT.
// Source: services.go

// Package mock_services is a generated GoMock package.
package mock

import (
	services "booleans/services"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepo is a mock of Repo interface
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// AddToDB mocks base method
func (m *MockRepo) AddToDB(name string, value bool) services.Boolean {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToDB", name, value)
	ret0, _ := ret[0].(services.Boolean)
	return ret0
}

// AddToDB indicates an expected call of AddToDB
func (mr *MockRepoMockRecorder) AddToDB(name, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToDB", reflect.TypeOf((*MockRepo)(nil).AddToDB), name, value)
}

// GetFromDB mocks base method
func (m *MockRepo) GetFromDB(uuid string) (services.Boolean, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFromDB", uuid)
	ret0, _ := ret[0].(services.Boolean)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFromDB indicates an expected call of GetFromDB
func (mr *MockRepoMockRecorder) GetFromDB(uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFromDB", reflect.TypeOf((*MockRepo)(nil).GetFromDB), uuid)
}

// UpdateInDB mocks base method
func (m *MockRepo) UpdateInDB(name string, value bool, uuid string) (services.Boolean, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInDB", name, value, uuid)
	ret0, _ := ret[0].(services.Boolean)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateInDB indicates an expected call of UpdateInDB
func (mr *MockRepoMockRecorder) UpdateInDB(name, value, uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInDB", reflect.TypeOf((*MockRepo)(nil).UpdateInDB), name, value, uuid)
}

// DeleteFromDB mocks base method
func (m *MockRepo) DeleteFromDB(uuid string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFromDB", uuid)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFromDB indicates an expected call of DeleteFromDB
func (mr *MockRepoMockRecorder) DeleteFromDB(uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFromDB", reflect.TypeOf((*MockRepo)(nil).DeleteFromDB), uuid)
}