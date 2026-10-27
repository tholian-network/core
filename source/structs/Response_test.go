package structs

import "testing"
import "tholian-core/matchers"
import "tholian-core/types"

func ResponseTestConnection() matchers.Connection {

	connection := matchers.NewConnection()
	connection.Host = "1.2.3.4"
	connection.Port = 1234
	connection.Protocol = "tcp"
	connection.Type = "client"

	return connection

}

func ResponseTestDistribution() matchers.Distribution {

	distribution := matchers.NewDistribution()
	distribution.SetName("debian")
	distribution.SetVersion("12")
	distribution.SetVendor("debian")

	return distribution

}

func ResponseTestNetwork() matchers.Network {

	network := matchers.NewNetwork()
	network.SetSubnet("1.1.1.0/24")

	return network

}

func ResponseTestPackage() matchers.Package {

	pkg := matchers.NewPackage()
	pkg.SetName("openssl")

	return pkg

}

func ResponseTestProduct() matchers.Product {

	product := matchers.NewProduct()
	product.SetVendor("acme")
	product.SetProduct("widget")
	product.SetVersion("1-2-3")
	product.SetType("software")

	return product

}

func ResponseTestProgram() matchers.Program {

	program := matchers.NewProgram()
	program.SetName("stealth.mjs")

	return program

}

func ResponseTestUser() matchers.User {

	user := matchers.NewUser()
	user.SetName("alice")

	return user

}

func ResponseTestVulnerability() matchers.Vulnerability {

	vulnerability := matchers.NewVulnerability()
	vulnerability.SetName("CVE-2021-1234")

	return vulnerability

}

func ResponseTestWeakness() matchers.Weakness {

	weakness := matchers.NewWeakness()
	weakness.SetName("CWE-79")

	return weakness

}

func ResponseTestBase(typ string) Response {

	response := NewResponse("machine", typ)
	response.SetDatetime("2024-01-02T03:04:05Z")

	return response

}

func ResponseTestRich() Response {

	response := ResponseTestBase("Permit")
	response.AddConnection(ResponseTestConnection())
	response.AddDistribution(ResponseTestDistribution())
	response.AddNetwork(ResponseTestNetwork())
	response.AddPackage(ResponseTestPackage())
	response.AddProduct(ResponseTestProduct())
	response.AddProgram(ResponseTestProgram())
	response.AddUser(ResponseTestUser())
	response.AddVulnerability(ResponseTestVulnerability())
	response.AddWeakness(ResponseTestWeakness())

	return response

}

func TestResponseNew(t *testing.T) {

	response := NewResponse(" machine ", "Update")

	if response.System != "machine" {
		t.Fatalf("unexpected system: %q", response.System)
	}

	if response.Type != "Update" {
		t.Fatalf("unexpected type: %q", response.Type)
	}

	if response.Datetime.IsValid() == false {
		t.Fatal("new response has an invalid datetime")
	}

	if len(response.Connections) != 0 ||
		len(response.Distributions) != 0 ||
		len(response.Networks) != 0 ||
		len(response.Packages) != 0 ||
		len(response.Products) != 0 ||
		len(response.Programs) != 0 ||
		len(response.Users) != 0 ||
		len(response.Vulnerabilities) != 0 ||
		len(response.Weaknesses) != 0 {
		t.Fatal("new response has non-empty collections")
	}

	invalid := NewResponse("machine", "Bogus")

	if invalid.Type != "none" {
		t.Fatalf("invalid type was accepted: %q", invalid.Type)
	}

}

func TestResponseSetSystemAndType(t *testing.T) {

	response := NewResponse("", "")

	response.SetSystem("  machine  ")

	if response.System != "machine" {
		t.Fatalf("unexpected system: %q", response.System)
	}

	types := []string{"Update", "Forbid", "Permit", "Recon", "Intel", "Conquer", "Persist", "Exfil", "Destroy"}

	for _, typ := range types {

		response.SetType(typ)

		if response.Type != typ {
			t.Fatalf("SetType(%q) = %q", typ, response.Type)
		}

	}

	response.SetType("Bogus")

	if response.Type != "Destroy" {
		t.Fatalf("invalid type overwrote value: %q", response.Type)
	}

}

