package structs

import "testing"
import "tholian-core/matchers"

func programUnitClientConnection() matchers.Connection {
	return matchers.ToConnection("1.2.3.4:443TC")
}

func programUnitServerConnection() matchers.Connection {
	return matchers.ToConnection("1.2.3.4:443TS")
}

func programUnitValidProgram() Program {

	program := NewProgram(1337, "stealth.mjs")

	program.SetCommand("/usr/bin/node")
	program.SetFolder("/home/user/stealth/")
	program.SetUser(matchers.ToUser("alice"))
	program.SetManager(matchers.ToManager("npm"))
	program.SetArguments([]string{"serve", "--port", "8080"})
	program.AddEnvironment("LD_PRELOAD", "/tmp/evil.so")
	program.AddFilesystem("/tmp/evil.so")

	return program

}

func TestProgramNewProgram(t *testing.T) {

	program := NewProgram(1337, "  stealth.mjs  ")

	if program.PID != 1337 {
		t.Fatalf("unexpected PID: %d", program.PID)
	}

	if program.Name != "stealth.mjs" {
		t.Fatalf("NewProgram did not trim the name: %q", program.Name)
	}

	if program.Arguments == nil ||
		program.Environment == nil ||
		program.Filesystem == nil ||
		program.Connections == nil ||
		program.Dependencies == nil ||
		program.Packages == nil {
		t.Fatal("NewProgram did not initialize all fields")
	}

	if program.IsValid() == false {
		t.Fatal("a program with a PID and name should be valid")
	}

	empty := NewProgram(0, "")

	if empty.IsValid() == true {
		t.Fatal("a program without a PID or name should be invalid")
	}

}

func TestProgramEqual(t *testing.T) {

	base := programUnitValidProgram()

	if base.Equal(programUnitValidProgram()) == false {
		t.Fatal("identical programs should be equal")
	}

	cases := []struct {
		name   string
		mutate func(*Program)
	}{
		{"PID", func(program *Program) { program.PID = 1338 }},
		{"name", func(program *Program) { program.Name = "other.mjs" }},
		{"command", func(program *Program) { program.Command = "/usr/bin/other" }},
		{"folder", func(program *Program) { program.Folder = "/home/user/other" }},
		{"arguments", func(program *Program) { program.Arguments = []string{"serve"} }},
		{"environment", func(program *Program) { program.Environment["OTHER"] = "value" }},
		{"user", func(program *Program) { program.User = matchers.ToUser("bob") }},
		{"manager", func(program *Program) { program.Manager = matchers.ToManager("apt") }},
		{"filesystem", func(program *Program) { program.Filesystem = append(program.Filesystem, "/tmp/other.so") }},
		{"connections", func(program *Program) { program.AddConnection(programUnitClientConnection()) }},
		{"dependencies", func(program *Program) { program.AddDependency(packageUnitMatcher("curl", "1.0")) }},
		{"packages", func(program *Program) { program.AddPackage(packageUnitValidPackage("curl", "1.0.0")) }},
	}

	for _, c := range cases {

		other := programUnitValidProgram()
		c.mutate(&other)

		if base.Equal(other) == true {
			t.Fatalf("programs with a different %s should not be equal", c.name)
		}

	}

}

func TestProgramIsProgramIsService(t *testing.T) {

	program := programUnitValidProgram()

	if program.IsProgram() == false {
		t.Fatal("a program without server connections should be a client program")
	}

	if program.IsService() == true {
		t.Fatal("a program without server connections should not be a service")
	}

	program.AddConnection(programUnitClientConnection())

	if program.IsProgram() == false || program.IsService() == true {
		t.Fatal("a client connection should not turn a program into a service")
	}

	program.AddConnection(programUnitServerConnection())

	if program.IsProgram() == true {
		t.Fatal("a server connection should not be a client program")
	}

	if program.IsService() == false {
		t.Fatal("a server connection should make a service")
	}

	incomplete := NewProgram(1, "incomplete")

	if incomplete.IsProgram() == true || incomplete.IsService() == true {
		t.Fatal("an incomplete program should be neither client nor service")
	}

	if (&Program{}).IsProgram() == true || (&Program{}).IsService() == true {
		t.Fatal("an empty program should be neither client nor service")
	}

}

