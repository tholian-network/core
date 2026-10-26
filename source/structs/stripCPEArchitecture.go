package structs

import "strings"

func stripCPEArchitecture(value string) string {

	if strings.HasPrefix(value, "x86_64_") {
		value = strings.TrimSpace(value[7:])
	} else if strings.HasPrefix(value, "x86_") {
		value = strings.TrimSpace(value[4:])
	} else if strings.HasSuffix(value, "_x86_64") {
		value = strings.TrimSpace(value[0 : len(value)-7])
	} else if strings.HasSuffix(value, "_x86") {
		value = strings.TrimSpace(value[0 : len(value)-4])
	} else if strings.HasSuffix(value, "_32-bit") {
		value = strings.TrimSpace(value[0 : len(value)-7])
	} else if strings.HasSuffix(value, "_64-bit") {
		value = strings.TrimSpace(value[0 : len(value)-7])
	} else if value == "x86_64" {
		value = ""
	} else if value == "x86" {
		value = ""
	} else if value == "32-bit" {
		value = ""
	} else if value == "64-bit" {
		value = ""
	} else if value == "*" {
		value = ""
	} else if value == "-" {
		value = ""
	}

	return value

}
