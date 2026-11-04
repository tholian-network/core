package parameters

import "tholian-core/matchers"

type GuardRemoveFiles struct {
	System string        `json:"system"`
	Search matchers.File `json:"search"`
}
