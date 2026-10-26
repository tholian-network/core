package structs

import "tholian-core/matchers"
import "tholian-core/types"
import utils_path "tholian-core/utils/path"
import "net/url"
import "strings"

type Package struct {
	Name         string                `json:"name"`
	Version      types.Version         `json:"version"`
	Architecture types.Architecture    `json:"architecture"`
	Manager      types.Manager         `json:"manager"`
	Vendor       string                `json:"vendor"`
	URL          string                `json:"url"`
	Datetime     types.Datetime        `json:"datetime"`
	Maintainers  []types.Maintainer    `json:"maintainers"`
	Filesystem   []string              `json:"filesystem"`
	Conflicts    []matchers.Package    `json:"conflicts"`
	Dependencies []matchers.Package    `json:"dependencies"`
	Provides     []matchers.Package    `json:"provides"`
	Replaces     []matchers.Package    `json:"replaces"`
	Unresolved   []matchers.Unresolved `json:"unresolved"`
}

func NewPackage(manager string) Package {

	var pkg Package

	pkg.SetManager(manager)

	pkg.Maintainers = make([]types.Maintainer, 0)
	pkg.Filesystem = make([]string, 0)

	pkg.Conflicts = make([]matchers.Package, 0)
	pkg.Dependencies = make([]matchers.Package, 0)
	pkg.Provides = make([]matchers.Package, 0)
	pkg.Replaces = make([]matchers.Package, 0)
	pkg.Unresolved = make([]matchers.Unresolved, 0)

	return pkg

}

func (pkg *Package) Equal(value Package) bool {

	if pkg.Name == value.Name &&
		pkg.Version.String() == value.Version.String() &&
		pkg.Architecture.String() == value.Architecture.String() &&
		pkg.Manager.String() == value.Manager.String() {
		return true
	}

	return false

}

func (pkg *Package) IsValid() bool {

	if pkg.Name != "" {

		result := true

		if pkg.Datetime.IsValid() == false {
			result = false
		}

		if pkg.Version.IsValid() == false {
			result = false
		}

		if pkg.Architecture.IsValid() == false {
			result = false
		}

		if pkg.Manager.IsValid() == false {
			result = false
		}

		if result == true {

			for m := 0; m < len(pkg.Maintainers); m++ {

				if pkg.Maintainers[m].IsValid() == false {
					result = false
					break
				}

			}

		}

		if result == true {

			for c := 0; c < len(pkg.Conflicts); c++ {

				if pkg.Conflicts[c].IsValid() == false {
					result = false
					break
				}

			}

		}

		if result == true {

			for d := 0; d < len(pkg.Dependencies); d++ {

				if pkg.Dependencies[d].IsValid() == false {
					result = false
					break
				}

			}

		}

		if result == true {

			for p := 0; p < len(pkg.Provides); p++ {

				if pkg.Provides[p].IsValid() == false {
					result = false
					break
				}

			}

		}

		if result == true {

			for r := 0; r < len(pkg.Replaces); r++ {

				if pkg.Replaces[r].IsValid() == false {
					result = false
					break
				}

			}

		}

		return result

	}

	return false

}

func (pkg *Package) SetArchitecture(value string) {

	architecture := types.ParseArchitecture(value)

	if architecture != nil {
		pkg.Architecture = *architecture
	}

}

func (pkg *Package) AddConflict(value matchers.Package) {

	if value.IsValid() {

		found := false

		for c := 0; c < len(pkg.Conflicts); c++ {

			if pkg.Conflicts[c].Equal(value) {
				found = true
				break
			}

		}

		if found == false {
			pkg.Conflicts = append(pkg.Conflicts, value)
		}

	}

}

func (pkg *Package) RemoveConflict(value matchers.Package) {

	index := -1

	for c := 0; c < len(pkg.Conflicts); c++ {

		if pkg.Conflicts[c].Equal(value) {
			index = c
			break
		}

	}

	if index != -1 {
		pkg.Conflicts = append(pkg.Conflicts[:index], pkg.Conflicts[index+1:]...)
	}

}

func (pkg *Package) SetConflicts(value []matchers.Package) {

	filtered := make([]matchers.Package, 0)

	for v := 0; v < len(value); v++ {

		if value[v].IsValid() {
			filtered = append(filtered, value[v])
		}

	}

	pkg.Conflicts = filtered

}

