package structs

import "testing"

func TestNetworkDeduplicatesSubnets(t *testing.T) {

	network := NewNetwork("AS13335")

	network.AddSubnet(ToSubnet("1.1.1.0/24"))
	network.AddSubnet(ToSubnet("1.1.1.0/24"))

	if len(network.Subnets) != 1 {
		t.Fatalf("expected 1 subnet, got %d", len(network.Subnets))
	}

	network.RemoveSubnet(ToSubnet("1.1.1.0/24"))

	if len(network.Subnets) != 0 {
		t.Fatalf("expected 0 subnets, got %d", len(network.Subnets))
	}

}

func TestNetworkPredicates(t *testing.T) {

	network := NewNetwork("")

	if network.HasName() == true || network.HasSubnets() == true || network.IsValid() == true {
		t.Fatal("an empty network should be invalid")
	}

	network.SetName("AS13335")
	network.AddSubnet(ToSubnet("1.1.1.0/24"))

	if network.HasName() == false || network.HasSubnets() == false || network.IsValid() == false {
		t.Fatal("a populated network should be valid")
	}

}

func TestNetworkNewNetwork(t *testing.T) {

	network := NewNetwork("  AS13335  ")

	if network.Name != "AS13335" {
		t.Fatalf("NewNetwork did not trim the name: %q", network.Name)
	}

	if network.Subnets == nil {
		t.Fatal("NewNetwork did not initialize the subnets")
	}

}

func TestNetworkEqual(t *testing.T) {

	first := NewNetwork("AS13335")
	first.AddSubnet(ToSubnet("10.0.0.0/8"))
	first.AddSubnet(ToSubnet("10.1.0.0/16"))

	second := NewNetwork("AS13335")
	second.AddSubnet(ToSubnet("10.1.0.0/16"))
	second.AddSubnet(ToSubnet("10.0.0.0/8"))

	if first.Equal(second) == false {
		t.Fatal("networks with the same subnets in a different order should be equal")
	}

	differentName := NewNetwork("AS64500")
	differentName.AddSubnet(ToSubnet("10.0.0.0/8"))
	differentName.AddSubnet(ToSubnet("10.1.0.0/16"))

	if first.Equal(differentName) == true {
		t.Fatal("networks with different names should not be equal")
	}

	differentLength := NewNetwork("AS13335")
	differentLength.AddSubnet(ToSubnet("10.0.0.0/8"))

	if first.Equal(differentLength) == true {
		t.Fatal("networks with a different subnet count should not be equal")
	}

	differentSubnet := NewNetwork("AS13335")
	differentSubnet.AddSubnet(ToSubnet("10.0.0.0/8"))
	differentSubnet.AddSubnet(ToSubnet("10.2.0.0/16"))

	if first.Equal(differentSubnet) == true {
		t.Fatal("networks with different subnets should not be equal")
	}

	empty := NewNetwork("")
	otherEmpty := NewNetwork("")

	if empty.Equal(otherEmpty) == false {
		t.Fatal("two empty networks should be equal")
	}

}

func TestNetworkIsValid(t *testing.T) {

	empty := NewNetwork("")

	if empty.IsValid() == true {
		t.Fatal("a network without a name should be invalid")
	}

	named := NewNetwork("AS13335")

	if named.IsValid() == false {
		t.Fatal("a named network without subnets should be valid")
	}

	invalid := NewNetwork("AS13335")
	invalid.Subnets = []Subnet{{}}

	if invalid.IsValid() == true {
		t.Fatal("a network with an invalid subnet should be invalid")
	}

}

func TestNetworkSetName(t *testing.T) {

	network := NewNetwork("")

	network.SetName("  AS13335  ")

	if network.Name != "AS13335" {
		t.Fatalf("SetName did not trim: %q", network.Name)
	}

}

func TestNetworkAddSubnet(t *testing.T) {

	network := NewNetwork("AS13335")

	network.AddSubnet(ToSubnet("10.0.0.0/8"))
	network.AddSubnet(ToSubnet("10.0.0.0/8"))
	network.AddSubnet(Subnet{})
	network.AddSubnet(ToSubnet("invalid"))

	if len(network.Subnets) != 1 {
		t.Fatalf("AddSubnet did not deduplicate invalid entries: %d", len(network.Subnets))
	}

}

