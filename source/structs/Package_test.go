package structs

import "testing"
import "tholian-core/matchers"
import "tholian-core/types"

func packageUnitValidPackage(name string, version string) Package {

	pkg := NewPackage("apt")

	pkg.SetName(name)
	pkg.SetVersion(version)
	pkg.SetArchitecture("x86_64")
	pkg.SetDatetime("2024-01-02T03:04:05Z")
	pkg.SetVendor("TestVendor")
	pkg.SetURL("https://example.com/" + name)

	return pkg

}

func packageUnitMatcher(name string, version string) matchers.Package {

	if version == "" {
		version = matchers.Any
	}

	return matchers.Package{
		Name:         name,
		Version:      version,
		Architecture: matchers.Any,
		Manager:      matchers.Any,
		Vendor:       matchers.Any,
	}

}

func TestPackageNewPackage(t *testing.T) {

	pkg := NewPackage("apt")

	if pkg.Manager != types.ManagerAPT {
		t.Fatalf("expected apt manager, got %q", pkg.Manager)
	}

	if pkg.Maintainers == nil ||
		pkg.Filesystem == nil ||
		pkg.Conflicts == nil ||
		pkg.Dependencies == nil ||
		pkg.Provides == nil ||
		pkg.Replaces == nil ||
		pkg.Unresolved == nil {
		t.Fatal("NewPackage did not initialize all slices")
	}

	if pkg.IsValid() == true {
		t.Fatal("a new package without a name should be invalid")
	}

	invalid := NewPackage("not-a-manager")

	if invalid.Manager != "" {
		t.Fatalf("invalid manager was accepted: %q", invalid.Manager)
	}

}

func TestPackageEqual(t *testing.T) {

	base := packageUnitValidPackage("curl", "1.2.3")

	if base.Equal(packageUnitValidPackage("curl", "1.2.3")) == false {
		t.Fatal("identical packages should be equal")
	}

	same := base
	same.Vendor = "OtherVendor"
	same.URL = "https://other.example.com/curl"
	same.Filesystem = []string{"/tmp/other.so"}
	same.Maintainers = []types.Maintainer{types.ToMaintainer("Other <other@example.com>")}

	if base.Equal(same) == false {
		t.Fatal("Equal should only compare name, version, architecture and manager")
	}

	cases := []struct {
		name   string
		mutate func(*Package)
	}{
		{"name", func(pkg *Package) { pkg.Name = "wget" }},
		{"version", func(pkg *Package) { pkg.Version = types.ToVersion("2.0.0") }},
		{"architecture", func(pkg *Package) { pkg.Architecture = types.ArchitectureX86 }},
		{"manager", func(pkg *Package) { pkg.Manager = types.ManagerDNF }},
	}

	for _, c := range cases {

		other := packageUnitValidPackage("curl", "1.2.3")
		c.mutate(&other)

		if base.Equal(other) == true {
			t.Fatalf("packages with a different %s should not be equal", c.name)
		}

	}

}

func TestPackageIsValid(t *testing.T) {

	valid := packageUnitValidPackage("curl", "1.2.3")

	if valid.IsValid() == false {
		t.Fatal("a fully populated package should be valid")
	}

	cases := []struct {
		name   string
		mutate func(*Package)
	}{
		{"no name", func(pkg *Package) { pkg.Name = "" }},
		{"invalid datetime", func(pkg *Package) { pkg.Datetime = types.Datetime{} }},
		{"invalid version", func(pkg *Package) { pkg.Version = types.Version{} }},
		{"invalid architecture", func(pkg *Package) { pkg.Architecture = "" }},
		{"invalid manager", func(pkg *Package) { pkg.Manager = "" }},
		{"invalid maintainer", func(pkg *Package) { pkg.Maintainers = []types.Maintainer{types.NewMaintainer()} }},
		{"invalid conflict", func(pkg *Package) { pkg.Conflicts = []matchers.Package{matchers.NewPackage()} }},
		{"invalid dependency", func(pkg *Package) { pkg.Dependencies = []matchers.Package{matchers.NewPackage()} }},
		{"invalid provide", func(pkg *Package) { pkg.Provides = []matchers.Package{matchers.NewPackage()} }},
		{"invalid replace", func(pkg *Package) { pkg.Replaces = []matchers.Package{matchers.NewPackage()} }},
	}

	for _, c := range cases {

		pkg := packageUnitValidPackage("curl", "1.2.3")
		c.mutate(&pkg)

		if pkg.IsValid() == true {
			t.Fatalf("package with %s should be invalid", c.name)
		}

	}

}

