package types

import (
	"encoding/json"
	"testing"
)

func TestConnection(t *testing.T) {

	connection := NewConnection()

	if connection.IsValid() == true {
		t.Fatal("a fresh connection should not be valid")
	}

	connection.SetType("client")
	connection.SetSource(NewSocket("0.0.0.0", 0))
	connection.SetTarget(NewSocket("1.3.3.7", 1338))
	connection.SetProtocol(ProtocolTCP)

	if connection.IsValid() == false {
		t.Fatal("client connection should be valid")
	}

	if connection.Protocol != ProtocolTCP {
		t.Fatalf("unexpected protocol: %q", connection.Protocol)
	}

	if connection.Equal(connection) == false {
		t.Fatal("connection should be identical to itself")
	}

}

func TestConnectionIsValid(t *testing.T) {

	client := NewConnection()
	client.SetType(ConnectionTypeClient)
	client.SetTarget(NewSocket("1.3.3.7", 1338))

	if client.IsValid() == false {
		t.Fatal("client with a target should be valid")
	}

	server := NewConnection()
	server.SetType(ConnectionTypeServer)
	server.SetSource(NewSocket("1.3.3.7", 1338))

	if server.IsValid() == false {
		t.Fatal("server with a source should be valid")
	}

	peer := NewConnection()
	peer.SetType(ConnectionTypePeer)
	peer.SetSource(NewSocket("1.3.3.7", 1338))
	peer.SetTarget(NewSocket("1.3.3.7", 1338))

	if peer.IsValid() == false {
		t.Fatal("peer with both endpoints should be valid")
	}

	peer.Protocol = Protocol("unknown")

	if peer.IsValid() == true {
		t.Fatal("peer with an unknown protocol should be invalid")
	}

}

func TestConnectionNewAndFactories(t *testing.T) {

	connection := NewConnection()

	if connection.Source.Host != UnspecifiedIPv4 || connection.Source.Port != 0 {
		t.Fatalf("unexpected default source: %+v", connection.Source)
	}

	if connection.Target.Host != UnspecifiedIPv4 || connection.Target.Port != 0 {
		t.Fatalf("unexpected default target: %+v", connection.Target)
	}

	if connection.Protocol != Protocol(Any) {
		t.Fatalf("unexpected default protocol: %q", connection.Protocol)
	}

	if connection.Type != "" {
		t.Fatalf("unexpected default type: %q", connection.Type)
	}

	parsed := ParseConnection("[tcp] 1.3.3.7:1338 -> 1.3.3.7:443")

	if parsed == nil {
		t.Fatal("ParseConnection() returned nil for a valid client connection")
	}

	if parsed.Source.Host != "1.3.3.7" || parsed.Source.Port != 1338 {
		t.Errorf("unexpected parsed source: %+v", parsed.Source)
	}

	if parsed.Target.Host != "1.3.3.7" || parsed.Target.Port != 443 {
		t.Errorf("unexpected parsed target: %+v", parsed.Target)
	}

	if parsed.Protocol != ProtocolTCP || parsed.Type != ConnectionTypeClient {
		t.Errorf("unexpected parsed protocol/type: %q/%q", parsed.Protocol, parsed.Type)
	}

	if ParseConnection("") != nil || ParseConnection("not a connection") != nil {
		t.Fatal("ParseConnection() should return nil for invalid values")
	}

	invalidConnection := ToConnection("not a connection")

	if invalidConnection.Equal(Connection{}) == false {
		t.Fatal("ToConnection() should return the zero value for invalid input")
	}

}

