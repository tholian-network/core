package scout

import "tholian-core/structs"

type CollectDevices struct {
	Devices []structs.Device `json:"devices"`
}
