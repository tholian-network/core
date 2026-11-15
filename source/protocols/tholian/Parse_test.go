package tholian

import "encoding/json"
import "testing"
import "tholian-core/protocols/tholian/methods"
import "tholian-core/protocols/tholian/parameters"

func TestParseRequestRoundTrip(t *testing.T) {

	request, err0 := NewRequest("id-1", methods.ScoutSearchPackages, PackageSearchFixture())

	if err0 != nil {
		t.Fatalf("NewRequest failed: %s", err0)
	}

	parsed := Parse(request.Bytes())

	if parsed.Version != Version {
		t.Fatalf("unexpected version: %q", parsed.Version)
	}

	if parsed.Type != TypeRequest {
		t.Fatalf("unexpected type: %q", parsed.Type)
	}

	if parsed.ID != "id-1" || parsed.Method != methods.ScoutSearchPackages {
		t.Fatal("request correlation fields were not preserved")
	}

	params, err1 := DecodeParameters[parameters.ScoutSearchPackages](parsed)

	if err1 != nil {
		t.Fatalf("DecodeParameters failed: %s", err1)
	}

	if params.System != "system-1" || params.Search.Name != "log4j" {
		t.Fatal("request parameters were not preserved")
	}

}

func TestParseResponseRoundTrip(t *testing.T) {

	response, err0 := NewResponse("id-1", methods.ScoutSearchPackages, map[string]string{"status": "ok"})

	if err0 != nil {
		t.Fatalf("NewResponse failed: %s", err0)
	}

	parsed := Parse(response.Bytes())

	if parsed.Type != TypeResponse {
		t.Fatalf("unexpected type: %q", parsed.Type)
	}

	result, err1 := DecodeResult[map[string]string](parsed)

	if err1 != nil {
		t.Fatalf("DecodeResult failed: %s", err1)
	}

	if result["status"] != "ok" {
		t.Fatal("response result was not preserved")
	}

}

func TestParseEventRoundTrip(t *testing.T) {

	event, err0 := NewEvent(methods.GuardOnline, map[string]string{"system": "system-1"})

	if err0 != nil {
		t.Fatalf("NewEvent failed: %s", err0)
	}

	parsed := Parse(event.Bytes())

	if parsed.Type != TypeEvent || parsed.Method != methods.GuardOnline {
		t.Fatal("event fields were not preserved")
	}

	params, err1 := DecodeParameters[map[string]string](parsed)

	if err1 != nil {
		t.Fatalf("DecodeParameters failed: %s", err1)
	}

	if params["system"] != "system-1" {
		t.Fatal("event parameters were not preserved")
	}

}

func TestParseErrorRoundTrip(t *testing.T) {

	failure := NewError(ErrorForbidden, "not allowed")
	packet := NewErrorResponse("id-1", methods.WatchtowerListAudits, failure)

	parsed := Parse(packet.Bytes())

	if parsed.Type != TypeError {
		t.Fatalf("unexpected type: %q", parsed.Type)
	}

	if parsed.Error == nil || parsed.Error.Code != ErrorForbidden || parsed.Error.Message != "not allowed" {
		t.Fatal("error details were not preserved")
	}

}

func TestParseInvalidBuffer(t *testing.T) {

	parsed := Parse([]byte("not-json"))

	if parsed.Version != Version || parsed.Type != TypeRequest {
		t.Fatal("invalid buffer did not return packet defaults")
	}

}

func TestParseWireFormat(t *testing.T) {

	packet := NewPacket()
	packet.SetID("id-1")

	var object map[string]any
	err := json.Unmarshal(packet.Bytes(), &object)

	if err != nil {
		t.Fatalf("Bytes returned invalid JSON: %s", err)
	}

	if object["type"] != TypeRequest {
		t.Fatal("packet did not use the type field")
	}

	if _, ok := object["kind"]; ok == true {
		t.Fatal("packet still used the kind field")
	}

	packet.SetParameters(map[string]string{"value": "test"})

	err1 := json.Unmarshal(packet.Bytes(), &object)

	if err1 != nil {
		t.Fatalf("Bytes returned invalid JSON: %s", err1)
	}

	if _, ok := object["parameters"]; ok == false {
		t.Fatal("packet did not use the parameters field")
	}

	if _, ok := object["params"]; ok == true {
		t.Fatal("packet still used the params field")
	}

}

func TestBytesDefaultsVersion(t *testing.T) {

	packet := Packet{
		Type:   TypeRequest,
		Method: methods.Ping,
	}

	parsed := Parse(packet.Bytes())

	if parsed.Version != Version {
		t.Fatalf("unexpected version: %q", parsed.Version)
	}

}
