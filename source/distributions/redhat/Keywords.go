package redhat

var Keywords map[string]string = map[string]string{
	"ID":          "rhel",
	"NAME":        "Red Hat Enterprise Linux",
	"PLATFORM_ID": "platform:el%version%",
}

var KeywordsLegacy map[string]string = map[string]string{
	"ID":         "rhel",
	"NAME":       "Red Hat Enterprise Linux Server",
	"VERSION_ID": "%version%",
}
