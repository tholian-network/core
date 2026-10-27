package matchers

import "tholian-core/types"
import "strings"

type System struct {
	Name     string `json:"name"`
	Hostname string `json:"hostname"`
	Datetime string `json:"datetime"`
	Country  string `json:"country"`
	Timezone string `json:"timezone"`
}

func NewSystem() System {

	var system System

	system.Name = Any
	system.Hostname = Any
	system.Datetime = Any
	system.Country = Any
	system.Timezone = Any

	return system

}

func ToSystem(value string) System {

	var system System

	system.Name = Any
	system.Hostname = Any
	system.Datetime = Any
	system.Country = Any
	system.Timezone = Any

	system.SetName(value)

	return system

}

func (system *System) Equal(value System) bool {

	if system.Name == value.Name &&
		system.Hostname == value.Hostname &&
		system.Datetime == value.Datetime &&
		system.Country == value.Country &&
		system.Timezone == value.Timezone {
		return true
	}

	return false

}

func (system *System) IsValid() bool {

	if system.Name != Any ||
		system.Hostname != Any ||
		system.Datetime != Any ||
		system.Country != Any ||
		system.Timezone != Any {
		return true
	}

	return false

}

func (system *System) Matches(name string, hostname string, datetime string, country string, timezone string) bool {
	return system.MatchesName(name) && system.MatchesHostname(hostname) && system.MatchesDatetime(datetime) && system.MatchesCountry(country) && system.MatchesTimezone(timezone)
}

func (system *System) MatchesCountry(value string) bool {

	if system.Country == value {
		return true
	} else if system.Country == Any {
		return true
	}

	return false

}


func (system *System) MatchesDatetime(value string) bool {

	matches_from := false
	matches_until := false

	// Compatibility with "<operator> <version>" syntax
	if strings.Contains(value, " ") {

		tmp := strings.Split(value, " ")

		// >= 2001-02-03 04:05:06
		if len(tmp) == 3 {
			value = strings.TrimSpace(tmp[1] + " " + tmp[2])
		}

	}

	if system.Datetime == Any {

		matches_from = true
		matches_until = true

	} else if strings.Contains(system.Datetime, " - ") {

		tmp1 := strings.TrimSpace(system.Datetime[0:strings.Index(system.Datetime, " - ")])
		tmp2 := strings.TrimSpace(system.Datetime[strings.Index(system.Datetime, " - ")+3:])

		from := types.ToDatetime(tmp1)
		until := types.ToDatetime(tmp2)
		current := types.ToDatetime(value)

		if current.IsValid() && from.IsBefore(until) {

			if from.IsBefore(current) {
				matches_from = true
			}

			if until.IsAfter(current) {
				matches_until = true
			}

		}

	} else if strings.HasPrefix(system.Datetime, "<= ") {

		current := types.ToDatetime(value)
		until := types.ToDatetime(strings.TrimSpace(system.Datetime[3:]))
		matches_from = true

		if until.IsValid() {

			if until.Equal(current) {
				matches_until = true
			} else if until.IsAfter(current) {
				matches_until = true
			}

		}

	} else if strings.HasPrefix(system.Datetime, "< ") {

		current := types.ToDatetime(value)
		until := types.ToDatetime(strings.TrimSpace(system.Datetime[2:]))
		matches_from = true

		if current.IsValid() && until.IsValid() {

			if until.IsAfter(current) {
				matches_until = true
			}

		}

	} else if strings.HasPrefix(system.Datetime, ">= ") {

		current := types.ToDatetime(value)
		from := types.ToDatetime(strings.TrimSpace(system.Datetime[3:]))
		matches_until = true

		if current.IsValid() && from.IsValid() {

			if from.Equal(current) {
				matches_from = true
			} else if from.IsBefore(current) {
				matches_from = true
			}

		}

	} else if strings.HasPrefix(system.Datetime, "> ") {

		current := types.ToDatetime(value)
		from := types.ToDatetime(strings.TrimSpace(system.Datetime[2:]))
		matches_until = true

		if current.IsValid() && from.IsValid() {

			if from.IsBefore(current) {
				matches_from = true
			}

		}

	} else if strings.HasPrefix(system.Datetime, "= ") {

		current := types.ToDatetime(value)
		datetime := types.ToDatetime(strings.TrimSpace(system.Datetime[2:]))

		if current.IsValid() && datetime.IsValid() {

			if datetime.Equal(current) {
				matches_from = true
				matches_until = true
			}

		}

	}

	return matches_from && matches_until

}

