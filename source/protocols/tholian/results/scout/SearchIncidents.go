package scout

import "tholian-core/structs"

type SearchIncidents struct {
	Incidents []structs.Incident `json:"incidents"`
}
