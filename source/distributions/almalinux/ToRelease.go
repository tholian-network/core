package almalinux

import "strconv"

func ToRelease(platform string) string {

	for name, val := range Releases {

		version := strconv.Itoa(val)

		if platform == "Alma Linux "+version ||
			platform == "AlmaLinux "+version {

			return name

		}

	}

	return ""

}
