package structs

import "testing"

func TestNewDistribution(t *testing.T) {

	distribution := NewDistribution()

	if distribution.Kernel != "any" {
		t.Errorf("NewDistribution() Kernel = %q, expected %q", distribution.Kernel, "any")
	}

	if distribution.KernelArchitecture != "any" {
		t.Errorf("NewDistribution() KernelArchitecture = %q, expected %q", distribution.KernelArchitecture, "any")
	}

	if distribution.KernelVersion != "any" {
		t.Errorf("NewDistribution() KernelVersion = %q, expected %q", distribution.KernelVersion, "any")
	}

	if distribution.Manager != "any" {
		t.Errorf("NewDistribution() Manager = %q, expected %q", distribution.Manager, "any")
	}

	if distribution.KernelModules == nil || len(distribution.KernelModules) != 0 {
		t.Errorf("NewDistribution() KernelModules = %#v, expected an empty slice", distribution.KernelModules)
	}

	if distribution.Name != "" || distribution.Version != "" || distribution.Vendor != "" {
		t.Errorf("NewDistribution() = %+v, expected empty name, version and vendor", distribution)
	}

}

func TestDistributionSetKernel(t *testing.T) {

	cases := []struct {
		value string
		want  string
	}{
		{"Android", "android"},
		{"android", "android"},
		{"Darwin", "darwin"},
		{"darwin", "darwin"},
		{"DragonFly", "dragonfly"},
		{"dragonfly", "dragonfly"},
		{"FreeBSD", "freebsd"},
		{"freebsd", "freebsd"},
		{"GNU", "hurd"},
		{"hurd", "hurd"},
		{"illumos", "illumos"},
		{"Linux", "linux"},
		{"linux", "linux"},
		{"NetBSD", "netbsd"},
		{"netbsd", "netbsd"},
		{"OpenBSD", "openbsd"},
		{"openbsd", "openbsd"},
		{"Solaris", "solaris"},
		{"solaris", "solaris"},
		{"SunOS", "solaris"},
		{"Windows", "windows"},
		{"MINGW32_NT", "windows"},
	}

	for _, c := range cases {

		distribution := NewDistribution()
		distribution.SetKernel(c.value)

		if distribution.Kernel != c.want {
			t.Errorf("SetKernel(%q) = %q, expected %q", c.value, distribution.Kernel, c.want)
		}

	}

	distribution := NewDistribution()
	distribution.SetKernel("windows")
	distribution.SetKernel("invalid")

	if distribution.Kernel != "any" {
		t.Errorf("SetKernel() accepted an invalid value: %q", distribution.Kernel)
	}

}

func TestDistributionSetters(t *testing.T) {

	distribution := NewDistribution()

	distribution.SetName("  Debian  ")
	distribution.SetVersion("  12  ")
	distribution.SetVendor("  Debian Project  ")
	distribution.SetKernelArchitecture("amd64")
	distribution.SetKernelVersion("  6.1.0  ")
	distribution.SetKernelModules([]string{"ext4", "nvme"})
	distribution.SetManager("apt")
	distribution.SetKeywords(map[string]string{"id": "debian"})

	if distribution.Name != "Debian" {
		t.Errorf("SetName() = %q, expected %q", distribution.Name, "Debian")
	}

	if distribution.Version != "12" {
		t.Errorf("SetVersion() = %q, expected %q", distribution.Version, "12")
	}

	if distribution.Vendor != "Debian Project" {
		t.Errorf("SetVendor() = %q, expected %q", distribution.Vendor, "Debian Project")
	}

	if distribution.KernelArchitecture != "x86_64" {
		t.Errorf("SetKernelArchitecture() = %q, expected %q", distribution.KernelArchitecture, "x86_64")
	}

	if distribution.KernelVersion != "6.1.0" {
		t.Errorf("SetKernelVersion() = %q, expected %q", distribution.KernelVersion, "6.1.0")
	}

	if len(distribution.KernelModules) != 2 || distribution.KernelModules[0] != "ext4" {
		t.Errorf("SetKernelModules() = %#v", distribution.KernelModules)
	}

	if distribution.Manager != "apt" {
		t.Errorf("SetManager() = %q, expected %q", distribution.Manager, "apt")
	}

	if distribution.Keywords == nil || (*distribution.Keywords)["id"] != "debian" {
		t.Errorf("SetKeywords() = %v", distribution.Keywords)
	}

	distribution.SetKernelArchitecture("invalid")

	if distribution.KernelArchitecture != "x86_64" {
		t.Errorf("SetKernelArchitecture() accepted an invalid value: %q", distribution.KernelArchitecture)
	}

	distribution.SetManager("invalid")

	if distribution.Manager != "apt" {
		t.Errorf("SetManager() accepted an invalid value: %q", distribution.Manager)
	}

}

