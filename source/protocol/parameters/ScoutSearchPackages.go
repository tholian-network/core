package parameters

import "tholian-core/matchers"

type ScoutSearchPackages struct {
	System string           `json:"system"`
	Search matchers.Package `json:"search"`
}
