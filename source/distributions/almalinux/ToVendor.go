package almalinux

import "strconv"

func ToVendor(platform string) string {

	for release, vendor := range Vendors {

		tmp_val, tmp_ok := Releases[release]

		if tmp_ok == true {

			version := strconv.Itoa(tmp_val)

			if platform == "Alma Linux "+version ||
				platform == "AlmaLinux "+version {

				return vendor

			}

		}

		if platform == release {
			return vendor
		}

	}

	return ""

}
