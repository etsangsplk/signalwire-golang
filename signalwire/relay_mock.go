// Code generated by MockGen. DO NOT EDIT.
// Source: relay.go

// Package signalwire is a generated GoMock package.
package signalwire

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIRelay is a mock of IRelay interface
type MockIRelay struct {
	ctrl     *gomock.Controller
	recorder *MockIRelayMockRecorder
}

// MockIRelayMockRecorder is the mock recorder for MockIRelay
type MockIRelayMockRecorder struct {
	mock *MockIRelay
}

// NewMockIRelay creates a new mock instance
func NewMockIRelay(ctrl *gomock.Controller) *MockIRelay {
	mock := &MockIRelay{ctrl: ctrl}
	mock.recorder = &MockIRelayMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIRelay) EXPECT() *MockIRelayMockRecorder {
	return m.recorder
}

// RelayPhoneDial mocks base method
func (m *MockIRelay) RelayPhoneDial(ctx context.Context, call *CallSession, fromNumber, toNumber string, timeout uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelayPhoneDial", ctx, call, fromNumber, toNumber, timeout)
	ret0, _ := ret[0].(error)
	return ret0
}

// RelayPhoneDial indicates an expected call of RelayPhoneDial
func (mr *MockIRelayMockRecorder) RelayPhoneDial(ctx, call, fromNumber, toNumber, timeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelayPhoneDial", reflect.TypeOf((*MockIRelay)(nil).RelayPhoneDial), ctx, call, fromNumber, toNumber, timeout)
}

// RelayPhoneConnect mocks base method
func (m *MockIRelay) RelayPhoneConnect(ctx context.Context, call *CallSession, fromNumber, toNumber string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelayPhoneConnect", ctx, call, fromNumber, toNumber)
	ret0, _ := ret[0].(error)
	return ret0
}

// RelayPhoneConnect indicates an expected call of RelayPhoneConnect
func (mr *MockIRelayMockRecorder) RelayPhoneConnect(ctx, call, fromNumber, toNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelayPhoneConnect", reflect.TypeOf((*MockIRelay)(nil).RelayPhoneConnect), ctx, call, fromNumber, toNumber)
}

// RelayCallEnd mocks base method
func (m *MockIRelay) RelayCallEnd(ctx context.Context, call *CallSession) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelayCallEnd", ctx, call)
	ret0, _ := ret[0].(error)
	return ret0
}

// RelayCallEnd indicates an expected call of RelayCallEnd
func (mr *MockIRelayMockRecorder) RelayCallEnd(ctx, call interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelayCallEnd", reflect.TypeOf((*MockIRelay)(nil).RelayCallEnd), ctx, call)
}

