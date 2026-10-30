package manjaro

import "strings"

func ToVendor(platform string) string {

	if strings.Contains(strings.ToLower(platform), "arm") {
		return "archlinux-arm"
	}

	return "archlinux"

}
