package parameters

import "tholian-core/matchers"

type FirewallPermitConnection struct {
	System     string              `json:"system"`
	Connection matchers.Connection `json:"connection"`
}
