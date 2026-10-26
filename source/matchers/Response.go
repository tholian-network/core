package matchers

import "tholian-core/types"
import "encoding/binary"
import "encoding/hex"
import "hash/crc32"
import "strings"

type Response struct {
	System   string `json:"system"`
	Type     string `json:"type"`
	Datetime string `json:"datetime"`
}

func NewResponse() Response {

	var response Response

	response.System = Any
	response.Type = Any
	response.Datetime = Any

	return response

}

func ToResponse(value string) Response {

	var response Response

	response.System = Any
	response.Type = Any
	response.Datetime = Any

	if strings.Contains(value, ":") {

		tmp := strings.Split(value, ":")

		if len(tmp) == 2 {

			response.SetSystem(tmp[0])
			response.SetType(tmp[1])

		}

	} else {
		response.SetSystem(value)
	}

	return response

}

func (response *Response) Equal(value Response) bool {

	if response.System == value.System &&
		response.Type == value.Type &&
		response.Datetime == value.Datetime {
		return true
	}

	return false

}

func (response *Response) IsValid() bool {

	if response.System != Any ||
		response.Type != Any ||
		response.Datetime != Any {
		return true
	}

	return false

}

func (response *Response) Matches(system string, typ string, datetime string) bool {
	return response.MatchesSystem(system) && response.MatchesType(typ) && response.MatchesDatetime(datetime)
}

func (response *Response) MatchesDatetime(value string) bool {

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

	if response.Datetime == Any {

		matches_from = true
		matches_until = true

	} else if strings.Contains(response.Datetime, " - ") {

		tmp1 := strings.TrimSpace(response.Datetime[0:strings.Index(response.Datetime, " - ")])
		tmp2 := strings.TrimSpace(response.Datetime[strings.Index(response.Datetime, " - ")+3:])

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

	} else if strings.HasPrefix(response.Datetime, "<= ") {

		current := types.ToDatetime(value)
		until := types.ToDatetime(strings.TrimSpace(response.Datetime[3:]))
		matches_from = true

		if until.IsValid() {

			if until.Equal(current) {
				matches_until = true
			} else if until.IsAfter(current) {
				matches_until = true
			}

		}

	} else if strings.HasPrefix(response.Datetime, "< ") {

		current := types.ToDatetime(value)
		until := types.ToDatetime(strings.TrimSpace(response.Datetime[2:]))
		matches_from = true

		if current.IsValid() && until.IsValid() {

			if until.IsAfter(current) {
				matches_until = true
			}

		}

	} else if strings.HasPrefix(response.Datetime, ">= ") {

		current := types.ToDatetime(value)
		from := types.ToDatetime(strings.TrimSpace(response.Datetime[3:]))
		matches_until = true

		if current.IsValid() && from.IsValid() {

			if from.Equal(current) {
				matches_from = true
			} else if from.IsBefore(current) {
				matches_from = true
			}

		}

	} else if strings.HasPrefix(response.Datetime, "> ") {

		current := types.ToDatetime(value)
		from := types.ToDatetime(strings.TrimSpace(response.Datetime[2:]))
		matches_until = true

		if current.IsValid() && from.IsValid() {

			if from.IsBefore(current) {
				matches_from = true
			}

		}

	} else if strings.HasPrefix(response.Datetime, "= ") {

		current := types.ToDatetime(value)
		datetime := types.ToDatetime(strings.TrimSpace(response.Datetime[2:]))

		if current.IsValid() && datetime.IsValid() {

			if datetime.Equal(current) {
				matches_from = true
				matches_until = true
			}

		}

	}

	return matches_from && matches_until

}

func (response *Response) MatchesSystem(value string) bool {

	if response.System == value {
		return true
	} else if response.System == Any {
		return true
	}

	return false

}

func (response *Response) MatchesType(value string) bool {

	if response.Type == value {
		return true
	} else if response.Type == Any {
		return true
	}

	return false

}

func (response *Response) SetDatetime(value string) {

	if value == "all" || value == "any" || value == "*" {

		response.Datetime = Any

	} else if strings.Contains(value, " - ") {

		tmp1 := strings.TrimSpace(value[0:strings.Index(value, " - ")])
		tmp2 := strings.TrimSpace(value[strings.Index(value, " - ")+3:])

		from := types.ToDatetime(tmp1)
		until := types.ToDatetime(tmp2)

		if from.IsBefore(until) {
			response.Datetime = from.String() + " - " + until.String()
		}

	} else if strings.HasPrefix(value, "<= ") {

		until := types.ToDatetime(strings.TrimSpace(value[3:]))

		if until.IsValid() {
			response.Datetime = "<= " + until.String()
		}

	} else if strings.HasPrefix(value, "< ") {

		until := types.ToDatetime(strings.TrimSpace(value[2:]))

		if until.IsValid() {
			response.Datetime = "< " + until.String()
		}

	} else if strings.HasPrefix(value, ">= ") {

		from := types.ToDatetime(strings.TrimSpace(value[3:]))

		if from.IsValid() {
			response.Datetime = ">= " + from.String()
		}

	} else if strings.HasPrefix(value, "> ") {

		from := types.ToDatetime(strings.TrimSpace(value[2:]))

		if from.IsValid() {
			response.Datetime = "> " + from.String()
		}

	} else if strings.HasPrefix(value, "= ") {

		datetime := types.ToDatetime(strings.TrimSpace(value[2:]))

		if datetime.IsValid() {
			response.Datetime = "= " + datetime.String()
		}

	} else {

		datetime := types.ToDatetime(strings.TrimSpace(value))

		if datetime.IsValid() {
			response.Datetime = "= " + datetime.String()
		}

	}

}

func (response *Response) SetSystem(value string) {
	response.System = strings.TrimSpace(value)
}

func (response *Response) SetType(value string) {

	if value == "all" || value == "any" || value == "*" {
		response.Type = Any
	} else if value != "" {
		response.Type = strings.TrimSpace(value)
	}

}

func (response *Response) Hash() string {

	hash := ""

	if response.Type != "" {

		checksum := crc32.ChecksumIEEE([]byte(strings.Join([]string{
			response.System,
			response.Type,
			response.Datetime,
		}, "-")))

		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, checksum)
		hash = hex.EncodeToString(tmp)

	}

	return hash

}
