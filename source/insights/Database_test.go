package insights

import (
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"tholian-core/caches"
	"tholian-core/matchers"
	"tholian-core/structs"
)

func testDatabase(t *testing.T) *Database {
	t.Helper()
	return NewDatabase(t.TempDir())
}

func testUser(name string) matchers.User {
	user := matchers.NewUser()
	user.Name = name
	return user
}

func testPackage(name string) matchers.Package {
	pkg := matchers.NewPackage()
	pkg.Name = name
	return pkg
}

func testDistribution(name string) matchers.Distribution {
	distribution := matchers.NewDistribution()
	distribution.Name = name
	return distribution
}

func newCredential(keys []string, passwords []string) *structs.Credential {
	credential := structs.NewCredential("alice", "user")
	for _, key := range keys {
		credential.AddKey([]byte(key))
	}
	for _, password := range passwords {
		credential.AddPassword(password)
	}
	return &credential
}

func newIncident(users ...string) *structs.Incident {
	incident := structs.NewIncident("system-a", "Destroy")
	incident.SetDatetime("2026-01-02T03:04:05Z")
	if len(users) == 0 {
		users = []string{"alice"}
	}
	for _, user := range users {
		incident.AddUser(testUser(user))
	}
	return &incident
}

func newResponse(packages ...string) *structs.Response {
	response := structs.NewResponse("system-a", "Update")
	response.SetDatetime("2026-01-02T03:04:05Z")
	if len(packages) == 0 {
		packages = []string{"alpha"}
	}
	for _, pkg := range packages {
		response.AddPackage(testPackage(pkg))
	}
	return &response
}

func newProduct() *structs.Product {
	product := structs.NewProduct("software")
	product.SetVendor("acme")
	product.SetProduct("widget")
	product.SetVersion("1.0")
	product.SetState("published")
	return &product
}

func newVulnerability(name string, distributions ...string) *structs.Vulnerability {
	vulnerability := structs.NewVulnerability(name)
	vulnerability.SetDatetime("2026-01-02T03:04:05Z")
	vulnerability.SetState("published")
	for _, distribution := range distributions {
		vulnerability.AddDistribution(testDistribution(distribution))
	}
	return &vulnerability
}

func readFile(t *testing.T, database *Database, file string) []byte {
	t.Helper()
	buffer := database.Read(file)
	if len(buffer) == 0 {
		t.Fatalf("expected file %s", file)
	}
	return buffer
}

func waitFor(t *testing.T, wait *sync.WaitGroup) {
	t.Helper()
	done := make(chan struct{})
	go func() {
		wait.Wait()
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("concurrent operations timed out")
	}
}

func TestFileBackedRoundTrip(t *testing.T) {

	database := testDatabase(t)

	credential := newCredential(nil, []string{"secret"})
	if database.WriteCredential(credential) == false {
		t.Fatal("cannot write credential")
	}
	if result := database.ReadCredential(credential.Name); result.Equal(*credential) == false {
		t.Fatal("credential round-trip failed")
	}

	incident := newIncident()
	if database.WriteIncident(incident) == false {
		t.Fatal("cannot write incident")
	}
	if result := database.ReadIncident(incident.Hash()); result.Equal(*incident) == false {
		t.Fatal("incident round-trip failed")
	}

	response := newResponse()
	if database.WriteResponse(response) == false {
		t.Fatal("cannot write response")
	}
	if result := database.ReadResponse(response.Hash()); result.Equal(*response) == false {
		t.Fatal("response round-trip failed")
	}

	product := newProduct()
	if database.WriteProduct(product) == false {
		t.Fatal("cannot write product")
	}
	if result := database.ReadProduct(product.Name); result.Equal(*product) == false {
		t.Fatal("product round-trip failed")
	}

	mitigation := structs.NewMitigation(*newIncident(), *newResponse())
	if database.WriteMitigation(&mitigation) == false {
		t.Fatal("cannot write mitigation")
	}
	if result := database.ReadMitigation(mitigation.Hash()); result.Equal(mitigation) == false {
		t.Fatal("mitigation round-trip failed")
	}

	vulnerability := newVulnerability("CVE-2026-1000", "ubuntu-24.04")
	if database.WriteVulnerability(vulnerability) == false {
		t.Fatal("cannot write vulnerability")
	}
	if result := database.ReadVulnerability(vulnerability.Name); result.Equal(*vulnerability) == false {
		t.Fatal("vulnerability round-trip failed")
	}

}

