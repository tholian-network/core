package scout

import "tholian-core/structs"

type CollectIncidentsSince struct {
	Incidents []structs.Incident `json:"incidents"`
}
