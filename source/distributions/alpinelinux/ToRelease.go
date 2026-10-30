package alpinelinux

import "strconv"

func ToRelease(platform string) string {

	for name, val := range Releases {

		version := strconv.FormatFloat(float64(val), 'f', 2, 32)

		if platform == "Alpine "+version ||
			platform == "Alpine Linux "+version {

			return name

		} else if platform == "alpine-"+version+"-main" ||
			platform == "alpine-"+version+"-community" {

			return name

		}

	}

	return ""

}
