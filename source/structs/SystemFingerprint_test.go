package structs

import "testing"

func TestSystemFingerprintEqual(t *testing.T) {

	base := SystemFingerprint{
		Country:  "DE",
		Locale:   "de_DE",
		Timezone: "Europe/Berlin",
		Token:    "token",
	}

	if base.Equal(base) == false {
		t.Fatal("a system fingerprint should equal itself")
	}

	if base.Equal(SystemFingerprint{
		Country:  "DE",
		Locale:   "de_DE",
		Timezone: "Europe/Berlin",
		Token:    "token",
	}) == false {
		t.Fatal("equal system fingerprints should be equal")
	}

	cases := []struct {
		name  string
		value SystemFingerprint
	}{
		{"country", SystemFingerprint{Country: "FR", Locale: "de_DE", Timezone: "Europe/Berlin", Token: "token"}},
		{"locale", SystemFingerprint{Country: "DE", Locale: "en_US", Timezone: "Europe/Berlin", Token: "token"}},
		{"timezone", SystemFingerprint{Country: "DE", Locale: "de_DE", Timezone: "UTC", Token: "token"}},
		{"token", SystemFingerprint{Country: "DE", Locale: "de_DE", Timezone: "Europe/Berlin", Token: "other"}},
	}

	for _, c := range cases {

		if base.Equal(c.value) == true {
			t.Errorf("Equal() returned true for a different %s", c.name)
		}

	}

}

func TestSystemFingerprintHash(t *testing.T) {

	empty := SystemFingerprint{}

	if empty.Hash() != "" {
		t.Fatal("Hash() should be empty for an empty fingerprint")
	}

	base := SystemFingerprint{
		Country:  "DE",
		Locale:   "de_DE",
		Timezone: "Europe/Berlin",
		Token:    "token",
	}

	hash := base.Hash()

	if hash == "" {
		t.Fatal("Hash() returned an empty string for a populated fingerprint")
	}

	if base.Hash() != hash {
		t.Fatal("Hash() is not deterministic")
	}

	cases := []struct {
		name  string
		value SystemFingerprint
	}{
		{"country", SystemFingerprint{Country: "FR", Locale: "de_DE", Timezone: "Europe/Berlin", Token: "token"}},
		{"locale", SystemFingerprint{Country: "DE", Locale: "en_US", Timezone: "Europe/Berlin", Token: "token"}},
		{"timezone", SystemFingerprint{Country: "DE", Locale: "de_DE", Timezone: "UTC", Token: "token"}},
		{"token", SystemFingerprint{Country: "DE", Locale: "de_DE", Timezone: "Europe/Berlin", Token: "other"}},
	}

	for _, c := range cases {

		if c.value.Hash() == hash {
			t.Errorf("Hash() did not change for a different %s", c.name)
		}

	}

}
