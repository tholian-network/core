package parameters

import "tholian-core/matchers"

type ScoutSearchDrives struct {
	System string         `json:"system"`
	Search matchers.Drive `json:"search"`
}
