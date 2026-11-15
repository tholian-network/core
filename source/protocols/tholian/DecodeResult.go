package tholian

import "encoding/json"
import "fmt"

func DecodeResult[T any](packet Packet) (T, error) {

	var value T

	if len(packet.Result) == 0 {
		return value, fmt.Errorf("tholian: missing result")
	}

	err := json.Unmarshal(packet.Result, &value)

	if err != nil {
		return value, fmt.Errorf("tholian: invalid result: %s", err)
	}

	return value, nil

}
