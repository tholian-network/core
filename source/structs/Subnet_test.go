package structs

import "testing"
import "tholian-core/types"

func TestToSubnet(t *testing.T) {

	cases := []struct {
		value   string
		valid   bool
		output  string
		address string
		prefix  uint8
	}{
		{"1.3.3.0/24", true, "1.3.3.0/24", "1.3.3.0", 24},
		{"1.3.3.7/32", true, "1.3.3.7/32", "1.3.3.7", 32},
		{"0.0.0.0/0", true, "0.0.0.0/0", "0.0.0.0", 0},
		{"[fe80::1]/64", true, "fe80:0000:0000:0000:0000:0000:0000:0001/64", "fe80:0000:0000:0000:0000:0000:0000:0001", 64},
		{"[::]/0", true, "0000:0000:0000:0000:0000:0000:0000:0000/0", "0000:0000:0000:0000:0000:0000:0000:0000", 0},
		{"1.3.3.7/33", false, "", "", 0},
		{"invalid", false, "", "", 0},
	}

	for _, c := range cases {

		subnet := ToSubnet(c.value)

		if subnet.IsValid() != c.valid {
			t.Errorf("ToSubnet(%q).IsValid() = %v, expected %v", c.value, subnet.IsValid(), c.valid)
			continue
		}

		if c.valid == false {
			continue
		}

		if subnet.String() != c.output {
			t.Errorf("ToSubnet(%q).String() = %q, expected %q", c.value, subnet.String(), c.output)
		}

		if subnet.Address != c.address || subnet.Prefix != c.prefix {
			t.Errorf("ToSubnet(%q) = %s/%d, expected %s/%d", c.value, subnet.Address, subnet.Prefix, c.address, c.prefix)
		}

	}

}

func TestSubnetContainsSubnet(t *testing.T) {

	outer := ToSubnet("1.3.3.0/24")
	inner := ToSubnet("1.3.3.7/32")
	other := ToSubnet("1.3.4.7/32")

	if outer.ContainsSubnet(inner) == false {
		t.Fatal("/24 should contain /32")
	}

	if outer.ContainsSubnet(other) == true {
		t.Fatal("/24 should not contain a different /24")
	}

	if outer.ContainsSubnet(outer) == false {
		t.Fatal("identical subnets should be contained")
	}

	if outer.Contains("1.3.3.7") == false {
		t.Fatal("/24 should contain its address")
	}

	if outer.Contains("1.3.4.7") == true {
		t.Fatal("/24 contained an unrelated address")
	}

	defaultRoute := ToSubnet("0.0.0.0/0")

	if defaultRoute.ContainsSubnet(inner) == false {
		t.Fatal("/0 should contain every IPv4 subnet")
	}

	ipv6 := ToSubnet("[fe80::1]/128")

	if ipv6.Contains("[fe80::1]") == false {
		t.Fatal("IPv6 /128 should contain its address")
	}

	if ipv6.Contains("[fe80::2]") == true {
		t.Fatal("IPv6 /128 contained an unrelated address")
	}

}

func TestSubnetPredicates(t *testing.T) {

	ipv4 := ToSubnet("0.0.0.0/0")

	if ipv4.IsUnspecified() == false || ipv4.IsDefaultRoute() == false || ipv4.IsValid() == false {
		t.Fatalf("unexpected IPv4 default route predicates: %+v", ipv4)
	}

	ipv6 := ToSubnet("[::]/0")

	if ipv6.IsUnspecified() == false || ipv6.IsDefaultRoute() == false || ipv6.IsValid() == false {
		t.Fatalf("unexpected IPv6 default route predicates: %+v", ipv6)
	}

	address := ToSubnet("1.3.3.7/32")

	if address.IsUnspecified() == true || address.IsDefaultRoute() == true || address.IsAddress() == false || address.IsValid() == false {
		t.Fatalf("unexpected address predicates: %+v", address)
	}

	if ipv4.IsAddress() == true {
		t.Fatal("a default route should not be an address")
	}

	zero := Subnet{}

	if zero.IsValid() == true {
		t.Fatal("the zero subnet should be invalid")
	}

}

func TestSubnetNewSubnet(t *testing.T) {

	subnet := NewSubnet("10.0.0.0", 24)

	if subnet.IsValid() == false {
		t.Fatal("NewSubnet produced an invalid subnet")
	}

	if subnet.Address != "10.0.0.0" ||
		subnet.Prefix != 24 ||
		subnet.Type != types.AddressTypeIPv4 ||
		subnet.Scope != "private" {
		t.Fatalf("unexpected subnet: %+v", subnet)
	}

	if subnet.String() != "10.0.0.0/24" {
		t.Fatalf("unexpected string: %q", subnet.String())
	}

	invalid := NewSubnet("not-an-address", 24)

	if invalid.IsValid() == true {
		t.Fatalf("NewSubnet accepted an invalid address: %+v", invalid)
	}

}

