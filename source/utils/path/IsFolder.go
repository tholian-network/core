package path

import "path"
import "strings"

func IsFolder(value string) bool {

	if strings.HasPrefix(value, "/") &&
		value == path.Clean(value) {
		return true
	}

	return false

}
