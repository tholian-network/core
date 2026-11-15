package scout

import "tholian-core/structs"

type UpdatePackage struct {
	Updates []structs.Update `json:"updates"`
}