func TestNetworkRemoveSubnet(t *testing.T) {

	values := []Subnet{
		ToSubnet("10.0.0.0/8"),
		ToSubnet("10.1.0.0/16"),
		ToSubnet("10.1.1.0/24"),
	}

	for _, position := range []int{0, 1, 2} {

		network := NewNetwork("AS13335")
		network.SetSubnets(values)
		network.RemoveSubnet(values[position])

		if len(network.Subnets) != 2 {
			t.Fatalf("RemoveSubnet at position %d left %d entries", position, len(network.Subnets))
		}

		for _, subnet := range network.Subnets {
			if subnet.Equal(values[position]) {
				t.Fatalf("RemoveSubnet at position %d left the removed subnet", position)
			}
		}

	}

	noop := NewNetwork("AS13335")
	noop.SetSubnets(values)
	noop.RemoveSubnet(ToSubnet("192.0.2.0/24"))

	if len(noop.Subnets) != 3 {
		t.Fatal("RemoveSubnet removed an absent subnet")
	}

}

func TestNetworkSetSubnets(t *testing.T) {

	network := NewNetwork("AS13335")

	network.SetSubnets([]Subnet{
		ToSubnet("10.1.1.0/24"),
		Subnet{},
		ToSubnet("10.0.0.0/8"),
		ToSubnet("10.1.0.0/16"),
	})

	if len(network.Subnets) != 3 {
		t.Fatalf("SetSubnets did not filter invalid entries: %d", len(network.Subnets))
	}

	expected := []uint8{8, 16, 24}

	for index, prefix := range expected {
		if network.Subnets[index].Prefix != prefix {
			t.Fatalf("SetSubnets did not sort at index %d: %v", index, network.Subnets)
		}
	}

	samePrefix := NewNetwork("AS13335")
	samePrefix.SetSubnets([]Subnet{
		ToSubnet("1.1.2.0/24"),
		ToSubnet("1.1.1.0/24"),
	})

	if samePrefix.Subnets[0].Address != "1.1.1.0" {
		t.Fatalf("SetSubnets did not sort equal prefixes by hash: %v", samePrefix.Subnets)
	}

	samePrefix.SetSubnets(nil)

	if len(samePrefix.Subnets) != 0 {
		t.Fatal("SetSubnets did not replace the subnets")
	}

}

func TestNetworkHasSubnet(t *testing.T) {

	network := NewNetwork("AS13335")

	subnet := ToSubnet("10.0.0.0/8")
	other := ToSubnet("10.1.0.0/16")

	if network.HasSubnet(subnet) == true {
		t.Fatal("HasSubnet did not report a missing subnet")
	}

	network.AddSubnet(subnet)

	if network.HasSubnet(subnet) == false {
		t.Fatal("HasSubnet did not find an existing subnet")
	}

	if network.HasSubnet(other) == true {
		t.Fatal("HasSubnet found a different subnet")
	}

}

func TestNetworkHash(t *testing.T) {

	empty := NewNetwork("")

	if empty.Hash() != "" {
		t.Fatal("an empty network should have an empty hash")
	}

	nameOnly := NewNetwork("AS13335")

	if nameOnly.Hash() == "" {
		t.Fatal("a named network should have a non-empty hash")
	}

	first := NewNetwork("AS13335")
	first.AddSubnet(ToSubnet("10.0.0.0/8"))
	first.AddSubnet(ToSubnet("10.1.0.0/16"))

	second := NewNetwork("AS13335")
	second.AddSubnet(ToSubnet("10.1.0.0/16"))
	second.AddSubnet(ToSubnet("10.0.0.0/8"))

	if first.Hash() == "" {
		t.Fatal("a populated network should have a non-empty hash")
	}

	if first.Hash() != second.Hash() {
		t.Fatal("Hash should not depend on subnet order")
	}

	different := NewNetwork("AS64500")
	different.AddSubnet(ToSubnet("10.0.0.0/8"))
	different.AddSubnet(ToSubnet("10.1.0.0/16"))

	if first.Hash() == different.Hash() {
		t.Fatal("changing the name should change the hash")
	}

}
