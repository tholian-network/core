package matchers

import "tholian-core/types"
import "strings"

type Evidence struct {
	Type     string `json:"type"`
	Datetime string `json:"datetime"`
}

func NewEvidence() Evidence {

	var evidence Evidence

	evidence.Type = Any
	evidence.Datetime = Any

	return evidence

}

func ToEvidence(value string) Evidence {

	var evidence Evidence

	evidence.Type = Any
	evidence.Datetime = Any

	evidence.SetType(value)

	return evidence

}

func (evidence *Evidence) Equal(value Evidence) bool {

	if evidence.Type == value.Type &&
		evidence.Datetime == value.Datetime {
		return true
	}

	return false

}

func (evidence *Evidence) IsValid() bool {

	if evidence.Type != Any ||
		evidence.Datetime != Any {
		return true
	}

	return false

}

func (evidence *Evidence) Matches(typ string, datetime string) bool {
	return evidence.MatchesType(typ) && evidence.MatchesDatetime(datetime)
}

func (evidence *Evidence) MatchesType(value string) bool {

	if evidence.Type == value {
		return true
	} else if evidence.Type == Any {
		return true
	}

	return false

}

func (evidence *Evidence) MatchesDatetime(value string) bool {

	if evidence.Datetime == Any {
		return true
	}

	current := types.ToDatetime(value)

	if current.IsValid() == false {
		return false
	}

	if strings.Contains(evidence.Datetime, " - ") {

		tmp1 := strings.TrimSpace(evidence.Datetime[0:strings.Index(evidence.Datetime, " - ")])
		tmp2 := strings.TrimSpace(evidence.Datetime[strings.Index(evidence.Datetime, " - ")+3:])

		from := types.ToDatetime(tmp1)
		until := types.ToDatetime(tmp2)

		if from.IsValid() && until.IsValid() {

			matches_from := from.IsBefore(current) || from.Equal(current)
			matches_until := until.IsAfter(current) || until.Equal(current)

			return matches_from && matches_until

		}

	} else if strings.HasPrefix(evidence.Datetime, "<= ") {

		until := types.ToDatetime(strings.TrimSpace(evidence.Datetime[3:]))

		if until.IsValid() {
			return until.IsAfter(current) || until.Equal(current)
		}

	} else if strings.HasPrefix(evidence.Datetime, "< ") {

		until := types.ToDatetime(strings.TrimSpace(evidence.Datetime[2:]))

		if until.IsValid() {
			return until.IsAfter(current)
		}

	} else if strings.HasPrefix(evidence.Datetime, ">= ") {

		from := types.ToDatetime(strings.TrimSpace(evidence.Datetime[3:]))

		if from.IsValid() {
			return from.IsBefore(current) || from.Equal(current)
		}

	} else if strings.HasPrefix(evidence.Datetime, "> ") {

		from := types.ToDatetime(strings.TrimSpace(evidence.Datetime[2:]))

		if from.IsValid() {
			return from.IsBefore(current)
		}

	} else if strings.HasPrefix(evidence.Datetime, "= ") {

		expected := types.ToDatetime(strings.TrimSpace(evidence.Datetime[2:]))

		if expected.IsValid() {
			return expected.Equal(current)
		}

	} else {

		expected := types.ToDatetime(evidence.Datetime)

		if expected.IsValid() {
			return expected.Equal(current)
		}

	}

	return false

}

func (evidence *Evidence) SetType(value string) {

	value = strings.TrimSpace(value)

	if value == "all" || value == "any" || value == "*" {
		evidence.Type = Any
	} else if value != "" {
		evidence.Type = value
	}

}

func (evidence *Evidence) SetDatetime(value string) {

	value = strings.TrimSpace(value)

	if value == "all" || value == "any" || value == "*" {
		evidence.Datetime = Any
	} else if value != "" {

		current := types.ToDatetime(value)
		matcher := ToDatetime(value)

		if current.IsValid() || matcher.IsValid() {
			evidence.Datetime = value
		}

	}

}

func (evidence *Evidence) Hash() string {

	hash := ""

	if evidence.Type != "" {

		hash = toHash(
			evidence.Type,
			evidence.Datetime,
		)

	}

	return hash

}
