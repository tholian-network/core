package capabilities

import "strings"
import "testing"

func TestCapabilitiesAreProjectQualified(t *testing.T) {

	supported := []Capability{
		CapabilitySystemRead,
		CapabilityPackagesRead,
		CapabilityPackagesUpdate,
		CapabilityProgramsRead,
		CapabilityServicesRead,
		CapabilityUsersRead,
		CapabilityUpdatesRead,
		CapabilityIncidentsRead,
		CapabilityFirewallRead,
		CapabilityFirewallWrite,
	}

	reserved := []Capability{
		CapabilityProgramsUpdate,
		CapabilityServicesUpdate,
		CapabilityUsersUpdate,
		CapabilityEvidenceRead,
		CapabilityEvidenceWrite,
		CapabilityShellExecute,
	}

	for _, capability := range supported {

		if strings.HasPrefix(capability.String(), "tholian-") == false {
			t.Fatalf("capability %q is not project-qualified", capability)
		}

		if IsSupportedCapability(capability) == false {
			t.Fatalf("capability %q is not supported", capability)
		}

	}

	for _, capability := range reserved {

		if strings.HasPrefix(capability.String(), "tholian-") == false {
			t.Fatalf("capability %q is not project-qualified", capability)
		}

		if IsReservedCapability(capability) == false {
			t.Fatalf("capability %q is not reserved", capability)
		}

	}

}
