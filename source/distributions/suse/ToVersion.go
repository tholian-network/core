package suse

import "strconv"

func ToVersion(platform string) string {

	release := extractRelease(platform)

	for name, val := range Releases {

		version := strconv.Itoa(val)

		if name == release {
			return version
		} else if name == platform {
			return version
		}

	}

	return ""

}
