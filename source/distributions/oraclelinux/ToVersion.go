package oraclelinux

import "strconv"

func ToVersion(platform string) string {

	for release, val := range Releases {

		version := strconv.Itoa(val)

		if platform == "Oracle Linux "+version {
			return version
		} else if platform == release {
			return version
		}

	}

	return ""

}