func TestDistributionSetKeywordsCopies(t *testing.T) {

	source := map[string]string{"id": "debian", "version": "12"}

	distribution := NewDistribution()
	distribution.SetKeywords(source)

	source["id"] = "changed"

	if distribution.Keywords == nil || (*distribution.Keywords)["id"] != "debian" {
		t.Fatalf("SetKeywords() did not copy the map: %v", distribution.Keywords)
	}

}

func TestDistributionIsValid(t *testing.T) {

	distribution := NewDistribution()

	if distribution.IsValid() == true {
		t.Fatal("an empty distribution should be invalid")
	}

	distribution.SetName("  Debian  ")

	if distribution.IsValid() == false {
		t.Fatal("a named distribution should be valid")
	}

}

func TestDistributionEqual(t *testing.T) {

	base := NewDistribution()
	base.SetName("Debian")
	base.SetVersion("12")
	base.SetVendor("Debian Project")
	base.SetKernel("linux")

	if base.Equal(base) == false {
		t.Fatal("a distribution should equal itself")
	}

	cases := []struct {
		name  string
		value Distribution
	}{
		{"name", func() Distribution { value := base; value.Name = "Ubuntu"; return value }()},
		{"version", func() Distribution { value := base; value.Version = "13"; return value }()},
		{"vendor", func() Distribution { value := base; value.Vendor = "Canonical"; return value }()},
		{"kernel", func() Distribution { value := base; value.Kernel = "freebsd"; return value }()},
	}

	for _, c := range cases {

		if base.Equal(c.value) == true {
			t.Errorf("Equal() returned true for a different %s", c.name)
		}

	}

	ignored := base
	ignored.SetKernelModules([]string{"ext4"})
	ignored.SetKernelVersion("6.1.0")
	ignored.SetManager("apt")
	ignored.SetKeywords(map[string]string{"id": "debian"})

	if base.Equal(ignored) == false {
		t.Fatal("Equal() should ignore kernel modules, kernel version, manager and keywords")
	}

}

func TestDistributionHash(t *testing.T) {

	base := NewDistribution()
	base.SetName("Debian")
	base.SetVersion("12")
	base.SetVendor("Debian Project")
	base.SetKernel("linux")

	hash := base.Hash()

	if hash == "" {
		t.Fatal("Hash() returned an empty string for a populated distribution")
	}

	if base.Hash() != hash {
		t.Fatal("Hash() is not deterministic")
	}

	empty := Distribution{}

	if empty.Hash() != "" {
		t.Fatal("Hash() should be empty without a name")
	}

	cases := []struct {
		name  string
		value Distribution
	}{
		{"name", func() Distribution { value := base; value.Name = "Ubuntu"; return value }()},
		{"version", func() Distribution { value := base; value.Version = "13"; return value }()},
		{"vendor", func() Distribution { value := base; value.Vendor = "Canonical"; return value }()},
		{"kernel", func() Distribution { value := base; value.Kernel = "freebsd"; return value }()},
	}

	for _, c := range cases {

		if c.value.Hash() == hash {
			t.Errorf("Hash() did not change for a different %s", c.name)
		}

	}

	ignored := base
	ignored.SetKernelModules([]string{"ext4"})
	ignored.SetKernelVersion("6.1.0")
	ignored.SetManager("apt")
	ignored.SetKeywords(map[string]string{"id": "debian"})

	if ignored.Hash() != hash {
		t.Fatal("Hash() should ignore kernel modules, kernel version, manager and keywords")
	}

}
