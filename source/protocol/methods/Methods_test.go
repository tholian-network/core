package methods

import "testing"

func TestMethodWireValues(t *testing.T) {

	cases := map[Method]string{
		FirewallCheck:               "tholian-firewall.Check",
		FirewallForbidNetworks:      "tholian-firewall.ForbidNetworks",
		ScoutCollectPackages:        "tholian-scout.CollectPackages",
		ScoutSearchPackages:         "tholian-scout.SearchPackages",
		GuardCollectVulnerabilities: "tholian-guard.CollectVulnerabilities",
		WatchtowerListAudits:        "tholian-watchtower.ListAudits",
	}

	for method, expected := range cases {
		if method.String() != expected {
			t.Fatalf("method %q has wire value %q, expected %q", method.Project()+"."+method.Action(), method, expected)
		}
	}

}
