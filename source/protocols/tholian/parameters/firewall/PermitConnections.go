package firewall

import "tholian-core/matchers"

type PermitConnections struct {
	System      string                `json:"system"`
	Connections []matchers.Connection `json:"connections"`
}
