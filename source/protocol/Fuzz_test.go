package protocol

import "testing"

func FuzzDecode(f *testing.F) {

	f.Add([]byte(`{"version":"1","kind":"request","method":"rpc.ping"}`))
	f.Add([]byte(`{"version":"1","kind":"event","method":"incident.new","params":{}}`))
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
