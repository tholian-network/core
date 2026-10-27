package structs

import "testing"

func TestNewTimeslot(t *testing.T) {

	timeslot := NewTimeslot("+01:00", 0.5)

	if timeslot.Offset != "+01:00" {
		t.Errorf("NewTimeslot() Offset = %q, expected %q", timeslot.Offset, "+01:00")
	}

	if timeslot.Weight != 0.5 {
		t.Errorf("NewTimeslot() Weight = %v, expected %v", timeslot.Weight, 0.5)
	}

	if timeslot.Population != 0 {
		t.Errorf("NewTimeslot() Population = %d, expected 0", timeslot.Population)
	}

	if timeslot.Countries == nil || len(timeslot.Countries) != 0 {
		t.Errorf("NewTimeslot() Countries = %#v, expected an empty slice", timeslot.Countries)
	}

	if timeslot.Timezones == nil || len(timeslot.Timezones) != 0 {
		t.Errorf("NewTimeslot() Timezones = %#v, expected an empty slice", timeslot.Timezones)
	}

	invalid := NewTimeslot("01:00", 0)

	if invalid.Offset != "" {
		t.Errorf("NewTimeslot() accepted an invalid offset: %q", invalid.Offset)
	}

}

func TestTimeslotAddCountry(t *testing.T) {

	timeslot := NewTimeslot("+01:00", 0.5)

	if timeslot.AddCountry("  de  ") == false {
		t.Fatal("AddCountry() rejected a valid country")
	}

	if timeslot.Countries[0] != "DE" {
		t.Errorf("AddCountry() = %q, expected %q", timeslot.Countries[0], "DE")
	}

	if timeslot.AddCountry("at") == false {
		t.Fatal("AddCountry() rejected a valid country")
	}

	if timeslot.AddCountry("fr") == false {
		t.Fatal("AddCountry() rejected a valid country")
	}

	expected := []string{"AT", "DE", "FR"}

	for c := 0; c < len(expected); c++ {

		if timeslot.Countries[c] != expected[c] {
			t.Errorf("AddCountry() countries[%d] = %q, expected %q", c, timeslot.Countries[c], expected[c])
		}

	}

	if timeslot.AddCountry("de") == true {
		t.Fatal("AddCountry() accepted a duplicate country")
	}

	if timeslot.AddCountry("") == true || timeslot.AddCountry("   ") == true {
		t.Fatal("AddCountry() accepted an empty country")
	}

}

func TestTimeslotAddTimezone(t *testing.T) {

	timeslot := NewTimeslot("+01:00", 0.5)

	if timeslot.AddTimezone("Europe/Berlin") == false {
		t.Fatal("AddTimezone() rejected a valid timezone")
	}

	if timeslot.AddTimezone("America/New_York") == false {
		t.Fatal("AddTimezone() rejected a valid timezone")
	}

	expected := []string{"America/New_York", "Europe/Berlin"}

	for tz := 0; tz < len(expected); tz++ {

		if timeslot.Timezones[tz] != expected[tz] {
			t.Errorf("AddTimezone() timezones[%d] = %q, expected %q", tz, timeslot.Timezones[tz], expected[tz])
		}

	}

	if timeslot.AddTimezone("Europe/Berlin") == true {
		t.Fatal("AddTimezone() accepted a duplicate timezone")
	}

	if timeslot.AddTimezone("") == true || timeslot.AddTimezone("UTC") == true {
		t.Fatal("AddTimezone() accepted an invalid timezone")
	}

}

func TestTimeslotRemove(t *testing.T) {

	timeslot := NewTimeslot("+01:00", 0.5)

	timeslot.AddCountry("DE")
	timeslot.AddCountry("AT")
	timeslot.AddTimezone("Europe/Berlin")
	timeslot.AddTimezone("America/New_York")

	if timeslot.RemoveCountry("DE") == false {
		t.Fatal("RemoveCountry() did not remove an existing country")
	}

	if timeslot.RemoveCountry("DE") == true {
		t.Fatal("RemoveCountry() removed a missing country")
	}

	if timeslot.RemoveCountry("") == true {
		t.Fatal("RemoveCountry() accepted an empty country")
	}

	if len(timeslot.Countries) != 1 || timeslot.Countries[0] != "AT" {
		t.Errorf("RemoveCountry() = %#v, expected [AT]", timeslot.Countries)
	}

	if timeslot.RemoveTimezone("Europe/Berlin") == false {
		t.Fatal("RemoveTimezone() did not remove an existing timezone")
	}

	if timeslot.RemoveTimezone("Europe/Berlin") == true {
		t.Fatal("RemoveTimezone() removed a missing timezone")
	}

	if timeslot.RemoveTimezone("") == true {
		t.Fatal("RemoveTimezone() accepted an empty timezone")
	}

	if len(timeslot.Timezones) != 1 || timeslot.Timezones[0] != "America/New_York" {
		t.Errorf("RemoveTimezone() = %#v, expected [America/New_York]", timeslot.Timezones)
	}

}

