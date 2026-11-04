package parameters

import "tholian-core/matchers"

type GuardUpdatePrograms struct {
	System string           `json:"system"`
	Search matchers.Program `json:"search"`
}
