package types

import (
	"encoding/json"
	"testing"
)

func TestIsManager(t *testing.T) {

	valid := []Manager{
		Manager(Any),
		ManagerAPK,
		ManagerAPT,
		ManagerDNF,
		ManagerPACMAN,
		ManagerRPM,
		ManagerTDNF,
		ManagerZYPPER,
		ManagerPKG,
		ManagerPKGSRC,
		ManagerMSI,
		ManagerCARGO,
		ManagerCHOCOLATEY,
		ManagerCOCOAPODS,
		ManagerCOMPOSER,
		ManagerCONAN,
		ManagerCONDA,
		ManagerCRAN,
		ManagerGEM,
		ManagerGO,
		ManagerGRADLE,
		ManagerHACKAGE,
		ManagerHEX,
		ManagerMAVEN,
		ManagerNPM,
		ManagerNUGET,
		ManagerPEAR,
		ManagerPHAR,
		ManagerPIP,
	}

	for _, manager := range valid {

		if IsManager(string(manager)) == false {
			t.Errorf("IsManager(%q) = false, expected true", manager)
		}

	}

	invalid := []string{"", "unknown", "APT", "npm "}

	for _, value := range invalid {

		if IsManager(value) == true {
			t.Errorf("IsManager(%q) = true, expected false", value)
		}

	}

}

func TestNewManager(t *testing.T) {

	manager := NewManager()

	if manager != Manager(Any) {
		t.Fatalf("NewManager() = %q, expected %q", manager, Any)
	}

	if manager.IsValid() == false {
		t.Fatal("NewManager() returned an invalid manager")
	}

}

func TestManagerParse(t *testing.T) {

	tests := []struct {
		input    string
		expected Manager
	}{
		{"any", Manager(Any)},
		{"apk", ManagerAPK},
		{"apt", ManagerAPT},
		{"dnf", ManagerDNF},
		{"pacman", ManagerPACMAN},
		{"rpm", ManagerRPM},
		{"tdnf", ManagerTDNF},
		{"zypper", ManagerZYPPER},
		{"pkg", ManagerPKG},
		{"pkgsrc", ManagerPKGSRC},
		{"msi", ManagerMSI},
		{"cargo", ManagerCARGO},
		{"chocolatey", ManagerCHOCOLATEY},
		{"cocoapods", ManagerCOCOAPODS},
		{"composer", ManagerCOMPOSER},
		{"conan", ManagerCONAN},
		{"conda", ManagerCONDA},
		{"cran", ManagerCRAN},
		{"gem", ManagerGEM},
		{"go", ManagerGO},
		{"gradle", ManagerGRADLE},
		{"hackage", ManagerHACKAGE},
		{"hex", ManagerHEX},
		{"maven", ManagerMAVEN},
		{"npm", ManagerNPM},
		{"nuget", ManagerNUGET},
		{"pear", ManagerPEAR},
		{"phar", ManagerPHAR},
		{"pip", ManagerPIP},
	}

	for _, test := range tests {

		var manager Manager

		if manager.Parse(test.input) == false {
			t.Errorf("Parse(%q) = false, expected true", test.input)
		} else if manager != test.expected {
			t.Errorf("Parse(%q) = %q, expected %q", test.input, manager, test.expected)
		}

	}

	invalid := []string{"", "unknown", "APT", "npm "}

	for _, value := range invalid {

		var manager Manager

		if manager.Parse(value) == true {
			t.Errorf("Parse(%q) = true, expected false", value)
		}

	}

}

func TestParseManager(t *testing.T) {

	manager := ParseManager("apt")

	if manager == nil || *manager != ManagerAPT {
		t.Fatalf("ParseManager() = %v, expected %q", manager, ManagerAPT)
	}

	if ParseManager("unknown") != nil {
		t.Fatal("ParseManager() returned a value for an invalid manager")
	}

	if ParseManager("") != nil {
		t.Fatal("ParseManager() returned a value for an empty manager")
	}

}

func TestToManager(t *testing.T) {

	if ToManager("npm") != ManagerNPM {
		t.Fatalf("ToManager(%q) = %q, expected %q", "npm", ToManager("npm"), ManagerNPM)
	}

	if ToManager("any") != Manager(Any) {
		t.Fatalf("ToManager(%q) = %q, expected %q", "any", ToManager("any"), Any)
	}

	if ToManager("unknown") != Manager("") {
		t.Fatalf("ToManager() = %q, expected zero value", ToManager("unknown"))
	}

}

func TestManagerEqual(t *testing.T) {

	if ManagerAPT.Equal(ManagerAPT) == false {
		t.Fatal("Equal() returned false for identical managers")
	}

	if ManagerAPT.Equal(ManagerDNF) == true {
		t.Fatal("Equal() returned true for different managers")
	}

	if NewManager().Equal(Manager(Any)) == false {
		t.Fatal("Equal() returned false for the any manager")
	}

}

func TestManagerHash(t *testing.T) {

	first := ManagerAPT.Hash()
	second := ManagerAPT.Hash()

	if first == "" {
		t.Fatal("Hash() returned an empty string")
	}

	if first != second {
		t.Fatalf("Hash() is not deterministic: %q != %q", first, second)
	}

	if first == ManagerDNF.Hash() {
		t.Fatal("Hash() returned the same value for different managers")
	}

}

func TestManagerString(t *testing.T) {

	if ManagerAPT.String() != "apt" {
		t.Fatalf("String() = %q, expected %q", ManagerAPT.String(), "apt")
	}

	if NewManager().String() != Any {
		t.Fatalf("String() = %q, expected %q", NewManager().String(), Any)
	}

}

func TestManagerMarshalJSON(t *testing.T) {

	data, err := ManagerAPT.MarshalJSON()

	if err != nil {
		t.Fatalf("MarshalJSON() returned an error: %v", err)
	}

	if string(data) != `"apt"` {
		t.Fatalf("MarshalJSON() = %s, expected %q", data, `"apt"`)
	}

	data, err = json.Marshal(NewManager())

	if err != nil {
		t.Fatalf("json.Marshal() returned an error: %v", err)
	}

	if string(data) != `"any"` {
		t.Fatalf("json.Marshal() = %s, expected %q", data, `"any"`)
	}

}

func TestManagerUnmarshalJSON(t *testing.T) {

	var manager Manager

	if err := json.Unmarshal([]byte(`"npm"`), &manager); err != nil {
		t.Fatalf("UnmarshalJSON() returned an error: %v", err)
	}

	if manager != ManagerNPM {
		t.Fatalf("UnmarshalJSON() = %q, expected %q", manager, ManagerNPM)
	}

	var zero Manager

	if err := json.Unmarshal([]byte(`"unknown"`), &zero); err != nil {
		t.Fatalf("UnmarshalJSON() returned an error for an invalid value: %v", err)
	}

	if zero != Manager("") {
		t.Fatalf("UnmarshalJSON() changed the value for an invalid manager: %q", zero)
	}

	if err := zero.UnmarshalJSON([]byte("not-quoted")); err == nil {
		t.Fatal("UnmarshalJSON() accepted malformed JSON")
	}

}

func TestManagerIsValid(t *testing.T) {

	valid := []Manager{
		Manager(Any),
		ManagerAPT,
		ManagerNPM,
		ManagerPIP,
	}

	for _, manager := range valid {

		value := manager

		if value.IsValid() == false {
			t.Errorf("IsValid() = false for %q", manager)
		}

	}

	invalid := []Manager{"", "unknown", "APT"}

	for _, manager := range invalid {

		value := manager

		if value.IsValid() == true {
			t.Errorf("IsValid() = true for %q", manager)
		}

	}

}
