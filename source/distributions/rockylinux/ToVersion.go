package rockylinux

import "strconv"

func ToVersion(platform string) string {

	for release, val := range Releases {

		version := strconv.Itoa(val)

		if platform == "Rocky Linux "+version ||
			platform == "Rocky Linux SIG Cloud "+version {

			return version

		} else if platform == release {

			return version

		}

	}

	return ""

}
