package tholian

import "tholian-core/protocols/tholian/methods"

func NewRequest(id string, method methods.Method, parameters any) (Packet, error) {

	packet := NewPacket()

	packet.ID = id
	packet.Type = TypeRequest
	packet.Method = method

	blob, err := toRawMessage(parameters)

	if err != nil {
		return packet, err
	}

	packet.Parameters = blob

	return packet, nil

}
