package structs

import "bytes"
import "encoding/json"
import "testing"
import "tholian-core/types"

func evidenceFixture() Evidence {

	evidence := NewEvidence(EvidenceTypePCAP, []byte{0x00, 0x01, 0x02, 0x03})
	evidence.SetDatetime("2026-10-29T12:00:00Z")

	return evidence

}

func TestEvidenceMIMETypes(t *testing.T) {

	cases := map[string]string{
		EvidenceTypePCAP:     "application/vnd.tcpdump.pcap",
		EvidenceTypePCAPGzip: "application/vnd.tcpdump.pcap+gzip",
		EvidenceTypePCAPNG:   "application/x-pcapng",
		EvidenceTypeJSON:     "application/json",
		EvidenceTypeText:     "text/plain",
	}

	for constant, expected := range cases {

		if constant != expected {
			t.Fatalf("evidence MIME constant = %q, expected %q", constant, expected)
		}

	}

}

func TestEvidenceNew(t *testing.T) {

	evidence := NewEvidence(EvidenceTypePCAPGzip, []byte{1, 2, 3})

	if evidence.Type != EvidenceTypePCAPGzip {
		t.Fatalf("unexpected type: %q", evidence.Type)
	}

	if evidence.Datetime.IsValid() == false {
		t.Fatal("new evidence has an invalid datetime")
	}

	if bytes.Equal(evidence.Data, []byte{1, 2, 3}) == false {
		t.Fatalf("unexpected data: %v", evidence.Data)
	}

	empty := NewEvidence("", nil)

	if empty.IsValid() == true {
		t.Fatal("empty evidence should be invalid")
	}

}

func TestEvidenceSetTypeAndDatetime(t *testing.T) {

	evidence := evidenceFixture()

	evidence.SetType("  " + EvidenceTypePCAPGzip + "  ")

	if evidence.Type != EvidenceTypePCAPGzip {
		t.Fatalf("unexpected type: %q", evidence.Type)
	}

	expected := types.ToDatetime("2024-01-02T03:04:05Z")
	evidence.SetDatetime("2024-01-02T03:04:05Z")

	if evidence.Datetime.Equal(expected) == false {
		t.Fatalf("unexpected datetime: %s", evidence.Datetime.String())
	}

	evidence.SetDatetime("invalid")

	if evidence.Datetime.Equal(expected) == false {
		t.Fatal("invalid datetime overwrote value")
	}

}

func TestEvidenceSetData(t *testing.T) {

	evidence := evidenceFixture()
	source := []byte{4, 5, 6}

	evidence.SetData(source)

	if bytes.Equal(evidence.Data, source) == false {
		t.Fatalf("unexpected data: %v", evidence.Data)
	}

	source[0] = 9

	if evidence.Data[0] != 4 {
		t.Fatal("SetData did not copy the source slice")
	}

	evidence.SetData(nil)

	if len(evidence.Data) != 0 {
		t.Fatal("SetData(nil) did not clear the data")
	}

}

func TestEvidenceEqual(t *testing.T) {

	left := evidenceFixture()
	right := evidenceFixture()

	if left.Equal(right) == false {
		t.Fatal("identical evidence should be equal")
	}

	if left.Equal(left) == false {
		t.Fatal("evidence should equal itself")
	}

	otherType := evidenceFixture()
	otherType.SetType(EvidenceTypeJSON)

	if left.Equal(otherType) == true {
		t.Fatal("evidence with a different type should differ")
	}

	otherDatetime := evidenceFixture()
	otherDatetime.SetDatetime("2020-01-01T00:00:00Z")

	if left.Equal(otherDatetime) == true {
		t.Fatal("evidence with a different datetime should differ")
	}

	otherData := evidenceFixture()
	otherData.SetData([]byte{0xff})

	if left.Equal(otherData) == true {
		t.Fatal("evidence with different data should differ")
	}

}

func TestEvidenceHash(t *testing.T) {

	evidence := evidenceFixture()
	hash := evidence.Hash()

	if hash == "" {
		t.Fatal("expected a non-empty hash")
	}

	if evidence.Hash() != hash {
		t.Fatal("expected a deterministic hash")
	}

	other := evidenceFixture()
	other.SetData([]byte{0xff})

	if other.Hash() == hash {
		t.Fatal("expected data to change the hash")
	}

	invalid := Evidence{}

	if invalid.Hash() != "" {
		t.Fatalf("expected an empty hash, got %q", invalid.Hash())
	}

}

func TestEvidenceIsValid(t *testing.T) {

	cases := []struct {
		name     string
		evidence Evidence
		valid    bool
	}{
		{"valid", evidenceFixture(), true},
		{"missing-type", Evidence{Datetime: evidenceFixture().Datetime, Data: []byte{1}}, false},
		{"missing-datetime", Evidence{Type: EvidenceTypePCAP, Data: []byte{1}}, false},
		{"missing-data", Evidence{Type: EvidenceTypePCAP, Datetime: evidenceFixture().Datetime}, false},
	}

	for _, c := range cases {

		t.Run(c.name, func(t *testing.T) {

			if c.evidence.IsValid() != c.valid {
				t.Fatalf("IsValid() = %v, expected %v", c.evidence.IsValid(), c.valid)
			}

		})

	}

}

func TestEvidenceJSONRoundTrip(t *testing.T) {

	evidence := evidenceFixture()
	data, err := json.Marshal(evidence)

	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var decoded Evidence

	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if evidence.Equal(decoded) == false {
		t.Fatalf("round-trip mismatch: %#v != %#v", evidence, decoded)
	}

}
