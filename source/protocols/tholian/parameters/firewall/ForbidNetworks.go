package firewall

import "tholian-core/matchers"

type ForbidNetworks struct {
	System   string             `json:"system"`
	Networks []matchers.Network `json:"networks"`
}
