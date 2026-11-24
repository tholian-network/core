package caches

import "tholian-core/structs"

func toMitigationIdentifier(value structs.Mitigation) string {

	result := ""
	result += value.Incident.System + ":" + value.Incident.Type + ":" + value.Incident.Datetime.String()
	result += "|"
	result += value.Response.System + ":" + value.Response.Type + ":" + value.Response.Datetime.String()

	return result

}
