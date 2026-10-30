package netbsd

import "strconv"

func ToRelease(platform string) string {

	for name, val := range Releases {

		version1 := strconv.FormatFloat(float64(val), 'f', 1, 32)
		version2 := strconv.FormatFloat(float64(val), 'f', 2, 32)

		if platform == "NetBSD "+version1 ||
			platform == "NetBSD "+version1+"-RELEASE" ||
			platform == "NetBSD "+version2 ||
			platform == "NetBSD "+version2+"-RELEASE" {

			return name

		}

	}

	return ""

}
