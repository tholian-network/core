package scout

import "tholian-core/matchers"

type SearchDrives struct {
	System string         `json:"system"`
	Search matchers.Drive `json:"search"`
}
