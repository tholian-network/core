package cblmariner

import "strconv"

func ToVendor(platform string) string {

	for release, vendor := range Vendors {

		tmp_val, tmp_ok := Releases[release]

		if tmp_ok == true {

			version := strconv.FormatFloat(float64(tmp_val), 'f', 1, 32)

			if platform == "CBL Mariner "+version ||
				platform == "CBL-Mariner "+version {

				return vendor

			}

		}

		if platform == release {
			return vendor
		}

	}

	return ""

}
