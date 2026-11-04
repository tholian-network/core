package parameters

import "tholian-core/matchers"

type GuardUpdatePackages struct {
	System string           `json:"system"`
	Search matchers.Package `json:"search"`
}
