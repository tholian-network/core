package parameters

import "tholian-core/matchers"

type ScoutSearchUsers struct {
	System string        `json:"system"`
	Search matchers.User `json:"search"`
}
