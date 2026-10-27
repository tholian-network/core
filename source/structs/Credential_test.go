package structs

import "testing"

func TestNewCredential(t *testing.T) {

	credential := NewCredential("  root  ", "user")

	if credential.Name != "root" {
		t.Errorf("NewCredential() Name = %q, expected %q", credential.Name, "root")
	}

	if credential.Type != "user" {
		t.Errorf("NewCredential() Type = %q, expected %q", credential.Type, "user")
	}

	if credential.Keys == nil || len(credential.Keys) != 0 {
		t.Errorf("NewCredential() Keys = %#v, expected an empty slice", credential.Keys)
	}

	if credential.Passwords == nil || len(credential.Passwords) != 0 {
		t.Errorf("NewCredential() Passwords = %#v, expected an empty slice", credential.Passwords)
	}

	invalid := NewCredential("root", "invalid")

	if invalid.Type != "" {
		t.Errorf("NewCredential() accepted an invalid type: %q", invalid.Type)
	}

}

func TestCredentialSetters(t *testing.T) {

	credential := NewCredential("root", "user")

	credential.SetName("  administrator  ")
	credential.SetType("program")

	if credential.Name != "administrator" {
		t.Errorf("SetName() = %q, expected %q", credential.Name, "administrator")
	}

	if credential.Type != "program" {
		t.Errorf("SetType() = %q, expected %q", credential.Type, "program")
	}

	credential.SetType("invalid")

	if credential.Type != "program" {
		t.Errorf("SetType() accepted an invalid type: %q", credential.Type)
	}

	credential.SetType("user")

	if credential.Type != "user" {
		t.Errorf("SetType() = %q, expected %q", credential.Type, "user")
	}

}

func TestCredentialIsValid(t *testing.T) {

	cases := []struct {
		name       string
		credential Credential
		want       bool
	}{
		{"default", Credential{}, false},
		{"name and type only", NewCredential("root", "user"), false},
		{"missing name", func() Credential {
			value := NewCredential("", "user")
			value.AddPassword("secret")
			return value
		}(), false},
		{"missing type", func() Credential {
			value := NewCredential("root", "")
			value.AddKey([]byte{1, 2, 3})
			return value
		}(), false},
		{"with key", func() Credential {
			value := NewCredential("root", "user")
			value.AddKey([]byte{1, 2, 3})
			return value
		}(), true},
		{"with password", func() Credential {
			value := NewCredential("root", "user")
			value.AddPassword("secret")
			return value
		}(), true},
	}

	for _, c := range cases {

		if c.credential.IsValid() != c.want {
			t.Errorf("%s: IsValid() = %v, expected %v", c.name, c.credential.IsValid(), c.want)
		}

	}

}

func TestCredentialKeys(t *testing.T) {

	credential := NewCredential("root", "user")

	credential.AddKey([]byte{1, 2, 3})
	credential.AddKey([]byte{1, 2, 3})
	credential.AddKey([]byte{4, 5, 6})

	if len(credential.Keys) != 2 {
		t.Fatalf("AddKey() did not deduplicate: %#v", credential.Keys)
	}

	credential.RemoveKey([]byte{1, 2, 3})

	if len(credential.Keys) != 1 {
		t.Fatalf("RemoveKey() missed an existing key: %#v", credential.Keys)
	}

	credential.RemoveKey([]byte{7, 8, 9})

	credential.SetKeys([][]byte{{1}, {2}, {1}})

	if len(credential.Keys) != 3 {
		t.Fatalf("SetKeys() = %#v, expected 3 entries", credential.Keys)
	}

}

func TestCredentialPasswords(t *testing.T) {

	credential := NewCredential("root", "user")

	credential.AddPassword("  secret  ")
	credential.AddPassword("secret")
	credential.AddPassword("other")

	if len(credential.Passwords) != 2 {
		t.Fatalf("AddPassword() did not trim or deduplicate: %#v", credential.Passwords)
	}

	if credential.Passwords[0] != "secret" {
		t.Errorf("AddPassword() = %q, expected %q", credential.Passwords[0], "secret")
	}

	credential.RemovePassword(" secret ")

	if len(credential.Passwords) != 1 {
		t.Fatalf("RemovePassword() missed an existing password: %#v", credential.Passwords)
	}

	credential.RemovePassword("missing")

	credential.SetPasswords([]string{"  one  ", "two", "one"})

	if len(credential.Passwords) != 3 {
		t.Fatalf("SetPasswords() = %#v, expected 3 entries", credential.Passwords)
	}

	if credential.Passwords[0] != "one" || credential.Passwords[1] != "two" {
		t.Errorf("SetPasswords() did not trim values: %#v", credential.Passwords)
	}

}

func TestCredentialEqual(t *testing.T) {

	base := NewCredential("root", "user")
	base.AddKey([]byte{1, 2, 3})
	base.AddPassword("secret")

	if base.Equal(base) == false {
		t.Fatal("a credential should equal itself")
	}

	cases := []struct {
		name  string
		value Credential
	}{
		{"name", func() Credential { value := base; value.Name = "other"; return value }()},
		{"type", func() Credential { value := base; value.Type = "program"; return value }()},
		{"keys", func() Credential { value := base; value.AddKey([]byte{4}); return value }()},
		{"key value", func() Credential { value := base; value.SetKeys([][]byte{{9, 9, 9}}); return value }()},
		{"passwords", func() Credential { value := base; value.AddPassword("other"); return value }()},
		{"password value", func() Credential { value := base; value.SetPasswords([]string{"other"}); return value }()},
	}

	for _, c := range cases {

		if base.Equal(c.value) == true {
			t.Errorf("Equal() returned true for a different %s", c.name)
		}

	}

}

func TestCredentialHash(t *testing.T) {

	base := NewCredential("root", "user")
	base.AddKey([]byte{1, 2, 3})
	base.AddPassword("secret")

	hash := base.Hash()

	if hash == "" {
		t.Fatal("Hash() returned an empty string for a populated credential")
	}

	if base.Hash() != hash {
		t.Fatal("Hash() is not deterministic")
	}

	empty := Credential{}

	if empty.Hash() != "" {
		t.Fatal("Hash() should be empty without a name")
	}

	cases := []struct {
		name  string
		value Credential
	}{
		{"name", func() Credential { value := base; value.Name = "other"; return value }()},
		{"type", func() Credential { value := base; value.Type = "program"; return value }()},
		{"keys", func() Credential { value := base; value.AddKey([]byte{4}); return value }()},
		{"passwords", func() Credential { value := base; value.AddPassword("other"); return value }()},
	}

	for _, c := range cases {

		if c.value.Hash() == hash {
			t.Errorf("Hash() did not change for a different %s", c.name)
		}

	}

}
