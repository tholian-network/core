package structs

import "testing"

func TestNewDrive(t *testing.T) {

	drive := NewDrive("  root  ", "local")

	if drive.Name != "root" {
		t.Errorf("NewDrive() Name = %q, expected %q", drive.Name, "root")
	}

	if drive.Type != "local" {
		t.Errorf("NewDrive() Type = %q, expected %q", drive.Type, "local")
	}

	if drive.Size != 0 || drive.Free != 0 || drive.Used != 0 || drive.Mountpoint != "" {
		t.Errorf("NewDrive() = %+v, expected zero size, free, used and mountpoint", drive)
	}

	remote := NewDrive("data", "remote")

	if remote.Type != "remote" {
		t.Errorf("NewDrive() Type = %q, expected %q", remote.Type, "remote")
	}

	invalid := NewDrive("data", "invalid")

	if invalid.Type != "" {
		t.Errorf("NewDrive() accepted an invalid type: %q", invalid.Type)
	}

}

func TestDriveSetters(t *testing.T) {

	drive := NewDrive("root", "local")

	drive.SetName("  data  ")
	drive.SetMountpoint("  /mnt/data  ")
	drive.SetSize(1000)
	drive.SetFree(400)
	drive.SetUsed(600)

	if drive.Name != "data" {
		t.Errorf("SetName() = %q, expected %q", drive.Name, "data")
	}

	if drive.Mountpoint != "/mnt/data" {
		t.Errorf("SetMountpoint() = %q, expected %q", drive.Mountpoint, "/mnt/data")
	}

	if drive.Size != 1000 || drive.Free != 400 || drive.Used != 600 {
		t.Errorf("SetSize(), SetFree(), SetUsed() = %d, %d, %d", drive.Size, drive.Free, drive.Used)
	}

	drive.SetType("remote")

	if drive.Type != "remote" {
		t.Errorf("SetType() = %q, expected %q", drive.Type, "remote")
	}

	drive.SetType("invalid")

	if drive.Type != "remote" {
		t.Errorf("SetType() accepted an invalid type: %q", drive.Type)
	}

}

func TestDriveIsValid(t *testing.T) {

	valid := func() Drive {

		drive := NewDrive("root", "local")
		drive.SetMountpoint("/")
		drive.SetSize(1000)
		drive.SetFree(400)
		drive.SetUsed(600)

		return drive

	}

	cases := []struct {
		name  string
		drive Drive
		want  bool
	}{
		{"default", Drive{}, false},
		{"valid", valid(), true},
		{"missing name", func() Drive { value := valid(); value.Name = ""; return value }(), false},
		{"relative mountpoint", func() Drive { value := valid(); value.Mountpoint = "mnt"; return value }(), false},
		{"empty mountpoint", func() Drive { value := valid(); value.Mountpoint = ""; return value }(), false},
		{"zero size", func() Drive { value := valid(); value.Size = 0; return value }(), false},
		{"zero free", func() Drive { value := valid(); value.Free = 0; return value }(), false},
		{"zero used", func() Drive { value := valid(); value.Used = 0; return value }(), false},
	}

	for _, c := range cases {

		if c.drive.IsValid() != c.want {
			t.Errorf("%s: IsValid() = %v, expected %v", c.name, c.drive.IsValid(), c.want)
		}

	}

}

func TestDriveEqual(t *testing.T) {

	base := NewDrive("root", "local")
	base.SetMountpoint("/")
	base.SetSize(1000)
	base.SetFree(400)
	base.SetUsed(600)

	if base.Equal(base) == false {
		t.Fatal("a drive should equal itself")
	}

	cases := []struct {
		name  string
		value Drive
	}{
		{"name", func() Drive { value := base; value.Name = "data"; return value }()},
		{"size", func() Drive { value := base; value.Size = 1; return value }()},
		{"free", func() Drive { value := base; value.Free = 1; return value }()},
		{"used", func() Drive { value := base; value.Used = 1; return value }()},
		{"mountpoint", func() Drive { value := base; value.Mountpoint = "/mnt"; return value }()},
		{"type", func() Drive { value := base; value.Type = "remote"; return value }()},
	}

	for _, c := range cases {

		if base.Equal(c.value) == true {
			t.Errorf("Equal() returned true for a different %s", c.name)
		}

	}

}

func TestDriveHash(t *testing.T) {

	base := NewDrive("root", "local")
	base.SetMountpoint("/")
	base.SetSize(1000)
	base.SetFree(400)
	base.SetUsed(600)

	hash := base.Hash()

	if hash == "" {
		t.Fatal("Hash() returned an empty string for a populated drive")
	}

	if base.Hash() != hash {
		t.Fatal("Hash() is not deterministic")
	}

	empty := Drive{}

	if empty.Hash() != "" {
		t.Fatal("Hash() should be empty without a name")
	}

	cases := []struct {
		name  string
		value Drive
	}{
		{"name", func() Drive { value := base; value.Name = "data"; return value }()},
		{"size", func() Drive { value := base; value.Size = 1; return value }()},
		{"free", func() Drive { value := base; value.Free = 1; return value }()},
		{"used", func() Drive { value := base; value.Used = 1; return value }()},
		{"mountpoint", func() Drive { value := base; value.Mountpoint = "/mnt"; return value }()},
		{"type", func() Drive { value := base; value.Type = "remote"; return value }()},
	}

	for _, c := range cases {

		if c.value.Hash() == hash {
			t.Errorf("Hash() did not change for a different %s", c.name)
		}

	}

}
