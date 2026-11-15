package scout

import "tholian-core/matchers"

type SearchDevices struct {
	System string          `json:"system"`
	Search matchers.Device `json:"search"`
}
