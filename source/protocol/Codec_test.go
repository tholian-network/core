package protocol

import "testing"
import "tholian-core/protocol/methods"
import "tholian-core/protocol/parameters"

func TestCodecRequestRoundTrip(t *testing.T) {

	envelope, err0 := NewRequest("id-1", methods.ScoutSearchPackages, PackageSearchFixture())

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

	if decoded.ID != "id-1" || decoded.Method != methods.ScoutSearchPackages {
		t.Fatal("request correlation fields were not preserved")
	}

	params, err3 := DecodeParams[parameters.ScoutSearchPackages](decoded)

	if err3 != nil {
		t.Fatalf("DecodeParams failed: %s", err3)
	}

	if params.System != "system-1" || params.Search.Name != "log4j" {
		t.Fatal("request params were not preserved")
	}

}

func TestCodecResponseRoundTrip(t *testing.T) {

	result := map[string]string{"status": "ok"}
	envelope, err0 := NewResponse("id-1", methods.ScoutSearchPackages, result)

	if err0 != nil {
		t.Fatalf("NewResponse failed: %s", err0)
	}

	blob, err1 := Encode(envelope)

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

func TestEncodeDefaultsVersion(t *testing.T) {

	envelope := Envelope{
		Kind:   KindRequest,
		Method: methods.RPCPing,
	}

	blob, err := Encode(envelope)

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

func TestDecodeRejectsUnsupportedVersion(t *testing.T) {

	_, err := Decode([]byte(`{"version":"999","kind":"request","method":"tholian-rpc.Ping"}`), 0)

	if err == nil {
		t.Fatal("unsupported version was accepted")
	}

}

func TestDecodeRejectsInvalidKind(t *testing.T) {

	_, err := Decode([]byte(`{"version":"2","kind":"nonsense","method":"tholian-rpc.Ping"}`), 0)

	if err == nil {
		t.Fatal("invalid kind was accepted")
	}

}

func TestDecodeRejectsOversizedMessage(t *testing.T) {

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

func TestDecodeParamsRejectsMissingOrInvalidData(t *testing.T) {

	_, err0 := DecodeParams[parameters.ScoutSearchPackages](Envelope{})

	if err0 == nil {
		t.Fatal("missing params were accepted")
	}

	envelope := Envelope{Parameters: []byte(`not-json`)}
	_, err1 := DecodeParams[parameters.ScoutSearchPackages](envelope)

	if err1 == nil {
		t.Fatal("invalid params were accepted")
	}

}

func TestDecodeResultRejectsMissingOrInvalidData(t *testing.T) {

	_, err0 := DecodeResult[map[string]string](Envelope{})

	if err0 == nil {
		t.Fatal("missing result was accepted")
	}

	envelope := Envelope{Result: []byte(`not-json`)}
	_, err1 := DecodeResult[map[string]string](envelope)

	if err1 == nil {
		t.Fatal("invalid result was accepted")
	}

}

func FuzzDecode(f *testing.F) {

	f.Add([]byte(`{"version":"1","kind":"request","method":"tholian-rpc.Ping"}`))
	f.Add([]byte(`{"version":"1","kind":"event","method":"tholian-guard.Online","params":{}}`))
	f.Add([]byte(`not-json`))

	f.Fuzz(func(t *testing.T, data []byte) {

		envelope, err := Decode(data, 4096)

		if err == nil {

			if envelope.Version != Version {
				t.Fatalf("decoded envelope with invalid version")
			}

			switch envelope.Kind {
			case KindRequest, KindResponse, KindEvent, KindError:
			default:
				t.Fatalf("decoded envelope with invalid kind")
			}

		}

	})

}
