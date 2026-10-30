package redhat

func ToVendor(platform string) string {

	release := extractRelease(platform)

	for name, vendor := range Vendors {

		if name == release {
			return vendor
		} else if name == platform {
			return vendor
		}

	}

	return ""

}