func TestPackageScalarSetters(t *testing.T) {

	t.Run("SetArchitecture", func(t *testing.T) {

		pkg := packageUnitValidPackage("curl", "1.2.3")
		pkg.SetArchitecture("arm64")

		if pkg.Architecture != types.ArchitectureARMv8 {
			t.Fatalf("unexpected architecture: %q", pkg.Architecture)
		}

		before := pkg.Architecture
		pkg.SetArchitecture("totally-invalid-architecture")

		if pkg.Architecture != before {
			t.Fatal("SetArchitecture should ignore invalid values")
		}

	})

	t.Run("SetDatetime", func(t *testing.T) {

		pkg := packageUnitValidPackage("curl", "1.2.3")
		pkg.SetDatetime("2025-06-07T08:09:10Z")

		if pkg.Datetime.Year != 2025 {
			t.Fatalf("unexpected datetime: %s", pkg.Datetime.String())
		}

		before := pkg.Datetime
		pkg.SetDatetime("not-a-datetime")

		if pkg.Datetime.Equal(before) == false {
			t.Fatal("SetDatetime should ignore invalid values")
		}

	})

	t.Run("SetManager", func(t *testing.T) {

		pkg := packageUnitValidPackage("curl", "1.2.3")
		pkg.SetManager("dnf")

		if pkg.Manager != types.ManagerDNF {
			t.Fatalf("unexpected manager: %q", pkg.Manager)
		}

		before := pkg.Manager
		pkg.SetManager("not-a-manager")

		if pkg.Manager != before {
			t.Fatal("SetManager should ignore invalid values")
		}

	})

	t.Run("SetName", func(t *testing.T) {

		pkg := NewPackage("apt")
		pkg.SetName("  curl  ")

		if pkg.Name != "curl" {
			t.Fatalf("SetName did not trim: %q", pkg.Name)
		}

	})

	t.Run("SetURL", func(t *testing.T) {

		pkg := NewPackage("apt")

		pkg.SetURL("https://example.com/curl")

		if pkg.URL != "https://example.com/curl" {
			t.Fatalf("unexpected https URL: %q", pkg.URL)
		}

		pkg.SetURL("http://example.com/curl")

		if pkg.URL != "http://example.com/curl" {
			t.Fatalf("unexpected http URL: %q", pkg.URL)
		}

		before := pkg.URL
		pkg.SetURL("ftp://example.com/curl")

		if pkg.URL != before {
			t.Fatal("SetURL should reject non-http schemes")
		}

		pkg.SetURL("://not-a-url")

		if pkg.URL != before {
			t.Fatal("SetURL should reject malformed URLs")
		}

	})

	t.Run("SetVendor", func(t *testing.T) {

		pkg := NewPackage("apt")
		pkg.SetVendor("  Tholian  ")

		if pkg.Vendor != "Tholian" {
			t.Fatalf("SetVendor did not trim: %q", pkg.Vendor)
		}

	})

	t.Run("SetVersion", func(t *testing.T) {

		pkg := NewPackage("apt")
		pkg.SetVersion("1.2.3")

		if pkg.Version.IsValid() == false {
			t.Fatal("SetVersion rejected a valid version")
		}

		before := pkg.Version
		pkg.SetVersion("not-a-version")

		if pkg.Version.Equal(before) == false {
			t.Fatal("SetVersion should ignore invalid values")
		}

	})

}

