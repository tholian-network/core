package caches

import "tholian-core/structs"

func toDeviceIdentifier(value structs.Device) string {

	result := ""

	if value.Bus != "" && value.System != nil && value.System.Vendor != "" && value.System.Device != "" {

		if value.Subsystem != nil && value.Subsystem.Vendor != "" && value.Subsystem.Device != "" {
			result = value.Bus + ":" + value.System.Vendor + "-" + value.System.Device + ":" + value.Subsystem.Vendor + "-" + value.Subsystem.Device
		} else {
			result = value.Bus + ":" + value.System.Vendor + "-" + value.System.Device + ":0000-0000"
		}

	}

	return result

}
