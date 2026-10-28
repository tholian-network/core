package types

import "strconv"
import "strings"

type Architecture string

const (
	ArchitectureX86      Architecture = "x86"
	ArchitectureX86_64   Architecture = "x86_64"
	ArchitectureARMv6    Architecture = "armv6"
	ArchitectureARMv7    Architecture = "armv7"
	ArchitectureARMv8    Architecture = "armv8"
	ArchitectureMIPS     Architecture = "mips"
	ArchitectureMIPSEL   Architecture = "mipsel"
	ArchitectureMIPS64   Architecture = "mips64"
	ArchitectureMIPS64EL Architecture = "mips64el"
	ArchitecturePPC      Architecture = "ppc"
	ArchitecturePPC64    Architecture = "ppc64"
	ArchitecturePPC64LE  Architecture = "ppc64le"
	ArchitectureRISCV32  Architecture = "riscv32"
	ArchitectureRISCV64  Architecture = "riscv64"
	ArchitectureS390X    Architecture = "s390x"
	ArchitectureSPARC32  Architecture = "sparc"
	ArchitectureSPARC64  Architecture = "sparc64"
)

func IsArchitecture(value string) bool {

	switch Architecture(value) {
	case Any,
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

	} else {

		switch {
		case strings.Contains(value, "amd64"),
			strings.Contains(value, "ia64"),
			strings.Contains(value, "x64"),
			strings.Contains(value, "x86_64"),
			strings.Contains(value, "x86-64"),
			strings.Contains(value, "64Bit"),
			strings.Contains(value, "64 Bit"),
			strings.Contains(value, "64-Bit"):

			*architecture = ArchitectureX86_64
			return true

		case strings.Contains(value, "386"),
			strings.Contains(value, "i386"),
			strings.Contains(value, "i686"),
			strings.Contains(value, "x32"),
			strings.Contains(value, "x86"),
			strings.Contains(value, "x86_32"),
			strings.Contains(value, "x86-32"),
			strings.Contains(value, "32Bit"),
			strings.Contains(value, "32 Bit"),
			strings.Contains(value, "32-Bit"):

			*architecture = ArchitectureX86
			return true

		case strings.Contains(value, "aarch64"),
			strings.Contains(value, "armv8"),
			strings.Contains(value, "armv8l"),
			strings.Contains(value, "arm64"),
			strings.Contains(value, "cortex-a53"),
			strings.Contains(value, "cortex-a72"),
			strings.Contains(value, "cortex-a73"):

			*architecture = ArchitectureARMv8
			return true

		case strings.Contains(value, "armhf"),
			strings.Contains(value, "armv7"),
			strings.Contains(value, "armv7h"),
			strings.Contains(value, "armv7l"),
			strings.Contains(value, "cortex-a7"),
			strings.Contains(value, "cortex-a15"):

			*architecture = ArchitectureARMv7
			return true

		case strings.Contains(value, "arm"),
			strings.Contains(value, "armel"),
			strings.Contains(value, "armv6"),
			strings.Contains(value, "armv6l"):

			*architecture = ArchitectureARMv6
			return true

		case strings.Contains(value, "mips64el"):

			*architecture = ArchitectureMIPS64EL
			return true

		case strings.Contains(value, "mips64"):

			*architecture = ArchitectureMIPS64
			return true

		case strings.Contains(value, "mipsel"):

			*architecture = ArchitectureMIPSEL
			return true

		case strings.Contains(value, "mips"):

			*architecture = ArchitectureMIPS
			return true

		case strings.Contains(value, "powerpc64le"),
			strings.Contains(value, "ppc64le"):

			*architecture = ArchitecturePPC64LE
			return true

		case strings.Contains(value, "powerpc64"),
			strings.Contains(value, "ppc64"):

			*architecture = ArchitecturePPC64
			return true

		case strings.Contains(value, "powerpc"),
			strings.Contains(value, "ppc"):

			*architecture = ArchitecturePPC
			return true

		case strings.Contains(value, "riscv64"):

			*architecture = ArchitectureRISCV64
			return true

		case strings.Contains(value, "riscv"),
			strings.Contains(value, "riscv32"):

			*architecture = ArchitectureRISCV32
			return true

		case strings.Contains(value, "s390x"):

			*architecture = ArchitectureS390X
			return true

		case strings.Contains(value, "sparc64"):

			*architecture = ArchitectureSPARC64
			return true

		case strings.Contains(value, "sparc"),
			strings.Contains(value, "sparc32"):

			*architecture = ArchitectureSPARC32
			return true

		}

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
