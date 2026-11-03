package parameters

import "tholian-core/matchers"

type ScoutSearchDevices struct {
	System string          `json:"system"`
	Search matchers.Device `json:"search"`
}
