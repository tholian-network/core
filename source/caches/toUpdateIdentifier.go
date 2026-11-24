package caches

import "tholian-core/structs"

func toUpdateIdentifier(value structs.Update) string {
	return value.Manager.String() + ":" + value.Vendor + ":" + value.Name + ":" + value.Version.String() + ":" + value.Architecture.String()
}
