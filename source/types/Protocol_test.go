package types

import "testing"

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

}
