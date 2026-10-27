package matchers

import "tholian-core/types"
import "strings"

type Timeslot struct {
	Offset   string `json:"offset"`
	Country  string `json:"country"`
	Timezone string `json:"timezone"`
}

func NewTimeslot() Timeslot {

	var timeslot Timeslot

	timeslot.Offset = Any
	timeslot.Country = Any
	timeslot.Timezone = Any

	return timeslot

}

func ToTimeslot(value string) Timeslot {

	var timeslot Timeslot

	timeslot.Offset = Any
	timeslot.Country = Any
	timeslot.Timezone = Any

	timeslot.Parse(value)

	return timeslot

}

func (timeslot *Timeslot) Equal(value Timeslot) bool {

	if timeslot.Offset == value.Offset &&
		timeslot.Country == value.Country &&
		timeslot.Timezone == value.Timezone {
		return true
	}

	return false

}

func (timeslot *Timeslot) IsValid() bool {

	if timeslot.Offset != Any ||
		timeslot.Country != Any ||
		timeslot.Timezone != Any {
		return true
	}

	return false

}

func (timeslot *Timeslot) Matches(offset string, country string, timezone string) bool {
	return timeslot.MatchesOffset(offset) && timeslot.MatchesCountry(country) && timeslot.MatchesTimezone(timezone)
}

func (timeslot *Timeslot) MatchesCountry(value string) bool {

	if timeslot.Country == value {
		return true
	} else if timeslot.Country == Any {
		return true
	}

	return false

}

func (timeslot *Timeslot) MatchesOffset(value string) bool {

	// Compatibility with "<operator> <offset>" syntax
	if strings.Contains(value, " ") {
		value = strings.TrimSpace(value[strings.Index(value, " ")+1:])
	}

	if timeslot.Offset == Any {

		return true

	} else if strings.HasPrefix(timeslot.Offset, "<= ") {

		timeslot_offset := types.NewTime()
		timeslot_offset.Offset(timeslot.Offset[3:])

		other_offset := types.NewTime()
		other_offset.Offset(value)

		if other_offset.Equal(timeslot_offset) {
			return true
		} else if other_offset.IsBefore(timeslot_offset) {
			return true
		}

	} else if strings.HasPrefix(timeslot.Offset, "< ") {

		timeslot_offset := types.NewTime()
		timeslot_offset.Offset(timeslot.Offset[2:])

		other_offset := types.NewTime()
		other_offset.Offset(value)

		if other_offset.IsBefore(timeslot_offset) {
			return true
		}

	} else if strings.HasPrefix(timeslot.Offset, ">= ") {

		timeslot_offset := types.NewTime()
		timeslot_offset.Offset(timeslot.Offset[3:])

		other_offset := types.NewTime()
		other_offset.Offset(value)

		if other_offset.Equal(timeslot_offset) {
			return true
		} else if other_offset.IsAfter(timeslot_offset) {
			return true
		}

	} else if strings.HasPrefix(timeslot.Offset, "> ") {

		timeslot_offset := types.NewTime()
		timeslot_offset.Offset(timeslot.Offset[2:])

		other_offset := types.NewTime()
		other_offset.Offset(value)

		if other_offset.IsAfter(timeslot_offset) {
			return true
		}

	} else if strings.HasPrefix(timeslot.Offset, "= ") {

		timeslot_offset := types.NewTime()
		timeslot_offset.Offset(timeslot.Offset[2:])

		other_offset := types.NewTime()
		other_offset.Offset(value)

		if other_offset.Equal(timeslot_offset) {
			return true
		}

	} else {

		timeslot_offset := types.NewTime()
		timeslot_offset.Offset(timeslot.Offset)

		other_offset := types.NewTime()
		other_offset.Offset(value)

		if other_offset.Equal(timeslot_offset) {
			return true
		}

	}

	return false

}

func (timeslot *Timeslot) MatchesTimezone(value string) bool {

	if timeslot.Timezone == value {

		return true

	} else if timeslot.Timezone == Any {

		return true

	} else if strings.HasPrefix(timeslot.Timezone, "*/") {

		if strings.HasSuffix(value, timeslot.Timezone[strings.Index(timeslot.Timezone, "*/")+1:]) {
			return true
		}

	} else if strings.HasSuffix(timeslot.Timezone, "/*") {

		if strings.HasPrefix(value, timeslot.Timezone[0:strings.Index(timeslot.Timezone, "/*")+1]) {
			return true
		}

	}

	return false

}

func (timeslot *Timeslot) Parse(value string) {

	timezone, offset := parseOffsetCondition(value)

	timeslot.Offset = offset
	timeslot.Timezone = timezone

}

func (timeslot *Timeslot) SetCountry(value string) {

	if value == "all" || value == "any" || value == "*" {
		timeslot.Country = Any
	} else if value != "" {
		timeslot.Country = value
	}

}

func (timeslot *Timeslot) SetOffset(value string) {

	if value == "all" || value == "any" || value == "*" {

		timeslot.Offset = Any

	} else if strings.HasPrefix(value, "+") || strings.HasPrefix(value, "-") {

		if len(value) == 6 && string(value[3]) == ":" {
			timeslot.Offset = value
		}

	}

}

func (timeslot *Timeslot) SetTimezone(value string) {

	if value == "all" || value == "any" || value == "*" {
		timeslot.Timezone = Any
	} else if value != "" {
		timeslot.Timezone = strings.TrimSpace(value)
	}

}

func (timeslot *Timeslot) Hash() string {

	hash := ""

	if timeslot.Offset != "" {

		hash = toHash(
			timeslot.Offset,
			timeslot.Country,
			timeslot.Timezone,
		)

	}

	return hash

}