func TestIsConnection(t *testing.T) {

	cases := []struct {
		value    string
		expected bool
	}{
		{"[tcp] 1.3.3.7:1338 -> 1.3.3.7:443", true},
		{"[udp] 1.3.3.7:1338 <- 1.3.3.7:443", true},
		{"[tcp] 1.3.3.7:1338 <-> 1.3.3.7:443", true},
		{"[tcp] *:1338 -> 1.3.3.7:443", true},
		{"[tcp] 1.3.3.7:* -> 1.3.3.7:443", true},
		{"[tcp] 1.3.3.7:1338 -> evil.example:443", true},
		{"[tcp] [fe80::1]:1338 <-> [fe80::1]:443", true},
		{"  [tcp] 1.3.3.7:1338 -> 1.3.3.7:443  ", true},
		{"", false},
		{"not a connection", false},
		{"[unknown] 1.3.3.7:1338 -> 1.3.3.7:443", false},
		{"[tcp] 1.3.3.7:1338 => 1.3.3.7:443", false},
		{"[tcp] 1.3.3.7:0 -> 1.3.3.7:443", false},
		{"[tcp] 1.3.3.7:1338 -> 1.3.3.7:65535", false},
		{"[tcp] evil.example:1338 -> 1.3.3.7:443", false},
		{"[tcp] 1.3.3.7:1338 ->", false},
	}

	for _, c := range cases {
		if result := IsConnection(c.value); result != c.expected {
			t.Errorf("IsConnection(%q) = %v, expected %v", c.value, result, c.expected)
		}
	}

}

func TestConnectionParseVariants(t *testing.T) {

	cases := []struct {
		value      string
		sourceHost string
		sourcePort uint16
		targetHost string
		targetPort uint16
		protocol   Protocol
		typ        string
	}{
		{"[tcp] 1.3.3.7:1338 -> 1.3.3.7:443", "1.3.3.7", 1338, "1.3.3.7", 443, ProtocolTCP, ConnectionTypeClient},
		{"[udp] 1.3.3.7:1338 <- 1.3.3.7:443", "1.3.3.7", 1338, "1.3.3.7", 443, ProtocolUDP, ConnectionTypeServer},
		{"[tcp] 1.3.3.7:1338 <-> 1.3.3.7:443", "1.3.3.7", 1338, "1.3.3.7", 443, ProtocolTCP, ConnectionTypePeer},
		{"[tcp] *:1338 -> 1.3.3.7:443", "0.0.0.0", 1338, "1.3.3.7", 443, ProtocolTCP, ConnectionTypeClient},
		{"[tcp] 1.3.3.7:1338 -> evil.example:443", "1.3.3.7", 1338, "evil.example", 443, ProtocolTCP, ConnectionTypeClient},
		{"[tcp] [fe80::1]:1338 <-> [2001:db8::1]:443", "fe80:0000:0000:0000:0000:0000:0000:0001", 1338, "2001:0db8:0000:0000:0000:0000:0000:0001", 443, ProtocolTCP, ConnectionTypePeer},
	}

	for _, c := range cases {

		connection := ParseConnection(c.value)

		if connection == nil {
			t.Fatalf("ParseConnection(%q) returned nil", c.value)
		}

		if connection.Source.Host != c.sourceHost || connection.Source.Port != c.sourcePort {
			t.Errorf("ParseConnection(%q) source = %+v, expected %s:%d", c.value, connection.Source, c.sourceHost, c.sourcePort)
		}

		if connection.Target.Host != c.targetHost || connection.Target.Port != c.targetPort {
			t.Errorf("ParseConnection(%q) target = %+v, expected %s:%d", c.value, connection.Target, c.targetHost, c.targetPort)
		}

		if connection.Protocol != c.protocol || connection.Type != c.typ {
			t.Errorf("ParseConnection(%q) protocol/type = %q/%q, expected %q/%q", c.value, connection.Protocol, connection.Type, c.protocol, c.typ)
		}

	}

	for _, value := range []string{
		"",
		"not a connection",
		"[unknown] 1.3.3.7:1338 -> 1.3.3.7:443",
		"[tcp] 1.3.3.7:0 -> 1.3.3.7:443",
		"[tcp] 1.3.3.7:1338 ->> 1.3.3.7:443",
		"[tcp] 1.3.3.7:1338 -> 1.3.3.7:65535",
	} {
		if connection := ParseConnection(value); connection != nil {
			t.Errorf("ParseConnection(%q) = %v, expected nil", value, connection.String())
		}
	}

}

