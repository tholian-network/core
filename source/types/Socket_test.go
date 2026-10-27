package types

import (
	"encoding/json"
	"testing"
)

func TestSocket(t *testing.T) {

	socket := NewSocket("1.3.3.7", 65535)

	if socket.Host != "1.3.3.7" || socket.Port != 65535 || socket.Type != AddressTypeIPv4 {
		t.Fatalf("unexpected socket: %+v", socket)
	}

	if socket.IsValid() == false {
		t.Fatal("IPv4 socket should be valid")
	}

	socket = NewSocket("evil.example", 443)

	if socket.Type != AddressTypeDomain || socket.Scope != "public" {
		t.Fatalf("unexpected domain socket: %+v", socket)
	}

	socket = NewSocket("*", 0)

	if socket.Host != HostWildcard || socket.Type != AddressTypeIPv4 {
		t.Fatalf("unexpected wildcard socket: %+v", socket)
	}

}

func TestToSocket(t *testing.T) {

	cases := []struct {
		value string
		host  string
		port  uint16
	}{
		{"1.3.3.7:1338", "1.3.3.7", 1338},
		{"[fe80::1]:1338", "fe80:0000:0000:0000:0000:0000:0000:0001", 1338},
		{"evil.example:1338", "evil.example", 1338},
	}

	for _, c := range cases {

		socket := ToSocket(c.value)

		if socket.Host != c.host || socket.Port != c.port {
			t.Errorf("ToSocket(%q) = %s:%d, expected %s:%d", c.value, socket.Host, socket.Port, c.host, c.port)
		}

	}

}

func TestSocketPredicates(t *testing.T) {

	unset := Socket{}

	if unset.IsUnset() == false || unset.IsValid() == true {
		t.Fatal("the zero socket should be unset and invalid")
	}

	wildcard := NewSocket("*", 1338)

	if wildcard.IsAny() == false || wildcard.HasHost() == true || wildcard.HasPort() == false {
		t.Fatalf("unexpected wildcard socket predicates: %+v", wildcard)
	}

	unspecified := NewSocket(UnspecifiedIPv4, 0)

	if unspecified.IsUnspecified() == false || unspecified.HasHost() == true {
		t.Fatalf("unexpected unspecified socket predicates: %+v", unspecified)
	}

	host_only := NewSocket("1.3.3.7", 0)

	if host_only.HasHost() == false || host_only.HasPort() == true {
		t.Fatalf("unexpected host-only socket predicates: %+v", host_only)
	}

	if host_only.IsValid() == false {
		t.Fatal("host-only IPv4 socket should be valid")
	}

	ipv6_unspecified := NewSocket(UnspecifiedIPv6, 0)

	if ipv6_unspecified.IsUnspecified() == false || ipv6_unspecified.IsValid() == false {
		t.Fatalf("unexpected IPv6 unspecified socket predicates: %+v", ipv6_unspecified)
	}

	invalid := Socket{Host: "", Type: AddressTypeIPv4}

	if invalid.IsValid() == true {
		t.Fatal("a socket without a host should be invalid")
	}

}

func TestSocketSetters(t *testing.T) {

	socket := NewSocket("1.3.3.7", 1338)

	socket.SetPort(443)
	socket.SetScope("private")

	if socket.Port != 443 {
		t.Errorf("Expected port to be %d, got %d", 443, socket.Port)
	}

	if socket.Scope != "private" {
		t.Errorf("Expected scope to be %s, got %s", "private", socket.Scope)
	}

	if socket.IsValid() == false {
		t.Fatal("socket should remain valid after setters")
	}

}

func TestSocketNewConstructor(t *testing.T) {

	socket := NewSocket("1.3.3.7", 0)

	if socket.Host != "1.3.3.7" || socket.Port != 0 {
		t.Fatalf("unexpected host-only socket: %+v", socket)
	}

	if socket.Type != AddressTypeIPv4 || socket.Scope != "public" {
		t.Errorf("unexpected host-only socket type/scope: %+v", socket)
	}

	socket = NewSocket("", 1338)

	if socket.Host != "" || socket.Port != 1338 {
		t.Fatalf("unexpected port-only socket: %+v", socket)
	}

	if socket.IsUnset() == true || socket.IsValid() == true {
		t.Fatalf("port-only socket should not be unset or valid: %+v", socket)
	}

}

