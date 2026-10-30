package linuxmint

import "strconv"

func ToRelease(platform string) string {

	for name, val := range Releases {

		version := strconv.FormatFloat(float64(val), 'f', 1, 32)

		if platform == "Linux Mint "+version {
			return name
		}

	}

	return ""

}