func TestGzipFallback(t *testing.T) {

	database := testDatabase(t)

	credential := newCredential(nil, []string{"secret"})
	if database.WriteCredential(credential) == false {
		t.Fatal("cannot write credential")
	}
	database.Remove("/credentials/embedded/" + credential.Name + ".json")
	if result := database.ReadCredential(credential.Name); result.Equal(*credential) == false {
		t.Fatal("credential gzip fallback failed")
	}

	incident := newIncident()
	if database.WriteIncident(incident) == false {
		t.Fatal("cannot write incident")
	}
	database.Remove("/incidents/embedded/" + incident.Hash() + ".json")
	if result := database.ReadIncident(incident.Hash()); result.Equal(*incident) == false {
		t.Fatal("incident gzip fallback failed")
	}

	response := newResponse()
	if database.WriteResponse(response) == false {
		t.Fatal("cannot write response")
	}
	database.Remove("/responses/embedded/" + response.Hash() + ".json")
	if result := database.ReadResponse(response.Hash()); result.Equal(*response) == false {
		t.Fatal("response gzip fallback failed")
	}

	product := newProduct()
	if database.WriteProduct(product) == false {
		t.Fatal("cannot write product")
	}
	database.Remove("/products/embedded/" + product.Name + ".json")
	if result := database.ReadProduct(product.Name); result.Equal(*product) == false {
		t.Fatal("product gzip fallback failed")
	}

	mitigation := structs.NewMitigation(*newIncident(), *newResponse())
	if database.WriteMitigation(&mitigation) == false {
		t.Fatal("cannot write mitigation")
	}
	database.Remove("/mitigations/embedded/" + mitigation.Hash() + ".json")
	if result := database.ReadMitigation(mitigation.Hash()); result.Equal(mitigation) == false {
		t.Fatal("mitigation gzip fallback failed")
	}

	vulnerability := newVulnerability("CVE-2026-1001")
	if database.WriteVulnerability(vulnerability) == false {
		t.Fatal("cannot write vulnerability")
	}
	database.Remove("/vulnerabilities/embedded/" + vulnerability.Name + ".json")
	if result := database.ReadVulnerability(vulnerability.Name); result.Equal(*vulnerability) == false {
		t.Fatal("vulnerability gzip fallback failed")
	}

}

func TestDeterministicSorting(t *testing.T) {

	first := testDatabase(t)
	second := testDatabase(t)

	credential_a := newCredential([]string{"b", "a"}, []string{"two", "one"})
	credential_b := newCredential([]string{"a", "b"}, []string{"one", "two"})
	if first.WriteCredential(credential_a) == false || second.WriteCredential(credential_b) == false {
		t.Fatal("cannot write credentials")
	}
	file_a := readFile(t, first, "/credentials/embedded/"+credential_a.Name+".json")
	file_b := readFile(t, second, "/credentials/embedded/"+credential_b.Name+".json")
	if string(file_a) != string(file_b) {
		t.Fatal("credential sorting is not deterministic")
	}

	vulnerability_a := newVulnerability("CVE-2026-2000")
	vulnerability_a.References = []string{"https://b.example", "https://a.example"}
	vulnerability_a.Packages = []matchers.Package{testPackage("zeta"), testPackage("alpha")}
	vulnerability_b := newVulnerability("CVE-2026-2000")
	vulnerability_b.References = []string{"https://a.example", "https://b.example"}
	vulnerability_b.Packages = []matchers.Package{testPackage("alpha"), testPackage("zeta")}
	if first.WriteVulnerability(vulnerability_a) == false || second.WriteVulnerability(vulnerability_b) == false {
		t.Fatal("cannot write vulnerabilities")
	}
	file_a = readFile(t, first, "/vulnerabilities/embedded/"+vulnerability_a.Name+".json")
	file_b = readFile(t, second, "/vulnerabilities/embedded/"+vulnerability_b.Name+".json")
	if string(file_a) != string(file_b) {
		t.Fatal("vulnerability sorting is not deterministic")
	}

	incident_a := newIncident("bob", "alice")
	incident_a.AddPackage(testPackage("zeta"))
	incident_a.AddPackage(testPackage("alpha"))
	incident_b := newIncident("alice", "bob")
	incident_b.AddPackage(testPackage("alpha"))
	incident_b.AddPackage(testPackage("zeta"))
	if first.WriteIncident(incident_a) == false || second.WriteIncident(incident_b) == false {
		t.Fatal("cannot write incidents")
	}
	if incident_a.Hash() != incident_b.Hash() {
		t.Fatal("incident sorting changes the hash")
	}
	file_a = readFile(t, first, "/incidents/embedded/"+incident_a.Hash()+".json")
	file_b = readFile(t, second, "/incidents/embedded/"+incident_b.Hash()+".json")
	if string(file_a) != string(file_b) {
		t.Fatal("incident sorting is not deterministic")
	}

	response_a := newResponse("zeta", "alpha")
	response_b := newResponse("alpha", "zeta")
	if first.WriteResponse(response_a) == false || second.WriteResponse(response_b) == false {
		t.Fatal("cannot write responses")
	}
	if response_a.Hash() != response_b.Hash() {
		t.Fatal("response sorting changes the hash")
	}
	file_a = readFile(t, first, "/responses/embedded/"+response_a.Hash()+".json")
	file_b = readFile(t, second, "/responses/embedded/"+response_b.Hash()+".json")
	if string(file_a) != string(file_b) {
		t.Fatal("response sorting is not deterministic")
	}

	mitigation_a := structs.NewMitigation(*newIncident("bob", "alice"), *newResponse("zeta", "alpha"))
	mitigation_b := structs.NewMitigation(*newIncident("alice", "bob"), *newResponse("alpha", "zeta"))
	if first.WriteMitigation(&mitigation_a) == false || second.WriteMitigation(&mitigation_b) == false {
		t.Fatal("cannot write mitigations")
	}
	if mitigation_a.Hash() != mitigation_b.Hash() {
		t.Fatal("mitigation sorting changes the hash")
	}
	file_a = readFile(t, first, "/mitigations/embedded/"+mitigation_a.Hash()+".json")
	file_b = readFile(t, second, "/mitigations/embedded/"+mitigation_b.Hash()+".json")
	if string(file_a) != string(file_b) {
		t.Fatal("mitigation sorting is not deterministic")
	}

}

