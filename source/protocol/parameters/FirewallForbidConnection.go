package parameters

import "tholian-core/matchers"

type FirewallForbidConnection struct {
	System     string              `json:"system"`
	Connection matchers.Connection `json:"connection"`
}
