package tholian

import "encoding/json"
import "fmt"

func Decode(data []byte, max_bytes int) (Packet, error) {

	var packet Packet

	if max_bytes <= 0 {
		max_bytes = DefaultMaxMessageBytes
	}

	if len(data) > max_bytes {
		return packet, fmt.Errorf("tholian: message exceeds %d bytes", max_bytes)
	}

	err := json.Unmarshal(data, &packet)

	if err != nil {
		return packet, fmt.Errorf("tholian: invalid packet: %s", err)
	}

	if packet.Version != Version {
		return packet, fmt.Errorf("tholian: unsupported version %q", packet.Version)
	}

	switch packet.Type {
	case TypeRequest, TypeResponse, TypeEvent, TypeError:
	default:
		return packet, fmt.Errorf("tholian: invalid type %q", packet.Type)
	}

	return packet, nil

}
