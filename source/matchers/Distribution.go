package matchers

import "tholian-core/types"
import "strings"

type Distribution struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Manager string `json:"manager"`
	Vendor  string `json:"vendor"`
}

func NewDistribution() Distribution {

	var distribution Distribution

	distribution.Name = Any
	distribution.Version = Any
	distribution.Manager = Any
	distribution.Vendor = Any

	return distribution

}

func ToDistribution(value string) Distribution {

	var distribution Distribution

	distribution.Name = Any
	distribution.Version = Any
	distribution.Manager = Any
	distribution.Vendor = Any

	distribution.SetName(value)

	return distribution

}

func (distribution *Distribution) Equal(value Distribution) bool {

	if distribution.Name == value.Name &&
		distribution.Version == value.Version &&
		distribution.Manager == value.Manager &&
		distribution.Vendor == value.Vendor {
		return true
	}

	return false

}

func (distribution *Distribution) IsValid() bool {

	if distribution.Name != Any ||
		distribution.Version != Any ||
		distribution.Manager != Any ||
		distribution.Vendor != Any {
		return true
	}

	return false

}

func (distribution *Distribution) Matches(name string, version string, manager string, vendor string) bool {
	return distribution.MatchesName(name) && distribution.MatchesVersion(version) && distribution.MatchesManager(manager) && distribution.MatchesVendor(vendor)
}

func (distribution *Distribution) MatchesManager(value string) bool {

	if distribution.Manager == value {
		return true
	} else if distribution.Manager == Any {
		return true
	}

	return false

}

func (distribution *Distribution) MatchesName(value string) bool {

	if distribution.Name == value {

		return true

	} else if distribution.Name == Any {

		return true

	} else if strings.HasSuffix(distribution.Name, "-*") {

		prefix := distribution.Name[0 : len(distribution.Name)-2]

		if strings.HasPrefix(value, prefix) {
			return true
		}

	}

	return false

}

func (distribution *Distribution) MatchesVendor(value string) bool {

	if distribution.Vendor == value {
		return true
	} else if distribution.Vendor == Any {
		return true
	}

	return false

}

func (distribution *Distribution) MatchesVersion(value string) bool {

	// Compatibility with "<operator> <version>" syntax
	if strings.Contains(value, " ") {
		value = strings.TrimSpace(value[strings.Index(value, " ")+1:])
	}

	if distribution.Version == Any {

		return true

	} else if strings.HasPrefix(distribution.Version, "<= ") {

		distro_version := types.ToVersion(distribution.Version[3:])
		other_version := types.ToVersion(value)

		if other_version.Equal(distro_version) {
			return true
		} else if other_version.IsBefore(distro_version) {
			return true
		}

	} else if strings.HasPrefix(distribution.Version, "< ") {

		distro_version := types.ToVersion(distribution.Version[2:])
		other_version := types.ToVersion(value)

		if other_version.IsBefore(distro_version) {
			return true
		}

	} else if strings.HasPrefix(distribution.Version, ">= ") {

		distro_version := types.ToVersion(distribution.Version[3:])
		other_version := types.ToVersion(value)

		if other_version.Equal(distro_version) {
			return true
		} else if other_version.IsAfter(distro_version) {
			return true
		}

	} else if strings.HasPrefix(distribution.Version, "> ") {

		distro_version := types.ToVersion(distribution.Version[2:])
		other_version := types.ToVersion(value)

		if other_version.IsAfter(distro_version) {
			return true
		}

	} else {

		distro_version := types.ToVersion(distribution.Version)
		other_version := types.ToVersion(value)

		if other_version.Equal(distro_version) {
			return true
		}

	}

	return false

}

func (distribution *Distribution) SetName(value string) {

	if value == "all" || value == Any || value == "*" {
		distribution.Name = Any
	} else if value != "" {
		distribution.Name = strings.TrimSpace(value)
	}

}

func (distribution *Distribution) SetVendor(value string) {
	distribution.Vendor = strings.TrimSpace(value)
}

func (distribution *Distribution) SetVersion(value string) {

	if value == "all" || value == Any || value == "*" {
		distribution.Version = Any
	} else if value != "" {
		distribution.Version = value
	}

}

func (distribution *Distribution) Hash() string {

	hash := ""

	if distribution.Name != "" {

		hash = distribution.Name

		if distribution.Version != Any {
			hash += "-" + distribution.Version
		}

		if distribution.Vendor != Any {
			hash += "-" + distribution.Vendor
		}

	}

	return hash

}
