package types

import (
	"encoding/json"
	"testing"
)

func TestIsArchitecture(t *testing.T) {

	valid := []Architecture{
		Architecture(Any),
		ArchitectureX86,
		ArchitectureX86_64,
		ArchitectureARMv6,
		ArchitectureARMv7,
		ArchitectureARMv8,
		ArchitectureMIPS,
		ArchitectureMIPSEL,
		ArchitectureMIPS64,
		ArchitectureMIPS64EL,
		ArchitecturePPC,
		ArchitecturePPC64,
		ArchitecturePPC64LE,
		ArchitectureRISCV32,
		ArchitectureRISCV64,
		ArchitectureS390X,
		ArchitectureSPARC32,
		ArchitectureSPARC64,
	}

	for _, value := range valid {

		if IsArchitecture(string(value)) == false {
			t.Errorf("IsArchitecture(%q) = false, expected true", value)
		}

	}

	if IsArchitecture("unknown") == true {
		t.Fatal("IsArchitecture() accepted an unknown architecture")
	}

}

func TestNewArchitecture(t *testing.T) {

	architecture := NewArchitecture()

	if architecture != Architecture(Any) {
		t.Fatalf("NewArchitecture() = %q, expected %q", architecture, Any)
	}

}

func TestArchitectureParse(t *testing.T) {

	tests := []struct {
		input    string
		expected Architecture
	}{
		{"*", Architecture(Any)},
		{"all", Architecture(Any)},
		{"any", Architecture(Any)},
		{"noarch", Architecture(Any)},
		{"386", ArchitectureX86},
		{"i386", ArchitectureX86},
		{"i686", ArchitectureX86},
		{"x32", ArchitectureX86},
		{"x86", ArchitectureX86},
		{"x86_32", ArchitectureX86},
		{"x86-32", ArchitectureX86},
		{"amd64", ArchitectureX86_64},
		{"ia64", ArchitectureX86_64},
		{"x64", ArchitectureX86_64},
		{"x86_64", ArchitectureX86_64},
		{"x86-64", ArchitectureX86_64},
		{"arm", ArchitectureARMv6},
		{"armel", ArchitectureARMv6},
		{"armv6", ArchitectureARMv6},
		{"armv6l", ArchitectureARMv6},
		{"armhf", ArchitectureARMv7},
		{"armv7", ArchitectureARMv7},
		{"armv7h", ArchitectureARMv7},
		{"armv7l", ArchitectureARMv7},
		{"aarch64", ArchitectureARMv8},
		{"armv8", ArchitectureARMv8},
		{"armv8l", ArchitectureARMv8},
		{"arm64", ArchitectureARMv8},
		{"mips", ArchitectureMIPS},
		{"mipsel", ArchitectureMIPSEL},
		{"mips64", ArchitectureMIPS64},
		{"mips64el", ArchitectureMIPS64EL},
		{"powerpc", ArchitecturePPC},
		{"ppc", ArchitecturePPC},
		{"powerpc64", ArchitecturePPC64},
		{"ppc64", ArchitecturePPC64},
		{"powerpc64le", ArchitecturePPC64LE},
		{"ppc64le", ArchitecturePPC64LE},
		{"s390x", ArchitectureS390X},
		{"riscv", ArchitectureRISCV32},
		{"riscv32", ArchitectureRISCV32},
		{"riscv64", ArchitectureRISCV64},
		{"sparc", ArchitectureSPARC32},
		{"sparc32", ArchitectureSPARC32},
		{"sparc64", ArchitectureSPARC64},
		{"mipsel_24kc", ArchitectureMIPSEL},
		{"mips_24kc", ArchitectureMIPS},
		{"mips64_octeon", ArchitectureMIPS64},
		{"mips64el_octeon", ArchitectureMIPS64EL},
		{"arm_cortex-a7", ArchitectureARMv7},
		{"arm_cortex-a15", ArchitectureARMv7},
		{"arm_cortex-a53", ArchitectureARMv8},
		{"aarch64_cortex-a53", ArchitectureARMv8},
		{"i386_pentium4", ArchitectureX86},
		{"x86_64-linux", ArchitectureX86_64},
		{"aarch64-linux", ArchitectureARMv8},
		{"armv7l-linux", ArchitectureARMv7},
		{"i686-linux", ArchitectureX86},
		{"riscv64-linux", ArchitectureRISCV64},
		{"powerpc64le-linux", ArchitecturePPC64LE},
		{"s390x-linux", ArchitectureS390X},
		{"aarch64-darwin", ArchitectureARMv8},
		{"x86_64-freebsd", ArchitectureX86_64},
		{"some 64Bit build", ArchitectureX86_64},
		{"some 64 Bit build", ArchitectureX86_64},
		{"some 64-Bit build", ArchitectureX86_64},
		{"some 32Bit build", ArchitectureX86},
		{"some 32 Bit build", ArchitectureX86},
		{"some 32-Bit build", ArchitectureX86},
	}

	for _, test := range tests {

		var architecture Architecture

		if architecture.Parse(test.input) == false {
			t.Errorf("Parse(%q) = false, expected true", test.input)
		} else if architecture != test.expected {
			t.Errorf("Parse(%q) = %q, expected %q", test.input, architecture, test.expected)
		}

	}

	invalid := []string{"", "unknown", "asdf 64"}

	for _, value := range invalid {

		var architecture Architecture

		if architecture.Parse(value) == true {
			t.Errorf("Parse(%q) = true, expected false", value)
		}

	}

}

