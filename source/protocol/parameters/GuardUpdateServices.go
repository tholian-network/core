package parameters

import "tholian-core/matchers"

type GuardUpdateServices struct {
	System string           `json:"system"`
	Search matchers.Program `json:"search"`
}
