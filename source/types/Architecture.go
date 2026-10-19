package types

import "strconv"
import "strings"

type Architecture string

const (
	ArchitectureANY     Architecture = "any"
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

	var result bool

	if value == string(ArchitectureANY) {
		result = true
	} else if value == string(ArchitectureX86) {
		result = true
	} else if value == string(ArchitectureANY) {
		result = true
	} else if value == string(ArchitectureX86) {
		result = true
	} else if value == string(ArchitectureX86_64) {
		result = true
	} else if value == string(ArchitectureARMv6) {
		result = true
	} else if value == string(ArchitectureARMv7) {
		result = true
	} else if value == string(ArchitectureARMv8) {
		result = true
	} else if value == string(ArchitectureRISCV32) {
		result = true
	} else if value == string(ArchitectureRISCV64) {
		result = true
	} else if value == string(ArchitectureSPARC32) {
		result = true
	} else if value == string(ArchitectureSPARC64) {
		result = true
	}

	return result

}

func ParseArchitecture(value string) *Architecture {

	var result *Architecture = nil

	if value == "*" || value == "all" || value == "any" || value == "noarch" {
		architecture := Architecture(ArchitectureANY)
		result = &architecture
	} else if value == "386" || value == "i386" || value == "i686" || value == "x32" || value == "x86" || value == "x86_32" || value == "x86-32" {
		architecture := Architecture(ArchitectureX86)
		result = &architecture
	} else if value == "amd64" || value == "ia64" || value == "x64" || value == "x86_64" || value == "x86-64" {
		architecture := Architecture(ArchitectureX86_64)
		result = &architecture
	} else if value == "arm" || value == "armel" || value == "armv6" {
		architecture := Architecture(ArchitectureARMv6)
		result = &architecture
	} else if value == "armhf" || value == "armv7" || value == "armv7h" {
		architecture := Architecture(ArchitectureARMv7)
		result = &architecture
	} else if value == "aarch64" || value == "armv8" || value == "arm64" {
		architecture := Architecture(ArchitectureARMv8)
		result = &architecture
	} else if value == "riscv" || value == "riscv32" {
		architecture := Architecture(ArchitectureRISCV32)
		result = &architecture
	} else if value == "riscv64" {
		architecture := Architecture(ArchitectureRISCV64)
		result = &architecture
	} else if value == "sparc" || value == "sparc32" {
		architecture := Architecture(ArchitectureSPARC32)
		result = &architecture
	} else if value == "sparc64" {
		architecture := Architecture(ArchitectureSPARC64)
		result = &architecture
	} else {

		if strings.Contains(value, "64Bit") || strings.Contains(value, "64 Bit") || strings.Contains(value, "64-Bit") || strings.Contains(value, "x64") || strings.Contains(value, "x86_64") {
			architecture := Architecture(ArchitectureX86_64)
			result = &architecture
		} else if strings.Contains(value, "32Bit") || strings.Contains(value, "32 Bit") || strings.Contains(value, "32-Bit") || strings.Contains(value, "x32") || strings.Contains(value, "x86") {
			architecture := Architecture(ArchitectureX86)
			result = &architecture
		} else {
			architecture := Architecture(ArchitectureANY)
			result = &architecture
		}

	}

	return result

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

func (architecture *Architecture) IsValid() bool {

	var result bool

	if IsArchitecture(architecture.String()) {
		result = true
	}

	return result

}
