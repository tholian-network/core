package types

import (
	"encoding/json"
	"testing"
)

func TestProtocolIsValid(t *testing.T) {

	if Protocol(Any).IsValid() == false || ProtocolTCP.IsValid() == false || ProtocolUDP.IsValid() == false {
		t.Fatal("known protocols should be valid")
	}

	if Protocol("unknown").IsValid() == true {
		t.Fatal("unknown protocol should be invalid")
	}

	if ProtocolTCP.IsAny() == true || Protocol(Any).IsAny() == false {
		t.Fatal("unexpected protocol any predicate")
	}

}

func TestIsProtocol(t *testing.T) {

	valid := []Protocol{
		Protocol(Any),
		ProtocolDNS,
		ProtocolDNSoverTLS,
		ProtocolHTTPS,
		ProtocolHTTP,
		ProtocolICMP,
		ProtocolSSH,
		ProtocolSOCKS,
		ProtocolTCP,
		ProtocolUDP,
		ProtocolWHOIS,
	}

	for _, value := range valid {

		if IsProtocol(string(value)) == false {
			t.Errorf("IsProtocol(%q) = false, expected true", value)
		}

	}

	if IsProtocol("unknown") == true {
		t.Fatal("IsProtocol() accepted an unknown protocol")
	}

	if IsProtocol("") == true {
		t.Fatal("IsProtocol() accepted an empty protocol")
	}

}

func TestNewProtocol(t *testing.T) {

	protocol := NewProtocol()

	if protocol != Protocol(Any) {
		t.Fatalf("NewProtocol() = %q, expected %q", protocol, Any)
	}

	if protocol.IsAny() == false {
		t.Fatal("NewProtocol() did not return the any protocol")
	}

}

func TestProtocolParse(t *testing.T) {

	tests := []struct {
		input    string
		expected Protocol
	}{
		{"*", Protocol(Any)},
		{"all", Protocol(Any)},
		{"any", Protocol(Any)},
		{"https", ProtocolHTTPS},
		{"http", ProtocolHTTP},
		{"dns", ProtocolDNS},
		{"dns-over-tls", ProtocolDNSoverTLS},
		{"dot", ProtocolDNSoverTLS},
		{"icmp", ProtocolICMP},
		{"ssh", ProtocolSSH},
		{"socks", ProtocolSOCKS},
		{"socks4", ProtocolSOCKS},
		{"socks5", ProtocolSOCKS},
		{"whois", ProtocolWHOIS},
		{"tcp", ProtocolTCP},
		{"tcp4", ProtocolTCP},
		{"tcp6", ProtocolTCP},
		{"udp", ProtocolUDP},
		{"udp4", ProtocolUDP},
		{"udp6", ProtocolUDP},
	}

	for _, test := range tests {

		var protocol Protocol

		if protocol.Parse(test.input) == false {
			t.Errorf("Parse(%q) = false, expected true", test.input)
		} else if protocol != test.expected {
			t.Errorf("Parse(%q) = %q, expected %q", test.input, protocol, test.expected)
		}

	}

	invalid := []string{"", "unknown", "HTTP", "HTTPS", "tls"}

	for _, value := range invalid {

		var protocol Protocol

		if protocol.Parse(value) == true {
			t.Errorf("Parse(%q) = true, expected false", value)
		}

	}

}

func TestParseProtocol(t *testing.T) {

	protocol := ParseProtocol("dot")

	if protocol == nil || *protocol != ProtocolDNSoverTLS {
		t.Fatalf("ParseProtocol() = %v, expected %q", protocol, ProtocolDNSoverTLS)
	}

	if ParseProtocol("unknown") != nil {
		t.Fatal("ParseProtocol() returned a value for an invalid protocol")
	}

	if ParseProtocol("") != nil {
		t.Fatal("ParseProtocol() returned a value for an empty protocol")
	}

}

