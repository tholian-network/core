package ubuntu

import "strconv"

func ToVendor(platform string) string {

	for release, vendor := range Vendors {

		tmp_val, tmp_ok := Releases[release]

		if tmp_ok == true {

			version := strconv.FormatFloat(float64(tmp_val), 'f', 2, 32)

			if platform == "Ubuntu "+version+" ESM" ||
				platform == "Ubuntu "+version+" LTS" ||
				platform == "Ubuntu "+version {

				return vendor

			}

		}

		if platform == release {
			return vendor
		}

	}

	return ""

}
