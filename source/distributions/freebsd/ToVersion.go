package freebsd

import "strconv"

func ToVersion(platform string) string {

	for _, val := range Releases {

		version1 := strconv.FormatFloat(float64(val), 'f', 1, 32)
		version2 := strconv.FormatFloat(float64(val), 'f', 2, 32)

		if platform == "FreeBSD "+version2 ||
			platform == "FreeBSD "+version2+"-RELEASE" ||
			platform == "FreeBSD "+version2+"-STABLE" ||
			platform == version2 {

			return version2

		} else if platform == "FreeBSD "+version1 ||
			platform == "FreeBSD "+version1+"-RELEASE" ||
			platform == "FreeBSD "+version1+"-STABLE" ||
			platform == version1 {

			return version1

		}

	}

	return ""

}
