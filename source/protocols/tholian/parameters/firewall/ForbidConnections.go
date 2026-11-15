package firewall

import "tholian-core/matchers"

type ForbidConnections struct {
	System      string                `json:"system"`
	Connections []matchers.Connection `json:"connections"`
}