func TestPackageConflicts(t *testing.T) {

	pkg := NewPackage("apt")

	alpha := packageUnitMatcher("alpha", "1.0")
	beta := packageUnitMatcher("beta", "2.0")
	gamma := packageUnitMatcher("gamma", "3.0")

	pkg.AddConflict(alpha)
	pkg.AddConflict(alpha)

	if len(pkg.Conflicts) != 1 {
		t.Fatalf("AddConflict did not deduplicate: %d", len(pkg.Conflicts))
	}

	pkg.AddConflict(matchers.NewPackage())

	if len(pkg.Conflicts) != 1 {
		t.Fatal("AddConflict accepted an invalid matcher")
	}

	pkg.SetConflicts([]matchers.Package{alpha, matchers.NewPackage(), beta})

	if len(pkg.Conflicts) != 2 {
		t.Fatalf("SetConflicts did not filter invalid entries: %d", len(pkg.Conflicts))
	}

	values := []matchers.Package{alpha, beta, gamma}

	for _, position := range []int{0, 1, 2} {

		current := NewPackage("apt")
		current.SetConflicts(values)
		current.RemoveConflict(values[position])

		if len(current.Conflicts) != 2 {
			t.Fatalf("RemoveConflict at position %d left %d entries", position, len(current.Conflicts))
		}

		for _, conflict := range current.Conflicts {
			if conflict.Equal(values[position]) {
				t.Fatalf("RemoveConflict at position %d left the removed matcher", position)
			}
		}

	}

	noop := NewPackage("apt")
	noop.SetConflicts(values)
	noop.RemoveConflict(packageUnitMatcher("absent", "1.0"))

	if len(noop.Conflicts) != 3 {
		t.Fatal("RemoveConflict removed an absent matcher")
	}

}

func TestPackageDependencies(t *testing.T) {

	pkg := NewPackage("apt")

	alpha := packageUnitMatcher("alpha", "1.0")
	beta := packageUnitMatcher("beta", "2.0")
	gamma := packageUnitMatcher("gamma", "3.0")

	pkg.AddDependency(alpha)
	pkg.AddDependency(alpha)

	if len(pkg.Dependencies) != 1 {
		t.Fatalf("AddDependency did not deduplicate: %d", len(pkg.Dependencies))
	}

	pkg.AddDependency(matchers.NewPackage())

	if len(pkg.Dependencies) != 1 {
		t.Fatal("AddDependency accepted an invalid matcher")
	}

	if pkg.HasDependency(alpha) == false {
		t.Fatal("HasDependency did not find an existing dependency")
	}

	if pkg.HasDependency(packageUnitMatcher("absent", "1.0")) == true {
		t.Fatal("HasDependency found an absent dependency")
	}

	pkg.SetDependencies([]matchers.Package{alpha, matchers.NewPackage(), beta})

	if len(pkg.Dependencies) != 2 {
		t.Fatalf("SetDependencies did not filter invalid entries: %d", len(pkg.Dependencies))
	}

	values := []matchers.Package{alpha, beta, gamma}

	for _, position := range []int{0, 1, 2} {

		current := NewPackage("apt")
		current.SetDependencies(values)
		current.RemoveDependency(values[position])

		if len(current.Dependencies) != 2 {
			t.Fatalf("RemoveDependency at position %d left %d entries", position, len(current.Dependencies))
		}

	}

	noop := NewPackage("apt")
	noop.SetDependencies(values)
	noop.RemoveDependency(packageUnitMatcher("absent", "1.0"))

	if len(noop.Dependencies) != 3 {
		t.Fatal("RemoveDependency removed an absent matcher")
	}

}

