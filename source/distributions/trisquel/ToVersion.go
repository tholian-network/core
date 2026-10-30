package trisquel

import "strconv"
import "strings"

func ToVersion(platform string) string {

	for release, val := range Releases {

		version := strconv.FormatFloat(float64(val), 'f', 1, 32)

		// matches 11.0 and 11.0.1
		if strings.HasPrefix(platform, "Trisquel GNU/Linux ("+version) {
			return version
		} else if platform == release {
			return version
		}

	}

	return ""

}
