package alpinelinux

import "strconv"

func ToVendor(platform string) string {

	for release, vendor := range Vendors {

		tmp_val, tmp_ok := Releases[release]

		if tmp_ok == true {

			version := strconv.FormatFloat(float64(tmp_val), 'f', 2, 32)

			if platform == "Alpine "+version ||
				platform == "Alpine Linux "+version {

				return vendor

			} else if platform == "alpine-"+version+"-main" ||
				platform == "alpine-"+version+"-community" {

				return vendor

			}

		}

		if platform == release {
			return vendor
		}

	}

	return ""

}
