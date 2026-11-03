package parameters

import "tholian-core/matchers"

type FirewallForbidNetwork struct {
	System  string           `json:"system"`
	Network matchers.Network `json:"network"`
}
