package structs

import "testing"
import "tholian-core/matchers"

func mitigationUnitPackageMatcher() matchers.Package {
	return matchers.ToPackage("curl")
}

func mitigationUnitIncident() Incident {

	incident := NewIncident("test-system", "Destroy")
	incident.AddPackage(mitigationUnitPackageMatcher())

	return incident

}

func mitigationUnitResponse() Response {

	response := NewResponse("test-system", "Update")
	response.AddPackage(mitigationUnitPackageMatcher())

	return response

}

func TestMitigationNewMitigation(t *testing.T) {

	incident := mitigationUnitIncident()
	response := mitigationUnitResponse()

	mitigation := NewMitigation(incident, response)

	if mitigation.IsValid() == false {
		t.Fatal("NewMitigation produced an invalid mitigation from valid parts")
	}

	if mitigation.Incident.Equal(incident) == false {
		t.Fatal("NewMitigation did not set the incident")
	}

	if mitigation.Response.Equal(response) == false {
		t.Fatal("NewMitigation did not set the response")
	}

	invalid := NewMitigation(Incident{}, Response{})

	if invalid.Incident.System != "" || invalid.Response.System != "" {
		t.Fatal("NewMitigation accepted invalid parts")
	}

	if invalid.IsValid() == true {
		t.Fatal("a mitigation with empty parts should be invalid")
	}

}

func TestMitigationEqual(t *testing.T) {

	incident := mitigationUnitIncident()
	response := mitigationUnitResponse()

	first := NewMitigation(incident, response)
	second := NewMitigation(incident, response)

	if first.Equal(second) == false {
		t.Fatal("identical mitigations should be equal")
	}

	otherIncident := incident
	otherIncident.System = "other-system"

	if first.Equal(Mitigation{Incident: otherIncident, Response: response}) == true {
		t.Fatal("mitigations with different incidents should not be equal")
	}

	otherResponse := response
	otherResponse.Type = "Destroy"

	if first.Equal(Mitigation{Incident: incident, Response: otherResponse}) == true {
		t.Fatal("mitigations with different responses should not be equal")
	}

}

func TestMitigationIsValid(t *testing.T) {

	incident := mitigationUnitIncident()
	response := mitigationUnitResponse()

	valid := NewMitigation(incident, response)

	if valid.IsValid() == false {
		t.Fatal("a complete mitigation should be valid")
	}

	invalidIncident := Mitigation{Incident: Incident{}, Response: response}

	if invalidIncident.IsValid() == true {
		t.Fatal("a mitigation with an invalid incident should be invalid")
	}

	invalidResponse := Mitigation{Incident: incident, Response: Response{}}

	if invalidResponse.IsValid() == true {
		t.Fatal("a mitigation with an invalid response should be invalid")
	}

}

func TestMitigationSetIncident(t *testing.T) {

	mitigation := NewMitigation(mitigationUnitIncident(), mitigationUnitResponse())
	original := mitigation.Incident

	mitigation.SetIncident(Incident{})

	if mitigation.Incident.Equal(original) == false {
		t.Fatal("SetIncident accepted an invalid incident")
	}

	replacement := NewIncident("other-system", "Destroy")
	replacement.AddPackage(mitigationUnitPackageMatcher())
	mitigation.SetIncident(replacement)

	if mitigation.Incident.System != "other-system" {
		t.Fatalf("SetIncident did not replace the incident: %+v", mitigation.Incident)
	}

}

func TestMitigationSetResponse(t *testing.T) {

	mitigation := NewMitigation(mitigationUnitIncident(), mitigationUnitResponse())
	original := mitigation.Response

	mitigation.SetResponse(Response{})

	if mitigation.Response.Equal(original) == false {
		t.Fatal("SetResponse accepted an invalid response")
	}

	replacement := NewResponse("other-system", "Update")
	replacement.AddPackage(mitigationUnitPackageMatcher())
	mitigation.SetResponse(replacement)

	if mitigation.Response.System != "other-system" {
		t.Fatalf("SetResponse did not replace the response: %+v", mitigation.Response)
	}

}

func TestMitigationHash(t *testing.T) {

	pkg := matchers.ToPackage("curl")
	program := matchers.ToProgram("node")
	connection := matchers.ToConnection("1.2.3.4:443TC")
	network := matchers.ToNetwork("AS64500")
	user := matchers.ToUser("alice")

	cases := []struct {
		name     string
		response Response
		want     string
	}{
		{"Update packages", Response{Type: "Update", Packages: []matchers.Package{pkg}}, "Update-Packages-" + pkg.Hash()},
		{"Update packages deduplicated", Response{Type: "Update", Packages: []matchers.Package{pkg, pkg}}, "Update-Packages-" + pkg.Hash()},
		{"Update programs", Response{Type: "Update", Programs: []matchers.Program{program}}, "Update-Programs-" + program.Hash()},
		{"Destroy packages", Response{Type: "Destroy", Packages: []matchers.Package{pkg}}, "Destroy-Packages-" + pkg.Hash()},
		{"Destroy programs", Response{Type: "Destroy", Programs: []matchers.Program{program}}, "Destroy-Programs-" + program.Hash()},
		{"Destroy users", Response{Type: "Destroy", Users: []matchers.User{user}}, "Destroy-Users-" + user.Hash()},
		{"Unknown type", Response{Type: "Unknown", Packages: []matchers.Package{pkg}}, ""},
		{"Update without targets", Response{Type: "Update"}, ""},
	}

	for _, typ := range []string{"Forbid", "Permit"} {

		cases = append(cases,
			struct {
				name     string
				response Response
				want     string
			}{typ + " connections", Response{Type: typ, Connections: []matchers.Connection{connection}}, typ + "-Connections-" + connection.Hash()},
			struct {
				name     string
				response Response
				want     string
			}{typ + " networks", Response{Type: typ, Networks: []matchers.Network{network}}, typ + "-Networks-" + network.Hash()},
			struct {
				name     string
				response Response
				want     string
			}{typ + " programs", Response{Type: typ, Programs: []matchers.Program{program}}, typ + "-Programs-" + program.Hash()},
			struct {
				name     string
				response Response
				want     string
			}{typ + " users", Response{Type: typ, Users: []matchers.User{user}}, typ + "-Users-" + user.Hash()},
		)

	}

	for _, typ := range []string{"Recon", "Intel", "Conquer", "Persist", "Exfil"} {

		cases = append(cases,
			struct {
				name     string
				response Response
				want     string
			}{typ + " connections", Response{Type: typ, Connections: []matchers.Connection{connection}}, typ + "-Connections-" + connection.Hash()},
			struct {
				name     string
				response Response
				want     string
			}{typ + " networks", Response{Type: typ, Networks: []matchers.Network{network}}, typ + "-Networks-" + network.Hash()},
		)

	}

	for _, c := range cases {

		mitigation := Mitigation{Response: c.response}

		if got := mitigation.Hash(); got != c.want {
			t.Errorf("%s: Hash() = %q, expected %q", c.name, got, c.want)
		}

	}

}
