// Code generated by MockGen. DO NOT EDIT.
// Source: k8scommerce/internal/repos (interfaces: OthersBought)

// Package mock_repos is a generated GoMock package.
package mock_repos

import (
	models "k8scommerce/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOthersBought is a mock of OthersBought interface.
type MockOthersBought struct {
	ctrl     *gomock.Controller
	recorder *MockOthersBoughtMockRecorder
}

// MockOthersBoughtMockRecorder is the mock recorder for MockOthersBought.
type MockOthersBoughtMockRecorder struct {
	mock *MockOthersBought
}

// NewMockOthersBought creates a new mock instance.
func NewMockOthersBought(ctrl *gomock.Controller) *MockOthersBought {
	mock := &MockOthersBought{ctrl: ctrl}
	mock.recorder = &MockOthersBoughtMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOthersBought) EXPECT() *MockOthersBoughtMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockOthersBought) Create(arg0 *models.Product) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockOthersBoughtMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOthersBought)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockOthersBought) Delete(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockOthersBoughtMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOthersBought)(nil).Delete), arg0)
}

// Deleted mocks base method.
func (m *MockOthersBought) Deleted() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deleted")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Deleted indicates an expected call of Deleted.
func (mr *MockOthersBoughtMockRecorder) Deleted() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deleted", reflect.TypeOf((*MockOthersBought)(nil).Deleted))
}

// Exists mocks base method.
func (m *MockOthersBought) Exists() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Exists indicates an expected call of Exists.
func (mr *MockOthersBoughtMockRecorder) Exists() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockOthersBought)(nil).Exists))
}

// Save mocks base method.
func (m *MockOthersBought) Save() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save")
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockOthersBoughtMockRecorder) Save() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockOthersBought)(nil).Save))
}

// Update mocks base method.
func (m *MockOthersBought) Update(arg0 *models.Product) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockOthersBoughtMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOthersBought)(nil).Update), arg0)
}

// Upsert mocks base method.
func (m *MockOthersBought) Upsert() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert")
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockOthersBoughtMockRecorder) Upsert() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockOthersBought)(nil).Upsert))
}
