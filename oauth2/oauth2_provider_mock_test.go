// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ory/fosite (interfaces: OAuth2Provider)

// Package oauth2_test is a generated GoMock package.
package oauth2_test

import (
	context "context"
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	fosite "github.com/ory/fosite"
)

// MockOAuth2Provider is a mock of OAuth2Provider interface.
type MockOAuth2Provider struct {
	ctrl     *gomock.Controller
	recorder *MockOAuth2ProviderMockRecorder
}

// MockOAuth2ProviderMockRecorder is the mock recorder for MockOAuth2Provider.
type MockOAuth2ProviderMockRecorder struct {
	mock *MockOAuth2Provider
}

// NewMockOAuth2Provider creates a new mock instance.
func NewMockOAuth2Provider(ctrl *gomock.Controller) *MockOAuth2Provider {
	mock := &MockOAuth2Provider{ctrl: ctrl}
	mock.recorder = &MockOAuth2ProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOAuth2Provider) EXPECT() *MockOAuth2ProviderMockRecorder {
	return m.recorder
}

// IntrospectToken mocks base method.
func (m *MockOAuth2Provider) IntrospectToken(arg0 context.Context, arg1 string, arg2 fosite.TokenType, arg3 fosite.Session, arg4 ...string) (fosite.TokenType, fosite.AccessRequester, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IntrospectToken", varargs...)
	ret0, _ := ret[0].(fosite.TokenType)
	ret1, _ := ret[1].(fosite.AccessRequester)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// IntrospectToken indicates an expected call of IntrospectToken.
func (mr *MockOAuth2ProviderMockRecorder) IntrospectToken(arg0, arg1, arg2, arg3 interface{}, arg4 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IntrospectToken", reflect.TypeOf((*MockOAuth2Provider)(nil).IntrospectToken), varargs...)
}

// NewAccessRequest mocks base method.
func (m *MockOAuth2Provider) NewAccessRequest(arg0 context.Context, arg1 *http.Request, arg2 fosite.Session) (fosite.AccessRequester, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAccessRequest", arg0, arg1, arg2)
	ret0, _ := ret[0].(fosite.AccessRequester)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewAccessRequest indicates an expected call of NewAccessRequest.
func (mr *MockOAuth2ProviderMockRecorder) NewAccessRequest(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAccessRequest", reflect.TypeOf((*MockOAuth2Provider)(nil).NewAccessRequest), arg0, arg1, arg2)
}

// NewAccessResponse mocks base method.
func (m *MockOAuth2Provider) NewAccessResponse(arg0 context.Context, arg1 fosite.AccessRequester) (fosite.AccessResponder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAccessResponse", arg0, arg1)
	ret0, _ := ret[0].(fosite.AccessResponder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewAccessResponse indicates an expected call of NewAccessResponse.
func (mr *MockOAuth2ProviderMockRecorder) NewAccessResponse(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAccessResponse", reflect.TypeOf((*MockOAuth2Provider)(nil).NewAccessResponse), arg0, arg1)
}

// NewAuthorizeRequest mocks base method.
func (m *MockOAuth2Provider) NewAuthorizeRequest(arg0 context.Context, arg1 *http.Request) (fosite.AuthorizeRequester, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAuthorizeRequest", arg0, arg1)
	ret0, _ := ret[0].(fosite.AuthorizeRequester)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewAuthorizeRequest indicates an expected call of NewAuthorizeRequest.
func (mr *MockOAuth2ProviderMockRecorder) NewAuthorizeRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAuthorizeRequest", reflect.TypeOf((*MockOAuth2Provider)(nil).NewAuthorizeRequest), arg0, arg1)
}

// NewAuthorizeResponse mocks base method.
func (m *MockOAuth2Provider) NewAuthorizeResponse(arg0 context.Context, arg1 fosite.AuthorizeRequester, arg2 fosite.Session) (fosite.AuthorizeResponder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAuthorizeResponse", arg0, arg1, arg2)
	ret0, _ := ret[0].(fosite.AuthorizeResponder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewAuthorizeResponse indicates an expected call of NewAuthorizeResponse.
func (mr *MockOAuth2ProviderMockRecorder) NewAuthorizeResponse(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAuthorizeResponse", reflect.TypeOf((*MockOAuth2Provider)(nil).NewAuthorizeResponse), arg0, arg1, arg2)
}

// NewDeviceUserRequest mocks base method.
func (m *MockOAuth2Provider) NewDeviceUserRequest(arg0 context.Context, arg1 *http.Request) (fosite.DeviceUserRequester, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDeviceUserRequest", arg0, arg1)
	ret0, _ := ret[0].(fosite.DeviceUserRequester)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewDeviceUserRequest indicates an expected call of NewDeviceUserRequest.
func (mr *MockOAuth2ProviderMockRecorder) NewDeviceUserRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDeviceUserRequest", reflect.TypeOf((*MockOAuth2Provider)(nil).NewDeviceUserRequest), arg0, arg1)
}

// NewDeviceUserResponse mocks base method.
func (m *MockOAuth2Provider) NewDeviceUserResponse(arg0 context.Context, arg1 fosite.DeviceUserRequester, arg2 fosite.Session) (fosite.DeviceUserResponder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDeviceUserResponse", arg0, arg1, arg2)
	ret0, _ := ret[0].(fosite.DeviceUserResponder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewDeviceUserResponse indicates an expected call of NewDeviceUserResponse.
func (mr *MockOAuth2ProviderMockRecorder) NewDeviceUserResponse(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDeviceUserResponse", reflect.TypeOf((*MockOAuth2Provider)(nil).NewDeviceUserResponse), arg0, arg1, arg2)
}

// NewDeviceRequest mocks base method.
func (m *MockOAuth2Provider) NewDeviceRequest(arg0 context.Context, arg1 *http.Request) (fosite.DeviceRequester, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDeviceRequest", arg0, arg1)
	ret0, _ := ret[0].(fosite.DeviceRequester)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewDeviceRequest indicates an expected call of NewDeviceRequest.
func (mr *MockOAuth2ProviderMockRecorder) NewDeviceRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDeviceRequest", reflect.TypeOf((*MockOAuth2Provider)(nil).NewDeviceRequest), arg0, arg1)
}

// NewDeviceResponse mocks base method.
func (m *MockOAuth2Provider) NewDeviceResponse(arg0 context.Context, arg1 fosite.DeviceRequester) (fosite.DeviceResponder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDeviceResponse", arg0, arg1)
	ret0, _ := ret[0].(fosite.DeviceResponder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewDeviceResponse indicates an expected call of NewDeviceResponse.
func (mr *MockOAuth2ProviderMockRecorder) NewDeviceResponse(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDeviceResponse", reflect.TypeOf((*MockOAuth2Provider)(nil).NewDeviceResponse), arg0, arg1)
}

// NewIntrospectionRequest mocks base method.
func (m *MockOAuth2Provider) NewIntrospectionRequest(arg0 context.Context, arg1 *http.Request, arg2 fosite.Session) (fosite.IntrospectionResponder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewIntrospectionRequest", arg0, arg1, arg2)
	ret0, _ := ret[0].(fosite.IntrospectionResponder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewIntrospectionRequest indicates an expected call of NewIntrospectionRequest.
func (mr *MockOAuth2ProviderMockRecorder) NewIntrospectionRequest(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewIntrospectionRequest", reflect.TypeOf((*MockOAuth2Provider)(nil).NewIntrospectionRequest), arg0, arg1, arg2)
}

// NewPushedAuthorizeRequest mocks base method.
func (m *MockOAuth2Provider) NewPushedAuthorizeRequest(arg0 context.Context, arg1 *http.Request) (fosite.AuthorizeRequester, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewPushedAuthorizeRequest", arg0, arg1)
	ret0, _ := ret[0].(fosite.AuthorizeRequester)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewPushedAuthorizeRequest indicates an expected call of NewPushedAuthorizeRequest.
func (mr *MockOAuth2ProviderMockRecorder) NewPushedAuthorizeRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewPushedAuthorizeRequest", reflect.TypeOf((*MockOAuth2Provider)(nil).NewPushedAuthorizeRequest), arg0, arg1)
}

// NewPushedAuthorizeResponse mocks base method.
func (m *MockOAuth2Provider) NewPushedAuthorizeResponse(arg0 context.Context, arg1 fosite.AuthorizeRequester, arg2 fosite.Session) (fosite.PushedAuthorizeResponder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewPushedAuthorizeResponse", arg0, arg1, arg2)
	ret0, _ := ret[0].(fosite.PushedAuthorizeResponder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewPushedAuthorizeResponse indicates an expected call of NewPushedAuthorizeResponse.
func (mr *MockOAuth2ProviderMockRecorder) NewPushedAuthorizeResponse(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewPushedAuthorizeResponse", reflect.TypeOf((*MockOAuth2Provider)(nil).NewPushedAuthorizeResponse), arg0, arg1, arg2)
}

// NewRevocationRequest mocks base method.
func (m *MockOAuth2Provider) NewRevocationRequest(arg0 context.Context, arg1 *http.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewRevocationRequest", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// NewRevocationRequest indicates an expected call of NewRevocationRequest.
func (mr *MockOAuth2ProviderMockRecorder) NewRevocationRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRevocationRequest", reflect.TypeOf((*MockOAuth2Provider)(nil).NewRevocationRequest), arg0, arg1)
}

// WriteAccessError mocks base method.
func (m *MockOAuth2Provider) WriteAccessError(arg0 context.Context, arg1 http.ResponseWriter, arg2 fosite.AccessRequester, arg3 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WriteAccessError", arg0, arg1, arg2, arg3)
}

// WriteAccessError indicates an expected call of WriteAccessError.
func (mr *MockOAuth2ProviderMockRecorder) WriteAccessError(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteAccessError", reflect.TypeOf((*MockOAuth2Provider)(nil).WriteAccessError), arg0, arg1, arg2, arg3)
}

// WriteAccessResponse mocks base method.
func (m *MockOAuth2Provider) WriteAccessResponse(arg0 context.Context, arg1 http.ResponseWriter, arg2 fosite.AccessRequester, arg3 fosite.AccessResponder) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WriteAccessResponse", arg0, arg1, arg2, arg3)
}

// WriteAccessResponse indicates an expected call of WriteAccessResponse.
func (mr *MockOAuth2ProviderMockRecorder) WriteAccessResponse(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteAccessResponse", reflect.TypeOf((*MockOAuth2Provider)(nil).WriteAccessResponse), arg0, arg1, arg2, arg3)
}

// WriteAuthorizeError mocks base method.
func (m *MockOAuth2Provider) WriteAuthorizeError(arg0 context.Context, arg1 http.ResponseWriter, arg2 fosite.AuthorizeRequester, arg3 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WriteAuthorizeError", arg0, arg1, arg2, arg3)
}

// WriteAuthorizeError indicates an expected call of WriteAuthorizeError.
func (mr *MockOAuth2ProviderMockRecorder) WriteAuthorizeError(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteAuthorizeError", reflect.TypeOf((*MockOAuth2Provider)(nil).WriteAuthorizeError), arg0, arg1, arg2, arg3)
}

// WriteAuthorizeResponse mocks base method.
func (m *MockOAuth2Provider) WriteAuthorizeResponse(arg0 context.Context, arg1 http.ResponseWriter, arg2 fosite.AuthorizeRequester, arg3 fosite.AuthorizeResponder) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WriteAuthorizeResponse", arg0, arg1, arg2, arg3)
}

// WriteAuthorizeResponse indicates an expected call of WriteAuthorizeResponse.
func (mr *MockOAuth2ProviderMockRecorder) WriteAuthorizeResponse(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteAuthorizeResponse", reflect.TypeOf((*MockOAuth2Provider)(nil).WriteAuthorizeResponse), arg0, arg1, arg2, arg3)
}

// WriteDeviceUserResponse mocks base method.
func (m *MockOAuth2Provider) WriteDeviceUserResponse(arg0 context.Context, arg1 *http.Request, arg2 http.ResponseWriter, arg3 fosite.DeviceUserRequester, arg4 fosite.DeviceUserResponder) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WriteDeviceUserResponse", arg0, arg1, arg2, arg3, arg4)
}

// WriteDeviceUserResponse indicates an expected call of WriteDeviceUserResponse.
func (mr *MockOAuth2ProviderMockRecorder) WriteDeviceUserResponse(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteDeviceUserResponse", reflect.TypeOf((*MockOAuth2Provider)(nil).WriteDeviceUserResponse), arg0, arg1, arg2, arg3, arg4)
}

// WriteDeviceResponse mocks base method.
func (m *MockOAuth2Provider) WriteDeviceResponse(arg0 context.Context, arg1 http.ResponseWriter, arg2 fosite.DeviceRequester, arg3 fosite.DeviceResponder) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WriteDeviceResponse", arg0, arg1, arg2, arg3)
}

// WriteDeviceResponse indicates an expected call of WriteDeviceResponse.
func (mr *MockOAuth2ProviderMockRecorder) WriteDeviceResponse(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteDeviceResponse", reflect.TypeOf((*MockOAuth2Provider)(nil).WriteDeviceResponse), arg0, arg1, arg2, arg3)
}

// WriteIntrospectionError mocks base method.
func (m *MockOAuth2Provider) WriteIntrospectionError(arg0 context.Context, arg1 http.ResponseWriter, arg2 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WriteIntrospectionError", arg0, arg1, arg2)
}

// WriteIntrospectionError indicates an expected call of WriteIntrospectionError.
func (mr *MockOAuth2ProviderMockRecorder) WriteIntrospectionError(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteIntrospectionError", reflect.TypeOf((*MockOAuth2Provider)(nil).WriteIntrospectionError), arg0, arg1, arg2)
}

// WriteIntrospectionResponse mocks base method.
func (m *MockOAuth2Provider) WriteIntrospectionResponse(arg0 context.Context, arg1 http.ResponseWriter, arg2 fosite.IntrospectionResponder) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WriteIntrospectionResponse", arg0, arg1, arg2)
}

// WriteIntrospectionResponse indicates an expected call of WriteIntrospectionResponse.
func (mr *MockOAuth2ProviderMockRecorder) WriteIntrospectionResponse(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteIntrospectionResponse", reflect.TypeOf((*MockOAuth2Provider)(nil).WriteIntrospectionResponse), arg0, arg1, arg2)
}

// WritePushedAuthorizeError mocks base method.
func (m *MockOAuth2Provider) WritePushedAuthorizeError(arg0 context.Context, arg1 http.ResponseWriter, arg2 fosite.AuthorizeRequester, arg3 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WritePushedAuthorizeError", arg0, arg1, arg2, arg3)
}

// WritePushedAuthorizeError indicates an expected call of WritePushedAuthorizeError.
func (mr *MockOAuth2ProviderMockRecorder) WritePushedAuthorizeError(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WritePushedAuthorizeError", reflect.TypeOf((*MockOAuth2Provider)(nil).WritePushedAuthorizeError), arg0, arg1, arg2, arg3)
}

// WritePushedAuthorizeResponse mocks base method.
func (m *MockOAuth2Provider) WritePushedAuthorizeResponse(arg0 context.Context, arg1 http.ResponseWriter, arg2 fosite.AuthorizeRequester, arg3 fosite.PushedAuthorizeResponder) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WritePushedAuthorizeResponse", arg0, arg1, arg2, arg3)
}

// WritePushedAuthorizeResponse indicates an expected call of WritePushedAuthorizeResponse.
func (mr *MockOAuth2ProviderMockRecorder) WritePushedAuthorizeResponse(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WritePushedAuthorizeResponse", reflect.TypeOf((*MockOAuth2Provider)(nil).WritePushedAuthorizeResponse), arg0, arg1, arg2, arg3)
}

// WriteRevocationResponse mocks base method.
func (m *MockOAuth2Provider) WriteRevocationResponse(arg0 context.Context, arg1 http.ResponseWriter, arg2 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WriteRevocationResponse", arg0, arg1, arg2)
}

// WriteRevocationResponse indicates an expected call of WriteRevocationResponse.
func (mr *MockOAuth2ProviderMockRecorder) WriteRevocationResponse(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteRevocationResponse", reflect.TypeOf((*MockOAuth2Provider)(nil).WriteRevocationResponse), arg0, arg1, arg2)
}