func TestGeneratedEmbeds(t *testing.T) {

	database := testDatabase(t)

	distribution := structs.NewDistribution()
	distribution.SetName("ubuntu")
	database.Distributions.Add(distribution)

	if database.WriteIncident(newIncident()) == false {
		t.Fatal("cannot write incident")
	}
	if database.WriteResponse(newResponse()) == false {
		t.Fatal("cannot write response")
	}
	if database.WriteProduct(newProduct()) == false {
		t.Fatal("cannot write product")
	}
	mitigation := structs.NewMitigation(*newIncident(), *newResponse())
	if database.WriteMitigation(&mitigation) == false {
		t.Fatal("cannot write mitigation")
	}
	if database.WriteCredential(newCredential(nil, []string{"secret"})) == false {
		t.Fatal("cannot write credential")
	}
	if database.WriteVulnerability(newVulnerability("CVE-2026-3000", "ubuntu-24.04")) == false {
		t.Fatal("cannot write ubuntu vulnerability")
	}
	if database.WriteVulnerability(newVulnerability("CVE-2026-3001", "debian-12")) == false {
		t.Fatal("cannot write debian vulnerability")
	}

	if database.WriteIncidents() == false {
		t.Fatal("cannot generate incidents")
	}
	if database.WriteResponses() == false {
		t.Fatal("cannot generate responses")
	}
	if database.WriteProducts() == false {
		t.Fatal("cannot generate products")
	}
	if database.WriteMitigations() == false {
		t.Fatal("cannot generate mitigations")
	}
	if database.WriteCredentials() == false {
		t.Fatal("cannot generate credentials")
	}
	if database.WriteVulnerabilities() == false {
		t.Fatal("cannot generate vulnerabilities")
	}

	incident_guard := string(readFile(t, database, "/incidents/Incidents_embed.guard.go"))
	incident_oversight := string(readFile(t, database, "/incidents/Incidents_embed.oversight.go"))
	if strings.Contains(incident_guard, "guard_ubuntu") == false {
		t.Fatal("incident guard embed is not distribution aware")
	}
	if strings.Contains(incident_oversight, "//go:build oversight") == false {
		t.Fatal("incident oversight embed has the wrong build tag")
	}

	response_guard := string(readFile(t, database, "/responses/Responses_embed.guard.go"))
	response_oversight := string(readFile(t, database, "/responses/Responses_embed.oversight.go"))
	if strings.Contains(response_guard, "guard_ubuntu") == false {
		t.Fatal("response guard embed is not distribution aware")
	}
	if strings.Contains(response_oversight, "//go:build oversight") == false {
		t.Fatal("response oversight embed has the wrong build tag")
	}

	product_scout := string(readFile(t, database, "/products/Products_embed.scout.go"))
	product_guard := string(readFile(t, database, "/products/Products_embed.guard.go"))
	product_oversight := string(readFile(t, database, "/products/Products_embed.oversight.go"))
	if strings.Contains(product_scout, "scout_ubuntu") == false {
		t.Fatal("product scout embed is not distribution aware")
	}
	if strings.Contains(product_guard, "guard_ubuntu") == false {
		t.Fatal("product guard embed is not distribution aware")
	}
	if strings.Contains(product_oversight, "//go:build oversight") == false {
		t.Fatal("product oversight embed has the wrong build tag")
	}
	if database.Exists("/products/Products_embed.scout_ubuntu.go") {
		t.Fatal("product embeds must not be distribution specific")
	}

	mitigation_guard := string(readFile(t, database, "/mitigations/Mitigations_embed.guard.go"))
	if strings.Contains(mitigation_guard, "guard_ubuntu") == false {
		t.Fatal("mitigation guard embed is not distribution aware")
	}

	vulnerability_scout := string(readFile(t, database, "/vulnerabilities/Vulnerabilities_embed.scout.go"))
	vulnerability_ubuntu := string(readFile(t, database, "/vulnerabilities/Vulnerabilities_embed.guard_ubuntu.go"))
	vulnerability_debian := string(readFile(t, database, "/vulnerabilities/Vulnerabilities_embed.guard_debian.go"))
	if strings.Contains(vulnerability_scout, "//go:build scout") == false {
		t.Fatal("vulnerability scout embed has the wrong build tag")
	}
	if strings.Contains(vulnerability_ubuntu, "CVE-2026-3000") == false || strings.Contains(vulnerability_ubuntu, "CVE-2026-3001") {
		t.Fatal("ubuntu vulnerability embed contains the wrong records")
	}
	if strings.Contains(vulnerability_debian, "CVE-2026-3001") == false || strings.Contains(vulnerability_debian, "CVE-2026-3000") {
		t.Fatal("debian vulnerability embed contains the wrong records")
	}

	generated := make([]string, 0)
	for _, folder := range []string{"/credentials", "/incidents", "/mitigations", "/products", "/responses", "/vulnerabilities"} {
		generated = append(generated, database.Scan(folder, "*.go")...)
	}

	fileset := token.NewFileSet()

	for _, file := range generated {

		content := readFile(t, database, file)

		if strings.Contains(string(content), "tholian-endpoint") {
			t.Fatalf("generated file %s uses the old module name", file)
		}

		if _, err := parser.ParseFile(fileset, file, content, parser.AllErrors); err != nil {
			t.Fatalf("generated file %s is invalid: %s", file, err.Error())
		}

	}

}

