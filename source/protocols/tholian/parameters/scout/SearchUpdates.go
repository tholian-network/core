package scout

import "tholian-core/matchers"

type SearchUpdates struct {
	System string          `json:"system"`
	Search matchers.Update `json:"search"`
}
