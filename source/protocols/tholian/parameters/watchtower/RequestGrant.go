package watchtower

import "tholian-core/protocols/tholian/permissions"

type RequestGrant struct {
	Request permissions.GrantRequest `json:"request"`
}