func TestGeneratedCacheConstructors(t *testing.T) {

	if caches.NewMitigations() == nil ||
		caches.NewIncidents() == nil ||
		caches.NewResponses() == nil {
		t.Fatal("generated cache constructors require arguments")
	}

}

func TestEmptyEmbeds(t *testing.T) {

	database := testDatabase(t)

	if database.WriteIncidents() == false {
		t.Fatal("cannot generate empty incidents")
	}
	if database.WriteProducts() == false {
		t.Fatal("cannot generate empty products")
	}

	for _, file := range []string{"/incidents/Incidents_embed.guard.go", "/products/Products_embed.scout.go"} {

		content := readFile(t, database, file)

		if strings.Contains(string(content), "//go:embed") {
			t.Fatalf("empty embed %s contains a go:embed directive", file)
		}

		if _, err := parser.ParseFile(token.NewFileSet(), file, content, parser.AllErrors); err != nil {
			t.Fatalf("empty embed %s is invalid: %s", file, err.Error())
		}

	}

}

func TestConcurrentAccess(t *testing.T) {

	database := testDatabase(t)

	t.Run("credential", func(t *testing.T) {

		var wait sync.WaitGroup

		for worker := 0; worker < 8; worker++ {
			wait.Add(1)
			go func() {
				defer wait.Done()
				for index := 0; index < 50; index++ {
					credential := newCredential(nil, []string{"secret"})
					database.WriteCredential(credential)
					database.ReadCredential(credential.Name)
					if index%10 == 0 {
						database.RemoveCredential(credential.Name)
					}
				}
			}()
		}

		waitFor(t, &wait)

	})

	t.Run("incident", func(t *testing.T) {

		var wait sync.WaitGroup

		for worker := 0; worker < 8; worker++ {
			wait.Add(1)
			go func() {
				defer wait.Done()
				for index := 0; index < 50; index++ {
					incident := newIncident()
					database.WriteIncident(incident)
					database.ReadIncident(incident.Hash())
					if index%10 == 0 {
						database.RemoveIncident(incident.Hash())
					}
				}
			}()
		}

		waitFor(t, &wait)

	})

	t.Run("vulnerability", func(t *testing.T) {

		var wait sync.WaitGroup

		for worker := 0; worker < 8; worker++ {
			wait.Add(1)
			go func() {
				defer wait.Done()
				for index := 0; index < 50; index++ {
					vulnerability := newVulnerability("CVE-2026-4000", "ubuntu-24.04")
					database.WriteVulnerability(vulnerability)
					database.ReadVulnerability(vulnerability.Name)
					if index%10 == 0 {
						database.RemoveVulnerability(vulnerability.Name)
					}
				}
			}()
		}

		waitFor(t, &wait)

	})

}

