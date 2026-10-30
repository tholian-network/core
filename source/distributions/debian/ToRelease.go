package debian

import "strconv"

func ToRelease(platform string) string {

	for name, version := range Releases {

		if platform == "Debian GNU/Linux "+strconv.Itoa(version) {
			return name
		}

	}

	return ""

}
