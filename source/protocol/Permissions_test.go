package protocol

import "testing"
import "tholian-core/types"
import "time"

func TestGrantAllowsCapability(t *testing.T) {

	purpose := NewPurpose(PurposeInvestigation, "Investigate")
	grant := NewGrant("system-1", "operator-1", purpose, []Capability{CapabilityPackagesRead}, NewUnrestrictedScope(), time.Hour)
	now := types.ToDatetime(time.Now().Format(time.RFC3339))

	if grant.IsValid() == false {
		t.Fatal("valid grant was rejected")
	}

	if grant.Allows("system-1", CapabilityPackagesRead, now) == false {
		t.Fatal("grant did not allow its capability")
	}

	if grant.Allows("system-2", CapabilityPackagesRead, now) == true {
		t.Fatal("grant allowed a different system")
	}

	if grant.Allows("system-1", CapabilityFirewallWrite, now) == true {
		t.Fatal("grant allowed a missing capability")
	}

}

func TestGrantExpiryAndRevocation(t *testing.T) {

	purpose := NewPurpose(PurposeIncident, "Incident")
	grant := NewGrant("system-1", "operator-1", purpose, []Capability{CapabilityFirewallWrite}, NewUnrestrictedScope(), -1*time.Minute)
	now := types.ToDatetime(time.Now().Format(time.RFC3339))

	if grant.Allows("system-1", CapabilityFirewallWrite, now) == true {
		t.Fatal("expired grant was accepted")
	}

	grant = NewGrant("system-1", "operator-1", purpose, []Capability{CapabilityFirewallWrite}, NewUnrestrictedScope(), time.Hour)
	grant.Revoked = true

	if grant.Allows("system-1", CapabilityFirewallWrite, now) == true {
		t.Fatal("revoked grant was accepted")
	}

}

func TestGrantMaxUses(t *testing.T) {

	purpose := NewPurpose(PurposeMitigation, "Mitigate")
	grant := NewGrant("system-1", "operator-1", purpose, []Capability{CapabilityPackagesUpdate}, NewUnrestrictedScope(), time.Hour)
	grant.MaxUses = 1
	grant.Uses = 1
	now := types.ToDatetime(time.Now().Format(time.RFC3339))

	if grant.Allows("system-1", CapabilityPackagesUpdate, now) == true {
		t.Fatal("exhausted grant was accepted")
	}

}

func TestMethodCapabilityMapping(t *testing.T) {

	cases := map[string]Capability{
		MethodSystemReport:          CapabilitySystemRead,
		MethodPackagesSearch:        CapabilityPackagesRead,
		MethodPackagesUpdate:        CapabilityPackagesUpdate,
		MethodFirewallStatus:        CapabilityFirewallRead,
		MethodFirewallForbid:        CapabilityFirewallWrite,
		MethodIncidentsCollectSince: CapabilityIncidentsRead,
	}

	for method, expected := range cases {

		capability, ok := MethodCapability(method)

		if ok == false {
			t.Fatalf("method %q has no capability", method)
		}

		if capability != expected {
			t.Fatalf("method %q maps to %q, expected %q", method, capability, expected)
		}

	}

	if _, ok := MethodCapability(MethodVulnerabilitiesScan); ok == true {
		t.Fatal("reserved method unexpectedly has a capability")
	}

}

func TestCapabilitySupport(t *testing.T) {

	if IsSupportedCapability(CapabilityPackagesRead) == false {
		t.Fatal("packages.read should be supported")
	}

	if IsSupportedCapability(Capability("nonsense")) == true {
		t.Fatal("unknown capability was reported as supported")
	}

	if IsReservedCapability(CapabilityShellExecute) == false {
		t.Fatal("shell.execute should be reserved")
	}

	if IsSupportedCapability(CapabilityShellExecute) == true {
		t.Fatal("reserved capability was reported as supported")
	}

}
