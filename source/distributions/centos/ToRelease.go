package centos

import "strconv"

func ToRelease(platform string) string {

	for name, val := range Releases {

		version := strconv.Itoa(val)

		if platform == "Cent OS "+version ||
			platform == "CentOS "+version {

			return name

		}

	}

	return ""

}
