package types

import "testing"

func TestNewConstructors(t *testing.T) {

	if NewArchitecture().Equal(Architecture(Any)) == false || NewArchitecture().IsValid() == false {
		t.Fatal("NewArchitecture() should be any and valid")
	}

	if NewASN().Equal(ASN(0)) == false {
		t.Fatal("NewASN() should be zero")
	}

	if NewDomain().Equal(Domain("")) == false {
		t.Fatal("NewDomain() should be empty")
	}

	if NewIPv4().Equal(IPv4{}) == false {
		t.Fatal("NewIPv4() should be unspecified")
	}

	if NewIPv6().Equal(IPv6{}) == false {
		t.Fatal("NewIPv6() should be unspecified")
	}

	manager := NewManager()

	if manager.Equal(Manager(Any)) == false || manager.IsValid() == false {
		t.Fatal("NewManager() should be any and valid")
	}

	if NewProtocol().Equal(Protocol(Any)) == false || NewProtocol().IsAny() == false {
		t.Fatal("NewProtocol() should be any")
	}

	version := NewVersion()

	if version.Equal(Version{}) == false {
		t.Fatal("NewVersion() should be zero")
	}

}

func TestCommonValidators(t *testing.T) {

	if IsDatetime("2026-10-19T11:59:59Z") == false || IsDatetime("not a datetime") == true {
		t.Fatal("unexpected datetime validation result")
	}

	if IsSocket("1.3.3.7:1338") == false || IsSocket("not a socket") == true {
		t.Fatal("unexpected socket validation result")
	}

	if IsTime("11:59:59") == false || IsTime("not a time") == true {
		t.Fatal("unexpected time validation result")
	}

	if IsVersion("1.23.45") == false || IsVersion("not a version") == true {
		t.Fatal("unexpected version validation result")
	}

	if IsManager(Any) == false || IsManager("unknown") == true {
		t.Fatal("unexpected manager validation result")
	}

	ipv6 := NewIPv6()

	if ipv6.IsValid() == false {
		t.Fatal("the unspecified IPv6 address should be valid")
	}

}

func TestCommonParsers(t *testing.T) {

	var architecture Architecture

	if architecture.Parse("x86_64") == false || architecture.Equal(ArchitectureX86_64) == false {
		t.Fatal("Architecture.Parse() failed")
	}

	var asn ASN

	if asn.Parse("AS1337") == false || asn.Equal(ASN(1337)) == false {
		t.Fatal("ASN.Parse() failed")
	}

	var connection Connection

	if connection.Parse("[tcp] 1.3.3.7:1338 -> 1.3.3.7:443") == false {
		t.Fatal("Connection.Parse() failed")
	}

	var domain Domain

	if domain.Parse("evil.example.") == false || domain.Equal(Domain("evil.example")) == false {
		t.Fatal("Domain.Parse() failed")
	}

	var ipv4 IPv4

	if ipv4.Parse("1.3.3.7") == false || ipv4.Equal(ToIPv4("1.3.3.7")) == false {
		t.Fatal("IPv4.Parse() failed")
	}

	var ipv6 IPv6

	if ipv6.Parse("[2001:db8::1]") == false || ipv6.Equal(ToIPv6("2001:db8::1")) == false {
		t.Fatal("IPv6.Parse() failed")
	}

	var manager Manager

	if manager.Parse(Any) == false || manager.Equal(Manager(Any)) == false {
		t.Fatal("Manager.Parse() failed")
	}

	var protocol Protocol

	if protocol.Parse("tcp") == false || protocol.Equal(ProtocolTCP) == false {
		t.Fatal("Protocol.Parse() failed")
	}

	var socket Socket

	if socket.Parse("1.3.3.7:1338") == false || socket.Equal(NewSocket("1.3.3.7", 1338)) == false {
		t.Fatal("Socket.Parse() failed")
	}

}