func TestResponseSetDatetime(t *testing.T) {

	response := NewResponse("machine", "Update")
	original := response.Datetime

	response.SetDatetime("invalid")

	if response.Datetime.Equal(original) == false {
		t.Fatal("invalid datetime overwrote value")
	}

	expected := types.ToDatetime("2020-02-03T04:05:06Z")
	response.SetDatetime("2020-02-03T04:05:06Z")

	if response.Datetime.Equal(expected) == false {
		t.Fatalf("unexpected datetime: %s", response.Datetime.String())
	}

}

func TestResponseIsAfterIsBefore(t *testing.T) {

	types := []string{"Update", "Forbid", "Permit", "Recon", "Intel", "Conquer", "Persist", "Exfil", "Destroy"}
	indices := map[string]int{
		"Update":  0,
		"Forbid":  0,
		"Permit":  0,
		"Recon":   1,
		"Intel":   2,
		"Conquer": 3,
		"Persist": 4,
		"Exfil":   5,
		"Destroy": 6,
	}

	for i := 0; i < len(types); i++ {

		for j := i + 1; j < len(types); j++ {

			earlier := ResponseTestBase(types[i])
			later := ResponseTestBase(types[j])

			if indices[types[i]] < indices[types[j]] {

				if earlier.IsBefore(later) == false {
					t.Errorf("%s should be before %s at the same datetime", types[i], types[j])
				}

				if later.IsAfter(earlier) == false {
					t.Errorf("%s should be after %s at the same datetime", types[j], types[i])
				}

			} else {

				if earlier.IsBefore(later) == true {
					t.Errorf("%s should not be before %s at the same datetime", types[i], types[j])
				}

				if later.IsAfter(earlier) == true {
					t.Errorf("%s should not be after %s at the same datetime", types[j], types[i])
				}

			}

		}

	}

	earlier := ResponseTestBase("Update")
	earlier.SetDatetime("2024-01-01T00:00:00Z")
	later := ResponseTestBase("Destroy")
	later.SetDatetime("2024-01-03T00:00:00Z")

	if earlier.IsBefore(later) == false || later.IsAfter(earlier) == false {
		t.Fatal("datetime ordering failed")
	}

	if earlier.IsAfter(later) == true || later.IsBefore(earlier) == true {
		t.Fatal("reverse datetime ordering failed")
	}

	if earlier.IsBefore(earlier) == true || earlier.IsAfter(earlier) == true {
		t.Fatal("a response should not be before or after itself")
	}

}

func TestResponseEqual(t *testing.T) {

	left := ResponseTestRich()
	right := ResponseTestRich()

	if left.Equal(right) == false {
		t.Fatal("identical responses should be equal")
	}

	if left.Equal(left) == false {
		t.Fatal("a response should equal itself")
	}

	cases := []struct {
		name   string
		mutate func(*Response)
	}{
		{"system", func(value *Response) { value.System = "other" }},
		{"type", func(value *Response) { value.Type = "Destroy" }},
		{"datetime", func(value *Response) { value.SetDatetime("2020-01-01T00:00:00Z") }},
		{"connections-length", func(value *Response) { value.Connections = append(value.Connections, value.Connections[0]) }},
		{"connections-value", func(value *Response) { value.Connections[0].Host = "9.9.9.9" }},
		{"distributions-length", func(value *Response) { value.Distributions = append(value.Distributions, value.Distributions[0]) }},
		{"distributions-value", func(value *Response) { value.Distributions[0].Name = "ubuntu" }},
		{"networks-length", func(value *Response) { value.Networks = append(value.Networks, value.Networks[0]) }},
		{"networks-value", func(value *Response) { value.Networks[0].Name = "other" }},
		{"packages-length", func(value *Response) { value.Packages = append(value.Packages, value.Packages[0]) }},
		{"packages-value", func(value *Response) { value.Packages[0].Name = "nginx" }},
		{"products-length", func(value *Response) { value.Products = append(value.Products, value.Products[0]) }},
		{"products-value", func(value *Response) { value.Products[0].Vendor = "other" }},
		{"programs-length", func(value *Response) { value.Programs = append(value.Programs, value.Programs[0]) }},
		{"programs-value", func(value *Response) { value.Programs[0].Name = "other" }},
		{"users-length", func(value *Response) { value.Users = append(value.Users, value.Users[0]) }},
		{"users-value", func(value *Response) { value.Users[0].Name = "bob" }},
		{"vulnerabilities-length", func(value *Response) { value.Vulnerabilities = append(value.Vulnerabilities, value.Vulnerabilities[0]) }},
		{"vulnerabilities-value", func(value *Response) { value.Vulnerabilities[0].Name = "CVE-2020-0000" }},
		{"weaknesses-length", func(value *Response) { value.Weaknesses = append(value.Weaknesses, value.Weaknesses[0]) }},
		{"weaknesses-value", func(value *Response) { value.Weaknesses[0].Name = "CWE-999" }},
	}

	for _, c := range cases {

		t.Run(c.name, func(t *testing.T) {

			other := ResponseTestRich()
			c.mutate(&other)

			if left.Equal(other) == true {
				t.Fatalf("responses should differ for %s", c.name)
			}

		})

	}

}

