package structs

import "testing"

func TestNewWeakness(t *testing.T) {

	weakness := NewWeakness()

	if weakness.Name != "" || weakness.Description != "" {
		t.Errorf("NewWeakness() = %+v, expected empty strings", weakness)
	}

	if weakness.Impacts == nil || len(weakness.Impacts) != 0 {
		t.Errorf("NewWeakness() Impacts = %#v, expected an empty slice", weakness.Impacts)
	}

	if weakness.References == nil || len(weakness.References) != 0 {
		t.Errorf("NewWeakness() References = %#v, expected an empty slice", weakness.References)
	}

	if weakness.Scopes == nil || len(weakness.Scopes) != 0 {
		t.Errorf("NewWeakness() Scopes = %#v, expected an empty slice", weakness.Scopes)
	}

}

func TestToWeakness(t *testing.T) {

	weakness := ToWeakness("CWE-79  ")

	if weakness.Name != "CWE-79" {
		t.Errorf("ToWeakness() Name = %q, expected %q", weakness.Name, "CWE-79")
	}

	if weakness.Impacts == nil || weakness.References == nil || weakness.Scopes == nil {
		t.Fatal("ToWeakness() did not initialize slices")
	}

	invalid := []string{"CWE", "cwe-79", "  CWE-79", "79", ""}

	for _, value := range invalid {

		weakness := ToWeakness(value)

		if weakness.Name != "" {
			t.Errorf("ToWeakness(%q) accepted an invalid value: %q", value, weakness.Name)
		}

	}

}

func TestWeaknessSetters(t *testing.T) {

	weakness := NewWeakness()

	weakness.SetName("CWE-79  ")
	weakness.SetDescription("  Cross-site scripting  ")

	if weakness.Name != "CWE-79" {
		t.Errorf("SetName() = %q, expected %q", weakness.Name, "CWE-79")
	}

	if weakness.Description != "Cross-site scripting" {
		t.Errorf("SetDescription() = %q, expected %q", weakness.Description, "Cross-site scripting")
	}

	weakness.SetName("not-a-cwe")
	weakness.SetName("")

	if weakness.Name != "CWE-79" {
		t.Errorf("SetName() accepted an invalid value: %q", weakness.Name)
	}

}

func TestWeaknessIsValid(t *testing.T) {

	cases := []struct {
		name     string
		weakness Weakness
		want     bool
	}{
		{"default", Weakness{}, false},
		{"name only", func() Weakness {
			value := NewWeakness()
			value.SetName("CWE-79")
			return value
		}(), false},
		{"name and impact", func() Weakness {
			value := NewWeakness()
			value.SetName("CWE-79")
			value.AddImpact("Execute Unauthorized Code or Commands")
			return value
		}(), true},
		{"name and scope", func() Weakness {
			value := NewWeakness()
			value.SetName("CWE-79")
			value.AddScope("Confidentiality")
			return value
		}(), true},
		{"description and reference only", func() Weakness {
			value := NewWeakness()
			value.SetName("CWE-79")
			value.SetDescription("description")
			value.AddReference("https://example.com")
			return value
		}(), false},
		{"impact without name", func() Weakness {
			value := NewWeakness()
			value.AddImpact("Modify Memory")
			return value
		}(), false},
	}

	for _, c := range cases {

		if c.weakness.IsValid() != c.want {
			t.Errorf("%s: IsValid() = %v, expected %v", c.name, c.weakness.IsValid(), c.want)
		}

	}

}

func TestWeaknessImpacts(t *testing.T) {

	weakness := NewWeakness()

	impact := "Execute Unauthorized Code or Commands"

	weakness.AddImpact(impact)
	weakness.AddImpact(impact)
	weakness.AddImpact("Bogus Impact")

	if len(weakness.Impacts) != 1 || weakness.Impacts[0] != impact {
		t.Fatalf("AddImpact() = %#v, expected one allowed impact", weakness.Impacts)
	}

	weakness.AddImpact("Read Application Data")

	if len(weakness.Impacts) != 2 {
		t.Fatalf("AddImpact() = %#v, expected two impacts", weakness.Impacts)
	}

	weakness.RemoveImpact(impact)

	if len(weakness.Impacts) != 1 || weakness.Impacts[0] != "Read Application Data" {
		t.Fatalf("RemoveImpact() = %#v, expected the remaining impact", weakness.Impacts)
	}

	weakness.RemoveImpact("missing")

	weakness.SetImpacts([]string{
		"Modify Memory",
		"Bogus Impact",
		"Modify Memory",
	})

	if len(weakness.Impacts) != 2 {
		t.Fatalf("SetImpacts() = %#v, expected two allowed entries", weakness.Impacts)
	}

	for _, value := range weakness.Impacts {

		if value != "Modify Memory" {
			t.Errorf("SetImpacts() kept an invalid impact: %q", value)
		}

	}

}