func TestProgramIsValid(t *testing.T) {

	cases := []struct {
		name    string
		program Program
		want    bool
	}{
		{"valid", NewProgram(1, "program"), true},
		{"no pid", NewProgram(0, "program"), false},
		{"no name", NewProgram(1, ""), false},
		{"empty", Program{}, false},
	}

	for _, c := range cases {

		if c.program.IsValid() != c.want {
			t.Fatalf("%s: IsValid() = %v, expected %v", c.name, c.program.IsValid(), c.want)
		}

	}

}

func TestProgramScalarSetters(t *testing.T) {

	t.Run("SetName", func(t *testing.T) {

		program := NewProgram(1, "program")
		program.SetName("  renamed  ")

		if program.Name != "renamed" {
			t.Fatalf("SetName did not trim: %q", program.Name)
		}

	})

	t.Run("SetArguments", func(t *testing.T) {

		program := NewProgram(1, "program")
		arguments := []string{"one", "two"}
		program.SetArguments(arguments)

		if len(program.Arguments) != 2 ||
			program.Arguments[0] != "one" ||
			program.Arguments[1] != "two" {
			t.Fatalf("unexpected arguments: %v", program.Arguments)
		}

	})

	t.Run("SetCommand", func(t *testing.T) {

		program := NewProgram(1, "program")
		program.SetCommand("  /usr/bin/node  ")

		if program.Command != "/usr/bin/node" {
			t.Fatalf("SetCommand did not trim: %q", program.Command)
		}

	})

	t.Run("SetFolder", func(t *testing.T) {

		program := NewProgram(1, "program")
		program.SetFolder("/home/user/project/")

		if program.Folder != "/home/user/project" {
			t.Fatalf("SetFolder did not trim the trailing slash: %q", program.Folder)
		}

		program.SetFolder("/")

		if program.Folder != "/" {
			t.Fatalf("SetFolder rejected the root folder: %q", program.Folder)
		}

		before := program.Folder
		program.SetFolder("relative/path")

		if program.Folder != before {
			t.Fatal("SetFolder accepted a relative path")
		}

	})

	t.Run("SetManager", func(t *testing.T) {

		program := NewProgram(1, "program")
		program.SetManager(matchers.ToManager("npm"))

		if program.Manager.Name != "npm" {
			t.Fatalf("unexpected manager: %+v", program.Manager)
		}

		before := program.Manager
		program.SetManager(matchers.NewManager())

		if program.Manager.Equal(before) == false {
			t.Fatal("SetManager accepted an invalid manager")
		}

	})

	t.Run("SetUser", func(t *testing.T) {

		program := NewProgram(1, "program")
		program.SetUser(matchers.ToUser("alice"))

		if program.User.Name != "alice" {
			t.Fatalf("unexpected user: %+v", program.User)
		}

		before := program.User
		program.SetUser(matchers.NewUser())

		if program.User.Equal(before) == false {
			t.Fatal("SetUser accepted an invalid user")
		}

	})

}

func TestProgramConnections(t *testing.T) {

	program := NewProgram(1, "program")

	client := programUnitClientConnection()
	server := programUnitServerConnection()
	peer := matchers.ToConnection("1.2.3.4:8080TP")

	program.AddConnection(client)
	program.AddConnection(client)

	if len(program.Connections) != 1 {
		t.Fatalf("AddConnection did not deduplicate: %d", len(program.Connections))
	}

	program.AddConnection(matchers.NewConnection())

	if len(program.Connections) != 1 {
		t.Fatal("AddConnection accepted an invalid connection")
	}

	program.SetConnections([]matchers.Connection{client, matchers.NewConnection(), server})

	if len(program.Connections) != 2 {
		t.Fatalf("SetConnections did not filter invalid entries: %d", len(program.Connections))
	}

	values := []matchers.Connection{client, server, peer}

	for _, position := range []int{0, 1, 2} {

		current := NewProgram(1, "program")
		current.SetConnections(values)
		current.RemoveConnection(values[position])

		if len(current.Connections) != 2 {
			t.Fatalf("RemoveConnection at position %d left %d entries", position, len(current.Connections))
		}

	}

	noop := NewProgram(1, "program")
	noop.SetConnections(values)
	noop.RemoveConnection(matchers.ToConnection("5.6.7.8:443TC"))

	if len(noop.Connections) != 3 {
		t.Fatal("RemoveConnection removed an absent connection")
	}

}

