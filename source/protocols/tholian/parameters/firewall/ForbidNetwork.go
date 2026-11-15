package firewall

import "tholian-core/matchers"

type ForbidNetwork struct {
	System  string           `json:"system"`
	Network matchers.Network `json:"network"`
}
