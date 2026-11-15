package firewall

import "tholian-core/matchers"

type PermitNetworks struct {
	System   string             `json:"system"`
	Networks []matchers.Network `json:"networks"`
}
