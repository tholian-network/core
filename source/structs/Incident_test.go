package structs

import "testing"
import "tholian-core/matchers"
import "tholian-core/types"

func IncidentTestConnection() matchers.Connection {

	connection := matchers.NewConnection()
	connection.Socket.Host = "1.2.3.4"
	connection.Socket.Port = 1234
	connection.Protocol = "tcp"
	connection.Type = "client"

	return connection

}

func IncidentTestDistribution() matchers.Distribution {

	distribution := matchers.NewDistribution()
	distribution.SetName("debian")
	distribution.SetVersion("12")
	distribution.SetVendor("debian")

	return distribution

}

func IncidentTestNetwork() matchers.Network {

	network := matchers.NewNetwork()
	network.SetSubnet("1.1.1.0/24")

	return network

}

func IncidentTestPackage() matchers.Package {

	pkg := matchers.NewPackage()
	pkg.SetName("openssl")

	return pkg

}

func IncidentTestProduct() matchers.Product {

	product := matchers.NewProduct()
	product.SetVendor("acme")
	product.SetProduct("widget")
	product.SetVersion("1-2-3")
	product.SetType("software")

	return product

}

func IncidentTestProgram() matchers.Program {

	program := matchers.NewProgram()
	program.SetName("stealth.mjs")

	return program

}

func IncidentTestUser() matchers.User {

	user := matchers.NewUser()
	user.SetName("alice")

	return user

}

func IncidentTestVulnerability() matchers.Vulnerability {

	vulnerability := matchers.NewVulnerability()
	vulnerability.SetName("CVE-2021-1234")

	return vulnerability

}

func IncidentTestWeakness() matchers.Weakness {

	weakness := matchers.NewWeakness()
	weakness.SetName("CWE-79")

	return weakness

}

func IncidentTestBase(typ string) Incident {

	incident := NewIncident("machine", typ)
	incident.SetDatetime("2024-01-02T03:04:05Z")

	return incident

}

func IncidentTestRich() Incident {

	incident := IncidentTestBase("Recon")
	incident.AddConnection(IncidentTestConnection())
	incident.AddDistribution(IncidentTestDistribution())
	incident.AddNetwork(IncidentTestNetwork())
	incident.AddPackage(IncidentTestPackage())
	incident.AddProduct(IncidentTestProduct())
	incident.AddProgram(IncidentTestProgram())
	incident.AddUser(IncidentTestUser())
	incident.AddVulnerability(IncidentTestVulnerability())
	incident.AddWeakness(IncidentTestWeakness())

	return incident

}

func TestIncidentNew(t *testing.T) {

	incident := NewIncident(" machine ", "Recon")

	if incident.System != "machine" {
		t.Fatalf("unexpected system: %q", incident.System)
	}

	if incident.Type != "Recon" {
		t.Fatalf("unexpected type: %q", incident.Type)
	}

	if incident.Datetime.IsValid() == false {
		t.Fatal("new incident has an invalid datetime")
	}

	if len(incident.Connections) != 0 ||
		len(incident.Distributions) != 0 ||
		len(incident.Networks) != 0 ||
		len(incident.Packages) != 0 ||
		len(incident.Products) != 0 ||
		len(incident.Programs) != 0 ||
		len(incident.Users) != 0 ||
		len(incident.Vulnerabilities) != 0 ||
		len(incident.Weaknesses) != 0 {
		t.Fatal("new incident has non-empty collections")
	}

	invalid := NewIncident("machine", "Bogus")

	if invalid.Type != "" {
		t.Fatalf("invalid type was accepted: %q", invalid.Type)
	}

}

func TestIncidentSetSystemAndType(t *testing.T) {

	incident := NewIncident("", "")

	incident.SetSystem("  machine  ")

	if incident.System != "machine" {
		t.Fatalf("unexpected system: %q", incident.System)
	}

	types := []string{"Recon", "Intel", "Conquer", "Persist", "Exfil", "Destroy"}

	for _, typ := range types {

		incident.SetType(typ)

		if incident.Type != typ {
			t.Fatalf("SetType(%q) = %q", typ, incident.Type)
		}

	}

	incident.SetType("Bogus")

	if incident.Type != "Destroy" {
		t.Fatalf("invalid type overwrote value: %q", incident.Type)
	}

}

