package tholian

import "tholian-core/protocols/tholian/methods"

func NewResponse(id string, method methods.Method, result any) (Packet, error) {

	packet := NewPacket()

	packet.ID = id
	packet.Type = TypeResponse
	packet.Method = method

	blob, err := toRawMessage(result)

	if err != nil {
		return packet, err
	}

	packet.Result = blob

	return packet, nil

}
