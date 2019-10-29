// Code generated by MockGen. DO NOT EDIT.
// Source: blade.go

// Package signalwire is a generated GoMock package.
package signalwire

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	websocket "github.com/gorilla/websocket"
	jsonrpc2 "github.com/sourcegraph/jsonrpc2"
	url "net/url"
	reflect "reflect"
)

// MockIBlade is a mock of IBlade interface
type MockIBlade struct {
	ctrl     *gomock.Controller
	recorder *MockIBladeMockRecorder
}

// MockIBladeMockRecorder is the mock recorder for MockIBlade
type MockIBladeMockRecorder struct {
	mock *MockIBlade
}

// NewMockIBlade creates a new mock instance
func NewMockIBlade(ctrl *gomock.Controller) *MockIBlade {
	mock := &MockIBlade{ctrl: ctrl}
	mock.recorder = &MockIBladeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIBlade) EXPECT() *MockIBladeMockRecorder {
	return m.recorder
}

// GetConnection mocks base method
func (m *MockIBlade) GetConnection() (*jsonrpc2.Conn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConnection")
	ret0, _ := ret[0].(*jsonrpc2.Conn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConnection indicates an expected call of GetConnection
func (mr *MockIBladeMockRecorder) GetConnection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConnection", reflect.TypeOf((*MockIBlade)(nil).GetConnection))
}

// BladeCleanup mocks base method
func (m *MockIBlade) BladeCleanup() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BladeCleanup")
	ret0, _ := ret[0].(error)
	return ret0
}

// BladeCleanup indicates an expected call of BladeCleanup
func (mr *MockIBladeMockRecorder) BladeCleanup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BladeCleanup", reflect.TypeOf((*MockIBlade)(nil).BladeCleanup))
}

