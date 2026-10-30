package linuxmint

import "strconv"

func ToVendor(platform string) string {

	for release, vendor := range Vendors {

		tmp_val, tmp_ok := Releases[release]

		if tmp_ok == true {

			version := strconv.FormatFloat(float64(tmp_val), 'f', 1, 32)

			if platform == "Linux Mint "+version {
				return vendor
			}

		}

		if platform == release {
			return vendor
		}

	}

	return ""

}
