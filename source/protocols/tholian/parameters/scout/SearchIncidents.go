package scout

import "tholian-core/matchers"

type SearchIncidents struct {
	System string            `json:"system"`
	Search matchers.Incident `json:"search"`
}