func TestIncidentSetDatetime(t *testing.T) {

	incident := NewIncident("machine", "Recon")
	original := incident.Datetime

	incident.SetDatetime("invalid")

	if incident.Datetime.Equal(original) == false {
		t.Fatal("invalid datetime overwrote value")
	}

	expected := types.ToDatetime("2020-02-03T04:05:06Z")
	incident.SetDatetime("2020-02-03T04:05:06Z")

	if incident.Datetime.Equal(expected) == false {
		t.Fatalf("unexpected datetime: %s", incident.Datetime.String())
	}

}

func TestIncidentIsAfterIsBefore(t *testing.T) {

	types := []string{"Recon", "Intel", "Conquer", "Persist", "Exfil", "Destroy"}

	for i := 0; i < len(types); i++ {

		for j := i + 1; j < len(types); j++ {

			earlier := IncidentTestBase(types[i])
			later := IncidentTestBase(types[j])

			if earlier.IsBefore(later) == false {
				t.Errorf("%s should be before %s at the same datetime", types[i], types[j])
			}

			if later.IsAfter(earlier) == false {
				t.Errorf("%s should be after %s at the same datetime", types[j], types[i])
			}

			if earlier.IsAfter(later) == true {
				t.Errorf("%s should not be after %s", types[i], types[j])
			}

			if later.IsBefore(earlier) == true {
				t.Errorf("%s should not be before %s", types[j], types[i])
			}

		}

	}

	earlier := IncidentTestBase("Recon")
	earlier.SetDatetime("2024-01-01T00:00:00Z")
	later := IncidentTestBase("Destroy")
	later.SetDatetime("2024-01-03T00:00:00Z")

	if earlier.IsBefore(later) == false || later.IsAfter(earlier) == false {
		t.Fatal("datetime ordering failed")
	}

	if earlier.IsAfter(later) == true || later.IsBefore(earlier) == true {
		t.Fatal("reverse datetime ordering failed")
	}

	if earlier.IsBefore(earlier) == true || earlier.IsAfter(earlier) == true {
		t.Fatal("an incident should not be before or after itself")
	}

}

func TestIncidentEqual(t *testing.T) {

	left := IncidentTestRich()
	right := IncidentTestRich()

	if left.Equal(right) == false {
		t.Fatal("identical incidents should be equal")
	}

	if left.Equal(left) == false {
		t.Fatal("an incident should equal itself")
	}

	cases := []struct {
		name   string
		mutate func(*Incident)
	}{
		{"system", func(value *Incident) { value.System = "other" }},
		{"type", func(value *Incident) { value.Type = "Destroy" }},
		{"datetime", func(value *Incident) { value.SetDatetime("2020-01-01T00:00:00Z") }},
		{"connections-length", func(value *Incident) { value.Connections = append(value.Connections, value.Connections[0]) }},
		{"connections-value", func(value *Incident) { value.Connections[0].Socket.Host = "9.9.9.9" }},
		{"distributions-length", func(value *Incident) { value.Distributions = append(value.Distributions, value.Distributions[0]) }},
		{"distributions-value", func(value *Incident) { value.Distributions[0].Name = "ubuntu" }},
		{"networks-length", func(value *Incident) { value.Networks = append(value.Networks, value.Networks[0]) }},
		{"networks-value", func(value *Incident) { value.Networks[0].Name = "other" }},
		{"packages-length", func(value *Incident) { value.Packages = append(value.Packages, value.Packages[0]) }},
		{"packages-value", func(value *Incident) { value.Packages[0].Name = "nginx" }},
		{"products-length", func(value *Incident) { value.Products = append(value.Products, value.Products[0]) }},
		{"products-value", func(value *Incident) { value.Products[0].Vendor = "other" }},
		{"programs-length", func(value *Incident) { value.Programs = append(value.Programs, value.Programs[0]) }},
		{"programs-value", func(value *Incident) { value.Programs[0].Name = "other" }},
		{"users-length", func(value *Incident) { value.Users = append(value.Users, value.Users[0]) }},
		{"users-value", func(value *Incident) { value.Users[0].Name = "bob" }},
		{"vulnerabilities-length", func(value *Incident) { value.Vulnerabilities = append(value.Vulnerabilities, value.Vulnerabilities[0]) }},
		{"vulnerabilities-value", func(value *Incident) { value.Vulnerabilities[0].Name = "CVE-2020-0000" }},
		{"weaknesses-length", func(value *Incident) { value.Weaknesses = append(value.Weaknesses, value.Weaknesses[0]) }},
		{"weaknesses-value", func(value *Incident) { value.Weaknesses[0].Name = "CWE-999" }},
	}

	for _, c := range cases {

		t.Run(c.name, func(t *testing.T) {

			other := IncidentTestRich()
			c.mutate(&other)

			if left.Equal(other) == true {
				t.Fatalf("incidents should differ for %s", c.name)
			}

		})

	}

}

