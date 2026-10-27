package caches

import "strconv"
import "strings"

type prefix uint8

func (value prefix) String() string {
	return "/" + strconv.FormatUint(uint64(value), 10)
}

func (prefix prefix) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote("/" + strconv.FormatUint(uint64(prefix), 10))), nil
}

func (prefix *prefix) UnmarshalJSON(data []byte) error {

	unquoted, err := strconv.Unquote(string(data))

	if err != nil {
		return err
	}

	tmp := parsePrefix(unquoted)

	if tmp != nil {
		*prefix = *tmp
	}

	return nil

}

func parsePrefix(value string) *prefix {

	var pointer *prefix = nil

	if strings.HasPrefix(value, "/") {

		tmp, err := strconv.ParseUint(value[1:], 10, 8)

		if err == nil {
			tmp_val := prefix(tmp)
			pointer = &tmp_val
		}

	} else if strings.HasPrefix(value, "[") && strings.Contains(value, "]/") {

		tmp, err := strconv.ParseUint(strings.Split(value, "]/")[1], 10, 8)

		if err == nil {
			tmp_val := prefix(tmp)
			pointer = &tmp_val
		}

	} else if strings.HasPrefix(value, "[") && strings.HasSuffix(value, "]") {

		tmp_val := prefix(128)
		pointer = &tmp_val

	} else if strings.Contains(value, ":") && strings.Contains(value, "/") {

		tmp, err := strconv.ParseUint(strings.Split(value, "/")[1], 10, 8)

		if err == nil {
			tmp_val := prefix(tmp)
			pointer = &tmp_val
		}

	} else if strings.Contains(value, ":") {

		tmp_val := prefix(128)
		pointer = &tmp_val

	} else if strings.Contains(value, ".") && strings.Contains(value, "/") {

		tmp, err := strconv.ParseUint(strings.Split(value, "/")[1], 10, 8)

		if err == nil {
			tmp_val := prefix(tmp)
			pointer = &tmp_val
		}

	} else if strings.Contains(value, ".") {

		tmp_val := prefix(32)
		pointer = &tmp_val

	} else {

		tmp, err := strconv.ParseUint(value, 10, 8)

		if err == nil {
			tmp_val := prefix(tmp)
			pointer = &tmp_val
		}

	}

	return pointer

}

