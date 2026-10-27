package structs

import utils_strings "tholian-core/utils/strings"
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

func (timezone *Timezone) Equal(value Timezone) bool {

	if timezone.Name == value.Name && timezone.Offset == value.Offset {
		return true
	}

	return false

}

func (timezone *Timezone) IsValid() bool {

	if strings.TrimSpace(timezone.Name) == "" {
		return false
	}

	return utils_strings.IsTimezoneOffset(timezone.Offset)

}

func (timezone *Timezone) String() string {

	result := ""

	if timezone.IsValid() {
		result = timezone.Name + " (UTC" + timezone.Offset + ")"
	}

	return result

}

func (timezone *Timezone) Hash() string {

	hash := ""

	if timezone.Name != "" {

		hash = toHash(
			timezone.Name,
			timezone.Offset,
		)

	}

	return hash

}

func (timezone *Timezone) SetName(value string) {

	value = strings.TrimSpace(value)

	if value != "" {
		timezone.Name = value
	}

}

func (timezone *Timezone) SetOffset(value string) {

	if utils_strings.IsTimezoneOffset(value) {
		timezone.Offset = value
	}

}

