package methods

import "testing"

func TestMethodString(t *testing.T) {

	method := Method("tholian-firewall.Check")

	if method.String() != "tholian-firewall.Check" {
		t.Fatalf("unexpected string: %q", method.String())
	}

}

func TestMethodProjectAndAction(t *testing.T) {

	method := FirewallForbidNetworks

	if method.Project() != "tholian-firewall" {
		t.Fatalf("unexpected project: %q", method.Project())
	}

	if method.Action() != "ForbidNetworks" {
		t.Fatalf("unexpected action: %q", method.Action())
	}

	invalid := Method("firewall")

	if invalid.Project() != "" || invalid.Action() != "" {
		t.Fatal("unqualified method was split into project and action")
	}

}

func TestMethodIsValid(t *testing.T) {

	if FirewallCheck.IsValid() == false {
		t.Fatal("known method was rejected")
	}

	if Method("tholian-firewall.Nonsense").IsValid() == true {
		t.Fatal("unknown method was accepted")
	}

	if Method("").IsValid() == true {
		t.Fatal("empty method was accepted")
	}

}

func TestParseMethod(t *testing.T) {

	parsed, ok := ParseMethod("tholian-firewall.ForbidNetworks")

	if ok == false {
		t.Fatal("known method was rejected")
	}

	if parsed != FirewallForbidNetworks {
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

	if _, ok := ParseMethod("tholian-firewall.Nonsense"); ok == true {
		t.Fatal("unknown method was accepted")
	}

}
