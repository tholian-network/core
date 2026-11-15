package guard

import "tholian-core/matchers"

type SearchFiles struct {
	System string        `json:"system"`
	Search matchers.File `json:"search"`
}
