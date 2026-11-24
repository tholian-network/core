package caches

import "tholian-core/structs"

func toResponseIdentifier(value structs.Response) string {
	return value.System + ":" + value.Type + ":" + value.Datetime.String()
}
