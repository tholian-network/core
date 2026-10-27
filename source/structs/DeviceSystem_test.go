package structs

import "testing"

func TestNewDeviceSystem(t *testing.T) {

	system := NewDeviceSystem("  Intel Ethernet  ", " 8086 ", " 1234 ")

	if system.Name != "Intel Ethernet" {
		t.Errorf("NewDeviceSystem() Name = %q, expected %q", system.Name, "Intel Ethernet")
	}

	if system.Vendor != "8086" {
		t.Errorf("NewDeviceSystem() Vendor = %q, expected %q", system.Vendor, "8086")
	}

	if system.Device != "1234" {
		t.Errorf("NewDeviceSystem() Device = %q, expected %q", system.Device, "1234")
	}

}

func TestDeviceSystemIsValid(t *testing.T) {

	cases := []struct {
		name   string
		system DeviceSystem
		want   bool
	}{
		{"zero value", DeviceSystem{}, false},
		{"name only", DeviceSystem{Name: "Intel Ethernet"}, false},
		{"missing name", DeviceSystem{Vendor: "8086", Device: "1234"}, false},
		{"missing vendor", DeviceSystem{Name: "Intel Ethernet", Device: "1234"}, false},
		{"missing device", DeviceSystem{Name: "Intel Ethernet", Vendor: "8086"}, false},
		{"valid", DeviceSystem{Name: "Intel Ethernet", Vendor: "8086", Device: "1234"}, true},
	}

	for _, c := range cases {

		if c.system.IsValid() != c.want {
			t.Errorf("%s: IsValid() = %v, expected %v", c.name, c.system.IsValid(), c.want)
		}

	}

}

func TestDeviceSystemEqual(t *testing.T) {

	base := DeviceSystem{Name: "Intel Ethernet", Vendor: "8086", Device: "1234"}

	if base.Equal(base) == false {
		t.Fatal("a device system should equal itself")
	}

	cases := []struct {
		name  string
		value DeviceSystem
	}{
		{"name", DeviceSystem{Name: "Other", Vendor: "8086", Device: "1234"}},
		{"vendor", DeviceSystem{Name: "Intel Ethernet", Vendor: "1234", Device: "1234"}},
		{"device", DeviceSystem{Name: "Intel Ethernet", Vendor: "8086", Device: "5678"}},
	}

	for _, c := range cases {

		if base.Equal(c.value) == true {
			t.Errorf("Equal() returned true for a different %s", c.name)
		}

	}

}

func TestDeviceSystemHash(t *testing.T) {

	empty_system := DeviceSystem{}

	if empty_system.Hash() != "" {
		t.Fatal("Hash() should be empty for an empty device system")
	}

	base := DeviceSystem{Name: "Intel Ethernet", Vendor: "8086", Device: "1234"}
	hash := base.Hash()

	if hash == "" {
		t.Fatal("Hash() returned an empty string for a populated device system")
	}

	if base.Hash() != hash {
		t.Fatal("Hash() is not deterministic")
	}

	cases := []struct {
		name  string
		value DeviceSystem
	}{
		{"name", DeviceSystem{Name: "Other", Vendor: "8086", Device: "1234"}},
		{"vendor", DeviceSystem{Name: "Intel Ethernet", Vendor: "1234", Device: "1234"}},
		{"device", DeviceSystem{Name: "Intel Ethernet", Vendor: "8086", Device: "5678"}},
	}

	for _, c := range cases {

		if c.value.Hash() == hash {
			t.Errorf("Hash() did not change for a different %s", c.name)
		}

	}

}
