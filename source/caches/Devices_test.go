package caches

import "tholian-core/matchers"
import "tholian-core/structs"
import "testing"

func mockupDevices() *Devices {

	devices := NewDevices()
	devices.Add(structs.Device{
		Name: "Leet Networks Network Card",
		Bus:  "pci",
		System: &structs.DeviceSystem{
			Name:   "Network Card",
			Device: "0101",
			Vendor: "1337",
		},
		Subsystem: nil,
	})
	devices.Add(structs.Device{
		Name: "Leet Networks Network Card v2.0",
		Bus:  "pci",
		System: &structs.DeviceSystem{
			Name:   "Network Card",
			Device: "0101",
			Vendor: "1337",
		},
		Subsystem: &structs.DeviceSystem{
			Name:   "Network Card v2.0",
			Device: "0102",
			Vendor: "1337",
		},
	})
	devices.Add(structs.Device{
		Name: "Universal Systems USB 2.0 Hub",
		Bus:  "usb",
		System: &structs.DeviceSystem{
			Name:   "USB Hub",
			Device: "0103",
			Vendor: "1338",
		},
		Subsystem: nil,
	})
	devices.Add(structs.Device{
		Name: "Universal Systems USB 3.0 Hub",
		Bus:  "usb",
		System: &structs.DeviceSystem{
			Name:   "USB Hub",
			Device: "0103",
			Vendor: "1338",
		},
		Subsystem: &structs.DeviceSystem{
			Name:   "USB Hub 3.0",
			Device: "0104",
			Vendor: "1338",
		},
	})

	return devices

}

