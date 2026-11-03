package parameters

import "tholian-core/matchers"

type ScoutSearchPrograms struct {
	System string           `json:"system"`
	Search matchers.Program `json:"search"`
}