func TestToProtocol(t *testing.T) {

	if ToProtocol("tcp4") != ProtocolTCP {
		t.Fatalf("ToProtocol(%q) = %q, expected %q", "tcp4", ToProtocol("tcp4"), ProtocolTCP)
	}

	if ToProtocol("udp6") != ProtocolUDP {
		t.Fatalf("ToProtocol(%q) = %q, expected %q", "udp6", ToProtocol("udp6"), ProtocolUDP)
	}

	if ToProtocol("unknown") != Protocol("") {
		t.Fatalf("ToProtocol() = %q, expected zero value", ToProtocol("unknown"))
	}

}

func TestProtocolBytes(t *testing.T) {

	bytes := ProtocolHTTPS.Bytes()

	if string(bytes) != "https" {
		t.Fatalf("Bytes() = %q, expected %q", string(bytes), "https")
	}

}

func TestProtocolEqual(t *testing.T) {

	if ProtocolTCP.Equal(ProtocolTCP) == false {
		t.Fatal("Equal() returned false for identical protocols")
	}

	if ProtocolTCP.Equal(ProtocolUDP) == true {
		t.Fatal("Equal() returned true for different protocols")
	}

}

func TestProtocolHash(t *testing.T) {

	first := ProtocolTCP.Hash()
	second := ProtocolTCP.Hash()

	if first == "" {
		t.Fatal("Hash() returned an empty string")
	}

	if first != second {
		t.Fatalf("Hash() is not deterministic: %q != %q", first, second)
	}

	if first == ProtocolUDP.Hash() {
		t.Fatal("Hash() returned the same value for different protocols")
	}

}

func TestProtocolIsAny(t *testing.T) {

	if Protocol(Any).IsAny() == false {
		t.Fatal("the any protocol was not reported as any")
	}

	if Protocol("").IsAny() == false {
		t.Fatal("an empty protocol was not reported as any")
	}

	if ProtocolTCP.IsAny() == true {
		t.Fatal("TCP was reported as any")
	}

}

func TestProtocolString(t *testing.T) {

	if ProtocolHTTPS.String() != "https" {
		t.Fatalf("String() = %q, expected %q", ProtocolHTTPS.String(), "https")
	}

	if NewProtocol().String() != Any {
		t.Fatalf("String() = %q, expected %q", NewProtocol().String(), Any)
	}

}

func TestProtocolMarshalJSON(t *testing.T) {

	data, err := ProtocolHTTPS.MarshalJSON()

	if err != nil {
		t.Fatalf("MarshalJSON() returned an error: %v", err)
	}

	if string(data) != `"https"` {
		t.Fatalf("MarshalJSON() = %s, expected %q", data, `"https"`)
	}

	data, err = json.Marshal(NewProtocol())

	if err != nil {
		t.Fatalf("json.Marshal() returned an error: %v", err)
	}

	if string(data) != `"any"` {
		t.Fatalf("json.Marshal() = %s, expected %q", data, `"any"`)
	}

}

func TestProtocolUnmarshalJSON(t *testing.T) {

	var protocol Protocol

	if err := json.Unmarshal([]byte(`"dot"`), &protocol); err != nil {
		t.Fatalf("UnmarshalJSON() returned an error: %v", err)
	}

	if protocol != ProtocolDNSoverTLS {
		t.Fatalf("UnmarshalJSON() = %q, expected %q", protocol, ProtocolDNSoverTLS)
	}

	var zero Protocol

	if err := json.Unmarshal([]byte(`"unknown"`), &zero); err != nil {
		t.Fatalf("UnmarshalJSON() returned an error for an invalid value: %v", err)
	}

	if zero != Protocol("") {
		t.Fatalf("UnmarshalJSON() changed the value for an invalid protocol: %q", zero)
	}

	if err := zero.UnmarshalJSON([]byte("not-quoted")); err == nil {
		t.Fatal("UnmarshalJSON() accepted malformed JSON")
	}

}
