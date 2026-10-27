package types

import (
	"encoding/json"
	"strings"
	"testing"
)

func FuzzParseDomain(f *testing.F) {

	f.Add("evil.example")
	f.Add("evil.example.")
	f.Add("")

	f.Fuzz(func(t *testing.T, value string) {
		_ = IsDomain(value)
		_ = ParseDomain(value)
	})

}

func TestIsDomain(t *testing.T) {

	valid := []string{
		"evil.example",
		"sub.evil.example",
		"a-b.example",
		"a_b.example",
		"evil.example.",
		"  evil.example  ",
		"UPPER.EXAMPLE",
	}

	for _, value := range valid {

		if IsDomain(value) == false {
			t.Errorf("IsDomain(%q) = false, expected true", value)
		}

	}

	invalid := []string{
		"",
		"localhost",
		".example",
		"evil..example",
		"evil.c",
		"evil.exa mple",
		"evil.exa$mple",
		strings.Repeat("a", 63) + "." + strings.Repeat("b", 63) + "." + strings.Repeat("c", 63) + "." + strings.Repeat("d", 62),
	}

	for _, value := range invalid {

		if IsDomain(value) == true {
			t.Errorf("IsDomain(%q) = true, expected false", value)
		}

	}

}

func TestIsDomainLimits(t *testing.T) {

	label63 := ""
	for l := 0; l < 63; l++ {
		label63 += "a"
	}

	label64 := label63 + "a"

	if IsDomain(label63+".example") == false {
		t.Fatal("a 63-byte label should be valid")
	}

	if IsDomain(label64+".example") == true {
		t.Fatal("a 64-byte label should be invalid")
	}

	if IsDomain("evil.example.") == false {
		t.Fatal("a trailing root dot should be accepted")
	}

	domain := ParseDomain("evil.example.")

	if domain == nil || domain.String() != "evil.example" {
		t.Fatalf("ParseDomain did not normalize the trailing dot: %v", domain)
	}

}

func TestIsDomainAndPort(t *testing.T) {

	if IsDomainAndPort("evil.example:443") == false {
		t.Fatal("expected a valid domain and port")
	}

	domain, port := ParseDomainAndPort("evil.example:443")

	if domain == nil || domain.String() != "evil.example" || port != 443 {
		t.Fatalf("unexpected domain parse: %v %d", domain, port)
	}

	if IsDomainAndPort("evil.example:0") == false {
		t.Fatal("port 0 parses but is rejected later by SetPort")
	}

}

func TestIsDomainAndPortInvalid(t *testing.T) {

	invalid := []string{
		"evil.example",
		"evil.example:",
		"evil.example:65536",
		"evil.example:-1",
		"evil.example:abc",
		"evil.example:80:90",
		"localhost:80",
		":80",
		"",
	}

	for _, value := range invalid {

		if IsDomainAndPort(value) == true {
			t.Errorf("IsDomainAndPort(%q) = true, expected false", value)
		}

		domain, port := ParseDomainAndPort(value)

		if domain != nil || port != 0 {
			t.Errorf("ParseDomainAndPort(%q) = %v %d, expected nil 0", value, domain, port)
		}

	}

	if IsDomainAndPort("evil.example:65535") == false {
		t.Fatal("the maximum port should be accepted")
	}

	domain, port := ParseDomainAndPort("evil.example:65535")

	if domain == nil || domain.String() != "evil.example" || port != 65535 {
		t.Fatalf("unexpected boundary domain parse: %v %d", domain, port)
	}

}

func TestNewDomain(t *testing.T) {

	domain := NewDomain()

	if domain != Domain("") {
		t.Fatalf("NewDomain() = %q, expected an empty domain", domain)
	}

	if domain.IsValid() == true {
		t.Fatal("NewDomain() returned a valid domain")
	}

}

func TestParseDomain(t *testing.T) {

	domain := ParseDomain("  evil.example.  ")

	if domain == nil || domain.String() != "evil.example" {
		t.Fatalf("ParseDomain() = %v, expected %q", domain, "evil.example")
	}

	if ParseDomain("localhost") != nil {
		t.Fatal("ParseDomain() returned a value for an invalid domain")
	}

	if ParseDomain("") != nil {
		t.Fatal("ParseDomain() returned a value for an empty domain")
	}

}

func TestDomainParse(t *testing.T) {

	tests := []struct {
		input    string
		expected Domain
	}{
		{"evil.example", Domain("evil.example")},
		{"evil.example.", Domain("evil.example")},
		{"  evil.example  ", Domain("evil.example")},
		{"UPPER.EXAMPLE", Domain("UPPER.EXAMPLE")},
	}

	for _, test := range tests {

		var domain Domain

		if domain.Parse(test.input) == false {
			t.Errorf("Parse(%q) = false, expected true", test.input)
		} else if domain != test.expected {
			t.Errorf("Parse(%q) = %q, expected %q", test.input, domain, test.expected)
		}

	}

	invalid := []string{"", "localhost", ".example", "evil..example", "evil.c"}

	for _, value := range invalid {

		var domain Domain

		if domain.Parse(value) == true {
			t.Errorf("Parse(%q) = true, expected false", value)
		}

	}

}

