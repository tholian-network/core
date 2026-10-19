package types

import "encoding/json"
import "reflect"
import "testing"

func assertJSONRoundTrip(t *testing.T, value any, expected string) {

	t.Helper()

	data, err := json.Marshal(value)

	if err != nil {
		t.Fatalf("Marshal() returned error: %v", err)
	}

	if string(data) != expected {
		t.Errorf("Marshal() = %s, expected %s", string(data), expected)
	}

	target := reflect.New(reflect.TypeOf(value))

	if err := json.Unmarshal(data, target.Interface()); err != nil {
		t.Fatalf("Unmarshal() returned error: %v", err)
	}

	if reflect.DeepEqual(value, target.Elem().Interface()) == false {
		t.Errorf("Unmarshal() = %#v, expected %#v", target.Elem().Interface(), value)
	}

}

func TestJSONRoundTrip(t *testing.T) {

	t.Run("Time", func(t *testing.T) {
		assertJSONRoundTrip(t, ToTime("11:59:59"), `"11:59:59"`)
	})

	t.Run("Datetime", func(t *testing.T) {
		assertJSONRoundTrip(t, ToDatetime("2026-10-19T11:59:59Z"), `"2026-10-19 11:59:59"`)
	})

	t.Run("IPv4", func(t *testing.T) {

		ipv4 := ParseIPv4("1.3.3.7")

		if ipv4 == nil {
			t.Fatal("ParseIPv4() returned nil")
		}

		assertJSONRoundTrip(t, *ipv4, `"1.3.3.7"`)

	})

	t.Run("IPv6", func(t *testing.T) {

		ipv6 := ParseIPv6("2001:db8::1")

		if ipv6 == nil {
			t.Fatal("ParseIPv6() returned nil")
		}

		assertJSONRoundTrip(t, *ipv6, `"2001:0db8:0000:0000:0000:0000:0000:0001"`)

	})

	t.Run("Socket", func(t *testing.T) {
		assertJSONRoundTrip(t, NewSocket("1.3.3.7", 1338), `"1.3.3.7:1338"`)
	})

	t.Run("Connection", func(t *testing.T) {

		connection := NewConnection()

		connection.SetType(ConnectionTypeClient)
		connection.SetSource(NewSocket("1.3.3.7", 1338))
		connection.SetTarget(NewSocket("1.3.3.7", 443))
		connection.SetProtocol(ProtocolTCP)

		assertJSONRoundTrip(t, connection, `"[tcp] 1.3.3.7:1338 -\u003e 1.3.3.7:443"`)

	})

	t.Run("Architecture", func(t *testing.T) {
		assertJSONRoundTrip(t, ArchitectureX86, `"x86"`)
	})

	t.Run("ASN", func(t *testing.T) {
		assertJSONRoundTrip(t, ASN(1337), `"AS1337"`)
	})

	t.Run("Domain", func(t *testing.T) {
		assertJSONRoundTrip(t, ToDomain("example.com"), `"example.com"`)
	})

	t.Run("Protocol", func(t *testing.T) {
		assertJSONRoundTrip(t, ProtocolTCP, `"tcp"`)
	})

}
