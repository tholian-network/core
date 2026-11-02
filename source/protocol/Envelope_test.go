package protocol

import "tholian-core/protocol/permissions"
import "testing"

func TestEnvelopeRequestRoundTrip(t *testing.T) {

	envelope, err0 := NewRequest("id-1", permissions.MethodScoutSearchPackages, PackageSearchFixture())

	if err0 != nil {
		t.Fatalf("NewRequest failed: %s", err0)
	}

	blob, err1 := Encode(envelope)

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

	if decoded.Kind != KindRequest {
		t.Fatalf("unexpected kind: %q", decoded.Kind)
	}

	if decoded.ID != "id-1" || decoded.Method != permissions.MethodScoutSearchPackages {
		t.Fatal("request correlation fields were not preserved")
	}

	params, err3 := DecodeParams[PackagesSearchParams](decoded)

	if err3 != nil {
		t.Fatalf("DecodeParams failed: %s", err3)
	}

	if params.SystemID != "system-1" || params.Search.Name != "log4j" {
		t.Fatal("request params were not preserved")
	}

}

func TestEnvelopeRejectsUnsupportedVersion(t *testing.T) {

	envelope := Envelope{
		Version: "999",
		Kind:    KindRequest,
		Method:  permissions.MethodRPCPing,
	}

	_, err := Decode([]byte(`{"version":"999","kind":"request","method":"tholian-rpc.Ping"}`), 0)

	if err == nil {
		t.Fatal("unsupported version was accepted")
	}

	_ = envelope

}

func TestEnvelopeRejectsInvalidKind(t *testing.T) {

	_, err := Decode([]byte(`{"version":"2","kind":"nonsense","method":"tholian-rpc.Ping"}`), 0)

	if err == nil {
		t.Fatal("invalid kind was accepted")
	}

}

func TestEnvelopeRejectsOversizedMessage(t *testing.T) {

	blob := []byte(`{"version":"2","kind":"request","method":"tholian-rpc.Ping","params":{"value":"`)

	for len(blob) < 64 {
		blob = append(blob, 'a')
	}

	blob = append(blob, []byte(`"}}`)...)

	_, err := Decode(blob, 16)

	if err == nil {
		t.Fatal("oversized message was accepted")
	}

}

func TestEnvelopeErrorResponse(t *testing.T) {

	failure := NewError(ErrorForbidden, "not allowed")
	envelope := NewErrorResponse("id-1", permissions.MethodWatchtowerCall, failure)

	if envelope.IsError() == false {
		t.Fatal("error response was not marked as error")
	}

	if envelope.Error == nil || envelope.Error.Code != ErrorForbidden {
		t.Fatal("error details were not preserved")
	}

	blob, err0 := Encode(envelope)

	if err0 != nil {
		t.Fatalf("Encode failed: %s", err0)
	}

	decoded, err1 := Decode(blob, 0)

	if err1 != nil {
		t.Fatalf("Decode failed: %s", err1)
	}

	if decoded.Error == nil || decoded.Error.Error() != "forbidden: not allowed" {
		t.Fatalf("unexpected error: %v", decoded.Error)
	}

}

func PackageSearchFixture() PackagesSearchParams {

	var params PackagesSearchParams

	params.SystemID = "system-1"
	params.Search.Name = "log4j"

	return params

}
