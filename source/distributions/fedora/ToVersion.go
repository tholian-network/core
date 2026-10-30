package fedora

import "strconv"

func ToVersion(platform string) string {

	for release, val := range Releases {

		version := strconv.Itoa(val)

		if platform == "Fedora "+version ||
			platform == "Fedora "+version+" (Cloud Edition)" ||
			platform == "Fedora "+version+" (Container Image)" ||
			platform == "Fedora "+version+" (Server Edition)" ||
			platform == "Fedora "+version+" (Workstation Edition)" {

			return version

		} else if platform == "Fedora Linux "+version ||
			platform == "Fedora Linux "+version+" (Cloud Edition)" ||
			platform == "Fedora Linux "+version+" (Container Image)" ||
			platform == "Fedora Linux "+version+" (Server Edition)" ||
			platform == "Fedora Linux "+version+" (Workstation Edition)" {

			return version

		} else if platform == release {

			return version

		}

	}

	return ""

}
