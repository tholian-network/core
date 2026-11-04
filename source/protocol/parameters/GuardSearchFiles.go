package parameters

import "tholian-core/matchers"

type GuardSearchFiles struct {
	System string        `json:"system"`
	Search matchers.File `json:"search"`
}
