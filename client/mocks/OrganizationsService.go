// Code generated by MockGen. DO NOT EDIT.
// Source: organizations.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gocancel "github.com/gocancel/gocancel-go"
	gomock "github.com/golang/mock/gomock"
)

// MockOrganizationsService is a mock of OrganizationsService interface.
type MockOrganizationsService struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationsServiceMockRecorder
}

// MockOrganizationsServiceMockRecorder is the mock recorder for MockOrganizationsService.
type MockOrganizationsServiceMockRecorder struct {
	mock *MockOrganizationsService
}

// NewMockOrganizationsService creates a new mock instance.
func NewMockOrganizationsService(ctrl *gomock.Controller) *MockOrganizationsService {
	mock := &MockOrganizationsService{ctrl: ctrl}
	mock.recorder = &MockOrganizationsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationsService) EXPECT() *MockOrganizationsServiceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockOrganizationsService) Get(organizationID string) (*gocancel.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", organizationID)
	ret0, _ := ret[0].(*gocancel.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockOrganizationsServiceMockRecorder) Get(organizationID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockOrganizationsService)(nil).Get), organizationID)
}

// GetProduct mocks base method.
func (m *MockOrganizationsService) GetProduct(organizationID, productID string) (*gocancel.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProduct", organizationID, productID)
	ret0, _ := ret[0].(*gocancel.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProduct indicates an expected call of GetProduct.
func (mr *MockOrganizationsServiceMockRecorder) GetProduct(organizationID, productID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProduct", reflect.TypeOf((*MockOrganizationsService)(nil).GetProduct), organizationID, productID)
}

// List mocks base method.
func (m *MockOrganizationsService) List(opts *gocancel.OrganizationsListOptions) ([]*gocancel.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", opts)
	ret0, _ := ret[0].([]*gocancel.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockOrganizationsServiceMockRecorder) List(opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockOrganizationsService)(nil).List), opts)
}

// ListProducts mocks base method.
func (m *MockOrganizationsService) ListProducts(organizationID string, opts *gocancel.OrganizationProductsListOptions) ([]*gocancel.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProducts", organizationID, opts)
	ret0, _ := ret[0].([]*gocancel.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProducts indicates an expected call of ListProducts.
func (mr *MockOrganizationsServiceMockRecorder) ListProducts(organizationID, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProducts", reflect.TypeOf((*MockOrganizationsService)(nil).ListProducts), organizationID, opts)
}
