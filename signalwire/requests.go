package signalwire

// DeviceParamsStruct TODO DESCRIPTION
type DeviceParamsStruct struct {
	ToNumber   string `json:"to_number"`
	FromNumber string `json:"from_number"`
	Timeout    uint   `json:"timeout"`
}

// DeviceStruct TODO DESCRIPTION
type DeviceStruct struct {
	Type   string             `json:"type"`
	Params DeviceParamsStruct `json:"params"`
}

// ParamsCallingBeginStruct TODO DESCRIPTION
type ParamsCallingBeginStruct struct {
	Device DeviceStruct `json:"device"`
	Tag    string       `json:"tag"`
}

// ParamsSignalwireReceive TODO DESCRIPTION
type ParamsSignalwireReceive struct {
	Contexts []string `json:"contexts"`
}

// ParamsCallConnectStruct TODO DESCRIPTION
type ParamsCallConnectStruct struct {
	Devices [][]DeviceStruct `json:"devices"`
	NodeID  string           `json:"node_id"`
	CallID  string           `json:"call_id"`
}

// ParamsCommandStruct TODO DESCRIPTION
type ParamsCommandStruct interface{}

// ParamsBladeExecuteStruct TODO DESCRIPTION
type ParamsBladeExecuteStruct struct {
	Protocol string              `json:"protocol"`
	Method   string              `json:"method"`
	Params   ParamsCommandStruct `json:"params"`
}

// ParamsSubscriptionStruct TODO DESCRIPTION
type ParamsSubscriptionStruct struct {
	Command  string   `json:"command"`
	Protocol string   `json:"protocol"`
	Channels []string `json:"channels"`
}

// ParamsSignalwireSetupStruct TODO DESCRIPTION
type ParamsSignalwireSetupStruct struct {
}

// ParamsSignalwireStruct TODO DESCRIPTION
type ParamsSignalwireStruct struct {
	Params ParamsSignalwireSetupStruct `json:"params"`

	Protocol string `json:"protocol"`
	Method   string `json:"method"`
}

// BladeVersionStruct TODO DESCRIPTION
type BladeVersionStruct struct {
	Major    int `json:"major"`
	Minor    int `json:"minor"`
	Revision int `json:"revision"`
}

// ParamsConnectStruct TODO DESCRIPTION
type ParamsConnectStruct struct {
	Version        BladeVersionStruct `json:"version"`
	SessionID      string             `json:"session_id"`
	Authentication AuthStruct         `json:"authentication"`
}

// AuthStruct TODO DESCRIPTION
type AuthStruct struct {
	Project string `json:"project"`
	Token   string `json:"token"`
}

// ReqBladeConnect TODO DESCRIPTION
type ReqBladeConnect struct {
	Method string              `json:"method"`
	Params ParamsConnectStruct `json:"params"`
}

// ParamsAuthStruct TODO DESCRIPTION
type ParamsAuthStruct struct {
	RequesterNodeID string `json:"requester_node_id"`
	ResponderNodeID string `json:"responder_node_id"`
	OriginalID      string `json:"original_id"`
	NodeID          string `json:"node_id"`
	ConnectionID    string `json:"connection_id"`
}

// ReqBladeAuthenticate TODO DESCRIPTION
type ReqBladeAuthenticate struct {
	Method string           `json:"method"`
	Params ParamsAuthStruct `json:"params"`
}

// ReqBladeSetup TODO DESCRIPTION
type ReqBladeSetup struct {
	Method string                 `json:"method"`
	Params ParamsSignalwireStruct `json:"params"`
}

// ErrorStruct is RPC error object
type ErrorStruct struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ReplyError TODO DESCRIPTION
type ReplyError struct {
	Error ErrorStruct `json:"error"`
}

// ReplyAuthStruct TODO DESCRIPTION
type ReplyAuthStruct struct {
	Project   string   `json:"project"`
	ExpiresAt string   `json:"expires_at"`
	Scopes    []string `json:"scopes"`
	Signature string   `json:"signature"`
}

