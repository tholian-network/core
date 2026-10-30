package ubuntu

import "strconv"

func ToVersion(platform string) string {

	for release, val := range Releases {

		version := strconv.FormatFloat(float64(val), 'f', 2, 32)

		if platform == "Ubuntu "+version+" ESM" ||
			platform == "Ubuntu "+version+" LTS" ||
			platform == "Ubuntu "+version {

			return version

		} else if platform == release {

			return version

		}

	}

	return ""

}
