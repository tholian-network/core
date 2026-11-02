package protocol

import "testing"

func FuzzDecode(f *testing.F) {

	f.Add([]byte(`{"version":"2","kind":"request","method":"tholian-rpc.Ping"}`))
	f.Add([]byte(`{"version":"2","kind":"event","method":"tholian-guard.IncidentNew","params":{}}`))
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