func TestResponseIsValidMatrix(t *testing.T) {

	cases := []struct {
		typ      string
		populate func(*Response)
	}{
		{"Update", func(value *Response) {
			value.AddPackage(ResponseTestPackage())
		}},
		{"Forbid", func(value *Response) {
			value.AddConnection(ResponseTestConnection())
		}},
		{"Permit", func(value *Response) {
			value.AddUser(ResponseTestUser())
		}},
		{"Recon", func(value *Response) {
			value.AddNetwork(ResponseTestNetwork())
		}},
		{"Intel", func(value *Response) {
			value.AddConnection(ResponseTestConnection())
		}},
		{"Conquer", func(value *Response) {
			value.AddDistribution(ResponseTestDistribution())
			value.AddVulnerability(ResponseTestVulnerability())
			value.AddConnection(ResponseTestConnection())
		}},
		{"Persist", func(value *Response) {
			value.AddDistribution(ResponseTestDistribution())
			value.AddVulnerability(ResponseTestVulnerability())
			value.AddNetwork(ResponseTestNetwork())
		}},
		{"Exfil", func(value *Response) {
			value.AddNetwork(ResponseTestNetwork())
			value.AddProgram(ResponseTestProgram())
		}},
		{"Destroy", func(value *Response) {
			value.AddProgram(ResponseTestProgram())
		}},
	}

	for _, c := range cases {

		t.Run(c.typ, func(t *testing.T) {

			valid := ResponseTestBase(c.typ)
			c.populate(&valid)

			if valid.IsValid() == false {
				t.Fatalf("%s response should be valid", c.typ)
			}

			empty := ResponseTestBase(c.typ)

			if empty.IsValid() == true {
				t.Fatalf("%s response without evidence should be invalid", c.typ)
			}

		})

	}

	noSystem := ResponseTestBase("Destroy")
	noSystem.SetSystem("")
	noSystem.AddProgram(ResponseTestProgram())

	if noSystem.IsValid() == true {
		t.Fatal("response without system should be invalid")
	}

	noDatetime := ResponseTestBase("Destroy")
	noDatetime.Datetime = types.Datetime{}
	noDatetime.AddProgram(ResponseTestProgram())

	if noDatetime.IsValid() == true {
		t.Fatal("response without datetime should be invalid")
	}

}

