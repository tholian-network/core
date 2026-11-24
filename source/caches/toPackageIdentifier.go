package caches

import "tholian-core/structs"

func toPackageIdentifier(value structs.Package) string {
	return value.Manager.String() + ":" + value.Vendor + ":" + value.Name + ":" + value.Version.String() + ":" + value.Architecture.String()
}
