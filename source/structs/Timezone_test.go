package structs

import "testing"

func TestNewTimezone(t *testing.T) {

	timezone := NewTimezone("  Europe/Berlin  ", "+01:00")

	if timezone.Name != "Europe/Berlin" {
		t.Errorf("NewTimezone() Name = %q, expected %q", timezone.Name, "Europe/Berlin")
	}

	if timezone.Offset != "+01:00" {
		t.Errorf("NewTimezone() Offset = %q, expected %q", timezone.Offset, "+01:00")
	}

	invalid := NewTimezone("UTC", "01:00")

	if invalid.Offset != "" {
		t.Errorf("NewTimezone() accepted an invalid offset: %q", invalid.Offset)
	}

}

func TestTimezoneSetName(t *testing.T) {

	timezone := NewTimezone("Europe/Berlin", "+01:00")

	timezone.SetName("  UTC  ")

	if timezone.Name != "UTC" {
		t.Errorf("SetName() = %q, expected %q", timezone.Name, "UTC")
	}

	timezone.SetName("   ")

	if timezone.Name != "UTC" {
		t.Errorf("SetName() accepted an empty name: %q", timezone.Name)
	}

}

func TestTimezoneSetOffset(t *testing.T) {

	timezone := NewTimezone("Europe/Berlin", "+01:00")

	timezone.SetOffset("-12:30")

	if timezone.Offset != "-12:30" {
		t.Errorf("SetOffset() = %q, expected %q", timezone.Offset, "-12:30")
	}

	cases := []string{
		"01:00",
		"+1:00",
		"+012:0",
		"+01-00",
		"garbage",
		"",
	}

	for _, value := range cases {

		timezone.SetOffset(value)

		if timezone.Offset != "-12:30" {
			t.Errorf("SetOffset(%q) accepted an invalid offset: %q", value, timezone.Offset)
		}

	}

}

func TestTimezoneIsValid(t *testing.T) {

	cases := []struct {
		name     string
		timezone Timezone
		want     bool
	}{
		{"zero value", Timezone{}, false},
		{"name only", Timezone{Name: "Europe/Berlin"}, false},
		{"offset only", Timezone{Offset: "+01:00"}, false},
		{"valid", NewTimezone("Europe/Berlin", "+01:00"), true},
		{"valid negative", NewTimezone("America/New_York", "-05:00"), true},
		{"invalid offset", Timezone{Name: "Europe/Berlin", Offset: "+1:00"}, false},
		{"invalid offset minutes", Timezone{Name: "Europe/Berlin", Offset: "+01:99"}, false},
	}

	for _, c := range cases {

		if c.timezone.IsValid() != c.want {
			t.Errorf("%s: IsValid() = %v, expected %v", c.name, c.timezone.IsValid(), c.want)
		}

	}

}

func TestTimezoneString(t *testing.T) {

	empty := Timezone{}

	if empty.String() != "" {
		t.Fatal("String() should be empty for an invalid timezone")
	}

	timezone := NewTimezone("Europe/Berlin", "+01:00")

	if timezone.String() != "Europe/Berlin (UTC+01:00)" {
		t.Fatalf("String() = %q", timezone.String())
	}

	negative := NewTimezone("America/New_York", "-05:00")

	if negative.String() != "America/New_York (UTC-05:00)" {
		t.Fatalf("String() = %q", negative.String())
	}

}

func TestTimezoneEqual(t *testing.T) {

	base := NewTimezone("Europe/Berlin", "+01:00")

	if base.Equal(base) == false {
		t.Fatal("a timezone should equal itself")
	}

	if base.Equal(NewTimezone("Europe/Berlin", "+01:00")) == false {
		t.Fatal("equal timezones should be equal")
	}

	cases := []struct {
		name  string
		value Timezone
	}{
		{"name", NewTimezone("UTC", "+01:00")},
		{"offset", NewTimezone("Europe/Berlin", "+02:00")},
	}

	for _, c := range cases {

		if base.Equal(c.value) == true {
			t.Errorf("Equal() returned true for a different %s", c.name)
		}

	}

}

func TestTimezoneHash(t *testing.T) {

	base := NewTimezone("Europe/Berlin", "+01:00")

	hash := base.Hash()

	if hash == "" {
		t.Fatal("Hash() returned an empty string for a populated timezone")
	}

	if base.Hash() != hash {
		t.Fatal("Hash() is not deterministic")
	}

	empty := Timezone{}

	if empty.Hash() != "" {
		t.Fatal("Hash() should be empty without a name")
	}

	cases := []struct {
		name  string
		value Timezone
	}{
		{"name", NewTimezone("UTC", "+01:00")},
		{"offset", NewTimezone("Europe/Berlin", "+02:00")},
	}

	for _, c := range cases {

		if c.value.Hash() == hash {
			t.Errorf("Hash() did not change for a different %s", c.name)
		}

	}

}
