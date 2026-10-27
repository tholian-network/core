package matchers

import "tholian-core/types"
import "strings"

type Update struct {
	Name         string `json:"name"`
	Version      string `json:"version"`
	Architecture string `json:"architecture"`
	Manager      string `json:"manager"`
	Vendor       string `json:"vendor"`
}

func NewUpdate() Update {

	var update Update

	update.Name = Any
	update.Version = Any
	update.Architecture = Any
	update.Manager = Any
	update.Vendor = Any

	return update

}

func ToUpdate(value string) Update {

	var update Update

	update.Name = Any
	update.Version = Any
	update.Architecture = Any
	update.Manager = Any
	update.Vendor = Any

	update.Parse(value)

	return update

}

func (update *Update) Equal(value Update) bool {

	if update.Name == value.Name &&
		update.Version == value.Version &&
		update.Architecture == value.Architecture &&
		update.Manager == value.Manager &&
		update.Vendor == value.Vendor {
		return true
	}

	return false

}

func (update *Update) IsValid() bool {

	if update.Name != Any ||
		update.Version != Any ||
		update.Architecture != Any ||
		update.Manager != Any ||
		update.Vendor != Any {
		return true
	}

	return false

}

func (update *Update) Matches(name string, version string, architecture string, manager string, vendor string) bool {
	return update.MatchesName(name) && update.MatchesVersion(version) && update.MatchesArchitecture(architecture) && update.MatchesManager(manager) && update.MatchesVendor(vendor)
}

func (update *Update) MatchesArchitecture(value string) bool {

	if update.Architecture == value {
		return true
	} else if update.Architecture == Any {
		return true
	}

	return false

}

func (update *Update) MatchesManager(value string) bool {

	if update.Manager == value {
		return true
	} else if update.Manager == Any {
		return true
	}

	return false

}

func (update *Update) MatchesName(value string) bool {

	if update.Name == value {
		return true
	} else if update.Name == Any {
		return true
	}

	return false

}

func (update *Update) MatchesVendor(value string) bool {

	if update.Vendor == value {
		return true
	} else if update.Vendor == Any {
		return true
	}

	return false

}

func (update *Update) MatchesVersion(value string) bool {

	// Compatibility with "<operator> <version>" syntax
	if strings.Contains(value, " ") {
		value = strings.TrimSpace(value[strings.Index(value, " ")+1:])
	}

	if update.Version == Any {

		return true

	} else if strings.HasPrefix(update.Version, "<= ") {

		update_version := types.ToVersion(update.Version[3:])
		other_version := types.ToVersion(value)

		if other_version.Equal(update_version) {
			return true
		} else if other_version.IsBefore(update_version) {
			return true
		}

	} else if strings.HasPrefix(update.Version, "< ") {

		update_version := types.ToVersion(update.Version[2:])
		other_version := types.ToVersion(value)

		if other_version.IsBefore(update_version) {
			return true
		}

	} else if strings.HasPrefix(update.Version, ">= ") {

		update_version := types.ToVersion(update.Version[3:])
		other_version := types.ToVersion(value)

		if other_version.Equal(update_version) {
			return true
		} else if other_version.IsAfter(update_version) {
			return true
		}

	} else if strings.HasPrefix(update.Version, "> ") {

		update_version := types.ToVersion(update.Version[2:])
		other_version := types.ToVersion(value)

		if other_version.IsAfter(update_version) {
			return true
		}

	} else if strings.HasPrefix(update.Version, "= ") {

		update_version := types.ToVersion(update.Version[2:])
		other_version := types.ToVersion(value)

		if other_version.Equal(update_version) {
			return true
		}

	} else {

		update_version := types.ToVersion(update.Version)
		other_version := types.ToVersion(value)

		if other_version.Equal(update_version) {
			return true
		}

	}

	return false

}

func (update *Update) Parse(value string) {

	name, version, architecture := parseVersionCondition(value)

	update.Name = name
	update.Version = version

	if architecture != "" {
		update.Architecture = architecture
	}

}

func (update *Update) SetArchitecture(value string) {

	architecture := types.ParseArchitecture(value)

	if architecture != nil {
		update.Architecture = architecture.String()
	}

}

func (update *Update) SetManager(value string) {

	manager := types.ParseManager(value)

	if manager != nil {
		update.Manager = manager.String()
	}

}

func (update *Update) SetName(value string) {

	if value == "all" || value == "any" || value == "*" {
		update.Name = Any
	} else if value != "" {
		update.Name = strings.TrimSpace(value)
	}

}

func (update *Update) SetVersion(value string) {

	if value == "all" || value == "any" || value == "*" {
		update.Version = Any
	} else if value != "" {
		update.Version = value
	}

}

func (update *Update) Hash() string {

	hash := ""

	if update.Name != "" {

		hash = toHash(
			update.Name,
			update.Version,
			update.Architecture,
			update.Manager,
		)

	}

	return hash

}
