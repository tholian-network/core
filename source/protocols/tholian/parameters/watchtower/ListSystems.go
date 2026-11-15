package watchtower

type ListSystems struct {
	System string `json:"system,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}
