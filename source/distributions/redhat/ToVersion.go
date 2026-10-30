package redhat

import "strconv"

func ToVersion(platform string) string {

	extracted_release := extractRelease(platform)

	for release, val := range Releases {

		version := strconv.Itoa(val)

		if release == extracted_release {
			return version
		} else if release == platform {
			return version
		}

	}

	return ""

}