func TestConnectionParseLeavesInvalidUnchanged(t *testing.T) {

	connection := ToConnection("[tcp] 1.3.3.7:1338 -> 1.3.3.7:443")
	expected := connection

	if connection.Parse("not a connection") != false {
		t.Fatal("Parse() should return false for invalid input")
	}

	if connection.Equal(expected) == false {
		t.Fatalf("Parse() modified the connection on invalid input: %+v", connection)
	}

}

func TestConnectionPredicates(t *testing.T) {

	client := ToConnection("[tcp] 1.3.3.7:1338 -> 1.3.3.7:443")
	server := ToConnection("[tcp] 1.3.3.7:1338 <- 1.3.3.7:443")
	peer := ToConnection("[tcp] 1.3.3.7:1338 <-> 1.3.3.7:443")

	if client.IsClient() != true || client.IsServer() != false || client.IsPeer() != false {
		t.Error("unexpected client connection predicates")
	}

	if server.IsClient() != false || server.IsServer() != true || server.IsPeer() != false {
		t.Error("unexpected server connection predicates")
	}

	if peer.IsClient() != false || peer.IsServer() != false || peer.IsPeer() != true {
		t.Error("unexpected peer connection predicates")
	}

}

func TestConnectionEqualAndHash(t *testing.T) {

	base := ToConnection("[tcp] 1.3.3.7:1338 -> 1.3.3.7:443")
	same := ToConnection("[tcp] 1.3.3.7:1338 -> 1.3.3.7:443")

	if base.Equal(same) == false {
		t.Fatal("identical connections should be equal")
	}

	if base.Hash() == "" || base.Hash() != same.Hash() {
		t.Fatal("Hash() should be deterministic for equal connections")
	}

	scoped := same
	scoped.Source.Scope = "private"

	if base.Equal(scoped) == false {
		t.Fatal("Equal() should ignore the socket scope")
	}

	differentHost := same
	differentHost.Source.Host = "1.3.3.8"

	differentPort := same
	differentPort.Target.Port = 444

	differentProtocol := same
	differentProtocol.Protocol = ProtocolUDP

	differentType := same
	differentType.Type = ConnectionTypeServer

	for name, other := range map[string]Connection{
		"host":     differentHost,
		"port":     differentPort,
		"protocol": differentProtocol,
		"type":     differentType,
	} {
		if base.Equal(other) {
			t.Errorf("connections with a different %s should not be equal", name)
		}
		if base.Hash() == other.Hash() {
			t.Errorf("connections with a different %s should not have the same hash", name)
		}
	}

}

func TestConnectionIsValidRules(t *testing.T) {

	client := NewConnection()
	client.SetType(ConnectionTypeClient)
	client.SetProtocol(ProtocolTCP)

	client.SetTarget(NewSocket("*", 1338))

	if client.IsValid() == false {
		t.Fatal("client with a wildcard target and a port should be valid")
	}

	client.SetTarget(NewSocket(UnspecifiedIPv4, 1338))

	if client.IsValid() == true {
		t.Fatal("client with an unspecified target should be invalid")
	}

	client.SetTarget(NewSocket("1.3.3.7", 0))

	if client.IsValid() == false {
		t.Fatal("client with a target host but no port should be valid")
	}

	client.Target = Socket{}

	if client.IsValid() == true {
		t.Fatal("client with an unset target should be invalid")
	}

	peer := NewConnection()
	peer.SetType(ConnectionTypePeer)
	peer.SetProtocol(ProtocolTCP)
	peer.SetSource(NewSocket("1.3.3.7", 1338))
	peer.SetTarget(NewSocket("*", 443))

	if peer.IsValid() == true {
		t.Fatal("peer with a wildcard endpoint should be invalid")
	}

	peer.SetTarget(NewSocket("1.3.3.7", 443))

	if peer.IsValid() == false {
		t.Fatal("peer with two concrete endpoints should be valid")
	}

	unknown := NewConnection()
	unknown.Type = "unknown"
	unknown.SetProtocol(ProtocolTCP)

	if unknown.IsValid() == true {
		t.Fatal("connection with an unknown type should be invalid")
	}

}

