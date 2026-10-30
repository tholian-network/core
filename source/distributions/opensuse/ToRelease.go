package opensuse

func ToRelease(platform string) string {

	release := extractRelease(platform)

	for name := range Releases {

		if name == release {
			return name
		}

	}

	return ""

}
