package amazonlinux

import "strconv"

func ToRelease(platform string) string {

	for name, val := range Releases {

		version := strconv.Itoa(val)

		if platform == "Amazon Linux "+version {
			return name
		}

	}

	return ""

}