func TestPackageResolveDependencies(t *testing.T) {

	t.Run("direct match", func(t *testing.T) {

		pkg := packageUnitValidPackage("application", "1.0.0")
		pkg.Unresolved = []matchers.Unresolved{matchers.ToUnresolved("curl")}

		pkg.ResolveDependencies([]Package{packageUnitValidPackage("curl", "1.2.3")})

		if len(pkg.Unresolved) != 0 {
			t.Fatalf("expected all dependencies to resolve, got %d unresolved", len(pkg.Unresolved))
		}

		if len(pkg.Dependencies) != 1 {
			t.Fatalf("expected 1 dependency, got %d", len(pkg.Dependencies))
		}

		dependency := pkg.Dependencies[0]

		if dependency.Name != "curl" || dependency.Version != "0:1.2.3" {
			t.Fatalf("unexpected dependency: %+v", dependency)
		}

		if dependency.Architecture != "x86_64" ||
			dependency.Manager != "apt" ||
			dependency.Vendor != "TestVendor" {
			t.Fatalf("dependency metadata was not resolved: %+v", dependency)
		}

	})

	t.Run("provided match", func(t *testing.T) {

		provider := packageUnitValidPackage("openssl", "3.0.0")
		provider.Provides = []matchers.Package{packageUnitMatcher("libssl", "")}

		pkg := packageUnitValidPackage("application", "1.0.0")
		pkg.Unresolved = []matchers.Unresolved{matchers.ToUnresolved("libssl")}
		pkg.ResolveDependencies([]Package{provider})

		if len(pkg.Unresolved) != 0 || len(pkg.Dependencies) != 1 {
			t.Fatalf("provided dependency was not resolved: %+v", pkg)
		}

		dependency := pkg.Dependencies[0]

		if dependency.Name != "openssl" || dependency.Version != "0:3.0.0" {
			t.Fatalf("unexpected provided dependency: %+v", dependency)
		}

	})

	t.Run("unresolved remains", func(t *testing.T) {

		pkg := packageUnitValidPackage("application", "1.0.0")
		pkg.Unresolved = []matchers.Unresolved{matchers.ToUnresolved("absent")}
		pkg.ResolveDependencies([]Package{packageUnitValidPackage("curl", "1.2.3")})

		if len(pkg.Unresolved) != 1 {
			t.Fatal("unmatched dependency should remain unresolved")
		}

		if len(pkg.Dependencies) != 0 {
			t.Fatal("unmatched dependency was added")
		}

	})

	t.Run("no packages", func(t *testing.T) {

		pkg := packageUnitValidPackage("application", "1.0.0")
		pkg.Unresolved = []matchers.Unresolved{matchers.ToUnresolved("curl")}
		pkg.ResolveDependencies(nil)

		if len(pkg.Unresolved) != 1 {
			t.Fatal("resolving without packages changed the unresolved list")
		}

	})

	t.Run("existing dependency is not duplicated", func(t *testing.T) {

		pkg := packageUnitValidPackage("application", "1.0.0")
		pkg.Unresolved = []matchers.Unresolved{matchers.ToUnresolved("curl")}
		pkg.Dependencies = []matchers.Package{packageUnitMatcher("curl", "0:1.2.3")}
		pkg.ResolveDependencies([]Package{packageUnitValidPackage("curl", "1.2.3")})

		if len(pkg.Dependencies) != 1 {
			t.Fatalf("existing dependency was duplicated: %d", len(pkg.Dependencies))
		}

	})

}

func TestPackageFilesystem(t *testing.T) {

	pkg := NewPackage("apt")

	pkg.AddFilesystem("/usr/bin/curl")

	if len(pkg.Filesystem) != 1 {
		t.Fatalf("unexpected filesystem entries: %v", pkg.Filesystem)
	}

	pkg.AddFilesystem("/tmp/evil.so")
	pkg.AddFilesystem("/tmp/evil.so")
	pkg.AddFilesystem("/usr/lib/sub/libssl.so")

	if len(pkg.Filesystem) != 3 {
		t.Fatalf("unexpected filesystem entries: %v", pkg.Filesystem)
	}

	if pkg.HasFilesystem("/tmp/evil.so") == false {
		t.Fatal("HasFilesystem did not find a watched file")
	}

	if pkg.HasFilesystem("/usr/bin/curl") == false {
		t.Fatal("HasFilesystem did not find a watched file")
	}

	values := []string{"/tmp/a.so", "/tmp/b.so", "/tmp/c.so"}

	for _, position := range []int{0, 1, 2} {

		current := NewPackage("apt")
		current.SetFilesystem(values)
		current.RemoveFilesystem(values[position])

		if len(current.Filesystem) != 2 {
			t.Fatalf("RemoveFilesystem at position %d left %d entries", position, len(current.Filesystem))
		}

	}

	noop := NewPackage("apt")
	noop.SetFilesystem(values)
	noop.RemoveFilesystem("/tmp/absent.so")

	if len(noop.Filesystem) != 3 {
		t.Fatal("RemoveFilesystem removed an absent path")
	}

	pkg.SetFilesystem([]string{
		"/tmp/a.so",
		"/tmp/a.so",
		"/usr/bin/curl",
		"/home/user/a.txt",
		"/tmp/b.so",
	})

	if len(pkg.Filesystem) != 3 {
		t.Fatalf("SetFilesystem did not filter and deduplicate: %v", pkg.Filesystem)
	}

	pkg.SetFilesystem(nil)

	if len(pkg.Filesystem) != 0 {
		t.Fatal("SetFilesystem did not replace the filesystem")
	}

}

