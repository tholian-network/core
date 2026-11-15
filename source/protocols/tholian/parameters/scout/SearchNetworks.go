package scout

import "tholian-core/matchers"

type SearchNetworks struct {
	System string           `json:"system"`
	Search matchers.Network `json:"search"`
}
