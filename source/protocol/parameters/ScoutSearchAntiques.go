package parameters

import "tholian-core/matchers"

type ScoutSearchAntiques struct {
	System string           `json:"system"`
	Search matchers.Antique `json:"search"`
}
