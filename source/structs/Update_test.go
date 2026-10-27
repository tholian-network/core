package structs

import "testing"

func TestNewUpdate(t *testing.T) {

	update := NewUpdate("apt")

	if update.Name != "" {
		t.Errorf("NewUpdate() Name = %q, expected empty", update.Name)
	}

	if update.Architecture.String() != "any" {
		t.Errorf("NewUpdate() Architecture = %q, expected %q", update.Architecture.String(), "any")
	}

	if update.Manager.String() != "apt" {
		t.Errorf("NewUpdate() Manager = %q, expected %q", update.Manager.String(), "apt")
	}

	invalid := NewUpdate("invalid-manager")

	if invalid.Manager.String() != "" {
		t.Errorf("NewUpdate() accepted an invalid manager: %q", invalid.Manager.String())
	}

}

func TestUpdateSetters(t *testing.T) {

	update := NewUpdate("apt")

	update.SetName("  openssl  ")
	update.SetVendor("  OpenSSL Foundation  ")
	update.SetVersion("1.0.0")
	update.SetArchitecture("x86_64")
	update.SetManager("apt")
	update.SetURL("https://example.com/security")

	if update.Name != "openssl" {
		t.Errorf("SetName() = %q, expected %q", update.Name, "openssl")
	}

	if update.Vendor != "OpenSSL Foundation" {
		t.Errorf("SetVendor() = %q, expected %q", update.Vendor, "OpenSSL Foundation")
	}

	if update.Version.IsValid() == false {
		t.Fatal("SetVersion() did not accept a valid version")
	}

	if update.Architecture.String() != "x86_64" {
		t.Errorf("SetArchitecture() = %q, expected %q", update.Architecture.String(), "x86_64")
	}

	if update.Manager.String() != "apt" {
		t.Errorf("SetManager() = %q, expected %q", update.Manager.String(), "apt")
	}

	if update.URL != "https://example.com/security" {
		t.Errorf("SetURL() = %q, expected %q", update.URL, "https://example.com/security")
	}

	before := update

	update.SetURL("ftp://example.com/file")
	update.SetURL("relative/path")
	update.SetVersion("not a version")
	update.SetArchitecture("not-an-architecture")
	update.SetManager("not-a-manager")

	if update.URL != before.URL {
		t.Errorf("SetURL() accepted an invalid URL: %q", update.URL)
	}

	if update.Version.Equal(before.Version) == false {
		t.Error("SetVersion() accepted an invalid version")
	}

	if update.Architecture.Equal(before.Architecture) == false {
		t.Error("SetArchitecture() accepted an invalid architecture")
	}

	if update.Manager.Equal(before.Manager) == false {
		t.Error("SetManager() accepted an invalid manager")
	}

}

func TestUpdateIsValid(t *testing.T) {

	valid := func() Update {

		update := NewUpdate("apt")
		update.SetName("openssl")
		update.SetVersion("1.0.0")

		return update

	}

	cases := []struct {
		name   string
		update Update
		want   bool
	}{
		{"default", Update{}, false},
		{"valid", valid(), true},
		{"missing name", func() Update { value := valid(); value.Name = ""; return value }(), false},
		{"missing version", func() Update {
			value := Update{}
			value.SetName("openssl")
			value.SetArchitecture("any")
			value.SetManager("apt")
			return value
		}(), false},
		{"missing architecture", func() Update { value := valid(); value.Architecture = ""; return value }(), false},
		{"missing manager", func() Update { value := valid(); value.Manager = ""; return value }(), false},
	}

	for _, c := range cases {

		if c.update.IsValid() != c.want {
			t.Errorf("%s: IsValid() = %v, expected %v", c.name, c.update.IsValid(), c.want)
		}

	}

}

func TestUpdateEqual(t *testing.T) {

	base := NewUpdate("apt")
	base.SetName("openssl")
	base.SetVersion("1.0.0")
	base.SetArchitecture("x86_64")
	base.SetManager("apt")
	base.SetVendor("OpenSSL Foundation")
	base.SetURL("https://example.com")

	if base.Equal(base) == false {
		t.Fatal("an update should equal itself")
	}

	cases := []struct {
		name  string
		value Update
	}{
		{"name", func() Update { value := base; value.Name = "other"; return value }()},
		{"version", func() Update { value := base; value.SetVersion("2.0.0"); return value }()},
		{"architecture", func() Update { value := base; value.SetArchitecture("arm64"); return value }()},
		{"manager", func() Update { value := base; value.SetManager("dnf"); return value }()},
		{"vendor", func() Update { value := base; value.Vendor = "Other"; return value }()},
		{"url", func() Update { value := base; value.URL = "https://other.example.com"; return value }()},
	}

	for _, c := range cases {

		if base.Equal(c.value) == true {
			t.Errorf("Equal() returned true for a different %s", c.name)
		}

	}

}

func TestUpdateHash(t *testing.T) {

	base := NewUpdate("apt")
	base.SetName("openssl")
	base.SetVersion("1.0.0")
	base.SetArchitecture("x86_64")
	base.SetManager("apt")
	base.SetVendor("OpenSSL Foundation")
	base.SetURL("https://example.com")

	hash := base.Hash()

	if hash == "" {
		t.Fatal("Hash() returned an empty string for a populated update")
	}

	if base.Hash() != hash {
		t.Fatal("Hash() is not deterministic")
	}

	empty := Update{}

	if empty.Hash() != "" {
		t.Fatal("Hash() should be empty without a name")
	}

	cases := []struct {
		name  string
		value Update
	}{
		{"name", func() Update { value := base; value.Name = "other"; return value }()},
		{"version", func() Update { value := base; value.SetVersion("2.0.0"); return value }()},
		{"architecture", func() Update { value := base; value.SetArchitecture("arm64"); return value }()},
		{"manager", func() Update { value := base; value.SetManager("dnf"); return value }()},
		{"vendor", func() Update { value := base; value.Vendor = "Other"; return value }()},
		{"url", func() Update { value := base; value.URL = "https://other.example.com"; return value }()},
	}

	for _, c := range cases {

		if c.value.Hash() == hash {
			t.Errorf("Hash() did not change for a different %s", c.name)
		}

	}

}
