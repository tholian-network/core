package matchers

import "tholian-core/types"
import "strings"

type Incident struct {
	System   string `json:"system"`
	Type     string `json:"type"`
	Datetime string `json:"datetime"`
}

func NewIncident() Incident {

	var incident Incident

	incident.System = Any
	incident.Type = Any
	incident.Datetime = Any

	return incident

}

func ToIncident(value string) Incident {

	var incident Incident

	incident.System = Any
	incident.Type = Any
	incident.Datetime = Any

	if strings.Contains(value, ":") {

		tmp := strings.Split(value, ":")

		if len(tmp) == 2 {

			incident.SetSystem(tmp[0])
			incident.SetType(tmp[1])

		}

	} else {
		incident.SetSystem(value)
	}

	return incident

}

func (incident *Incident) Equal(value Incident) bool {

	if incident.System == value.System &&
		incident.Type == value.Type &&
		incident.Datetime == value.Datetime {
		return true
	}

	return false

}

func (incident *Incident) IsValid() bool {

	if incident.System != Any ||
		incident.Type != Any ||
		incident.Datetime != Any {
		return true
	}

	return false

}

func (incident *Incident) Matches(system string, typ string, datetime string) bool {
	return incident.MatchesSystem(system) && incident.MatchesType(typ) && incident.MatchesDatetime(datetime)
}

func (incident *Incident) MatchesDatetime(value string) bool {

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

	if incident.Datetime == Any {

		matches_from = true
		matches_until = true

	} else if strings.Contains(incident.Datetime, " - ") {

		tmp1 := strings.TrimSpace(incident.Datetime[0:strings.Index(incident.Datetime, " - ")])
		tmp2 := strings.TrimSpace(incident.Datetime[strings.Index(incident.Datetime, " - ")+3:])

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

	} else if strings.HasPrefix(incident.Datetime, "<= ") {

		current := types.ToDatetime(value)
		until := types.ToDatetime(strings.TrimSpace(incident.Datetime[3:]))
		matches_from = true

		if until.IsValid() {

			if until.Equal(current) {
				matches_until = true
			} else if until.IsAfter(current) {
				matches_until = true
			}

		}

	} else if strings.HasPrefix(incident.Datetime, "< ") {

		current := types.ToDatetime(value)
		until := types.ToDatetime(strings.TrimSpace(incident.Datetime[2:]))
		matches_from = true

		if current.IsValid() && until.IsValid() {

			if until.IsAfter(current) {
				matches_until = true
			}

		}

	} else if strings.HasPrefix(incident.Datetime, ">= ") {

		current := types.ToDatetime(value)
		from := types.ToDatetime(strings.TrimSpace(incident.Datetime[3:]))
		matches_until = true

		if current.IsValid() && from.IsValid() {

			if from.Equal(current) {
				matches_from = true
			} else if from.IsBefore(current) {
				matches_from = true
			}

		}

	} else if strings.HasPrefix(incident.Datetime, "> ") {

		current := types.ToDatetime(value)
		from := types.ToDatetime(strings.TrimSpace(incident.Datetime[2:]))
		matches_until = true

		if current.IsValid() && from.IsValid() {

			if from.IsBefore(current) {
				matches_from = true
			}

		}

	} else if strings.HasPrefix(incident.Datetime, "= ") {

		current := types.ToDatetime(value)
		datetime := types.ToDatetime(strings.TrimSpace(incident.Datetime[2:]))

		if current.IsValid() && datetime.IsValid() {

			if datetime.Equal(current) {
				matches_from = true
				matches_until = true
			}

		}

	}

	return matches_from && matches_until

}

func (incident *Incident) MatchesSystem(value string) bool {

	if incident.System == value {
		return true
	} else if incident.System == Any {
		return true
	}

	return false

}

func (incident *Incident) MatchesType(value string) bool {

	if incident.Type == value {
		return true
	} else if incident.Type == Any {
		return true
	}

	return false

}

func (incident *Incident) SetDatetime(value string) {

	if value == "all" || value == "any" || value == "*" {

		incident.Datetime = Any

	} else if strings.Contains(value, " - ") {

		tmp1 := strings.TrimSpace(value[0:strings.Index(value, " - ")])
		tmp2 := strings.TrimSpace(value[strings.Index(value, " - ")+3:])

		from := types.ToDatetime(tmp1)
		until := types.ToDatetime(tmp2)

		if from.IsBefore(until) {
			incident.Datetime = from.String() + " - " + until.String()
		}

	} else if strings.HasPrefix(value, "<= ") {

		until := types.ToDatetime(strings.TrimSpace(value[3:]))

		if until.IsValid() {
			incident.Datetime = "<= " + until.String()
		}

	} else if strings.HasPrefix(value, "< ") {

		until := types.ToDatetime(strings.TrimSpace(value[2:]))

		if until.IsValid() {
			incident.Datetime = "< " + until.String()
		}

	} else if strings.HasPrefix(value, ">= ") {

		from := types.ToDatetime(strings.TrimSpace(value[3:]))

		if from.IsValid() {
			incident.Datetime = ">= " + from.String()
		}

	} else if strings.HasPrefix(value, "> ") {

		from := types.ToDatetime(strings.TrimSpace(value[2:]))

		if from.IsValid() {
			incident.Datetime = "> " + from.String()
		}

	} else if strings.HasPrefix(value, "= ") {

		datetime := types.ToDatetime(strings.TrimSpace(value[2:]))

		if datetime.IsValid() {
			incident.Datetime = "= " + datetime.String()
		}

	} else {

		datetime := types.ToDatetime(strings.TrimSpace(value))

		if datetime.IsValid() {
			incident.Datetime = "= " + datetime.String()
		}

	}

}

func (incident *Incident) SetSystem(value string) {
	incident.System = strings.TrimSpace(value)
}

func (incident *Incident) SetType(value string) {

	if value == "all" || value == "any" || value == "*" {
		incident.Type = Any
	} else if value != "" {
		incident.Type = strings.TrimSpace(value)
	}

}

func (incident *Incident) Hash() string {

	hash := ""

	if incident.Type != "" {

		hash = toHash(
			incident.System,
			incident.Type,
			incident.Datetime,
		)

	}

	return hash

}
