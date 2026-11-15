package tholian

import "tholian-core/protocols/tholian/methods"

func NewEvent(method methods.Method, data any) (Packet, error) {

	packet := NewPacket()

	packet.Type = TypeEvent
	packet.Method = method

	blob, err := toRawMessage(data)

	if err != nil {
		return packet, err
	}

	packet.Parameters = blob

	return packet, nil

}
