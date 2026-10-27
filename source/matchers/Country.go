package matchers

import "strings"

type Country struct {
	Name       string `json:"name"`
	Continent  string `json:"continent"`
	Allegiance string `json:"allegiance"`
	Subnet     string `json:"subnet"`
	Timezone   string `json:"timezone"`
}

func NewCountry() Country {

	var country Country

	country.Name = Any
	country.Continent = Any
	country.Allegiance = Any
	country.Subnet = Any
	country.Timezone = Any

	return country

}

func ToCountry(value string) Country {

	var country Country

	country.Name = Any
	country.Continent = Any
	country.Allegiance = Any
	country.Subnet = Any
	country.Timezone = Any

	country.SetName(value)

	return country

}

func (country *Country) Equal(value Country) bool {

	if country.Name == value.Name &&
		country.Continent == value.Continent &&
		country.Allegiance == value.Allegiance &&
		country.Subnet == value.Subnet &&
		country.Timezone == value.Timezone {
		return true
	}

	return false

}

func (country *Country) IsValid() bool {

	if country.Name != Any ||
		country.Continent != Any ||
		country.Allegiance != Any ||
		country.Subnet != Any ||
		country.Timezone != Any {
		return true
	}

	return false

}

func (country *Country) Matches(name string, continent string, allegiance string, subnet string, timezone string) bool {
	return country.MatchesName(name) && country.MatchesContinent(continent) && country.MatchesAllegiance(allegiance) && country.MatchesSubnet(subnet) && country.MatchesTimezone(timezone)
}

func (country *Country) MatchesAllegiance(value string) bool {

	if country.Allegiance == value {
		return true
	} else if country.Allegiance == Any {
		return true
	}

	return false

}

func (country *Country) MatchesContinent(value string) bool {

	if country.Continent == value {
		return true
	} else if country.Continent == Any {
		return true
	}

	return false

}

func (country *Country) MatchesName(value string) bool {

	if country.Name == value {
		return true
	} else if country.Name == Any {
		return true
	}

	return false

}

func (country *Country) MatchesSubnet(value string) bool {

	if country.Subnet != Any && value != "any" {
		return containsSubnet(value, country.Subnet)
	} else if country.Subnet == Any {
		return true
	}

	return false

}

func (country *Country) MatchesTimezone(value string) bool {

	if country.Timezone == value {
		return true
	} else if country.Timezone == Any {
		return true
	}

	return false

}

func (country *Country) SetAllegiance(value string) {

	if value == "all" || value == "any" || value == "*" {
		country.Allegiance = Any
	} else if value != "" {
		country.Allegiance = value
	}

}

func (country *Country) SetName(value string) {
	country.Name = strings.TrimSpace(value)
}

func (country *Country) SetContinent(value string) {

	if value == "all" || value == "any" || value == "*" {
		country.Continent = Any
	} else if value != "" {
		country.Continent = value
	}

}

func (country *Country) SetSubnet(value string) {

	address, prefix := toSubnet(value)

	if value == "all" || value == "any" || value == "*" {
		country.Subnet = Any
	} else if address != "" && prefix != 0 {
		country.Subnet = value
	}

}

func (country *Country) SetTimezone(value string) {

	if value == "all" || value == "any" || value == "*" {
		country.Timezone = Any
	} else if value != "" {
		country.Timezone = value
	}

}

func (country *Country) Hash() string {

	hash := ""

	if country.Name != "" {

		hash = toHash(
			country.Name,
			country.Continent,
			country.Allegiance,
			country.Subnet,
			country.Timezone,
		)

	}

	return hash

}