func TestIncidentIsValidMatrix(t *testing.T) {

	cases := []struct {
		typ      string
		populate func(*Incident)
	}{
		{"Recon", func(value *Incident) {
			value.AddConnection(IncidentTestConnection())
			value.AddProgram(IncidentTestProgram())
		}},
		{"Intel", func(value *Incident) {
			value.AddVulnerability(IncidentTestVulnerability())
			value.AddPackage(IncidentTestPackage())
		}},
		{"Conquer", func(value *Incident) {
			value.AddDistribution(IncidentTestDistribution())
			value.AddVulnerability(IncidentTestVulnerability())
			value.AddPackage(IncidentTestPackage())
		}},
		{"Persist", func(value *Incident) {
			value.AddDistribution(IncidentTestDistribution())
			value.AddVulnerability(IncidentTestVulnerability())
			value.AddUser(IncidentTestUser())
		}},
		{"Exfil", func(value *Incident) {
			value.AddConnection(IncidentTestConnection())
			value.AddUser(IncidentTestUser())
		}},
		{"Destroy", func(value *Incident) {
			value.AddPackage(IncidentTestPackage())
		}},
	}

	for _, c := range cases {

		t.Run(c.typ, func(t *testing.T) {

			valid := IncidentTestBase(c.typ)
			c.populate(&valid)

			if valid.IsValid() == false {
				t.Fatalf("%s incident should be valid", c.typ)
			}

			empty := IncidentTestBase(c.typ)

			if empty.IsValid() == true {
				t.Fatalf("%s incident without evidence should be invalid", c.typ)
			}

		})

	}

	noSystem := IncidentTestBase("Destroy")
	noSystem.SetSystem("")
	noSystem.AddPackage(IncidentTestPackage())

	if noSystem.IsValid() == true {
		t.Fatal("incident without system should be invalid")
	}

	noDatetime := IncidentTestBase("Destroy")
	noDatetime.Datetime = types.Datetime{}
	noDatetime.AddPackage(IncidentTestPackage())

	if noDatetime.IsValid() == true {
		t.Fatal("incident without datetime should be invalid")
	}

}

