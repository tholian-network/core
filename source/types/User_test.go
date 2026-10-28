package types

import "encoding/json"
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

func TestToUserMinimal(t *testing.T) {

	user1 := ToUser("alice")
	user2 := ToUser("alice:user")
	user3 := ToUser("alice:user:1001")

	if user1.ID != 0 || user1.Name != "alice" || user1.Type != "" {
		t.Fatalf("ToUser(%q) = %#v, expected only the name", "alice", user1)
	}

	if user2.ID != 0 || user2.Name != "alice" || user2.Type != "user" {
		t.Fatalf("ToUser(%q) = %#v, expected the name and type", "alice:user", user2)
	}

	if user3.ID != 1001 || user3.Name != "alice" || user3.Type != "user" {
		t.Fatalf("ToUser(%q) = %#v, expected the name, type and ID", "alice:user:1001", user3)
	}

	if user3.IsValid() == false {
		t.Fatal("ToUser() returned an invalid user")
	}

	for _, user := range []User{user1, user2, user3} {

		if user.Groups == nil || len(user.Groups) != 0 {
			t.Errorf("ToUser() groups = %#v, expected an empty slice", user.Groups)
		}

	}

}

func TestToUser(t *testing.T) {

	user := ToUser("alice:user:1000:%2Fhome%2Falice:%2Fbin%2Fbash:secret:wheel:user:1000:groupsecret")

	if user.ID != 1000 {
		t.Fatalf("ToUser().ID = %d, expected 1000", user.ID)
	}

	if user.Name != "alice" {
		t.Fatalf("ToUser().Name = %q, expected %q", user.Name, "alice")
	}

	if user.Password != "secret" {
		t.Fatalf("ToUser().Password = %q, expected %q", user.Password, "secret")
	}

	if user.Folder != "/home/alice" {
		t.Fatalf("ToUser().Folder = %q, expected %q", user.Folder, "/home/alice")
	}

	if user.Shell != "/bin/bash" {
		t.Fatalf("ToUser().Shell = %q, expected %q", user.Shell, "/bin/bash")
	}

	if user.Type != "user" {
		t.Fatalf("ToUser().Type = %q, expected %q", user.Type, "user")
	}

	if len(user.Groups) != 1 || user.Groups[0].ID != 1000 || user.Groups[0].Name != "wheel" {
		t.Fatalf("ToUser().Groups = %#v, expected the wheel group", user.Groups)
	}

	if user.IsValid() == false {
		t.Fatal("ToUser() returned an invalid user")
	}

	invalid := []string{
		"",
		"  alice  ",
	}

	for _, value := range invalid {

		user := ToUser(value)

		if user.IsValid() == true {
			t.Errorf("ToUser(%q) returned a valid user", value)
		}

		if user.Groups == nil || len(user.Groups) != 0 {
			t.Errorf("ToUser(%q).Groups = %#v, expected an empty slice", value, user.Groups)
		}

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

	if user.ID != 65535 {
		t.Fatalf("SetID(65535) = %d, expected 65535", user.ID)
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

func TestUserParse(t *testing.T) {

	tests := []struct {
		input    string
		expected User
	}{
		{
			"alice:user:1000:%2Fhome%2Falice:%2Fbin%2Fbash:secret:wheel:user:1000:groupsecret",
			User{ID: 1000, Name: "alice", Password: "secret", Folder: "/home/alice", Groups: []Group{{ID: 1000, Name: "wheel", Password: "groupsecret", Type: "user"}}, Shell: "/bin/bash", Type: "user"},
		},
		{
			" alice :invalid:65535:home/alice:bin/bash:secret:wheel:user:1000:groupsecret",
			User{ID: 65535, Name: "alice", Password: "secret", Groups: []Group{{ID: 1000, Name: "wheel", Password: "groupsecret", Type: "user"}}},
		},
		{
			"alice:user:1000:%2Fhome%2Falice:%2Fbin%2Fbash:secret:",
			User{ID: 1000, Name: "alice", Password: "secret", Folder: "/home/alice", Shell: "/bin/bash", Type: "user"},
		},
		{
			"alice",
			User{Name: "alice"},
		},
		{
			"alice:user",
			User{Name: "alice", Type: "user"},
		},
		{
			"alice:user:1000",
			User{ID: 1000, Name: "alice", Type: "user"},
		},
		{
			"alice:user:1000:%2Fhome%2Falice:%2Fbin%2Fbash:secret:wheel:user",
			User{ID: 1000, Name: "alice", Password: "secret", Folder: "/home/alice", Shell: "/bin/bash", Type: "user"},
		},
	}

	for _, test := range tests {

		var user User

		if user.Parse(test.input) == false {
			t.Fatalf("Parse(%q) returned false", test.input)
		}

		if user.Hash() != test.expected.Hash() {
			t.Fatalf("Parse(%q) = %#v, expected %#v", test.input, user, test.expected)
		}

		if test.expected.Groups != nil && user.Groups == nil {
			t.Fatalf("Parse(%q) resulted in nil Groups", test.input)
		}

	}

	invalid := []string{
		"alice:user:not-a-number:%2Fhome%2Falice:%2Fbin%2Fbash:secret:",
		"alice:user:70000:%2Fhome%2Falice:%2Fbin%2Fbash:secret:",
		"alice:user:1000:%zz:%2Fbin%2Fbash:secret:",
	}

	for _, input := range invalid {

		user := User{ID: 1000, Name: "alice", Password: "secret", Folder: "/home/alice", Shell: "/bin/bash", Type: "user"}

		original := user

		if user.Parse(input) == true {
			t.Errorf("Parse(%q) returned true", input)
		}

		if user.Hash() != original.Hash() {
			t.Errorf("Parse(%q) changed the value to %#v", input, user)
		}

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

func TestUserString(t *testing.T) {

	tests := []struct {
		user     User
		expected string
	}{
		{User{}, "::0::::"},
		{
			User{ID: 1000, Name: "alice", Password: "secret", Folder: "/home/alice", Shell: "/bin/bash", Type: "user"},
			"alice:user:1000:%2Fhome%2Falice:%2Fbin%2Fbash:secret:",
		},
		{
			User{
				ID:       1000,
				Name:     "alice",
				Password: "se:cret",
				Folder:   "/home/alice",
				Groups:   []Group{{ID: 1000, Name: "wheel", Password: "group:secret", Type: "user"}},
				Shell:    "/bin/bash",
				Type:     "user",
			},
			"alice:user:1000:%2Fhome%2Falice:%2Fbin%2Fbash:se%3Acret:wheel:user:1000:group%3Asecret",
		},
	}

	for _, test := range tests {

		if result := test.user.String(); result != test.expected {
			t.Errorf("User%#v.String() = %q, expected %q", test.user, result, test.expected)
		}

	}

}

func TestUserMarshalJSON(t *testing.T) {

	user := User{
		ID:       1000,
		Name:     "alice",
		Password: "secret",
		Folder:   "/home/alice",
		Groups:   []Group{{ID: 1000, Name: "wheel", Password: "secret", Type: "user"}},
		Shell:    "/bin/bash",
		Type:     "user",
	}

	data, err := json.Marshal(user)

	if err != nil {
		t.Fatalf("Marshal() returned an error: %v", err)
	}

	expected := `"alice:user:1000:%2Fhome%2Falice:%2Fbin%2Fbash:secret:wheel:user:1000:secret"`

	if string(data) != expected {
		t.Fatalf("Marshal() = %s, expected %s", data, expected)
	}

}

func TestUserUnmarshalJSON(t *testing.T) {

	original := User{
		ID:       1000,
		Name:     "alice",
		Password: "se:cret",
		Folder:   "/home/alice",
		Groups: []Group{
			{ID: 1000, Name: "wheel", Password: "group:secret", Type: "user"},
			{ID: 1001, Name: "staff", Password: "other,secret", Type: "program"},
		},
		Shell: "/bin/bash",
		Type:  "user",
	}

	data, err := json.Marshal(original)

	if err != nil {
		t.Fatalf("Marshal() returned an error: %v", err)
	}

	var decoded User

	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal() returned an error: %v", err)
	}

	if decoded.ID != original.ID ||
		decoded.Name != original.Name ||
		decoded.Password != original.Password ||
		decoded.Folder != original.Folder ||
		decoded.Shell != original.Shell ||
		decoded.Type != original.Type {
		t.Fatalf("Unmarshal() = %#v, expected %#v", decoded, original)
	}

	if len(decoded.Groups) != len(original.Groups) {
		t.Fatalf("Unmarshal() groups = %#v, expected %#v", decoded.Groups, original.Groups)
	}

	for g := 0; g < len(original.Groups); g++ {

		if decoded.Groups[g] != original.Groups[g] {
			t.Fatalf("Unmarshal() group %d = %#v, expected %#v", g, decoded.Groups[g], original.Groups[g])
		}

	}

}

func TestUserUnmarshalJSONSanitizes(t *testing.T) {

	var user User

	input := `" alice :invalid:65535:home/alice:bin/bash:secret:wheel:user:1000:groupsecret"`

	if err := json.Unmarshal([]byte(input), &user); err != nil {
		t.Fatalf("Unmarshal() returned an error: %v", err)
	}

	expected := User{
		ID:       65535,
		Name:     "alice",
		Password: "secret",
		Folder:   "",
		Groups:   []Group{{ID: 1000, Name: "wheel", Password: "groupsecret", Type: "user"}},
		Shell:    "",
		Type:     "",
	}

	if user.ID != expected.ID ||
		user.Name != expected.Name ||
		user.Password != expected.Password ||
		user.Folder != expected.Folder ||
		user.Shell != expected.Shell ||
		user.Type != expected.Type {
		t.Fatalf("Unmarshal() = %#v, expected %#v", user, expected)
	}

	if len(user.Groups) != len(expected.Groups) || user.Groups[0] != expected.Groups[0] {
		t.Fatalf("Unmarshal() groups = %#v, expected %#v", user.Groups, expected.Groups)
	}

}

func TestUserUnmarshalJSONDropsInvalidGroups(t *testing.T) {

	var user User

	input := `"alice:user:1000:%2Fhome%2Falice:%2Fbin%2Fbash:secret:::0:"`

	if err := json.Unmarshal([]byte(input), &user); err != nil {
		t.Fatalf("Unmarshal() returned an error: %v", err)
	}

	if user.Groups == nil || len(user.Groups) != 0 {
		t.Fatalf("Unmarshal() groups = %#v, expected an empty slice", user.Groups)
	}

}

func TestUserUnmarshalJSONErrors(t *testing.T) {

	original := User{ID: 1000, Name: "alice", Password: "secret", Folder: "/home/alice", Shell: "/bin/bash", Type: "user"}

	inputs := []string{
		`not-quoted`,
		`123`,
	}

	for _, input := range inputs {

		user := original

		if err := json.Unmarshal([]byte(input), &user); err == nil {
			t.Errorf("Unmarshal(%s) accepted malformed JSON", input)
		}

		if user.Hash() != original.Hash() {
			t.Errorf("Unmarshal(%s) changed the value to %#v", input, user)
		}

	}

}

func TestUserUnmarshalJSONIgnoresInvalid(t *testing.T) {

	original := User{ID: 1000, Name: "alice", Password: "secret", Folder: "/home/alice", Shell: "/bin/bash", Type: "user"}

	inputs := []string{
		`"alice:user:not-a-number:%2Fhome%2Falice:%2Fbin%2Fbash:secret:"`,
		`"alice:user:70000:%2Fhome%2Falice:%2Fbin%2Fbash:secret:"`,
	}

	for _, input := range inputs {

		user := original

		if err := json.Unmarshal([]byte(input), &user); err != nil {
			t.Errorf("Unmarshal(%s) returned an error: %v", input, err)
		}

		if user.Hash() != original.Hash() {
			t.Errorf("Unmarshal(%s) changed the value to %#v", input, user)
		}

	}

}

func TestUserUnmarshalJSONPartial(t *testing.T) {

	tests := []struct {
		input    string
		expected User
	}{
		{
			`"alice:user:1000"`,
			User{ID: 1000, Name: "alice", Groups: []Group{}, Type: "user"},
		},
		{
			`"alice:user:1000:%zz:%2Fbin%2Fbash:secret:"`,
			User{ID: 1000, Name: "alice", Groups: []Group{}, Type: "user"},
		},
		{
			`"alice:user:1000:%2Fhome%2Falice:%2Fbin%2Fbash:secret:wheel:user"`,
			User{ID: 1000, Name: "alice", Password: "secret", Folder: "/home/alice", Groups: []Group{}, Shell: "/bin/bash", Type: "user"},
		},
	}

	for _, test := range tests {

		var user User

		if err := json.Unmarshal([]byte(test.input), &user); err != nil {
			t.Fatalf("Unmarshal(%s) returned an error: %v", test.input, err)
		}

		if user.ID != test.expected.ID ||
			user.Name != test.expected.Name ||
			user.Password != test.expected.Password ||
			user.Folder != test.expected.Folder ||
			user.Shell != test.expected.Shell ||
			user.Type != test.expected.Type {
			t.Fatalf("Unmarshal(%s) = %#v, expected %#v", test.input, user, test.expected)
		}

		if user.Groups == nil || len(user.Groups) != 0 {
			t.Fatalf("Unmarshal(%s) groups = %#v, expected an empty slice", test.input, user.Groups)
		}

	}

}
