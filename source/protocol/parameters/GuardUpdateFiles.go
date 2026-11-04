package parameters

import "tholian-core/matchers"

type GuardUpdateFiles struct {
	System string        `json:"system"`
	Search matchers.File `json:"search"`
}