func TestIncidentMatches(t *testing.T) {

	incident := IncidentTestBase("Recon")
	incident.AddConnection(IncidentTestConnection())
	incident.AddDistribution(IncidentTestDistribution())
	incident.AddNetwork(IncidentTestNetwork())
	incident.AddPackage(IncidentTestPackage())
	incident.AddProduct(IncidentTestProduct())
	incident.AddProgram(IncidentTestProgram())
	incident.AddUser(IncidentTestUser())
	incident.AddVulnerability(IncidentTestVulnerability())
	incident.AddWeakness(IncidentTestWeakness())

	antique := NewAntique("apt", "openssl")
	antique.SetName("openssl")
	antique.SetVersion("1.0.0")
	antique.SetArchitecture("amd64")
	antique.SetVendor("debian")

	if incident.MatchesAntique(antique) == false {
		t.Fatal("MatchesAntique should match")
	}

	otherAntique := antique
	otherAntique.SetName("nginx")

	if incident.MatchesAntique(otherAntique) == true {
		t.Fatal("MatchesAntique should not match")
	}

	connection := types.ToConnection("[tcp] 1.2.3.4:1234 -> 5.6.7.8:443")

	if incident.MatchesConnection(connection) == false {
		t.Fatal("MatchesConnection should match")
	}

	otherConnection := types.ToConnection("[udp] 9.9.9.9:53 -> 8.8.8.8:53")

	if incident.MatchesConnection(otherConnection) == true {
		t.Fatal("MatchesConnection should not match")
	}

	targetConnection := types.ToConnection("[tcp] 9.9.9.9:9999 -> 1.2.3.4:1234")

	if incident.MatchesConnection(targetConnection) == false {
		t.Fatal("MatchesConnection should match the target endpoint")
	}

	distribution := NewDistribution()
	distribution.SetName("debian")
	distribution.SetVersion("12")
	distribution.SetManager("apt")
	distribution.SetVendor("debian")

	if incident.MatchesDistribution(distribution) == false {
		t.Fatal("MatchesDistribution should match")
	}

	otherDistribution := distribution
	otherDistribution.SetName("ubuntu")

	if incident.MatchesDistribution(otherDistribution) == true {
		t.Fatal("MatchesDistribution should not match")
	}

	network := NewNetwork("trusted")
	network.AddSubnet(ToSubnet("1.1.1.0/24"))

	if incident.MatchesNetwork(network) == false {
		t.Fatal("MatchesNetwork should match")
	}

	otherNetwork := NewNetwork("other")
	otherNetwork.AddSubnet(ToSubnet("2.2.2.0/24"))

	if incident.MatchesNetwork(otherNetwork) == true {
		t.Fatal("MatchesNetwork should not match")
	}

	pkg := NewPackage("apt")
	pkg.SetName("openssl")
	pkg.SetVersion("1.2.3")
	pkg.SetArchitecture("amd64")
	pkg.SetVendor("debian")
	pkg.SetDatetime("2024-01-02T03:04:05Z")

	if incident.MatchesPackage(pkg) == false {
		t.Fatal("MatchesPackage should match")
	}

	otherPackage := pkg
	otherPackage.SetName("nginx")

	if incident.MatchesPackage(otherPackage) == true {
		t.Fatal("MatchesPackage should not match")
	}

	product := NewProduct("software")
	product.SetVendor("acme")
	product.SetProduct("widget")
	product.SetVersion("1-2-3")
	product.SetState("published")

	if incident.MatchesProduct(product) == false {
		t.Fatal("MatchesProduct should match")
	}

	otherProduct := product
	otherProduct.SetVendor("other")

	if incident.MatchesProduct(otherProduct) == true {
		t.Fatal("MatchesProduct should not match")
	}

	program := NewProgram(1234, "stealth.mjs")
	program.SetCommand("/usr/bin/node")
	program.SetArguments([]string{"serve"})

	if incident.MatchesProgram(program) == false {
		t.Fatal("MatchesProgram should match")
	}

	otherProgram := program
	otherProgram.SetName("other.mjs")

	if incident.MatchesProgram(otherProgram) == true {
		t.Fatal("MatchesProgram should not match")
	}

	update := NewUpdate("apt")
	update.SetName("openssl")
	update.SetVersion("1.2.3")
	update.SetArchitecture("amd64")

	if incident.MatchesUpdate(update) == false {
		t.Fatal("MatchesUpdate should match")
	}

	otherUpdate := update
	otherUpdate.SetName("nginx")

	if incident.MatchesUpdate(otherUpdate) == true {
		t.Fatal("MatchesUpdate should not match")
	}

	user := types.ToUser("alice", 1000)

	if incident.MatchesUser(user) == false {
		t.Fatal("MatchesUser should match")
	}

	otherUser := types.ToUser("bob", 1001)

	if incident.MatchesUser(otherUser) == true {
		t.Fatal("MatchesUser should not match")
	}

	vulnerability := NewVulnerability("CVE-2021-1234")
	vulnerability.SetState("published")

	if incident.MatchesVulnerability(vulnerability) == false {
		t.Fatal("MatchesVulnerability should match")
	}

	otherVulnerability := NewVulnerability("CVE-2020-0000")
	otherVulnerability.SetState("published")

	if incident.MatchesVulnerability(otherVulnerability) == true {
		t.Fatal("MatchesVulnerability should not match")
	}

	weakness := NewWeakness()
	weakness.SetName("CWE-79")
	weakness.AddImpact("Read Application Data")
	weakness.AddScope("Confidentiality")

	if incident.MatchesWeakness(weakness) == false {
		t.Fatal("MatchesWeakness should match")
	}

	otherWeakness := weakness
	otherWeakness.SetName("CWE-89")

	if incident.MatchesWeakness(otherWeakness) == true {
		t.Fatal("MatchesWeakness should not match")
	}

}