// ReplyResultConnect TODO DESCRIPTION
type ReplyResultConnect struct {
	SessionRestored      bool            `json:"session_restored"`
	SessionID            string          `json:"session_id"`
	NodeID               string          `json:"node_id"`
	MasterNodeID         string          `json:"master_node_id"`
	Authorization        ReplyAuthStruct `json:"authorization"`
	Routes               []string        `json:"routes"`
	Protocols            []string        `json:"protocols"`
	Subscriptions        []string        `json:"subscriptions"`
	Authorities          []string        `json:"authorities"`
	Authorizations       []string        `json:"authorizations"`
	Accesses             []string        `json:"accesses"`
	ProtocolsUncertified []string        `json:"protocols_uncertified"`
}

// ReplyBladeConnect TODO DESCRIPTION
type ReplyBladeConnect struct {
	Result ReplyResultConnect `json:"result"`
}

// ReplyResultResultSetup TODO DESCRIPTION
type ReplyResultResultSetup struct {
	Protocol string `json:"protocol"`
}

// ReplyResultSetup TODO DESCRIPTION
type ReplyResultSetup struct {
	RequesterNodeID string                 `json:"requester_node_id"`
	ResponderNodeID string                 `json:"responder_node_id"`
	Result          ReplyResultResultSetup `json:"result"`
}

// ReplyResultSubscription TODO DESCRIPTION
type ReplyResultSubscription struct {
	Protocol          string   `json:"protocol"`
	Command           string   `json:"command"`
	SubscribeChannels []string `json:"subscribe_channels"`
}

// ReplyBladeSetup TODO DESCRIPTION
type ReplyBladeSetup struct {
	Result ReplyResultSetup
}

// ReplyBladeExecuteResult TODO DESCRIPTION
type ReplyBladeExecuteResult struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ReplyBladeExecute TODO DESCRIPTION
type ReplyBladeExecute struct {
	RequesterNodeID string                  `json:"requester_node_id"`
	ResponderNodeID string                  `json:"responder_node_id"`
	Result          ReplyBladeExecuteResult `json:"result"`
}

// PeerStruct  TODO DESCRIPTION
type PeerStruct struct {
	CallID string `json:"call_id"`
	NodeID string `json:"node_id"`
}

// PeerDeviceStruct  TODO DESCRIPTION
type PeerDeviceStruct struct {
	CallID string       `json:"call_id"`
	NodeID string       `json:"node_id"`
	Device DeviceStruct `json:"device"`
}

// ParamsEventCallingCallConnect  TODO DESCRIPTION
type ParamsEventCallingCallConnect struct {
	ConnectState string           `json:"connect_state"`
	CallID       string           `json:"call_id"`
	NodeID       string           `json:"node_id"`
	TagID        string           `json:"tag"`
	Peer         PeerDeviceStruct `json:"peer"`
}

// ParamsEventCallingCallState TODO DESCRIPTION
type ParamsEventCallingCallState struct {
	CallState string       `json:"call_state"`
	Direction string       `json:"direction"`
	Device    DeviceStruct `json:"device"`
	EndReason string       `json:"end_reason"`
	CallID    string       `json:"call_id"`
	NodeID    string       `json:"node_id"`
	TagID     string       `json:"tag"`
}

// ParamsEventCallingCallReceive TODO DESCRIPTION
type ParamsEventCallingCallReceive struct {
	CallState string       `json:"call_state"`
	Direction string       `json:"direction"`
	Device    DeviceStruct `json:"device"`
	CallID    string       `json:"call_id"`
	NodeID    string       `json:"node_id"`
	TagID     string       `json:"tag"`
}

// ParamsGenericAction TODO DESCRIPTION
type ParamsGenericAction struct {
	CallID    string `json:"call_id"`
	NodeID    string `json:"node_id"`
	ControlID string `json:"control_id"`
}

// ParamsEventCallingCallPlay TODO DESCRIPTION
type ParamsEventCallingCallPlay struct {
	PlayState string `json:"state"`
	CallID    string `json:"call_id"`
	NodeID    string `json:"node_id"`
	ControlID string `json:"control_id"`
}

// AudioStruct TODO DESCRIPTION
type AudioStruct struct {
	Format    string `json:"format,omitempty"`
	Direction string `json:"direction,omitempty"`
	Stereo    bool   `json:"stereo,omitempty"`
}

// ParamsRecord TODO DESCRIPTION
type ParamsRecord struct {
	Audio AudioStruct `json:"audio"`
}

