package types

import (
	"encoding/json"
	"testing"
)

func TestMaintainer(t *testing.T) {

	t.Run("IsMaintainer()", func(t *testing.T) {

		result1 := IsMaintainer("Harold Fynch")
		result2 := IsMaintainer("harold@fynch.io")
		result3 := IsMaintainer("Harold Fynch <harold@fynch.io>")
		result4 := IsMaintainer("<harold@fynch.io>")
		result5 := IsMaintainer("<harold@fynch.io> Harold Fynch")

		if result1 != true {
			t.Errorf("Expected %t to be true", result1)
		}

		if result2 != true {
			t.Errorf("Expected %t to be true", result2)
		}

		if result3 != true {
			t.Errorf("Expected %t to be true", result3)
		}

		if result3 != true {
			t.Errorf("Expected %t to be true", result4)
		}

		if result5 != false {
			t.Errorf("Expected %t to be false", result5)
		}

	})

	t.Run("ParseMaintainer()", func(t *testing.T) {

		maintainer1 := ParseMaintainer("Harold Fynch")
		maintainer2 := ParseMaintainer("harold@fynch.io")
		maintainer3 := ParseMaintainer("Harold Fynch <harold@fynch.io>")
		maintainer4 := ParseMaintainer("<harold@fynch.io>")
		maintainer5 := ParseMaintainer("<harold@fynch.io> Harold Fynch")

		if maintainer1 == nil {
			t.Errorf("Expected %s to be valid", "Harold Fynch")
		} else if maintainer1.String() != "Harold Fynch" {
			t.Errorf("Expected %s to be %s", maintainer1.String(), "Harold Fynch")
		}

		if maintainer2 == nil {
			t.Errorf("Expected %s to be valid", "harold@fynch.io")
		} else if maintainer2.String() != "harold@fynch.io" {
			t.Errorf("Expected %s to be %s", maintainer2.String(), "harold@fynch.io")
		}

		if maintainer3 == nil {
			t.Errorf("Expected %s to be valid", "Harold Fynch <harold@fynch.io>")
		} else if maintainer3.String() != "Harold Fynch <harold@fynch.io>" {
			t.Errorf("Expected %s to be %s", maintainer3.String(), "Harold Fynch <harold@fynch.io>")
		}

		if maintainer4 == nil {
			t.Errorf("Expected %s to be valid", "<harold@fynch.io>")
		} else if maintainer4.String() != "harold@fynch.io" {
			t.Errorf("Expected %s to be %s", maintainer4.String(), "harold@fynch.io")
		}

		if maintainer5 != nil {
			t.Errorf("Expected %s to be invalid", maintainer5.String())
		}

	})

}

func TestNewMaintainer(t *testing.T) {

	maintainer := NewMaintainer()

	if maintainer.Name != "" || maintainer.Email != "" {
		t.Fatalf("NewMaintainer() = %#v, expected empty fields", maintainer)
	}

	if maintainer.IsValid() == true {
		t.Fatal("NewMaintainer() returned a valid maintainer")
	}

}

func TestToMaintainer(t *testing.T) {

	maintainer := ToMaintainer("Harold Fynch <harold@fynch.io>")

	if maintainer.Name != "Harold Fynch" || maintainer.Email != "harold@fynch.io" {
		t.Fatalf("ToMaintainer() = %#v, expected name and email", maintainer)
	}

	email := ToMaintainer("<harold@fynch.io>")

	if email.Name != "" || email.Email != "harold@fynch.io" {
		t.Fatalf("ToMaintainer() = %#v, expected only an email", email)
	}

	name := ToMaintainer("Harold Fynch")

	if name.Name != "Harold Fynch" || name.Email != "" {
		t.Fatalf("ToMaintainer() = %#v, expected only a name", name)
	}

	empty := ToMaintainer("")

	if empty.Name != "" || empty.Email != "" {
		t.Fatalf("ToMaintainer(%q) = %#v, expected zero value", "", empty)
	}

}

func TestMaintainerParse(t *testing.T) {

	tests := []struct {
		input string
		name  string
		email string
		valid bool
	}{
		{"Harold Fynch <harold@fynch.io>", "Harold Fynch", "harold@fynch.io", true},
		{"<harold@fynch.io>", "", "harold@fynch.io", true},
		{"harold@fynch.io", "", "harold@fynch.io", true},
		{"Harold Fynch", "Harold Fynch", "", true},
		{"  Harold Fynch  ", "Harold Fynch", "", true},
		{"", "", "", false},
		{"   ", "", "", false},
		{"<>", "", "", false},
		{"<harold@fynch.io> Harold Fynch", "", "", false},
	}

	for _, test := range tests {

		var maintainer Maintainer

		if result := maintainer.Parse(test.input); result != test.valid {
			t.Errorf("Parse(%q) = %t, expected %t", test.input, result, test.valid)
			continue
		}

		if maintainer.Name != test.name || maintainer.Email != test.email {
			t.Errorf("Parse(%q) = %#v, expected name %q and email %q", test.input, maintainer, test.name, test.email)
		}

	}

}

