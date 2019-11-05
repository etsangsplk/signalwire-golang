package signalwire

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"sync"
)

// CollectResultType keeps the result type of a Collect action
type CollectResultType int

// Collect Result Type constants
const (
	CollectResultError CollectResultType = iota
	CollectResultNoInput
	CollectResultNoMatch
	CollectResultDigit
	CollectResultSpeech
	CollectResultStartOfSpeech
)

func (s CollectResultType) String() string {
	return [...]string{"Error", "No_input", "No_match", "Digit", "Speech", "Start_of_speech"}[s]
}

// CollectResult TODO DESCRIPTION
type CollectResult struct {
	Successful bool
	Terminator string
	Confidence float64
	Result     string
	ResultType CollectResultType
	Continue   CollectContinue
	Event      json.RawMessage
}

// PromptResult TODO DESCRIPTION
type PromptResult CollectResult

// CollectContinue  TODO DESCRIPTION
type CollectContinue int

// Collect Result Type constants
const (
	CollectPartial CollectContinue = iota
	CollectFinal
)

// PromptAction TODO DESCRIPTION
type PromptAction struct {
	CallObj   *CallObj
	ControlID string
	Completed bool
	Result    CollectResult
	err       error
	sync.RWMutex
}

// IPromptAction TODO DESCRIPTION
type IPromptAction interface {
	playAndCollectAsyncStop() error
	Stop()
	GetCompleted() bool
	GetResult() PlayResult
	Volume(vol float64) (*PlayVolumeResult, error)
	GetEvent() *json.RawMessage
}

// Prompt TODO DESCRIPTION
func (callobj *CallObj) Prompt(playlist *[]PlayStruct, collect *CollectStruct) (*CollectResult, error) {
	a := new(PromptAction)

	if callobj.Calling == nil {
		return &a.Result, errors.New("nil Calling object")
	}

	if callobj.Calling.Relay == nil {
		return &a.Result, errors.New("nil Relay object")
	}

	ctrlID, _ := GenUUIDv4()

	err := callobj.Calling.Relay.RelayPlayAndCollect(callobj.Calling.Ctx, callobj.call, ctrlID, playlist, collect)

	if err != nil {
		return &a.Result, err
	}

	callobj.callbacksRunPlayAndCollect(callobj.Calling.Ctx, ctrlID, a)

	return &a.Result, nil
}

// PromptStop TODO DESCRIPTION
func (callobj *CallObj) PromptStop(ctrlID *string) error {
	if callobj.Calling == nil {
		return errors.New("nil Calling object")
	}

	if callobj.Calling.Relay == nil {
		return errors.New("nil Relay object")
	}

	return callobj.Calling.Relay.RelayPlayAndCollectStop(callobj.Calling.Ctx, callobj.call, ctrlID)
}

// callbacksRunPlayAndCollect TODO DESCRIPTION
func (callobj *CallObj) callbacksRunPlayAndCollect(_ context.Context, ctrlID string, res *PromptAction) {
	var cont bool

	var out bool

	for {
		select {
		case resType := <-callobj.call.CallPlayAndCollectChans[ctrlID]:
			switch resType {
			case CollectResultError:
				fallthrough
			case CollectResultNoInput:
				res.Lock()

				res.Result.ResultType = resType
				res.Result.Successful = false
				res.Completed = true

				res.Unlock()

				out = true

			case CollectResultNoMatch:
				fallthrough
			case CollectResultDigit:
				fallthrough
			case CollectResultSpeech:
				fallthrough
			case CollectResultStartOfSpeech:
				res.Lock()

				res.Result.ResultType = resType
				res.Result.Successful = true

				if !cont {
					res.Completed = true
				}

				res.Unlock()

				Log.Debug("Prompt finished. ctrlID: %s res [%p] Completed [%v] Successful [%v]\n", ctrlID, res, res.Completed, res.Result.Successful)

				out = true

				if callobj.OnPrompt != nil {
					callobj.OnPrompt(res)
				}

			default:
				Log.Debug("Unknown state. ctrlID: %s\n", ctrlID)
			}
		case params := <-callobj.call.CallPlayAndCollectEventChans[ctrlID]:
			Log.Debug("got params for ctrlID : %s params: %v\n", ctrlID, params)

			res.Lock()

			if strings.EqualFold(params.Result.Type, CollectResultSpeech.String()) {
				speech := params.Result.Params

				confidence, ok1 := speech["confidence"].(float64)
				if !ok1 {
					Log.Error("type assertion error")

					out = true
				} else {
					res.Result.Confidence = confidence
				}

				text, ok2 := speech["text"].(string)
				if !ok2 {
					Log.Error("type assertion error")

					out = true
				} else {
					res.Result.Result = text
				}
			} else if strings.EqualFold(params.Result.Type, CollectResultDigit.String()) {
				digit := params.Result.Params

				terminator, ok1 := digit["terminator"].(string)
				if !ok1 {
					Log.Error("type assertion error")

					out = true
				} else {
					res.Result.Terminator = terminator
				}

				digits, ok2 := digit["digits"].(string)
				if !ok2 {
					Log.Error("type assertion error")

					out = true
				} else {
					res.Result.Result = digits
				}
			}

			res.Unlock()

			callobj.call.CallPlayAndCollectReadyChans[ctrlID] <- struct{}{}

		case rawEvent := <-callobj.call.CallPlayAndCollectRawEventChans[ctrlID]:
			res.Lock()
			res.Result.Event = *rawEvent
			res.Unlock()

			callobj.call.CallPlayAndCollectReadyChans[ctrlID] <- struct{}{}
		case <-callobj.call.Hangup:
			out = true
		}

		if out {
			break
		}
	}
}

