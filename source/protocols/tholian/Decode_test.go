package tholian

import "testing"
import "tholian-core/protocols/tholian/parameters"

func TestDecodeRejectsUnsupportedVersion(t *testing.T) {

	_, err := Decode([]byte(`{"version":"999","type":"request","method":"tholian-rpc.Ping"}`), 0)

	if err == nil {
		t.Fatal("unsupported version was accepted")
	}

}

func TestDecodeRejectsInvalidType(t *testing.T) {

	_, err := Decode([]byte(`{"version":"1","type":"nonsense","method":"tholian-rpc.Ping"}`), 0)

	if err == nil {
		t.Fatal("invalid type was accepted")
	}

}

func TestDecodeRejectsMissingType(t *testing.T) {

	_, err := Decode([]byte(`{"version":"1","method":"tholian-rpc.Ping"}`), 0)

	if err == nil {
		t.Fatal("missing type was accepted")
	}

}

func TestDecodeRejectsLegacyKind(t *testing.T) {

	_, err := Decode([]byte(`{"version":"1","kind":"request","method":"tholian-rpc.Ping"}`), 0)

	if err == nil {
		t.Fatal("legacy kind field was accepted")
	}

}

func TestDecodeRejectsOversizedMessage(t *testing.T) {

	blob := []byte(`{"version":"1","type":"request","method":"tholian-rpc.Ping","parameters":{"value":"`)

	for len(blob) < 64 {
		blob = append(blob, 'a')
	}

	blob = append(blob, []byte(`"}}`)...)

	_, err := Decode(blob, 16)

	if err == nil {
		t.Fatal("oversized message was accepted")
	}

}

func TestDecodeParametersRejectsMissingOrInvalidData(t *testing.T) {

	_, err0 := DecodeParameters[parameters.ScoutSearchPackages](Packet{})

	if err0 == nil {
		t.Fatal("missing parameters were accepted")
	}

	packet := Packet{Parameters: []byte(`not-json`)}
	_, err1 := DecodeParameters[parameters.ScoutSearchPackages](packet)

	if err1 == nil {
		t.Fatal("invalid parameters were accepted")
	}

}

func TestDecodeResultRejectsMissingOrInvalidData(t *testing.T) {

	_, err0 := DecodeResult[map[string]string](Packet{})

	if err0 == nil {
		t.Fatal("missing result was accepted")
	}

	packet := Packet{Result: []byte(`not-json`)}
	_, err1 := DecodeResult[map[string]string](packet)

	if err1 == nil {
		t.Fatal("invalid result was accepted")
	}

}

func FuzzDecode(f *testing.F) {

	f.Add([]byte(`{"version":"1","type":"request","method":"tholian-rpc.Ping"}`))
	f.Add([]byte(`{"version":"1","type":"event","method":"tholian-guard.Online","parameters":{}}`))
	f.Add([]byte(`not-json`))

	f.Fuzz(func(t *testing.T, data []byte) {

		packet, err := Decode(data, 4096)

		if err == nil {

			if packet.Version != Version {
				t.Fatalf("decoded packet with invalid version")
			}

			switch packet.Type {
			case TypeRequest, TypeResponse, TypeEvent, TypeError:
			default:
				t.Fatalf("decoded packet with invalid type")
			}

		}

	})

}
