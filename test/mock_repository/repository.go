// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/repo.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	proto "github.com/cyneruxyz/address-book/gen/proto"
	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateAddressField mocks base method.
func (m *MockRepository) CreateAddressField(field *proto.AddressField) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAddressField", field)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAddressField indicates an expected call of CreateAddressField.
func (mr *MockRepositoryMockRecorder) CreateAddressField(field interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAddressField", reflect.TypeOf((*MockRepository)(nil).CreateAddressField), field)
}

// DeleteAddressField mocks base method.
func (m *MockRepository) DeleteAddressField(phone *proto.Phone) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAddressField", phone)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAddressField indicates an expected call of DeleteAddressField.
func (mr *MockRepositoryMockRecorder) DeleteAddressField(phone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAddressField", reflect.TypeOf((*MockRepository)(nil).DeleteAddressField), phone)
}

// GetAddressFields mocks base method.
func (m *MockRepository) GetAddressFields(param string) ([]*proto.AddressField, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddressFields", param)
	ret0, _ := ret[0].([]*proto.AddressField)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAddressFields indicates an expected call of GetAddressFields.
func (mr *MockRepositoryMockRecorder) GetAddressFields(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddressFields", reflect.TypeOf((*MockRepository)(nil).GetAddressFields), param)
}

// UpdateAddressField mocks base method.
func (m *MockRepository) UpdateAddressField(phone *proto.Phone, replace *proto.AddressField) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAddressField", phone, replace)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAddressField indicates an expected call of UpdateAddressField.
func (mr *MockRepositoryMockRecorder) UpdateAddressField(phone, replace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAddressField", reflect.TypeOf((*MockRepository)(nil).UpdateAddressField), phone, replace)
}