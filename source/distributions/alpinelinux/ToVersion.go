package alpinelinux

import "strconv"

func ToVersion(platform string) string {

	for release, val := range Releases {

		version := strconv.FormatFloat(float64(val), 'f', 2, 32)

		if platform == "Alpine "+version ||
			platform == "Alpine Linux "+version {

			return version

		} else if platform == "alpine-"+version+"-main" ||
			platform == "alpine-"+version+"-community" {

			return version

		} else if platform == release {

			return version

		}

	}

	return ""

}
