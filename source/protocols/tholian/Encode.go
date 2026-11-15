package tholian

import "encoding/json"

func Encode(packet Packet) ([]byte, error) {

	if packet.Version == "" {
		packet.Version = Version
	}

	return json.Marshal(packet)

}
