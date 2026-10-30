package debian

var Keywords map[string]string = map[string]string{
	"ID":               "debian",
	"NAME":             "Debian GNU/Linux",
	"VERSION_CODENAME": "%release%",
}

var KeywordsLegacy map[string]string = map[string]string{
	"ID":         "debian",
	"NAME":       "Debian GNU/Linux",
	"VERSION_ID": "%version%",
}
