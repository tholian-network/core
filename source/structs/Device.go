package structs

import "strings"

type DeviceSystem struct {
	Name   string `json:"name"`
	Vendor string `json:"vendor"`
	Device string `json:"device"`
}

type Device struct {
	Name      string        `json:"name"`
	Bus       string        `json:"bus"`
	System    *DeviceSystem `json:"system,omitempty"`
	Subsystem *DeviceSystem `json:"subsystem,omitempty"`
}

func NewDevice(bus string) Device {

	var device Device

	device.SetBus(bus)

	return device

}

func (device *Device) Equal(value Device) bool {

	if device.Name == value.Name {

		if device.Bus == "pci" && value.Bus == "pci" {

			if device.System.Vendor == value.System.Vendor && device.System.Device == value.System.Device {

				if device.Subsystem.Vendor == value.Subsystem.Vendor && device.Subsystem.Device == value.Subsystem.Device {
					return true
				}

			}

		} else if device.Bus == "usb" && value.Bus == "usb" {

			if device.System.Vendor == value.System.Vendor && device.System.Device == value.System.Device {
				return true
			}

		}

	}

	return false

}

func (device *Device) IsValid() bool {

	if device.Name != "" {

		if device.Bus == "pci" {

			if device.System.Vendor != "" && device.System.Device != "" {

				if device.Subsystem.Vendor != "" && device.Subsystem.Device != "" {
					return true
				}

			}

		} else if device.Bus == "usb" {

			if device.System.Vendor != "" && device.System.Device != "" {
				return true
			}

		}

	}

	return false

}

func (device *Device) SetName(value string) {
	device.Name = strings.TrimSpace(value)
}

func (device *Device) SetSystem(value_vendor string, value_device string, value_name string) {

	var system DeviceSystem

	system.Vendor = strings.TrimSpace(value_vendor)
	system.Device = strings.TrimSpace(value_device)
	system.Name = strings.TrimSpace(value_name)

	device.System = &system

}

func (device *Device) SetSubsystem(value_vendor string, value_device string, value_name string) {

	var system DeviceSystem

	system.Vendor = strings.TrimSpace(value_vendor)
	system.Device = strings.TrimSpace(value_device)
	system.Name = strings.TrimSpace(value_name)

	device.Subsystem = &system

}

func (device *Device) SetBus(value string) {

	if value == "hid" {
		device.Bus = "hid"
	} else if value == "i2c" {
		device.Bus = "i2c"
	} else if value == "pci" {
		device.Bus = "pci"
	} else if value == "scsi" {
		device.Bus = "scsi"
	} else if value == "usb" {
		device.Bus = "usb"
	} else if value == "other" {
		device.Bus = "other"
	}

}
