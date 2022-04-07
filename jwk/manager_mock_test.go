// Code generated by MockGen. DO NOT EDIT.
// Source: jwk/manager.go

// Package mock_jwk is a generated GoMock package.
package jwk_test

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	jose "gopkg.in/square/go-jose.v2"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// AddKey mocks base method.
func (m *MockManager) AddKey(ctx context.Context, set string, key *jose.JSONWebKey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddKey", ctx, set, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddKey indicates an expected call of AddKey.
func (mr *MockManagerMockRecorder) AddKey(ctx, set, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddKey", reflect.TypeOf((*MockManager)(nil).AddKey), ctx, set, key)
}

// AddKeySet mocks base method.
func (m *MockManager) AddKeySet(ctx context.Context, set string, keys *jose.JSONWebKeySet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddKeySet", ctx, set, keys)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddKeySet indicates an expected call of AddKeySet.
func (mr *MockManagerMockRecorder) AddKeySet(ctx, set, keys interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddKeySet", reflect.TypeOf((*MockManager)(nil).AddKeySet), ctx, set, keys)
}

// DeleteKey mocks base method.
func (m *MockManager) DeleteKey(ctx context.Context, set, kid string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteKey", ctx, set, kid)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteKey indicates an expected call of DeleteKey.
func (mr *MockManagerMockRecorder) DeleteKey(ctx, set, kid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteKey", reflect.TypeOf((*MockManager)(nil).DeleteKey), ctx, set, kid)
}

// DeleteKeySet mocks base method.
func (m *MockManager) DeleteKeySet(ctx context.Context, set string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteKeySet", ctx, set)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteKeySet indicates an expected call of DeleteKeySet.
func (mr *MockManagerMockRecorder) DeleteKeySet(ctx, set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteKeySet", reflect.TypeOf((*MockManager)(nil).DeleteKeySet), ctx, set)
}

// GenerateAndPersistKeySet mocks base method.
func (m *MockManager) GenerateAndPersistKeySet(ctx context.Context, set, kid, alg, use string) (*jose.JSONWebKeySet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAndPersistKeySet", ctx, set, kid, alg, use)
	ret0, _ := ret[0].(*jose.JSONWebKeySet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateAndPersistKeySet indicates an expected call of GenerateAndPersistKeySet.
func (mr *MockManagerMockRecorder) GenerateAndPersistKeySet(ctx, set, kid, alg, use interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAndPersistKeySet", reflect.TypeOf((*MockManager)(nil).GenerateAndPersistKeySet), ctx, set, kid, alg, use)
}

// GetKey mocks base method.
func (m *MockManager) GetKey(ctx context.Context, set, kid string) (*jose.JSONWebKeySet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKey", ctx, set, kid)
	ret0, _ := ret[0].(*jose.JSONWebKeySet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKey indicates an expected call of GetKey.
func (mr *MockManagerMockRecorder) GetKey(ctx, set, kid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKey", reflect.TypeOf((*MockManager)(nil).GetKey), ctx, set, kid)
}

// GetKeySet mocks base method.
func (m *MockManager) GetKeySet(ctx context.Context, set string) (*jose.JSONWebKeySet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKeySet", ctx, set)
	ret0, _ := ret[0].(*jose.JSONWebKeySet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKeySet indicates an expected call of GetKeySet.
func (mr *MockManagerMockRecorder) GetKeySet(ctx, set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKeySet", reflect.TypeOf((*MockManager)(nil).GetKeySet), ctx, set)
}

// UpdateKey mocks base method.
func (m *MockManager) UpdateKey(ctx context.Context, set string, key *jose.JSONWebKey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateKey", ctx, set, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateKey indicates an expected call of UpdateKey.
func (mr *MockManagerMockRecorder) UpdateKey(ctx, set, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateKey", reflect.TypeOf((*MockManager)(nil).UpdateKey), ctx, set, key)
}

// UpdateKeySet mocks base method.
func (m *MockManager) UpdateKeySet(ctx context.Context, set string, keys *jose.JSONWebKeySet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateKeySet", ctx, set, keys)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateKeySet indicates an expected call of UpdateKeySet.
func (mr *MockManagerMockRecorder) UpdateKeySet(ctx, set, keys interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateKeySet", reflect.TypeOf((*MockManager)(nil).UpdateKeySet), ctx, set, keys)
}
