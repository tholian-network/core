package centos

var Keywords map[string]string = map[string]string{
	"ID": "centos",
	// "NAME": "CentOS Linux",
	// "NAME": "CentOS Stream",
	"VERSION_ID":  "%version%",
	"PLATFORM_ID": "platform:el%version%",
}

var KeywordsLegacy map[string]string = map[string]string{
	"ID":         "centos",
	"VERSION_ID": "%version%",
}
