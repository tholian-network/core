package strings

import "strings"

func ExtractNumber(value string, delimiter string) string {

	result := ""
	tmp := strings.Split(value, delimiter)

	for t := 0; t < len(tmp); t++ {

		chunk := tmp[t]

		if strings.Contains(chunk, "-") {
			chunk = chunk[0:strings.Index(chunk, "-")]
		}

		is_number := true

		for c := 0; c < len(chunk); c++ {

			chr := string(chunk[c])

			if chr >= "0" && chr <= "9" {
				// Do Nothing
			} else {
				is_number = false
				break
			}

		}

		if is_number {
			result = chunk
			break
		}

	}

	return result

}
