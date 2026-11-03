package parameters

import "tholian-core/matchers"

type FirewallForbidConnections struct {
	System      string                `json:"system"`
	Connections []matchers.Connection `json:"connections"`
}