func TestParseArchitecture(t *testing.T) {

	architecture := ParseArchitecture("amd64")

	if architecture == nil || *architecture != ArchitectureX86_64 {
		t.Fatalf("ParseArchitecture(%q) = %v, expected %q", "amd64", architecture, ArchitectureX86_64)
	}

	if ParseArchitecture("unknown") != nil {
		t.Fatal("ParseArchitecture() returned a value for an invalid architecture")
	}

}

func TestToArchitecture(t *testing.T) {

	if ToArchitecture("amd64") != ArchitectureX86_64 {
		t.Fatalf("ToArchitecture(%q) = %q, expected %q", "amd64", ToArchitecture("amd64"), ArchitectureX86_64)
	}

	if ToArchitecture("aarch64") != ArchitectureARMv8 {
		t.Fatalf("ToArchitecture(%q) = %q, expected %q", "aarch64", ToArchitecture("aarch64"), ArchitectureARMv8)
	}

	if ToArchitecture("unknown") != Architecture("") {
		t.Fatalf("ToArchitecture(%q) = %q, expected zero value", "unknown", ToArchitecture("unknown"))
	}

}

func TestArchitectureBytes(t *testing.T) {

	bytes := ArchitectureX86_64.Bytes()

	if string(bytes) != "x86_64" {
		t.Fatalf("Bytes() = %q, expected %q", string(bytes), "x86_64")
	}

}

func TestArchitectureEqual(t *testing.T) {

	if ArchitectureX86.Equal(ArchitectureX86) == false {
		t.Fatal("Equal() returned false for identical architectures")
	}

	if ArchitectureX86.Equal(ArchitectureX86_64) == true {
		t.Fatal("Equal() returned true for different architectures")
	}

	if NewArchitecture().Equal(Architecture(Any)) == false {
		t.Fatal("Equal() returned false for the any architecture")
	}

}

func TestArchitectureHash(t *testing.T) {

	first := ArchitectureX86_64.Hash()
	second := ArchitectureX86_64.Hash()

	if first == "" {
		t.Fatal("Hash() returned an empty string")
	}

	if first != second {
		t.Fatalf("Hash() is not deterministic: %q != %q", first, second)
	}

	if first == ArchitectureX86.Hash() {
		t.Fatal("Hash() returned the same value for different architectures")
	}

}

func TestArchitectureString(t *testing.T) {

	if ArchitectureX86_64.String() != "x86_64" {
		t.Fatalf("String() = %q, expected %q", ArchitectureX86_64.String(), "x86_64")
	}

	if NewArchitecture().String() != Any {
		t.Fatalf("String() = %q, expected %q", NewArchitecture().String(), Any)
	}

}

func TestArchitectureMarshalJSON(t *testing.T) {

	data, err := ArchitectureX86_64.MarshalJSON()

	if err != nil {
		t.Fatalf("MarshalJSON() returned an error: %v", err)
	}

	if string(data) != `"x86_64"` {
		t.Fatalf("MarshalJSON() = %s, expected %q", data, `"x86_64"`)
	}

	data, err = json.Marshal(NewArchitecture())

	if err != nil {
		t.Fatalf("json.Marshal() returned an error: %v", err)
	}

	if string(data) != `"any"` {
		t.Fatalf("json.Marshal() = %s, expected %q", data, `"any"`)
	}

}

func TestArchitectureUnmarshalJSON(t *testing.T) {

	var architecture Architecture

	if err := json.Unmarshal([]byte(`"amd64"`), &architecture); err != nil {
		t.Fatalf("UnmarshalJSON() returned an error: %v", err)
	}

	if architecture != ArchitectureX86_64 {
		t.Fatalf("UnmarshalJSON() = %q, expected %q", architecture, ArchitectureX86_64)
	}

	var zero Architecture

	if err := json.Unmarshal([]byte(`"unknown"`), &zero); err != nil {
		t.Fatalf("UnmarshalJSON() returned an error for an invalid value: %v", err)
	}

	if zero != Architecture("") {
		t.Fatalf("UnmarshalJSON() changed the value for an invalid architecture: %q", zero)
	}

	if err := zero.UnmarshalJSON([]byte("not-quoted")); err == nil {
		t.Fatal("UnmarshalJSON() accepted malformed JSON")
	}

}

func TestArchitectureIsValid(t *testing.T) {

	valid := []Architecture{
		Architecture(Any),
		ArchitectureX86,
		ArchitectureX86_64,
		ArchitectureARMv6,
		ArchitectureARMv7,
		ArchitectureARMv8,
		ArchitectureMIPS,
		ArchitectureMIPSEL,
		ArchitectureMIPS64,
		ArchitectureMIPS64EL,
		ArchitecturePPC,
		ArchitecturePPC64,
		ArchitecturePPC64LE,
		ArchitectureRISCV32,
		ArchitectureRISCV64,
		ArchitectureS390X,
		ArchitectureSPARC32,
		ArchitectureSPARC64,
	}

	for _, architecture := range valid {

		if architecture.IsValid() == false {
			t.Errorf("IsValid() = false for %q", architecture)
		}

	}

	invalid := []Architecture{"", "unknown"}

	for _, architecture := range invalid {

		if architecture.IsValid() == true {
			t.Errorf("IsValid() = true for %q", architecture)
		}

	}

}
