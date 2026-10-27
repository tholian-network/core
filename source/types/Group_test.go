package types

import "encoding/json"
import "testing"

func TestNewGroup(t *testing.T) {

	group := NewGroup()

	if group.ID != 0 || group.Name != "" || group.Password != "" {
		t.Fatalf("NewGroup() = %#v, expected empty fields", group)
	}

	if group.Type != "user" {
		t.Fatalf("NewGroup().Type = %q, expected %q", group.Type, "user")
	}

	if group.IsValid() == true {
		t.Fatal("NewGroup() returned a valid group")
	}

}

func TestToGroupMinimal(t *testing.T) {

	group1 := ToGroup("wheel")
	group2 := ToGroup("wheel:user")
	group3 := ToGroup("wheel:user:1000")

	if group1.ID != 0 || group1.Name != "wheel" || group1.Type != "" || group1.Password != "" {
		t.Fatalf("ToGroup(%q) = %#v, expected only the name", "wheel", group1)
	}

	if group2.ID != 0 || group2.Name != "wheel" || group2.Type != "user" || group2.Password != "" {
		t.Fatalf("ToGroup(%q) = %#v, expected the name and type", "wheel:user", group2)
	}

	if group3.ID != 1000 || group3.Name != "wheel" || group3.Type != "user" || group3.Password != "" {
		t.Fatalf("ToGroup(%q) = %#v, expected the name, type and ID", "wheel:user:1000", group3)
	}

	if group3.IsValid() == false {
		t.Fatal("ToGroup() returned an invalid group")
	}

}

func TestToGroup(t *testing.T) {

	group := ToGroup("wheel:user:1000:secret")

	if group.Name != "wheel" {
		t.Fatalf("ToGroup().Name = %q, expected %q", group.Name, "wheel")
	}

	if group.Type != "user" {
		t.Fatalf("ToGroup().Type = %q, expected %q", group.Type, "user")
	}

	if group.ID != 1000 {
		t.Fatalf("ToGroup().ID = %d, expected 1000", group.ID)
	}

	if group.Password != "secret" {
		t.Fatalf("ToGroup().Password = %q, expected %q", group.Password, "secret")
	}

	if group.IsValid() == false {
		t.Fatal("ToGroup() returned an invalid group")
	}

	empty := ToGroup("")

	if empty.IsValid() == true {
		t.Fatal("ToGroup(\"\") returned a valid group")
	}

}

func TestGroupIsValid(t *testing.T) {

	tests := []struct {
		group    Group
		expected bool
	}{
		{Group{}, false},
		{Group{ID: 1000}, false},
		{Group{ID: 0, Name: "root"}, true},
		{Group{ID: 1000, Name: "wheel"}, true},
		{Group{ID: 65534, Name: "nobody"}, true},
		{Group{ID: 1, Name: " "}, true},
	}

	for _, test := range tests {

		group := test.group

		if result := group.IsValid(); result != test.expected {
			t.Errorf("Group%#v.IsValid() = %t, expected %t", test.group, result, test.expected)
		}

	}

}

func TestGroupSetID(t *testing.T) {

	var group Group

	group.SetID(0)

	if group.ID != 0 {
		t.Fatalf("SetID(0) = %d, expected 0", group.ID)
	}

	group.SetID(65534)

	if group.ID != 65534 {
		t.Fatalf("SetID(65534) = %d, expected 65534", group.ID)
	}

	group.ID = 42
	group.SetID(65535)

	if group.ID != 42 {
		t.Fatalf("SetID(65535) = %d, expected the previous value 42", group.ID)
	}

}

func TestGroupSetName(t *testing.T) {

	var group Group

	group.SetName("  wheel  ")

	if group.Name != "wheel" {
		t.Fatalf("SetName() = %q, expected %q", group.Name, "wheel")
	}

	group.SetName("")

	if group.Name != "" {
		t.Fatalf("SetName(\"\") = %q, expected an empty string", group.Name)
	}

}

func TestGroupSetPassword(t *testing.T) {

	var group Group

	group.SetPassword("secret")

	if group.Password != "secret" {
		t.Fatalf("SetPassword() = %q, expected %q", group.Password, "secret")
	}

	group.SetPassword("")

	if group.Password != "" {
		t.Fatalf("SetPassword(\"\") = %q, expected an empty string", group.Password)
	}

}

func TestGroupSetType(t *testing.T) {

	group := NewGroup()

	group.SetType("program")

	if group.Type != "program" {
		t.Fatalf("SetType(%q) = %q, expected %q", "program", group.Type, "program")
	}

	group.SetType("invalid")

	if group.Type != "program" {
		t.Fatalf("SetType(%q) = %q, expected the previous value %q", "invalid", group.Type, "program")
	}

	group.SetType("user")

	if group.Type != "user" {
		t.Fatalf("SetType(%q) = %q, expected %q", "user", group.Type, "user")
	}

}

