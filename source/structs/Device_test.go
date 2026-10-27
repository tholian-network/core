package structs

import "testing"

func TestNewDevice(t *testing.T) {

	device := NewDevice("pci")

	if device.Bus != "pci" {
		t.Errorf("NewDevice() Bus = %q, expected %q", device.Bus, "pci")
	}

	if device.Name != "" || device.System != nil || device.Subsystem != nil {
		t.Errorf("NewDevice() = %+v, expected empty fields", device)
	}

	invalid := NewDevice("invalid")

	if invalid.Bus != "" {
		t.Errorf("NewDevice() accepted an invalid bus: %q", invalid.Bus)
	}

}

func TestDeviceSetBus(t *testing.T) {

	buses := []string{"hid", "i2c", "pci", "scsi", "usb", "other"}

	for _, bus := range buses {

		device := NewDevice(bus)

		if device.Bus != bus {
			t.Errorf("SetBus(%q) = %q", bus, device.Bus)
		}

	}

	device := NewDevice("usb")
	device.SetBus("invalid")

	if device.Bus != "usb" {
		t.Errorf("SetBus() accepted an invalid bus: %q", device.Bus)
	}

}

func TestDeviceSetters(t *testing.T) {

	device := NewDevice("pci")

	device.SetName("  Ethernet  ")
	device.SetSystem(" 8086 ", " 1234 ", " Intel ")
	device.SetSubsystem(" 8086 ", " 5678 ", " Sub ")

	if device.Name != "Ethernet" {
		t.Errorf("SetName() = %q, expected %q", device.Name, "Ethernet")
	}

	if device.System == nil {
		t.Fatal("SetSystem() did not set the system")
	}

	if device.System.Vendor != "8086" || device.System.Device != "1234" || device.System.Name != "Intel" {
		t.Errorf("SetSystem() = %+v, expected trimmed values", device.System)
	}

	if device.Subsystem == nil {
		t.Fatal("SetSubsystem() did not set the subsystem")
	}

	if device.Subsystem.Vendor != "8086" || device.Subsystem.Device != "5678" || device.Subsystem.Name != "Sub" {
		t.Errorf("SetSubsystem() = %+v, expected trimmed values", device.Subsystem)
	}

}

func TestDeviceHasSystemAndSubsystem(t *testing.T) {

	device := NewDevice("pci")

	if device.HasSystem() == true || device.HasSubsystem() == true {
		t.Fatal("HasSystem() or HasSubsystem() reported a missing value")
	}

	device.SetSystem("8086", "1234", "Intel")

	if device.HasSystem() == false {
		t.Fatal("HasSystem() did not find an existing system")
	}

	if device.HasSubsystem() == true {
		t.Fatal("HasSubsystem() found a missing subsystem")
	}

	device.SetSubsystem("8086", "5678", "Sub")

	if device.HasSubsystem() == false {
		t.Fatal("HasSubsystem() did not find an existing subsystem")
	}

}

func TestDeviceIsValid(t *testing.T) {

	cases := []struct {
		name   string
		device Device
		want   bool
	}{
		{"zero value", Device{}, false},
		{"name only", Device{Name: "Ethernet"}, false},
		{"pci without system and subsystem", Device{Name: "Ethernet", Bus: "pci"}, false},
		{"pci without optional subsystem", func() Device {
			value := NewDevice("pci")
			value.SetName("Ethernet")
			value.SetSystem("8086", "1234", "Intel")
			return value
		}(), true},
		{"pci without system but with subsystem", func() Device {
			value := NewDevice("pci")
			value.SetName("Ethernet")
			value.SetSubsystem("8086", "5678", "Sub")
			return value
		}(), false},
		{"pci with empty subsystem", func() Device {
			value := NewDevice("pci")
			value.SetName("Ethernet")
			value.SetSystem("8086", "1234", "Intel")
			value.Subsystem = &DeviceSystem{}
			return value
		}(), false},
		{"pci with partial subsystem", func() Device {
			value := NewDevice("pci")
			value.SetName("Ethernet")
			value.SetSystem("8086", "1234", "Intel")
			value.SetSubsystem("8086", "", "Sub")
			return value
		}(), false},
		{"pci valid", func() Device {
			value := NewDevice("pci")
			value.SetName("Ethernet")
			value.SetSystem("8086", "1234", "Intel")
			value.SetSubsystem("8086", "5678", "Sub")
			return value
		}(), true},
		{"usb without system", Device{Name: "Keyboard", Bus: "usb"}, false},
		{"usb with empty system", func() Device {
			value := NewDevice("usb")
			value.SetName("Keyboard")
			value.System = &DeviceSystem{}
			return value
		}(), false},
		{"usb with partial system", func() Device {
			value := NewDevice("usb")
			value.SetName("Keyboard")
			value.SetSystem("1234", "", "Keyboard")
			return value
		}(), false},
		{"usb valid", func() Device {
			value := NewDevice("usb")
			value.SetName("Keyboard")
			value.SetSystem("1234", "5678", "Keyboard")
			return value
		}(), true},
		{"unknown bus", func() Device {
			value := NewDevice("hid")
			value.SetName("Keyboard")
			value.SetSystem("1234", "5678", "Keyboard")
			return value
		}(), false},
	}

	for _, c := range cases {

		if c.device.IsValid() != c.want {
			t.Errorf("%s: IsValid() = %v, expected %v", c.name, c.device.IsValid(), c.want)
		}

	}

}