func TestToDomain(t *testing.T) {

	if ToDomain("evil.example.") != Domain("evil.example") {
		t.Fatalf("ToDomain() = %q, expected %q", ToDomain("evil.example."), "evil.example")
	}

	if ToDomain("localhost") != Domain("") {
		t.Fatalf("ToDomain() = %q, expected zero value", ToDomain("localhost"))
	}

}

func TestDomainEqual(t *testing.T) {

	if Domain("evil.example").Equal(Domain("evil.example")) == false {
		t.Fatal("Equal() returned false for identical domains")
	}

	if Domain("evil.example").Equal(Domain("other.example")) == true {
		t.Fatal("Equal() returned true for different domains")
	}

}

func TestDomainHash(t *testing.T) {

	first := Domain("evil.example").Hash()
	second := Domain("evil.example").Hash()

	if first == "" {
		t.Fatal("Hash() returned an empty string")
	}

	if first != second {
		t.Fatalf("Hash() is not deterministic: %q != %q", first, second)
	}

	if first == Domain("other.example").Hash() {
		t.Fatal("Hash() returned the same value for different domains")
	}

}

func TestDomainIsValid(t *testing.T) {

	if Domain("evil.example").IsValid() == false {
		t.Fatal("valid domain reported invalid")
	}

	if Domain("invalid").IsValid() == true {
		t.Fatal("invalid domain reported valid")
	}

	if NewDomain().IsValid() == true {
		t.Fatal("empty domain reported valid")
	}

}

func TestDomainBytes(t *testing.T) {

	bytes := Domain("evil.example").Bytes()

	if string(bytes) != "evil.example" {
		t.Fatalf("Bytes() = %q, expected %q", string(bytes), "evil.example")
	}

}

func TestDomainScope(t *testing.T) {

	tests := []struct {
		input    Domain
		expected string
	}{
		{"1.10.in-addr.arpa", "private"},
		{"8.8.8.8.in-addr.arpa", "public"},
		{"x.b.e.f.ip6.arpa", "private"},
		{"8.8.8.8.ip6.arpa", "public"},
		{"host.local", "private"},
		{"sub.home.arpa", "private"},
		{"evil.example", "private"},
		{"evil.invalid", "private"},
		{"evil.localhost", "private"},
		{"evil.test", "private"},
		{"evil.alt", "private"},
		{"evil.com", "public"},
		{"localhost", "private"},
		{"", "private"},
	}

	for _, test := range tests {

		if result := test.input.Scope(); result != test.expected {
			t.Errorf("Domain(%q).Scope() = %q, expected %q", test.input, result, test.expected)
		}

	}

}

func TestDomainString(t *testing.T) {

	if Domain("evil.example").String() != "evil.example" {
		t.Fatalf("String() = %q, expected %q", Domain("evil.example").String(), "evil.example")
	}

	if NewDomain().String() != "" {
		t.Fatalf("String() = %q, expected an empty string", NewDomain().String())
	}

}

func TestDomainMarshalJSON(t *testing.T) {

	data, err := Domain("evil.example").MarshalJSON()

	if err != nil {
		t.Fatalf("MarshalJSON() returned an error: %v", err)
	}

	if string(data) != `"evil.example"` {
		t.Fatalf("MarshalJSON() = %s, expected %q", data, `"evil.example"`)
	}

	data, err = json.Marshal(NewDomain())

	if err != nil {
		t.Fatalf("json.Marshal() returned an error: %v", err)
	}

	if string(data) != `""` {
		t.Fatalf("json.Marshal() = %s, expected %q", data, `""`)
	}

}

func TestDomainUnmarshalJSON(t *testing.T) {

	var domain Domain

	if err := json.Unmarshal([]byte(`"evil.example."`), &domain); err != nil {
		t.Fatalf("UnmarshalJSON() returned an error: %v", err)
	}

	if domain != Domain("evil.example") {
		t.Fatalf("UnmarshalJSON() = %q, expected %q", domain, "evil.example")
	}

	var zero Domain

	if err := json.Unmarshal([]byte(`"invalid"`), &zero); err != nil {
		t.Fatalf("UnmarshalJSON() returned an error for an invalid value: %v", err)
	}

	if zero != Domain("") {
		t.Fatalf("UnmarshalJSON() changed the value for an invalid domain: %q", zero)
	}

	if err := zero.UnmarshalJSON([]byte("not-quoted")); err == nil {
		t.Fatal("UnmarshalJSON() accepted malformed JSON")
	}

}
