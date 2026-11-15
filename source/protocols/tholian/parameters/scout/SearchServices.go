package scout

import "tholian-core/matchers"

type SearchServices struct {
	System string           `json:"system"`
	Search matchers.Program `json:"search"`
}
