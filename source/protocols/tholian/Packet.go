package tholian

import "encoding/json"
import "tholian-core/protocols/tholian/methods"

type Packet struct {
	Version    string            `json:"version"`
	ID         string            `json:"id,omitempty"`
	Type       string            `json:"type"`
	Method     methods.Method    `json:"method,omitempty"`
	Parameters json.RawMessage   `json:"parameters,omitempty"`
	Result     json.RawMessage   `json:"result,omitempty"`
	Error      *Error            `json:"error,omitempty"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

func NewPacket() Packet {

	var packet Packet

	packet.Version = Version
	packet.Type = TypeRequest
	packet.Metadata = make(map[string]string)

	return packet

}

func (packet *Packet) Bytes() []byte {

	blob, err := Encode(*packet)

	if err != nil {
		return nil
	}

	return blob

}

func (packet *Packet) IsRequest() bool {

	if packet != nil {
		return packet.Type == TypeRequest
	}

	return false

}

func (packet *Packet) IsResponse() bool {

	if packet != nil {
		return packet.Type == TypeResponse
	}

	return false

}

func (packet *Packet) IsEvent() bool {

	if packet != nil {
		return packet.Type == TypeEvent
	}

	return false

}

func (packet *Packet) IsError() bool {

	if packet != nil {
		return packet.Type == TypeError
	}

	return false

}

func (packet *Packet) SetVersion(value string) {
	packet.Version = value
}

func (packet *Packet) SetID(value string) {
	packet.ID = value
}

func (packet *Packet) SetType(value string) {

	switch value {
	case TypeRequest, TypeResponse, TypeEvent, TypeError:
		packet.Type = value
	}

}

func (packet *Packet) SetMethod(value methods.Method) {
	packet.Method = value
}

func (packet *Packet) SetParameters(value any) error {

	blob, err := toRawMessage(value)

	if err != nil {
		return err
	}

	packet.Parameters = blob

	return nil

}

func (packet *Packet) SetResult(value any) error {

	blob, err := toRawMessage(value)

	if err != nil {
		return err
	}

	packet.Result = blob

	return nil

}

func (packet *Packet) SetError(value *Error) {
	packet.Error = value
}

func (packet *Packet) SetMetadata(key string, value string) {

	if packet.Metadata == nil {
		packet.Metadata = make(map[string]string)
	}

	packet.Metadata[key] = value

}

func (packet *Packet) GetMetadata(key string) string {

	if packet.Metadata != nil {
		return packet.Metadata[key]
	}

	return ""

}

func toRawMessage(value any) (json.RawMessage, error) {

	if value == nil {
		return nil, nil
	}

	if raw, ok := value.(json.RawMessage); ok {
		return raw, nil
	}

	if raw, ok := value.([]byte); ok {
		return json.RawMessage(raw), nil
	}

	blob, err := json.Marshal(value)

	if err != nil {
		return nil, err
	}

	return blob, nil

}
