package parameters

import "tholian-core/matchers"

type ScoutSearchServices struct {
	System string           `json:"system"`
	Search matchers.Program `json:"search"`
}
