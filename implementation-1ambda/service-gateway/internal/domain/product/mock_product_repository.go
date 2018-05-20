// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/product/product_repository.go

// Package product is a generated GoMock package.
package product

import (
	exception "github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/exception"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// AddCategory mocks base method
func (m *MockRepository) AddCategory(record *Category) (*Category, exception.Exception) {
	ret := m.ctrl.Call(m, "AddCategory", record)
	ret0, _ := ret[0].(*Category)
	ret1, _ := ret[1].(exception.Exception)
	return ret0, ret1
}

// AddCategory indicates an expected call of AddCategory
func (mr *MockRepositoryMockRecorder) AddCategory(record interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCategory", reflect.TypeOf((*MockRepository)(nil).AddCategory), record)
}

// FindCategoryById mocks base method
func (m *MockRepository) FindCategoryById(id uint) (*Category, exception.Exception) {
	ret := m.ctrl.Call(m, "FindCategoryById", id)
	ret0, _ := ret[0].(*Category)
	ret1, _ := ret[1].(exception.Exception)
	return ret0, ret1
}

// FindCategoryById indicates an expected call of FindCategoryById
func (mr *MockRepositoryMockRecorder) FindCategoryById(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCategoryById", reflect.TypeOf((*MockRepository)(nil).FindCategoryById), id)
}

// AddImage mocks base method
func (m *MockRepository) AddImage(record *Image) (*Image, exception.Exception) {
	ret := m.ctrl.Call(m, "AddImage", record)
	ret0, _ := ret[0].(*Image)
	ret1, _ := ret[1].(exception.Exception)
	return ret0, ret1
}

// AddImage indicates an expected call of AddImage
func (mr *MockRepositoryMockRecorder) AddImage(record interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddImage", reflect.TypeOf((*MockRepository)(nil).AddImage), record)
}

// FindImageById mocks base method
func (m *MockRepository) FindImageById(id uint) (*Image, exception.Exception) {
	ret := m.ctrl.Call(m, "FindImageById", id)
	ret0, _ := ret[0].(*Image)
	ret1, _ := ret[1].(exception.Exception)
	return ret0, ret1
}

// FindImageById indicates an expected call of FindImageById
func (mr *MockRepositoryMockRecorder) FindImageById(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindImageById", reflect.TypeOf((*MockRepository)(nil).FindImageById), id)
}

// AddProduct mocks base method
func (m *MockRepository) AddProduct(record *Product) (*Product, exception.Exception) {
	ret := m.ctrl.Call(m, "AddProduct", record)
	ret0, _ := ret[0].(*Product)
	ret1, _ := ret[1].(exception.Exception)
	return ret0, ret1
}

// AddProduct indicates an expected call of AddProduct
func (mr *MockRepositoryMockRecorder) AddProduct(record interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProduct", reflect.TypeOf((*MockRepository)(nil).AddProduct), record)
}

// FindProduct mocks base method
func (m *MockRepository) FindProduct(id uint) (*Product, exception.Exception) {
	ret := m.ctrl.Call(m, "FindProduct", id)
	ret0, _ := ret[0].(*Product)
	ret1, _ := ret[1].(exception.Exception)
	return ret0, ret1
}

// FindProduct indicates an expected call of FindProduct
func (mr *MockRepositoryMockRecorder) FindProduct(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindProduct", reflect.TypeOf((*MockRepository)(nil).FindProduct), id)
}
