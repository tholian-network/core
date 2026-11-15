package tholian

import "encoding/json"

func Parse(buffer []byte) Packet {

	packet := NewPacket()

	if len(buffer) > 0 {
		_ = json.Unmarshal(buffer, &packet)
	}

	return packet

}
