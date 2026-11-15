package guard

import "tholian-core/matchers"

type ProtectServices struct {
	System string           `json:"system"`
	Search matchers.Program `json:"search"`
}
