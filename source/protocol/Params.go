package protocol

import "encoding/json"
import "tholian-core/matchers"
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

type ChallengeParams struct {
	Nonce string `json:"nonce"`
}

type WelcomeParams struct {
	ServerID  string `json:"server-id"`
	PublicKey string `json:"public-key"`
	Signature string `json:"signature"`
	SystemID  string `json:"system-id,omitempty"`
	SessionID string `json:"session-id,omitempty"`
}

type SystemReportParams struct {
	SystemID string `json:"system"`
}

type SystemCollectParams struct {
	SystemID string   `json:"system"`
	Sections []string `json:"sections,omitempty"`
}

type IncidentsCollectSinceParams struct {
	SystemID string            `json:"system"`
	Search   matchers.Datetime `json:"search"`
}

type PackagesSearchParams struct {
	SystemID string           `json:"system"`
	Search   matchers.Package `json:"search"`
}

type PackagesUpdateParams struct {
	SystemID string             `json:"system"`
	Packages []matchers.Package `json:"packages"`
}

type ProgramsSearchParams struct {
	SystemID string           `json:"system"`
	Search   matchers.Program `json:"search"`
}

type ServicesSearchParams struct {
	SystemID string           `json:"system"`
	Search   matchers.Program `json:"search"`
}

type UsersSearchParams struct {
	SystemID string        `json:"system"`
	Search   matchers.User `json:"search"`
}

type UpdatesSearchParams struct {
	SystemID string          `json:"system"`
	Search   matchers.Update `json:"search"`
}

type NetworksSearchParams struct {
	SystemID string           `json:"system"`
	Search   matchers.Network `json:"search"`
}

type AntiquesSearchParams struct {
	SystemID string           `json:"system"`
	Search   matchers.Antique `json:"search"`
}

type DevicesSearchParams struct {
	SystemID string          `json:"system"`
	Search   matchers.Device `json:"search"`
}

type DrivesSearchParams struct {
	SystemID string `json:"system"`
	Search   string `json:"search"`
}

type FirewallStatusParams struct {
	SystemID string `json:"system"`
}

type FirewallCheckParams struct {
	SystemID string `json:"system"`
	Target   string `json:"target"`
}

type FirewallScopeParams struct {
	SystemID    string                `json:"system"`
	Connections []matchers.Connection `json:"connections,omitempty"`
	Networks    []matchers.Network    `json:"networks,omitempty"`
}

type CallParams struct {
	Systems []string        `json:"systems"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
	GrantID string          `json:"grant"`
	Purpose string          `json:"purpose"`
}

type SystemsReportParams struct {
	SystemID string `json:"system"`
}

type SystemsIncidentsParams struct {
	SystemID string `json:"system"`
	Limit    int    `json:"limit,omitempty"`
}

type SessionCreateParams struct {
	Systems []string `json:"systems"`
	Purpose Purpose  `json:"purpose"`
}

type SessionCloseParams struct {
	SessionID string `json:"session"`
}

type GrantsRequestParams struct {
	Request GrantRequest `json:"request"`
}

type GrantsListParams struct {
	SessionID string `json:"session,omitempty"`
	SystemID  string `json:"system,omitempty"`
}

type GrantsRevokeParams struct {
	GrantID string `json:"grant"`
}

type EventsSubscribeParams struct {
	Systems []string `json:"systems"`
}

type AuditListParams struct {
	SystemID string `json:"system,omitempty"`
	Limit    int    `json:"limit,omitempty"`
}
