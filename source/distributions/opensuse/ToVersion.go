package opensuse

import "strconv"

func ToVersion(platform string) string {

	extracted_release := extractRelease(platform)

	for release, val := range Releases {

		version := strconv.FormatFloat(float64(val), 'f', 1, 32)

		if release == extracted_release {

			if version == "0.0" {
				return ""
			} else {
				return version
			}

		} else if release == platform {

			if version == "0.0" {
				return ""
			} else {
				return version
			}

		}

	}

	return ""

}
