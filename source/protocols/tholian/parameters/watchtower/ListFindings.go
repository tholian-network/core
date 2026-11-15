package watchtower

type ListFindings struct {
	System string `json:"system,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}