func TestMaintainerEqual(t *testing.T) {

	first := Maintainer{Name: "Harold Fynch", Email: "harold@fynch.io"}

	if first.Equal(Maintainer{Name: "Harold Fynch", Email: "harold@fynch.io"}) == false {
		t.Fatal("Equal() returned false for identical maintainers")
	}

	if first.Equal(Maintainer{Name: "Harold Fynch", Email: "other@fynch.io"}) == false {
		t.Fatal("Equal() should only compare names")
	}

	if first.Equal(Maintainer{Name: "Other Fynch", Email: "harold@fynch.io"}) == true {
		t.Fatal("Equal() returned true for different names")
	}

}

func TestMaintainerHash(t *testing.T) {

	maintainer := Maintainer{Name: "Harold Fynch", Email: "harold@fynch.io"}

	first := maintainer.Hash()
	second := maintainer.Hash()

	if first == "" {
		t.Fatal("Hash() returned an empty string")
	}

	if first != second {
		t.Fatalf("Hash() is not deterministic: %q != %q", first, second)
	}

	if first == (Maintainer{Name: "Harold Fynch", Email: "other@fynch.io"}).Hash() {
		t.Fatal("Hash() ignored the email field")
	}

	if first == (Maintainer{Name: "Other Fynch", Email: "harold@fynch.io"}).Hash() {
		t.Fatal("Hash() ignored the name field")
	}

}

func TestMaintainerIsValid(t *testing.T) {

	valid := Maintainer{Name: "Harold Fynch", Email: "harold@fynch.io"}

	if valid.IsValid() == false {
		t.Fatal("valid maintainer reported invalid")
	}

	invalid := []Maintainer{
		{},
		{Name: "Harold Fynch"},
		{Email: "harold@fynch.io"},
		{Name: "Harold Fynch", Email: "harold.fynch.io"},
	}

	for _, maintainer := range invalid {

		if maintainer.IsValid() == true {
			t.Errorf("IsValid() = true for %#v", maintainer)
		}

	}

}

func TestMaintainerMarshalJSON(t *testing.T) {

	tests := []struct {
		maintainer Maintainer
		expected   string
	}{
		{Maintainer{Name: "Harold Fynch", Email: "harold@fynch.io"}, `"Harold Fynch <harold@fynch.io>"`},
		{Maintainer{Name: "Harold Fynch"}, `"Harold Fynch"`},
		{Maintainer{Email: "harold@fynch.io"}, `"harold@fynch.io"`},
		{Maintainer{}, `""`},
	}

	for _, test := range tests {

		data, err := test.maintainer.MarshalJSON()

		if err != nil {
			t.Errorf("MarshalJSON() returned an error: %v", err)
			continue
		}

		if string(data) != test.expected {
			t.Errorf("MarshalJSON() = %s, expected %s", data, test.expected)
		}

	}

	data, err := json.Marshal(Maintainer{Name: "Harold Fynch", Email: "harold@fynch.io"})

	if err != nil {
		t.Fatalf("json.Marshal() returned an error: %v", err)
	}

	if string(data) != `"Harold Fynch \u003charold@fynch.io\u003e"` {
		t.Fatalf("json.Marshal() = %s, expected an HTML-escaped maintainer", data)
	}

}

func TestMaintainerUnmarshalJSON(t *testing.T) {

	var maintainer Maintainer

	if err := json.Unmarshal([]byte(`"Harold Fynch <harold@fynch.io>"`), &maintainer); err != nil {
		t.Fatalf("UnmarshalJSON() returned an error: %v", err)
	}

	if maintainer.Name != "Harold Fynch" || maintainer.Email != "harold@fynch.io" {
		t.Fatalf("UnmarshalJSON() = %#v, expected name and email", maintainer)
	}

	var email Maintainer

	if err := json.Unmarshal([]byte(`"<harold@fynch.io>"`), &email); err != nil {
		t.Fatalf("UnmarshalJSON() returned an error: %v", err)
	}

	if email.Name != "" || email.Email != "harold@fynch.io" {
		t.Fatalf("UnmarshalJSON() = %#v, expected only an email", email)
	}

	unchanged := Maintainer{Name: "Keep", Email: "keep@example.com"}

	if err := json.Unmarshal([]byte(`"<broken> John"`), &unchanged); err != nil {
		t.Fatalf("UnmarshalJSON() returned an error for an invalid value: %v", err)
	}

	if unchanged.Name != "Keep" || unchanged.Email != "keep@example.com" {
		t.Fatalf("UnmarshalJSON() changed fields for an invalid value: %#v", unchanged)
	}

	if err := maintainer.UnmarshalJSON([]byte("not-quoted")); err == nil {
		t.Fatal("UnmarshalJSON() accepted malformed JSON")
	}

}

func TestMaintainerString(t *testing.T) {

	tests := []struct {
		maintainer Maintainer
		expected   string
	}{
		{Maintainer{Name: "Harold Fynch", Email: "harold@fynch.io"}, "Harold Fynch <harold@fynch.io>"},
		{Maintainer{Name: "Harold Fynch"}, "Harold Fynch"},
		{Maintainer{Email: "harold@fynch.io"}, "harold@fynch.io"},
		{Maintainer{}, ""},
	}

	for _, test := range tests {

		if result := test.maintainer.String(); result != test.expected {
			t.Errorf("String() = %q, expected %q", result, test.expected)
		}

	}

}