func TestGroupHash(t *testing.T) {

	group := Group{ID: 1000, Name: "wheel", Password: "secret", Type: "user"}

	first := group.Hash()
	second := group.Hash()

	if first == "" {
		t.Fatal("Hash() returned an empty string")
	}

	if first != second {
		t.Fatalf("Hash() is not deterministic: %q != %q", first, second)
	}

	variants := []Group{
		{ID: 1001, Name: "wheel", Password: "secret", Type: "user"},
		{ID: 1000, Name: "other", Password: "secret", Type: "user"},
		{ID: 1000, Name: "wheel", Password: "other", Type: "user"},
		{ID: 1000, Name: "wheel", Password: "secret", Type: "program"},
	}

	for _, variant := range variants {

		if first == variant.Hash() {
			t.Errorf("Hash() did not change for %#v", variant)
		}

	}

}

func TestGroupString(t *testing.T) {

	tests := []struct {
		group    Group
		expected string
	}{
		{Group{}, "::0:"},
		{Group{ID: 1000, Name: "wheel", Password: "secret", Type: "user"}, "wheel:user:1000:secret"},
		{Group{ID: 1000, Name: "wheel", Password: "sec:ret,with space%", Type: "user"}, "wheel:user:1000:sec%3Aret%2Cwith+space%25"},
	}

	for _, test := range tests {

		if result := test.group.String(); result != test.expected {
			t.Errorf("Group%#v.String() = %q, expected %q", test.group, result, test.expected)
		}

	}

}

func TestGroupMarshalJSON(t *testing.T) {

	group := Group{ID: 1000, Name: "wheel", Password: "secret", Type: "user"}

	data, err := json.Marshal(group)

	if err != nil {
		t.Fatalf("Marshal() returned an error: %v", err)
	}

	if string(data) != `"wheel:user:1000:secret"` {
		t.Fatalf("Marshal() = %s, expected %q", data, `"wheel:user:1000:secret"`)
	}

}

func TestGroupUnmarshalJSON(t *testing.T) {

	original := Group{ID: 1000, Name: "wheel", Password: "sec:ret,with space%", Type: "user"}

	data, err := json.Marshal(original)

	if err != nil {
		t.Fatalf("Marshal() returned an error: %v", err)
	}

	var decoded Group

	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal() returned an error: %v", err)
	}

	if decoded != original {
		t.Fatalf("Unmarshal() = %#v, expected %#v", decoded, original)
	}

}

func TestGroupUnmarshalJSONSanitizes(t *testing.T) {

	var group Group

	if err := json.Unmarshal([]byte(`" wheel :invalid:65535:secret"`), &group); err != nil {
		t.Fatalf("Unmarshal() returned an error: %v", err)
	}

	expected := Group{ID: 65535, Name: "wheel", Password: "secret", Type: ""}

	if group != expected {
		t.Fatalf("Unmarshal() = %#v, expected %#v", group, expected)
	}

}

func TestGroupUnmarshalJSONErrors(t *testing.T) {

	original := Group{ID: 1000, Name: "wheel", Password: "secret", Type: "user"}

	inputs := []string{
		`not-quoted`,
		`123`,
	}

	for _, input := range inputs {

		group := original

		if err := json.Unmarshal([]byte(input), &group); err == nil {
			t.Errorf("Unmarshal(%s) accepted malformed JSON", input)
		}

		if group != original {
			t.Errorf("Unmarshal(%s) changed the value to %#v", input, group)
		}

	}

}

func TestGroupUnmarshalJSONIgnoresInvalid(t *testing.T) {

	original := Group{ID: 1000, Name: "wheel", Password: "secret", Type: "user"}

	inputs := []string{
		`"wheel:user"`,
		`"wheel:user:not-a-number:secret"`,
		`"wheel:user:70000:secret"`,
	}

	for _, input := range inputs {

		group := original

		if err := json.Unmarshal([]byte(input), &group); err != nil {
			t.Errorf("Unmarshal(%s) returned an error: %v", input, err)
		}

		if group != original {
			t.Errorf("Unmarshal(%s) changed the value to %#v", input, group)
		}

	}

}

func TestGroupUnmarshalJSONPartial(t *testing.T) {

	tests := []struct {
		input    string
		expected Group
	}{
		{
			`"wheel:user:1000"`,
			Group{ID: 1000, Name: "wheel", Type: "user"},
		},
		{
			`"wheel:user:1000:%zz"`,
			Group{ID: 1000, Name: "wheel", Type: "user"},
		},
	}

	for _, test := range tests {

		var group Group

		if err := json.Unmarshal([]byte(test.input), &group); err != nil {
			t.Fatalf("Unmarshal(%s) returned an error: %v", test.input, err)
		}

		if group != test.expected {
			t.Fatalf("Unmarshal(%s) = %#v, expected %#v", test.input, group, test.expected)
		}

	}

}
