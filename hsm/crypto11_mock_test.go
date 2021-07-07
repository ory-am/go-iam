// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ThalesIgnite/crypto11 (interfaces: SignerDecrypter)

// Package hsm_test is a generated GoMock package.
package hsm_test

import (
	crypto "crypto"
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSignerDecrypter is a mock of SignerDecrypter interface.
type MockSignerDecrypter struct {
	ctrl     *gomock.Controller
	recorder *MockSignerDecrypterMockRecorder
}

// MockSignerDecrypterMockRecorder is the mock recorder for MockSignerDecrypter.
type MockSignerDecrypterMockRecorder struct {
	mock *MockSignerDecrypter
}

// NewMockSignerDecrypter creates a new mock instance.
func NewMockSignerDecrypter(ctrl *gomock.Controller) *MockSignerDecrypter {
	mock := &MockSignerDecrypter{ctrl: ctrl}
	mock.recorder = &MockSignerDecrypterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSignerDecrypter) EXPECT() *MockSignerDecrypterMockRecorder {
	return m.recorder
}

// Decrypt mocks base method.
func (m *MockSignerDecrypter) Decrypt(arg0 io.Reader, arg1 []byte, arg2 crypto.DecrypterOpts) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decrypt", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decrypt indicates an expected call of Decrypt.
func (mr *MockSignerDecrypterMockRecorder) Decrypt(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decrypt", reflect.TypeOf((*MockSignerDecrypter)(nil).Decrypt), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockSignerDecrypter) Delete() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete")
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockSignerDecrypterMockRecorder) Delete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSignerDecrypter)(nil).Delete))
}

// Public mocks base method.
func (m *MockSignerDecrypter) Public() crypto.PublicKey {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Public")
	ret0, _ := ret[0].(crypto.PublicKey)
	return ret0
}

// Public indicates an expected call of Public.
func (mr *MockSignerDecrypterMockRecorder) Public() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Public", reflect.TypeOf((*MockSignerDecrypter)(nil).Public))
}

// Sign mocks base method.
func (m *MockSignerDecrypter) Sign(arg0 io.Reader, arg1 []byte, arg2 crypto.SignerOpts) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sign", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sign indicates an expected call of Sign.
func (mr *MockSignerDecrypterMockRecorder) Sign(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sign", reflect.TypeOf((*MockSignerDecrypter)(nil).Sign), arg0, arg1, arg2)
}
