package parameters

import "tholian-core/matchers"

type ScoutSearchUpdates struct {
	System string          `json:"system"`
	Search matchers.Update `json:"search"`
}