func TestTimeslotSetters(t *testing.T) {

	timeslot := NewTimeslot("+01:00", 0)

	timeslot.SetCountries([]string{"FR", "DE", "AT"})
	timeslot.SetTimezones([]string{"Europe/Berlin", "America/New_York"})
	timeslot.SetOffset("-05:00")
	timeslot.SetPopulation(1000)
	timeslot.SetWeight(0.25)

	expected_countries := []string{"AT", "DE", "FR"}

	for c := 0; c < len(expected_countries); c++ {

		if timeslot.Countries[c] != expected_countries[c] {
			t.Errorf("SetCountries() countries[%d] = %q, expected %q", c, timeslot.Countries[c], expected_countries[c])
		}

	}

	expected_timezones := []string{"America/New_York", "Europe/Berlin"}

	for tz := 0; tz < len(expected_timezones); tz++ {

		if timeslot.Timezones[tz] != expected_timezones[tz] {
			t.Errorf("SetTimezones() timezones[%d] = %q, expected %q", tz, timeslot.Timezones[tz], expected_timezones[tz])
		}

	}

	if timeslot.Offset != "-05:00" {
		t.Errorf("SetOffset() = %q, expected %q", timeslot.Offset, "-05:00")
	}

	if timeslot.Population != 1000 {
		t.Errorf("SetPopulation() = %d, expected %d", timeslot.Population, 1000)
	}

	if timeslot.Weight != 0.25 {
		t.Errorf("SetWeight() = %v, expected %v", timeslot.Weight, 0.25)
	}

	timeslot.SetOffset("invalid")

	if timeslot.Offset != "-05:00" {
		t.Errorf("SetOffset() accepted an invalid offset: %q", timeslot.Offset)
	}

}

func TestTimeslotIsValid(t *testing.T) {

	cases := []struct {
		name     string
		timeslot Timeslot
		want     bool
	}{
		{"default", Timeslot{}, false},
		{"missing weight", NewTimeslot("+01:00", 0), false},
		{"missing offset", NewTimeslot("", 0.5), false},
		{"valid", NewTimeslot("+01:00", 0.5), true},
	}

	for _, c := range cases {

		if c.timeslot.IsValid() != c.want {
			t.Errorf("%s: IsValid() = %v, expected %v", c.name, c.timeslot.IsValid(), c.want)
		}

	}

}

func TestTimeslotEqual(t *testing.T) {

	base := NewTimeslot("+01:00", 0.5)
	base.SetPopulation(1000)
	base.AddCountry("DE")
	base.AddTimezone("Europe/Berlin")

	if base.Equal(base) == false {
		t.Fatal("a timeslot should equal itself")
	}

	cases := []struct {
		name  string
		value Timeslot
	}{
		{"offset", func() Timeslot { value := base; value.SetOffset("+02:00"); return value }()},
		{"population", func() Timeslot { value := base; value.SetPopulation(1); return value }()},
		{"weight", func() Timeslot { value := base; value.SetWeight(1.0); return value }()},
		{"countries", func() Timeslot { value := base; value.AddCountry("AT"); return value }()},
		{"timezones", func() Timeslot { value := base; value.AddTimezone("Asia/Tokyo"); return value }()},
	}

	for _, c := range cases {

		if base.Equal(c.value) == true {
			t.Errorf("Equal() returned true for a different %s", c.name)
		}

	}

}

func TestTimeslotHash(t *testing.T) {

	base := NewTimeslot("+01:00", 0.5)
	base.SetPopulation(1000)
	base.AddCountry("DE")
	base.AddTimezone("Europe/Berlin")

	hash := base.Hash()

	if hash == "" {
		t.Fatal("Hash() returned an empty string for a populated timeslot")
	}

	if base.Hash() != hash {
		t.Fatal("Hash() is not deterministic")
	}

	empty := Timeslot{}

	if empty.Hash() != "" {
		t.Fatal("Hash() should be empty without an offset")
	}

	cases := []struct {
		name  string
		value Timeslot
	}{
		{"offset", func() Timeslot { value := base; value.SetOffset("+02:00"); return value }()},
		{"population", func() Timeslot { value := base; value.SetPopulation(1); return value }()},
		{"weight", func() Timeslot { value := base; value.SetWeight(1.0); return value }()},
		{"countries", func() Timeslot { value := base; value.AddCountry("AT"); return value }()},
		{"timezones", func() Timeslot { value := base; value.AddTimezone("Asia/Tokyo"); return value }()},
	}

	for _, c := range cases {

		if c.value.Hash() == hash {
			t.Errorf("Hash() did not change for a different %s", c.name)
		}

	}

}
