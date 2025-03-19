// Code generated by MockGen. DO NOT EDIT.
// Source: internal/services/services.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	dto "github.com/Axel791/auth/internal/usecases/auth/dto"
	gomock "github.com/golang/mock/gomock"
)

// MockHashPasswordService is a mock of HashPasswordService interface.
type MockHashPasswordService struct {
	ctrl     *gomock.Controller
	recorder *MockHashPasswordServiceMockRecorder
}

// MockHashPasswordServiceMockRecorder is the mock recorder for MockHashPasswordService.
type MockHashPasswordServiceMockRecorder struct {
	mock *MockHashPasswordService
}

// NewMockHashPasswordService creates a new mock instance.
func NewMockHashPasswordService(ctrl *gomock.Controller) *MockHashPasswordService {
	mock := &MockHashPasswordService{ctrl: ctrl}
	mock.recorder = &MockHashPasswordServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHashPasswordService) EXPECT() *MockHashPasswordServiceMockRecorder {
	return m.recorder
}

// Hash mocks base method.
func (m *MockHashPasswordService) Hash(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hash", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Hash indicates an expected call of Hash.
func (mr *MockHashPasswordServiceMockRecorder) Hash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hash", reflect.TypeOf((*MockHashPasswordService)(nil).Hash), arg0)
}

// MockTokenService is a mock of TokenService interface.
type MockTokenService struct {
	ctrl     *gomock.Controller
	recorder *MockTokenServiceMockRecorder
}

// MockTokenServiceMockRecorder is the mock recorder for MockTokenService.
type MockTokenServiceMockRecorder struct {
	mock *MockTokenService
}

// NewMockTokenService creates a new mock instance.
func NewMockTokenService(ctrl *gomock.Controller) *MockTokenService {
	mock := &MockTokenService{ctrl: ctrl}
	mock.recorder = &MockTokenServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenService) EXPECT() *MockTokenServiceMockRecorder {
	return m.recorder
}

// GenerateToken mocks base method.
func (m *MockTokenService) GenerateToken(claimsDTO dto.ClaimsDTO) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", claimsDTO)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockTokenServiceMockRecorder) GenerateToken(claimsDTO interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockTokenService)(nil).GenerateToken), claimsDTO)
}

// ValidateToken mocks base method.
func (m *MockTokenService) ValidateToken(tokenStr string) (dto.ClaimsDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateToken", tokenStr)
	ret0, _ := ret[0].(dto.ClaimsDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateToken indicates an expected call of ValidateToken.
func (mr *MockTokenServiceMockRecorder) ValidateToken(tokenStr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateToken", reflect.TypeOf((*MockTokenService)(nil).ValidateToken), tokenStr)
}
