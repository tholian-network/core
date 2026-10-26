package matchers

import utils_fmt "tholian-core/utils/fmt"
import "strings"

func parseCPEValue(value string) string {

	value = utils_fmt.FormatASCII(strings.TrimSpace(value))

	if value == "*" {
		value = ""
	} else if value == "-" {
		value = ""
	}

	return value

}