func TestSocketParseFactory(t *testing.T) {

	cases := []struct {
		value string
		host  string
		port  uint16
		typ   string
	}{
		{"1.3.3.7:1338", "1.3.3.7", 1338, AddressTypeIPv4},
		{"[fe80::1]:1338", "fe80:0000:0000:0000:0000:0000:0000:0001", 1338, AddressTypeIPv6},
		{"evil.example:1338", "evil.example", 1338, AddressTypeDomain},
		{"1.3.3.7", "1.3.3.7", 0, AddressTypeIPv4},
		{"[fe80::1]", "fe80:0000:0000:0000:0000:0000:0000:0001", 0, AddressTypeIPv6},
		{"evil.example.", "evil.example", 0, AddressTypeDomain},
		{"*:1338", HostWildcard, 1338, AddressTypeIPv4},
	}

	for _, c := range cases {

		socket := ParseSocket(c.value)

		if socket == nil {
			t.Fatalf("ParseSocket(%q) returned nil", c.value)
		}

		if socket.Host != c.host || socket.Port != c.port || socket.Type != c.typ {
			t.Errorf("ParseSocket(%q) = %+v, expected %s:%d (%s)", c.value, *socket, c.host, c.port, c.typ)
		}

		if IsSocket(c.value) == false {
			t.Errorf("IsSocket(%q) = false, expected true", c.value)
		}

		toSocket := ToSocket(c.value)

		if toSocket.Equal(*socket) == false {
			t.Errorf("ToSocket(%q) does not match ParseSocket(%q)", c.value, c.value)
		}

	}

	for _, value := range []string{
		"",
		"localhost",
		"*",
		"1.3.3.7:notaport",
		"evil.example:70000",
	} {
		if socket := ParseSocket(value); socket != nil {
			t.Errorf("ParseSocket(%q) = %+v, expected nil", value, *socket)
		}

		if IsSocket(value) == true {
			t.Errorf("IsSocket(%q) = true, expected false", value)
		}
	}

	var socket Socket

	if socket.Parse("1.3.3.7:1338") == false {
		t.Fatal("Socket.Parse() returned false for valid input")
	}

	if socket.Host != "1.3.3.7" || socket.Port != 1338 {
		t.Errorf("Socket.Parse() = %+v", socket)
	}

	if socket.Parse("not a socket") == true {
		t.Fatal("Socket.Parse() returned true for invalid input")
	}

}

func TestSocketEqualAndHash(t *testing.T) {

	base := NewSocket("1.3.3.7", 1338)

	if base.Equal(NewSocket("1.3.3.7", 1338)) == false {
		t.Fatal("identical sockets should be equal")
	}

	if base.Hash() == "" || base.Hash() != NewSocket("1.3.3.7", 1338).Hash() {
		t.Fatal("Hash() should be deterministic for equal sockets")
	}

	differentHost := NewSocket("1.3.3.8", 1338)
	differentPort := NewSocket("1.3.3.7", 443)
	differentScope := base
	differentScope.Scope = "private"
	differentType := base
	differentType.Type = AddressTypeDomain

	if base.Equal(differentHost) || base.Equal(differentPort) || base.Equal(differentScope) || base.Equal(differentType) {
		t.Fatal("Equal() should compare host, port, scope and type")
	}

	if base.Hash() == differentHost.Hash() || base.Hash() == differentPort.Hash() || base.Hash() == differentScope.Hash() || base.Hash() == differentType.Hash() {
		t.Fatal("unequal sockets should not share a hash")
	}

}

func TestSocketPredicatesExtra(t *testing.T) {

	unset := Socket{Scope: "public"}

	if unset.IsUnset() == false {
		t.Fatal("a socket without host, port and type should be unset regardless of scope")
	}

	for _, host := range []string{Any, HostWildcard} {
		socket := Socket{Host: host}

		if socket.IsAny() == false || socket.IsWildcard() == false {
			t.Errorf("Socket{Host:%q} should be any and wildcard", host)
		}

		if socket.HasHost() == true {
			t.Errorf("Socket{Host:%q} should not have a host", host)
		}
	}

	domain := NewSocket("evil.example", 443)

	if domain.IsAny() || domain.IsWildcard() || domain.IsUnspecified() {
		t.Fatal("domain socket should be neither any, wildcard nor unspecified")
	}

	if domain.HasHost() == false || domain.HasPort() == false {
		t.Fatal("domain socket should have a host and port")
	}

	if domain.IsValid() == false {
		t.Fatal("domain socket should be valid")
	}

	invalidDomain := Socket{Host: "evil.example", Type: AddressTypeIPv4}

	if invalidDomain.IsValid() == true {
		t.Fatal("IPv4 socket with a domain host should be invalid")
	}

	wildcardIPv6 := Socket{Host: HostWildcard, Type: AddressTypeIPv6}

	if wildcardIPv6.IsValid() == false {
		t.Fatal("wildcard socket with an IPv6 type should be valid")
	}

	unspecifiedIPv6 := Socket{Host: UnspecifiedIPv4, Type: AddressTypeIPv6}

	if unspecifiedIPv6.IsValid() == false {
		t.Fatal("unspecified socket with an IPv6 type should be valid")
	}

	unknownType := Socket{Host: "", Port: 1338, Type: "unknown"}

	if unknownType.IsValid() == true {
		t.Fatal("socket with an unknown type should be invalid")
	}

}

