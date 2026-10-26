package structs

import "sort"
import "strconv"
import "strings"

type Timeslot struct {
	Offset     string   `json:"offset"`
	Population uint     `json:"population"`
	Weight     float64  `json:"weight"`
	Countries  []string `json:"countries"`
	Timezones  []string `json:"timezones"`
}

func NewTimeslot(offset string, weight float64) Timeslot {

	var timeslot Timeslot

	timeslot.Population = 0
	timeslot.Weight = 0.0
	timeslot.Countries = make([]string, 0)
	timeslot.Timezones = make([]string, 0)

	timeslot.SetOffset(offset)
	timeslot.SetWeight(weight)

	return timeslot

}

func (timeslot *Timeslot) Equal(value Timeslot) bool {

	if timeslot.Offset != value.Offset ||
		timeslot.Population != value.Population ||
		timeslot.Weight != value.Weight {
		return false
	}

	if len(timeslot.Countries) != len(value.Countries) {
		return false
	}

	for c := 0; c < len(timeslot.Countries); c++ {

		if timeslot.Countries[c] != value.Countries[c] {
			return false
		}

	}

	if len(timeslot.Timezones) != len(value.Timezones) {
		return false
	}

	for t := 0; t < len(timeslot.Timezones); t++ {

		if timeslot.Timezones[t] != value.Timezones[t] {
			return false
		}

	}

	return true

}

func (timeslot *Timeslot) IsValid() bool {

	if timeslot.Offset != "" &&
		timeslot.Weight != 0.0 {
		return true
	}

	return false

}

func (timeslot *Timeslot) AddCountry(raw string) bool {

	value := strings.ToUpper(strings.TrimSpace(raw))

	if value != "" {

		found := false

		for _, iso := range timeslot.Countries {

			if iso == value {
				found = true
				break
			}

		}

		if found == false {

			timeslot.Countries = append(timeslot.Countries, value)
			sort.Strings(timeslot.Countries)

			return true

		}

	}

	return false

}

func (timeslot *Timeslot) AddTimezone(value string) bool {

	if value != "" && strings.Contains(value, "/") {

		found := false

		for _, timezone := range timeslot.Timezones {

			if timezone == value {
				found = true
				break
			}

		}

		if found == false {

			timeslot.Timezones = append(timeslot.Timezones, value)
			sort.Strings(timeslot.Timezones)

			return true

		}

	}

	return false

}

func (timeslot *Timeslot) RemoveCountry(value string) bool {

	if value != "" {

		index := -1

		for c := 0; c < len(timeslot.Countries); c++ {

			if timeslot.Countries[c] == value {
				index = c
				break
			}

		}

		if index != -1 {

			timeslot.Countries = append(timeslot.Countries[:index], timeslot.Countries[index+1:]...)

			return true

		}

	}

	return false

}

func (timeslot *Timeslot) RemoveTimezone(value string) bool {

	if value != "" {

		index := -1

		for t := 0; t < len(timeslot.Timezones); t++ {

			if timeslot.Timezones[t] == value {
				index = t
				break
			}

		}

		if index != -1 {

			timeslot.Timezones = append(timeslot.Timezones[:index], timeslot.Timezones[index+1:]...)

			return true

		}

	}

	return false

}

func (timeslot *Timeslot) SetCountries(values []string) {
	timeslot.Countries = values
	sort.Strings(timeslot.Countries)
}

func (timeslot *Timeslot) SetOffset(value string) {

	if strings.HasPrefix(value, "+") || strings.HasPrefix(value, "-") {

		if len(value) == 6 && string(value[3]) == ":" {
			timeslot.Offset = value
		}

	}

}

func (timeslot *Timeslot) SetTimezones(values []string) {
	timeslot.Timezones = values
	sort.Strings(timeslot.Timezones)
}

func (timeslot *Timeslot) SetPopulation(value uint) {
	timeslot.Population = value
}

func (timeslot *Timeslot) SetWeight(value float64) {
	timeslot.Weight = value
}

func (timeslot *Timeslot) Hash() string {

	hash := ""

	if timeslot.Offset != "" {

		values := make([]string, 0)

		values = append(values, timeslot.Offset)
		values = append(values, strconv.FormatUint(uint64(timeslot.Population), 10))
		values = append(values, strconv.FormatFloat(timeslot.Weight, 'f', -1, 64))

		for c := 0; c < len(timeslot.Countries); c++ {
			values = append(values, timeslot.Countries[c])
		}

		for t := 0; t < len(timeslot.Timezones); t++ {
			values = append(values, timeslot.Timezones[t])
		}

		hash = toHash(values...)

	}

	return hash

}