func TestPackageMaintainers(t *testing.T) {

	pkg := NewPackage("apt")

	alice := types.ToMaintainer("Alice <alice@example.com>")
	bob := types.ToMaintainer("Bob <bob@example.com>")
	carol := types.ToMaintainer("Carol <carol@example.com>")

	pkg.AddMaintainer(alice)
	pkg.AddMaintainer(alice)

	if len(pkg.Maintainers) != 1 {
		t.Fatalf("AddMaintainer did not deduplicate: %d", len(pkg.Maintainers))
	}

	pkg.AddMaintainer(types.NewMaintainer())

	if len(pkg.Maintainers) != 1 {
		t.Fatal("AddMaintainer accepted an invalid maintainer")
	}

	pkg.SetMaintainers([]types.Maintainer{alice, types.NewMaintainer(), bob})

	if len(pkg.Maintainers) != 2 {
		t.Fatalf("SetMaintainers did not filter invalid entries: %d", len(pkg.Maintainers))
	}

	values := []types.Maintainer{alice, bob, carol}

	for _, position := range []int{0, 1, 2} {

		current := NewPackage("apt")
		current.SetMaintainers(values)
		current.RemoveMaintainer(values[position])

		if len(current.Maintainers) != 2 {
			t.Fatalf("RemoveMaintainer at position %d left %d entries", position, len(current.Maintainers))
		}

	}

	noop := NewPackage("apt")
	noop.SetMaintainers(values)
	noop.RemoveMaintainer(types.ToMaintainer("Absent <absent@example.com>"))

	if len(noop.Maintainers) != 3 {
		t.Fatal("RemoveMaintainer removed an absent maintainer")
	}

}

func TestPackageProvides(t *testing.T) {

	pkg := NewPackage("apt")

	alpha := packageUnitMatcher("alpha", "1.0")
	beta := packageUnitMatcher("beta", "2.0")
	gamma := packageUnitMatcher("gamma", "3.0")

	pkg.AddProvide(alpha)
	pkg.AddProvide(alpha)

	if len(pkg.Provides) != 1 {
		t.Fatalf("AddProvide did not deduplicate: %d", len(pkg.Provides))
	}

	pkg.AddProvide(matchers.NewPackage())

	if len(pkg.Provides) != 1 {
		t.Fatal("AddProvide accepted an invalid matcher")
	}

	pkg.SetProvides([]matchers.Package{alpha, matchers.NewPackage(), beta})

	if len(pkg.Provides) != 2 {
		t.Fatalf("SetProvides did not filter invalid entries: %d", len(pkg.Provides))
	}

	values := []matchers.Package{alpha, beta, gamma}

	for _, position := range []int{0, 1, 2} {

		current := NewPackage("apt")
		current.SetProvides(values)
		current.RemoveProvide(values[position])

		if len(current.Provides) != 2 {
			t.Fatalf("RemoveProvide at position %d left %d entries", position, len(current.Provides))
		}

	}

	noop := NewPackage("apt")
	noop.SetProvides(values)
	noop.RemoveProvide(packageUnitMatcher("absent", "1.0"))

	if len(noop.Provides) != 3 {
		t.Fatal("RemoveProvide removed an absent matcher")
	}

}

func TestPackageReplaces(t *testing.T) {

	pkg := NewPackage("apt")

	alpha := packageUnitMatcher("alpha", "1.0")
	beta := packageUnitMatcher("beta", "2.0")
	gamma := packageUnitMatcher("gamma", "3.0")

	pkg.AddReplace(alpha)
	pkg.AddReplace(alpha)

	if len(pkg.Replaces) != 1 {
		t.Fatalf("AddReplace did not deduplicate: %d", len(pkg.Replaces))
	}

	pkg.AddReplace(matchers.NewPackage())

	if len(pkg.Replaces) != 1 {
		t.Fatal("AddReplace accepted an invalid matcher")
	}

	pkg.SetReplaces([]matchers.Package{alpha, matchers.NewPackage(), beta})

	if len(pkg.Replaces) != 2 {
		t.Fatalf("SetReplaces did not filter invalid entries: %d", len(pkg.Replaces))
	}

	values := []matchers.Package{alpha, beta, gamma}

	for _, position := range []int{0, 1, 2} {

		current := NewPackage("apt")
		current.SetReplaces(values)
		current.RemoveReplace(values[position])

		if len(current.Replaces) != 2 {
			t.Fatalf("RemoveReplace at position %d left %d entries", position, len(current.Replaces))
		}

	}

	noop := NewPackage("apt")
	noop.SetReplaces(values)
	noop.RemoveReplace(packageUnitMatcher("absent", "1.0"))

	if len(noop.Replaces) != 3 {
		t.Fatal("RemoveReplace removed an absent matcher")
	}

}