func (system *System) MatchesHostname(value string) bool {

	if system.Hostname == value {
		return true
	} else if system.Hostname == Any {
		return true
	}

	return false

}

func (system *System) MatchesName(value string) bool {

	if system.Name == value {
		return true
	} else if system.Name == Any {
		return true
	}

	return false

}

func (system *System) MatchesTimezone(value string) bool {

	if system.Timezone == value {

		return true

	} else if system.Timezone == Any {

		return true

	} else if strings.HasPrefix(system.Timezone, "*/") {

		if strings.HasSuffix(value, system.Timezone[strings.Index(system.Timezone, "*/")+1:]) {
			return true
		}

	} else if strings.HasSuffix(system.Timezone, "/*") {

		if strings.HasPrefix(value, system.Timezone[0:strings.Index(system.Timezone, "/*")+1]) {
			return true
		}

	}

	return false

}

func (system *System) SetCountry(value string) {

	if value == "all" || value == "any" || value == "*" {
		system.Country = Any
	} else if value != "" {
		system.Country = value
	}

}

func (system *System) SetHostname(value string) {

	if value == "all" || value == "any" || value == "*" {
		system.Hostname = Any
	} else if value != "" {
		system.Hostname = strings.TrimSpace(value)
	}

}

func (system *System) SetName(value string) {
	system.Name = strings.TrimSpace(value)
}

func (system *System) SetDatetime(value string) {

	if value == "all" || value == "any" || value == "*" {

		system.Datetime = Any

	} else if strings.Contains(value, " - ") {

		tmp1 := strings.TrimSpace(value[0:strings.Index(value, " - ")])
		tmp2 := strings.TrimSpace(value[strings.Index(value, " - ")+3:])

		from := types.ToDatetime(tmp1)
		until := types.ToDatetime(tmp2)

		if from.IsBefore(until) {
			system.Datetime = from.String() + " - " + until.String()
		}

	} else if strings.HasPrefix(value, "<= ") {

		until := types.ToDatetime(strings.TrimSpace(value[3:]))

		if until.IsValid() {
			system.Datetime = "<= " + until.String()
		}

	} else if strings.HasPrefix(value, "< ") {

		until := types.ToDatetime(strings.TrimSpace(value[2:]))

		if until.IsValid() {
			system.Datetime = "< " + until.String()
		}

	} else if strings.HasPrefix(value, ">= ") {

		from := types.ToDatetime(strings.TrimSpace(value[3:]))

		if from.IsValid() {
			system.Datetime = ">= " + from.String()
		}

	} else if strings.HasPrefix(value, "> ") {

		from := types.ToDatetime(strings.TrimSpace(value[2:]))

		if from.IsValid() {
			system.Datetime = "> " + from.String()
		}

	} else if strings.HasPrefix(value, "= ") {

		datetime := types.ToDatetime(strings.TrimSpace(value[2:]))

		if datetime.IsValid() {
			system.Datetime = "= " + datetime.String()
		}

	} else {

		datetime := types.ToDatetime(strings.TrimSpace(value))

		if datetime.IsValid() {
			system.Datetime = "= " + datetime.String()
		}

	}

}

func (system *System) SetTimezone(value string) {

	if value == "all" || value == "any" || value == "*" {
		system.Timezone = Any
	} else if value != "" {
		system.Timezone = strings.TrimSpace(value)
	}

}

func (system *System) Hash() string {

	hash := ""

	if system.Name != "" {

		hash = toHash(
			system.Name,
			system.Hostname,
			system.Datetime,
			system.Country,
			system.Timezone,
		)

	}

	return hash

}
