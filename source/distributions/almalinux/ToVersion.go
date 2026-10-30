package almalinux

import "strconv"

func ToVersion(platform string) string {

	for release, val := range Releases {

		version := strconv.Itoa(val)

		if platform == "AlmaLinux "+version ||
			platform == "Alma Linux "+version {

			return version

		} else if platform == release {

			return version

		}

	}

	return ""

}