func TestPackageHasMatchers(t *testing.T) {

	pkg := NewPackage("apt")

	alpha := packageUnitMatcher("alpha", "1.0")
	beta := packageUnitMatcher("beta", "2.0")
	alice := types.ToMaintainer("Alice <alice@example.com>")
	bob := types.ToMaintainer("Bob <bob@example.com>")

	if pkg.HasConflict(alpha) == true || pkg.HasProvide(alpha) == true || pkg.HasReplace(alpha) == true || pkg.HasMaintainer(alice) == true {
		t.Fatal("empty package reported an existing entry")
	}

	pkg.AddConflict(alpha)
	pkg.AddProvide(alpha)
	pkg.AddReplace(alpha)
	pkg.AddMaintainer(alice)

	if pkg.HasConflict(alpha) == false {
		t.Fatal("HasConflict did not find an existing conflict")
	}

	if pkg.HasProvide(alpha) == false {
		t.Fatal("HasProvide did not find an existing provide")
	}

	if pkg.HasReplace(alpha) == false {
		t.Fatal("HasReplace did not find an existing replace")
	}

	if pkg.HasMaintainer(alice) == false {
		t.Fatal("HasMaintainer did not find an existing maintainer")
	}

	if pkg.HasConflict(beta) == true || pkg.HasProvide(beta) == true || pkg.HasReplace(beta) == true || pkg.HasMaintainer(bob) == true {
		t.Fatal("package reported a missing entry")
	}

}

func TestPackageUnresolved(t *testing.T) {

	pkg := NewPackage("apt")

	alpha := matchers.ToUnresolved("curl")
	beta := matchers.ToUnresolved("wget")
	gamma := matchers.ToUnresolved("git")

	pkg.AddUnresolved(alpha)
	pkg.AddUnresolved(matchers.ToUnresolved("curl"))

	if len(pkg.Unresolved) != 1 {
		t.Fatalf("AddUnresolved did not deduplicate: %d", len(pkg.Unresolved))
	}

	pkg.AddUnresolved(matchers.NewUnresolved())

	if len(pkg.Unresolved) != 1 {
		t.Fatal("AddUnresolved accepted an empty unresolved value")
	}

	if pkg.HasUnresolved(alpha) == false {
		t.Fatal("HasUnresolved did not find an existing value")
	}

	if pkg.HasUnresolved(beta) == true {
		t.Fatal("HasUnresolved found a missing value")
	}

	pkg.SetUnresolved([]matchers.Unresolved{alpha, matchers.NewUnresolved(), beta})

	if len(pkg.Unresolved) != 2 {
		t.Fatalf("SetUnresolved did not filter invalid entries: %d", len(pkg.Unresolved))
	}

	pkg.RemoveUnresolved(alpha)

	if len(pkg.Unresolved) != 1 || pkg.HasUnresolved(alpha) == true || pkg.HasUnresolved(beta) == false {
		t.Fatalf("RemoveUnresolved left unexpected entries: %+v", pkg.Unresolved)
	}

	noop := NewPackage("apt")
	noop.SetUnresolved([]matchers.Unresolved{alpha, beta, gamma})
	noop.RemoveUnresolved(matchers.ToUnresolved("absent"))

	if len(noop.Unresolved) != 3 {
		t.Fatal("RemoveUnresolved removed an absent value")
	}

}

func TestPackageHash(t *testing.T) {

	empty := NewPackage("apt")

	if empty.Hash() != "" {
		t.Fatal("a package without a name should have an empty hash")
	}

	first := packageUnitValidPackage("curl", "1.2.3")

	if first.Hash() == "" {
		t.Fatal("a valid package should have a non-empty hash")
	}

	second := packageUnitValidPackage("curl", "1.2.3")

	if first.Hash() != second.Hash() {
		t.Fatal("identical packages should have identical hashes")
	}

	differentVendor := packageUnitValidPackage("curl", "1.2.3")
	differentVendor.Vendor = "OtherVendor"

	if first.Hash() != differentVendor.Hash() {
		t.Fatal("Hash should not depend on the vendor")
	}

	differentVersion := packageUnitValidPackage("curl", "2.0.0")

	if first.Hash() == differentVersion.Hash() {
		t.Fatal("changing the version should change the hash")
	}

}
