package structs

import "testing"

func TestNewSystemFingerprint(t *testing.T) {

	fingerprint := NewSystemFingerprint()

	if fingerprint.Country != "" ||
		fingerprint.Locale != "" ||
		fingerprint.Timezone != "" ||
		fingerprint.Token != "" {
		t.Fatalf("NewSystemFingerprint() = %+v, expected an empty fingerprint", fingerprint)
	}

	if fingerprint.IsValid() == true {
		t.Fatal("an empty system fingerprint should be invalid")
	}

}

func TestSystemFingerprintIsValid(t *testing.T) {

	cases := []struct {
		name        string
		fingerprint SystemFingerprint
		want        bool
	}{
		{"empty", SystemFingerprint{}, false},
		{"country", SystemFingerprint{Country: "DE"}, false},
		{"country and locale", SystemFingerprint{Country: "DE", Locale: "de_DE"}, false},
		{"country and timezone", SystemFingerprint{Country: "DE", Timezone: "Europe/Berlin"}, false},
		{"locale and timezone", SystemFingerprint{Locale: "de_DE", Timezone: "Europe/Berlin"}, false},
		{"identity", SystemFingerprint{Country: "DE", Locale: "de_DE", Timezone: "Europe/Berlin"}, true},
		{"identity with token", SystemFingerprint{Country: "DE", Locale: "de_DE", Timezone: "Europe/Berlin", Token: "token"}, true},
		{"token only", SystemFingerprint{Token: "token"}, false},
	}

	for _, c := range cases {

		t.Run(c.name, func(t *testing.T) {

			if c.fingerprint.IsValid() != c.want {
				t.Fatalf("IsValid() = %t, expected %t for %+v", c.fingerprint.IsValid(), c.want, c.fingerprint)
			}

		})

	}

}

func TestSystemFingerprintSetters(t *testing.T) {

	fingerprint := NewSystemFingerprint()

	fingerprint.SetCountry("  DE  ")
	fingerprint.SetLocale("  de_DE  ")
	fingerprint.SetTimezone("  Europe/Berlin  ")
	fingerprint.SetToken("  token  ")

	if fingerprint.Country != "DE" {
		t.Fatalf("SetCountry() = %q, expected %q", fingerprint.Country, "DE")
	}

	if fingerprint.Locale != "de_DE" {
		t.Fatalf("SetLocale() = %q, expected %q", fingerprint.Locale, "de_DE")
	}

	if fingerprint.Timezone != "Europe/Berlin" {
		t.Fatalf("SetTimezone() = %q, expected %q", fingerprint.Timezone, "Europe/Berlin")
	}

	if fingerprint.Token != "token" {
		t.Fatalf("SetToken() = %q, expected %q", fingerprint.Token, "token")
	}

	fingerprint.SetCountry("   ")
	fingerprint.SetLocale("   ")
	fingerprint.SetTimezone("   ")
	fingerprint.SetToken("   ")

	if fingerprint.Country != "DE" ||
		fingerprint.Locale != "de_DE" ||
		fingerprint.Timezone != "Europe/Berlin" ||
		fingerprint.Token != "token" {
		t.Fatalf("empty setter values should be ignored, got %+v", fingerprint)
	}

}

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
