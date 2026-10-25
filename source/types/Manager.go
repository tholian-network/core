package types

import "strconv"

type Manager string

const (
	// Supported Distributions
	ManagerAPK    Manager = "apk"
	ManagerAPT    Manager = "apt"
	ManagerDNF    Manager = "dnf"
	ManagerPACMAN Manager = "pacman"
	ManagerRPM    Manager = "rpm"
	ManagerTDNF   Manager = "tdnf"
	ManagerZYPPER Manager = "zypper"
	ManagerPKG    Manager = "pkg"    // FreeBSD
	ManagerPKGSRC Manager = "pkgsrc" // NetBSD, OpenBSD
	ManagerMSI    Manager = "msi"    // Microsoft Installer

	// Programming Languages
	ManagerCARGO      Manager = "cargo"      // Rust
	ManagerCHOCOLATEY Manager = "chocolatey" // C#
	ManagerCOCOAPODS  Manager = "cocoapods"  // Cocoa
	ManagerCOMPOSER   Manager = "composer"   // PHP
	ManagerCONAN      Manager = "conan"      // C++
	ManagerCONDA      Manager = "conda"      // Python
	ManagerCRAN       Manager = "cran"       // R
	ManagerGEM        Manager = "gem"        // Ruby
	ManagerGO         Manager = "go"         // Go
	ManagerGRADLE     Manager = "gradle"     // Java
	ManagerHACKAGE    Manager = "hackage"    // Haskell
	ManagerHEX        Manager = "hex"        // Erlang
	ManagerMAVEN      Manager = "maven"      // Java
	ManagerNPM        Manager = "npm"        // node.js
	ManagerNUGET      Manager = "nuget"      // C#
	ManagerPEAR       Manager = "pear"       // PHP
	ManagerPHAR       Manager = "phar"       // PHP
	ManagerPIP        Manager = "pip"        // Python

)

func IsManager(value string) bool {

	switch Manager(value) {
	case Any,
		ManagerAPK,
		ManagerAPT,
		ManagerDNF,
		ManagerPACMAN,
		ManagerRPM,
		ManagerTDNF,
		ManagerZYPPER,
		ManagerPKG,
		ManagerPKGSRC,
		ManagerMSI,

		ManagerCARGO,
		ManagerCHOCOLATEY,
		ManagerCOCOAPODS,
		ManagerCOMPOSER,
		ManagerCONAN,
		ManagerCONDA,
		ManagerCRAN,
		ManagerGEM,
		ManagerGO,
		ManagerGRADLE,
		ManagerHACKAGE,
		ManagerHEX,
		ManagerMAVEN,
		ManagerNPM,
		ManagerNUGET,
		ManagerPEAR,
		ManagerPHAR,
		ManagerPIP:
		return true
	}

	return false

}

func NewManager() Manager {
	return Manager(Any)
}

func ParseManager(value string) *Manager {

	var result *Manager = nil
	var manager Manager

	if manager.Parse(value) == true {
		result = &manager
	}

	return result

}

func ToManager(value string) Manager {

	var manager Manager

	manager.Parse(value)

	return manager

}

func (manager *Manager) Parse(value string) bool {

	if IsManager(value) {

		*manager = Manager(value)

		return true

	}

	return false

}

func (manager Manager) Equal(other Manager) bool {
	return manager == other
}

func (manager Manager) String() string {
	return string(manager)
}

func (manager Manager) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(string(manager))), nil
}

func (manager *Manager) UnmarshalJSON(data []byte) error {

	unquoted, err := strconv.Unquote(string(data))

	if err != nil {
		return err
	}

	tmp := ParseManager(unquoted)

	if tmp != nil {
		*manager = *tmp
	}

	return nil

}

func (manager *Manager) IsValid() bool {

	if IsManager(manager.String()) {
		return true
	}

	return false

}
