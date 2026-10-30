package openwrt

import "strings"

func ToVersion(platform string) string {

	for release := range Releases {

		if platform == "OpenWrt "+release ||
			strings.HasPrefix(platform, "OpenWrt "+release+".") ||
			platform == "LEDE "+release ||
			strings.HasPrefix(platform, "LEDE "+release+".") ||
			platform == "OpenWrt/LEDE "+release ||
			strings.HasPrefix(platform, "OpenWrt/LEDE "+release+".") {

			return release

		} else if platform == release {
			return release
		}

	}

	return ""

}
