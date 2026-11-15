package watchtower

import "tholian-core/protocols/tholian/permissions"

type ConnectSystems struct {
	Systems []string            `json:"systems"`
	Purpose permissions.Purpose `json:"purpose"`
}
