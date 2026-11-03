package protocol

import "encoding/json"
import "tholian-core/protocol/methods"
import "tholian-core/protocol/permissions"
import "tholian-core/structs"
import "tholian-core/types"

type ActionResult struct {
	Success bool                     `json:"success"`
	Data    json.RawMessage          `json:"data,omitempty"`
	Console []structs.ConsoleMessage `json:"console,omitempty"`
	Error   string                   `json:"error,omitempty"`
}

type SystemSummary struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	Hostname  string         `json:"hostname"`
	Hash      string         `json:"hash,omitempty"`
	Online    bool           `json:"online"`
	LastSeen  types.Datetime `json:"last-seen"`
	Incidents int            `json:"incidents"`
}

type SystemsListResult struct {
	Systems []SystemSummary `json:"systems"`
}

type SystemReportResult struct {
	System structs.System `json:"system"`
}

type IncidentsResult struct {
	Incidents []structs.Incident `json:"incidents"`
}

type PackagesSearchResult struct {
	Packages []structs.Package `json:"packages"`
}

type ProgramsSearchResult struct {
	Programs []structs.Program `json:"programs"`
}

type ServicesSearchResult struct {
	Services []structs.Program `json:"services"`
}

type UsersSearchResult struct {
	Users []types.User `json:"users"`
}

type UpdatesSearchResult struct {
	Updates []structs.Update `json:"updates"`
}

type NetworksSearchResult struct {
	Networks []structs.Network `json:"networks"`
}

type AntiquesSearchResult struct {
	Antiques []structs.Antique `json:"antiques"`
}

type DevicesSearchResult struct {
	Devices []structs.Device `json:"devices"`
}

type DrivesSearchResult struct {
	Drives []structs.Drive `json:"drives"`
}

type FirewallStatusResult struct {
	Entries []string `json:"entries"`
}

type FirewallCheckResult struct {
	Forbidden bool `json:"forbidden"`
}

type SystemCallResult struct {
	SystemID string          `json:"system"`
	Success  bool            `json:"success"`
	Result   json.RawMessage `json:"result,omitempty"`
	Error    *Error          `json:"error,omitempty"`
}

type CallResult struct {
	Results []SystemCallResult `json:"results"`
}

type SessionResult struct {
	SessionID string              `json:"session"`
	Purpose   permissions.Purpose `json:"purpose"`
}

type GrantsResult struct {
	Rsults []permissions.GrantResponse `json:"results"`
}

type AuditEntry struct {
	Datetime  types.Datetime `json:"datetime"`
	Actor     string         `json:"actor"`
	SessionID string         `json:"session,omitempty"`
	SystemID  string         `json:"system,omitempty"`
	Purpose   string         `json:"purpose,omitempty"`
	Method    methods.Method `json:"method,omitempty"`
	Grant     string         `json:"grant,omitempty"`
	Decision  string         `json:"decision"`
	Params    string         `json:"params,omitempty"`
	Outcome   string         `json:"outcome,omitempty"`
}

type AuditResult struct {
	Entries []AuditEntry `json:"entries"`
}

type IncidentEvent struct {
	SystemID string           `json:"system"`
	Incident structs.Incident `json:"incident"`
}

type ReportEvent struct {
	SystemID string         `json:"system"`
	Report   structs.System `json:"report"`
}

type GuardEvent struct {
	SystemID string `json:"system"`
	Name     string `json:"name"`
	Hostname string `json:"hostname"`
	Online   bool   `json:"online"`
}

type GrantEvent struct {
	SystemID string                    `json:"system"`
	Grant    permissions.GrantResponse `json:"grant"`
	Action   string                    `json:"action"`
}
