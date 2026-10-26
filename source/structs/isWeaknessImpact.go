package structs

var weakness_impacts = []string{
	"Alter Execution Logic",
	"Bypass Protection Mechanism",
	"DoS: Amplification",
	"DoS: Crash, Exit, or Restart",
	"DoS: Instability",
	"DoS: Resource Consumption (CPU)",
	"DoS: Resource Consumption (Memory)",
	"DoS: Resource Consumption (Other)",
	"Execute Unauthorized Code or Commands",
	"Gain Privileges or Assume Identity",
	"Hide Activities",
	"Modify Application Data",
	"Modify Files or Directories",
	"Modify Memory",
	"Quality Degradation",
	"Read Application Data",
	"Read Files or Directories",
	"Read Memory",
	"Reduce Maintainability",
	"Reduce Performance",
	"Reduce Reliability",
	"Unexpected State",
	"Varies by Context",
}

func isWeaknessImpact(value string) bool {

	for i := 0; i < len(weakness_impacts); i++ {

		if weakness_impacts[i] == value {
			return true
		}

	}

	return false

}
