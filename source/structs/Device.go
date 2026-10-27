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

func (system *DeviceSystem) Equal(value DeviceSystem) bool {

	if system.Name == value.Name &&
		system.Vendor == value.Vendor &&
		system.Device == value.Device {
		return true
	}

	return false

}

func (system *DeviceSystem) Hash() string {

	hash := ""

	if system.Name != "" || system.Vendor != "" || system.Device != "" {
		hash = toHash(system.Name, system.Vendor, system.Device)
	}

	return hash

}

func NewDevice(bus string) Device {

	var device Device

	device.SetBus(bus)

	return device

}

func (device *Device) Equal(value Device) bool {

	if device.Name != value.Name || device.Bus != value.Bus {
		return false
	}

	if device.Bus == "pci" {

		if (device.System == nil) != (value.System == nil) {
			return false
		}

		if device.System != nil {
			if device.System.Vendor != value.System.Vendor || device.System.Device != value.System.Device {
				return false
			}
		}

		if (device.Subsystem == nil) != (value.Subsystem == nil) {
			return false
		}

		if device.Subsystem != nil {
			if device.Subsystem.Vendor != value.Subsystem.Vendor || device.Subsystem.Device != value.Subsystem.Device {
				return false
			}
		}

		return true

	} else if device.Bus == "usb" {

		if (device.System == nil) != (value.System == nil) {
			return false
		}

		if device.System != nil {
			if device.System.Vendor != value.System.Vendor || device.System.Device != value.System.Device {
				return false
			}
		}

		return true

	}

	return true

}

func (device *Device) IsValid() bool {

	if device.Name != "" {

		if device.Bus == "pci" {

			// Subsystem is optional: some devices have no subsystem data at all.
			// Subsystem == nil is valid, but a present Subsystem must be complete.
			if device.System != nil && device.System.Vendor != "" && device.System.Device != "" {

				if device.Subsystem == nil {
					return true
				}

				if device.Subsystem.Vendor != "" && device.Subsystem.Device != "" {
					return true
				}

			}

		} else if device.Bus == "usb" {

			if device.System != nil && device.System.Vendor != "" && device.System.Device != "" {
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

func (device *Device) Hash() string {

	hash := ""

	if device.Name != "" {

		values := make([]string, 0)

		values = append(values, device.Name)
		values = append(values, device.Bus)

		if device.Bus == "pci" {

			if device.System != nil {
				values = append(values, device.System.Vendor)
				values = append(values, device.System.Device)
			} else {
				values = append(values, "")
				values = append(values, "")
			}

			if device.Subsystem != nil {
				values = append(values, device.Subsystem.Vendor)
				values = append(values, device.Subsystem.Device)
			} else {
				values = append(values, "")
				values = append(values, "")
			}

		} else if device.Bus == "usb" {

			if device.System != nil {
				values = append(values, device.System.Vendor)
				values = append(values, device.System.Device)
			} else {
				values = append(values, "")
				values = append(values, "")
			}

		}

		hash = toHash(values...)

	}

	return hash

}
