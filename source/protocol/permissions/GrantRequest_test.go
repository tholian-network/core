package permissions

import "testing"
import "tholian-core/protocol/capabilities"
import "tholian-core/protocol/methods"

func TestIsSupportedCapability(t *testing.T) {

	supported := []capabilities.Capability{
		capabilities.FirewallInstall,
		capabilities.FirewallRead,
		capabilities.FirewallWrite,
		capabilities.GuardInstall,
		capabilities.GuardRead,
		capabilities.GuardWrite,
		capabilities.ScoutInstall,
		capabilities.ScoutRead,
		capabilities.ScoutWrite,
	}

	for _, capability := range supported {
		if IsSupportedCapability(capability) == false {
			t.Fatalf("capability %q is not supported", capability)
		}
	}

	unsupported := []capabilities.Capability{
		capabilities.Capability(""),
		capabilities.Capability("tholian-firewall.Nonsense"),
		capabilities.Capability("packages.read"),
	}

	for _, capability := range unsupported {
		if IsSupportedCapability(capability) == true {
			t.Fatalf("capability %q was supported", capability)
		}
	}

}

func TestIsReservedCapability(t *testing.T) {

	reserved := []capabilities.Capability{
		capabilities.GuardWrite,
		capabilities.ScoutWrite,
	}

	for _, capability := range reserved {
		if IsReservedCapability(capability) == false {
			t.Fatalf("capability %q is not reserved", capability)
		}
	}

	unreserved := []capabilities.Capability{
		capabilities.FirewallInstall,
		capabilities.FirewallRead,
		capabilities.FirewallWrite,
		capabilities.GuardInstall,
		capabilities.GuardRead,
		capabilities.ScoutInstall,
		capabilities.ScoutRead,
	}

	for _, capability := range unreserved {
		if IsReservedCapability(capability) == true {
			t.Fatalf("capability %q was reserved", capability)
		}
	}

}

func TestMethodCapability(t *testing.T) {

	cases := map[methods.Method]capabilities.Capability{
		methods.FirewallInstall:             capabilities.FirewallInstall,
		methods.FirewallCheck:               capabilities.FirewallRead,
		methods.FirewallForbidNetworks:      capabilities.FirewallWrite,
		methods.GuardInstall:                capabilities.GuardInstall,
		methods.GuardCollectVulnerabilities: capabilities.GuardRead,
		methods.GuardExecuteShell:           capabilities.GuardWrite,
		methods.ScoutInstall:                capabilities.ScoutInstall,
		methods.ScoutSearchPackages:         capabilities.ScoutRead,
		methods.ScoutUpdatePackages:         capabilities.ScoutWrite,
	}

	for method, expected := range cases {

		capability, ok := MethodCapability(method)

		if ok == false {
			t.Fatalf("method %q has no capability mapping", method)
		}

		if capability != expected {
			t.Fatalf("method %q mapped to %q, expected %q", method, capability, expected)
		}

	}

	if _, ok := MethodCapability(methods.RPCPing); ok == true {
		t.Fatal("RPC method was mapped to a capability")
	}

	if _, ok := MethodCapability(methods.WatchtowerListAudits); ok == true {
		t.Fatal("watchtower method was mapped to a capability")
	}

}

func TestGrantRequestIsValid(t *testing.T) {

	var nil_request *GrantRequest

	if nil_request.IsValid() == true {
		t.Fatal("nil grant request was valid")
	}

	purpose := NewPurpose(PurposeInvestigation, "Investigate log4j")
	caps := []capabilities.Capability{capabilities.ScoutRead}

	valid := GrantRequest{
		SystemID:     "system-1",
		Purpose:      purpose,
		Capabilities: caps,
	}

	if valid.IsValid() == false {
		t.Fatal("valid grant request was rejected")
	}

	without_system := GrantRequest{
		Purpose:      purpose,
		Capabilities: caps,
	}

	if without_system.IsValid() == true {
		t.Fatal("grant request without system was valid")
	}

	without_purpose := GrantRequest{
		SystemID:     "system-1",
		Capabilities: caps,
	}

	if without_purpose.IsValid() == true {
		t.Fatal("grant request without purpose was valid")
	}

	without_capabilities := GrantRequest{
		SystemID: "system-1",
		Purpose:  purpose,
	}

	if without_capabilities.IsValid() == true {
		t.Fatal("grant request without capabilities was valid")
	}

}
