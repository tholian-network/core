package guard

import "tholian-core/matchers"

type UpdateServices struct {
	System string           `json:"system"`
	Search matchers.Program `json:"search"`
}
