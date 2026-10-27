package types

import (
	"encoding/json"
	"testing"
)

func FuzzParseIPv4(f *testing.F) {

	f.Add("1.3.3.7")
	f.Add("1.3.3.256")
	f.Add("")

	f.Fuzz(func(t *testing.T, value string) {
		_ = IsIPv4(value)
		_ = ParseIPv4(value)
	})

}

func TestIsIPv4(t *testing.T) {

	cases := map[string]bool{
		"1.3.3.7":         true,
		"0.0.0.0":         true,
		"255.255.255.255": true,
		"1.3.3":           false,
		"1.3.3.256":       false,
		"1.3.3.7/24":      false,
		"example.com":     false,
	}

	for input, expected := range cases {

		if result := IsIPv4(input); result != expected {
			t.Errorf("IsIPv4(%q) = %v, expected %v", input, result, expected)
		}

	}

}

func TestIsIPv4AndPrefix(t *testing.T) {

	cases := map[string]bool{
		"1.3.3.0/24": true,
		"1.3.3.7/32": true,
		"1.3.3.7/8":  true,
		"1.3.3.7/7":  true,
		"1.3.3.7/0":  true,
		"1.3.3.7/33": false,
		"1.3.3.7":    false,
	}

	for input, expected := range cases {

		if result := IsIPv4AndPrefix(input); result != expected {
			t.Errorf("IsIPv4AndPrefix(%q) = %v, expected %v", input, result, expected)
		}

	}

}

func TestIsIPv4AndPort(t *testing.T) {

	cases := map[string]bool{
		"1.3.3.7:1338": true,
		"1.3.3.7:":     false,
		"1.3.3.7":      false,
	}

	for input, expected := range cases {

		if result := IsIPv4AndPort(input); result != expected {
			t.Errorf("IsIPv4AndPort(%q) = %v, expected %v", input, result, expected)
		}

	}

}

func TestIPv4Bytes(t *testing.T) {

	ipv4 := ParseIPv4("1.3.3.7")

	if ipv4 == nil {
		t.Fatalf("ParseIPv4 returned nil")
	}

	masked := ipv4.Bytes(24)

	if len(masked) != 4 {
		t.Fatalf("IPv4.Bytes(24) length = %d, expected 4", len(masked))
	}

	if masked[0] != 1 || masked[1] != 3 || masked[2] != 3 || masked[3] != 0 {
		t.Errorf("IPv4.Bytes(24) = %v, expected [1 3 3 0]", masked)
	}

}

func TestParseIPv4NonOctetPrefix(t *testing.T) {

	ipv4 := ParseIPv4("1.3.255.7")

	if ipv4 == nil {
		t.Fatal("ParseIPv4 returned nil")
	}

	cases := []struct {
		prefix   uint8
		expected []byte
	}{
		{0, []byte{0, 0, 0, 0}},
		{1, []byte{0, 0, 0, 0}},
		{12, []byte{1, 0, 0, 0}},
		{17, []byte{1, 3, 128, 0}},
		{20, []byte{1, 3, 240, 0}},
		{24, []byte{1, 3, 255, 0}},
		{31, []byte{1, 3, 255, 6}},
		{32, []byte{1, 3, 255, 7}},
	}

	for _, c := range cases {

		result := ipv4.Bytes(c.prefix)

		if len(result) != len(c.expected) {
			t.Fatalf("IPv4.Bytes(%d) length = %d", c.prefix, len(result))
		}

		for b := 0; b < len(result); b++ {
			if result[b] != c.expected[b] {
				t.Errorf("IPv4.Bytes(%d) = %v, expected %v", c.prefix, result, c.expected)
				break
			}
		}

	}

}

