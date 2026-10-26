package types

import "strconv"
import "strings"

type Architecture string

const (
	ArchitectureX86     Architecture = "x86"
	ArchitectureX86_64  Architecture = "x86_64"
	ArchitectureARMv6   Architecture = "armv6"
	ArchitectureARMv7   Architecture = "armv7"
	ArchitectureARMv8   Architecture = "armv8"
	ArchitectureRISCV32 Architecture = "riscv32"
	ArchitectureRISCV64 Architecture = "riscv64"
	ArchitectureSPARC32 Architecture = "sparc"
	ArchitectureSPARC64 Architecture = "sparc64"
)

func IsArchitecture(value string) bool {

	switch Architecture(value) {
	case Any,
		ArchitectureX86,
		ArchitectureX86_64,
		ArchitectureARMv6,
		ArchitectureARMv7,
		ArchitectureARMv8,
		ArchitectureRISCV32,
		ArchitectureRISCV64,
		ArchitectureSPARC32,
		ArchitectureSPARC64:
		return true
	}

	return false

}

func NewArchitecture() Architecture {
	return Architecture(Any)
}

func ParseArchitecture(value string) *Architecture {

	var result *Architecture = nil
	var architecture Architecture

	if architecture.Parse(value) == true {
		result = &architecture
	}

	return result

}

func (architecture *Architecture) Parse(value string) bool {

	if value == "*" || value == "all" || value == Any || value == "noarch" {

		*architecture = Architecture(Any)

		return true

	} else if value == "386" || value == "i386" || value == "i686" || value == "x32" || value == "x86" || value == "x86_32" || value == "x86-32" {

		*architecture = ArchitectureX86

		return true

	} else if value == "amd64" || value == "ia64" || value == "x64" || value == "x86_64" || value == "x86-64" {

		*architecture = ArchitectureX86_64

		return true

	} else if value == "arm" || value == "armel" || value == "armv6" {

		*architecture = ArchitectureARMv6

		return true

	} else if value == "armhf" || value == "armv7" || value == "armv7h" {

		*architecture = ArchitectureARMv7

		return true

	} else if value == "aarch64" || value == "armv8" || value == "arm64" {

		*architecture = ArchitectureARMv8

		return true

	} else if value == "riscv" || value == "riscv32" {

		*architecture = ArchitectureRISCV32

		return true

	} else if value == "riscv64" {

		*architecture = ArchitectureRISCV64

		return true

	} else if value == "sparc" || value == "sparc32" {

		*architecture = ArchitectureSPARC32

		return true

	} else if value == "sparc64" {

		*architecture = ArchitectureSPARC64

		return true

	} else if strings.Contains(value, "64Bit") || strings.Contains(value, "64 Bit") || strings.Contains(value, "64-Bit") || strings.Contains(value, "x64") || strings.Contains(value, "x86_64") {

		*architecture = ArchitectureX86_64

		return true

	} else if strings.Contains(value, "32Bit") || strings.Contains(value, "32 Bit") || strings.Contains(value, "32-Bit") || strings.Contains(value, "x32") || strings.Contains(value, "x86") {

		*architecture = ArchitectureX86

		return true

	}

	return false

}

func ToArchitecture(value string) Architecture {

	var architecture Architecture

	tmp := ParseArchitecture(value)

	if tmp != nil {
		architecture = *tmp
	}

	return architecture

}

func (architecture Architecture) Bytes() []byte {
	return []byte(architecture)
}

func (architecture Architecture) Equal(other Architecture) bool {
	return architecture == other
}

func (architecture Architecture) Hash() string {
	return toHash(architecture.String())
}

func (architecture Architecture) String() string {
	return string(architecture)
}

func (architecture Architecture) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(string(architecture))), nil
}

func (architecture *Architecture) UnmarshalJSON(data []byte) error {

	unquoted, err := strconv.Unquote(string(data))

	if err != nil {
		return err
	}

	tmp := ParseArchitecture(unquoted)

	if tmp != nil {
		*architecture = *tmp
	}

	return nil

}

func (architecture Architecture) IsValid() bool {

	var result bool

	if IsArchitecture(architecture.String()) {
		result = true
	}

	return result

}
