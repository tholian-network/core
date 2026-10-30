package freebsd

func ToVendor(platform string) string {

	for release, vendor := range Vendors {

		if platform == "FreeBSD "+release ||
			platform == "FreeBSD "+release+"-RELEASE" ||
			platform == "FreeBSD "+release+"-STABLE" ||
			platform == release {

			return vendor

		}

	}

	return ""

}
