package guard

import "tholian-core/matchers"

type UpdatePrograms struct {
	System string           `json:"system"`
	Search matchers.Program `json:"search"`
}
