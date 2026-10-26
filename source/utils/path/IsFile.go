package path

import "strings"

// From: golang.org/x/module/module.go
// See: https://learn.microsoft.com/en-us/windows/win32/fileio/naming-a-file

var bad_windows_names = []string{
	"aux",
	"con",
	"com0",
	"com1",
	"com2",
	"com3",
	"com4",
	"com5",
	"com6",
	"com7",
	"com8",
	"com9",
	"lpt0",
	"lpt1",
	"lpt2",
	"lpt3",
	"lpt4",
	"lpt5",
	"lpt6",
	"lpt7",
	"lpt8",
	"lpt9",
	"nul",
	"prn",
}

func IsFile(value string) bool {

	if strings.Contains(value, "/") {
		value = value[strings.LastIndex(value, "/")+1:]
	}

	result := true

	for v := 0; v < len(value); v++ {

		character := string(value[v])

		if character >= "0" && character <= "9" {
			continue
		} else if character >= "A" && character <= "Z" {
			continue
		} else if character >= "a" && character <= "z" {
			continue
		} else if character == ":" || character == ";" {
			continue
		} else if character == "." || character == "," {
			continue
		} else if character == "+" || character == "-" {
			continue
		} else if character == "_" {
			continue
		} else {
			result = false
			break
		}

	}

	if result == true {

		if strings.Contains(value, ".") {

			basename := strings.ToLower(value[0:strings.Index(value, ".")])

			for b := 0; b < len(bad_windows_names); b++ {

				if bad_windows_names[b] == basename {
					result = false
					break
				}

			}

		} else {

			for b := 0; b < len(bad_windows_names); b++ {

				if bad_windows_names[b] == value {
					result = false
					break
				}

			}

		}

	}

	return result

}
