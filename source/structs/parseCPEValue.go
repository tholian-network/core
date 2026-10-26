package structs

import utils_strings "tholian-core/utils/strings"
import "strings"

func parseCPEValue(value string) string {

	value = utils_strings.ToASCII(strings.TrimSpace(value))

	if value == "*" {
		value = ""
	} else if value == "-" {
		value = ""
	}

	return value

}
