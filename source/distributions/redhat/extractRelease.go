package redhat

import utils_strings "tholian-detective/utils/strings"
import "strings"

func extractRelease(platform string) string {

	if strings.HasPrefix(platform, "Oracle Java for Red Hat Enterprise Linux") {

		tmp := utils_strings.ExtractNumber(platform, " ")

		if tmp != "" {
			return tmp
		}

	} else if strings.HasPrefix(platform, "Red Hat Enterprise Linux 6.0 EUS") {

		return "6"

	} else if strings.HasPrefix(platform, "Red Hat Enterprise Linux") {

		tmp := utils_strings.ExtractNumber(platform, " ")

		if tmp != "" {
			return tmp
		}

	} else if strings.HasPrefix(platform, "Supplementary for Red Hat Enterprise Linux") {

		tmp := utils_strings.ExtractNumber(platform, " ")

		if tmp != "" {
			return tmp
		}

	}

	return ""

}
