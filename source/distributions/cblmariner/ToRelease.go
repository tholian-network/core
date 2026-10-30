package cblmariner

import "strconv"

func ToRelease(platform string) string {

	for name, val := range Releases {

		version := strconv.FormatFloat(float64(val), 'f', 1, 32)

		if platform == "CBL Mariner "+version ||
			platform == "CBL-Mariner "+version {

			return name

		}

	}

	return ""

}