func TestCommonFactories(t *testing.T) {

	if ToIPv4("1.3.3.7").String() != "1.3.3.7" {
		t.Fatal("ToIPv4() failed")
	}

	if ToIPv6("2001:db8::1").String() != "2001:0db8:0000:0000:0000:0000:0000:0001" {
		t.Fatal("ToIPv6() failed")
	}

	if ToProtocol("tcp").Equal(ProtocolTCP) == false {
		t.Fatal("ToProtocol() failed")
	}

}

func TestEqual(t *testing.T) {

	if ArchitectureX86.Equal(ArchitectureX86) == false || ArchitectureX86.Equal(ArchitectureX86_64) == true {
		t.Fatal("Architecture.Equal() failed")
	}

	if ASN(1337).Equal(ASN(1337)) == false || ASN(1337).Equal(ASN(1338)) == true {
		t.Fatal("ASN.Equal() failed")
	}

	connection := ToConnection("[tcp] 1.3.3.7:1338 -> 1.3.3.7:443")

	if connection.Equal(ToConnection("[tcp] 1.3.3.7:1338 -> 1.3.3.7:443")) == false {
		t.Fatal("Connection.Equal() failed")
	}

	datetime := ToDatetime("2026-10-19T11:59:59Z")

	if datetime.Equal(ToDatetime("2026-10-19T11:59:59Z")) == false {
		t.Fatal("Datetime.Equal() failed")
	}

	if ToDomain("example.com").Equal(ToDomain("example.com")) == false {
		t.Fatal("Domain.Equal() failed")
	}

	if ToIPv4("1.3.3.7").Equal(ToIPv4("1.3.3.7")) == false || ToIPv4("1.3.3.7").Equal(ToIPv4("1.3.3.8")) == true {
		t.Fatal("IPv4.Equal() failed")
	}

	if ToIPv6("2001:db8::1").Equal(ToIPv6("2001:db8::1")) == false {
		t.Fatal("IPv6.Equal() failed")
	}

	maintainer1 := Maintainer{Name: "Harold Fynch", Email: "harold@fynch.io"}
	maintainer2 := Maintainer{Name: "Harold Fynch", Email: "harold@fynch.dev"}

	if maintainer1.Equal(maintainer2) == false {
		t.Fatal("Maintainer.Equal() should compare names")
	}

	if Manager(Any).Equal(Manager(Any)) == false || Manager(Any).Equal(ManagerAPK) == true {
		t.Fatal("Manager.Equal() failed")
	}

	if ProtocolTCP.Equal(ProtocolTCP) == false || ProtocolTCP.Equal(ProtocolUDP) == true {
		t.Fatal("Protocol.Equal() failed")
	}

	socket := NewSocket("1.3.3.7", 1338)

	if socket.Equal(NewSocket("1.3.3.7", 1338)) == false || socket.Equal(NewSocket("1.3.3.7", 443)) == true {
		t.Fatal("Socket.Equal() failed")
	}

	time := ToTime("11:59:59")

	if time.Equal(ToTime("11:59:59")) == false || time.Equal(ToTime("12:00:00")) == true {
		t.Fatal("Time.Equal() failed")
	}

	version1 := ToVersion("1.23.45")
	version2 := ToVersion("1.23.46")

	if version1.Equal(ToVersion("1.23.45")) == false || version1.Equal(version2) == true {
		t.Fatal("Version.Equal() failed")
	}

}

func TestManagerAny(t *testing.T) {

	manager := Manager(Any)

	if manager.IsValid() == false {
		t.Fatal("Manager(Any) should be valid")
	}

	parsed := ParseManager(Any)

	if parsed == nil || parsed.Equal(Manager(Any)) == false {
		t.Fatal("ParseManager(Any) should return Manager(Any)")
	}

	if ToManager(Any).Equal(Manager(Any)) == false {
		t.Fatal("ToManager(Any) should return Manager(Any)")
	}

}
