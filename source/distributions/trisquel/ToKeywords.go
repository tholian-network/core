package trisquel

import "strings"

func ToKeywords(release string, ubuntu_release string) map[string]string {

	result := make(map[string]string)

	for key, val := range Keywords {

		if strings.Contains(val, "{release}") {

			if release != "" {
				result[key] = strings.ReplaceAll(val, "{release}", release)
			}

		} else if strings.Contains(val, "{ubuntu_release}") {

			if ubuntu_release != "" {
				result[key] = strings.ReplaceAll(val, "{ubuntu_release}", ubuntu_release)
			}

		}

	}

	return result

}
