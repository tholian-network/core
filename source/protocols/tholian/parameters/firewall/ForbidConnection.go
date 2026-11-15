package firewall

import "tholian-core/matchers"

type ForbidConnection struct {
	System     string              `json:"system"`
	Connection matchers.Connection `json:"connection"`
}