// ParamsEventCallingCallRecord TODO DESCRIPTION
type ParamsEventCallingCallRecord struct {
	CallID      string       `json:"call_id"`
	NodeID      string       `json:"node_id"`
	ControlID   string       `json:"control_id"`
	TagID       string       `json:"tag"`
	Params      ParamsRecord `json:"params"`
	RecordState string       `json:"state"`
	Duration    uint         `json:"duration"`
	URL         string       `json:"url"`
	Size        uint         `json:"size"`
}

// ParamsEventDetect TODO DESCRIPTION
type ParamsEventDetect struct {
	Event string `json:"event"`
}

// DetectEventStruct TODO DESCRIPTION
type DetectEventStruct struct {
	Type   string            `json:"type"`
	Params ParamsEventDetect `json:"params"`
}

// ParamsEventCallingCallDetect TODO DESCRIPTION
type ParamsEventCallingCallDetect struct {
	CallID    string            `json:"call_id"`
	NodeID    string            `json:"node_id"`
	ControlID string            `json:"control_id"`
	Detect    DetectEventStruct `json:"detect"`
}

type FaxTypeParamsPage struct {
	Direction string `json:"direction"`
	Number    uint16 `json:"number"`
}

type FaxTypeParamsFinished struct {
	Direction      string `json:"direction"`
	Identity       string `json:"identity"`
	RemoteIdentity string `json:"remote_identity"`
	Document       string `json:"document"`
	Pages          uint16 `json:"pages"`
	Success        bool   `json:"success"`
	Result         uint16 `json:"result"`
	ResultText     string `json:"result_text"`
	Format         string `json:"format"`
}

type FaxTypeParamsError struct {
	Description string `json:"description"`
}

type FaxEventStruct struct {
	EventType string                 `json:"type"`
	Params    map[string]interface{} `json:"params"`
}

// ParamsEventCallingFax TODO DESCRIPTION
type ParamsEventCallingFax struct {
	CallID    string         `json:"call_id"`
	NodeID    string         `json:"node_id"`
	ControlID string         `json:"control_id"`
	Fax       FaxEventStruct `json:"fax"`
}

// ParamsQueueingRelayEvents TODO DESCRIPTION
type ParamsQueueingRelayEvents struct {
	EventType    string      `json:"event_type"`
	EventChannel string      `json:"event_channel"`
	Timestamp    float64     `json:"timestamp"`
	Project      string      `json:"project_id"`
	Space        string      `json:"space_id"`
	Params       interface{} `json:"params"`
}

// NotifParamsBladeBroadcast TODO DESCRIPTION
type NotifParamsBladeBroadcast struct {
	BroadcasterNodeID string                    `json:"broadcaster_nodeid"`
	Protocol          string                    `json:"protocol"`
	Channel           string                    `json:"channel"`
	Event             string                    `json:"event"`
	Params            ParamsQueueingRelayEvents `json:"params"`
}

// ParamsNetcastEvent TODO DESCRIPTION
type ParamsNetcastEvent struct {
	Protocol string `json:"protocol"`
}

// NotifParamsBladeNetcast TODO DESCRIPTION
type NotifParamsBladeNetcast struct {
	NetcasterNodeID string             `json:"netcaster_nodeid"`
	Command         string             `json:"command"`
	Params          ParamsNetcastEvent `json:"params"`
}

// ParamsCallEndStruct TODO DESCRIPTION
type ParamsCallEndStruct struct {
	CallID string `json:"call_id"`
	NodeID string `json:"node_id"`
	Reason string `json:"reason"`
}

// ParamsDisconnect - empty
type ParamsDisconnect struct{}

// ReplyResultDisconnect - empty
type ReplyResultDisconnect struct{}

// ParamsCallAnswer TODO DESCRIPTION
type ParamsCallAnswer struct {
	CallID string `json:"call_id"`
	NodeID string `json:"node_id"`
}

// PlayAudioParams TODO DESCRIPTION
type PlayAudioParams struct {
	URL string `json:"url"`
}

// PlayTTSParams TODO DESCRIPTION
type PlayTTSParams struct {
	Text     string `json:"text"`
	Language string `json:"language"`
	Gender   string `json:"gender"`
}

