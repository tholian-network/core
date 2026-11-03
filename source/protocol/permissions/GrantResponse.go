package permissions

import "tholian-core/protocol/capabilities"
import "tholian-core/protocol/id"
import "tholian-core/types"
import "time"

type GrantResponse struct {
	ID           string                    `json:"id"`
	System       string                    `json:"system"`
	Subject      string                    `json:"subject"`
	Purpose      Purpose                   `json:"purpose"`
	Capabilities []capabilities.Capability `json:"capabilities"`
	Scope        Scope                     `json:"scope"`
	IssuedAt     types.Datetime            `json:"issued-at"`
	ExpiresAt    types.Datetime            `json:"expires-at"`
	MaxUses      int                       `json:"max-uses"`
	Uses         int                       `json:"uses"`
	Revoked      bool                      `json:"revoked"`
}

func NewGrantResponse(system string, subject string, purpose Purpose, caps []capabilities.Capability, scope Scope, expires_in time.Duration) GrantResponse {

	var grant GrantResponse

	now := time.Now()

	grant.ID = id.NewID()
	grant.System = system
	grant.Subject = subject
	grant.Purpose = purpose
	grant.Capabilities = append([]capabilities.Capability(nil), caps...)
	grant.Scope = scope
	grant.IssuedAt = types.ToDatetime(now.Format(time.RFC3339))
	grant.ExpiresAt = types.ToDatetime(now.Add(expires_in).Format(time.RFC3339))
	grant.MaxUses = 0
	grant.Uses = 0
	grant.Revoked = false

	return grant

}

func (grant *GrantResponse) IsValid() bool {

	if grant == nil {
		return false
	}

	if grant.ID != "" && grant.System != "" && grant.Subject != "" && grant.Purpose.IsValid() && len(grant.Capabilities) > 0 {
		return true
	}

	return false

}

func (grant *GrantResponse) IsExpired(now types.Datetime) bool {

	if grant == nil {
		return true
	}

	return grant.ExpiresAt.IsBefore(now)

}

func (grant *GrantResponse) IsExhausted() bool {

	if grant == nil {
		return true
	}

	if grant.MaxUses > 0 && grant.Uses >= grant.MaxUses {
		return true
	}

	return false

}

func (grant *GrantResponse) HasCapability(capability capabilities.Capability) bool {

	if grant != nil {

		for _, granted := range grant.Capabilities {
			if granted == capability {
				return true
			}
		}

	}

	return false

}

func (grant *GrantResponse) HasSystem(system string) bool {

	if grant != nil {
		return grant.System == system
	}

	return false

}

func (grant *GrantResponse) Allows(system string, capability capabilities.Capability, now types.Datetime) bool {

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
