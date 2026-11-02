package permissions

import "regexp"
import "testing"

var methodPattern = regexp.MustCompile(`^tholian-[a-z0-9-]+\.[A-Z][A-Za-z0-9]*$`)

func TestMethodsAreUniqueAndWellFormed(t *testing.T) {

	seen := make(map[Method]bool)

	for _, method := range AllMethods() {

		if methodPattern.MatchString(method.String()) == false {
			t.Fatalf("method %q does not match the project-qualified pattern", method)
		}

		if seen[method] == true {
			t.Fatalf("method %q is defined more than once", method)
		}

		seen[method] = true

		if method.Project() == "" || method.Action() == "" {
			t.Fatalf("method %q cannot be split into project and action", method)
		}

	}

}

func TestMethodWireValues(t *testing.T) {

	cases := map[Method]string{
		MethodFirewallCheck:          "tholian-firewall.Check",
		MethodFirewallForbidNetworks: "tholian-firewall.ForbidNetworks",
		MethodScoutCollectPackages:   "tholian-scout.CollectPackages",
		MethodScoutSearchPackages:    "tholian-scout.SearchPackages",
		MethodGuardReport:            "tholian-guard.Report",
		MethodWatchtowerCall:         "tholian-watchtower.Call",
	}

	for method, expected := range cases {
		if method.String() != expected {
			t.Fatalf("method %q has wire value %q, expected %q", method.Project()+"."+method.Action(), method, expected)
		}
	}

}

func TestParseMethod(t *testing.T) {
	parsed, ok := ParseMethod("tholian-firewall.ForbidNetworks")

	if ok == false {
		t.Fatal("known method was rejected")
	}

	if parsed != MethodFirewallForbidNetworks {
		t.Fatalf("unexpected method: %q", parsed)
	}

	if parsed.Project() != "tholian-firewall" {
		t.Fatalf("unexpected project: %q", parsed.Project())
	}

	if parsed.Action() != "ForbidNetworks" {
		t.Fatalf("unexpected action: %q", parsed.Action())
	}

	if _, ok := ParseMethod("firewall.forbid"); ok == true {
		t.Fatal("legacy method was accepted")
	}

}

func TestActionMethodsHaveCapabilities(t *testing.T) {

	reserved := map[Method]bool{
		MethodGuardCollectVulnerabilities: true,
		MethodGuardReadEvidence:           true,
		MethodGuardReadFile:               true,
		MethodGuardWriteFile:              true,
		MethodGuardExecuteShell:           true,
	}

	for _, method := range AllMethods() {

		project := method.Project()

		if project != "tholian-firewall" && project != "tholian-scout" && method != MethodGuardReport {
			continue
		}

		if reserved[method] == true {
			continue
		}

		if _, ok := MethodCapability(method); ok == false {
			t.Fatalf("action method %q has no capability mapping", method)
		}

	}

}