func TestResponseMatches(t *testing.T) {

	response := ResponseTestBase("Permit")
	response.AddConnection(ResponseTestConnection())
	response.AddDistribution(ResponseTestDistribution())
	response.AddNetwork(ResponseTestNetwork())
	response.AddPackage(ResponseTestPackage())
	response.AddProduct(ResponseTestProduct())
	response.AddProgram(ResponseTestProgram())
	response.AddUser(ResponseTestUser())
	response.AddVulnerability(ResponseTestVulnerability())
	response.AddWeakness(ResponseTestWeakness())

	antique := NewAntique("apt", "openssl")
	antique.SetName("openssl")
	antique.SetVersion("1.0.0")
	antique.SetArchitecture("amd64")
	antique.SetVendor("debian")

	if response.MatchesAntique(antique) == false {
		t.Fatal("MatchesAntique should match")
	}

	otherAntique := antique
	otherAntique.SetName("nginx")

	if response.MatchesAntique(otherAntique) == true {
		t.Fatal("MatchesAntique should not match")
	}

	connection := types.ToConnection("[tcp] 1.2.3.4:1234 -> 5.6.7.8:443")

	if response.MatchesConnection(connection) == false {
		t.Fatal("MatchesConnection should match")
	}

	otherConnection := types.ToConnection("[udp] 9.9.9.9:53 -> 8.8.8.8:53")

	if response.MatchesConnection(otherConnection) == true {
		t.Fatal("MatchesConnection should not match")
	}

	targetConnection := types.ToConnection("[tcp] 9.9.9.9:9999 -> 1.2.3.4:1234")

	if response.MatchesConnection(targetConnection) == false {
		t.Fatal("MatchesConnection should match the target endpoint")
	}

	distribution := NewDistribution()
	distribution.SetName("debian")
	distribution.SetVersion("12")
	distribution.SetManager("apt")
	distribution.SetVendor("debian")

	if response.MatchesDistribution(distribution) == false {
		t.Fatal("MatchesDistribution should match")
	}

	otherDistribution := distribution
	otherDistribution.SetName("ubuntu")

	if response.MatchesDistribution(otherDistribution) == true {
		t.Fatal("MatchesDistribution should not match")
	}

	network := NewNetwork("trusted")
	network.AddSubnet(ToSubnet("1.1.1.0/24"))

	if response.MatchesNetwork(network) == false {
		t.Fatal("MatchesNetwork should match")
	}

	otherNetwork := NewNetwork("other")
	otherNetwork.AddSubnet(ToSubnet("2.2.2.0/24"))

	if response.MatchesNetwork(otherNetwork) == true {
		t.Fatal("MatchesNetwork should not match")
	}

	pkg := NewPackage("apt")
	pkg.SetName("openssl")
	pkg.SetVersion("1.2.3")
	pkg.SetArchitecture("amd64")
	pkg.SetVendor("debian")
	pkg.SetDatetime("2024-01-02T03:04:05Z")

	if response.MatchesPackage(pkg) == false {
		t.Fatal("MatchesPackage should match")
	}

	otherPackage := pkg
	otherPackage.SetName("nginx")

	if response.MatchesPackage(otherPackage) == true {
		t.Fatal("MatchesPackage should not match")
	}

	product := NewProduct("software")
	product.SetVendor("acme")
	product.SetProduct("widget")
	product.SetVersion("1-2-3")
	product.SetState("published")

	if response.MatchesProduct(product) == false {
		t.Fatal("MatchesProduct should match")
	}

	otherProduct := product
	otherProduct.SetVendor("other")

	if response.MatchesProduct(otherProduct) == true {
		t.Fatal("MatchesProduct should not match")
	}

	program := NewProgram(1234, "stealth.mjs")
	program.SetCommand("/usr/bin/node")
	program.SetArguments([]string{"serve"})

	if response.MatchesProgram(program) == false {
		t.Fatal("MatchesProgram should match")
	}

	otherProgram := program
	otherProgram.SetName("other.mjs")

	if response.MatchesProgram(otherProgram) == true {
		t.Fatal("MatchesProgram should not match")
	}

	update := NewUpdate("apt")
	update.SetName("openssl")
	update.SetVersion("1.2.3")
	update.SetArchitecture("amd64")

	if response.MatchesUpdate(update) == false {
		t.Fatal("MatchesUpdate should match")
	}

	otherUpdate := update
	otherUpdate.SetName("nginx")

	if response.MatchesUpdate(otherUpdate) == true {
		t.Fatal("MatchesUpdate should not match")
	}

	user := types.ToUser("alice:user:1000")

	if response.MatchesUser(user) == false {
		t.Fatal("MatchesUser should match")
	}

	otherUser := types.ToUser("bob:user:1001")

	if response.MatchesUser(otherUser) == true {
		t.Fatal("MatchesUser should not match")
	}

	vulnerability := NewVulnerability("CVE-2021-1234")
	vulnerability.SetState("published")

	if response.MatchesVulnerability(vulnerability) == false {
		t.Fatal("MatchesVulnerability should match")
	}

	otherVulnerability := NewVulnerability("CVE-2020-0000")
	otherVulnerability.SetState("published")

	if response.MatchesVulnerability(otherVulnerability) == true {
		t.Fatal("MatchesVulnerability should not match")
	}

	weakness := NewWeakness()
	weakness.SetName("CWE-79")
	weakness.AddImpact("Read Application Data")
	weakness.AddScope("Confidentiality")

	if response.MatchesWeakness(weakness) == false {
		t.Fatal("MatchesWeakness should match")
	}

	otherWeakness := weakness
	otherWeakness.SetName("CWE-89")

	if response.MatchesWeakness(otherWeakness) == true {
		t.Fatal("MatchesWeakness should not match")
	}

}

