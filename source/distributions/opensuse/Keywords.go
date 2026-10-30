package opensuse

var KeywordsLeap map[string]string = map[string]string{
	"NAME":       "openSUSE Leap",
	"ID":         "opensuse-leap",
	"ID_LIKE":    "suse opensuse",
	"VERSION_ID": "%version%",
}

var KeywordsLeapLegacy map[string]string = map[string]string{
	"NAME":       "openSUSE Leap",
	"ID":         "opensuse",
	"ID_LIKE":    "suse",
	"VERSION_ID": "%version%",
}

var KeywordsTumbleweed map[string]string = map[string]string{
	"NAME":    "openSUSE Tumbleweed",
	"ID":      "opensuse-tumbleweed",
	"ID_LIKE": "opensuse suse",
}

var Keywords map[string]string = map[string]string{
	"ID":         "opensuse",
	"ID_LIKE":    "suse",
	"VERSION_ID": "%version%",
}