func TestProgramDependencies(t *testing.T) {

	program := NewProgram(1, "program")

	alpha := packageUnitMatcher("alpha", "1.0")
	beta := packageUnitMatcher("beta", "2.0")
	gamma := packageUnitMatcher("gamma", "3.0")

	program.AddDependency(alpha)
	program.AddDependency(alpha)

	if len(program.Dependencies) != 1 {
		t.Fatalf("AddDependency did not deduplicate: %d", len(program.Dependencies))
	}

	program.AddDependency(matchers.NewPackage())

	if len(program.Dependencies) != 1 {
		t.Fatal("AddDependency accepted an invalid matcher")
	}

	if program.HasDependency(alpha) == false {
		t.Fatal("HasDependency did not find an existing dependency")
	}

	if program.HasDependency(packageUnitMatcher("absent", "1.0")) == true {
		t.Fatal("HasDependency found an absent dependency")
	}

	program.SetDependencies([]matchers.Package{alpha, matchers.NewPackage(), beta})

	if len(program.Dependencies) != 2 {
		t.Fatalf("SetDependencies did not filter invalid entries: %d", len(program.Dependencies))
	}

	values := []matchers.Package{alpha, beta, gamma}

	for _, position := range []int{0, 1, 2} {

		current := NewProgram(1, "program")
		current.SetDependencies(values)
		current.RemoveDependency(values[position])

		if len(current.Dependencies) != 2 {
			t.Fatalf("RemoveDependency at position %d left %d entries", position, len(current.Dependencies))
		}

	}

	noop := NewProgram(1, "program")
	noop.SetDependencies(values)
	noop.RemoveDependency(packageUnitMatcher("absent", "1.0"))

	if len(noop.Dependencies) != 3 {
		t.Fatal("RemoveDependency removed an absent matcher")
	}

}

func TestProgramResolveDependencies(t *testing.T) {

	t.Run("matching filesystem", func(t *testing.T) {

		pkg := packageUnitValidPackage("curl", "1.2.3")
		pkg.Filesystem = []string{"/tmp/evil.so"}

		program := programUnitValidProgram()
		program.ResolveDependencies([]Package{pkg})

		if len(program.Dependencies) != 1 {
			t.Fatalf("expected 1 dependency, got %d", len(program.Dependencies))
		}

		dependency := program.Dependencies[0]

		if dependency.Name != "curl" || dependency.Version != "0:1.2.3" {
			t.Fatalf("unexpected dependency: %+v", dependency)
		}

		if dependency.Manager != "apt" || dependency.Vendor != "TestVendor" {
			t.Fatalf("dependency metadata was not resolved: %+v", dependency)
		}

	})

	t.Run("no matching filesystem", func(t *testing.T) {

		pkg := packageUnitValidPackage("curl", "1.2.3")
		pkg.Filesystem = []string{"/tmp/other.so"}

		program := programUnitValidProgram()
		program.ResolveDependencies([]Package{pkg})

		if len(program.Dependencies) != 0 {
			t.Fatal("unmatched filesystem added a dependency")
		}

	})

	t.Run("no packages", func(t *testing.T) {

		program := programUnitValidProgram()
		program.ResolveDependencies(nil)

		if len(program.Dependencies) != 0 {
			t.Fatal("resolving without packages added a dependency")
		}

	})

	t.Run("existing dependency is not duplicated", func(t *testing.T) {

		pkg := packageUnitValidPackage("curl", "1.2.3")
		pkg.Filesystem = []string{"/tmp/evil.so"}

		program := programUnitValidProgram()
		program.Dependencies = []matchers.Package{packageUnitMatcher("curl", "0:1.2.3")}
		program.ResolveDependencies([]Package{pkg})

		if len(program.Dependencies) != 1 {
			t.Fatalf("existing dependency was duplicated: %d", len(program.Dependencies))
		}

	})

}