func TestPathSafety(t *testing.T) {

	database := testDatabase(t)

	credential := structs.NewCredential("../escape", "user")
	credential.AddPassword("secret")

	if database.WriteCredential(&credential) {
		t.Fatal("credential with an unsafe name was written")
	}

	if _, err := os.Stat(filepath.Join(filepath.Dir(database.Folder), "escape.json")); err == nil {
		t.Fatal("credential write escaped the database folder")
	}

	if database.ReadCredential("../escape").IsValid() {
		t.Fatal("credential with an unsafe name was read")
	}

}

func TestVulnerabilityIndex(t *testing.T) {

	database := testDatabase(t)

	vulnerability := newVulnerability("CVE-2026-6000")
	distribution := testDistribution("ubuntu-24.04")
	distribution.Vendor = "Canonical"
	vulnerability.AddDistribution(distribution)

	if database.WriteVulnerability(vulnerability) == false {
		t.Fatal("cannot write vulnerability")
	}

	if len(database.QueryVulnerabilitiesByDistributionVendor("Canonical")) != 1 {
		t.Fatal("vulnerability was not indexed")
	}

	database.RemoveVulnerability(vulnerability.Name)

	if len(database.QueryVulnerabilitiesByDistributionVendor("Canonical")) != 0 {
		t.Fatal("vulnerability is still indexed after removal")
	}

}

func TestCleanup(t *testing.T) {

	database := testDatabase(t)

	published := newProduct()
	if database.WriteProduct(published) == false {
		t.Fatal("cannot write product")
	}

	edited := newProduct()
	edited.SetProduct("edited")
	edited.SetState("edited")
	if database.WriteProduct(edited) == false {
		t.Fatal("cannot write edited product")
	}

	database.CleanupProducts()

	if database.Exists("/products/embedded/" + published.Name + ".json") {
		t.Fatal("cleanup did not remove the published product")
	}
	if database.Exists("/products/embedded/"+edited.Name+".json") == false {
		t.Fatal("cleanup removed the edited product")
	}

	vulnerability := newVulnerability("CVE-2026-5000")
	if database.WriteVulnerability(vulnerability) == false {
		t.Fatal("cannot write vulnerability")
	}

	edited_vulnerability := newVulnerability("CVE-2026-5001")
	edited_vulnerability.SetState("edited")
	if database.WriteVulnerability(edited_vulnerability) == false {
		t.Fatal("cannot write edited vulnerability")
	}

	database.CleanupVulnerabilities()

	if database.Exists("/vulnerabilities/embedded/" + vulnerability.Name + ".json") {
		t.Fatal("cleanup did not remove the published vulnerability")
	}
	if database.Exists("/vulnerabilities/embedded/"+edited_vulnerability.Name+".json") == false {
		t.Fatal("cleanup removed the edited vulnerability")
	}

	incident := newIncident()
	if database.WriteIncident(incident) == false {
		t.Fatal("cannot write incident")
	}

	database.CleanupIncidents()

	if database.Exists("/incidents/embedded/" + incident.Hash() + ".json") {
		t.Fatal("cleanup did not remove the incident")
	}

}
