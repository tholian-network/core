package matchers

import "testing"

func TestSubnetDefaults(t *testing.T) {

	subnet := NewSubnet()

	if subnet.Name != Any || subnet.Subnet != Any {
		t.Fatalf("unexpected defaults: %+v", subnet)
	}

	if subnet.IsAny() == false || subnet.IsValid() == true {
		t.Fatal("a fresh subnet matcher should be any and invalid")
	}

}

func TestSubnetValidation(t *testing.T) {

	subnet := NewSubnet()
	subnet.SetName("AS1337")

	if subnet.HasName() == false || subnet.HasSubnet() == true || subnet.IsValid() == false {
		t.Fatalf("unexpected name state: %+v", subnet)
	}

	subnet.SetSubnet("[2001:db8::]/32")

	if subnet.HasName() == true || subnet.HasSubnet() == false || subnet.IsValid() == false {
		t.Fatalf("unexpected subnet state: %+v", subnet)
	}

	invalid := Subnet{Name: "AS1337", Subnet: "1.3.3.0/24"}

	if invalid.IsValid() == true {
		t.Fatal("a subnet with both selectors should be invalid")
	}

	invalid = Subnet{Name: "NOT-AN-ASN", Subnet: Any}

	if invalid.IsAny() == true || invalid.IsValid() == true || invalid.HasName() == true {
		t.Fatal("an invalid name should not be a valid selector")
	}

}

func TestSubnetSetSubnet(t *testing.T) {

	subnet := NewSubnet()
	subnet.SetSubnet("1.3.3.0/24")

	if subnet.HasSubnet() == false {
		t.Fatalf("expected a valid subnet: %+v", subnet)
	}

	subnet.SetSubnet("1.3.3.7")

	if subnet.Subnet != "1.3.3.0/24" {
		t.Fatalf("bare address replaced the previous value: %q", subnet.Subnet)
	}

	subnet.SetSubnet("[2001:db8::]/129")

	if subnet.Subnet != "1.3.3.0/24" {
		t.Fatalf("out of range prefix replaced the previous value: %q", subnet.Subnet)
	}

}

func TestToSubnetMatcher(t *testing.T) {

	subnet := ToSubnet("2001:db8::/32")

	if subnet.HasSubnet() == false || subnet.IsValid() == false {
		t.Fatalf("unexpected subnet: %+v", subnet)
	}

	subnet = ToSubnet("not a subnet")

	if subnet.IsAny() == false || subnet.IsValid() == true {
		t.Fatalf("unexpected invalid subnet: %+v", subnet)
	}

}
