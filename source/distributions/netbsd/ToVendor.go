package netbsd

func ToVendor(platform string) string {

	for release, vendor := range Vendors {

		if platform == "NetBSD "+release ||
			platform == "NetBSD "+release+"-RELEASE" ||
			platform == release {

			return vendor

		}

	}

	return ""

}
