package parameters

import "tholian-core/matchers"

type ScoutUpdatePackage struct {
	System string           `json:"system"`
	Search matchers.Package `json:"search"`
}