// PromptAsync TODO DESCRIPTION
func (callobj *CallObj) PromptAsync(playlist *[]PlayStruct, collect *CollectStruct) (*PromptAction, error) {
	res := new(PromptAction)

	if callobj.Calling == nil {
		return res, errors.New("nil Calling object")
	}

	if callobj.Calling.Relay == nil {
		return res, errors.New("nil Relay object")
	}

	res.CallObj = callobj

	done := make(chan struct{}, 1)

	go func() {
		go func() {
			// wait to get control ID (buffered channel)
			ctrlID := <-callobj.call.CallPlayAndCollectControlID

			callobj.callbacksRunPlayAndCollect(callobj.Calling.Ctx, ctrlID, res)
		}()

		newCtrlID, _ := GenUUIDv4()
		res.Lock()
		res.ControlID = newCtrlID
		res.Unlock()

		err := callobj.Calling.Relay.RelayPlayAndCollect(callobj.Calling.Ctx, callobj.call, newCtrlID, playlist, collect)

		if err != nil {
			res.Lock()

			res.err = err

			res.Completed = true

			res.Unlock()
		}
		done <- struct{}{}
	}()

	<-done

	return res, res.err
}

// ctrlIDCopy TODO DESCRIPTION
func (action *PromptAction) ctrlIDCopy() (string, error) {
	action.RLock()

	if len(action.ControlID) == 0 {
		action.RUnlock()
		return "", errors.New("no controlID")
	}

	c := action.ControlID

	action.RUnlock()

	return c, nil
}

// playAndCollectAsyncStop TODO DESCRIPTION
func (action *PromptAction) playAndCollectAsyncStop() error {
	if action.CallObj.Calling == nil {
		return errors.New("nil Calling object")
	}

	if action.CallObj.Calling.Relay == nil {
		return errors.New("nil Relay object")
	}

	c, err := action.ctrlIDCopy()
	if err != nil {
		return err
	}

	call := action.CallObj.call

	return action.CallObj.Calling.Relay.RelayPlayAndCollectStop(action.CallObj.Calling.Ctx, call, &c)
}

// Stop TODO DESCRIPTION
func (action *PromptAction) Stop() {
	action.err = action.playAndCollectAsyncStop()
}

// GetCompleted TODO DESCRIPTION
func (action *PromptAction) GetCompleted() bool {
	action.RLock()

	ret := action.Completed

	action.RUnlock()

	return ret
}

// GetResult TODO DESCRIPTION
func (action *PromptAction) GetResult() CollectResult {
	action.RLock()

	ret := action.Result

	action.RUnlock()

	return ret
}

// GetCollectResult TODO DESCRIPTION
func (action *PromptAction) GetCollectResult() string {
	action.RLock()

	ret := action.Result.Result

	action.RUnlock()

	return ret
}

// GetConfidence TODO DESCRIPTION
func (action *PromptAction) GetConfidence() float64 {
	action.RLock()

	ret := action.Result.Confidence

	action.RUnlock()

	return ret
}

// GetTerminator TODO DESCRIPTION
func (action *PromptAction) GetTerminator() string {
	action.RLock()

	ret := action.Result.Terminator

	action.RUnlock()

	return ret
}

// GetResultType TODO DESCRIPTION
func (action *PromptAction) GetResultType() CollectResultType {
	action.RLock()

	ret := action.Result.ResultType

	action.RUnlock()

	return ret
}

// Volume TODO DESCRIPTION
func (action *PromptAction) Volume(vol float64) (*PlayVolumeResult, error) {
	res := new(PlayVolumeResult)

	if action.CallObj.Calling == nil {
		return res, errors.New("nil Calling object")
	}

	if action.CallObj.Calling.Relay == nil {
		return res, errors.New("nil Relay object")
	}

	c, err := action.ctrlIDCopy()
	if err != nil {
		return res, err
	}

	call := action.CallObj.call

	err = action.CallObj.Calling.Relay.RelayPlayAndCollectVolume(action.CallObj.Calling.Ctx, call, &c, vol)

	if err != nil {
		return res, err
	}

	res.Successful = true

	return res, nil
}

// GetEvent TODO DESCRIPTION
func (action *PromptAction) GetEvent() *json.RawMessage {
	action.RLock()

	ret := &action.Result.Event

	action.RUnlock()

	return ret
}
