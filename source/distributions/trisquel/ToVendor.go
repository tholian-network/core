package trisquel

import "strconv"
import "strings"

func ToVendor(platform string) string {

	for release, vendor := range Vendors {

		tmp_val, tmp_ok := Releases[release]

		if tmp_ok == true {

			version := strconv.FormatFloat(float64(tmp_val), 'f', 1, 32)

			// matches 11.0 and 11.0.1
			if strings.HasPrefix(platform, "Trisquel GNU/Linux ("+version) {
				return vendor
			}

		}

		if platform == release {
			return vendor
		}

	}

	return ""

}