func TestIPv4FactoryAndParse(t *testing.T) {

	zero := NewIPv4()

	if zero.Equal(IPv4{}) == false || zero.IsUnspecified() == false {
		t.Fatal("NewIPv4() should return the unspecified address")
	}

	if zero.IsValid() == false || zero.String() != UnspecifiedIPv4 {
		t.Fatalf("unexpected zero IPv4: %q valid=%v", zero.String(), zero.IsValid())
	}

	cases := []struct {
		value    string
		expected string
	}{
		{"1.3.3.7", "1.3.3.7"},
		{"0.0.0.0", "0.0.0.0"},
		{"255.255.255.255", "255.255.255.255"},
		{"01.02.03.004", "1.2.3.4"},
	}

	for _, c := range cases {

		ipv4 := ParseIPv4(c.value)

		if ipv4 == nil {
			t.Fatalf("ParseIPv4(%q) returned nil", c.value)
		}

		if ipv4.String() != c.expected {
			t.Errorf("ParseIPv4(%q).String() = %q, expected %q", c.value, ipv4.String(), c.expected)
		}

		if ToIPv4(c.value).Equal(*ipv4) == false {
			t.Errorf("ToIPv4(%q) does not match ParseIPv4(%q)", c.value, c.value)
		}

	}

	for _, value := range []string{"", "1.3.3", "1.3.3.256", "1.3.3.7.8", "a.b.c.d", " 1.3.3.7"} {
		if ipv4 := ParseIPv4(value); ipv4 != nil {
			t.Errorf("ParseIPv4(%q) = %v, expected nil", value, ipv4.String())
		}
	}

	if ToIPv4("not an address").Equal(IPv4{}) == false {
		t.Fatal("ToIPv4() should return the zero value for invalid input")
	}

	var ipv4 IPv4

	if ipv4.Parse("1.3.3.7") == false || ipv4.Equal(ToIPv4("1.3.3.7")) == false {
		t.Fatal("IPv4.Parse() failed for valid input")
	}

	if ipv4.Parse("not an address") == true {
		t.Fatal("IPv4.Parse() returned true for invalid input")
	}

}

func TestIPv4AndPrefixPortEdges(t *testing.T) {

	prefixCases := []struct {
		value  string
		valid  bool
		prefix uint8
	}{
		{"1.3.3.0/0", true, 0},
		{"1.3.3.7/24", true, 24},
		{"1.3.3.7/32", true, 32},
		{"1.3.3.7/33", false, 0},
		{"1.3.3.7/", false, 0},
		{"1.3.3.7", false, 0},
		{"1.3.3.7/24/25", false, 0},
		{"1.3.3.256/24", false, 0},
		{"1.3.3.7/abc", false, 0},
	}

	for _, c := range prefixCases {

		ipv4, prefix := ParseIPv4AndPrefix(c.value)

		if c.valid {

			if ipv4 == nil {
				t.Errorf("ParseIPv4AndPrefix(%q) returned nil address", c.value)
			} else if prefix != c.prefix {
				t.Errorf("ParseIPv4AndPrefix(%q) prefix = %d, expected %d", c.value, prefix, c.prefix)
			}

		} else if ipv4 != nil || prefix != 0 {
			t.Errorf("ParseIPv4AndPrefix(%q) = %v/%d, expected nil/0", c.value, ipv4, prefix)
		}

		if IsIPv4AndPrefix(c.value) != c.valid {
			t.Errorf("IsIPv4AndPrefix(%q) = %v, expected %v", c.value, IsIPv4AndPrefix(c.value), c.valid)
		}

	}

	portCases := []struct {
		value string
		port  uint16
		valid bool
	}{
		{"1.3.3.7:0", 0, true},
		{"1.3.3.7:1338", 1338, true},
		{"1.3.3.7:65535", 65535, true},
		{"1.3.3.7:65536", 0, false},
		{"1.3.3.7:", 0, false},
		{"1.3.3.7", 0, false},
		{"1.3.3.7:13:38", 0, false},
		{"1.3.3.7:abc", 0, false},
	}

	for _, c := range portCases {

		ipv4, port := ParseIPv4AndPort(c.value)

		if c.valid {

			if ipv4 == nil {
				t.Errorf("ParseIPv4AndPort(%q) returned nil address", c.value)
			} else if port != c.port {
				t.Errorf("ParseIPv4AndPort(%q) port = %d, expected %d", c.value, port, c.port)
			}

		} else if ipv4 != nil || port != 0 {
			t.Errorf("ParseIPv4AndPort(%q) = %v/%d, expected nil/0", c.value, ipv4, port)
		}

		if IsIPv4AndPort(c.value) != c.valid {
			t.Errorf("IsIPv4AndPort(%q) = %v, expected %v", c.value, IsIPv4AndPort(c.value), c.valid)
		}

	}

}

