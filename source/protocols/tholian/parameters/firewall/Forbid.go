package firewall

type Forbid struct {
	System string `json:"system"`
	Target string `json:"target"`
}
