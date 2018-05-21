// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/user/auth_handler.go

// Package user is a generated GoMock package.
package user

import (
	exception "github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/exception"
	auth "github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/pkg/generated/swagger/swagserver/swagapi/auth"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAuthHandler is a mock of AuthHandler interface
type MockAuthHandler struct {
	ctrl     *gomock.Controller
	recorder *MockAuthHandlerMockRecorder
}

// MockAuthHandlerMockRecorder is the mock recorder for MockAuthHandler
type MockAuthHandlerMockRecorder struct {
	mock *MockAuthHandler
}

// NewMockAuthHandler creates a new mock instance
func NewMockAuthHandler(ctrl *gomock.Controller) *MockAuthHandler {
	mock := &MockAuthHandler{ctrl: ctrl}
	mock.recorder = &MockAuthHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthHandler) EXPECT() *MockAuthHandlerMockRecorder {
	return m.recorder
}

// Register mocks base method
func (m *MockAuthHandler) Register(params auth.RegisterParams) exception.Exception {
	ret := m.ctrl.Call(m, "Register", params)
	ret0, _ := ret[0].(exception.Exception)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockAuthHandlerMockRecorder) Register(params interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockAuthHandler)(nil).Register), params)
}

// Login mocks base method
func (m *MockAuthHandler) Login(params auth.LoginParams) (*AuthClaim, exception.Exception) {
	ret := m.ctrl.Call(m, "Login", params)
	ret0, _ := ret[0].(*AuthClaim)
	ret1, _ := ret[1].(exception.Exception)
	return ret0, ret1
}

// Login indicates an expected call of Login
func (mr *MockAuthHandlerMockRecorder) Login(params interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthHandler)(nil).Login), params)
}

// Logout mocks base method
func (m *MockAuthHandler) Logout(params auth.LogoutParams) exception.Exception {
	ret := m.ctrl.Call(m, "Logout", params)
	ret0, _ := ret[0].(exception.Exception)
	return ret0
}

// Logout indicates an expected call of Logout
func (mr *MockAuthHandlerMockRecorder) Logout(params interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockAuthHandler)(nil).Logout), params)
}
