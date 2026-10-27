package types

import (
	"encoding/json"
	"testing"
)

func FuzzParseIPv6(f *testing.F) {

	f.Add("[fe80::1]")
	f.Add("::1")
	f.Add("1:2:3:4:5:6:7:8:9")

	f.Fuzz(func(t *testing.T, value string) {
		_ = IsIPv6(value)
		_ = ParseIPv6(value)
	})

}

func TestIsIPv6(t *testing.T) {

	cases := map[string]bool{
		"[fe80::1]":     true,
		"fe80::1":       true,
		"::1":           true,
		"[::1]":         true,
		"[2001:db8::1]": true,
		"2001:0db8:0000:0000:0000:0000:0000:0001": true,
		"[fe80::1]:1338": false,
		"[fe80::1]/64":   false,
		"1.2.3.4":        false,
		"[]":             false,
		"":               false,
	}

	for input, expected := range cases {

		if result := IsIPv6(input); result != expected {
			t.Errorf("IsIPv6(%q) = %v, expected %v", input, result, expected)
		}

	}

}

func TestIsIPv6AndPrefix(t *testing.T) {

	cases := map[string]bool{
		"[fe80::1]/64":  true,
		"[fe80::1]/8":   true,
		"[fe80::1]/128": true,
		"[fe80::1]/7":   true,
		"[fe80::1]/129": false,
		"[fe80::1]/0":   true,
		"fe80::1/64":    false,
		"[fe80::1]":     false,
	}

	for input, expected := range cases {

		if result := IsIPv6AndPrefix(input); result != expected {
			t.Errorf("IsIPv6AndPrefix(%q) = %v, expected %v", input, result, expected)
		}

	}

}

func TestIsIPv6AndPort(t *testing.T) {

	cases := map[string]bool{
		"[fe80::1]:1338": true,
		"[fe80::1]":      false,
		"fe80::1:1338":   false,
		"[fe80::1]:":     false,
	}

	for input, expected := range cases {

		if result := IsIPv6AndPort(input); result != expected {
			t.Errorf("IsIPv6AndPort(%q) = %v, expected %v", input, result, expected)
		}

	}

}

func TestParseIPv6AcceptsEquivalents(t *testing.T) {

	values := []string{
		"::1",
		"[::1]",
		"0:0:0:0:0:0:0:1",
		"2001:db8::1",
	}

	expected := "0000:0000:0000:0000:0000:0000:0000:0001"

	for _, value := range values {

		ipv6 := ParseIPv6(value)

		if ipv6 == nil {
			t.Fatalf("ParseIPv6(%q) returned nil", value)
		}

		if value == "2001:db8::1" {
			expected = "2001:0db8:0000:0000:0000:0000:0000:0001"
		}

		if ipv6.String() != expected {
			t.Errorf("ParseIPv6(%q).String() = %q, expected %q", value, ipv6.String(), expected)
		}

	}

}

func TestParseIPv6AndPrefix(t *testing.T) {

	ipv6, prefix := ParseIPv6AndPrefix("[fe80::1]/64")

	if ipv6 == nil {
		t.Fatalf("ParseIPv6AndPrefix returned nil address")
	}

	if prefix != 64 {
		t.Errorf("ParseIPv6AndPrefix prefix = %d, expected 64", prefix)
	}

	if ipv6.String() != "fe80:0000:0000:0000:0000:0000:0000:0001" {
		t.Errorf("ParseIPv6AndPrefix address = %q", ipv6.String())
	}

}

func TestParseIPv6AndPort(t *testing.T) {

	ipv6, port := ParseIPv6AndPort("[fe80::1]:1338")

	if ipv6 == nil {
		t.Fatalf("ParseIPv6AndPort returned nil address")
	}

	if port != 1338 {
		t.Errorf("ParseIPv6AndPort port = %d, expected 1338", port)
	}

}

func TestParseIPv6NonOctetPrefix(t *testing.T) {

	ipv6 := ParseIPv6("2001:db8:1234:5678::1")

	if ipv6 == nil {
		t.Fatal("ParseIPv6 returned nil")
	}

	result := ipv6.Bytes(68)

	if len(result) != 16 {
		t.Fatalf("IPv6.Bytes(68) length = %d", len(result))
	}

	if result[8] != 0 || result[9] != 0 {
		t.Errorf("IPv6.Bytes(68) did not mask the host bits: %v", result)
	}

	if result[0] != 0x20 || result[1] != 0x01 || result[2] != 0x0d || result[3] != 0xb8 {
		t.Errorf("IPv6.Bytes(68) lost the network bits: %v", result)
	}

}

