package core

import "tholian-core/structs"

type Hello struct {
	Role            string         `json:"role"`
	NodeID          string         `json:"node-id"`
	PublicKey       string         `json:"public-key"`
	Signature       string         `json:"signature"`
	Nonce           string         `json:"nonce"`
	EnrollmentToken string         `json:"enrollment-token,omitempty"`
	ProtocolVersion string         `json:"protocol-version"`
	System          structs.System `json:"system,omitempty"`
}
