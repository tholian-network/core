package protocol

import "encoding/json"
import "fmt"

func Encode(envelope Envelope) ([]byte, error) {

	if envelope.Version == "" {
		envelope.Version = Version
	}

	return json.Marshal(envelope)

}

func Decode(data []byte, max_bytes int) (Envelope, error) {

	var envelope Envelope

	if max_bytes <= 0 {
		max_bytes = DefaultMaxMessageBytes
	}

	if len(data) > max_bytes {
		return envelope, fmt.Errorf("protocol: message exceeds %d bytes", max_bytes)
	}

	err := json.Unmarshal(data, &envelope)

	if err != nil {
		return envelope, fmt.Errorf("protocol: invalid envelope: %s", err)
	}

	if envelope.Version != Version {
		return envelope, fmt.Errorf("protocol: unsupported version %q", envelope.Version)
	}

	switch envelope.Kind {
	case KindRequest, KindResponse, KindEvent, KindError:
	default:
		return envelope, fmt.Errorf("protocol: invalid kind %q", envelope.Kind)
	}

	return envelope, nil

}

func DecodeParams[T any](envelope Envelope) (T, error) {

	var value T

	if len(envelope.Parameters) == 0 {
		return value, fmt.Errorf("protocol: missing params")
	}

	err := json.Unmarshal(envelope.Parameters, &value)

	if err != nil {
		return value, fmt.Errorf("protocol: invalid params: %s", err)
	}

	return value, nil

}

func DecodeResult[T any](envelope Envelope) (T, error) {

	var value T

	if len(envelope.Result) == 0 {
		return value, fmt.Errorf("protocol: missing result")
	}

	err := json.Unmarshal(envelope.Result, &value)

	if err != nil {
		return value, fmt.Errorf("protocol: invalid result: %s", err)
	}

	return value, nil

}
