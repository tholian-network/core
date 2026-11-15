package firewall

import "tholian-core/matchers"

type PermitNetwork struct {
	System  string           `json:"system"`
	Network matchers.Network `json:"network"`
}
