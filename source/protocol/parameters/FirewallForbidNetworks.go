package parameters

import "tholian-core/matchers"

type FirewallForbidNetworks struct {
	System   string             `json:"system"`
	Networks []matchers.Network `json:"networks"`
}
