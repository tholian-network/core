package firewall

import "tholian-core/matchers"

type PermitConnection struct {
	System     string              `json:"system"`
	Connection matchers.Connection `json:"connection"`
}
