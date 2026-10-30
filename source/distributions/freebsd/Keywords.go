package freebsd

var Keywords map[string]string = map[string]string{
	"ID":   "freebsd",
	"NAME": "FreeBSD",
	// Versions can be 4.1 and 4.10 which are the same float32 values
	"VERSION_ID": "%release%",
}
