package scout

import "tholian-core/structs"

type SearchPackages struct {
	Packages []structs.Package `json:"packages"`
}
