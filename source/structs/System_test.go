package structs

import "tholian-core/matchers"
import "tholian-core/types"
import "encoding/json"
import "strings"
import "testing"

func systemTestValidDevice() Device {

	device := NewDevice("usb")
	device.SetName("Example USB Device")
	device.SetSystem("1234", "5678", "Example System")

	return device

}

func systemTestValidDrive() Drive {

	drive := NewDrive("sda", "local")
	drive.SetMountpoint("/")
	drive.SetSize(1000)
	drive.SetFree(400)
	drive.SetUsed(600)

	return drive

}

func systemTestValidDistribution() Distribution {

	distribution := NewDistribution()
	distribution.SetName("ExampleOS")

	return distribution

}

func systemTestValidIncident() Incident {

	incident := NewIncident("example-system", "Recon")
	incident.AddNetwork(matchers.ToNetwork("192.168.0.0/24"))
	incident.AddProgram(matchers.ToProgram("/usr/bin/example"))

	return incident

}

func systemTestValidResponse() Response {

	response := NewResponse("example-system", "Update")
	response.AddProgram(matchers.ToProgram("/usr/bin/example"))

	return response

}

func systemTestValidMitigation() Mitigation {
	return NewMitigation(systemTestValidIncident(), systemTestValidResponse())
}

func systemTestValidNetwork() Network {
	return NewNetwork("eth0")
}

func systemTestValidPackage() Package {

	pkg := NewPackage("apt")
	pkg.SetName("example-package")
	pkg.SetVersion("1.0.0")
	pkg.SetArchitecture("any")
	pkg.SetDatetime("2024-01-02T03:04:05Z")

	return pkg

}

func systemTestValidProgram() Program {

	program := NewProgram(1234, "example-program")
	program.SetCommand("/usr/bin/example")
	program.SetFolder("/opt/example")

	return program

}

func systemTestValidService() Program {

	service := NewProgram(5678, "example-service")
	service.SetCommand("/usr/bin/example-service")
	service.SetFolder("/opt/example-service")

	connection := matchers.NewConnection()
	connection.SetHost("127.0.0.1")
	connection.SetPort(8080)
	connection.SetType("server")
	service.AddConnection(connection)

	return service

}

func systemTestValidAntique() Antique {

	antique := NewAntique("apt", "http")
	antique.SetName("example-antique")
	antique.SetService("http")
	antique.SetVersion("1.0.0")

	return antique

}

func systemTestValidUpdate() Update {

	update := NewUpdate("apt")
	update.SetName("example-update")
	update.SetVersion("1.0.0")

	return update

}

func systemTestValidUser() types.User {
	return types.ToUser("example-user", 1000)
}

func systemTestPopulatedSystem() System {

	system := System{}

	system.SetName("example-system")
	system.SetHostname("example-host")
	system.SetDatetime("2024-01-02T03:04:05Z")
	system.SetDistribution(systemTestValidDistribution())
	system.SetBIOS(systemTestValidDevice())
	system.SetBoard(systemTestValidDevice())
	system.SetTimezone("UTC")
	system.SetCountry("DE")
	system.SetLocale("en_US")

	system.SetDevices([]Device{systemTestValidDevice()})
	system.SetDrives([]Drive{systemTestValidDrive()})
	system.SetIncidents([]Incident{systemTestValidIncident()})
	system.SetMitigations([]Mitigation{systemTestValidMitigation()})
	system.SetNetworks([]Network{systemTestValidNetwork()})
	system.SetPackages([]Package{systemTestValidPackage()})
	system.SetPrograms([]Program{systemTestValidProgram()})
	system.SetResponses([]Response{systemTestValidResponse()})
	system.SetServices([]Program{systemTestValidService()})
	system.SetAntiques([]Antique{systemTestValidAntique()})
	system.SetUpdates([]Update{systemTestValidUpdate()})
	system.SetUsers([]types.User{systemTestValidUser()})

	return system

}

