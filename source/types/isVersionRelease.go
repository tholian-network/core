package types

import "slices"

func isVersionRelease(value string) bool {

	var result bool

	chunks := toVersionReleaseParts(value)

	if len(chunks) == 1 {

		if slices.Contains(version_letters, chunks[0]) {
			result = true
		}

		if result == false {

			if slices.Contains(version_prereleases, chunks[0]) {
				result = true
			}

		}

	} else if len(chunks) == 2 {

		if isNumber(chunks[0]) {

			if result == false {

				if slices.Contains(version_letters, chunks[1]) {
					result = true
				}

			}

			if result == false {

				if slices.Contains(version_prereleases, chunks[1]) {
					result = true
				}

			}

		} else if isNumber(chunks[1]) {

			if chunks[0] == "rc" || chunks[0] == "r" || chunks[0] == "pre" {
				result = true
			}

			if result == false {

				if slices.Contains(version_letters, chunks[0]) {
					result = true
				}

			}

			if result == false {

				if slices.Contains(version_prereleases, chunks[0]) {
					result = true
				}

			}

		}

	} else if len(chunks) == 3 {

		if isNumber(chunks[0]) && isNumber(chunks[2]) {

			if chunks[1] == "rc" || chunks[1] == "r" || chunks[1] == "pre" {
				result = true
			}

			if result == false {

				if slices.Contains(version_letters, chunks[1]) {
					result = true
				}

			}

			if result == false {

				if slices.Contains(version_prereleases, chunks[1]) {
					result = true
				}

			}

		}

	}

	return result

}
