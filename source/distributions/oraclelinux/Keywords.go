package oraclelinux

var Keywords map[string]string = map[string]string{
	"ID":          "ol",
	"NAME":        "Oracle Linux Server",
	"PLATFORM_ID": "platform:el%version%",
}

var KeywordsLegacy map[string]string = map[string]string{
	"ID":         "ol",
	"NAME":       "Oracle Linux Server",
	"VERSION_ID": "%version%",
}
