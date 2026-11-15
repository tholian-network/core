package guard

import "tholian-core/matchers"

type RemoveFiles struct {
	System string        `json:"system"`
	Search matchers.File `json:"search"`
}
