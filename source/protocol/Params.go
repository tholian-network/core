package protocol

import "encoding/json"
import "tholian-core/matchers"
import "tholian-core/protocol/methods"
import "tholian-core/protocol/permissions"
import "tholian-core/structs"

type HelloParams struct {
	Role            string         `json:"role"`
	NodeID          string         `json:"node-id"`
	PublicKey       string         `json:"public-key"`
	Signature       string         `json:"signature"`
	Nonce           string         `json:"nonce"`
	EnrollmentToken string         `json:"enrollment-token,omitempty"`
	ProtocolVersion string         `json:"protocol-version"`
	System          structs.System `json:"system,omitempty"`
}

type SystemReportParams struct {
	System string `json:"system"`
}

type SystemCollectParams struct {
	System   string   `json:"system"`
	Sections []string `json:"sections,omitempty"`
}

type FirewallScopeParams struct {
	System      string                `json:"system"`
	Connections []matchers.Connection `json:"connections,omitempty"`
	Networks    []matchers.Network    `json:"networks,omitempty"`
}

type CallParams struct {
	Systems []string        `json:"systems"`
	Method  methods.Method  `json:"method"`
	Params  json.RawMessage `json:"params"`
	Grant   string          `json:"grant"`
	Purpose string          `json:"purpose"`
}

type SessionCreateParams struct {
	Systems []string            `json:"systems"`
	Purpose permissions.Purpose `json:"purpose"`
}

type SessionCloseParams struct {
	Session string `json:"session"`
}

type GrantsRequestParams struct {
	Request permissions.GrantRequest `json:"request"`
}

type GrantsListParams struct {
	Session string `json:"session,omitempty"`
	System  string `json:"system,omitempty"`
}

type GrantsRevokeParams struct {
	Grant string `json:"grant"`
}

type EventsSubscribeParams struct {
	Systems []string `json:"systems"`
}

type AuditListParams struct {
	System string `json:"system,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}

type SystemActionParams struct {
	System string `json:"system"`
}

