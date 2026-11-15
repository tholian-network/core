package tholian

import "testing"
import "tholian-core/protocols/tholian/methods"
import "tholian-core/protocols/tholian/parameters"

func TestEncodeDefaultsVersion(t *testing.T) {

	packet := Packet{
		Type:   TypeRequest,
		Method: methods.Ping,
	}

	blob, err := Encode(packet)

	if err != nil {
		t.Fatalf("Encode failed: %s", err)
	}

	decoded, err1 := Decode(blob, 0)

	if err1 != nil {
		t.Fatalf("Decode failed: %s", err1)
	}

	if decoded.Version != Version {
		t.Fatalf("unexpected version: %q", decoded.Version)
	}

}

func TestEncodeDecodeRequestRoundTrip(t *testing.T) {

	packet, err0 := NewRequest("id-1", methods.ScoutSearchPackages, PackageSearchFixture())

	if err0 != nil {
		t.Fatalf("NewRequest failed: %s", err0)
	}

	blob, err1 := Encode(packet)

	if err1 != nil {
		t.Fatalf("Encode failed: %s", err1)
	}

	decoded, err2 := Decode(blob, 0)

	if err2 != nil {
		t.Fatalf("Decode failed: %s", err2)
	}

	if decoded.Version != Version {
		t.Fatalf("unexpected version: %q", decoded.Version)
	}

	if decoded.Type != TypeRequest {
		t.Fatalf("unexpected type: %q", decoded.Type)
	}

	if decoded.ID != "id-1" || decoded.Method != methods.ScoutSearchPackages {
		t.Fatal("request correlation fields were not preserved")
	}

	params, err3 := DecodeParameters[parameters.ScoutSearchPackages](decoded)

	if err3 != nil {
		t.Fatalf("DecodeParameters failed: %s", err3)
	}

	if params.System != "system-1" || params.Search.Name != "log4j" {
		t.Fatal("request parameters were not preserved")
	}

}

func TestEncodeDecodeResponseRoundTrip(t *testing.T) {

	result := map[string]string{"status": "ok"}
	packet, err0 := NewResponse("id-1", methods.ScoutSearchPackages, result)

	if err0 != nil {
		t.Fatalf("NewResponse failed: %s", err0)
	}

	blob, err1 := Encode(packet)

	if err1 != nil {
		t.Fatalf("Encode failed: %s", err1)
	}

	decoded, err2 := Decode(blob, 0)

	if err2 != nil {
		t.Fatalf("Decode failed: %s", err2)
	}

	value, err3 := DecodeResult[map[string]string](decoded)

	if err3 != nil {
		t.Fatalf("DecodeResult failed: %s", err3)
	}

	if value["status"] != "ok" {
		t.Fatal("response result was not preserved")
	}

}
