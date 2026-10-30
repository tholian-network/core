package rockylinux

import "strconv"

func ToRelease(platform string) string {

	for name, val := range Releases {

		version := strconv.Itoa(val)

		if platform == "Rocky Linux "+version ||
			platform == "Rocky Linux SIG Cloud "+version {

			return name

		}

	}

	return ""

}