// RelayStop mocks base method
func (m *MockIRelay) RelayStop(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelayStop", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// RelayStop indicates an expected call of RelayStop
func (mr *MockIRelayMockRecorder) RelayStop(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelayStop", reflect.TypeOf((*MockIRelay)(nil).RelayStop), ctx)
}

// RelayOnInboundAnswer mocks base method
func (m *MockIRelay) RelayOnInboundAnswer(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RelayOnInboundAnswer", ctx)
}

// RelayOnInboundAnswer indicates an expected call of RelayOnInboundAnswer
func (mr *MockIRelayMockRecorder) RelayOnInboundAnswer(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelayOnInboundAnswer", reflect.TypeOf((*MockIRelay)(nil).RelayOnInboundAnswer), ctx)
}

// RelayPlayAudio mocks base method
func (m *MockIRelay) RelayPlayAudio(ctx context.Context, call *CallSession, ctrlID, url string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelayPlayAudio", ctx, call, ctrlID, url)
	ret0, _ := ret[0].(error)
	return ret0
}

// RelayPlayAudio indicates an expected call of RelayPlayAudio
func (mr *MockIRelayMockRecorder) RelayPlayAudio(ctx, call, ctrlID, url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelayPlayAudio", reflect.TypeOf((*MockIRelay)(nil).RelayPlayAudio), ctx, call, ctrlID, url)
}

// RelayPlayAudioStop mocks base method
func (m *MockIRelay) RelayPlayAudioStop(ctx context.Context, call *CallSession, ctrlID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelayPlayAudioStop", ctx, call, ctrlID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RelayPlayAudioStop indicates an expected call of RelayPlayAudioStop
func (mr *MockIRelayMockRecorder) RelayPlayAudioStop(ctx, call, ctrlID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelayPlayAudioStop", reflect.TypeOf((*MockIRelay)(nil).RelayPlayAudioStop), ctx, call, ctrlID)
}

// RelayRecordAudio mocks base method
func (m *MockIRelay) RelayRecordAudio(ctx context.Context, call *CallSession, ctrlID string, rec RecordParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelayRecordAudio", ctx, call, ctrlID, rec)
	ret0, _ := ret[0].(error)
	return ret0
}

// RelayRecordAudio indicates an expected call of RelayRecordAudio
func (mr *MockIRelayMockRecorder) RelayRecordAudio(ctx, call, ctrlID, rec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelayRecordAudio", reflect.TypeOf((*MockIRelay)(nil).RelayRecordAudio), ctx, call, ctrlID, rec)
}

// RelayRecordAudioStop mocks base method
func (m *MockIRelay) RelayRecordAudioStop(ctx context.Context, call *CallSession, ctrlID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelayRecordAudioStop", ctx, call, ctrlID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RelayRecordAudioStop indicates an expected call of RelayRecordAudioStop
func (mr *MockIRelayMockRecorder) RelayRecordAudioStop(ctx, call, ctrlID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelayRecordAudioStop", reflect.TypeOf((*MockIRelay)(nil).RelayRecordAudioStop), ctx, call, ctrlID)
}

// RelayConnect mocks base method
func (m *MockIRelay) RelayConnect(ctx context.Context, call *CallSession, ringback *[]RingbackStruct, devices *[][]DeviceStruct) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelayConnect", ctx, call, ringback, devices)
	ret0, _ := ret[0].(error)
	return ret0
}

// RelayConnect indicates an expected call of RelayConnect
func (mr *MockIRelayMockRecorder) RelayConnect(ctx, call, ringback, devices interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelayConnect", reflect.TypeOf((*MockIRelay)(nil).RelayConnect), ctx, call, ringback, devices)
}

// RelayCallAnswer mocks base method
func (m *MockIRelay) RelayCallAnswer(ctx context.Context, call *CallSession) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelayCallAnswer", ctx, call)
	ret0, _ := ret[0].(error)
	return ret0
}

// RelayCallAnswer indicates an expected call of RelayCallAnswer
func (mr *MockIRelayMockRecorder) RelayCallAnswer(ctx, call interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelayCallAnswer", reflect.TypeOf((*MockIRelay)(nil).RelayCallAnswer), ctx, call)
}

// RelayPlayTTS mocks base method
func (m *MockIRelay) RelayPlayTTS(ctx context.Context, call *CallSession, ctrlID, text, language, gender string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelayPlayTTS", ctx, call, ctrlID, text, language, gender)
	ret0, _ := ret[0].(error)
	return ret0
}

// RelayPlayTTS indicates an expected call of RelayPlayTTS
func (mr *MockIRelayMockRecorder) RelayPlayTTS(ctx, call, ctrlID, text, language, gender interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelayPlayTTS", reflect.TypeOf((*MockIRelay)(nil).RelayPlayTTS), ctx, call, ctrlID, text, language, gender)
}

// RelayPlayRingtone mocks base method
func (m *MockIRelay) RelayPlayRingtone(ctx context.Context, call *CallSession, ctrlID, name string, duration float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelayPlayRingtone", ctx, call, ctrlID, name, duration)
	ret0, _ := ret[0].(error)
	return ret0
}

// RelayPlayRingtone indicates an expected call of RelayPlayRingtone
func (mr *MockIRelayMockRecorder) RelayPlayRingtone(ctx, call, ctrlID, name, duration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelayPlayRingtone", reflect.TypeOf((*MockIRelay)(nil).RelayPlayRingtone), ctx, call, ctrlID, name, duration)
}

// RelayPlaySilence mocks base method
func (m *MockIRelay) RelayPlaySilence(ctx context.Context, call *CallSession, ctrlID string, duration float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelayPlaySilence", ctx, call, ctrlID, duration)
	ret0, _ := ret[0].(error)
	return ret0
}

// RelayPlaySilence indicates an expected call of RelayPlaySilence
func (mr *MockIRelayMockRecorder) RelayPlaySilence(ctx, call, ctrlID, duration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelayPlaySilence", reflect.TypeOf((*MockIRelay)(nil).RelayPlaySilence), ctx, call, ctrlID, duration)
}

// RelayPlay mocks base method
func (m *MockIRelay) RelayPlay(ctx context.Context, call *CallSession, controlID string, play []PlayStruct) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelayPlay", ctx, call, controlID, play)
	ret0, _ := ret[0].(error)
	return ret0
}

