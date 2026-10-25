package matchers

import "tholian-endpoint/types"
import "encoding/binary"
import "encoding/hex"
import "hash/crc32"
import "strings"

type Datetime struct {
	From  string `json:"from"`
	Until string `json:"until"`
}

func NewDatetime() Datetime {

	var datetime Datetime

	datetime.From = Any
	datetime.Until = Any

	return datetime

}

func ToDatetime(value string) Datetime {

	var datetime Datetime

	datetime.From = Any
	datetime.Until = Any

	datetime.Parse(value)

	return datetime

}

func (datetime *Datetime) IsIdentical(value Datetime) bool {

	if datetime.From == value.From && datetime.Until == value.Until {
		return true
	}

	return false

}

func (datetime *Datetime) IsValid() bool {

	if datetime.From != Any || datetime.Until != Any {
		return true
	}

	return false

}

func (datetime *Datetime) Matches(datetime_ string) bool {
	return datetime.MatchesDatetime(datetime_)
}

func (datetime *Datetime) MatchesDatetime(value string) bool {

	var matches_from bool
	var matches_until bool

	current := types.ToDatetime(value)

	if current.IsValid() {

		if datetime.From == Any {

			matches_from = true

		} else {

			from := types.ToDatetime(datetime.From)

			if from.IsBefore(current) {
				matches_from = true
			}

		}

		if datetime.Until == Any {

			matches_until = true

		} else {

			until := types.ToDatetime(datetime.Until)

			if until.IsBefore(current) {
				matches_until = true
			}

		}

	}

	return matches_from && matches_until

}

func (datetime *Datetime) Parse(value string) {

	if strings.Contains(value, " - ") {

		tmp1 := strings.TrimSpace(value[0:strings.Index(value, " - ")])
		tmp2 := strings.TrimSpace(value[strings.Index(value, " - ")+3:])

		from := types.ToDatetime(tmp1)
		until := types.ToDatetime(tmp2)

		if from.IsBefore(until) {
			datetime.From = from.String()
			datetime.Until = until.String()
		}

	} else if strings.HasPrefix(value, "< ") {

		tmp1 := strings.TrimSpace(value[2:])
		until := types.ToDatetime(tmp1)

		if until.IsValid() {
			datetime.Until = until.String()
		}

	} else if strings.HasPrefix(value, "> ") {

		tmp1 := strings.TrimSpace(value[2:])
		from := types.ToDatetime(tmp1)

		if from.IsValid() {
			datetime.From = from.String()
		}

	}

}

func (datetime *Datetime) SetFrom(value string) {

	from := types.ToDatetime(value)

	if from.IsValid() {
		datetime.From = from.String()
	}

}

func (datetime *Datetime) SetUntil(value string) {

	until := types.ToDatetime(value)

	if until.IsValid() {
		datetime.Until = until.String()
	}

}

func (datetime *Datetime) Hash() string {

	var hash string

	checksum := crc32.ChecksumIEEE([]byte(strings.Join([]string{
		datetime.From,
		datetime.Until,
	}, "-")))

	tmp := make([]byte, 4)
	binary.LittleEndian.PutUint32(tmp, checksum)
	hash = hex.EncodeToString(tmp)

	return hash

}
