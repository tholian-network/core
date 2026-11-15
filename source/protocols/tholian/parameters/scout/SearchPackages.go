package scout

import "tholian-core/matchers"

type SearchPackages struct {
	System string           `json:"system"`
	Search matchers.Package `json:"search"`
}
