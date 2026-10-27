package structs

import "testing"

func TestNewAntique(t *testing.T) {

	antique := NewAntique("apt", " openssl ")

	if antique.Name != "" {
		t.Errorf("NewAntique() Name = %q, expected empty", antique.Name)
	}

	if antique.Service != "openssl" {
		t.Errorf("NewAntique() Service = %q, expected %q", antique.Service, "openssl")
	}

	if antique.Architecture.String() != "any" {
		t.Errorf("NewAntique() Architecture = %q, expected %q", antique.Architecture.String(), "any")
	}

	if antique.Manager.String() != "apt" {
		t.Errorf("NewAntique() Manager = %q, expected %q", antique.Manager.String(), "apt")
	}

	invalid := NewAntique("invalid-manager", "service")

	if invalid.Manager.String() != "" {
		t.Errorf("NewAntique() accepted an invalid manager: %q", invalid.Manager.String())
	}

}

func TestAntiqueSetters(t *testing.T) {

	antique := NewAntique("apt", "openssl")

	antique.SetName("  OpenSSL  ")
	antique.SetService("  libssl  ")
	antique.SetVendor("  OpenSSL Foundation  ")
	antique.SetVersion("1.0.0")
	antique.SetArchitecture("x86_64")
	antique.SetManager("apt")
	antique.SetURL("https://example.com/security")

	if antique.Name != "OpenSSL" {
		t.Errorf("SetName() = %q, expected %q", antique.Name, "OpenSSL")
	}

	if antique.Service != "libssl" {
		t.Errorf("SetService() = %q, expected %q", antique.Service, "libssl")
	}

	if antique.Vendor != "OpenSSL Foundation" {
		t.Errorf("SetVendor() = %q, expected %q", antique.Vendor, "OpenSSL Foundation")
	}

	if antique.Version.IsValid() == false {
		t.Fatal("SetVersion() did not accept a valid version")
	}

	if antique.Architecture.String() != "x86_64" {
		t.Errorf("SetArchitecture() = %q, expected %q", antique.Architecture.String(), "x86_64")
	}

	if antique.Manager.String() != "apt" {
		t.Errorf("SetManager() = %q, expected %q", antique.Manager.String(), "apt")
	}

	if antique.URL != "https://example.com/security" {
		t.Errorf("SetURL() = %q, expected %q", antique.URL, "https://example.com/security")
	}

	before := antique

	antique.SetURL("ftp://example.com/file")
	antique.SetURL("relative/path")
	antique.SetVersion("not a version")
	antique.SetArchitecture("not-an-architecture")
	antique.SetManager("not-a-manager")

	if antique.URL != before.URL {
		t.Errorf("SetURL() accepted an invalid URL: %q", antique.URL)
	}

	if antique.Version.Equal(before.Version) == false {
		t.Error("SetVersion() accepted an invalid version")
	}

	if antique.Architecture.Equal(before.Architecture) == false {
		t.Error("SetArchitecture() accepted an invalid architecture")
	}

	if antique.Manager.Equal(before.Manager) == false {
		t.Error("SetManager() accepted an invalid manager")
	}

}

func TestAntiqueIsValid(t *testing.T) {

	valid := func() Antique {

		antique := NewAntique("apt", "openssl")
		antique.SetName("OpenSSL")
		antique.SetVersion("1.0.0")

		return antique

	}

	cases := []struct {
		name    string
		antique Antique
		want    bool
	}{
		{"default", Antique{}, false},
		{"valid", valid(), true},
		{"missing name", func() Antique { value := valid(); value.Name = ""; return value }(), false},
		{"missing service", func() Antique { value := valid(); value.Service = ""; return value }(), false},
		{"missing version", func() Antique {
			value := Antique{}
			value.SetName("OpenSSL")
			value.SetService("openssl")
			value.SetArchitecture("any")
			value.SetManager("apt")
			return value
		}(), false},
		{"missing architecture", func() Antique { value := valid(); value.Architecture = ""; return value }(), false},
		{"missing manager", func() Antique { value := valid(); value.Manager = ""; return value }(), false},
	}

	for _, c := range cases {

		if c.antique.IsValid() != c.want {
			t.Errorf("%s: IsValid() = %v, expected %v", c.name, c.antique.IsValid(), c.want)
		}

	}

}

func TestAntiqueEqual(t *testing.T) {

	base := NewAntique("apt", "openssl")
	base.SetName("OpenSSL")
	base.SetVersion("1.0.0")
	base.SetArchitecture("x86_64")
	base.SetManager("apt")
	base.SetVendor("OpenSSL Foundation")
	base.SetURL("https://example.com")

	if base.Equal(base) == false {
		t.Fatal("an antique should equal itself")
	}

	cases := []struct {
		name  string
		value Antique
	}{
		{"name", func() Antique { value := base; value.Name = "Other"; return value }()},
		{"version", func() Antique { value := base; value.SetVersion("2.0.0"); return value }()},
		{"architecture", func() Antique { value := base; value.SetArchitecture("arm64"); return value }()},
		{"manager", func() Antique { value := base; value.SetManager("dnf"); return value }()},
		{"vendor", func() Antique { value := base; value.Vendor = "Other"; return value }()},
		{"service", func() Antique { value := base; value.Service = "other"; return value }()},
		{"url", func() Antique { value := base; value.URL = "https://other.example.com"; return value }()},
	}

	for _, c := range cases {

		if base.Equal(c.value) == true {
			t.Errorf("Equal() returned true for a different %s", c.name)
		}

	}

}

func TestAntiqueHash(t *testing.T) {

	base := NewAntique("apt", "openssl")
	base.SetName("OpenSSL")
	base.SetVersion("1.0.0")
	base.SetArchitecture("x86_64")
	base.SetManager("apt")
	base.SetVendor("OpenSSL Foundation")
	base.SetURL("https://example.com")

	hash := base.Hash()

	if hash == "" {
		t.Fatal("Hash() returned an empty string for a populated antique")
	}

	if base.Hash() != hash {
		t.Fatal("Hash() is not deterministic")
	}

	empty := Antique{}

	if empty.Hash() != "" {
		t.Fatal("Hash() should be empty without a name")
	}

	cases := []struct {
		name  string
		value Antique
	}{
		{"name", func() Antique { value := base; value.Name = "Other"; return value }()},
		{"version", func() Antique { value := base; value.SetVersion("2.0.0"); return value }()},
		{"architecture", func() Antique { value := base; value.SetArchitecture("arm64"); return value }()},
		{"manager", func() Antique { value := base; value.SetManager("dnf"); return value }()},
		{"vendor", func() Antique { value := base; value.Vendor = "Other"; return value }()},
		{"service", func() Antique { value := base; value.Service = "other"; return value }()},
		{"url", func() Antique { value := base; value.URL = "https://other.example.com"; return value }()},
	}

	for _, c := range cases {

		if c.value.Hash() == hash {
			t.Errorf("Hash() did not change for a different %s", c.name)
		}

	}

}