func TestParseIPv6RejectsMalformedInput(t *testing.T) {

	values := []string{
		"",
		"[]",
		"[fe80::1",
		"fe80::1]",
		"1:2:3:4:5:6:7:8:9",
		"::1::2",
		"fe80::1%eth0",
		"fe80::1/64",
		"[::ffff:1.2.3.4]",
		"fe80:::1",
	}

	for _, value := range values {

		if ipv6 := ParseIPv6(value); ipv6 != nil {
			t.Errorf("ParseIPv6(%q) = %v, expected nil", value, ipv6.String())
		}

	}

}

func TestIPv6Scope(t *testing.T) {

	private := []string{
		"[fe80::1]",
		"[::1]",
		"[fc00::1]",
		"[fd00::1]",
	}

	for p := 0; p < len(private); p++ {

		ipv6 := ParseIPv6(private[p])

		if ipv6 == nil {
			t.Fatalf("ParseIPv6(%q) returned nil", private[p])
		}

		if scope := ipv6.Scope(); scope != "private" {
			t.Errorf("ParseIPv6(%q).Scope() = %q, expected private", private[p], scope)
		}

	}

	public := "[2001:db8::1]"
	ipv6 := ParseIPv6(public)

	if ipv6 == nil {
		t.Fatalf("ParseIPv6(%q) returned nil", public)
	}

	if scope := ipv6.Scope(); scope != "public" {
		t.Errorf("ParseIPv6(%q).Scope() = %q, expected public", public, scope)
	}

}

func TestIPv6String(t *testing.T) {

	ipv6 := ParseIPv6("[2001:db8::1]")

	if ipv6 == nil {
		t.Fatalf("ParseIPv6 returned nil")
	}

	if ipv6.String() != "2001:0db8:0000:0000:0000:0000:0000:0001" {
		t.Errorf("IPv6.String() = %q", ipv6.String())
	}

}

func TestIPv6FactoryAndParse(t *testing.T) {

	zero := NewIPv6()

	if zero.Equal(IPv6{}) == false || zero.IsUnspecified() == false {
		t.Fatal("NewIPv6() should return the unspecified address")
	}

	if zero.IsValid() == false {
		t.Fatal("the unspecified IPv6 address should be valid")
	}

	if zero.String() != UnspecifiedIPv6 {
		t.Errorf("NewIPv6().String() = %q, expected %q", zero.String(), UnspecifiedIPv6)
	}

	cases := []struct {
		value    string
		expected string
	}{
		{"::", UnspecifiedIPv6},
		{"::1", "0000:0000:0000:0000:0000:0000:0000:0001"},
		{"[::1]", "0000:0000:0000:0000:0000:0000:0000:0001"},
		{"2001:db8::1", "2001:0db8:0000:0000:0000:0000:0000:0001"},
	}

	for _, c := range cases {

		ipv6 := ParseIPv6(c.value)

		if ipv6 == nil {
			t.Fatalf("ParseIPv6(%q) returned nil", c.value)
		}

		if ipv6.String() != c.expected {
			t.Errorf("ParseIPv6(%q).String() = %q, expected %q", c.value, ipv6.String(), c.expected)
		}

		if ToIPv6(c.value).Equal(*ipv6) == false {
			t.Errorf("ToIPv6(%q) does not match ParseIPv6(%q)", c.value, c.value)
		}

	}

	if ToIPv6("not an address").Equal(IPv6{}) == false {
		t.Fatal("ToIPv6() should return the zero value for invalid input")
	}

	var ipv6 IPv6

	if ipv6.Parse("[2001:db8::1]") == false || ipv6.Equal(ToIPv6("2001:db8::1")) == false {
		t.Fatal("IPv6.Parse() failed for valid input")
	}

	if ipv6.Parse("not an address") == true {
		t.Fatal("IPv6.Parse() returned true for invalid input")
	}

}

