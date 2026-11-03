package parameters

import "tholian-core/matchers"

type FirewallPermitNetwork struct {
	System  string           `json:"system"`
	Network matchers.Network `json:"network"`
}