// RelayPlay indicates an expected call of RelayPlay
func (mr *MockIRelayMockRecorder) RelayPlay(ctx, call, controlID, play interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelayPlay", reflect.TypeOf((*MockIRelay)(nil).RelayPlay), ctx, call, controlID, play)
}

// RelayPlayVolume mocks base method
func (m *MockIRelay) RelayPlayVolume(ctx context.Context, call *CallSession, ctrlID *string, vol float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelayPlayVolume", ctx, call, ctrlID, vol)
	ret0, _ := ret[0].(error)
	return ret0
}

// RelayPlayVolume indicates an expected call of RelayPlayVolume
func (mr *MockIRelayMockRecorder) RelayPlayVolume(ctx, call, ctrlID, vol interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelayPlayVolume", reflect.TypeOf((*MockIRelay)(nil).RelayPlayVolume), ctx, call, ctrlID, vol)
}

// RelayPlayResume mocks base method
func (m *MockIRelay) RelayPlayResume(ctx context.Context, call *CallSession, ctrlID *string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelayPlayResume", ctx, call, ctrlID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RelayPlayResume indicates an expected call of RelayPlayResume
func (mr *MockIRelayMockRecorder) RelayPlayResume(ctx, call, ctrlID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelayPlayResume", reflect.TypeOf((*MockIRelay)(nil).RelayPlayResume), ctx, call, ctrlID)
}

// RelayPlayPause mocks base method
func (m *MockIRelay) RelayPlayPause(ctx context.Context, call *CallSession, ctrlID *string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelayPlayPause", ctx, call, ctrlID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RelayPlayPause indicates an expected call of RelayPlayPause
func (mr *MockIRelayMockRecorder) RelayPlayPause(ctx, call, ctrlID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelayPlayPause", reflect.TypeOf((*MockIRelay)(nil).RelayPlayPause), ctx, call, ctrlID)
}

// RelayPlayStop mocks base method
func (m *MockIRelay) RelayPlayStop(ctx context.Context, call *CallSession, ctrlID *string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelayPlayStop", ctx, call, ctrlID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RelayPlayStop indicates an expected call of RelayPlayStop
func (mr *MockIRelayMockRecorder) RelayPlayStop(ctx, call, ctrlID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelayPlayStop", reflect.TypeOf((*MockIRelay)(nil).RelayPlayStop), ctx, call, ctrlID)
}

// RelayDetectDigit mocks base method
func (m *MockIRelay) RelayDetectDigit(ctx context.Context, call *CallSession, controlID, digits string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelayDetectDigit", ctx, call, controlID, digits)
	ret0, _ := ret[0].(error)
	return ret0
}

// RelayDetectDigit indicates an expected call of RelayDetectDigit
func (mr *MockIRelayMockRecorder) RelayDetectDigit(ctx, call, controlID, digits interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelayDetectDigit", reflect.TypeOf((*MockIRelay)(nil).RelayDetectDigit), ctx, call, controlID, digits)
}

// RelayDetectFax mocks base method
func (m *MockIRelay) RelayDetectFax(ctx context.Context, call *CallSession, controlID, faxtone string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelayDetectFax", ctx, call, controlID, faxtone)
	ret0, _ := ret[0].(error)
	return ret0
}

// RelayDetectFax indicates an expected call of RelayDetectFax
func (mr *MockIRelayMockRecorder) RelayDetectFax(ctx, call, controlID, faxtone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelayDetectFax", reflect.TypeOf((*MockIRelay)(nil).RelayDetectFax), ctx, call, controlID, faxtone)
}

// RelayDetectMachine mocks base method
func (m *MockIRelay) RelayDetectMachine(ctx context.Context, call *CallSession, controlID string, det *DetectMachineParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelayDetectMachine", ctx, call, controlID, det)
	ret0, _ := ret[0].(error)
	return ret0
}

// RelayDetectMachine indicates an expected call of RelayDetectMachine
func (mr *MockIRelayMockRecorder) RelayDetectMachine(ctx, call, controlID, det interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelayDetectMachine", reflect.TypeOf((*MockIRelay)(nil).RelayDetectMachine), ctx, call, controlID, det)
}

// RelayDetect mocks base method
func (m *MockIRelay) RelayDetect(ctx context.Context, call *CallSession, controlID string, detect DetectStruct) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelayDetect", ctx, call, controlID, detect)
	ret0, _ := ret[0].(error)
	return ret0
}

// RelayDetect indicates an expected call of RelayDetect
func (mr *MockIRelayMockRecorder) RelayDetect(ctx, call, controlID, detect interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelayDetect", reflect.TypeOf((*MockIRelay)(nil).RelayDetect), ctx, call, controlID, detect)
}

// RelayDetectStop mocks base method
func (m *MockIRelay) RelayDetectStop(ctx context.Context, call *CallSession, ctrlID *string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelayDetectStop", ctx, call, ctrlID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RelayDetectStop indicates an expected call of RelayDetectStop
func (mr *MockIRelayMockRecorder) RelayDetectStop(ctx, call, ctrlID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelayDetectStop", reflect.TypeOf((*MockIRelay)(nil).RelayDetectStop), ctx, call, ctrlID)
}

// RelaySendFax mocks base method
func (m *MockIRelay) RelaySendFax(ctx context.Context, call *CallSession, ctrlID *string, doc, id, headerInfo string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelaySendFax", ctx, call, ctrlID, doc, id, headerInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// RelaySendFax indicates an expected call of RelaySendFax
func (mr *MockIRelayMockRecorder) RelaySendFax(ctx, call, ctrlID, doc, id, headerInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelaySendFax", reflect.TypeOf((*MockIRelay)(nil).RelaySendFax), ctx, call, ctrlID, doc, id, headerInfo)
}

// RelayReceiveFax mocks base method
func (m *MockIRelay) RelayReceiveFax(ctx context.Context, call *CallSession, ctrlID *string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelayReceiveFax", ctx, call, ctrlID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RelayReceiveFax indicates an expected call of RelayReceiveFax
func (mr *MockIRelayMockRecorder) RelayReceiveFax(ctx, call, ctrlID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelayReceiveFax", reflect.TypeOf((*MockIRelay)(nil).RelayReceiveFax), ctx, call, ctrlID)
}

// RelaySendFaxStop mocks base method
func (m *MockIRelay) RelaySendFaxStop(ctx context.Context, call *CallSession, ctrlID *string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelaySendFaxStop", ctx, call, ctrlID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RelaySendFaxStop indicates an expected call of RelaySendFaxStop
func (mr *MockIRelayMockRecorder) RelaySendFaxStop(ctx, call, ctrlID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelaySendFaxStop", reflect.TypeOf((*MockIRelay)(nil).RelaySendFaxStop), ctx, call, ctrlID)
}

// RelayReceiveFaxStop mocks base method
func (m *MockIRelay) RelayReceiveFaxStop(ctx context.Context, call *CallSession, ctrlID *string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelayReceiveFaxStop", ctx, call, ctrlID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RelayReceiveFaxStop indicates an expected call of RelayReceiveFaxStop
func (mr *MockIRelayMockRecorder) RelayReceiveFaxStop(ctx, call, ctrlID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelayReceiveFaxStop", reflect.TypeOf((*MockIRelay)(nil).RelayReceiveFaxStop), ctx, call, ctrlID)
}

// RelayTapAudio mocks base method
func (m *MockIRelay) RelayTapAudio(ctx context.Context, call *CallSession, ctrlID, direction string, device *TapDevice) (TapDevice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelayTapAudio", ctx, call, ctrlID, direction, device)
	ret0, _ := ret[0].(TapDevice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RelayTapAudio indicates an expected call of RelayTapAudio
func (mr *MockIRelayMockRecorder) RelayTapAudio(ctx, call, ctrlID, direction, device interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelayTapAudio", reflect.TypeOf((*MockIRelay)(nil).RelayTapAudio), ctx, call, ctrlID, direction, device)
}

// RelayTap mocks base method
func (m *MockIRelay) RelayTap(ctx context.Context, call *CallSession, controlID string, tap TapStruct, device *TapDevice) (TapDevice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelayTap", ctx, call, controlID, tap, device)
	ret0, _ := ret[0].(TapDevice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RelayTap indicates an expected call of RelayTap
func (mr *MockIRelayMockRecorder) RelayTap(ctx, call, controlID, tap, device interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelayTap", reflect.TypeOf((*MockIRelay)(nil).RelayTap), ctx, call, controlID, tap, device)
}

// RelayTapStop mocks base method
func (m *MockIRelay) RelayTapStop(ctx context.Context, call *CallSession, ctrlID *string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelayTapStop", ctx, call, ctrlID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RelayTapStop indicates an expected call of RelayTapStop
func (mr *MockIRelayMockRecorder) RelayTapStop(ctx, call, ctrlID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelayTapStop", reflect.TypeOf((*MockIRelay)(nil).RelayTapStop), ctx, call, ctrlID)
}

// RelaySendDigits mocks base method
func (m *MockIRelay) RelaySendDigits(ctx context.Context, call *CallSession, controlID, digits string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelaySendDigits", ctx, call, controlID, digits)
	ret0, _ := ret[0].(error)
	return ret0
}

// RelaySendDigits indicates an expected call of RelaySendDigits
func (mr *MockIRelayMockRecorder) RelaySendDigits(ctx, call, controlID, digits interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelaySendDigits", reflect.TypeOf((*MockIRelay)(nil).RelaySendDigits), ctx, call, controlID, digits)
}

// RelayPlayAndCollect mocks base method
func (m *MockIRelay) RelayPlayAndCollect(ctx context.Context, call *CallSession, controlID string, playlist *[]PlayStruct, collect *CollectStruct) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelayPlayAndCollect", ctx, call, controlID, playlist, collect)
	ret0, _ := ret[0].(error)
	return ret0
}

// RelayPlayAndCollect indicates an expected call of RelayPlayAndCollect
func (mr *MockIRelayMockRecorder) RelayPlayAndCollect(ctx, call, controlID, playlist, collect interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelayPlayAndCollect", reflect.TypeOf((*MockIRelay)(nil).RelayPlayAndCollect), ctx, call, controlID, playlist, collect)
}

// RelayPlayAndCollectVolume mocks base method
func (m *MockIRelay) RelayPlayAndCollectVolume(ctx context.Context, call *CallSession, ctrlID *string, vol float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelayPlayAndCollectVolume", ctx, call, ctrlID, vol)
	ret0, _ := ret[0].(error)
	return ret0
}

// RelayPlayAndCollectVolume indicates an expected call of RelayPlayAndCollectVolume
func (mr *MockIRelayMockRecorder) RelayPlayAndCollectVolume(ctx, call, ctrlID, vol interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelayPlayAndCollectVolume", reflect.TypeOf((*MockIRelay)(nil).RelayPlayAndCollectVolume), ctx, call, ctrlID, vol)
}

// RelayPlayAndCollectStop mocks base method
func (m *MockIRelay) RelayPlayAndCollectStop(ctx context.Context, call *CallSession, ctrlID *string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelayPlayAndCollectStop", ctx, call, ctrlID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RelayPlayAndCollectStop indicates an expected call of RelayPlayAndCollectStop
func (mr *MockIRelayMockRecorder) RelayPlayAndCollectStop(ctx, call, ctrlID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelayPlayAndCollectStop", reflect.TypeOf((*MockIRelay)(nil).RelayPlayAndCollectStop), ctx, call, ctrlID)
}