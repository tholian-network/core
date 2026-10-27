package structs

import "strings"

type DeviceSystem struct {
	Name   string `json:"name"`
	Vendor string `json:"vendor"`
	Device string `json:"device"`
}

func NewDeviceSystem(name string, vendor string, device string) DeviceSystem {

	var system DeviceSystem

	system.Name = strings.TrimSpace(name)
	system.Vendor = strings.TrimSpace(vendor)
	system.Device = strings.TrimSpace(device)

	return system

}

func (system *DeviceSystem) Equal(value DeviceSystem) bool {

	if system.Name == value.Name &&
		system.Vendor == value.Vendor &&
		system.Device == value.Device {
		return true
	}

	return false

}

func (system *DeviceSystem) IsValid() bool {

	if system.Name != "" &&
		system.Vendor != "" &&
		system.Device != "" {
		return true
	}

	return false

}

func (system *DeviceSystem) Hash() string {

	hash := ""

	if system.Name != "" || system.Vendor != "" || system.Device != "" {

		hash = toHash(
			system.Name,
			system.Vendor,
			system.Device,
		)

	}

	return hash

}
