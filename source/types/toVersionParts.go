package types

import "strings"

func toVersionParts(value string) []string {

	result := make([]string, 0)
	section := ""

	if strings.HasPrefix(value, "v") {
		value = value[1:]
	}

	for _, chr := range value {

		character := string(chr)

		if character == "~" || character == "+" || character == "-" || character == "_" {

			if section != "" {
				result = append(result, section)
				section = ""
			}

		} else if character == "." {

			if section != "" {
				result = append(result, section)
				section = ""
			}

		} else {
			section += character
		}

	}

	if section != "" {
		result = append(result, section)
	}

	return result

}