func TestProgramEnvironment(t *testing.T) {

	program := NewProgram(1, "program")

	program.AddEnvironment("KEY", "value")
	program.AddEnvironment("KEY", "value")

	if len(program.Environment) != 1 || program.Environment["KEY"] != "value" {
		t.Fatalf("AddEnvironment did not deduplicate: %v", program.Environment)
	}

	program.AddEnvironment("KEY", "other")

	if program.Environment["KEY"] != "other" {
		t.Fatal("AddEnvironment did not replace a changed value")
	}

	program.SetEnvironment("SECOND", "value")

	if program.Environment["SECOND"] != "value" {
		t.Fatal("SetEnvironment did not add a value")
	}

	program.SetEnvironment("", "value")
	program.SetEnvironment("THIRD", "")

	if len(program.Environment) != 2 {
		t.Fatalf("SetEnvironment accepted an empty key or value: %v", program.Environment)
	}

	program.RemoveEnvironment("KEY")

	if _, ok := program.Environment["KEY"]; ok {
		t.Fatal("RemoveEnvironment did not remove the key")
	}

	program.RemoveEnvironment("ABSENT")

	if len(program.Environment) != 1 {
		t.Fatal("RemoveEnvironment changed an absent key")
	}

}

func TestProgramFilesystem(t *testing.T) {

	program := NewProgram(1, "program")

	program.AddFilesystem("/tmp/a.so")
	program.AddFilesystem("/tmp/a.so")

	if len(program.Filesystem) != 1 {
		t.Fatalf("AddFilesystem did not deduplicate: %d", len(program.Filesystem))
	}

	program.AddFilesystem("/usr/bin/curl")
	program.AddFilesystem("/home/user/a.txt")

	if len(program.Filesystem) != 2 {
		t.Fatal("AddFilesystem accepted a path that is not watched")
	}

	values := []string{"/tmp/a.so", "/tmp/b.so", "/tmp/c.so"}

	for _, position := range []int{0, 1, 2} {

		current := NewProgram(1, "program")
		current.SetFilesystem(values)
		current.RemoveFilesystem(values[position])

		if len(current.Filesystem) != 2 {
			t.Fatalf("RemoveFilesystem at position %d left %d entries", position, len(current.Filesystem))
		}

	}

	noop := NewProgram(1, "program")
	noop.SetFilesystem(values)
	noop.RemoveFilesystem("/tmp/absent.so")

	if len(noop.Filesystem) != 3 {
		t.Fatal("RemoveFilesystem removed an absent path")
	}

	program.SetFilesystem([]string{
		"/tmp/a.so",
		"/tmp/a.so",
		"/usr/bin/curl",
		"/home/user/a.txt",
		"/tmp/b.so",
	})

	if len(program.Filesystem) != 3 {
		t.Fatalf("SetFilesystem did not filter and deduplicate: %v", program.Filesystem)
	}

}

func TestProgramPackages(t *testing.T) {

	alpha := packageUnitValidPackage("alpha", "1.0.0")
	beta := packageUnitValidPackage("beta", "2.0.0")
	gamma := packageUnitValidPackage("gamma", "3.0.0")
	invalid := Package{}

	t.Run("AddPackage", func(t *testing.T) {

		program := NewProgram(1, "program")

		program.AddPackage(alpha)
		program.AddPackage(alpha)

		if len(program.Packages) != 1 {
			t.Fatalf("AddPackage did not deduplicate: %d", len(program.Packages))
		}

		program.AddPackage(invalid)

		if len(program.Packages) != 1 {
			t.Fatal("AddPackage accepted an invalid package")
		}

	})

	t.Run("RemovePackage", func(t *testing.T) {

		values := []Package{alpha, beta, gamma}

		for _, position := range []int{0, 1, 2} {

			program := NewProgram(1, "program")
			program.SetPackages(values)
			program.RemovePackage(values[position])

			if len(program.Packages) != 2 {
				t.Fatalf("RemovePackage at position %d left %d entries", position, len(program.Packages))
			}

			for _, pkg := range program.Packages {
				if pkg.Equal(values[position]) {
					t.Fatalf("RemovePackage at position %d left the removed package", position)
				}
			}

		}

		program := NewProgram(1, "program")
		program.SetPackages(values)
		program.RemovePackage(packageUnitValidPackage("absent", "1.0.0"))

		if len(program.Packages) != 3 {
			t.Fatal("RemovePackage removed an absent package")
		}

	})

	t.Run("AddPackages deduplicates", func(t *testing.T) {

		program := NewProgram(1, "program")
		program.AddPackages([]Package{alpha, alpha, beta, invalid, gamma, beta})

		if len(program.Packages) != 3 {
			t.Fatalf("AddPackages did not deduplicate: %d", len(program.Packages))
		}

		program.AddPackages([]Package{alpha, beta, gamma})

		if len(program.Packages) != 3 {
			t.Fatalf("AddPackages duplicated existing packages: %d", len(program.Packages))
		}

		program.AddPackages([]Package{alpha, packageUnitValidPackage("delta", "4.0.0")})

		if len(program.Packages) != 4 {
			t.Fatalf("AddPackages did not add a new package: %d", len(program.Packages))
		}

	})

	t.Run("SetPackages filters invalid", func(t *testing.T) {

		program := NewProgram(1, "program")
		program.SetPackages([]Package{alpha, invalid, beta})

		if len(program.Packages) != 2 {
			t.Fatalf("SetPackages did not filter invalid entries: %d", len(program.Packages))
		}

		program.SetPackages(nil)

		if len(program.Packages) != 0 {
			t.Fatal("SetPackages did not replace the packages")
		}

	})

}

