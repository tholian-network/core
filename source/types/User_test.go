package types

import "testing"

func TestNewUser(t *testing.T) {

	user := NewUser()

	if user.ID != 0 || user.Name != "" || user.Password != "" || user.Folder != "" || user.Shell != "" || user.Type != "" {
		t.Fatalf("NewUser() = %#v, expected empty fields", user)
	}

	if user.Groups == nil {
		t.Fatal("NewUser().Groups is nil, expected an empty slice")
	}

	if len(user.Groups) != 0 {
		t.Fatalf("NewUser().Groups has %d entries, expected 0", len(user.Groups))
	}

	if user.IsValid() == true {
		t.Fatal("NewUser() returned a valid user")
	}

}

func TestToUser(t *testing.T) {

	user := ToUser("  alice  ", 1000)

	if user.ID != 1000 {
		t.Fatalf("ToUser().ID = %d, expected 1000", user.ID)
	}

	if user.Name != "alice" {
		t.Fatalf("ToUser().Name = %q, expected %q", user.Name, "alice")
	}

	if user.Type != "user" {
		t.Fatalf("ToUser().Type = %q, expected %q", user.Type, "user")
	}

	if user.Groups == nil || len(user.Groups) != 0 {
		t.Fatalf("ToUser().Groups = %#v, expected an empty slice", user.Groups)
	}

	boundary := ToUser("nobody", 65535)

	if boundary.ID != 65535 {
		t.Fatalf("ToUser().ID = %d, expected 65535", boundary.ID)
	}

	trimmed := ToUser("  bob  ", 65534)

	if trimmed.Name != "bob" || trimmed.ID != 65534 {
		t.Fatalf("ToUser() = %#v, expected trimmed name and ID 65534", trimmed)
	}

	empty := ToUser("", 0)

	if empty.IsValid() == true {
		t.Fatal("ToUser(\"\", 0) returned a valid user")
	}

}

func TestUserIsValid(t *testing.T) {

	tests := []struct {
		user     User
		expected bool
	}{
		{User{}, false},
		{User{ID: 1000}, false},
		{User{ID: 0, Name: "root"}, true},
		{User{ID: 1000, Name: "alice"}, true},
		{User{ID: 65534, Name: "nobody"}, true},
		{User{ID: 1, Name: " "}, true},
	}

	for _, test := range tests {

		user := test.user

		if result := user.IsValid(); result != test.expected {
			t.Errorf("User%#v.IsValid() = %t, expected %t", test.user, result, test.expected)
		}

	}

}

func TestUserSetID(t *testing.T) {

	var user User

	user.SetID(0)

	if user.ID != 0 {
		t.Fatalf("SetID(0) = %d, expected 0", user.ID)
	}

	user.SetID(65534)

	if user.ID != 65534 {
		t.Fatalf("SetID(65534) = %d, expected 65534", user.ID)
	}

	user.ID = 42
	user.SetID(65535)

	if user.ID != 42 {
		t.Fatalf("SetID(65535) = %d, expected the previous value 42", user.ID)
	}

}

func TestUserSetName(t *testing.T) {

	var user User

	user.SetName("  alice  ")

	if user.Name != "alice" {
		t.Fatalf("SetName() = %q, expected %q", user.Name, "alice")
	}

	user.SetName("")

	if user.Name != "" {
		t.Fatalf("SetName(\"\") = %q, expected an empty string", user.Name)
	}

}

func TestUserSetPassword(t *testing.T) {

	var user User

	user.SetPassword("secret")

	if user.Password != "secret" {
		t.Fatalf("SetPassword() = %q, expected %q", user.Password, "secret")
	}

	user.SetPassword("")

	if user.Password != "" {
		t.Fatalf("SetPassword(\"\") = %q, expected an empty string", user.Password)
	}

}

func TestUserSetFolder(t *testing.T) {

	var user User

	user.SetFolder("/home/alice")

	if user.Folder != "/home/alice" {
		t.Fatalf("SetFolder() = %q, expected %q", user.Folder, "/home/alice")
	}

	user.SetFolder("home/bob")

	if user.Folder != "/home/alice" {
		t.Fatalf("SetFolder() changed to %q for a relative path", user.Folder)
	}

	user.SetFolder(" /home/bob")

	if user.Folder != "/home/alice" {
		t.Fatalf("SetFolder() changed to %q for a leading space", user.Folder)
	}

	user.SetFolder("")

	if user.Folder != "/home/alice" {
		t.Fatalf("SetFolder(\"\") changed to %q", user.Folder)
	}

}

func TestUserAddGroup(t *testing.T) {

	user := NewUser()
	wheel := Group{ID: 1000, Name: "wheel"}
	staff := Group{ID: 1001, Name: "staff"}

	user.AddGroup(wheel)

	if len(user.Groups) != 1 {
		t.Fatalf("AddGroup() added %d groups, expected 1", len(user.Groups))
	}

	user.AddGroup(wheel)

	if len(user.Groups) != 1 {
		t.Fatalf("AddGroup() duplicated a group, got %d groups", len(user.Groups))
	}

	user.AddGroup(Group{ID: 1000, Name: "other"})

	if len(user.Groups) != 1 {
		t.Fatalf("AddGroup() did not deduplicate by ID, got %d groups", len(user.Groups))
	}

	user.AddGroup(staff)

	if len(user.Groups) != 2 {
		t.Fatalf("AddGroup() added %d groups, expected 2", len(user.Groups))
	}

	user.AddGroup(Group{})

	if len(user.Groups) != 2 {
		t.Fatalf("AddGroup() added an invalid group, got %d groups", len(user.Groups))
	}

	root := NewUser()
	root.AddGroup(Group{ID: 0, Name: "root"})

	if len(root.Groups) != 1 {
		t.Fatalf("AddGroup() rejected the root group, got %d groups", len(root.Groups))
	}

}

