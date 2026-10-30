package distributions_test

import (
	"strings"
	"testing"
)

func parseOSRelease(t *testing.T, data []byte) map[string]string {

	t.Helper()

	fields := make(map[string]string)

	for _, line := range strings.Split(string(data), "\n") {

		line = strings.TrimSpace(line)

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		key, value, ok := strings.Cut(line, "=")

		if ok == true {

			value = strings.TrimSpace(value)

			if len(value) >= 2 {

				if value[0] == '"' && value[len(value)-1] == '"' {
					value = value[1 : len(value)-1]
				} else if value[0] == '\'' && value[len(value)-1] == '\'' {
					value = value[1 : len(value)-1]
				}

			}

			fields[strings.TrimSpace(key)] = value

		}

	}

	return fields

}

func assertKeywords(t *testing.T, keywords map[string]string, data []byte, variables map[string]string) {

	t.Helper()

	fields := parseOSRelease(t, data)

	for key, expected := range keywords {

		contains := false

		for placeholder, value := range variables {

			if strings.Contains(expected, placeholder) {
				contains = true
				expected = strings.ReplaceAll(expected, placeholder, value)
			}

		}

		actual, ok := fields[key]

		if ok == false {

			t.Errorf("missing %s in os-release", key)
			continue

		}

		if contains == true {

			if strings.Contains(actual, expected) == false {
				t.Errorf("%s = %q, expected to contain %q", key, actual, expected)
			}

		} else if actual != expected {
			t.Errorf("%s = %q, expected %q", key, actual, expected)
		}

	}

}
