package ubuntu

import "strconv"

func ToRelease(platform string) string {

	for name, val := range Releases {

		version := strconv.FormatFloat(float64(val), 'f', 2, 32)

		if platform == "Ubuntu "+version+" ESM" ||
			platform == "Ubuntu "+version+" LTS" ||
			platform == "Ubuntu "+version {

			return name

		}

	}

	return ""

}
