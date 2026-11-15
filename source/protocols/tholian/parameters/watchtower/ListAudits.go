package watchtower

type ListAudits struct {
	System string `json:"system,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}
