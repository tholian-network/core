package matchers

import "tholian-core/types"
import "strings"

type Package struct {
	Name         string `json:"name"`
	Version      string `json:"version"`
	Architecture string `json:"architecture"`
	Manager      string `json:"manager"`
	Vendor       string `json:"vendor"`
}

func NewPackage() Package {

	var pkg Package

	pkg.Name = Any
	pkg.Version = Any
	pkg.Architecture = Any
	pkg.Manager = Any
	pkg.Vendor = Any

	return pkg

}

func ToPackage(value string) Package {

	var pkg Package

	pkg.Name = Any
	pkg.Version = Any
	pkg.Architecture = Any
	pkg.Manager = Any
	pkg.Vendor = Any

	pkg.Parse(value)

	return pkg

}

func (pkg *Package) Equal(value Package) bool {

	if pkg.Name == value.Name &&
		pkg.Version == value.Version &&
		pkg.Architecture == value.Architecture &&
		pkg.Manager == value.Manager &&
		pkg.Vendor == value.Vendor {
		return true
	}

	return false

}

func (pkg *Package) IsValid() bool {

	if pkg.Name != Any ||
		pkg.Version != Any ||
		pkg.Architecture != Any ||
		pkg.Manager != Any ||
		pkg.Vendor != Any {
		return true
	}

	return false

}

func (pkg *Package) Matches(name string, version string, architecture string, manager string, vendor string) bool {
	return pkg.MatchesName(name) && pkg.MatchesVersion(version) && pkg.MatchesArchitecture(architecture) && pkg.MatchesManager(manager) && pkg.MatchesVendor(vendor)
}

func (pkg *Package) MatchesArchitecture(value string) bool {

	if pkg.Architecture == value {
		return true
	} else if pkg.Architecture == Any {
		return true
	}

	return false

}

func (pkg *Package) MatchesManager(value string) bool {

	if pkg.Manager == value {
		return true
	} else if pkg.Manager == Any {
		return true
	}

	return false

}

func (pkg *Package) MatchesName(value string) bool {

	if pkg.Name == value {
		return true
	} else if pkg.Name == Any {
		return true
	}

	return false

}

func (pkg *Package) MatchesVendor(value string) bool {

	if pkg.Vendor == value {
		return true
	} else if pkg.Vendor == Any {
		return true
	}

	return false

}

func (pkg *Package) MatchesVersion(value string) bool {

	// Compatibility with "<operator> <version>" syntax
	if strings.Contains(value, " ") {
		value = strings.TrimSpace(value[strings.Index(value, " ")+1:])
	}

	if pkg.Version == Any {

		return true

	} else if strings.HasPrefix(pkg.Version, "<= ") {

		pkg_version := types.ToVersion(pkg.Version[3:])
		other_version := types.ToVersion(value)

		if other_version.Equal(pkg_version) {
			return true
		} else if other_version.IsBefore(pkg_version) {
			return true
		}

	} else if strings.HasPrefix(pkg.Version, "< ") {

		pkg_version := types.ToVersion(pkg.Version[2:])
		other_version := types.ToVersion(value)

		if other_version.IsBefore(pkg_version) {
			return true
		}

	} else if strings.HasPrefix(pkg.Version, ">= ") {

		pkg_version := types.ToVersion(pkg.Version[3:])
		other_version := types.ToVersion(value)

		if other_version.Equal(pkg_version) {
			return true
		} else if other_version.IsAfter(pkg_version) {
			return true
		}

	} else if strings.HasPrefix(pkg.Version, "> ") {

		pkg_version := types.ToVersion(pkg.Version[2:])
		other_version := types.ToVersion(value)

		if other_version.IsAfter(pkg_version) {
			return true
		}

	} else if strings.HasPrefix(pkg.Version, "= ") {

		pkg_version := types.ToVersion(pkg.Version[2:])
		other_version := types.ToVersion(value)

		if other_version.Equal(pkg_version) {
			return true
		}

	} else {

		pkg_version := types.ToVersion(pkg.Version)
		other_version := types.ToVersion(value)

		if other_version.Equal(pkg_version) {
			return true
		}

	}

	return false

}

func (pkg *Package) Parse(value string) {

	name, version, architecture := parseVersionCondition(value)

	pkg.Name = name
	pkg.Version = version

	if architecture != "" {
		pkg.Architecture = architecture
	}

}

func (pkg *Package) SetArchitecture(value string) {

	architecture := types.ParseArchitecture(value)

	if architecture != nil {
		pkg.Architecture = architecture.String()
	}

}

func (pkg *Package) SetManager(value string) {

	manager := types.ParseManager(value)

	if manager != nil {
		pkg.Manager = manager.String()
	}

}

func (pkg *Package) SetName(value string) {
	pkg.Name = strings.TrimSpace(value)
}

func (pkg *Package) SetVendor(value string) {

	if value == "all" || value == "any" || value == "*" {
		pkg.Vendor = Any
	} else if value != "" {
		pkg.Vendor = value
	}

}

func (pkg *Package) SetVersion(value string) {

	if value == "all" || value == "any" || value == "*" {
		pkg.Version = Any
	} else if value != "" {
		pkg.Version = value
	}

}

func (pkg *Package) Hash() string {

	hash := ""

	if pkg.Name != "" {

		hash = toHash(
			pkg.Name,
			pkg.Version,
			pkg.Architecture,
			pkg.Manager,
			pkg.Vendor,
		)

	}

	return hash

}
