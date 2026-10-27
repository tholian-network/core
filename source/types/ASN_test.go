package types

import (
	"encoding/json"
	"testing"
)

func TestIsASN(t *testing.T) {

	if IsASN("AS1337") == false || IsASN("as1337") == true || IsASN("AS") == true || IsASN("AS12X") == true {
		t.Fatal("unexpected ASN validation result")
	}

}

func TestIsASNTable(t *testing.T) {

	tests := []struct {
		input    string
		expected bool
	}{
		{"AS1337", true},
		{"AS0", true},
		{"AS4294967295", true},
		{"as1337", false},
		{"AS", false},
		{"AS12X", false},
		{"AS-1", false},
		{"1337", false},
		{"", false},
	}

	for _, test := range tests {

		if result := IsASN(test.input); result != test.expected {
			t.Errorf("IsASN(%q) = %t, expected %t", test.input, result, test.expected)
		}

	}

}

func TestNewASN(t *testing.T) {

	asn := NewASN()

	if asn != ASN(0) {
		t.Fatalf("NewASN() = %d, expected 0", asn)
	}

	if asn.String() != "AS0" {
		t.Fatalf("NewASN().String() = %q, expected %q", asn.String(), "AS0")
	}

}

func TestParseASN(t *testing.T) {

	asn := ParseASN("AS1337")

	if asn == nil || asn.String() != "AS1337" {
		t.Fatalf("unexpected ASN parse result: %v", asn)
	}

	asn = ParseASN("1337")

	if asn == nil || asn.String() != "AS1337" {
		t.Fatalf("unexpected numeric ASN parse result: %v", asn)
	}

	if ParseASN("invalid") != nil {
		t.Fatal("ParseASN() returned a value for an invalid ASN")
	}

	if ParseASN("") != nil {
		t.Fatal("ParseASN() returned a value for an empty ASN")
	}

	if ParseASN("as1337") != nil {
		t.Fatal("ParseASN() accepted a lowercase prefix")
	}

}

func TestASNParse(t *testing.T) {

	tests := []struct {
		input    string
		expected ASN
	}{
		{"AS1337", ASN(1337)},
		{"1337", ASN(1337)},
		{"AS0", ASN(0)},
		{"0", ASN(0)},
		{"AS-1", ASN(-1)},
		{"-1", ASN(-1)},
	}

	for _, test := range tests {

		var asn ASN

		if asn.Parse(test.input) == false {
			t.Errorf("Parse(%q) = false, expected true", test.input)
		} else if asn != test.expected {
			t.Errorf("Parse(%q) = %d, expected %d", test.input, asn, test.expected)
		}

	}

	invalid := []string{"", "AS", "AS12X", "as1337", "invalid"}

	for _, value := range invalid {

		var asn ASN

		if asn.Parse(value) == true {
			t.Errorf("Parse(%q) = true, expected false", value)
		}

	}

}

func TestToASN(t *testing.T) {

	if ToASN("AS1337") != ASN(1337) {
		t.Fatalf("ToASN(%q) = %d, expected 1337", "AS1337", ToASN("AS1337"))
	}

	if ToASN("1337") != ASN(1337) {
		t.Fatalf("ToASN(%q) = %d, expected 1337", "1337", ToASN("1337"))
	}

	if ToASN("invalid") != ASN(0) {
		t.Fatalf("ToASN(%q) = %d, expected zero value", "invalid", ToASN("invalid"))
	}

	if ToASN("AS-1") != ASN(-1) {
		t.Fatalf("ToASN(%q) = %d, expected -1", "AS-1", ToASN("AS-1"))
	}

}

func TestASNEqual(t *testing.T) {

	if ASN(1337).Equal(ASN(1337)) == false {
		t.Fatal("Equal() returned false for identical ASNs")
	}

	if ASN(1337).Equal(ASN(1338)) == true {
		t.Fatal("Equal() returned true for different ASNs")
	}

	if NewASN().Equal(ASN(0)) == false {
		t.Fatal("Equal() returned false for zero ASNs")
	}

}

func TestASNHash(t *testing.T) {

	first := ASN(1337).Hash()
	second := ASN(1337).Hash()

	if first == "" {
		t.Fatal("Hash() returned an empty string")
	}

	if first != second {
		t.Fatalf("Hash() is not deterministic: %q != %q", first, second)
	}

	if first == ASN(1338).Hash() {
		t.Fatal("Hash() returned the same value for different ASNs")
	}

}

func TestASNIsValid(t *testing.T) {

	if ASN(1337).IsValid() == false {
		t.Fatal("valid ASN reported invalid")
	}

	if ASN(-1).IsValid() == true {
		t.Fatal("negative ASN reported valid")
	}

	if NewASN().IsValid() == false {
		t.Fatal("zero ASN reported invalid")
	}

}

func TestASNString(t *testing.T) {

	if ASN(1337).String() != "AS1337" {
		t.Fatalf("String() = %q, expected %q", ASN(1337).String(), "AS1337")
	}

	if NewASN().String() != "AS0" {
		t.Fatalf("String() = %q, expected %q", NewASN().String(), "AS0")
	}

	if ASN(-1).String() != "AS-1" {
		t.Fatalf("String() = %q, expected %q", ASN(-1).String(), "AS-1")
	}

}

func TestASNMarshalJSON(t *testing.T) {

	data, err := ASN(1337).MarshalJSON()

	if err != nil {
		t.Fatalf("MarshalJSON() returned an error: %v", err)
	}

	if string(data) != `"AS1337"` {
		t.Fatalf("MarshalJSON() = %s, expected %q", data, `"AS1337"`)
	}

	data, err = json.Marshal(NewASN())

	if err != nil {
		t.Fatalf("json.Marshal() returned an error: %v", err)
	}

	if string(data) != `"AS0"` {
		t.Fatalf("json.Marshal() = %s, expected %q", data, `"AS0"`)
	}

}

func TestASNUnmarshalJSON(t *testing.T) {

	var asn ASN

	if err := json.Unmarshal([]byte(`"AS42"`), &asn); err != nil {
		t.Fatalf("UnmarshalJSON() returned an error: %v", err)
	}

	if asn != ASN(42) {
		t.Fatalf("UnmarshalJSON() = %d, expected 42", asn)
	}

	var numeric ASN

	if err := json.Unmarshal([]byte(`"42"`), &numeric); err != nil {
		t.Fatalf("UnmarshalJSON() returned an error: %v", err)
	}

	if numeric != ASN(42) {
		t.Fatalf("UnmarshalJSON() = %d, expected 42", numeric)
	}

	var zero ASN

	if err := json.Unmarshal([]byte(`"invalid"`), &zero); err != nil {
		t.Fatalf("UnmarshalJSON() returned an error for an invalid value: %v", err)
	}

	if zero != ASN(0) {
		t.Fatalf("UnmarshalJSON() changed the value for an invalid ASN: %d", zero)
	}

	if err := zero.UnmarshalJSON([]byte("not-quoted")); err == nil {
		t.Fatal("UnmarshalJSON() accepted malformed JSON")
	}

}
