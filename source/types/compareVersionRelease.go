package types

import "slices"

func compareVersionRelease(a string, b string) int {

	// returns -1 if a is older than b
	// returns 1 if a is newer than b

	chunks_a := toVersionReleaseParts(a)
	chunks_b := toVersionReleaseParts(b)
	length_a := len(chunks_a)
	length_b := len(chunks_b)

	var check_a string
	var check_b string

	if length_a == 1 && length_b == 1 {

		check_a = chunks_a[0]
		check_b = chunks_b[0]

	} else if length_a == 2 && length_b == 2 {

		if isNumber(chunks_a[0]) && isNumber(chunks_b[0]) {

			if chunks_a[0] < chunks_b[0] {
				return -1
			} else if chunks_a[0] > chunks_b[0] {
				return 1
			} else {
				check_a = chunks_a[1]
				check_b = chunks_b[1]
			}

		} else if isNumber(chunks_a[1]) && isNumber(chunks_b[1]) {

			if chunks_a[1] < chunks_b[1] {
				return -1
			} else if chunks_a[1] > chunks_b[1] {
				return 1
			} else {
				check_a = chunks_a[0]
				check_b = chunks_b[0]
			}

		}

	} else if length_a == 3 && length_b == 3 {

		if isNumber(chunks_a[0]) && isNumber(chunks_b[0]) && isNumber(chunks_a[2]) && isNumber(chunks_b[2]) {

			if chunks_a[0] < chunks_b[0] {
				return -1
			} else if chunks_a[0] > chunks_b[0] {
				return 1
			} else {

				if chunks_a[1] == chunks_b[1] {

					if chunks_a[2] < chunks_b[2] {
						return -1
					} else if chunks_a[2] > chunks_b[2] {
						return 1
					} else {
						return 0
					}

				} else {
					check_a = chunks_a[1]
					check_b = chunks_b[1]
				}

			}

		}

	} else {

		if length_a == 0 && length_b > 0 {
			return -1
		} else if length_a > 0 && length_b == 0 {
			return 1
		}

	}

	if check_a != check_b {

		if slices.Contains(version_letters, check_a) && slices.Contains(version_letters, check_b) {

			index_a := slices.Index(version_letters, check_a)
			index_b := slices.Index(version_letters, check_b)

			if index_a < index_b {
				return -1
			} else if index_a > index_b {
				return 1
			}

		} else if slices.Contains(version_prereleases, check_a) && slices.Contains(version_prereleases, check_b) {

			index_a := slices.Index(version_prereleases, check_a)
			index_b := slices.Index(version_prereleases, check_b)

			if index_a < index_b {
				return -1
			} else if index_a > index_b {
				return 1
			}

		}

	}

	return 0

}