func TestProgramCollectionHelpers(t *testing.T) {

	program := NewProgram(1, "program")

	program.SetPID(1337)

	if program.PID != 1337 {
		t.Fatalf("SetPID did not update the PID: %d", program.PID)
	}

	client := programUnitClientConnection()
	other := matchers.ToConnection("5.6.7.8:443TC")

	program.AddConnection(client)

	if program.HasConnection(client) == false {
		t.Fatal("HasConnection did not find an existing connection")
	}

	if program.HasConnection(other) == true {
		t.Fatal("HasConnection found a missing connection")
	}

	program.AddFilesystem("/tmp/a.so")

	if program.HasFilesystem("/tmp/a.so") == false {
		t.Fatal("HasFilesystem did not find an existing file")
	}

	if program.HasFilesystem("/tmp/b.so") == true {
		t.Fatal("HasFilesystem found a missing file")
	}

	alpha := packageUnitValidPackage("alpha", "1.0.0")
	beta := packageUnitValidPackage("beta", "2.0.0")
	gamma := packageUnitValidPackage("gamma", "3.0.0")

	program.AddPackages([]Package{alpha, beta, gamma})

	if program.HasPackage(alpha) == false {
		t.Fatal("HasPackage did not find an existing package")
	}

	if program.HasPackage(packageUnitValidPackage("delta", "4.0.0")) == true {
		t.Fatal("HasPackage found a missing package")
	}

	program.RemovePackages([]Package{alpha, gamma})

	if len(program.Packages) != 1 || program.HasPackage(beta) == false {
		t.Fatalf("RemovePackages left unexpected packages: %+v", program.Packages)
	}

	program.RemovePackages([]Package{alpha})

	if len(program.Packages) != 1 {
		t.Fatal("RemovePackages removed an absent package")
	}

}

func TestProgramHash(t *testing.T) {

	empty := NewProgram(0, "")

	if empty.Hash() != "" {
		t.Fatal("a program without a name should have an empty hash")
	}

	first := programUnitValidProgram()

	if first.Hash() == "" {
		t.Fatal("a program with a name should have a non-empty hash")
	}

	if first.Hash() != first.Hash() {
		t.Fatal("Hash should be deterministic")
	}

	second := programUnitValidProgram()

	if first.Hash() != second.Hash() {
		t.Fatal("identical programs should have identical hashes")
	}

	ordered := programUnitValidProgram()
	ordered.Environment = map[string]string{"ALPHA": "1", "BETA": "2"}

	reordered := programUnitValidProgram()
	reordered.Environment = map[string]string{"BETA": "2", "ALPHA": "1"}

	if ordered.Hash() != reordered.Hash() {
		t.Fatal("Hash should not depend on environment insertion order")
	}

	changed := programUnitValidProgram()
	changed.SetCommand("/usr/bin/other")

	if changed.Hash() == first.Hash() {
		t.Fatal("changing the command should change the hash")
	}

}
