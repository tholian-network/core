package scout

import "tholian-core/structs"

type UpdatePackages struct {
	Updates []structs.Update `json:"updates"`
}
