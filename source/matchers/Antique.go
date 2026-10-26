package matchers

import "tholian-core/types"
import "encoding/binary"
import "encoding/hex"
import "hash/crc32"
import "strings"

type Antique struct {
	Name         string `json:"name"`
	Version      string `json:"version"`
	Architecture string `json:"architecture"`
	Manager      string `json:"manager"`
	Vendor       string `json:"vendor"`
	Service      string `json:"service"`
}

func NewAntique() Antique {

	var antique Antique

	antique.Name = Any
	antique.Version = Any
	antique.Architecture = Any
	antique.Manager = Any
	antique.Service = Any
	antique.Vendor = Any

	return antique

}

func ToAntique(value string) Antique {

	var antique Antique

	antique.Name = Any
	antique.Version = Any
	antique.Architecture = Any
	antique.Manager = Any
	antique.Vendor = Any
	antique.Service = Any

	antique.Parse(value)

	return antique

}

func (antique *Antique) Equal(value Antique) bool {

	if antique.Name == value.Name &&
		antique.Version == value.Version &&
		antique.Architecture == value.Architecture &&
		antique.Manager == value.Manager &&
		antique.Vendor == value.Vendor &&
		antique.Service == value.Service {
		return true
	}

	return false

}

func (antique *Antique) IsValid() bool {

	if antique.Name != Any ||
		antique.Version != Any ||
		antique.Architecture != Any ||
		antique.Manager != Any ||
		antique.Vendor != Any ||
		antique.Service != Any {
		return true
	}

	return false

}

func (antique *Antique) Matches(name string, version string, architecture string, manager string, vendor string, service string) bool {
	return antique.MatchesName(name) && antique.MatchesVersion(version) && antique.MatchesArchitecture(architecture) && antique.MatchesManager(manager) && antique.MatchesVendor(vendor) && antique.MatchesService(service)
}

func (antique *Antique) MatchesArchitecture(value string) bool {

	if antique.Architecture == value {
		return true
	} else if antique.Architecture == Any {
		return true
	}

	return false

}

func (antique *Antique) MatchesManager(value string) bool {

	if antique.Manager == value {
		return true
	} else if antique.Manager == Any {
		return true
	}

	return false

}

func (antique *Antique) MatchesName(value string) bool {

	if antique.Name == value {
		return true
	} else if antique.Name == Any {
		return true
	}

	return false

}

func (antique *Antique) MatchesService(value string) bool {

	if antique.Service == value {
		return true
	} else if antique.Service == Any {
		return true
	}

	return false

}

func (antique *Antique) MatchesVendor(value string) bool {

	if antique.Vendor == value {
		return true
	} else if antique.Vendor == Any {
		return true
	}

	return false

}

func (antique *Antique) MatchesVersion(value string) bool {

	// Compatibility with "<operator> <version>" syntax
	if strings.Contains(value, " ") {
		value = strings.TrimSpace(value[strings.Index(value, " ")+1:])
	}

	if antique.Version == Any {

		return true

	} else if strings.HasPrefix(antique.Version, "<= ") {

		antique_version := types.ToVersion(antique.Version[3:])
		other_version := types.ToVersion(value)

		if other_version.Equal(antique_version) {
			return true
		} else if other_version.IsBefore(antique_version) {
			return true
		}

	} else if strings.HasPrefix(antique.Version, "< ") {

		antique_version := types.ToVersion(antique.Version[2:])
		other_version := types.ToVersion(value)

		if other_version.IsBefore(antique_version) {
			return true
		}

	} else if strings.HasPrefix(antique.Version, ">= ") {

		antique_version := types.ToVersion(antique.Version[3:])
		other_version := types.ToVersion(value)

		if other_version.Equal(antique_version) {
			return true
		} else if other_version.IsAfter(antique_version) {
			return true
		}

	} else if strings.HasPrefix(antique.Version, "> ") {

		antique_version := types.ToVersion(antique.Version[2:])
		other_version := types.ToVersion(value)

		if other_version.IsAfter(antique_version) {
			return true
		}

	} else if strings.HasPrefix(antique.Version, "= ") {

		antique_version := types.ToVersion(antique.Version[2:])
		other_version := types.ToVersion(value)

		if other_version.Equal(antique_version) {
			return true
		}

	} else {

		antique_version := types.ToVersion(antique.Version)
		other_version := types.ToVersion(value)

		if other_version.Equal(antique_version) {
			return true
		}

	}

	return false

}

func (antique *Antique) Parse(value string) {

	name, version, architecture := parseVersionCondition(value)

	antique.Name = name
	antique.Version = version

	if architecture != "" {
		antique.Architecture = architecture
	}

}

func (antique *Antique) SetArchitecture(value string) {

	architecture := types.ParseArchitecture(value)

	if architecture != nil {
		antique.Architecture = architecture.String()
	}

}

func (antique *Antique) SetManager(value string) {

	manager := types.ParseManager(value)

	if manager != nil {
		antique.Manager = manager.String()
	}

}

func (antique *Antique) SetName(value string) {
	antique.Name = strings.TrimSpace(value)
}

func (antique *Antique) SetService(value string) {

	if value == "all" || value == Any || value == "*" {
		antique.Service = Any
	} else if value != "" {
		antique.Service = value
	}

}

func (antique *Antique) SetVersion(value string) {

	if value == "all" || value == Any || value == "*" {
		antique.Version = Any
	} else if value != "" {
		antique.Version = value
	}

}

func (antique *Antique) Hash() string {

	hash := ""

	if antique.Name != "" {

		checksum := crc32.ChecksumIEEE([]byte(strings.Join([]string{
			antique.Name,
			antique.Version,
			antique.Architecture,
			antique.Manager,
			antique.Service,
		}, "-")))

		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, checksum)
		hash = hex.EncodeToString(tmp)

	}

	return hash

}
