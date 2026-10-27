package matchers

import "tholian-core/types"
import "strings"

type Time struct {
	From  string `json:"from"`
	Until string `json:"until"`
}

func NewTime() Time {

	var time Time

	time.From = Any
	time.Until = Any

	return time

}

func ToTime(value string) Time {

	var time Time

	time.From = Any
	time.Until = Any

	time.Parse(value)

	return time

}

func (time *Time) Equal(value Time) bool {

	if time.From == value.From &&
		time.Until == value.Until {
		return true
	}

	return false

}

func (time *Time) IsValid() bool {

	if time.From != Any ||
		time.Until != Any {
		return true
	}

	return false

}

func (time *Time) Matches(time_ string) bool {
	return time.MatchesTime(time_)
}

func (time *Time) MatchesTime(value string) bool {

	matches_from := false
	matches_until := false

	current := types.ToTime(value)

	if current.IsValid() {

		if time.From == Any {

			matches_from = true

		} else {

			from := types.ToTime(time.From)

			if from.IsBefore(current) {
				matches_from = true
			}

		}

		if time.Until == Any {

			matches_until = true

		} else {

			until := types.ToTime(time.Until)

			if until.IsBefore(current) {
				matches_until = true
			}

		}

	}

	return matches_from && matches_until

}

func (time *Time) Parse(value string) {

	if strings.Contains(value, " - ") {

		tmp1 := strings.TrimSpace(value[0:strings.Index(value, " - ")])
		tmp2 := strings.TrimSpace(value[strings.Index(value, " - ")+3:])

		from := types.ToTime(tmp1)
		until := types.ToTime(tmp2)

		if from.IsBefore(until) {
			time.From = from.String()
			time.Until = until.String()
		}

	} else if strings.HasPrefix(value, "< ") {

		tmp1 := strings.TrimSpace(value[2:])
		until := types.ToTime(tmp1)

		if until.IsValid() {
			time.Until = until.String()
		}

	} else if strings.HasPrefix(value, "> ") {

		tmp1 := strings.TrimSpace(value[2:])
		from := types.ToTime(tmp1)

		if from.IsValid() {
			time.From = from.String()
		}

	}

}

func (time *Time) SetFrom(value string) {

	from := types.ToTime(value)

	if from.IsValid() {
		time.From = from.String()
	}

}

func (time *Time) SetUntil(value string) {

	until := types.ToTime(value)

	if until.IsValid() {
		time.Until = until.String()
	}

}

func (time *Time) Hash() string {

	hash := ""

	hash = toHash(
		time.From,
		time.Until,
	)

	return hash

}
