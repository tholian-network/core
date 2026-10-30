package suse

var KeywordsDesktop map[string]string = map[string]string{
	"NAME": "SLED",
	"ID":   "sled",
	// Versions are matched with optional trailing (.[0-9])
	"VERSION_ID": "%version%",
}

var KeywordsServer map[string]string = map[string]string{
	"NAME": "SLES",
	"ID":   "sles",
	// Versions are matched with optional trailing (.[0-9])
	"VERSION_ID": "%version%",
}
