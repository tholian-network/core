package matchers

import "strings"

func parseOffsetCondition(value string) (string, string) {

	result_name := ""
	result_offset := ""

	if strings.Contains(value, ">=") {

		// "Europe/* >= +01:00"

		offset := strings.TrimSpace(value[strings.Index(value, ">=")+2:])

		if len(offset) == 6 {

			operator := string(offset[0])

			if operator == "+" || operator == "-" {
				result_name = strings.TrimSpace(value[0:strings.Index(value, ">=")])
				result_offset = ">= " + offset
			}

		}

	} else if strings.Contains(value, "<=") {

		// "Europe/* <= +02:00"

		offset := strings.TrimSpace(value[strings.Index(value, "<=")+2:])

		if len(offset) == 6 {

			operator := string(offset[0])

			if operator == "+" || operator == "-" {
				result_name = strings.TrimSpace(value[0:strings.Index(value, "<=")])
				result_offset = "<= " + offset
			}

		}

	} else if strings.Contains(value, ">") {

		// "Europe/* > +01:00"

		offset := strings.TrimSpace(value[strings.Index(value, ">")+1:])

		if len(offset) == 6 {

			operator := string(offset[0])

			if operator == "+" || operator == "-" {
				result_name = strings.TrimSpace(value[0:strings.Index(value, ">")])
				result_offset = "> " + offset
			}

		}

	} else if strings.Contains(value, "<") {

		// "Europe/* < +01:00"

		offset := strings.TrimSpace(value[strings.Index(value, "<")+1:])

		if len(offset) == 6 {

			operator := string(offset[0])

			if operator == "+" || operator == "-" {
				result_name = strings.TrimSpace(value[0:strings.Index(value, "<")])
				result_offset = "< " + offset
			}

		}

	} else if strings.Contains(value, "=") {

		// "Europe/* = +01:00"

		offset := strings.TrimSpace(value[strings.Index(value, "=")+1:])

		if len(offset) == 6 {

			operator := string(offset[0])

			if operator == "+" || operator == "-" {
				result_name = strings.TrimSpace(value[0:strings.Index(value, "=")])
				result_offset = "= " + offset
			}

		}

	} else if strings.Contains(value, " ") {

		// "Europe/* +01:00"

		offset := strings.TrimSpace(value[strings.Index(value, " ")+1:])

		if len(offset) == 6 {

			operator := string(offset[0])

			if operator == "+" || operator == "-" {
				result_name = strings.TrimSpace(value[0:strings.Index(value, " ")])
				result_offset = "= " + offset
			}

		}

	} else {

		// "Europe/*"

		result_name = value
		result_offset = "any"

	}

	return result_name, result_offset

}
