package scout

import "tholian-core/matchers"

type SearchAntiques struct {
	System string           `json:"system"`
	Search matchers.Antique `json:"search"`
}
