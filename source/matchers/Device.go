package matchers

import "encoding/binary"
import "encoding/hex"
import "hash/crc32"
import "strings"

type Device struct {
	Name   string `json:"name"`
	Vendor string `json:"vendor"`
	Device string `json:"device"`
	Bus    string `json:"bus"`
}

func NewDevice() Device {

	var device Device

	device.Name = Any
	device.Vendor = Any
	device.Device = Any
	device.Bus = Any

	return device

}

func ToDevice(value string) Device {

	var device Device

	device.Name = Any
	device.Vendor = Any
	device.Device = Any
	device.Bus = Any

	device.SetName(value)

	return device

}

func (device *Device) Equal(value Device) bool {

	if device.Name == value.Name &&
		device.Vendor == value.Vendor &&
		device.Device == value.Device &&
		device.Bus == value.Bus {
		return true
	}

	return false

}

func (device *Device) IsValid() bool {

	if device.Name != Any ||
		device.Vendor != Any ||
		device.Device != Any ||
		device.Bus != Any {
		return true
	}

	return false

}

func (device *Device) Matches(name string, vendor string, device_ string, bus string) bool {
	return device.MatchesName(name) && device.MatchesVendor(vendor) && device.MatchesDevice(device_) && device.MatchesBus(bus)
}

func (device *Device) MatchesBus(value string) bool {

	if device.Bus == value {
		return true
	} else if device.Bus == Any {
		return true
	}

	return false

}

func (device *Device) MatchesDevice(value string) bool {

	if device.Device == value {
		return true
	} else if device.Device == Any {
		return true
	}

	return false

}

func (device *Device) MatchesName(value string) bool {

	if device.Name == value {
		return true
	} else if device.Name == Any {
		return true
	}

	return false

}

func (device *Device) MatchesVendor(value string) bool {

	if device.Vendor == value {
		return true
	} else if device.Vendor == Any {
		return true
	}

	return false

}

func (device *Device) SetBus(value string) {

	if value == "any" {
		device.Bus = Any
	} else if value == "hid" {
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

func (device *Device) SetDevice(value string) {
	device.Device = strings.TrimSpace(value)
}

func (device *Device) SetName(value string) {
	device.Name = strings.TrimSpace(value)
}

func (device *Device) SetVendor(value string) {
	device.Vendor = strings.TrimSpace(value)
}

func (device *Device) Hash() string {

	hash := ""

	if device.Name != "" {

		checksum := crc32.ChecksumIEEE([]byte(strings.Join([]string{
			device.Name,
			device.Vendor,
			device.Device,
			device.Bus,
		}, "-")))

		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, checksum)
		hash = hex.EncodeToString(tmp)

	}

	return hash

}
