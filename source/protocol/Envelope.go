package protocol

import "encoding/json"
import "tholian-core/protocol/permissions"

const (
	KindRequest  = "request"
	KindResponse = "response"
	KindEvent    = "event"
	KindError    = "error"
)

type Envelope struct {
	Version  string             `json:"version"`
	ID       string             `json:"id,omitempty"`
	Kind     string             `json:"kind"`
	Method   permissions.Method `json:"method,omitempty"`
	Params   json.RawMessage    `json:"params,omitempty"`
	Result   json.RawMessage    `json:"result,omitempty"`
	Error    *Error             `json:"error,omitempty"`
	Metadata map[string]string  `json:"metadata,omitempty"`
}

func NewRequest(id string, method permissions.Method, params any) (Envelope, error) {

	var envelope Envelope

	envelope.Version = Version
	envelope.ID = id
	envelope.Kind = KindRequest
	envelope.Method = method
	envelope.Metadata = make(map[string]string)

	if params != nil {

		blob, err := json.Marshal(params)

		if err == nil {
			envelope.Params = blob
		} else {
			return envelope, err
		}

	}

	return envelope, nil

}

func NewResponse(id string, method permissions.Method, result any) (Envelope, error) {

	var envelope Envelope

	envelope.Version = Version
	envelope.ID = id
	envelope.Kind = KindResponse
	envelope.Method = method
	envelope.Metadata = make(map[string]string)

	if result != nil {

		blob, err := json.Marshal(result)

		if err == nil {
			envelope.Result = blob
		} else {
			return envelope, err
		}

	}

	return envelope, nil

}

func NewEvent(method permissions.Method, data any) (Envelope, error) {

	var envelope Envelope

	envelope.Version = Version
	envelope.Kind = KindEvent
	envelope.Method = method
	envelope.Metadata = make(map[string]string)

	if data != nil {

		blob, err := json.Marshal(data)

		if err == nil {
			envelope.Params = blob
		} else {
			return envelope, err
		}

	}

	return envelope, nil

}

func NewErrorResponse(id string, method permissions.Method, err *Error) Envelope {

	var envelope Envelope

	envelope.Version = Version
	envelope.ID = id
	envelope.Kind = KindError
	envelope.Method = method
	envelope.Error = err
	envelope.Metadata = make(map[string]string)

	return envelope

}

func (envelope *Envelope) IsRequest() bool {

	if envelope != nil {
		return envelope.Kind == KindRequest
	}

	return false

}

func (envelope *Envelope) IsResponse() bool {

	if envelope != nil {
		return envelope.Kind == KindResponse
	}

	return false

}

func (envelope *Envelope) IsEvent() bool {

	if envelope != nil {
		return envelope.Kind == KindEvent
	}

	return false

}

func (envelope *Envelope) IsError() bool {

	if envelope != nil {
		return envelope.Kind == KindError
	}

	return false

}

func (envelope *Envelope) SetMetadata(key string, value string) {

	if envelope.Metadata == nil {
		envelope.Metadata = make(map[string]string)
	}

	envelope.Metadata[key] = value

}

func (envelope *Envelope) GetMetadata(key string) string {

	if envelope.Metadata != nil {
		return envelope.Metadata[key]
	}

	return ""

}
