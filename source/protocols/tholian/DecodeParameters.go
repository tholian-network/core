package tholian

import "encoding/json"
import "fmt"

func DecodeParameters[T any](packet Packet) (T, error) {

	var value T

	if len(packet.Parameters) == 0 {
		return value, fmt.Errorf("tholian: missing parameters")
	}

	err := json.Unmarshal(packet.Parameters, &value)

	if err != nil {
		return value, fmt.Errorf("tholian: invalid parameters: %s", err)
	}

	return value, nil

}
