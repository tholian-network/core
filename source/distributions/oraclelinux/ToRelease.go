package oraclelinux

import "strconv"

func ToRelease(platform string) string {

	for name, version := range Releases {

		if platform == "Oracle Linux "+strconv.Itoa(version) {
			return name
		}

	}

	return ""

}
