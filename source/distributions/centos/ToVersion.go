package centos

import "strconv"

func ToVersion(platform string) string {

	for release, val := range Releases {

		version := strconv.Itoa(val)

		if platform == "Cent OS "+version ||
			platform == "CentOS "+version {

			return version

		} else if platform == release {

			return version

		}

	}

	return ""

}
