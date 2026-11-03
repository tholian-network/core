package capabilities

import "strings"
import "testing"

func TestCapabilityString(t *testing.T) {

	if FirewallRead.String() != "tholian-firewall.read" {
		t.Fatalf("unexpected capability string: %q", FirewallRead.String())
	}

	if GuardWrite.String() != "tholian-guard.write" {
		t.Fatalf("unexpected capability string: %q", GuardWrite.String())
	}

	if ScoutInstall.String() != "tholian-scout.install" {
		t.Fatalf("unexpected capability string: %q", ScoutInstall.String())
	}

}

func TestCapabilitiesAreProjectQualified(t *testing.T) {

	expected := map[Capability]string{
		FirewallInstall: "tholian-firewall",
		FirewallRead:    "tholian-firewall",
		FirewallWrite:   "tholian-firewall",
		GuardInstall:    "tholian-guard",
		GuardRead:       "tholian-guard",
		GuardWrite:      "tholian-guard",
		ScoutInstall:    "tholian-scout",
		ScoutRead:       "tholian-scout",
		ScoutWrite:      "tholian-scout",
	}

	seen := make(map[Capability]bool)

	for capability, project := range expected {

		value := capability.String()

		if strings.HasPrefix(value, "tholian-") == false {
			t.Fatalf("capability %q is not project-qualified", capability)
		}

		actual, action, ok := strings.Cut(value, ".")

		if ok == false || action == "" {
			t.Fatalf("capability %q cannot be split into project and action", capability)
		}

		if actual != project {
			t.Fatalf("capability %q belongs to project %q, expected %q", capability, actual, project)
		}

		if seen[capability] == true {
			t.Fatalf("capability %q is defined more than once", capability)
		}

		seen[capability] = true

	}

}

func TestCapabilitiesAreUnique(t *testing.T) {

	all := []Capability{
		FirewallInstall,
		FirewallRead,
		FirewallWrite,
		GuardInstall,
		GuardRead,
		GuardWrite,
		ScoutInstall,
		ScoutRead,
		ScoutWrite,
	}

	seen := make(map[string]bool)

	for _, capability := range all {

		value := capability.String()

		if seen[value] == true {
			t.Fatalf("capability value %q is defined more than once", value)
		}

		seen[value] = true

	}

}
