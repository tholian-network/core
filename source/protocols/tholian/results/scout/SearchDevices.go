package scout

import "tholian-core/structs"

type SearchDevices struct {
	Devices []structs.Device `json:"devices"`
}
