package protocol

import "crypto/rand"
import "encoding/hex"

func NewID() string {

	buffer := make([]byte, 16)

	_, err := rand.Read(buffer)

	if err == nil {
		return hex.EncodeToString(buffer)
	}

	return ""

}

func NewNonce() string {

	buffer := make([]byte, 32)

	_, err := rand.Read(buffer)

	if err == nil {
		return hex.EncodeToString(buffer)
	}

	return ""

}
