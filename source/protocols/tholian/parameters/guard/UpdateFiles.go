package guard

import "tholian-core/matchers"

type UpdateFiles struct {
	System string        `json:"system"`
	Search matchers.File `json:"search"`
}
