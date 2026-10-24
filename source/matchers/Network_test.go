package matchers

import "testing"

func TestNetworkDefaults(t *testing.T) {

	network := NewNetwork()

	if network.Name != "any" || network.Subnet != "any" {
		t.Fatalf("unexpected defaults: %+v", network)
	}

	if network.IsAny() == false || network.IsValid() == true {
		t.Fatal("a fresh network matcher should be any and invalid")
	}

}

func TestNetworkSetName(t *testing.T) {

	network := NewNetwork()
	network.SetName("AS1337")

	if network.Name != "AS1337" || network.Subnet != Any {
		t.Fatalf("unexpected name: %+v", network)
	}

	network.SetName("AS12X")

	if network.Name != "AS1337" || network.Subnet != Any {
		t.Fatalf("invalid name replaced the previous value: %+v", network)
	}

	network.SetName("as1337")

	if network.Name != "AS1337" || network.Subnet != Any {
		t.Fatalf("lowercase name replaced the previous value: %+v", network)
	}

}

func TestNetworkSetSubnet(t *testing.T) {

	network := NewNetwork()
	network.SetSubnet("1.3.3.0/24")

	if network.Subnet != "1.3.3.0/24" {
		t.Fatalf("unexpected subnet: %q", network.Subnet)
	}

	network.SetSubnet("0.0.0.0/0")

	if network.Subnet != "0.0.0.0/0" {
		t.Fatalf("unexpected default route: %q", network.Subnet)
	}

	network.SetSubnet("not a subnet")

	if network.Subnet != "0.0.0.0/0" {
		t.Fatalf("invalid subnet replaced the previous value: %q", network.Subnet)
	}

	network.SetSubnet("1.3.3.0/33")

	if network.Subnet != "0.0.0.0/0" {
		t.Fatalf("out of range prefix replaced the previous value: %q", network.Subnet)
	}

	network.SetSubnet("1.3.3.7")

	if network.Subnet != "0.0.0.0/0" {
		t.Fatalf("bare address replaced the previous value: %q", network.Subnet)
	}

}

func TestNetworkSetSubnetIPv6(t *testing.T) {

	network := NewNetwork()
	network.SetSubnet("[2001:db8::]/32")

	if network.Subnet != "[2001:db8::]/32" {
		t.Fatalf("unexpected IPv6 subnet: %q", network.Subnet)
	}

	network.SetSubnet("2001:db8::/32")

	if network.Subnet != "2001:db8::/32" {
		t.Fatalf("unexpected unbracketed IPv6 subnet: %q", network.Subnet)
	}

	network.SetSubnet("[2001:db8::]/129")

	if network.Subnet != "2001:db8::/32" {
		t.Fatalf("out of range IPv6 prefix replaced the previous value: %q", network.Subnet)
	}

}

func TestNetworkHash(t *testing.T) {

	network := NewNetwork()
	network.SetSubnet("1.3.3.0/24")

	if network.Hash() != "01030300" {
		t.Fatalf("unexpected network hash: %q", network.Hash())
	}

	network = NewNetwork()
	network.SetName("AS1337")

	if network.Hash() != "AS1337" {
		t.Fatalf("unexpected name hash: %q", network.Hash())
	}

}

func TestNetworkMatches(t *testing.T) {

	network := NewNetwork()
	network.SetSubnet("1.3.3.0/24")

	if network.Matches("any", "1.3.3.7/32") == false {
		t.Fatal("expected the subnet to match a contained address")
	}

	if network.Matches("any", "1.3.4.7/32") == true {
		t.Fatal("subnet matched an unrelated address")
	}

	network = NewNetwork()
	network.SetName("AS1337")

	if network.Matches("AS1337", "any") == false {
		t.Fatal("expected the name to match")
	}

	if network.Matches("AS1338", "any") == true {
		t.Fatal("name matched an unrelated ASN")
	}

	network = NewNetwork()

	if network.Matches("AS1337", "1.3.3.7/32") == false {
		t.Fatal("an any network should match everything")
	}

}

func TestNetworkValidation(t *testing.T) {

	network := NewNetwork()

	if network.IsAny() == false || network.IsValid() == true || network.HasName() == true || network.HasSubnet() == true {
		t.Fatalf("unexpected fresh network state: %+v", network)
	}

	network.SetName("AS1337")

	if network.IsAny() == true || network.IsValid() == false || network.HasName() == false || network.HasSubnet() == true {
		t.Fatalf("unexpected name network state: %+v", network)
	}

	network.SetSubnet("1.3.3.0/24")

	if network.IsAny() == true || network.IsValid() == false || network.HasName() == true || network.HasSubnet() == false {
		t.Fatalf("unexpected subnet network state: %+v", network)
	}

	invalid := Network{Name: "AS1337", Subnet: "1.3.3.0/24"}

	if invalid.IsValid() == true {
		t.Fatal("a network with both selectors should be invalid")
	}

	invalid = Network{Name: "NOT-AN-ASN", Subnet: Any}

	if invalid.IsAny() == true || invalid.IsValid() == true || invalid.HasName() == true {
		t.Fatal("an invalid name should not be a valid selector")
	}

}

func TestNetworkSubnetValidation(t *testing.T) {

	valid := []string{
		"1.3.3.0/24",
		"0.0.0.0/0",
		"[2001:db8::]/32",
		"2001:db8::/32",
		"[::1]/128",
		"::1/128",
	}

	for _, value := range valid {

		network := NewNetwork()
		network.SetSubnet(value)

		if network.HasSubnet() == false {
			t.Errorf("expected %q to be a valid subnet", value)
		}

	}

	invalid := []string{
		"1.3.3.7",
		"1.3.3.0/33",
		"1.3.3.0/24/25",
		"[2001:db8::]/129",
		"2001:db8::/129",
		"[::1]",
		"not a subnet",
	}

	for _, value := range invalid {

		network := NewNetwork()
		network.SetSubnet(value)

		if network.HasSubnet() == true || network.IsAny() == false {
			t.Errorf("expected %q to be rejected", value)
		}

	}

}

func TestToNetwork(t *testing.T) {

	network := ToNetwork("AS1337")

	if network.HasName() == false || network.IsValid() == false {
		t.Fatalf("unexpected ASN network: %+v", network)
	}

	network = ToNetwork("1.3.3.0/24")

	if network.HasSubnet() == false || network.IsValid() == false {
		t.Fatalf("unexpected subnet network: %+v", network)
	}

	network = ToNetwork("1.3.3.0/33")

	if network.IsAny() == false || network.IsValid() == true {
		t.Fatalf("unexpected invalid network: %+v", network)
	}

}
