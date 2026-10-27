package structs

import "testing"

func TestInternetAddSubnet(t *testing.T) {

	internet := NewInternet()
	subnet := NewSubnet("1.1.1.0", 24)

	if internet.AddSubnet(subnet) == false {
		t.Fatal("AddSubnet failed")
	}

	if internet.AddSubnet(subnet) == true {
		t.Fatal("AddSubnet accepted a duplicate")
	}

	if len(internet.Mapv4[24]) != 1 {
		t.Fatalf("unexpected IPv4 map: %v", internet.Mapv4)
	}

}

func TestInternetAddNetwork(t *testing.T) {

	internet := NewInternet()
	network := NewNetwork("AS13335")
	network.AddSubnet(NewSubnet("1.1.1.0", 24))

	if internet.AddNetwork(network) == false {
		t.Fatal("AddNetwork failed")
	}

	if internet.SearchASN("AS13335").Name != "AS13335" {
		t.Fatal("SearchASN did not find the network")
	}

	invalid := NewNetwork("NOT-AN-ASN")

	if internet.AddNetwork(invalid) == true {
		t.Fatal("AddNetwork accepted a network without an AS prefix")
	}

}

func TestInternetSearchSubnet(t *testing.T) {

	internet := NewInternet()
	internet.AddSubnet(NewSubnet("1.1.0.0", 16))
	internet.AddSubnet(NewSubnet("1.1.1.0", 24))

	subnet := internet.Search("1.1.1.1")

	if subnet.Address != "1.1.1.0" || subnet.Prefix != 24 {
		t.Fatalf("longest prefix match failed: %s/%d", subnet.Address, subnet.Prefix)
	}

	subnet = internet.Search("1.1.2.1")

	if subnet.Address != "1.1.0.0" || subnet.Prefix != 16 {
		t.Fatalf("fallback prefix match failed: %s/%d", subnet.Address, subnet.Prefix)
	}

}

func TestInternetSearchDefaultRoute(t *testing.T) {

	internet := NewInternet()

	if internet.AddSubnet(ToSubnet("0.0.0.0/0")) == false {
		t.Fatal("AddSubnet failed for the default route")
	}

	network := NewNetwork("")
	network.AddSubnet(ToSubnet("1.3.3.7/32"))

	results := internet.SearchNetwork(network)

	if len(results) != 1 {
		t.Fatalf("unexpected default route search results: %v", results)
	}

	if results[0].IsDefaultRoute() == false {
		t.Fatalf("expected a default route, got %+v", results[0])
	}

}

func TestInternetSearchNetwork(t *testing.T) {

	internet := NewInternet()
	network := NewNetwork("AS13335")
	network.AddSubnet(NewSubnet("1.1.1.0", 24))
	internet.AddNetwork(network)

	results := internet.SearchNetwork(network)

	if len(results) != 1 {
		t.Fatalf("unexpected search results: %v", results)
	}

	empty := NewNetwork("")

	results = internet.SearchNetwork(empty)

	if len(results) != 0 {
		t.Fatalf("empty network returned results: %v", results)
	}

}

func TestInternetNewInternet(t *testing.T) {

	internet := NewInternet()

	if internet.Mapv4 == nil || internet.Mapv6 == nil || internet.Networks == nil {
		t.Fatal("NewInternet did not initialize all maps")
	}

	if internet.Hash() != "" {
		t.Fatal("a new internet should have an empty hash")
	}

}