func TestIPv6PrefixPortEdges(t *testing.T) {

	prefixCases := []struct {
		value  string
		valid  bool
		prefix uint8
	}{
		{"[fe80::1]/0", true, 0},
		{"[fe80::1]/64", true, 64},
		{"[fe80::1]/128", true, 128},
		{"[fe80::1]/129", false, 0},
		{"[fe80::1]/", false, 0},
		{"fe80::1/64", false, 0},
		{"[fe80::1]/64/128", false, 0},
		{"[fe80::1]/abc", false, 0},
	}

	for _, c := range prefixCases {

		ipv6, prefix := ParseIPv6AndPrefix(c.value)

		if c.valid {

			if ipv6 == nil {
				t.Errorf("ParseIPv6AndPrefix(%q) returned nil address", c.value)
			} else if prefix != c.prefix {
				t.Errorf("ParseIPv6AndPrefix(%q) prefix = %d, expected %d", c.value, prefix, c.prefix)
			}

		} else if ipv6 != nil || prefix != 0 {
			t.Errorf("ParseIPv6AndPrefix(%q) = %v/%d, expected nil/0", c.value, ipv6, prefix)
		}

		if IsIPv6AndPrefix(c.value) != c.valid {
			t.Errorf("IsIPv6AndPrefix(%q) = %v, expected %v", c.value, IsIPv6AndPrefix(c.value), c.valid)
		}

	}

	portCases := []struct {
		value string
		port  uint16
		valid bool
	}{
		{"[fe80::1]:0", 0, true},
		{"[fe80::1]:1338", 1338, true},
		{"[fe80::1]:65535", 65535, true},
		{"[fe80::1]:65536", 0, false},
		{"[fe80::1]:", 0, false},
		{"fe80::1:1338", 0, false},
		{"[fe80::1]:abc", 0, false},
	}

	for _, c := range portCases {

		ipv6, port := ParseIPv6AndPort(c.value)

		if c.valid {

			if ipv6 == nil {
				t.Errorf("ParseIPv6AndPort(%q) returned nil address", c.value)
			} else if port != c.port {
				t.Errorf("ParseIPv6AndPort(%q) port = %d, expected %d", c.value, port, c.port)
			}

		} else if ipv6 != nil || port != 0 {
			t.Errorf("ParseIPv6AndPort(%q) = %v/%d, expected nil/0", c.value, ipv6, port)
		}

		if IsIPv6AndPort(c.value) != c.valid {
			t.Errorf("IsIPv6AndPort(%q) = %v, expected %v", c.value, IsIPv6AndPort(c.value), c.valid)
		}

	}

}

func TestIPv6EqualHashPredicates(t *testing.T) {

	base := ToIPv6("2001:db8::1")
	same := ToIPv6("2001:0db8:0000:0000:0000:0000:0000:0001")
	different := ToIPv6("2001:db8::2")

	if base.Equal(same) == false || base.Equal(different) == true {
		t.Fatal("IPv6.Equal() failed")
	}

	if base.Hash() == "" || base.Hash() != same.Hash() {
		t.Fatal("IPv6.Hash() should be deterministic for equal values")
	}

	if base.Hash() == different.Hash() {
		t.Fatal("different IPv6 addresses should not share a hash")
	}

	if (IPv6{}).IsUnspecified() == false || base.IsUnspecified() == true {
		t.Fatal("IPv6.IsUnspecified() failed")
	}

	if ToIPv6("::1").IsUnspecified() == true {
		t.Fatal("::1 should not be unspecified")
	}

	if (IPv6{}).IsValid() == false || base.IsValid() == false || different.IsValid() == false {
		t.Fatal("all IPv6 values should be valid")
	}

}

func TestIPv6StringJSON(t *testing.T) {

	cases := []struct {
		value    string
		expected string
	}{
		{"::", UnspecifiedIPv6},
		{"::1", "0000:0000:0000:0000:0000:0000:0000:0001"},
		{"2001:db8::1", "2001:0db8:0000:0000:0000:0000:0000:0001"},
		{"[fe80::1]", "fe80:0000:0000:0000:0000:0000:0000:0001"},
	}

	for _, c := range cases {
		if result := ToIPv6(c.value).String(); result != c.expected {
			t.Errorf("ToIPv6(%q).String() = %q, expected %q", c.value, result, c.expected)
		}
	}

	original := ToIPv6("2001:db8::1")

	data, err := json.Marshal(original)

	if err != nil {
		t.Fatalf("Marshal() returned error: %v", err)
	}

	expected := `"2001:0db8:0000:0000:0000:0000:0000:0001"`

	if string(data) != expected {
		t.Errorf("Marshal() = %s, expected %s", string(data), expected)
	}

	var decoded IPv6

	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal() returned error: %v", err)
	}

	if decoded.Equal(original) == false {
		t.Errorf("Unmarshal() = %v, expected %v", decoded.String(), original.String())
	}

	var malformed IPv6

	if err := json.Unmarshal([]byte("123"), &malformed); err == nil {
		t.Fatal("Unmarshal() should fail for malformed JSON")
	}

	unchanged := original

	if err := json.Unmarshal([]byte(`"not an address"`), &unchanged); err != nil {
		t.Fatalf("Unmarshal() of an invalid string returned error: %v", err)
	}

	if unchanged.Equal(original) == false {
		t.Fatal("Unmarshal() of an invalid string should leave the address unchanged")
	}

}