func TestSocketSetHost(t *testing.T) {

	cases := []struct {
		value string
		host  string
		scope string
		typ   string
	}{
		{"1.3.3.7", "1.3.3.7", "public", AddressTypeIPv4},
		{"10.0.0.1", "10.0.0.1", "private", AddressTypeIPv4},
		{"2001:db8::1", "2001:0db8:0000:0000:0000:0000:0000:0001", "public", AddressTypeIPv6},
		{"fe80::1", "fe80:0000:0000:0000:0000:0000:0000:0001", "private", AddressTypeIPv6},
		{"evil.example.", "evil.example", "public", AddressTypeDomain},
		{Any, HostWildcard, "public", AddressTypeIPv4},
		{"all", HostWildcard, "public", AddressTypeIPv4},
		{HostWildcard, HostWildcard, "public", AddressTypeIPv4},
	}

	for _, c := range cases {

		socket := NewSocket("1.3.3.7", 1338)
		socket.SetHost(c.value)

		if socket.Host != c.host || socket.Scope != c.scope || socket.Type != c.typ {
			t.Errorf("SetHost(%q) = %+v, expected %s (%s/%s)", c.value, socket, c.host, c.scope, c.typ)
		}

	}

	socket := NewSocket("1.3.3.7", 1338)
	expected := socket

	socket.SetHost("not a host")

	if socket != expected {
		t.Fatalf("SetHost() with an invalid value should be a no-op: %+v", socket)
	}

}

func TestSocketSetPortAndScope(t *testing.T) {

	socket := NewSocket("1.3.3.7", 1338)

	socket.SetPort(0)

	if socket.Port != 1338 {
		t.Fatal("SetPort(0) should be a no-op")
	}

	socket.SetPort(443)

	if socket.Port != 443 {
		t.Fatal("SetPort(443) did not update the port")
	}

	socket.SetScope("public")

	if socket.Scope != "public" {
		t.Fatal("SetScope(public) did not update the scope")
	}

	socket.SetScope("private")

	if socket.Scope != "private" {
		t.Fatal("SetScope(private) did not update the scope")
	}

	socket.SetScope("unknown")

	if socket.Scope != "private" {
		t.Fatal("SetScope() with an invalid value should be a no-op")
	}

}

func TestSocketString(t *testing.T) {

	cases := []struct {
		socket   Socket
		expected string
	}{
		{Socket{}, "*:0"},
		{NewSocket("", 1338), "*:1338"},
		{NewSocket("1.3.3.7", 1338), "1.3.3.7:1338"},
		{NewSocket("evil.example", 443), "evil.example:443"},
		{NewSocket(HostWildcard, 1338), "*:1338"},
	}

	for _, c := range cases {
		if result := c.socket.String(); result != c.expected {
			t.Errorf("Socket.String() = %q, expected %q", result, c.expected)
		}
	}

}

func TestSocketJSON(t *testing.T) {

	cases := []struct {
		socket   Socket
		expected string
	}{
		{NewSocket("1.3.3.7", 1338), `"1.3.3.7:1338"`},
		{NewSocket("evil.example", 443), `"evil.example:443"`},
		{NewSocket(HostWildcard, 1338), `"*:1338"`},
	}

	for _, c := range cases {

		data, err := json.Marshal(c.socket)

		if err != nil {
			t.Fatalf("Marshal() returned error: %v", err)
		}

		if string(data) != c.expected {
			t.Errorf("Marshal() = %s, expected %s", string(data), c.expected)
		}

		var decoded Socket

		if err := json.Unmarshal(data, &decoded); err != nil {
			t.Fatalf("Unmarshal() returned error: %v", err)
		}

		if decoded.Equal(c.socket) == false {
			t.Errorf("Unmarshal() = %+v, expected %+v", decoded, c.socket)
		}

	}

	var malformed Socket

	if err := json.Unmarshal([]byte("123"), &malformed); err == nil {
		t.Fatal("Unmarshal() should fail for malformed JSON")
	}

	unchanged := NewSocket("1.3.3.7", 1338)
	expected := unchanged

	if err := json.Unmarshal([]byte(`"not a socket"`), &unchanged); err != nil {
		t.Fatalf("Unmarshal() of an invalid string returned error: %v", err)
	}

	if unchanged != expected {
		t.Fatal("Unmarshal() of an invalid string should leave the socket unchanged")
	}

}