func TestResponseCollectionHelpers(t *testing.T) {

	t.Run("connections", func(t *testing.T) {

		response := ResponseTestBase("Permit")
		valid := ResponseTestConnection()
		invalid := matchers.NewConnection()

		response.AddConnection(valid)
		response.AddConnection(valid)
		response.AddConnection(invalid)

		if len(response.Connections) != 1 {
			t.Fatalf("AddConnection did not filter or deduplicate: %d", len(response.Connections))
		}

		response.SetConnections([]matchers.Connection{valid, invalid, valid})

		if len(response.Connections) != 2 {
			t.Fatalf("SetConnections did not filter: %d", len(response.Connections))
		}

		response.RemoveConnection(valid)

		if len(response.Connections) != 1 {
			t.Fatalf("RemoveConnection failed: %d", len(response.Connections))
		}

		response.RemoveConnection(invalid)

		if len(response.Connections) != 1 {
			t.Fatalf("RemoveConnection removed an unknown value: %d", len(response.Connections))
		}

		response.RemoveConnection(valid)

		if len(response.Connections) != 0 {
			t.Fatalf("RemoveConnection did not remove the last value: %d", len(response.Connections))
		}

	})

	t.Run("distributions", func(t *testing.T) {

		response := ResponseTestBase("Conquer")
		valid := ResponseTestDistribution()
		invalid := matchers.NewDistribution()

		response.AddDistribution(valid)
		response.AddDistribution(valid)
		response.AddDistribution(invalid)

		if len(response.Distributions) != 1 {
			t.Fatalf("AddDistribution did not filter or deduplicate: %d", len(response.Distributions))
		}

		response.SetDistributions([]matchers.Distribution{valid, invalid, valid})

		if len(response.Distributions) != 2 {
			t.Fatalf("SetDistributions did not filter: %d", len(response.Distributions))
		}

		response.RemoveDistribution(valid)

		if len(response.Distributions) != 1 {
			t.Fatalf("RemoveDistribution failed: %d", len(response.Distributions))
		}

		response.RemoveDistribution(invalid)

		if len(response.Distributions) != 1 {
			t.Fatalf("RemoveDistribution removed an unknown value: %d", len(response.Distributions))
		}

		response.RemoveDistribution(valid)

		if len(response.Distributions) != 0 {
			t.Fatalf("RemoveDistribution did not remove the last value: %d", len(response.Distributions))
		}

	})

	t.Run("networks", func(t *testing.T) {

		response := ResponseTestBase("Permit")
		valid := ResponseTestNetwork()
		invalid := matchers.NewNetwork()

		response.AddNetwork(valid)
		response.AddNetwork(valid)
		response.AddNetwork(invalid)

		if len(response.Networks) != 1 {
			t.Fatalf("AddNetwork did not filter or deduplicate: %d", len(response.Networks))
		}

		response.SetNetworks([]matchers.Network{valid, invalid, valid})

		if len(response.Networks) != 2 {
			t.Fatalf("SetNetworks did not filter: %d", len(response.Networks))
		}

		response.RemoveNetwork(valid)

		if len(response.Networks) != 1 {
			t.Fatalf("RemoveNetwork failed: %d", len(response.Networks))
		}

		response.RemoveNetwork(invalid)

		if len(response.Networks) != 1 {
			t.Fatalf("RemoveNetwork removed an unknown value: %d", len(response.Networks))
		}

		response.RemoveNetwork(valid)

		if len(response.Networks) != 0 {
			t.Fatalf("RemoveNetwork did not remove the last value: %d", len(response.Networks))
		}

	})

	t.Run("packages", func(t *testing.T) {

		response := ResponseTestBase("Update")
		valid := ResponseTestPackage()
		invalid := matchers.NewPackage()

		response.AddPackage(valid)
		response.AddPackage(valid)
		response.AddPackage(invalid)

		if len(response.Packages) != 1 {
			t.Fatalf("AddPackage did not filter or deduplicate: %d", len(response.Packages))
		}

		response.SetPackages([]matchers.Package{valid, invalid, valid})

		if len(response.Packages) != 2 {
			t.Fatalf("SetPackages did not filter: %d", len(response.Packages))
		}

		response.RemovePackage(valid)

		if len(response.Packages) != 1 {
			t.Fatalf("RemovePackage failed: %d", len(response.Packages))
		}

		response.RemovePackage(invalid)

		if len(response.Packages) != 1 {
			t.Fatalf("RemovePackage removed an unknown value: %d", len(response.Packages))
		}

		response.RemovePackage(valid)

		if len(response.Packages) != 0 {
			t.Fatalf("RemovePackage did not remove the last value: %d", len(response.Packages))
		}

	})

	t.Run("products", func(t *testing.T) {

		response := ResponseTestBase("Permit")
		valid := ResponseTestProduct()
		invalid := matchers.NewProduct()

		response.AddProduct(valid)
		response.AddProduct(valid)
		response.AddProduct(invalid)

		if len(response.Products) != 1 {
			t.Fatalf("AddProduct did not filter or deduplicate: %d", len(response.Products))
		}

		response.SetProducts([]matchers.Product{valid, invalid, valid})

		if len(response.Products) != 2 {
			t.Fatalf("SetProducts did not filter: %d", len(response.Products))
		}

		response.RemoveProduct(valid)

		if len(response.Products) != 1 {
			t.Fatalf("RemoveProduct failed: %d", len(response.Products))
		}

		response.RemoveProduct(invalid)

		if len(response.Products) != 1 {
			t.Fatalf("RemoveProduct removed an unknown value: %d", len(response.Products))
		}

		response.RemoveProduct(valid)

		if len(response.Products) != 0 {
			t.Fatalf("RemoveProduct did not remove the last value: %d", len(response.Products))
		}

	})

	t.Run("programs", func(t *testing.T) {

		response := ResponseTestBase("Destroy")
		valid := ResponseTestProgram()
		invalid := matchers.NewProgram()

		response.AddProgram(valid)
		response.AddProgram(valid)
		response.AddProgram(invalid)

		if len(response.Programs) != 1 {
			t.Fatalf("AddProgram did not filter or deduplicate: %d", len(response.Programs))
		}

		response.SetPrograms([]matchers.Program{valid, invalid, valid})

		if len(response.Programs) != 2 {
			t.Fatalf("SetPrograms did not filter: %d", len(response.Programs))
		}

		response.RemoveProgram(valid)

		if len(response.Programs) != 1 {
			t.Fatalf("RemoveProgram failed: %d", len(response.Programs))
		}

		response.RemoveProgram(invalid)

		if len(response.Programs) != 1 {
			t.Fatalf("RemoveProgram removed an unknown value: %d", len(response.Programs))
		}

		response.RemoveProgram(valid)

		if len(response.Programs) != 0 {
			t.Fatalf("RemoveProgram did not remove the last value: %d", len(response.Programs))
		}

	})

	t.Run("users", func(t *testing.T) {

		response := ResponseTestBase("Permit")
		valid := ResponseTestUser()
		invalid := matchers.NewUser()

		response.AddUser(valid)
		response.AddUser(valid)
		response.AddUser(invalid)

		if len(response.Users) != 1 {
			t.Fatalf("AddUser did not filter or deduplicate: %d", len(response.Users))
		}

		response.SetUsers([]matchers.User{valid, invalid, valid})

		if len(response.Users) != 2 {
			t.Fatalf("SetUsers did not filter: %d", len(response.Users))
		}

		response.RemoveUser(valid)

		if len(response.Users) != 1 {
			t.Fatalf("RemoveUser failed: %d", len(response.Users))
		}

		response.RemoveUser(invalid)

		if len(response.Users) != 1 {
			t.Fatalf("RemoveUser removed an unknown value: %d", len(response.Users))
		}

		response.RemoveUser(valid)

		if len(response.Users) != 0 {
			t.Fatalf("RemoveUser did not remove the last value: %d", len(response.Users))
		}

	})

	t.Run("vulnerabilities", func(t *testing.T) {

		response := ResponseTestBase("Conquer")
		valid := ResponseTestVulnerability()
		invalid := matchers.NewVulnerability()

		response.AddVulnerability(valid)
		response.AddVulnerability(valid)
		response.AddVulnerability(invalid)

		if len(response.Vulnerabilities) != 1 {
			t.Fatalf("AddVulnerability did not filter or deduplicate: %d", len(response.Vulnerabilities))
		}

		response.SetVulnerabilities([]matchers.Vulnerability{valid, invalid, valid})

		if len(response.Vulnerabilities) != 2 {
			t.Fatalf("SetVulnerabilities did not filter: %d", len(response.Vulnerabilities))
		}

		response.RemoveVulnerability(valid)

		if len(response.Vulnerabilities) != 1 {
			t.Fatalf("RemoveVulnerability failed: %d", len(response.Vulnerabilities))
		}

		response.RemoveVulnerability(invalid)

		if len(response.Vulnerabilities) != 1 {
			t.Fatalf("RemoveVulnerability removed an unknown value: %d", len(response.Vulnerabilities))
		}

		response.RemoveVulnerability(valid)

		if len(response.Vulnerabilities) != 0 {
			t.Fatalf("RemoveVulnerability did not remove the last value: %d", len(response.Vulnerabilities))
		}

	})

	t.Run("weaknesses", func(t *testing.T) {

		response := ResponseTestBase("Permit")
		valid := ResponseTestWeakness()
		invalid := matchers.NewWeakness()

		response.AddWeakness(valid)
		response.AddWeakness(valid)
		response.AddWeakness(invalid)

		if len(response.Weaknesses) != 1 {
			t.Fatalf("AddWeakness did not filter or deduplicate: %d", len(response.Weaknesses))
		}

		response.SetWeaknesses([]matchers.Weakness{valid, invalid, valid})

		if len(response.Weaknesses) != 2 {
			t.Fatalf("SetWeaknesses did not filter: %d", len(response.Weaknesses))
		}

		response.RemoveWeakness(valid)

		if len(response.Weaknesses) != 1 {
			t.Fatalf("RemoveWeakness failed: %d", len(response.Weaknesses))
		}

		response.RemoveWeakness(invalid)

		if len(response.Weaknesses) != 1 {
			t.Fatalf("RemoveWeakness removed an unknown value: %d", len(response.Weaknesses))
		}

		response.RemoveWeakness(valid)

		if len(response.Weaknesses) != 0 {
			t.Fatalf("RemoveWeakness did not remove the last value: %d", len(response.Weaknesses))
		}

	})

}

func TestResponseHash(t *testing.T) {

	response := ResponseTestRich()
	hash := response.Hash()

	if hash == "" {
		t.Fatal("expected a non-empty hash")
	}

	if response.Hash() != hash {
		t.Fatal("expected a deterministic hash")
	}

	empty := Response{}

	if empty.Hash() != "" {
		t.Fatalf("expected an empty hash, got %q", empty.Hash())
	}

	withPackage := ResponseTestBase("Update")
	withPackage.AddPackage(ResponseTestPackage())

	if withPackage.Hash() == "" {
		t.Fatal("expected a non-empty hash for a response with a package")
	}

	withProduct := ResponseTestBase("Update")
	withProduct.AddProduct(ResponseTestProduct())

	if withProduct.Hash() == withPackage.Hash() {
		t.Fatal("expected products to change the hash")
	}

	withWeakness := ResponseTestBase("Update")
	withWeakness.AddWeakness(ResponseTestWeakness())

	if withWeakness.Hash() == withPackage.Hash() {
		t.Fatal("expected weaknesses to change the hash")
	}

}