func TestSubnetEqual(t *testing.T) {

	first := ToSubnet("10.0.0.0/8")

	if first.Equal(ToSubnet("10.0.0.0/8")) == false {
		t.Fatal("identical subnets should be equal")
	}

	second := ToSubnet("10.0.0.0/8")
	second.Name = "Other"
	second.Country = "DE"
	second.Scope = "public"

	if first.Equal(second) == false {
		t.Fatal("Equal should only compare address and prefix")
	}

	if first.Equal(ToSubnet("10.0.0.0/16")) == true {
		t.Fatal("subnets with different prefixes should not be equal")
	}

	if first.Equal(ToSubnet("11.0.0.0/8")) == true {
		t.Fatal("subnets with different addresses should not be equal")
	}

}

func TestSubnetParse(t *testing.T) {

	cases := []struct {
		value   string
		valid   bool
		address string
		prefix  uint8
		typ     string
	}{
		{"192.168.1.0/24", true, "192.168.1.0", 24, types.AddressTypeIPv4},
		{"192.168.1.1", true, "192.168.1.1", 32, types.AddressTypeIPv4},
		{"[2001:db8::1]/64", true, "2001:0db8:0000:0000:0000:0000:0000:0001", 64, types.AddressTypeIPv6},
		{"[2001:db8::1]", true, "2001:0db8:0000:0000:0000:0000:0000:0001", 128, types.AddressTypeIPv6},
		{"2001:db8::1/64", true, "2001:0db8:0000:0000:0000:0000:0000:0001", 64, types.AddressTypeIPv6},
		{"2001:db8::1", true, "2001:0db8:0000:0000:0000:0000:0000:0001", 128, types.AddressTypeIPv6},
		{"10.0.0.0/33", false, "", 0, ""},
		{"2001:db8::1/129", false, "", 0, ""},
		{"10.0.0.0/not-a-prefix", false, "", 0, ""},
		{"not-a-subnet", false, "", 0, ""},
	}

	for _, c := range cases {

		subnet := ToSubnet(c.value)

		if subnet.IsValid() != c.valid {
			t.Errorf("ToSubnet(%q).IsValid() = %v, expected %v", c.value, subnet.IsValid(), c.valid)
			continue
		}

		if c.valid == false {
			continue
		}

		if subnet.Address != c.address || subnet.Prefix != c.prefix || subnet.Type != c.typ {
			t.Errorf("ToSubnet(%q) = %+v, expected %s/%d (%s)", c.value, subnet, c.address, c.prefix, c.typ)
		}

	}

}

func TestSubnetSetters(t *testing.T) {

	t.Run("SetCountry", func(t *testing.T) {

		subnet := ToSubnet("10.0.0.0/8")
		subnet.SetCountry("DE")

		if subnet.Country != "DE" {
			t.Fatalf("unexpected country: %q", subnet.Country)
		}

		subnet.SetCountry("XX")

		if subnet.Country != "DE" {
			t.Fatal("SetCountry accepted XX")
		}

		subnet.SetCountry("USA")

		if subnet.Country != "DE" {
			t.Fatal("SetCountry accepted an invalid country length")
		}

	})

	t.Run("SetName", func(t *testing.T) {

		subnet := ToSubnet("10.0.0.0/8")
		subnet.SetName("  Internal  ")

		if subnet.Name != "Internal" {
			t.Fatalf("SetName did not trim: %q", subnet.Name)
		}

		subnet.SetName("   ")

		if subnet.Name != "Internal" {
			t.Fatal("SetName accepted an empty name")
		}

	})

	t.Run("SetAddress", func(t *testing.T) {

		subnet := NewSubnet("10.0.0.0", 8)
		subnet.SetAddress("2001:db8::1")

		if subnet.Type != types.AddressTypeIPv6 ||
			subnet.Address != "2001:0db8:0000:0000:0000:0000:0000:0001" ||
			subnet.Scope != "public" {
			t.Fatalf("unexpected IPv6 address: %+v", subnet)
		}

		subnet.SetAddress("192.168.1.1")

		if subnet.Type != types.AddressTypeIPv4 ||
			subnet.Address != "192.168.1.1" ||
			subnet.Scope != "private" {
			t.Fatalf("unexpected IPv4 address: %+v", subnet)
		}

		before := subnet
		subnet.SetAddress("not-an-address")

		if subnet.Address != before.Address ||
			subnet.Type != before.Type ||
			subnet.Scope != before.Scope {
			t.Fatal("SetAddress accepted an invalid address")
		}

	})

	t.Run("SetPrefix", func(t *testing.T) {

		ipv4 := NewSubnet("10.0.0.0", 8)
		ipv4.SetPrefix(24)

		if ipv4.Prefix != 24 {
			t.Fatalf("unexpected IPv4 prefix: %d", ipv4.Prefix)
		}

		ipv4.SetPrefix(33)

		if ipv4.Prefix != 24 {
			t.Fatal("SetPrefix accepted an oversized IPv4 prefix")
		}

		ipv6 := NewSubnet("2001:db8::1", 64)

		if ipv6.Prefix != 64 {
			t.Fatalf("unexpected IPv6 prefix: %d", ipv6.Prefix)
		}

		ipv6.SetPrefix(129)

		if ipv6.Prefix != 64 {
			t.Fatal("SetPrefix accepted an oversized IPv6 prefix")
		}

		invalid := Subnet{}
		invalid.SetPrefix(24)

		if invalid.Prefix != 0 {
			t.Fatal("SetPrefix changed the prefix without an address type")
		}

	})

	t.Run("SetScope", func(t *testing.T) {

		subnet := ToSubnet("10.0.0.0/8")
		subnet.SetScope("public")

		if subnet.Scope != "public" {
			t.Fatalf("unexpected scope: %q", subnet.Scope)
		}

		subnet.SetScope("private")

		if subnet.Scope != "private" {
			t.Fatalf("unexpected scope: %q", subnet.Scope)
		}

		subnet.SetScope("invalid")

		if subnet.Scope != "private" {
			t.Fatal("SetScope accepted an invalid scope")
		}

	})

}