func TestIPv4EqualHashPredicates(t *testing.T) {

	base := ToIPv4("1.3.3.7")
	same := ToIPv4("1.3.3.7")
	different := ToIPv4("1.3.3.8")

	if base.Equal(same) == false || base.Equal(different) == true {
		t.Fatal("IPv4.Equal() failed")
	}

	if base.Hash() == "" || base.Hash() != same.Hash() {
		t.Fatal("IPv4.Hash() should be deterministic for equal values")
	}

	if base.Hash() == different.Hash() {
		t.Fatal("different IPv4 addresses should not share a hash")
	}

	if (IPv4{}).IsUnspecified() == false || base.IsUnspecified() == true {
		t.Fatal("IPv4.IsUnspecified() failed")
	}

	if ParseIPv4("0.0.0.0").IsUnspecified() == false {
		t.Fatal("0.0.0.0 should be unspecified")
	}

	if (IPv4{}).IsValid() == false || base.IsValid() == false {
		t.Fatal("all in-range IPv4 values should be valid")
	}

}

func TestIPv4Scope(t *testing.T) {

	private := []string{
		"0.0.0.0",
		"127.0.0.1",
		"10.0.0.1",
		"172.16.0.1",
		"172.31.255.255",
		"192.168.1.1",
		"169.254.1.1",
		"100.64.0.1",
		"100.128.0.1",
		"100.192.0.1",
		"192.0.2.1",
		"192.88.99.1",
		"198.0.0.1",
		"198.51.100.1",
		"198.254.0.1",
		"203.0.113.1",
	}

	for _, value := range private {
		if scope := ToIPv4(value).Scope(); scope != "private" {
			t.Errorf("ToIPv4(%q).Scope() = %q, expected private", value, scope)
		}
	}

	public := []string{
		"1.3.3.7",
		"8.8.8.8",
		"11.0.0.1",
		"100.63.0.1",
		"100.127.0.1",
		"172.15.0.1",
		"172.32.0.1",
		"192.169.1.1",
		"198.1.0.1",
		"198.3.0.1",
		"199.0.0.1",
	}

	for _, value := range public {
		if scope := ToIPv4(value).Scope(); scope != "public" {
			t.Errorf("ToIPv4(%q).Scope() = %q, expected public", value, scope)
		}
	}

}

func TestIPv4StringJSON(t *testing.T) {

	cases := []struct {
		value    string
		expected string
	}{
		{"1.3.3.7", "1.3.3.7"},
		{"0.0.0.0", "0.0.0.0"},
		{"255.255.255.255", "255.255.255.255"},
		{"001.002.003.004", "1.2.3.4"},
	}

	for _, c := range cases {
		if result := ToIPv4(c.value).String(); result != c.expected {
			t.Errorf("ToIPv4(%q).String() = %q, expected %q", c.value, result, c.expected)
		}
	}

	original := ToIPv4("1.3.3.7")

	data, err := json.Marshal(original)

	if err != nil {
		t.Fatalf("Marshal() returned error: %v", err)
	}

	if string(data) != `"1.3.3.7"` {
		t.Errorf("Marshal() = %s, expected \"1.3.3.7\"", string(data))
	}

	var decoded IPv4

	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal() returned error: %v", err)
	}

	if decoded.Equal(original) == false {
		t.Errorf("Unmarshal() = %v, expected %v", decoded.String(), original.String())
	}

	var malformed IPv4

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