func TestUserRemoveGroup(t *testing.T) {

	user := NewUser()
	user.AddGroup(Group{ID: 1000, Name: "wheel"})
	user.AddGroup(Group{ID: 1001, Name: "staff"})

	user.RemoveGroup(Group{ID: 1000, Name: "wheel"})

	if len(user.Groups) != 1 || user.Groups[0].ID != 1001 {
		t.Fatalf("RemoveGroup() = %#v, expected only group 1001", user.Groups)
	}

	user.RemoveGroup(Group{ID: 9999, Name: "missing"})

	if len(user.Groups) != 1 || user.Groups[0].ID != 1001 {
		t.Fatalf("RemoveGroup() changed groups for a missing ID: %#v", user.Groups)
	}

	empty := NewUser()
	empty.RemoveGroup(Group{ID: 1000, Name: "wheel"})

	if len(empty.Groups) != 0 {
		t.Fatalf("RemoveGroup() on an empty user produced %#v", empty.Groups)
	}

}

func TestUserSetGroups(t *testing.T) {

	user := NewUser()

	user.SetGroups([]Group{
		{ID: 1000, Name: "wheel"},
		{},
		{ID: 1001, Name: "staff"},
		{ID: 0, Name: "root"},
	})

	if len(user.Groups) != 3 {
		t.Fatalf("SetGroups() kept %d groups, expected 3", len(user.Groups))
	}

	if user.Groups[0].ID != 1000 || user.Groups[1].ID != 1001 || user.Groups[2].ID != 0 {
		t.Fatalf("SetGroups() = %#v, expected groups 1000, 1001 and 0", user.Groups)
	}

	user.SetGroups([]Group{{}, {}})

	if user.Groups == nil || len(user.Groups) != 0 {
		t.Fatalf("SetGroups() with only invalid groups = %#v, expected an empty slice", user.Groups)
	}

}

func TestUserSetShell(t *testing.T) {

	var user User

	user.SetShell("/bin/bash")

	if user.Shell != "/bin/bash" {
		t.Fatalf("SetShell() = %q, expected %q", user.Shell, "/bin/bash")
	}

	user.SetShell("/bin/zsh  ")

	if user.Shell != "/bin/zsh" {
		t.Fatalf("SetShell() = %q, expected %q", user.Shell, "/bin/zsh")
	}

	user.SetShell("bin/sh")

	if user.Shell != "/bin/zsh" {
		t.Fatalf("SetShell() changed to %q for a relative path", user.Shell)
	}

	user.SetShell(" /bin/fish")

	if user.Shell != "/bin/zsh" {
		t.Fatalf("SetShell() changed to %q for a leading space", user.Shell)
	}

	user.SetShell("")

	if user.Shell != "/bin/zsh" {
		t.Fatalf("SetShell(\"\") changed to %q", user.Shell)
	}

}

func TestUserSetType(t *testing.T) {

	user := NewUser()

	user.SetType("program")

	if user.Type != "program" {
		t.Fatalf("SetType(%q) = %q, expected %q", "program", user.Type, "program")
	}

	user.SetType("invalid")

	if user.Type != "program" {
		t.Fatalf("SetType(%q) = %q, expected the previous value %q", "invalid", user.Type, "program")
	}

	user.SetType("user")

	if user.Type != "user" {
		t.Fatalf("SetType(%q) = %q, expected %q", "user", user.Type, "user")
	}

}

func TestUserHash(t *testing.T) {

	user := User{
		ID:       1000,
		Name:     "alice",
		Password: "secret",
		Folder:   "/home/alice",
		Groups:   []Group{{ID: 1000, Name: "wheel", Type: "user"}},
		Shell:    "/bin/bash",
		Type:     "user",
	}

	first := user.Hash()
	second := user.Hash()

	if first == "" {
		t.Fatal("Hash() returned an empty string")
	}

	if first != second {
		t.Fatalf("Hash() is not deterministic: %q != %q", first, second)
	}

	variants := []User{
		{ID: 1001, Name: "alice", Password: "secret", Folder: "/home/alice", Groups: []Group{{ID: 1000, Name: "wheel", Type: "user"}}, Shell: "/bin/bash", Type: "user"},
		{ID: 1000, Name: "bob", Password: "secret", Folder: "/home/alice", Groups: []Group{{ID: 1000, Name: "wheel", Type: "user"}}, Shell: "/bin/bash", Type: "user"},
		{ID: 1000, Name: "alice", Password: "other", Folder: "/home/alice", Groups: []Group{{ID: 1000, Name: "wheel", Type: "user"}}, Shell: "/bin/bash", Type: "user"},
		{ID: 1000, Name: "alice", Password: "secret", Folder: "/home/bob", Groups: []Group{{ID: 1000, Name: "wheel", Type: "user"}}, Shell: "/bin/bash", Type: "user"},
		{ID: 1000, Name: "alice", Password: "secret", Folder: "/home/alice", Groups: []Group{{ID: 1001, Name: "staff", Type: "user"}}, Shell: "/bin/bash", Type: "user"},
		{ID: 1000, Name: "alice", Password: "secret", Folder: "/home/alice", Groups: []Group{{ID: 1000, Name: "wheel", Type: "user"}}, Shell: "/bin/sh", Type: "user"},
		{ID: 1000, Name: "alice", Password: "secret", Folder: "/home/alice", Groups: []Group{{ID: 1000, Name: "wheel", Type: "user"}}, Shell: "/bin/bash", Type: "program"},
	}

	for _, variant := range variants {

		if first == variant.Hash() {
			t.Errorf("Hash() did not change for %#v", variant)
		}

	}

}