func TestDeviceEqual(t *testing.T) {

	pci := NewDevice("pci")
	pci.SetName("Ethernet")
	pci.SetSystem("8086", "1234", "Intel")
	pci.SetSubsystem("8086", "5678", "Sub")

	if pci.Equal(pci) == false {
		t.Fatal("a pci device should equal itself")
	}

	empty_pci := Device{Name: "Ethernet", Bus: "pci"}

	if empty_pci.Equal(Device{Name: "Ethernet", Bus: "pci"}) == false {
		t.Fatal("two equal pci devices without systems should be equal")
	}

	if empty_pci.Equal(pci) == true {
		t.Fatal("a pci device without a system should not equal a populated one")
	}

	cases := []struct {
		name  string
		value Device
	}{
		{"name", func() Device { value := pci; value.Name = "Other"; return value }()},
		{"bus", func() Device { value := pci; value.Bus = "usb"; return value }()},
		{"system missing", func() Device { value := pci; value.System = nil; return value }()},
		{"subsystem missing", func() Device { value := pci; value.Subsystem = nil; return value }()},
		{"system vendor", func() Device { value := pci; value.SetSystem("1234", "1234", "Intel"); return value }()},
		{"system device", func() Device { value := pci; value.SetSystem("8086", "9999", "Intel"); return value }()},
		{"subsystem vendor", func() Device { value := pci; value.SetSubsystem("1234", "5678", "Sub"); return value }()},
		{"subsystem device", func() Device { value := pci; value.SetSubsystem("8086", "9999", "Sub"); return value }()},
	}

	for _, c := range cases {

		if pci.Equal(c.value) == true {
			t.Errorf("Equal() returned true for a different %s", c.name)
		}

	}

	usb := NewDevice("usb")
	usb.SetName("Keyboard")
	usb.SetSystem("1234", "5678", "Keyboard")
	usb.SetSubsystem("1234", "9999", "Ignored")

	usb_other_subsystem := usb
	usb_other_subsystem.SetSubsystem("1234", "0000", "Other")

	if usb.Equal(usb_other_subsystem) == false {
		t.Fatal("Equal() should ignore the subsystem for usb devices")
	}

	other := NewDevice("other")
	other.SetName("Device")
	other.SetSystem("1234", "5678", "Device")

	other_system := other
	other_system.SetSystem("0000", "0000", "Other")

	if other.Equal(other_system) == false {
		t.Fatal("Equal() should ignore systems for unsupported buses")
	}

}

func TestDeviceHash(t *testing.T) {

	empty_device := Device{}

	if empty_device.Hash() != "" {
		t.Fatal("Hash() should be empty without a name")
	}

	pci := NewDevice("pci")
	pci.SetName("Ethernet")
	pci.SetSystem("8086", "1234", "Intel")
	pci.SetSubsystem("8086", "5678", "Sub")

	hash := pci.Hash()

	if hash == "" {
		t.Fatal("Hash() returned an empty string for a populated device")
	}

	if pci.Hash() != hash {
		t.Fatal("Hash() is not deterministic")
	}

	cases := []struct {
		name  string
		value Device
	}{
		{"name", func() Device { value := pci; value.Name = "Other"; return value }()},
		{"bus", func() Device { value := pci; value.Bus = "usb"; return value }()},
		{"system missing", func() Device { value := pci; value.System = nil; return value }()},
		{"subsystem missing", func() Device { value := pci; value.Subsystem = nil; return value }()},
		{"system vendor", func() Device { value := pci; value.SetSystem("1234", "1234", "Intel"); return value }()},
		{"subsystem device", func() Device { value := pci; value.SetSubsystem("8086", "9999", "Sub"); return value }()},
	}

	for _, c := range cases {

		if c.value.Hash() == hash {
			t.Errorf("Hash() did not change for a different %s", c.name)
		}

	}

	usb := NewDevice("usb")
	usb.SetName("Keyboard")
	usb.SetSystem("1234", "5678", "Keyboard")

	if usb.Hash() == "" {
		t.Fatal("Hash() returned an empty string for a populated usb device")
	}

	usb_system := usb
	usb_system.SetSystem("0000", "0000", "Other")

	if usb_system.Hash() == usb.Hash() {
		t.Fatal("Hash() did not change for a different usb system")
	}

}
