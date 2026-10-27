package types

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

func TestToGroup(t *testing.T) {

	group := ToGroup("  wheel  ")

	if group.Name != "wheel" {
		t.Fatalf("ToGroup().Name = %q, expected %q", group.Name, "wheel")
	}

	if group.Type != "user" {
		t.Fatalf("ToGroup().Type = %q, expected %q", group.Type, "user")
	}

	if group.ID != 0 {
		t.Fatalf("ToGroup().ID = %d, expected 0", group.ID)
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
