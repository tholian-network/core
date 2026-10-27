package caches

import "strconv"
import "strings"

func isPrefix(value string) bool {

	if strings.HasPrefix(value, "/") {

		_, err := strconv.ParseUint(value[1:], 10, 8)

		if err == nil {
			return true
		}

	} else if strings.HasPrefix(value, "[") && strings.Contains(value, "]/") {

		_, err := strconv.ParseUint(strings.Split(value, "]/")[1], 10, 8)

		if err == nil {
			return true
		}

	} else if strings.HasPrefix(value, "[") && strings.HasSuffix(value, "]") {

		return true

	} else if strings.Contains(value, ":") && strings.Contains(value, "/") {

		_, err := strconv.ParseUint(strings.Split(value, "/")[1], 10, 8)

		if err == nil {
			return true
		}

	} else if strings.Contains(value, ":") {

		return true

	} else if strings.Contains(value, ".") && strings.Contains(value, "/") {

		_, err := strconv.ParseUint(strings.Split(value, "/")[1], 10, 8)

		if err == nil {
			return true
		}

	} else if strings.Contains(value, ".") {

		return true

	} else {

		_, err := strconv.ParseUint(value, 10, 8)

		if err == nil {
			return true
		}

	}

	return false

}