func (pkg *Package) SetDatetime(value string) {

	datetime := types.ToDatetime(value)

	if datetime.IsValid() {
		pkg.Datetime = datetime
	}

}

func (pkg *Package) AddDependency(value matchers.Package) {

	if value.IsValid() {

		found := false

		for d := 0; d < len(pkg.Dependencies); d++ {

			if pkg.Dependencies[d].Equal(value) {
				found = true
				break
			}

		}

		if found == false {
			pkg.Dependencies = append(pkg.Dependencies, value)
		}

	}

}

func (pkg *Package) HasDependency(dependency matchers.Package) bool {

	var result bool

	for d := 0; d < len(pkg.Dependencies); d++ {

		other := pkg.Dependencies[d]

		if dependency.Name == other.Name &&
			dependency.Version == other.Version {
			result = true
			break
		}

	}

	return result

}

func (pkg *Package) RemoveDependency(value matchers.Package) {

	index := -1

	for d := 0; d < len(pkg.Dependencies); d++ {

		if pkg.Dependencies[d].Equal(value) {
			index = d
			break
		}

	}

	if index != -1 {
		pkg.Dependencies = append(pkg.Dependencies[:index], pkg.Dependencies[index+1:]...)
	}

}

func (pkg *Package) ResolveDependencies(packages []Package) {

	if len(pkg.Unresolved) > 0 && len(packages) > 0 {

		remaining := make([]matchers.Unresolved, 0)

		for u := 0; u < len(pkg.Unresolved); u++ {

			unresolved := pkg.Unresolved[u]

			var resolved matchers.Package

			for p := 0; p < len(packages); p++ {

				other := packages[p]

				if unresolved.Matches(other.Name, other.Version.String(), other.Architecture.String(), other.Manager.String(), other.Vendor) {

					resolved.Name = other.Name
					resolved.Version = other.Version.String()

					if other.Architecture != "" {
						resolved.Architecture = other.Architecture.String()
					} else {
						resolved.Architecture = "any"
					}

					if other.Manager != "" {
						resolved.Manager = other.Manager.String()
					} else {
						resolved.Manager = "any"
					}

					if other.Vendor != "" {
						resolved.Vendor = other.Vendor
					} else {
						resolved.Vendor = "any"
					}

				} else {

					for p := 0; p < len(other.Provides); p++ {

						provide := other.Provides[p]

						if unresolved.Matches(provide.Name, provide.Version, provide.Architecture, provide.Manager, provide.Vendor) {

							resolved.Name = other.Name
							resolved.Version = other.Version.String()

							if other.Architecture != "" {
								resolved.Architecture = other.Architecture.String()
							} else {
								resolved.Architecture = "any"
							}

							if other.Manager != "" {
								resolved.Manager = other.Manager.String()
							} else {
								resolved.Manager = "any"
							}

							if other.Vendor != "" {
								resolved.Vendor = other.Vendor
							} else {
								resolved.Vendor = "any"
							}

						}

					}

				}

				if resolved.Name != "" {
					break
				}

			}

			if resolved.Name != "" {

				if !pkg.HasDependency(resolved) {
					pkg.AddDependency(resolved)
				}

			} else {

				remaining = append(remaining, unresolved)

			}

		}

		pkg.Unresolved = remaining

	}

}

func (pkg *Package) SetDependencies(value []matchers.Package) {

	filtered := make([]matchers.Package, 0)

	for v := 0; v < len(value); v++ {

		if value[v].IsValid() {
			filtered = append(filtered, value[v])
		}

	}

	pkg.Dependencies = filtered

}

func (pkg *Package) HasFilesystem(value string) bool {

	var result bool

	for f := 0; f < len(pkg.Filesystem); f++ {

		if pkg.Filesystem[f] == value {
			result = true
			break
		}

	}

	return result

}

func (pkg *Package) AddFilesystem(value string) {

	if utils_path.IsWatchedFile(value) {

		found := false

		for f := 0; f < len(pkg.Filesystem); f++ {

			file := pkg.Filesystem[f]

			if file == value {
				found = true
				break
			}

		}

		if found == false {
			pkg.Filesystem = append(pkg.Filesystem, value)
		}

	}

}

func (pkg *Package) RemoveFilesystem(value string) {

	index := -1

	for f := 0; f < len(pkg.Filesystem); f++ {

		if pkg.Filesystem[f] == value {
			index = f
			break
		}

	}

	if index != -1 {
		pkg.Filesystem = append(pkg.Filesystem[:index], pkg.Filesystem[index+1:]...)
	}

}

