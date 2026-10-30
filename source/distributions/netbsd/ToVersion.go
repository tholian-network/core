package netbsd

import "strconv"

func ToVersion(platform string) string {

	for _, val := range Releases {

		version1 := strconv.FormatFloat(float64(val), 'f', 1, 32)
		version2 := strconv.FormatFloat(float64(val), 'f', 2, 32)

		if platform == "NetBSD "+version2 ||
			platform == "NetBSD "+version2+"-RELEASE" ||
			platform == version2 {

			return version2

		} else if platform == "NetBSD "+version1 ||
			platform == "NetBSD "+version1+"-RELEASE" ||
			platform == version1 {

			return version1

		}

	}

	return ""

}
