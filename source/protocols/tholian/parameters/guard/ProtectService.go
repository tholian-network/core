package guard

import "tholian-core/matchers"

type ProtectService struct {
	System string           `json:"system"`
	Search matchers.Program `json:"search"`
}
