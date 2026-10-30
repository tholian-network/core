package opensuse

import "strings"

func extractRelease(platform string) string {

	if strings.HasPrefix(platform, "openSUSE Leap") && strings.HasSuffix(platform, "NonFree") {

		return "leap-" + strings.TrimSpace(platform[13:len(platform)-7])

	} else if strings.HasPrefix(platform, "openSUSE Leap") {

		return "leap-" + strings.TrimSpace(platform[13:])

	} else if strings.HasPrefix(platform, "openSUSE Tumbleweed") {

		return "tumbleweed"

	} else if strings.HasPrefix(platform, "openSUSE") && strings.HasSuffix(platform, "NonFree") {

		return strings.TrimSpace(platform[8 : len(platform)-7])

	} else if strings.HasPrefix(platform, "openSUSE") && strings.HasSuffix(platform, "Update") {

		return strings.TrimSpace(platform[8 : len(platform)-6])

	} else if strings.HasPrefix(platform, "openSUSE") {

		return strings.TrimSpace(platform[8:])

	}

	return ""

}
