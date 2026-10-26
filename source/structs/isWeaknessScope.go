package structs

var weakness_scopes = []string{
	"Access Control",
	"Accountability",
	"Authentication",
	"Authorization",
	"Availability",
	"Confidentiality",
	"Integrity",
	"Non-Repudiation",
}

func isWeaknessScope(value string) bool {

	for s := 0; s < len(weakness_scopes); s++ {

		if weakness_scopes[s] == value {
			return true
		}

	}

	return false

}