func TestIncidentCollectionHelpers(t *testing.T) {

	t.Run("connections", func(t *testing.T) {

		incident := IncidentTestBase("Recon")
		valid := IncidentTestConnection()
		invalid := matchers.NewConnection()

		incident.AddConnection(valid)
		incident.AddConnection(valid)
		incident.AddConnection(invalid)

		if len(incident.Connections) != 1 {
			t.Fatalf("AddConnection did not filter or deduplicate: %d", len(incident.Connections))
		}

		incident.SetConnections([]matchers.Connection{valid, invalid, valid})

		if len(incident.Connections) != 2 {
			t.Fatalf("SetConnections did not filter: %d", len(incident.Connections))
		}

		incident.RemoveConnection(valid)

		if len(incident.Connections) != 1 {
			t.Fatalf("RemoveConnection failed: %d", len(incident.Connections))
		}

		incident.RemoveConnection(invalid)

		if len(incident.Connections) != 1 {
			t.Fatalf("RemoveConnection removed an unknown value: %d", len(incident.Connections))
		}

		incident.RemoveConnection(valid)

		if len(incident.Connections) != 0 {
			t.Fatalf("RemoveConnection did not remove the last value: %d", len(incident.Connections))
		}

	})

	t.Run("distributions", func(t *testing.T) {

		incident := IncidentTestBase("Recon")
		valid := IncidentTestDistribution()
		invalid := matchers.NewDistribution()

		incident.AddDistribution(valid)
		incident.AddDistribution(valid)
		incident.AddDistribution(invalid)

		if len(incident.Distributions) != 1 {
			t.Fatalf("AddDistribution did not filter or deduplicate: %d", len(incident.Distributions))
		}

		incident.SetDistributions([]matchers.Distribution{valid, invalid, valid})

		if len(incident.Distributions) != 2 {
			t.Fatalf("SetDistributions did not filter: %d", len(incident.Distributions))
		}

		incident.RemoveDistribution(valid)

		if len(incident.Distributions) != 1 {
			t.Fatalf("RemoveDistribution failed: %d", len(incident.Distributions))
		}

		incident.RemoveDistribution(invalid)

		if len(incident.Distributions) != 1 {
			t.Fatalf("RemoveDistribution removed an unknown value: %d", len(incident.Distributions))
		}

		incident.RemoveDistribution(valid)

		if len(incident.Distributions) != 0 {
			t.Fatalf("RemoveDistribution did not remove the last value: %d", len(incident.Distributions))
		}

	})

	t.Run("networks", func(t *testing.T) {

		incident := IncidentTestBase("Recon")
		valid := IncidentTestNetwork()
		invalid := matchers.NewNetwork()

		incident.AddNetwork(valid)
		incident.AddNetwork(valid)
		incident.AddNetwork(invalid)

		if len(incident.Networks) != 1 {
			t.Fatalf("AddNetwork did not filter or deduplicate: %d", len(incident.Networks))
		}

		incident.SetNetworks([]matchers.Network{valid, invalid, valid})

		if len(incident.Networks) != 2 {
			t.Fatalf("SetNetworks did not filter: %d", len(incident.Networks))
		}

		incident.RemoveNetwork(valid)

		if len(incident.Networks) != 1 {
			t.Fatalf("RemoveNetwork failed: %d", len(incident.Networks))
		}

		incident.RemoveNetwork(invalid)

		if len(incident.Networks) != 1 {
			t.Fatalf("RemoveNetwork removed an unknown value: %d", len(incident.Networks))
		}

		incident.RemoveNetwork(valid)

		if len(incident.Networks) != 0 {
			t.Fatalf("RemoveNetwork did not remove the last value: %d", len(incident.Networks))
		}

	})

	t.Run("packages", func(t *testing.T) {

		incident := IncidentTestBase("Recon")
		valid := IncidentTestPackage()
		invalid := matchers.NewPackage()

		incident.AddPackage(valid)
		incident.AddPackage(valid)
		incident.AddPackage(invalid)

		if len(incident.Packages) != 1 {
			t.Fatalf("AddPackage did not filter or deduplicate: %d", len(incident.Packages))
		}

		incident.SetPackages([]matchers.Package{valid, invalid, valid})

		if len(incident.Packages) != 2 {
			t.Fatalf("SetPackages did not filter: %d", len(incident.Packages))
		}

		incident.RemovePackage(valid)

		if len(incident.Packages) != 1 {
			t.Fatalf("RemovePackage failed: %d", len(incident.Packages))
		}

		incident.RemovePackage(invalid)

		if len(incident.Packages) != 1 {
			t.Fatalf("RemovePackage removed an unknown value: %d", len(incident.Packages))
		}

		incident.RemovePackage(valid)

		if len(incident.Packages) != 0 {
			t.Fatalf("RemovePackage did not remove the last value: %d", len(incident.Packages))
		}

	})

	t.Run("products", func(t *testing.T) {

		incident := IncidentTestBase("Recon")
		valid := IncidentTestProduct()
		invalid := matchers.NewProduct()

		incident.AddProduct(valid)
		incident.AddProduct(valid)
		incident.AddProduct(invalid)

		if len(incident.Products) != 1 {
			t.Fatalf("AddProduct did not filter or deduplicate: %d", len(incident.Products))
		}

		incident.SetProducts([]matchers.Product{valid, invalid, valid})

		if len(incident.Products) != 2 {
			t.Fatalf("SetProducts did not filter: %d", len(incident.Products))
		}

		incident.RemoveProduct(valid)

		if len(incident.Products) != 1 {
			t.Fatalf("RemoveProduct failed: %d", len(incident.Products))
		}

		incident.RemoveProduct(invalid)

		if len(incident.Products) != 1 {
			t.Fatalf("RemoveProduct removed an unknown value: %d", len(incident.Products))
		}

		incident.RemoveProduct(valid)

		if len(incident.Products) != 0 {
			t.Fatalf("RemoveProduct did not remove the last value: %d", len(incident.Products))
		}

	})

	t.Run("programs", func(t *testing.T) {

		incident := IncidentTestBase("Recon")
		valid := IncidentTestProgram()
		invalid := matchers.NewProgram()

		incident.AddProgram(valid)
		incident.AddProgram(valid)
		incident.AddProgram(invalid)

		if len(incident.Programs) != 1 {
			t.Fatalf("AddProgram did not filter or deduplicate: %d", len(incident.Programs))
		}

		incident.SetPrograms([]matchers.Program{valid, invalid, valid})

		if len(incident.Programs) != 2 {
			t.Fatalf("SetPrograms did not filter: %d", len(incident.Programs))
		}

		incident.RemoveProgram(valid)

		if len(incident.Programs) != 1 {
			t.Fatalf("RemoveProgram failed: %d", len(incident.Programs))
		}

		incident.RemoveProgram(invalid)

		if len(incident.Programs) != 1 {
			t.Fatalf("RemoveProgram removed an unknown value: %d", len(incident.Programs))
		}

		incident.RemoveProgram(valid)

		if len(incident.Programs) != 0 {
			t.Fatalf("RemoveProgram did not remove the last value: %d", len(incident.Programs))
		}

	})

	t.Run("users", func(t *testing.T) {

		incident := IncidentTestBase("Recon")
		valid := IncidentTestUser()
		invalid := matchers.NewUser()

		incident.AddUser(valid)
		incident.AddUser(valid)
		incident.AddUser(invalid)

		if len(incident.Users) != 1 {
			t.Fatalf("AddUser did not filter or deduplicate: %d", len(incident.Users))
		}

		incident.SetUsers([]matchers.User{valid, invalid, valid})

		if len(incident.Users) != 2 {
			t.Fatalf("SetUsers did not filter: %d", len(incident.Users))
		}

		incident.RemoveUser(valid)

		if len(incident.Users) != 1 {
			t.Fatalf("RemoveUser failed: %d", len(incident.Users))
		}

		incident.RemoveUser(invalid)

		if len(incident.Users) != 1 {
			t.Fatalf("RemoveUser removed an unknown value: %d", len(incident.Users))
		}

		incident.RemoveUser(valid)

		if len(incident.Users) != 0 {
			t.Fatalf("RemoveUser did not remove the last value: %d", len(incident.Users))
		}

	})

	t.Run("vulnerabilities", func(t *testing.T) {

		incident := IncidentTestBase("Recon")
		valid := IncidentTestVulnerability()
		invalid := matchers.NewVulnerability()

		incident.AddVulnerability(valid)
		incident.AddVulnerability(valid)
		incident.AddVulnerability(invalid)

		if len(incident.Vulnerabilities) != 1 {
			t.Fatalf("AddVulnerability did not filter or deduplicate: %d", len(incident.Vulnerabilities))
		}

		incident.SetVulnerabilities([]matchers.Vulnerability{valid, invalid, valid})

		if len(incident.Vulnerabilities) != 2 {
			t.Fatalf("SetVulnerabilities did not filter: %d", len(incident.Vulnerabilities))
		}

		incident.RemoveVulnerability(valid)

		if len(incident.Vulnerabilities) != 1 {
			t.Fatalf("RemoveVulnerability failed: %d", len(incident.Vulnerabilities))
		}

		incident.RemoveVulnerability(invalid)

		if len(incident.Vulnerabilities) != 1 {
			t.Fatalf("RemoveVulnerability removed an unknown value: %d", len(incident.Vulnerabilities))
		}

		incident.RemoveVulnerability(valid)

		if len(incident.Vulnerabilities) != 0 {
			t.Fatalf("RemoveVulnerability did not remove the last value: %d", len(incident.Vulnerabilities))
		}

	})

	t.Run("weaknesses", func(t *testing.T) {

		incident := IncidentTestBase("Recon")
		valid := IncidentTestWeakness()
		invalid := matchers.NewWeakness()

		incident.AddWeakness(valid)
		incident.AddWeakness(valid)
		incident.AddWeakness(invalid)

		if len(incident.Weaknesses) != 1 {
			t.Fatalf("AddWeakness did not filter or deduplicate: %d", len(incident.Weaknesses))
		}

		incident.SetWeaknesses([]matchers.Weakness{valid, invalid, valid})

		if len(incident.Weaknesses) != 2 {
			t.Fatalf("SetWeaknesses did not filter: %d", len(incident.Weaknesses))
		}

		incident.RemoveWeakness(valid)

		if len(incident.Weaknesses) != 1 {
			t.Fatalf("RemoveWeakness failed: %d", len(incident.Weaknesses))
		}

		incident.RemoveWeakness(invalid)

		if len(incident.Weaknesses) != 1 {
			t.Fatalf("RemoveWeakness removed an unknown value: %d", len(incident.Weaknesses))
		}

		incident.RemoveWeakness(valid)

		if len(incident.Weaknesses) != 0 {
			t.Fatalf("RemoveWeakness did not remove the last value: %d", len(incident.Weaknesses))
		}

	})

}

func TestIncidentHash(t *testing.T) {

	incident := IncidentTestRich()
	hash := incident.Hash()

	if hash == "" {
		t.Fatal("expected a non-empty hash")
	}

	if incident.Hash() != hash {
		t.Fatal("expected a deterministic hash")
	}

	empty := Incident{}

	if empty.Hash() != "" {
		t.Fatalf("expected an empty hash, got %q", empty.Hash())
	}

	withPackage := IncidentTestBase("Recon")
	withPackage.AddPackage(IncidentTestPackage())

	if withPackage.Hash() == "" {
		t.Fatal("expected a non-empty hash for an incident with a package")
	}

	withProduct := IncidentTestBase("Recon")
	withProduct.AddProduct(IncidentTestProduct())

	if withProduct.Hash() == withPackage.Hash() {
		t.Fatal("expected products to change the hash")
	}

	withWeakness := IncidentTestBase("Recon")
	withWeakness.AddWeakness(IncidentTestWeakness())

	if withWeakness.Hash() == withPackage.Hash() {
		t.Fatal("expected weaknesses to change the hash")
	}

}
