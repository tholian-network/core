package parameters

import "tholian-core/matchers"

type GuardSearchVulnerabilities struct {
	System string                 `json:"system"`
	Search matchers.Vulnerability `json:"search"`
}