func TestDevices(t *testing.T) {

	t.Run("Add()", func(t *testing.T) {

		devices := mockupDevices()

		device1 := devices.Get("pci:1338-0001:0000-0000")

		if device1 != nil {
			t.Errorf("Expected %s to be nil", device1.Name)
		}

		devices.Add(structs.Device{
			Name: "Super Duper Hyper Network Card v1.0",
			Bus:  "pci",
			System: &structs.DeviceSystem{
				Name:   "Hyper Network Card",
				Device: "0001",
				Vendor: "1338",
			},
			Subsystem: nil,
		})

		device2 := devices.Get("pci:1338-0001:0000-0000")

		if device2 == nil {
			t.Errorf("Expected nil to be %s", "Super Duper Hyper Network Card v1.0")
		} else if device2.Name != "Super Duper Hyper Network Card v1.0" {
			t.Errorf("Expected %s to be %s", device2.Name, "Super Duper Hyper Network Card v1.0")
		}

	})

	t.Run("Get()", func(t *testing.T) {

		devices := mockupDevices()

		device1 := devices.Get("pci:1337-0101:1337-0102")
		device2 := devices.Get("usb:1338-0103:1338-0104")

		if device1 == nil {
			t.Errorf("Expected nil to be %s", "Leet Networks Network Card v2.0")
		} else if device1.Name != "Leet Networks Network Card v2.0" {
			t.Errorf("Expected %s to be %s", device1.Name, "Leet Networks Network Card v2.0")
		}

		if device2 == nil {
			t.Errorf("Expected nil to be %s", "Universal Systems USB 3.0 Hub")
		} else if device2.Name != "Universal Systems USB 3.0 Hub" {
			t.Errorf("Expected %s to be %s", device2.Name, "Universal Systems USB 3.0 Hub")
		}

	})

	t.Run("Query()", func(t *testing.T) {

		devices := mockupDevices()

		found1 := devices.Query(matchers.Device{
			Name: "any",
			Vendor: "1337",
			Device: "any",
			Bus: "pci",
		})

		found2 := devices.Query(matchers.Device{
			Name: "any",
			Vendor: "1337",
			Device: "0101",
			Bus: "pci",
		})

		found3 := devices.Query(matchers.Device{
			Name: "any",
			Vendor: "1338",
			Device: "any",
			Bus: "usb",
		})

		found4 := devices.Query(matchers.Device{
			Name: "any",
			Vendor: "1338",
			Device: "0103",
			Bus: "usb",
		})

		found5 := devices.Query(matchers.Device{
			Name: "any",
			Vendor: "any",
			Device: "any",
			Bus: "pci",
		})

		found6 := devices.Query(matchers.Device{
			Name: "any",
			Vendor: "any",
			Device: "any",
			Bus: "usb",
		})

		if len(found1) == 2 {

			if found1[0].Name != "Leet Networks Network Card" {
				t.Errorf("Expected %s to be %s", found1[0].Name, "Leet Networks Network Card")
			}

			if found1[1].Name != "Leet Networks Network Card v2.0" {
				t.Errorf("Expected %s to be %s", found1[1].Name, "Leet Networks Network Card v2.0")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 2, "Vendor=1337,Device=any,Bus=pci")
		}

		if len(found2) == 2 {

			if found2[0].Name != "Leet Networks Network Card" {
				t.Errorf("Expected %s to be %s", found2[0].Name, "Leet Networks Network Card")
			}

			if found2[1].Name != "Leet Networks Network Card v2.0" {
				t.Errorf("Expected %s to be %s", found2[1].Name, "Leet Networks Network Card v2.0")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 2, "Vendor=1337,Device=0101,Bus=pci")
		}

		if len(found3) == 2 {

			if found3[0].Name != "Universal Systems USB 2.0 Hub" {
				t.Errorf("Expected %s to be %s", found3[0].Name, "Universal Systems USB 2.0 Hub")
			}

			if found3[1].Name != "Universal Systems USB 3.0 Hub" {
				t.Errorf("Expected %s to be %s", found3[1].Name, "Universal Systems USB 3.0 Hub")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found3), 2, "Vendor=1338,Device=any,Bus=usb")
		}

		if len(found4) == 2 {

			if found4[0].Name != "Universal Systems USB 2.0 Hub" {
				t.Errorf("Expected %s to be %s", found4[0].Name, "Universal Systems USB 2.0 Hub")
			}

			if found4[1].Name != "Universal Systems USB 3.0 Hub" {
				t.Errorf("Expected %s to be %s", found4[1].Name, "Universal Systems USB 3.0 Hub")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found4), 2, "Vendor=1338,Device=0103,Bus=usb")
		}

		if len(found5) != 2 {
			t.Errorf("Expected %d results to be %d for query %s", len(found5), 2, "Bus=pci")
		}

		if len(found6) != 2 {
			t.Errorf("Expected %d results to be %d for query %s", len(found6), 2, "Bus=usb")
		}

	})

	t.Run("QueryBySubsystem()", func(t *testing.T) {

		devices := mockupDevices()

		found1 := devices.QueryBySubsystem(matchers.Device{
			Name: "any",
			Vendor: "1337",
			Device: "any",
			Bus: "pci",
		})

		found2 := devices.QueryBySubsystem(matchers.Device{
			Name: "any",
			Vendor: "1337",
			Device: "0102",
			Bus: "pci",
		})

		found3 := devices.QueryBySubsystem(matchers.Device{
			Name: "any",
			Vendor: "1338",
			Device: "any",
			Bus: "usb",
		})

		found4 := devices.QueryBySubsystem(matchers.Device{
			Name: "any",
			Vendor: "1338",
			Device: "0104",
			Bus: "usb",
		})

		if len(found1) == 1 {

			if found1[0].Name != "Leet Networks Network Card v2.0" {
				t.Errorf("Expected %s to be %s", found1[0].Name, "Leet Networks Network Card v2.0")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 1, "Vendor=1337,Bus=pci")
		}

		if len(found2) == 1 {

			if found2[0].Name != "Leet Networks Network Card v2.0" {
				t.Errorf("Expected %s to be %s", found2[0].Name, "Leet Networks Network Card v2.0")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 1, "Vendor=1337,Device=0102,Bus=pci")
		}

		if len(found3) == 1 {

			if found3[0].Name != "Universal Systems USB 3.0 Hub" {
				t.Errorf("Expected %s to be %s", found3[0].Name, "Universal Systems USB 3.0 Hub")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found3), 1, "Vendor=1338,Bus=usb")
		}

		if len(found4) == 1 {

			if found4[0].Name != "Universal Systems USB 3.0 Hub" {
				t.Errorf("Expected %s to be %s", found4[0].Name, "Universal Systems USB 3.0 Hub")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found4), 1, "Vendor=1338,Device=0104,Bus=usb")
		}

	})

	t.Run("Remove()", func(t *testing.T) {

		devices := mockupDevices()

		device1 := devices.Get("usb:1338-0103:1338-0104")

		if device1 == nil {
			t.Errorf("Expected nil to be %s", "Universal Systems USB 3.0 Hub")
		} else if device1.Name != "Universal Systems USB 3.0 Hub" {
			t.Errorf("Expected %s to be %s", device1.Name, "Universal Systems USB 3.0 Hub")
		}

		devices.Remove("usb:1338-0103:1338-0104")

		device2 := devices.Get("usb:1338-0103:1338-0104")

		if device2 != nil {
			t.Errorf("Expected %s to be nil", device2.Name)
		}

	})

}
