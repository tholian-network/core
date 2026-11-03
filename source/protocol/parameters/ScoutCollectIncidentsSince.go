package parameters

import "tholian-core/matchers"

type ScoutCollectIncidentsSince struct {
	System string            `json:"system"`
	Search matchers.Datetime `json:"search"`
}
