package types

import "testing"

func TestIsArchitecture(t *testing.T) {

	valid := []Architecture{
		Architecture(Any),
		ArchitectureX86,
		ArchitectureX86_64,
		ArchitectureARMv6,
		ArchitectureARMv7,
		ArchitectureARMv8,
		ArchitectureRISCV32,
		ArchitectureRISCV64,
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