func TestInternetIsValid(t *testing.T) {

	zero := Internet{}

	if zero.IsValid() == true {
		t.Fatal("a zero internet should be invalid")
	}

	empty := NewInternet()

	if empty.IsValid() == false {
		t.Fatal("an empty initialized internet should be valid")
	}

	valid := NewInternet()
	valid.AddSubnet(ToSubnet("10.0.0.0/8"))
	valid.AddSubnet(ToSubnet("[2001:db8::]/32"))

	network := NewNetwork("AS64500")
	network.AddSubnet(ToSubnet("192.0.2.0/24"))
	valid.AddNetwork(network)

	if valid.IsValid() == false {
		t.Fatal("a populated internet should be valid")
	}

	cases := []struct {
		name   string
		mutate func(*Internet)
	}{
		{"missing IPv4 map", func(value *Internet) { value.Mapv4 = nil }},
		{"missing IPv6 map", func(value *Internet) { value.Mapv6 = nil }},
		{"missing networks map", func(value *Internet) { value.Networks = nil }},
		{"nil subnet map", func(value *Internet) { value.Mapv4[8] = nil }},
		{"nil subnet", func(value *Internet) { value.Mapv4[8] = map[string]*Subnet{"missing": nil} }},
		{"wrong subnet key", func(value *Internet) {
			subnet := ToSubnet("10.0.0.0/8")
			value.Mapv4[8] = map[string]*Subnet{"wrong": &subnet}
		}},
		{"wrong subnet prefix", func(value *Internet) {
			subnet := ToSubnet("10.0.0.0/9")
			value.Mapv4[8] = map[string]*Subnet{subnet.Hash(): &subnet}
		}},
		{"wrong subnet type", func(value *Internet) {
			subnet := ToSubnet("[2001:db8::]/32")
			value.Mapv4[32] = map[string]*Subnet{subnet.Hash(): &subnet}
		}},
		{"invalid subnet", func(value *Internet) {
			subnet := Subnet{Address: "not-an-address", Type: "ipv4", Prefix: 8}
			value.Mapv4[8] = map[string]*Subnet{subnet.Hash(): &subnet}
		}},
		{"nil network", func(value *Internet) { value.Networks["AS64500"] = nil }},
		{"wrong network key", func(value *Internet) {
			value.Networks["AS64501"] = value.Networks["AS64500"]
			delete(value.Networks, "AS64500")
		}},
		{"invalid network", func(value *Internet) {
			value.Networks["AS64500"] = &Network{Name: "AS64500", Subnets: []Subnet{{}}}
		}},
		{"unregistered network subnet", func(value *Internet) {
			network := NewNetwork("AS64501")
			network.AddSubnet(ToSubnet("198.51.100.0/24"))
			value.Networks["AS64501"] = &network
		}},
	}

	for _, c := range cases {

		t.Run(c.name, func(t *testing.T) {

			other := NewInternet()
			other.AddSubnet(ToSubnet("10.0.0.0/8"))
			other.AddSubnet(ToSubnet("[2001:db8::]/32"))

			otherNetwork := NewNetwork("AS64500")
			otherNetwork.AddSubnet(ToSubnet("192.0.2.0/24"))
			other.AddNetwork(otherNetwork)

			c.mutate(&other)

			if other.IsValid() == true {
				t.Fatalf("internet should be invalid for %s", c.name)
			}

		})

	}

}

func TestInternetEqual(t *testing.T) {

	first := NewInternet()
	second := NewInternet()

	if first.Equal(second) == false {
		t.Fatal("two empty internets should be equal")
	}

	first.AddSubnet(ToSubnet("10.0.0.0/8"))
	second.AddSubnet(ToSubnet("10.0.0.0/8"))

	if first.Equal(second) == false {
		t.Fatal("internets with the same IPv4 subnet should be equal")
	}

	second.AddSubnet(ToSubnet("10.1.0.0/16"))

	if first.Equal(second) == true {
		t.Fatal("internets with different subnet counts should not be equal")
	}

	first.AddSubnet(ToSubnet("10.1.0.0/16"))

	if first.Equal(second) == false {
		t.Fatal("internets with the same IPv4 subnets should be equal")
	}

	first.AddSubnet(ToSubnet("[2001:db8::]/32"))
	second.AddSubnet(ToSubnet("[2001:db8::]/32"))

	if first.Equal(second) == false {
		t.Fatal("internets with the same IPv6 subnet should be equal")
	}

	third := NewInternet()
	fourth := NewInternet()
	network := NewNetwork("AS64500")
	network.AddSubnet(ToSubnet("192.0.2.0/24"))

	third.AddNetwork(network)
	fourth.AddNetwork(network)

	if third.Equal(fourth) == false {
		t.Fatal("internets with the same network should be equal")
	}

	fourth.Networks["AS64500"].SetName("AS64501")

	if third.Equal(fourth) == true {
		t.Fatal("internets with different networks should not be equal")
	}

}

func TestInternetAddSubnetIPv6(t *testing.T) {

	internet := NewInternet()
	subnet := ToSubnet("[2001:db8::]/32")

	if internet.AddSubnet(subnet) == false {
		t.Fatal("AddSubnet failed for an IPv6 subnet")
	}

	if internet.AddSubnet(subnet) == true {
		t.Fatal("AddSubnet accepted a duplicate IPv6 subnet")
	}

	if len(internet.Mapv6[32]) != 1 {
		t.Fatalf("unexpected IPv6 map: %v", internet.Mapv6)
	}

	if internet.AddSubnet(Subnet{}) == true {
		t.Fatal("AddSubnet accepted an invalid subnet")
	}

}

func TestInternetSearchASNInvalid(t *testing.T) {

	internet := NewInternet()
	internet.AddNetwork(NewNetwork("AS64500"))

	found := internet.SearchASN("AS64500")

	if found.Name != "AS64500" {
		t.Fatal("SearchASN did not find the network")
	}

	missingPrefix := internet.SearchASN("64500")

	if missingPrefix.HasName() == true {
		t.Fatal("SearchASN accepted a value without the AS prefix")
	}

	absent := internet.SearchASN("AS64501")

	if absent.HasName() == true {
		t.Fatal("SearchASN found an absent network")
	}

}

