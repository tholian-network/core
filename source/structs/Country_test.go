package structs

import "testing"

func TestNewCountry(t *testing.T) {

	country := NewCountry("  Germany  ")

	if country.Name != "Germany" {
		t.Errorf("NewCountry() Name = %q, expected %q", country.Name, "Germany")
	}

	if country.Allegiances == nil || len(country.Allegiances) != 0 {
		t.Errorf("NewCountry() Allegiances = %#v, expected an empty slice", country.Allegiances)
	}

	if country.Subnets == nil || len(country.Subnets) != 0 {
		t.Errorf("NewCountry() Subnets = %#v, expected an empty slice", country.Subnets)
	}

	if country.Timezones == nil || len(country.Timezones) != 0 {
		t.Errorf("NewCountry() Timezones = %#v, expected an empty slice", country.Timezones)
	}

	if country.Registry != nil {
		t.Errorf("NewCountry() Registry = %v, expected nil", country.Registry)
	}

}

func TestCountrySetters(t *testing.T) {

	country := NewCountry("Germany")

	country.SetISO("de")
	country.SetContinent("  Europe  ")
	country.SetGeolocation(51.0, 10.0)
	country.SetPopulation(83000000)
	country.SetRegistry("de.registry")

	if country.ISO != "DE" {
		t.Errorf("SetISO() = %q, expected %q", country.ISO, "DE")
	}

	if country.Continent != "Europe" {
		t.Errorf("SetContinent() = %q, expected %q", country.Continent, "Europe")
	}

	if country.Geolocation.Latitude != 51.0 || country.Geolocation.Longitude != 10.0 {
		t.Errorf("SetGeolocation() = %+v, expected 51.0,10.0", country.Geolocation)
	}

	if country.Population != 83000000 {
		t.Errorf("SetPopulation() = %d, expected %d", country.Population, 83000000)
	}

	if country.Registry == nil || *country.Registry != "de.registry" {
		t.Errorf("SetRegistry() = %v, expected %q", country.Registry, "de.registry")
	}

	country.SetISO("USA")
	country.SetISO("1")

	if country.ISO != "DE" {
		t.Errorf("SetISO() accepted an invalid ISO code: %q", country.ISO)
	}

}

func TestCountryIsValid(t *testing.T) {

	valid := func() Country {

		country := NewCountry("Germany")
		country.SetISO("DE")
		country.SetContinent("Europe")
		country.SetGeolocation(51.0, 10.0)

		return country

	}

	cases := []struct {
		name    string
		country Country
		want    bool
	}{
		{"default", Country{}, false},
		{"valid", valid(), true},
		{"missing iso", func() Country { value := valid(); value.ISO = ""; return value }(), false},
		{"missing name", func() Country { value := valid(); value.Name = ""; return value }(), false},
		{"missing continent", func() Country { value := valid(); value.Continent = ""; return value }(), false},
		{"zero geolocation", func() Country { value := valid(); value.SetGeolocation(0, 0); return value }(), false},
	}

	for _, c := range cases {

		if c.country.IsValid() != c.want {
			t.Errorf("%s: IsValid() = %v, expected %v", c.name, c.country.IsValid(), c.want)
		}

	}

}

func TestCountryAllegiances(t *testing.T) {

	country := NewCountry("Germany")

	country.AddAllegiance("EU")
	country.AddAllegiance("EU")
	country.AddAllegiance("NATO")
	country.AddAllegiance("SCO")
	country.AddAllegiance("SCO")

	if len(country.Allegiances) != 3 {
		t.Fatalf("AddAllegiance() did not deduplicate: %#v", country.Allegiances)
	}

	if country.HasAllegiance("EU") == false || country.HasAllegiance("NATO") == false || country.HasAllegiance("SCO") == false {
		t.Fatal("HasAllegiance() did not find an added allegiance")
	}

	if country.HasAllegiance("UN") == true {
		t.Fatal("HasAllegiance() found an allegiance that was not added")
	}

	country.RemoveAllegiance("EU")

	if country.HasAllegiance("EU") == true {
		t.Fatal("RemoveAllegiance() did not remove an allegiance")
	}

	if country.HasAllegiance("SCO") == false {
		t.Fatal("RemoveAllegiance() removed an unrelated allegiance")
	}

	country.RemoveAllegiance("UN")

	country.SetAllegiances([]string{"  EU  ", "", "NATO", "EU", "   ", "UN", "NATO", " SCO ", "SCO"})

	if len(country.Allegiances) != 4 {
		t.Fatalf("SetAllegiances() = %#v, expected 4 entries", country.Allegiances)
	}

	expected := []string{"EU", "NATO", "UN", "SCO"}

	for a := 0; a < len(expected); a++ {

		if country.Allegiances[a] != expected[a] {
			t.Errorf("SetAllegiances()[%d] = %q, expected %q", a, country.Allegiances[a], expected[a])
		}

	}

}

func TestCountrySubnets(t *testing.T) {

	country := NewCountry("Germany")

	subnet := ToSubnet("1.3.3.0/24")

	country.AddSubnet(subnet)
	country.AddSubnet(subnet)

	if len(country.Subnets) != 1 {
		t.Fatalf("AddSubnet() did not deduplicate: %#v", country.Subnets)
	}

	country.AddSubnet(Subnet{})

	if len(country.Subnets) != 1 {
		t.Fatal("AddSubnet() accepted an invalid subnet")
	}

	country.RemoveSubnet(subnet)

	if len(country.Subnets) != 0 {
		t.Fatal("RemoveSubnet() did not remove the subnet")
	}

	country.RemoveSubnet(subnet)

	country.SetSubnets([]Subnet{
		ToSubnet("1.3.3.0/24"),
		Subnet{},
		ToSubnet("invalid"),
		ToSubnet("[fe80::1]/64"),
	})

	if len(country.Subnets) != 2 {
		t.Fatalf("SetSubnets() = %#v, expected 2 valid subnets", country.Subnets)
	}

}