func TestWeaknessReferences(t *testing.T) {

	weakness := NewWeakness()

	weakness.AddReference("https://example.com/advisory")
	weakness.AddReference("https://example.com/advisory")
	weakness.AddReference("http://example.com/advisory")
	weakness.AddReference("ftp://example.com/advisory")
	weakness.AddReference("example.com/advisory")

	if len(weakness.References) != 2 {
		t.Fatalf("AddReference() = %#v, expected two URL references", weakness.References)
	}

	weakness.RemoveReference("https://example.com/advisory")

	if len(weakness.References) != 1 || weakness.References[0] != "http://example.com/advisory" {
		t.Fatalf("RemoveReference() = %#v, expected the remaining reference", weakness.References)
	}

	weakness.RemoveReference("missing")

	weakness.SetReferences([]string{
		"https://example.com/one",
		"ftp://example.com/two",
		"http://example.com/three",
		"example.com/four",
	})

	if len(weakness.References) != 2 {
		t.Fatalf("SetReferences() = %#v, expected two URL references", weakness.References)
	}

	if weakness.References[0] != "https://example.com/one" || weakness.References[1] != "http://example.com/three" {
		t.Errorf("SetReferences() = %#v, expected the http and https references", weakness.References)
	}

}

func TestWeaknessScopes(t *testing.T) {

	weakness := NewWeakness()

	weakness.AddScope("Confidentiality")
	weakness.AddScope("Confidentiality")
	weakness.AddScope("Bogus Scope")
	weakness.AddScope("Integrity")

	if len(weakness.Scopes) != 2 {
		t.Fatalf("AddScope() = %#v, expected two allowed scopes", weakness.Scopes)
	}

	weakness.RemoveScope("Confidentiality")

	if len(weakness.Scopes) != 1 || weakness.Scopes[0] != "Integrity" {
		t.Fatalf("RemoveScope() = %#v, expected the remaining scope", weakness.Scopes)
	}

	weakness.RemoveScope("missing")

	weakness.SetScopes([]string{
		"Confidentiality",
		"Bogus Scope",
		"Availability",
	})

	if len(weakness.Scopes) != 2 {
		t.Fatalf("SetScopes() = %#v, expected two allowed scopes", weakness.Scopes)
	}

	if weakness.Scopes[0] != "Confidentiality" || weakness.Scopes[1] != "Availability" {
		t.Errorf("SetScopes() = %#v, expected the allowed scopes", weakness.Scopes)
	}

}

func TestWeaknessHasImpactAndReference(t *testing.T) {

	weakness := NewWeakness()

	impact := "Execute Unauthorized Code or Commands"
	reference := "https://example.com/advisory"

	if weakness.HasImpact(impact) == true || weakness.HasReference(reference) == true {
		t.Fatal("HasImpact() or HasReference() found a missing value")
	}

	weakness.AddImpact(impact)
	weakness.AddReference(reference)

	if weakness.HasImpact(impact) == false {
		t.Fatal("HasImpact() did not find an existing impact")
	}

	if weakness.HasImpact("Modify Memory") == true {
		t.Fatal("HasImpact() found a missing impact")
	}

	if weakness.HasReference(reference) == false {
		t.Fatal("HasReference() did not find an existing reference")
	}

	if weakness.HasReference("https://example.com/other") == true {
		t.Fatal("HasReference() found a missing reference")
	}

}

func TestWeaknessEqual(t *testing.T) {

	base := NewWeakness()
	base.SetName("CWE-79")
	base.SetDescription("Cross-site scripting")
	base.AddImpact("Execute Unauthorized Code or Commands")
	base.AddReference("https://example.com/advisory")
	base.AddScope("Confidentiality")

	if base.Equal(base) == false {
		t.Fatal("a weakness should equal itself")
	}

	cases := []struct {
		name  string
		value Weakness
	}{
		{"name", func() Weakness { value := base; value.Name = "CWE-80"; return value }()},
		{"description", func() Weakness { value := base; value.Description = "other"; return value }()},
		{"impacts", func() Weakness { value := base; value.AddImpact("Modify Memory"); return value }()},
		{"references", func() Weakness { value := base; value.AddReference("https://example.com/other"); return value }()},
		{"scopes", func() Weakness { value := base; value.AddScope("Integrity"); return value }()},
	}

	for _, c := range cases {

		if base.Equal(c.value) == true {
			t.Errorf("Equal() returned true for a different %s", c.name)
		}

	}

}

func TestWeaknessHash(t *testing.T) {

	base := NewWeakness()
	base.SetName("CWE-79")
	base.SetDescription("Cross-site scripting")
	base.AddImpact("Execute Unauthorized Code or Commands")
	base.AddReference("https://example.com/advisory")
	base.AddScope("Confidentiality")

	hash := base.Hash()

	if hash == "" {
		t.Fatal("Hash() returned an empty string for a populated weakness")
	}

	if base.Hash() != hash {
		t.Fatal("Hash() is not deterministic")
	}

	empty := Weakness{}

	if empty.Hash() != "" {
		t.Fatal("Hash() should be empty without a name")
	}

	cases := []struct {
		name  string
		value Weakness
	}{
		{"name", func() Weakness { value := base; value.Name = "CWE-80"; return value }()},
		{"description", func() Weakness { value := base; value.Description = "other"; return value }()},
		{"impacts", func() Weakness { value := base; value.AddImpact("Modify Memory"); return value }()},
		{"references", func() Weakness { value := base; value.AddReference("https://example.com/other"); return value }()},
		{"scopes", func() Weakness { value := base; value.AddScope("Integrity"); return value }()},
	}

	for _, c := range cases {

		if c.value.Hash() == hash {
			t.Errorf("Hash() did not change for a different %s", c.name)
		}

	}

}
