package suse

import utils_strings "tholian-detective/utils/strings"
import "strings"

func extractRelease(platform string) string {

	if strings.HasPrefix(platform, "SUSE CORE") {

		tmp := utils_strings.ExtractNumber(platform, " ")

		if tmp != "" {
			return "server-" + tmp
		}

	} else if strings.HasPrefix(platform, "SUSE Linux Enterprise Desktop") {

		tmp := utils_strings.ExtractNumber(platform, " ")

		if tmp != "" {
			return "desktop-" + tmp
		}

	} else if strings.HasPrefix(platform, "SUSE Linux Enterprise Thin Client") {

		tmp := utils_strings.ExtractNumber(platform, " ")

		if tmp != "" {
			return "desktop-" + tmp
		}

	} else if strings.HasPrefix(platform, "SUSE Linux Enterprise Workstation Extension") || strings.HasPrefix(platform, "SUSE Linux Enterprise Workstation") {

		tmp := utils_strings.ExtractNumber(platform, " ")

		if tmp != "" {
			return "desktop-" + tmp
		}

	} else if strings.HasPrefix(platform, "SUSE Linux Enterprise High Performance Computing") {

		tmp := utils_strings.ExtractNumber(platform, " ")

		if tmp != "" {
			return "server-" + tmp
		}

	} else if strings.HasPrefix(platform, "SUSE Linux Enterprise Java") {

		tmp := utils_strings.ExtractNumber(platform, " ")

		if tmp != "" {
			return "server-" + tmp
		}

	} else if strings.HasPrefix(platform, "SUSE Linux Enterprise Live Patching") {

		tmp := utils_strings.ExtractNumber(platform, " ")

		if tmp != "" {
			return "server-" + tmp
		}

	} else if strings.HasPrefix(platform, "SUSE Linux Enterprise Module for Python 2") || strings.HasPrefix(platform, "SUSE Linux Enterprise Module for Python 3") {

		tmp := utils_strings.ExtractNumber(strings.TrimSpace(platform[41:]), " ")

		if tmp != "" {
			return "server-" + tmp
		}

	} else if strings.HasPrefix(platform, "SUSE Linux Enterprise Module for") {

		tmp := utils_strings.ExtractNumber(platform, " ")

		if tmp != "" {
			return "server-" + tmp
		}

	} else if strings.HasPrefix(platform, "SUSE Linux Enterprise Realtime") || strings.HasPrefix(platform, "SUSE Linux Enterprise Real Time") {

		tmp := utils_strings.ExtractNumber(platform, " ")

		if tmp != "" {
			return "server-" + tmp
		}

	} else if strings.HasPrefix(platform, "SUSE Linux Enterprise Server RT Solution") || strings.HasPrefix(platform, "SUSE Linux Enterprise Server") {

		tmp := utils_strings.ExtractNumber(platform, " ")

		if tmp != "" {
			return "server-" + tmp
		}

	} else if strings.HasPrefix(platform, "SUSE Linux Enterprise") {

		tmp := utils_strings.ExtractNumber(platform, " ")

		if tmp != "" {
			return "server-" + tmp
		}

	}

	return ""

}
