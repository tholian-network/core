package types

func toVersionReleaseParts(value string) []string {

	result := make([]string, 0)

	chunk := ""
	mode := "unknown"

	for _, chr := range value {

		character := string(chr)

		if mode == "unknown" {

			if character >= "0" && character <= "9" {
				chunk += character
				mode = "number"
			} else if character >= "a" && character <= "z" {
				chunk += character
				mode = "alphabet"
			}

		} else if mode == "alphabet" {

			if character >= "0" && character <= "9" {
				result = append(result, chunk)
				chunk = character
				mode = "number"
			} else if character >= "a" && character <= "z" {
				chunk += character
			}

		} else if mode == "number" {

			if character >= "0" && character <= "9" {
				chunk += character
			} else if character >= "a" && character <= "z" {
				result = append(result, chunk)
				chunk = character
				mode = "alphabet"
			}

		}

	}

	if chunk != "" {
		result = append(result, chunk)
	}

	return result

}
