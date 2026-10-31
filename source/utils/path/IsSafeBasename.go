package path

import "strings"

func IsSafeBasename(value string) bool {

	if value == "" || value == "." || value == ".." {
		return false
	}

	if strings.ContainsAny(value, "/\\\x00") {
		return false
	}

	return true

}
