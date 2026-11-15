package scout

import "tholian-core/matchers"

type UpdatePackage struct {
	System string           `json:"system"`
	Search matchers.Package `json:"search"`
}
