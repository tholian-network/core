package ubuntu

import "strings"

func ToKeywords(release string, version string, keywords map[string]string) map[string]string {

	result := make(map[string]string)

	for key, val := range keywords {

		if strings.Contains(val, "{release}") {

			if release != "" {
				result[key] = strings.ReplaceAll(val, "{release}", release)
			}

		} else if strings.Contains(val, "{version}") {

			if version != "" {
				result[key] = strings.ReplaceAll(val, "{version}", version)
			}

		}

	}

	return result

}
