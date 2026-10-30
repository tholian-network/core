package openwrt

import "strings"

func ToVendor(platform string) string {

	for release, vendor := range Vendors {

		if platform == "OpenWrt "+release ||
			strings.HasPrefix(platform, "OpenWrt "+release+".") ||
			platform == "LEDE "+release ||
			strings.HasPrefix(platform, "LEDE "+release+".") ||
			platform == "OpenWrt/LEDE "+release ||
			strings.HasPrefix(platform, "OpenWrt/LEDE "+release+".") {

			return vendor

		} else if platform == release {
			return vendor
		}

	}

	return ""

}
