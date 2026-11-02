package protocol

import "tholian-core/types"
import "time"

type Capability string

const (
	CapabilitySystemRead     Capability = "system.read"
	CapabilityPackagesRead   Capability = "packages.read"
	CapabilityPackagesUpdate Capability = "packages.update"
	CapabilityProgramsRead   Capability = "programs.read"
	CapabilityServicesRead   Capability = "services.read"
	CapabilityUsersRead      Capability = "users.read"
	CapabilityUpdatesRead    Capability = "updates.read"
	CapabilityIncidentsRead  Capability = "incidents.read"
	CapabilityFirewallRead   Capability = "firewall.read"
	CapabilityFirewallWrite  Capability = "firewall.write"
)

const (
	CapabilityProgramsUpdate Capability = "programs.update"
	CapabilityServicesUpdate Capability = "services.update"
	CapabilityUsersUpdate    Capability = "users.update"
	CapabilityEvidenceRead   Capability = "evidence.read"
	CapabilityEvidenceWrite  Capability = "evidence.write"
	CapabilityShellExecute   Capability = "shell.execute"
)

const (
	PurposeInvestigation = "investigation"
	PurposeIncident      = "incident"
	PurposeMitigation    = "mitigation"
	PurposeTask          = "task"
)

type Purpose struct {
	ID          string `json:"id"`
	Kind        string `json:"kind"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Incident    string `json:"incident,omitempty"`
}

func NewPurpose(kind string, title string) Purpose {

	return Purpose{
		ID:          NewID(),
		Kind:        kind,
		Title:       title,
		Description: "",
		Incident:    "",
	}

}

func (purpose *Purpose) IsValid() bool {

	if purpose == nil {
		return false
	}

	if purpose.ID != "" && purpose.Kind != "" && purpose.Title != "" {

		if isPurposeKind(purpose.Kind) {
			return true
		}

	}

	return false

}

type Grant struct {
	ID           string         `json:"id"`
	SystemID     string         `json:"system"`
	Subject      string         `json:"subject"`
	Purpose      Purpose        `json:"purpose"`
	Capabilities []Capability   `json:"capabilities"`
	Scope        Scope          `json:"scope"`
	IssuedAt     types.Datetime `json:"issued-at"`
	ExpiresAt    types.Datetime `json:"expires-at"`
	MaxUses      int            `json:"max-uses"`
	Uses         int            `json:"uses"`
	Revoked      bool           `json:"revoked"`
}

func NewGrant(system string, subject string, purpose Purpose, capabilities []Capability, scope Scope, expires_in time.Duration) Grant {

	var grant Grant

	now := time.Now()

	grant.ID = NewID()
	grant.SystemID = system
	grant.Subject = subject
	grant.Purpose = purpose
	grant.Capabilities = append([]Capability(nil), capabilities...)
	grant.Scope = scope
	grant.IssuedAt = types.ToDatetime(now.Format(time.RFC3339))
	grant.ExpiresAt = types.ToDatetime(now.Add(expires_in).Format(time.RFC3339))
	grant.MaxUses = 0
	grant.Uses = 0
	grant.Revoked = false

	return grant

}

func (grant *Grant) IsValid() bool {

	if grant == nil {
		return false
	}

	if grant.ID != "" && grant.SystemID != "" && grant.Subject != "" && grant.Purpose.IsValid() && len(grant.Capabilities) > 0 {
		return true
	}

	return false

}

func (grant *Grant) IsExpired(now types.Datetime) bool {

	if grant == nil {
		return true
	}

	return grant.ExpiresAt.IsBefore(now)

}

func (grant *Grant) IsExhausted() bool {

	if grant == nil {
		return true
	}

	if grant.MaxUses > 0 && grant.Uses >= grant.MaxUses {
		return true
	}

	return false

}

func (grant *Grant) HasCapability(capability Capability) bool {

	if grant != nil {

		for _, granted := range grant.Capabilities {
			if granted == capability {
				return true
			}
		}

	}

	return false

}

func (grant *Grant) HasSystem(system string) bool {

	if grant != nil {
		return grant.SystemID == system
	}

	return false

}

func (grant *Grant) Allows(system string, capability Capability, now types.Datetime) bool {

	if grant == nil {
		return false
	}

	if grant.Revoked == true {
		return false
	}

	if grant.IsExpired(now) == true || grant.IsExhausted() == true {
		return false
	}

	if grant.HasSystem(system) == false {
		return false
	}

	return grant.HasCapability(capability)

}

type GrantRequest struct {
	SystemID     string       `json:"system"`
	Purpose      Purpose      `json:"purpose"`
	Capabilities []Capability `json:"capabilities"`
	Scope        Scope        `json:"scope"`
	ExpiresIn    int          `json:"expires-in"`
}

func (request *GrantRequest) IsValid() bool {

	if request == nil {
		return false
	}

	if request.SystemID != "" && request.Purpose.IsValid() && len(request.Capabilities) > 0 {
		return true
	}

	return false

}

func IsSupportedCapability(value Capability) bool {

	switch value {
	case CapabilitySystemRead,
		CapabilityPackagesRead,
		CapabilityPackagesUpdate,
		CapabilityProgramsRead,
		CapabilityServicesRead,
		CapabilityUsersRead,
		CapabilityUpdatesRead,
		CapabilityIncidentsRead,
		CapabilityFirewallRead,
		CapabilityFirewallWrite:
		return true
	}

	return false

}

func IsReservedCapability(value Capability) bool {

	switch value {
	case CapabilityProgramsUpdate,
		CapabilityServicesUpdate,
		CapabilityUsersUpdate,
		CapabilityEvidenceRead,
		CapabilityEvidenceWrite,
		CapabilityShellExecute:
		return true
	}

	return false

}

func MethodCapability(method string) (Capability, bool) {

	switch method {
	case MethodSystemReport, MethodSystemCollect:
		return CapabilitySystemRead, true
	case MethodIncidentsCollectSince:
		return CapabilityIncidentsRead, true
	case MethodPackagesSearch:
		return CapabilityPackagesRead, true
	case MethodPackagesUpdate:
		return CapabilityPackagesUpdate, true
	case MethodProgramsSearch:
		return CapabilityProgramsRead, true
	case MethodServicesSearch:
		return CapabilityServicesRead, true
	case MethodUsersSearch:
		return CapabilityUsersRead, true
	case MethodUpdatesSearch:
		return CapabilityUpdatesRead, true
	case MethodNetworksSearch, MethodDevicesSearch, MethodDrivesSearch:
		return CapabilitySystemRead, true
	case MethodAntiquesSearch:
		return CapabilityPackagesRead, true
	case MethodFirewallStatus, MethodFirewallCheck:
		return CapabilityFirewallRead, true
	case MethodFirewallPermit, MethodFirewallForbid:
		return CapabilityFirewallWrite, true
	}

	return Capability(""), false

}

func isPurposeKind(value string) bool {

	switch value {
	case PurposeInvestigation, PurposeIncident, PurposeMitigation, PurposeTask:
		return true
	}

	return false

}
