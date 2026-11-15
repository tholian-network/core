package scout

import "tholian-core/structs"

type CollectPackages struct {
	Packages []structs.Package `json:"packages"`
}