func TestSystemNewSystemInvariants(t *testing.T) {

	system := NewSystem()

	if system.Devices == nil ||
		system.Drives == nil ||
		system.Incidents == nil ||
		system.Mitigations == nil ||
		system.Networks == nil ||
		system.Packages == nil ||
		system.Programs == nil ||
		system.Responses == nil ||
		system.Services == nil ||
		system.Antiques == nil ||
		system.Updates == nil ||
		system.Users == nil {
		t.Fatal("expected NewSystem to initialize all collections")
	}

	if strings.TrimSpace(system.Name) != system.Name {
		t.Fatalf("expected trimmed name, got %q", system.Name)
	}

	if strings.TrimSpace(system.Hostname) != system.Hostname {
		t.Fatalf("expected trimmed hostname, got %q", system.Hostname)
	}

	if system.Datetime.IsValid() == false {
		t.Fatal("expected NewSystem to set a valid datetime")
	}

}

func TestSystemMarshalJSONIndented(t *testing.T) {

	system := systemTestPopulatedSystem()

	data, err := system.MarshalJSON()

	if err != nil {
		t.Fatalf("MarshalJSON returned an error: %v", err)
	}

	if json.Valid(data) == false {
		t.Fatal("expected MarshalJSON to return valid JSON")
	}

	if strings.Contains(string(data), "\n\t\"name\":") == false {
		t.Fatalf("expected tab-indented JSON, got:\n%s", string(data))
	}

	var fields map[string]json.RawMessage

	if err := json.Unmarshal(data, &fields); err != nil {
		t.Fatalf("could not decode MarshalJSON output: %v", err)
	}

	representative := []string{
		"name",
		"hostname",
		"datetime",
		"distribution",
		"fingerprint",
		"bios",
		"board",
		"devices",
		"users",
	}

	for _, key := range representative {

		if _, ok := fields[key]; ok == false {
			t.Fatalf("expected JSON field %q", key)
		}

	}

	compact, err := json.Marshal(system)

	if err != nil {
		t.Fatalf("json.Marshal returned an error: %v", err)
	}

	if json.Valid(compact) == false {
		t.Fatal("expected json.Marshal to return valid JSON")
	}

	var compactFields map[string]json.RawMessage

	if err := json.Unmarshal(compact, &compactFields); err != nil {
		t.Fatalf("could not decode json.Marshal output: %v", err)
	}

	for _, key := range representative {

		if _, ok := compactFields[key]; ok == false {
			t.Fatalf("expected json.Marshal JSON field %q", key)
		}

	}

}

func TestSystemEqualAndHashCases(t *testing.T) {

	base := systemTestPopulatedSystem()
	other := base

	if base.Equal(other) == false {
		t.Fatal("expected identical systems to be equal")
	}

	if base.Hash() == "" {
		t.Fatal("expected a non-empty hash")
	}

	if base.Hash() != other.Hash() {
		t.Fatal("expected identical systems to have identical hashes")
	}

	var zero_a System
	var zero_b System

	if zero_a.Equal(zero_b) == false {
		t.Fatal("expected zero-value systems to be equal")
	}

	if zero_a.Hash() != "" {
		t.Fatal("expected zero-value system to have an empty hash")
	}

	tests := []struct {
		name   string
		mutate func(*System)
	}{
		{"name", func(system *System) { system.Name = "changed" }},
		{"hostname", func(system *System) { system.Hostname = "changed" }},
		{"datetime", func(system *System) { system.SetDatetime("2024-02-03T04:05:06Z") }},
		{"distribution", func(system *System) {
			distribution := NewDistribution()
			distribution.SetName("ChangedOS")
			system.Distribution = distribution
		}},
		{"fingerprint", func(system *System) { system.Fingerprint.Timezone = "Europe/Berlin" }},
		{"bios", func(system *System) {
			device := systemTestValidDevice()
			device.SetName("Changed BIOS")
			system.BIOS = device
		}},
		{"board", func(system *System) {
			device := systemTestValidDevice()
			device.SetName("Changed Board")
			system.Board = device
		}},
		{"devices", func(system *System) { system.Devices = nil }},
		{"drives", func(system *System) { system.Drives = nil }},
		{"incidents", func(system *System) { system.Incidents = nil }},
		{"mitigations", func(system *System) { system.Mitigations = nil }},
		{"networks", func(system *System) { system.Networks = nil }},
		{"packages", func(system *System) { system.Packages = nil }},
		{"programs", func(system *System) { system.Programs = nil }},
		{"responses", func(system *System) { system.Responses = nil }},
		{"services", func(system *System) { system.Services = nil }},
		{"antiques", func(system *System) { system.Antiques = nil }},
		{"updates", func(system *System) { system.Updates = nil }},
		{"users", func(system *System) { system.Users = nil }},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			mutated := base
			test.mutate(&mutated)

			if mutated.Equal(base) == true {
				t.Fatalf("expected systems with a different %s to be unequal", test.name)
			}

		})

	}

	mutated := base
	mutated.Name = "changed"

	if mutated.Hash() == base.Hash() {
		t.Fatal("expected a different name to produce a different hash")
	}

}

