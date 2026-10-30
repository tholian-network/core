package fedora

var Keywords map[string]string = map[string]string{
	"ID": "fedora",
	// "NAME": "Fedora",
	// "NAME": "Fedora Linux",
	// "VERSION": "28 (Cloud Edition)",
	// "VERSION": "32 (Container Image)",
	// "VERSION": "38 (Workstation Edition)",
	"VERSION_ID":  "%version%",
	"PLATFORM_ID": "platform:f%version%",
}
