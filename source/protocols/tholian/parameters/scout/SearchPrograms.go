package scout

import "tholian-core/matchers"

type SearchPrograms struct {
	System string           `json:"system"`
	Search matchers.Program `json:"search"`
}
