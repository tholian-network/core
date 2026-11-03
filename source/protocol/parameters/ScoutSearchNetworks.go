package parameters

import "tholian-core/matchers"

type ScoutSearchNetworks struct {
	System string           `json:"system"`
	Search matchers.Network `json:"search"`
}
