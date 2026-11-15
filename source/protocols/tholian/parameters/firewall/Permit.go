package firewall

type Permit struct {
	System string `json:"system"`
	Target string `json:"target"`
}
