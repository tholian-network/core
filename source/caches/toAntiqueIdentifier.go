package caches

import "tholian-core/structs"

func toAntiqueIdentifier(value structs.Antique) string {
	return value.Manager.String() + ":" + value.Vendor + ":" + value.Name + ":" + value.Version.String() + ":" + value.Architecture.String()
}
