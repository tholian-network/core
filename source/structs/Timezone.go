package structs

import "strings"

type Timezone struct {
	Name   string `json:"name"`
	Offset string `json:"offset"`
}

func NewTimezone(name string, offset string) Timezone {

	var timezone Timezone

	timezone.SetName(name)
	timezone.SetOffset(offset)

	return timezone

}

func (timezone *Timezone) SetName(value string) {

	value = strings.TrimSpace(value)

	if value != "" {
		timezone.Name = value
	}

}

func (timezone *Timezone) SetOffset(value string) {

	if strings.HasPrefix(value, "+") || strings.HasPrefix(value, "-") {

		if len(value) == 6 && string(value[3]) == ":" {
			timezone.Offset = value
		}

	}

}
