package structs

import "tholian-core/types"
import "strings"

type Country struct {
	ISO         string            `json:"iso"`
	Name        string            `json:"name"`
	Continent   string            `json:"continent"`
	Geolocation types.Geolocation `json:"geolocation"`
	Population  uint              `json:"population"`
	Allegiances []string          `json:"allegiances"`
	Subnets     []Subnet          `json:"subnets"`
	Timezones   []Timezone        `json:"timezones"`
	Registry    *string           `json:"registry"`
}

func NewCountry(name string) Country {

	var country Country

	country.SetName(name)
	country.Allegiances = make([]string, 0)
	country.Subnets = make([]Subnet, 0)
	country.Timezones = make([]Timezone, 0)

	return country

}

func (country *Country) IsValid() bool {

	var result bool

	if country.ISO != "" && country.Name != "" {

		result = true

		if country.Continent == "" {
			result = false
		}

		if country.Geolocation.Latitude == 0 && country.Geolocation.Longitude == 0 {
			result = false
		}

	}

	return result

}

func (country *Country) SetISO(value string) {

	if len(value) == 2 {
		country.ISO = strings.ToUpper(strings.TrimSpace(value))
	}

}

func (country *Country) SetName(value string) {
	country.Name = strings.TrimSpace(value)
}

func (country *Country) SetContinent(value string) {
	country.Continent = strings.TrimSpace(value)
}

func (country *Country) SetGeolocation(latitude float64, longitude float64) {

	country.Geolocation.Latitude = latitude
	country.Geolocation.Longitude = longitude

}

func (country *Country) SetPopulation(value uint) {
	country.Population = value
}

func (country *Country) AddAllegiance(value string) {

	found := false

	for a := 0; a < len(country.Allegiances); a++ {

		if country.Allegiances[a] == value {
			found = true
			break
		}

	}

	if found == false {
		country.Allegiances = append(country.Allegiances, value)
	}

}

func (country *Country) HasAllegiance(value string) bool {

	var result bool

	for a := 0; a < len(country.Allegiances); a++ {

		if country.Allegiances[a] == value {
			result = true
			break
		}

	}

	return result

}

func (country *Country) RemoveAllegiance(value string) {

	index := -1

	for a := 0; a < len(country.Allegiances); a++ {

		if country.Allegiances[a] == value {
			index = a
			break
		}

	}

	if index != -1 {
		country.Allegiances = append(country.Allegiances[:index], country.Allegiances[index+1:]...)
	}

}

func (country *Country) SetAllegiances(value []string) {
	country.Allegiances = value
}

func (country *Country) SetRegistry(value string) {
	country.Registry = &value
}

func (country *Country) AddSubnet(value Subnet) {

	if value.IsValid() {

		found := false

		for s := 0; s < len(country.Subnets); s++ {

			if country.Subnets[s].Equal(value) {
				found = true
				break
			}

		}

		if found == false {
			country.Subnets = append(country.Subnets, value)
		}

	}

}

func (country *Country) RemoveSubnet(value Subnet) {

	index := -1

	for s := 0; s < len(country.Subnets); s++ {

		if country.Subnets[s].Equal(value) {
			index = s
			break
		}

	}

	if index != -1 {
		country.Subnets = append(country.Subnets[:index], country.Subnets[index+1:]...)
	}

}

func (country *Country) SetSubnets(values []Subnet) {

	filtered := make([]Subnet, 0)

	for _, value := range values {

		if value.IsValid() {
			filtered = append(filtered, value)
		}

	}

	country.Subnets = filtered

}

func (country *Country) SetTimezones(values []Timezone) {
	country.Timezones = values
}

func (country *Country) HasTimezone(value string) bool {

	var result bool

	for t := 0; t < len(country.Timezones); t++ {

		if country.Timezones[t].Name == value {
			result = true
			break
		}

	}

	return result

}

func (country *Country) HasTimezoneOffset(value string) bool {

	var result bool

	for t := 0; t < len(country.Timezones); t++ {

		if country.Timezones[t].Offset == value {
			result = true
			break
		}

	}

	return result

}
