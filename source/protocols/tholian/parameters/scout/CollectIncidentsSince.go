package scout

import "tholian-core/matchers"

type CollectIncidentsSince struct {
	System string            `json:"system"`
	Search matchers.Datetime `json:"search"`
}