func TestConnectionSetters(t *testing.T) {

	connection := NewConnection()
	expected := connection

	connection.SetSource(Socket{})
	connection.SetTarget(Socket{})
	connection.SetProtocol(Protocol("unknown"))
	connection.SetType("unknown")

	if connection.Source != expected.Source || connection.Target != expected.Target {
		t.Fatal("invalid socket setters should be no-ops")
	}

	if connection.Protocol != expected.Protocol || connection.Type != expected.Type {
		t.Fatal("invalid protocol/type setters should be no-ops")
	}

	source := NewSocket("1.3.3.7", 1338)
	target := NewSocket("1.3.3.8", 443)

	connection.SetSource(source)
	connection.SetTarget(target)
	connection.SetProtocol(ProtocolTCP)
	connection.SetType(ConnectionTypePeer)

	if connection.Source != source || connection.Target != target {
		t.Fatal("valid socket setters did not update the connection")
	}

	if connection.Protocol != ProtocolTCP || connection.Type != ConnectionTypePeer {
		t.Fatal("valid protocol/type setters did not update the connection")
	}

}

func TestConnectionString(t *testing.T) {

	connection := NewConnection()

	if connection.String() != "" {
		t.Fatalf("empty typed connection String() = %q, expected an empty string", connection.String())
	}

	connection.SetType(ConnectionTypeClient)

	expected := "[any] 0.0.0.0:0 -> 0.0.0.0:0"

	if connection.String() != expected {
		t.Errorf("Connection.String() = %q, expected %q", connection.String(), expected)
	}

	connection.SetType(ConnectionTypeServer)

	expected = "[any] 0.0.0.0:0 <- 0.0.0.0:0"

	if connection.String() != expected {
		t.Errorf("Connection.String() = %q, expected %q", connection.String(), expected)
	}

	connection.SetType(ConnectionTypePeer)

	expected = "[any] 0.0.0.0:0 <-> 0.0.0.0:0"

	if connection.String() != expected {
		t.Errorf("Connection.String() = %q, expected %q", connection.String(), expected)
	}

	parsed := ToConnection("[tcp] 1.3.3.7:1338 -> 1.3.3.7:443")

	if parsed.String() != "[tcp] 1.3.3.7:1338 -> 1.3.3.7:443" {
		t.Errorf("Connection.String() = %q", parsed.String())
	}

	parsed.Type = "unknown"

	if parsed.String() != "[tcp] 1.3.3.7:1338 - 1.3.3.7:443" {
		t.Errorf("Connection.String() with unknown type = %q", parsed.String())
	}

}

func TestConnectionJSON(t *testing.T) {

	connection := ToConnection("[tcp] 1.3.3.7:1338 -> 1.3.3.7:443")

	data, err := json.Marshal(connection)

	if err != nil {
		t.Fatalf("Marshal() returned error: %v", err)
	}

	expected := `"[tcp] 1.3.3.7:1338 -\u003e 1.3.3.7:443"`

	if string(data) != expected {
		t.Errorf("Marshal() = %s, expected %s", string(data), expected)
	}

	var decoded Connection

	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal() returned error: %v", err)
	}

	if decoded.Equal(connection) == false {
		t.Errorf("Unmarshal() = %+v, expected %+v", decoded, connection)
	}

	var malformed Connection

	if err := json.Unmarshal([]byte("123"), &malformed); err == nil {
		t.Fatal("Unmarshal() should fail for malformed JSON")
	}

	unchanged := connection

	if err := json.Unmarshal([]byte(`"not a connection"`), &unchanged); err != nil {
		t.Fatalf("Unmarshal() of an invalid string returned error: %v", err)
	}

	if unchanged.Equal(connection) == false {
		t.Fatal("Unmarshal() of an invalid string should leave the connection unchanged")
	}

}
