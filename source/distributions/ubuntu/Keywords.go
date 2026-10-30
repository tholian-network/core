package ubuntu

var Keywords map[string]string = map[string]string{
	"ID":               "ubuntu",
	"NAME":             "Ubuntu",
	"VERSION_CODENAME": "%release%",
}

var KeywordsLegacy map[string]string = map[string]string{
	"ID":         "ubuntu",
	"NAME":       "Ubuntu",
	"VERSION_ID": "%version%",
}
