package fedora

import "strconv"

func ToVendor(platform string) string {

	for release, vendor := range Vendors {

		tmp_val, tmp_ok := Releases[release]

		if tmp_ok == true {

			version := strconv.Itoa(tmp_val)

			if platform == "Fedora "+version ||
				platform == "Fedora "+version+" (Cloud Edition)" ||
				platform == "Fedora "+version+" (Container Image)" ||
				platform == "Fedora "+version+" (Server Edition)" ||
				platform == "Fedora "+version+" (Workstation Edition)" {

				return vendor

			} else if platform == "Fedora Linux "+version ||
				platform == "Fedora Linux "+version+" (Cloud Edition)" ||
				platform == "Fedora Linux "+version+" (Container Image)" ||
				platform == "Fedora Linux "+version+" (Server Edition)" ||
				platform == "Fedora Linux "+version+" (Workstation Edition)" {

				return vendor

			}

		}

		if platform == release {
			return vendor
		}

	}

	return ""

}
