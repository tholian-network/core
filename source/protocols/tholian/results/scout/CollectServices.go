package scout

import "tholian-core/structs"

type CollectServices struct {
	Services []structs.Program `json:"services"`
}
