package matchers

import "testing"
import "tholian-core/types"

func TestEvidenceNew(t *testing.T) {

	evidence := NewEvidence()

	if evidence.Type != Any || evidence.Datetime != Any {
		t.Fatalf("unexpected defaults: %#v", evidence)
	}

	if evidence.IsValid() == true {
		t.Fatal("new evidence matcher should be invalid")
	}

}

func TestEvidenceToEvidence(t *testing.T) {

	evidence := ToEvidence("application/vnd.tcpdump.pcap")

	if evidence.Type != "application/vnd.tcpdump.pcap" {
		t.Fatalf("unexpected type: %q", evidence.Type)
	}

	if evidence.Datetime != Any {
		t.Fatalf("unexpected datetime: %q", evidence.Datetime)
	}

	if evidence.IsValid() == false {
		t.Fatal("evidence matcher with a type should be valid")
	}

}

func TestEvidenceSetTypeAndDatetime(t *testing.T) {

	evidence := NewEvidence()

	evidence.SetType("  application/vnd.tcpdump.pcap+gzip  ")

	if evidence.Type != "application/vnd.tcpdump.pcap+gzip" {
		t.Fatalf("unexpected type: %q", evidence.Type)
	}

	evidence.SetType("*")

	if evidence.Type != Any {
		t.Fatalf("wildcard did not reset the type: %q", evidence.Type)
	}

	evidence.SetDatetime("2026-10-29T12:00:00Z")

	if evidence.Datetime != "2026-10-29T12:00:00Z" {
		t.Fatalf("unexpected datetime: %q", evidence.Datetime)
	}

	evidence.SetDatetime("not-a-datetime")

	if evidence.Datetime != "2026-10-29T12:00:00Z" {
		t.Fatal("invalid datetime overwrote the value")
	}

}

func TestEvidenceEqual(t *testing.T) {

	left := Evidence{Type: "application/vnd.tcpdump.pcap", Datetime: "2026-10-29T12:00:00Z"}
	right := Evidence{Type: "application/vnd.tcpdump.pcap", Datetime: "2026-10-29T12:00:00Z"}

	if left.Equal(right) == false {
		t.Fatal("identical evidence matchers should be equal")
	}

	right.Type = "application/json"

	if left.Equal(right) == true {
		t.Fatal("evidence matchers with different types should differ")
	}

}

func TestEvidenceMatchesType(t *testing.T) {

	evidence := Evidence{Type: "application/vnd.tcpdump.pcap", Datetime: Any}

	if evidence.MatchesType("application/vnd.tcpdump.pcap") == false {
		t.Fatal("MatchesType should match the exact type")
	}

	if evidence.MatchesType("application/json") == true {
		t.Fatal("MatchesType should not match a different type")
	}

	any := NewEvidence()

	if any.MatchesType("anything") == false {
		t.Fatal("any type matcher should match everything")
	}

}

func TestEvidenceMatchesDatetime(t *testing.T) {

	exact := Evidence{Type: Any, Datetime: "2026-10-29T12:00:00Z"}

	if exact.MatchesDatetime("2026-10-29T12:00:00Z") == false {
		t.Fatal("exact datetime should match")
	}

	if exact.MatchesDatetime("2026-10-30T12:00:00Z") == true {
		t.Fatal("exact datetime should not match a different value")
	}

	ranged := Evidence{Type: Any, Datetime: "2026-01-01T00:00:00Z - 2026-12-31T23:59:59Z"}

	if ranged.MatchesDatetime("2026-10-29T12:00:00Z") == false {
		t.Fatal("datetime range should match")
	}

	if ranged.MatchesDatetime("2027-01-01T00:00:00Z") == true {
		t.Fatal("datetime range should not match outside the range")
	}

	any := NewEvidence()

	if any.MatchesDatetime("anything") == false {
		t.Fatal("any datetime matcher should match everything")
	}

	current := types.ToDatetime("2026-10-29T12:00:00Z")

	if any.MatchesDatetime(current.String()) == false {
		t.Fatal("any datetime matcher should match a valid datetime")
	}

}

func TestEvidenceMatchesDatetimeOperators(t *testing.T) {

	value := "2026-10-29T12:00:00Z"

	cases := []struct {
		datetime string
		matches  bool
	}{
		{"< 2026-10-30T00:00:00Z", true},
		{"< 2026-10-29T12:00:00Z", false},
		{"<= 2026-10-29T12:00:00Z", true},
		{"> 2026-10-28T00:00:00Z", true},
		{"> 2026-10-29T12:00:00Z", false},
		{">= 2026-10-29T12:00:00Z", true},
		{"= 2026-10-29T12:00:00Z", true},
		{"= 2026-10-30T00:00:00Z", false},
	}

	for _, c := range cases {

		evidence := Evidence{Type: Any, Datetime: c.datetime}

		if evidence.MatchesDatetime(value) != c.matches {
			t.Fatalf("MatchesDatetime(%q) = %v, expected %v", c.datetime, evidence.MatchesDatetime(value), c.matches)
		}

	}

}

func TestEvidenceMatches(t *testing.T) {

	evidence := Evidence{Type: "application/vnd.tcpdump.pcap", Datetime: "2026-10-29T12:00:00Z"}

	if evidence.Matches("application/vnd.tcpdump.pcap", "2026-10-29T12:00:00Z") == false {
		t.Fatal("central Matches should match")
	}

	if evidence.Matches("application/json", "2026-10-29T12:00:00Z") == true {
		t.Fatal("central Matches should not match a different type")
	}

}

func TestEvidenceHash(t *testing.T) {

	evidence := Evidence{Type: "application/vnd.tcpdump.pcap", Datetime: "2026-10-29T12:00:00Z"}
	hash := evidence.Hash()

	if hash == "" {
		t.Fatal("expected a non-empty hash")
	}

	if evidence.Hash() != hash {
		t.Fatal("expected a deterministic hash")
	}

	other := evidence
	other.Type = "application/json"

	if other.Hash() == hash {
		t.Fatal("expected the type to change the hash")
	}

	empty := Evidence{}

	if empty.Hash() != "" {
		t.Fatalf("expected an empty hash, got %q", empty.Hash())
	}

}