func TestSystemIsValidCases(t *testing.T) {

	base := System{}
	base.SetName("example-system")
	base.SetDatetime("2024-01-02T03:04:05Z")
	base.SetDistribution(systemTestValidDistribution())

	if base.IsValid() == false {
		t.Fatal("expected valid base system")
	}

	var zero System

	if zero.IsValid() == true {
		t.Fatal("expected zero-value system to be invalid")
	}

	tests := []struct {
		name   string
		mutate func(*System)
	}{
		{"datetime", func(system *System) { system.Datetime = types.Datetime{} }},
		{"distribution", func(system *System) { system.Distribution = Distribution{} }},
		{"incidents", func(system *System) { system.Incidents = []Incident{{}} }},
		{"mitigations", func(system *System) { system.Mitigations = []Mitigation{{}} }},
		{"networks", func(system *System) { system.Networks = []Network{{}} }},
		{"packages", func(system *System) { system.Packages = []Package{{}} }},
		{"programs", func(system *System) { system.Programs = []Program{{}} }},
		{"responses", func(system *System) { system.Responses = []Response{{}} }},
		{"services", func(system *System) { system.Services = []Program{{}} }},
		{"antiques", func(system *System) { system.Antiques = []Antique{{}} }},
		{"updates", func(system *System) { system.Updates = []Update{{}} }},
		{"users", func(system *System) { system.Users = []types.User{{}} }},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			mutated := base
			test.mutate(&mutated)

			if mutated.IsValid() == true {
				t.Fatalf("expected system with an invalid %s to be invalid", test.name)
			}

		})

	}

}

func TestSystemScalarSetters(t *testing.T) {

	system := System{}

	system.SetName("  example-system  ")

	if system.Name != "example-system" {
		t.Fatalf("expected trimmed name, got %q", system.Name)
	}

	system.SetHostname("  example-host  ")

	if system.Hostname != "example-host" {
		t.Fatalf("expected trimmed hostname, got %q", system.Hostname)
	}

	system.SetDatetime("2024-01-02T03:04:05Z")

	if system.Datetime.IsValid() == false {
		t.Fatal("expected valid datetime")
	}

	previous := system.Datetime
	system.SetDatetime("not-a-datetime")

	if system.Datetime.Equal(previous) == false {
		t.Fatal("expected invalid datetime to be ignored")
	}

	system.SetTimezone("  UTC  ")
	system.SetCountry("  DE  ")
	system.SetLocale("  en_US  ")

	if system.Fingerprint.Timezone != "UTC" {
		t.Fatalf("expected trimmed timezone, got %q", system.Fingerprint.Timezone)
	}

	if system.Fingerprint.Country != "DE" {
		t.Fatalf("expected trimmed country, got %q", system.Fingerprint.Country)
	}

	if system.Fingerprint.Locale != "en_US" {
		t.Fatalf("expected trimmed locale, got %q", system.Fingerprint.Locale)
	}

	device := systemTestValidDevice()
	system.SetBIOS(device)

	if system.BIOS.Equal(device) == false {
		t.Fatal("expected BIOS to be assigned")
	}

	system.SetBoard(device)

	if system.Board.Equal(device) == false {
		t.Fatal("expected Board to be assigned")
	}

	system.SetDistribution(Distribution{})

	if system.Distribution.IsValid() == true {
		t.Fatal("expected invalid distribution to be ignored")
	}

	distribution := systemTestValidDistribution()
	system.SetDistribution(distribution)

	if system.Distribution.Equal(distribution) == false {
		t.Fatal("expected valid distribution to be assigned")
	}

}

