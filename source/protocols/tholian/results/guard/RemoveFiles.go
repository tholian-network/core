package guard

import "tholian-core/matchers"

type RemoveFiles struct {
	Files []structs.File `json:"files"`
}