// PlaySilenceParams TODO DESCRIPTION
type PlaySilenceParams struct {
	Duration float64 `json:"duration"`
}

// PlayRingtoneParams TODO DESCRIPTION
type PlayRingtoneParams struct {
	Name     string  `json:"name"`
	Duration float64 `json:"duration"`
}

// PlayParams TODO DESCRIPTION
type PlayParams interface{}

// PlayStruct TODO DESCRIPTION
type PlayStruct struct {
	Type   string     `json:"type"`
	Params PlayParams `json:"params"`
}

// ParamsCallPlay TODO DESCRIPTION
type ParamsCallPlay struct {
	CallID    string       `json:"call_id"`
	NodeID    string       `json:"node_id"`
	ControlID string       `json:"control_id"`
	Volume    float64      `json:"volume"`
	Play      []PlayStruct `json:"play"`
}

// ParamsCallPlayStop TODO DESCRIPTION
type ParamsCallPlayStop ParamsGenericAction

// ParamsCallPlayPause TODO DESCRIPTION
type ParamsCallPlayPause ParamsGenericAction

// ParamsCallPlayResume TODO DESCRIPTION
type ParamsCallPlayResume ParamsGenericAction

// ParamsCallPlayVolume TODO DESCRIPTION
type ParamsCallPlayVolume struct {
	CallID    string  `json:"call_id"`
	NodeID    string  `json:"node_id"`
	ControlID string  `json:"control_id"`
	Volume    float64 `json:"volume"`
}

// RecordParams TODO DESCRIPTION
type RecordParams struct {
	Format            string `json:"format,omitempty"`
	Direction         string `json:"direction,omitempty"`
	Terminators       string `json:"terminators,omitempty"`
	InitialTimeout    uint16 `json:"initial_timeout,omitempty"`
	EndSilenceTimeout uint16 `json:"end_silence_timeout,omitempty"`
	Beep              bool   `json:"beep,omitempty"`
	Stereo            bool   `json:"stereo,omitempty"`
}

// RecordStruct TODO DESCRIPTION
type RecordStruct struct {
	Audio RecordParams `json:"audio"`
}

// ParamsCallRecord TODO DESCRIPTION
type ParamsCallRecord struct {
	CallID    string       `json:"call_id"`
	NodeID    string       `json:"node_id"`
	ControlID string       `json:"control_id"`
	Record    RecordStruct `json:"record"`
}

// ParamsCallRecordStop TODO DESCRIPTION
type ParamsCallRecordStop ParamsGenericAction

// DetectMachineParams TODO DESCRIPTION
type DetectMachineParams struct {
	InitialTimeout        float64 `json:"initial_timeout,omitempty"`
	EndSilenceTimeout     float64 `json:"end_silence_timeout,omitempty"`
	MachineVoiceThreshold float64 `json:"machine_voice_threshold,omitempty"`
	MachineWordsThreshold float64 `json:"machine_words_threshold,omitempty"`
}

// DetectFaxParams TODO DESCRIPTION
type DetectFaxParams struct {
	Tone string `json:"tone,omitempty"`
}

// DetectDigitParams TODO DESCRIPTION
type DetectDigitParams struct {
	Digits string `json:"digits,omitempty"`
}

// DetectStruct TODO DESCRIPTION
type DetectStruct struct {
	Type   string      `json:"type"`
	Params interface{} `json:"params"`
}

// ParamsCallDetect TODO DESCRIPTION
type ParamsCallDetect struct {
	CallID    string       `json:"call_id"`
	NodeID    string       `json:"node_id"`
	ControlID string       `json:"control_id"`
	Detect    DetectStruct `json:"detect"`
	Timeout   float64      `json:"timeout,omitempty"`
}

// ParamsCallDetectStop TODO DESCRIPTION
type ParamsCallDetectStop ParamsGenericAction

type ParamsSendFax struct {
	CallID     string `json:"call_id"`
	NodeID     string `json:"node_id"`
	ControlID  string `json:"control_id"`
	Document   string `json:"document"`
	Identity   string `json:"identity"`
	HeaderInfo string `json:"header_info"`
}

// ParamsSendFaxStop TODO DESCRIPTION
type ParamsFaxStop ParamsGenericAction