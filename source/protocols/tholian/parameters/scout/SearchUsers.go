package scout

import "tholian-core/matchers"

type SearchUsers struct {
	System string        `json:"system"`
	Search matchers.User `json:"search"`
}
