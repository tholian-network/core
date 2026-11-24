package caches

import "tholian-core/structs"

func toIncidentIdentifier(value structs.Incident) string {
	return value.System + ":" + value.Type + ":" + value.Datetime.String()
}
