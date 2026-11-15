package scout

import "tholian-core/structs"

type CollectUpdates struct {
	Updates []structs.Update `json:"updates"`
}