func TestSystemCollectionSetters(t *testing.T) {

	system := System{}

	device := systemTestValidDevice()
	system.SetDevices([]Device{{}, device})

	if len(system.Devices) != 1 || system.Devices[0].Equal(device) == false {
		t.Fatal("expected SetDevices to filter invalid devices")
	}

	drive := systemTestValidDrive()
	system.SetDrives([]Drive{{}, drive})

	if len(system.Drives) != 1 || system.Drives[0].Equal(drive) == false {
		t.Fatal("expected SetDrives to filter invalid drives")
	}

	incident := systemTestValidIncident()
	system.SetIncidents([]Incident{{}, incident})

	if len(system.Incidents) != 1 || system.Incidents[0].Equal(incident) == false {
		t.Fatal("expected SetIncidents to filter invalid incidents")
	}

	mitigation := systemTestValidMitigation()
	system.SetMitigations([]Mitigation{{}, mitigation})

	if len(system.Mitigations) != 1 || system.Mitigations[0].Equal(mitigation) == false {
		t.Fatal("expected SetMitigations to filter invalid mitigations")
	}

	network := systemTestValidNetwork()
	system.SetNetworks([]Network{{}, network})

	if len(system.Networks) != 1 || system.Networks[0].Equal(network) == false {
		t.Fatal("expected SetNetworks to filter invalid networks")
	}

	pkg := systemTestValidPackage()
	system.SetPackages([]Package{{}, pkg})

	if len(system.Packages) != 1 || system.Packages[0].Equal(pkg) == false {
		t.Fatal("expected SetPackages to filter invalid packages")
	}

	program := systemTestValidProgram()
	service := systemTestValidService()
	system.SetPrograms([]Program{{}, program, service})

	if len(system.Programs) != 1 || system.Programs[0].Equal(program) == false {
		t.Fatal("expected SetPrograms to filter non-programs")
	}

	response := systemTestValidResponse()
	system.SetResponses([]Response{{}, response})

	if len(system.Responses) != 1 || system.Responses[0].Equal(response) == false {
		t.Fatal("expected SetResponses to filter invalid responses")
	}

	system.SetServices([]Program{{}, program, service})

	if len(system.Services) != 1 || system.Services[0].Equal(service) == false {
		t.Fatal("expected SetServices to filter non-services")
	}

	antique := systemTestValidAntique()
	system.SetAntiques([]Antique{{}, antique})

	if len(system.Antiques) != 1 || system.Antiques[0].Equal(antique) == false {
		t.Fatal("expected SetAntiques to filter invalid antiques")
	}

	update := systemTestValidUpdate()
	system.SetUpdates([]Update{{}, update})

	if len(system.Updates) != 1 || system.Updates[0].Equal(update) == false {
		t.Fatal("expected SetUpdates to filter invalid updates")
	}

	user := systemTestValidUser()
	system.SetUsers([]types.User{{}, user})

	if len(system.Users) != 1 || system.Users[0].ID != user.ID || system.Users[0].Name != user.Name {
		t.Fatal("expected SetUsers to filter invalid users")
	}

}

func TestSystemAddMethods(t *testing.T) {

	system := System{
		Incidents:   make([]Incident, 0),
		Mitigations: make([]Mitigation, 0),
		Responses:   make([]Response, 0),
	}

	incident := systemTestValidIncident()
	system.AddIncident(Incident{})

	if len(system.Incidents) != 0 {
		t.Fatal("expected invalid incident to be ignored")
	}

	system.AddIncident(incident)

	if len(system.Incidents) != 1 || system.Incidents[0].Equal(incident) == false {
		t.Fatal("expected valid incident to be added")
	}

	mitigation := systemTestValidMitigation()
	system.AddMitigation(Mitigation{})

	if len(system.Mitigations) != 0 {
		t.Fatal("expected invalid mitigation to be ignored")
	}

	system.AddMitigation(mitigation)

	if len(system.Mitigations) != 1 || system.Mitigations[0].Equal(mitigation) == false {
		t.Fatal("expected valid mitigation to be added")
	}

	response := systemTestValidResponse()
	system.AddResponse(Response{})

	if len(system.Responses) != 0 {
		t.Fatal("expected invalid response to be ignored")
	}

	system.AddResponse(response)

	if len(system.Responses) != 1 || system.Responses[0].Equal(response) == false {
		t.Fatal("expected valid response to be added")
	}

}