func (pkg *Package) SetFilesystem(value []string) {

	filtered := make([]string, 0)

	for v := 0; v < len(value); v++ {

		file := value[v]

		if utils_path.IsWatchedFile(file) == true {

			found := false

			for f := 0; f < len(filtered); f++ {

				if filtered[f] == file {
					found = true
					break
				}

			}

			if found == false {
				filtered = append(filtered, file)
			}

		}

	}

	pkg.Filesystem = filtered

}

func (pkg *Package) AddMaintainer(value types.Maintainer) {

	if value.IsValid() {

		found := false

		for m := 0; m < len(pkg.Maintainers); m++ {

			if pkg.Maintainers[m].Equal(value) {
				found = true
				break
			}

		}

		if found == false {
			pkg.Maintainers = append(pkg.Maintainers, value)
		}

	}

}

func (pkg *Package) RemoveMaintainer(value types.Maintainer) {

	index := -1

	for m := 0; m < len(pkg.Maintainers); m++ {

		if pkg.Maintainers[m].Equal(value) {
			index = m
			break
		}

	}

	if index != -1 {
		pkg.Maintainers = append(pkg.Maintainers[:index], pkg.Maintainers[index+1:]...)
	}

}

func (pkg *Package) SetMaintainers(value []types.Maintainer) {

	filtered := make([]types.Maintainer, 0)

	for v := 0; v < len(value); v++ {

		if value[v].IsValid() {
			filtered = append(filtered, value[v])
		}

	}

	pkg.Maintainers = filtered

}

func (pkg *Package) SetManager(value string) {

	manager := types.ParseManager(value)

	if manager != nil {
		pkg.Manager = *manager
	}

}

func (pkg *Package) SetName(value string) {
	pkg.Name = strings.TrimSpace(value)
}

func (pkg *Package) AddProvide(value matchers.Package) {

	if value.IsValid() {

		found := false

		for p := 0; p < len(pkg.Provides); p++ {

			if pkg.Provides[p].Equal(value) {
				found = true
				break
			}

		}

		if found == false {
			pkg.Provides = append(pkg.Provides, value)
		}

	}

}

func (pkg *Package) RemoveProvide(value matchers.Package) {

	index := -1

	for p := 0; p < len(pkg.Provides); p++ {

		if pkg.Provides[p].Equal(value) {
			index = p
			break
		}

	}

	if index != -1 {
		pkg.Provides = append(pkg.Provides[:index], pkg.Provides[index+1:]...)
	}

}

func (pkg *Package) SetProvides(value []matchers.Package) {

	filtered := make([]matchers.Package, 0)

	for v := 0; v < len(value); v++ {

		if value[v].IsValid() {
			filtered = append(filtered, value[v])
		}

	}

	pkg.Provides = filtered

}

func (pkg *Package) AddReplace(value matchers.Package) {

	if value.IsValid() {

		found := false

		for r := 0; r < len(pkg.Replaces); r++ {

			if pkg.Replaces[r].Equal(value) {
				found = true
				break
			}

		}

		if found == false {
			pkg.Replaces = append(pkg.Replaces, value)
		}

	}

}

func (pkg *Package) RemoveReplace(value matchers.Package) {

	index := -1

	for r := 0; r < len(pkg.Replaces); r++ {

		if pkg.Replaces[r].Equal(value) {
			index = r
			break
		}

	}

	if index != -1 {
		pkg.Replaces = append(pkg.Replaces[:index], pkg.Replaces[index+1:]...)
	}

}

func (pkg *Package) SetReplaces(value []matchers.Package) {

	filtered := make([]matchers.Package, 0)

	for v := 0; v < len(value); v++ {

		if value[v].IsValid() {
			filtered = append(filtered, value[v])
		}

	}

	pkg.Replaces = filtered

}

func (pkg *Package) SetURL(value string) {

	tmp, err := url.ParseRequestURI(value)

	if err == nil {

		if tmp.Scheme == "https" || tmp.Scheme == "http" {
			pkg.URL = tmp.String()
		}

	}

}

func (pkg *Package) SetVendor(value string) {
	pkg.Vendor = strings.TrimSpace(value)
}

func (pkg *Package) SetVersion(value string) {

	version := types.ToVersion(value)

	if version.IsValid() {
		pkg.Version = version
	}

}
