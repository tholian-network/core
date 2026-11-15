package scout

import "tholian-core/structs"

type CollectIncidents struct {
	Incidents []structs.Incident `json:"incidents"`
}
