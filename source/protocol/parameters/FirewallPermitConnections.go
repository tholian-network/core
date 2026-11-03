package parameters

import "tholian-core/matchers"

type FirewallPermitConnections struct {
	System      string                `json:"system"`
	Connections []matchers.Connection `json:"connections"`
}