func TestCountryTimezones(t *testing.T) {

	country := NewCountry("Germany")

	berlin := NewTimezone("Europe/Berlin", "+01:00")
	empty := NewTimezone("", "+02:00")
	offset_only := Timezone{Offset: "+03:00"}

	country.SetTimezones([]Timezone{
		berlin,
		empty,
		offset_only,
		berlin,
		NewTimezone("Europe/Berlin", "+01:00"),
		NewTimezone("UTC", "+00:00"),
	})

	if len(country.Timezones) != 2 {
		t.Fatalf("SetTimezones() = %#v, expected 2 entries", country.Timezones)
	}

	if country.HasTimezone("Europe/Berlin") == false {
		t.Fatal("HasTimezone() did not find Europe/Berlin")
	}

	if country.HasTimezone("UTC") == false {
		t.Fatal("HasTimezone() did not find UTC")
	}

	if country.HasTimezone("") == true {
		t.Fatal("HasTimezone() found an empty timezone name")
	}

	if country.HasTimezoneOffset("+01:00") == false {
		t.Fatal("HasTimezoneOffset() did not find +01:00")
	}

	if country.HasTimezoneOffset("+02:00") == true {
		t.Fatal("HasTimezoneOffset() found a filtered timezone offset")
	}

}

func TestCountryEqual(t *testing.T) {

	base := NewCountry("Germany")
	base.SetISO("DE")
	base.SetContinent("Europe")
	base.SetGeolocation(51.0, 10.0)
	base.SetPopulation(83000000)
	base.SetRegistry("de.registry")
	base.AddAllegiance("EU")
	base.AddSubnet(ToSubnet("1.3.3.0/24"))
	base.SetTimezones([]Timezone{NewTimezone("Europe/Berlin", "+01:00")})

	if base.Equal(base) == false {
		t.Fatal("a country should equal itself")
	}

	cases := []struct {
		name  string
		value Country
	}{
		{"iso", func() Country { value := base; value.ISO = "FR"; return value }()},
		{"name", func() Country { value := base; value.Name = "France"; return value }()},
		{"continent", func() Country { value := base; value.Continent = "Asia"; return value }()},
		{"geolocation", func() Country { value := base; value.SetGeolocation(48.0, 2.0); return value }()},
		{"population", func() Country { value := base; value.SetPopulation(1); return value }()},
		{"registry missing", func() Country { value := base; value.Registry = nil; return value }()},
		{"registry value", func() Country { value := base; value.SetRegistry("other"); return value }()},
		{"allegiances", func() Country { value := base; value.AddAllegiance("NATO"); return value }()},
		{"allegiances sco", func() Country { value := base; value.AddAllegiance("SCO"); return value }()},
		{"subnets", func() Country { value := base; value.AddSubnet(ToSubnet("1.3.4.0/24")); return value }()},
		{"timezones", func() Country {
			value := base
			value.SetTimezones([]Timezone{NewTimezone("UTC", "+00:00")})
			return value
		}()},
	}

	for _, c := range cases {

		if base.Equal(c.value) == true {
			t.Errorf("Equal() returned true for a different %s", c.name)
		}

	}

}

func TestCountryHash(t *testing.T) {

	base := NewCountry("Germany")
	base.SetISO("DE")
	base.SetContinent("Europe")
	base.SetGeolocation(51.0, 10.0)
	base.SetPopulation(83000000)
	base.SetRegistry("de.registry")
	base.AddAllegiance("EU")
	base.AddSubnet(ToSubnet("1.3.3.0/24"))
	base.SetTimezones([]Timezone{NewTimezone("Europe/Berlin", "+01:00")})

	hash := base.Hash()

	if hash == "" {
		t.Fatal("Hash() returned an empty string for a populated country")
	}

	if base.Hash() != hash {
		t.Fatal("Hash() is not deterministic")
	}

	empty := Country{}

	if empty.Hash() != "" {
		t.Fatal("Hash() should be empty without a name")
	}

	cases := []struct {
		name  string
		value Country
	}{
		{"iso", func() Country { value := base; value.ISO = "FR"; return value }()},
		{"name", func() Country { value := base; value.Name = "France"; return value }()},
		{"continent", func() Country { value := base; value.Continent = "Asia"; return value }()},
		{"geolocation", func() Country { value := base; value.SetGeolocation(48.0, 2.0); return value }()},
		{"population", func() Country { value := base; value.SetPopulation(1); return value }()},
		{"registry", func() Country { value := base; value.SetRegistry("other"); return value }()},
		{"registry missing", func() Country { value := base; value.Registry = nil; return value }()},
		{"allegiances", func() Country { value := base; value.AddAllegiance("NATO"); return value }()},
		{"allegiances sco", func() Country { value := base; value.AddAllegiance("SCO"); return value }()},
		{"subnets", func() Country { value := base; value.AddSubnet(ToSubnet("1.3.4.0/24")); return value }()},
		{"timezones", func() Country {
			value := base
			value.SetTimezones([]Timezone{NewTimezone("UTC", "+00:00")})
			return value
		}()},
	}

	for _, c := range cases {

		if c.value.Hash() == hash {
			t.Errorf("Hash() did not change for a different %s", c.name)
		}

	}

}
