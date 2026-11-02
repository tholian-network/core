package protocol

import "tholian-core/protocol/id"

func NewID() string {
	return id.NewID()
}

func NewNonce() string {
	return id.NewNonce()
}
