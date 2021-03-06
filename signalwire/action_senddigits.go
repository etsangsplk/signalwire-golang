package signalwire

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"sync"
)

// SendDigitsState keeps the state of a SendDigits action
type SendDigitsState int

// TODO DESCRIPTION
const (
	SendDigitsFinished SendDigitsState = iota
)

func (s SendDigitsState) String() string {
	return [...]string{"Finished"}[s]
}

// SendDigitsResult TODO DESCRIPTION
type SendDigitsResult struct {
	Successful bool
	Event      json.RawMessage
}

// SendDigitsAction TODO DESCRIPTION
type SendDigitsAction struct {
	CallObj   *CallObj
	ControlID string
	Completed bool
	Result    SendDigitsResult
	State     SendDigitsState
	Payload   *json.RawMessage
	err       error
	sync.RWMutex
}

// ISendDigits TODO DESCRIPTION
type ISendDigits interface {
	GetCompleted() bool
	GetResult() SendDigitsResult
}

func checkDtmf(s string) bool {
	allowed := "wW1234567890*#ABCD"

	for _, c := range s {
		if !strings.Contains(allowed, string(c)) {
			return false
		}
	}

	return true
}

// SendDigits TODO DESCRIPTION
func (callobj *CallObj) SendDigits(digits string) (*SendDigitsResult, error) {
	if !checkDtmf(digits) {
		return nil, errors.New("invalid DTMF")
	}

	a := new(SendDigitsAction)

	if callobj.Calling == nil {
		return &a.Result, errors.New("nil Calling object")
	}

	if callobj.Calling.Relay == nil {
		return &a.Result, errors.New("nil Relay object")
	}

	ctrlID, _ := GenUUIDv4()
	err := callobj.Calling.Relay.RelaySendDigits(callobj.Calling.Ctx, callobj.call, ctrlID, digits, nil)

	if err != nil {
		return &a.Result, err
	}

	callobj.callbacksRunSendDigits(callobj.Calling.Ctx, ctrlID, a, true)

	return &a.Result, nil
}

// callbacksRunSendDigits TODO DESCRIPTION
func (callobj *CallObj) callbacksRunSendDigits(_ context.Context, ctrlID string, res *SendDigitsAction, norunCB bool) {
	var out bool

	for {
		select {
		case state := <-callobj.call.CallSendDigitsChans[ctrlID]:
			res.RLock()

			prevstate := res.State

			res.RUnlock()

			switch state {
			case SendDigitsFinished:
				res.Lock()

				res.State = state
				res.Result.Successful = true
				res.Completed = true

				res.Unlock()

				Log.Debug("SendDigits finished. ctrlID: %s res [%p] Completed [%v] Successful [%v]\n", ctrlID, res, res.Completed, res.Result.Successful)

				out = true

				if callobj.OnSendDigitsFinished != nil && !norunCB {
					callobj.OnSendDigitsFinished(res)
				}

			default:
				Log.Debug("Unknown state. ctrlID: %s\n", ctrlID)
			}

			if prevstate != state && callobj.OnSendDigitsStateChange != nil && !norunCB {
				callobj.OnSendDigitsStateChange(res)
			}
		case rawEvent := <-callobj.call.CallSendDigitsRawEventChans[ctrlID]:
			res.Lock()
			res.Result.Event = *rawEvent
			res.Unlock()

			callobj.call.CallSendDigitsReadyChans[ctrlID] <- struct{}{}
		case <-callobj.call.Hangup:
			out = true
		}

		if out {
			break
		}
	}
}

// SendDigitsAsync TODO DESCRIPTION
func (callobj *CallObj) SendDigitsAsync(digits string) (*SendDigitsAction, error) {
	if !checkDtmf(digits) {
		return nil, errors.New("invalid DTMF")
	}

	res := new(SendDigitsAction)

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
			ctrlID := <-callobj.call.CallSendDigitsControlIDs

			callobj.callbacksRunSendDigits(callobj.Calling.Ctx, ctrlID, res, false)
		}()

		newCtrlID, _ := GenUUIDv4()

		res.Lock()

		res.ControlID = newCtrlID

		res.Unlock()

		err := callobj.Calling.Relay.RelaySendDigits(callobj.Calling.Ctx, callobj.call, newCtrlID, digits, &res.Payload)

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

// GetCompleted TODO DESCRIPTION
func (action *SendDigitsAction) GetCompleted() bool {
	action.RLock()

	ret := action.Completed

	action.RUnlock()

	return ret
}

// GetResult TODO DESCRIPTION
func (action *SendDigitsAction) GetResult() SendDigitsResult {
	action.RLock()

	ret := action.Result

	action.RUnlock()

	return ret
}

// GetSuccessful TODO DESCRIPTION
func (action *SendDigitsAction) GetSuccessful() bool {
	action.RLock()

	ret := action.Result.Successful

	action.RUnlock()

	return ret
}

// GetEvent TODO DESCRIPTION
func (action *SendDigitsAction) GetEvent() *json.RawMessage {
	action.RLock()

	ret := &action.Result.Event

	action.RUnlock()

	return ret
}

// GetPayload TODO DESCRIPTION
func (action *SendDigitsAction) GetPayload() *json.RawMessage {
	action.RLock()

	ret := action.Payload

	action.RUnlock()

	return ret
}

// GetControlID TODO DESCRIPTION
func (action *SendDigitsAction) GetControlID() string {
	action.RLock()

	ret := action.ControlID

	action.RUnlock()

	return ret
}