func TestInternetSearchIPv6(t *testing.T) {

	internet := NewInternet()
	internet.AddSubnet(ToSubnet("[2001:db8::]/32"))
	internet.AddSubnet(ToSubnet("[2001:db8:1::]/48"))

	subnet := internet.Search("[2001:db8:1::1]")

	if subnet.Prefix != 48 {
		t.Fatalf("IPv6 longest prefix match failed: %s", subnet.String())
	}

	subnet = internet.Search("[2001:db8:2::1]/128")

	if subnet.Prefix != 32 {
		t.Fatalf("IPv6 fallback prefix match failed: %s", subnet.String())
	}

	invalid := internet.Search("not-an-address")

	if invalid.IsValid() == true {
		t.Fatal("Search returned a subnet for an invalid address")
	}

	invalidSubnet := internet.SearchSubnet(Subnet{Address: "not-an-address", Type: "ipv4", Prefix: 24})

	if invalidSubnet.IsValid() == true {
		t.Fatal("SearchSubnet returned a subnet for an invalid IPv4 address")
	}

}

func TestInternetSearchNetworkSubnets(t *testing.T) {

	internet := NewInternet()
	internet.AddSubnet(ToSubnet("10.0.0.0/8"))
	internet.AddSubnet(ToSubnet("192.168.1.0/24"))

	query := NewNetwork("")
	query.AddSubnet(ToSubnet("10.1.2.3/32"))
	query.AddSubnet(ToSubnet("172.16.0.0/16"))

	results := internet.SearchNetwork(query)

	if len(results) != 1 {
		t.Fatalf("unexpected subnet search results: %v", results)
	}

	if results[0].Address != "10.0.0.0" || results[0].Prefix != 8 {
		t.Fatalf("unexpected subnet result: %+v", results[0])
	}

	absent := NewNetwork("AS64500")

	if len(internet.SearchNetwork(absent)) != 0 {
		t.Fatal("SearchNetwork found an absent named network")
	}

}

func TestInternetSetSubnets(t *testing.T) {

	internet := NewInternet()
	valid := ToSubnet("10.0.0.0/8")

	if internet.SetSubnets([]Subnet{valid, valid, Subnet{}}) == true {
		t.Fatal("SetSubnets should report duplicates and invalid entries")
	}

	if len(internet.Mapv4[8]) != 1 {
		t.Fatalf("SetSubnets did not add the valid subnet: %v", internet.Mapv4)
	}

}

func TestInternetSetNetworks(t *testing.T) {

	internet := NewInternet()

	valid := NewNetwork("AS64500")
	valid.AddSubnet(ToSubnet("10.0.0.0/8"))

	invalid := NewNetwork("not-an-asn")

	if internet.SetNetworks([]Network{valid, invalid}) == true {
		t.Fatal("SetNetworks should report an invalid network")
	}

	if len(internet.Networks) != 1 {
		t.Fatalf("unexpected networks: %v", internet.Networks)
	}

	if _, ok := internet.Networks["not-an-asn"]; ok == true {
		t.Fatal("SetNetworks accepted a network without an AS prefix")
	}

	if internet.SetNetworks([]Network{valid}) == true {
		t.Fatal("SetNetworks should report a duplicate subnet")
	}

}

func TestInternetHash(t *testing.T) {

	first := NewInternet()
	first.AddSubnet(ToSubnet("10.0.0.0/8"))
	first.AddSubnet(ToSubnet("[2001:db8::]/32"))

	firstNetwork := NewNetwork("AS64500")
	firstNetwork.AddSubnet(ToSubnet("192.0.2.0/24"))
	first.AddNetwork(firstNetwork)

	second := NewInternet()

	secondNetwork := NewNetwork("AS64500")
	secondNetwork.AddSubnet(ToSubnet("192.0.2.0/24"))
	second.AddNetwork(secondNetwork)
	second.AddSubnet(ToSubnet("[2001:db8::]/32"))
	second.AddSubnet(ToSubnet("10.0.0.0/8"))

	if first.Hash() == "" {
		t.Fatal("a populated internet should have a non-empty hash")
	}

	if first.Hash() != first.Hash() {
		t.Fatal("Hash should be deterministic")
	}

	if first.Hash() != second.Hash() {
		t.Fatal("Hash should not depend on insertion order")
	}

	different := NewInternet()
	different.AddSubnet(ToSubnet("10.0.0.0/8"))

	if first.Hash() == different.Hash() {
		t.Fatal("changing the subnets should change the hash")
	}

}
