package strings

import "strconv"
import "strings"

func IsTimezoneOffset(value string) bool {

	if len(value) != 6 || string(value[3]) != ":" {
		return false
	}

	if strings.HasPrefix(value, "+") || strings.HasPrefix(value, "-") {

		hours, err_hours := strconv.ParseUint(value[1:3], 10, 8)
		minutes, err_minutes := strconv.ParseUint(value[4:6], 10, 8)

		if err_hours == nil && err_minutes == nil {
			return hours <= 23 && minutes <= 59
		}

	}

	return false

}
