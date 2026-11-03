package parameters

import "tholian-core/matchers"

type ScoutSearchIncidents struct {
	System string            `json:"system"`
	Search matchers.Incident `json:"search"`
}