func TestSubnetIPv6Predicates(t *testing.T) {

	address := ToSubnet("[2001:db8::1]/128")

	if address.IsAddress() == false ||
		address.IsUnspecified() == true ||
		address.IsDefaultRoute() == true ||
		address.IsValid() == false {
		t.Fatalf("unexpected IPv6 address predicates: %+v", address)
	}

	network := ToSubnet("[2001:db8::]/64")

	if network.IsAddress() == true || network.IsValid() == false {
		t.Fatalf("unexpected IPv6 network predicates: %+v", network)
	}

	zero := ToSubnet("[::]/0")

	if zero.IsUnspecified() == false || zero.IsDefaultRoute() == false {
		t.Fatalf("unexpected unspecified IPv6 predicates: %+v", zero)
	}

}

func TestSubnetAddresses(t *testing.T) {

	cases := []struct {
		value  string
		expect uint
	}{
		{"10.0.0.0/24", 256},
		{"10.0.0.0/30", 4},
		{"10.0.0.1/32", 1},
		{"0.0.0.0/0", 0},
		{"[2001:db8::]/126", 4},
		{"[2001:db8::1]/128", 1},
		{"[2001:db8::]/0", 0},
	}

	for _, c := range cases {

		subnet := ToSubnet(c.value)

		if subnet.Addresses() != c.expect {
			t.Errorf("Addresses(%q) = %d, expected %d", c.value, subnet.Addresses(), c.expect)
		}

	}

	invalid := Subnet{}

	if invalid.Addresses() != 0 {
		t.Fatal("an invalid subnet should have no addresses")
	}

}

func TestSubnetHash(t *testing.T) {

	invalid := Subnet{}

	if invalid.Hash() != "" {
		t.Fatal("an invalid subnet should have an empty hash")
	}

	first := ToSubnet("1.1.1.0/24")

	if first.Hash() != "01010100" {
		t.Fatalf("unexpected IPv4 hash: %q", first.Hash())
	}

	masked := ToSubnet("1.1.1.255/24")

	if masked.Hash() != first.Hash() {
		t.Fatal("Hash should mask the address to the prefix")
	}

	wider := ToSubnet("1.1.0.0/16")

	if wider.Hash() == first.Hash() {
		t.Fatal("changing the prefix should change the hash")
	}

	ipv6 := ToSubnet("[2001:db8::]/64")

	if ipv6.Hash() == "" {
		t.Fatal("an IPv6 subnet should have a non-empty hash")
	}

	if ipv6.Hash() != ipv6.Hash() {
		t.Fatal("Hash should be deterministic")
	}

}

func TestSubnetString(t *testing.T) {

	subnet := ToSubnet("10.0.0.0/8")

	if subnet.String() != "10.0.0.0/8" {
		t.Fatalf("unexpected string: %q", subnet.String())
	}

	invalid := Subnet{}

	if invalid.String() != "" {
		t.Fatal("an invalid subnet should have an empty string")
	}

}