// BladeWSOpenConn mocks base method
func (m *MockIBlade) BladeWSOpenConn(ctx context.Context, u url.URL) (*websocket.Conn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BladeWSOpenConn", ctx, u)
	ret0, _ := ret[0].(*websocket.Conn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BladeWSOpenConn indicates an expected call of BladeWSOpenConn
func (mr *MockIBladeMockRecorder) BladeWSOpenConn(ctx, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BladeWSOpenConn", reflect.TypeOf((*MockIBlade)(nil).BladeWSOpenConn), ctx, u)
}

// BladeInit mocks base method
func (m *MockIBlade) BladeInit(ctx context.Context, addr string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BladeInit", ctx, addr)
	ret0, _ := ret[0].(error)
	return ret0
}

// BladeInit indicates an expected call of BladeInit
func (mr *MockIBladeMockRecorder) BladeInit(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BladeInit", reflect.TypeOf((*MockIBlade)(nil).BladeInit), ctx, addr)
}

// BladeConnect mocks base method
func (m *MockIBlade) BladeConnect(ctx context.Context, bladeAuth *BladeAuth) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BladeConnect", ctx, bladeAuth)
	ret0, _ := ret[0].(error)
	return ret0
}

// BladeConnect indicates an expected call of BladeConnect
func (mr *MockIBladeMockRecorder) BladeConnect(ctx, bladeAuth interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BladeConnect", reflect.TypeOf((*MockIBlade)(nil).BladeConnect), ctx, bladeAuth)
}

// BladeSetup mocks base method
func (m *MockIBlade) BladeSetup(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BladeSetup", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// BladeSetup indicates an expected call of BladeSetup
func (mr *MockIBladeMockRecorder) BladeSetup(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BladeSetup", reflect.TypeOf((*MockIBlade)(nil).BladeSetup), ctx)
}

// BladeAddSubscription mocks base method
func (m *MockIBlade) BladeAddSubscription(ctx context.Context, signalwireChannels []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BladeAddSubscription", ctx, signalwireChannels)
	ret0, _ := ret[0].(error)
	return ret0
}

// BladeAddSubscription indicates an expected call of BladeAddSubscription
func (mr *MockIBladeMockRecorder) BladeAddSubscription(ctx, signalwireChannels interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BladeAddSubscription", reflect.TypeOf((*MockIBlade)(nil).BladeAddSubscription), ctx, signalwireChannels)
}

// BladeExecute mocks base method
func (m *MockIBlade) BladeExecute(ctx context.Context, v, res interface{}) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BladeExecute", ctx, v, res)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BladeExecute indicates an expected call of BladeExecute
func (mr *MockIBladeMockRecorder) BladeExecute(ctx, v, res interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BladeExecute", reflect.TypeOf((*MockIBlade)(nil).BladeExecute), ctx, v, res)
}

// BladeSignalwireReceive mocks base method
func (m *MockIBlade) BladeSignalwireReceive(ctx context.Context, signalwireContexts []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BladeSignalwireReceive", ctx, signalwireContexts)
	ret0, _ := ret[0].(error)
	return ret0
}

// BladeSignalwireReceive indicates an expected call of BladeSignalwireReceive
func (mr *MockIBladeMockRecorder) BladeSignalwireReceive(ctx, signalwireContexts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BladeSignalwireReceive", reflect.TypeOf((*MockIBlade)(nil).BladeSignalwireReceive), ctx, signalwireContexts)
}

// BladeWaitDisconnect mocks base method
func (m *MockIBlade) BladeWaitDisconnect(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BladeWaitDisconnect", ctx)
}

// BladeWaitDisconnect indicates an expected call of BladeWaitDisconnect
func (mr *MockIBladeMockRecorder) BladeWaitDisconnect(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BladeWaitDisconnect", reflect.TypeOf((*MockIBlade)(nil).BladeWaitDisconnect), ctx)
}

// BladeDisconnect mocks base method
func (m *MockIBlade) BladeDisconnect(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BladeDisconnect", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// BladeDisconnect indicates an expected call of BladeDisconnect
func (mr *MockIBladeMockRecorder) BladeDisconnect(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BladeDisconnect", reflect.TypeOf((*MockIBlade)(nil).BladeDisconnect), ctx)
}

// BladeWaitInboundCall mocks base method
func (m *MockIBlade) BladeWaitInboundCall(ctx context.Context) (*CallSession, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BladeWaitInboundCall", ctx)
	ret0, _ := ret[0].(*CallSession)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BladeWaitInboundCall indicates an expected call of BladeWaitInboundCall
func (mr *MockIBladeMockRecorder) BladeWaitInboundCall(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BladeWaitInboundCall", reflect.TypeOf((*MockIBlade)(nil).BladeWaitInboundCall), ctx)
}

// handleBladeBroadcast mocks base method
func (m *MockIBlade) handleBladeBroadcast(ctx context.Context, req *jsonrpc2.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "handleBladeBroadcast", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// handleBladeBroadcast indicates an expected call of handleBladeBroadcast
func (mr *MockIBladeMockRecorder) handleBladeBroadcast(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleBladeBroadcast", reflect.TypeOf((*MockIBlade)(nil).handleBladeBroadcast), ctx, req)
}

// handleBladeNetcast mocks base method
func (m *MockIBlade) handleBladeNetcast(ctx context.Context, req *jsonrpc2.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "handleBladeNetcast", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// handleBladeNetcast indicates an expected call of handleBladeNetcast
func (mr *MockIBladeMockRecorder) handleBladeNetcast(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleBladeNetcast", reflect.TypeOf((*MockIBlade)(nil).handleBladeNetcast), ctx, req)
}

// handleBladeDisconnect mocks base method
func (m *MockIBlade) handleBladeDisconnect(ctx context.Context, c *jsonrpc2.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "handleBladeDisconnect", ctx, c)
	ret0, _ := ret[0].(error)
	return ret0
}

// handleBladeDisconnect indicates an expected call of handleBladeDisconnect
func (mr *MockIBladeMockRecorder) handleBladeDisconnect(ctx, c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleBladeDisconnect", reflect.TypeOf((*MockIBlade)(nil).handleBladeDisconnect), ctx, c)
}

// handleInboundCall mocks base method
func (m *MockIBlade) handleInboundCall(ctx context.Context, callID string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "handleInboundCall", ctx, callID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// handleInboundCall indicates an expected call of handleInboundCall
func (mr *MockIBladeMockRecorder) handleInboundCall(ctx, callID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleInboundCall", reflect.TypeOf((*MockIBlade)(nil).handleInboundCall), ctx, callID)
}

// handleInboundMessage mocks base method
func (m *MockIBlade) handleInboundMessage(ctx context.Context, callID string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "handleInboundMessage", ctx, callID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// handleInboundMessage indicates an expected call of handleInboundMessage
func (mr *MockIBladeMockRecorder) handleInboundMessage(ctx, callID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleInboundMessage", reflect.TypeOf((*MockIBlade)(nil).handleInboundMessage), ctx, callID)
}

// eventNotif mocks base method
func (m *MockIBlade) eventNotif(ctx context.Context, broadcast NotifParamsBladeBroadcast) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "eventNotif", ctx, broadcast)
	ret0, _ := ret[0].(error)
	return ret0
}

// eventNotif indicates an expected call of eventNotif
func (mr *MockIBladeMockRecorder) eventNotif(ctx, broadcast interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "eventNotif", reflect.TypeOf((*MockIBlade)(nil).eventNotif), ctx, broadcast)
}

// MockISessionControl is a mock of ISessionControl interface
type MockISessionControl struct {
	ctrl     *gomock.Controller
	recorder *MockISessionControlMockRecorder
}

// MockISessionControlMockRecorder is the mock recorder for MockISessionControl
type MockISessionControlMockRecorder struct {
	mock *MockISessionControl
}

// NewMockISessionControl creates a new mock instance
func NewMockISessionControl(ctrl *gomock.Controller) *MockISessionControl {
	mock := &MockISessionControl{ctrl: ctrl}
	mock.recorder = &MockISessionControlMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockISessionControl) EXPECT() *MockISessionControlMockRecorder {
	return m.recorder
}

// addBlade mocks base method
func (m *MockISessionControl) addBlade(c *jsonrpc2.Conn, b *BladeSession) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "addBlade", c, b)
}

// addBlade indicates an expected call of addBlade
func (mr *MockISessionControlMockRecorder) addBlade(c, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "addBlade", reflect.TypeOf((*MockISessionControl)(nil).addBlade), c, b)
}

// getBlade mocks base method
func (m *MockISessionControl) getBlade(c *jsonrpc2.Conn) *BladeSession {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getBlade", c)
	ret0, _ := ret[0].(*BladeSession)
	return ret0
}

// getBlade indicates an expected call of getBlade
func (mr *MockISessionControlMockRecorder) getBlade(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getBlade", reflect.TypeOf((*MockISessionControl)(nil).getBlade), c)
}

// removeBlade mocks base method
func (m *MockISessionControl) removeBlade(c *jsonrpc2.Conn) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "removeBlade", c)
}

// removeBlade indicates an expected call of removeBlade
func (mr *MockISessionControlMockRecorder) removeBlade(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "removeBlade", reflect.TypeOf((*MockISessionControl)(nil).removeBlade), c)
}
