package guard

import "tholian-core/matchers"

type UpdatePackages struct {
	System string           `json:"system"`
	Search matchers.Package `json:"search"`
}
