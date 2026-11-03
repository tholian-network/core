package parameters

import "tholian-core/matchers"

type ScoutUpdatePackages struct {
	System   string             `json:"system"`
	Searches []matchers.Package `json:"searches"`
}
