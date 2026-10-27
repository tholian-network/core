package caches

import "tholian-core/matchers"
import "tholian-core/structs"
import "encoding/json"
import "sort"
import "strings"
import "sync"

func toPackageIdentifier(value structs.Package) string {
	return value.Manager.String() + ":" + value.Vendor + ":" + value.Name + ":" + value.Version.String() + ":" + value.Architecture.String()
}

type Packages struct {
	Packages        map[string]*structs.Package    `json:"packages"`
	mapdependencies map[string]map[string][]string `json:"-"`
	mappackages     map[string]map[string][]string `json:"-"`
	mutex           *sync.RWMutex                  `json:"-"`
}

func NewPackages() *Packages {

	var cache Packages

	cache.Packages = make(map[string]*structs.Package)
	cache.mapdependencies = make(map[string]map[string][]string, 0)
	cache.mappackages = make(map[string]map[string][]string, 0)
	cache.mutex = &sync.RWMutex{}

	return &cache

}

func (cache *Packages) Add(value structs.Package) bool {

	cache.mutex.Lock()

	var result bool

	if value.Vendor != "" && value.Name != "" && value.Version.String() != "" {

		identifier := toPackageIdentifier(value)

		cache.Packages[identifier] = &value
		result = true

		_, ok1 := cache.mappackages[value.Vendor]

		if ok1 == false {
			cache.mappackages[value.Vendor] = make(map[string][]string, 0)
		}

		_, ok2 := cache.mappackages[value.Vendor][value.Name]

		if ok2 == false {
			cache.mappackages[value.Vendor][value.Name] = make([]string, 0)
		}

		cache.mappackages[value.Vendor][value.Name] = append(cache.mappackages[value.Vendor][value.Name], identifier)

		for _, dependency := range value.Dependencies {

			_, ok3 := cache.mapdependencies[dependency.Vendor]

			if ok3 == false {
				cache.mapdependencies[dependency.Vendor] = make(map[string][]string, 0)
			}

			_, ok4 := cache.mapdependencies[dependency.Vendor][dependency.Name]

			if ok4 == false {
				cache.mapdependencies[dependency.Vendor][dependency.Name] = make([]string, 0)
			}

			cache.mapdependencies[dependency.Vendor][dependency.Name] = append(cache.mapdependencies[dependency.Vendor][dependency.Name], identifier)

		}

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Packages) Get(identifier string) *structs.Package {

	cache.mutex.RLock()

	var result *structs.Package = nil

	if strings.Contains(identifier, ":") {

		tmp := strings.Split(identifier, ":")

		if len(tmp) == 5 || len(tmp) == 6 {

			pkg, ok := cache.Packages[identifier]

			if ok == true {
				result = pkg
			}

		}

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Packages) Length() int {
	return len(cache.Packages)
}

func (cache *Packages) Query(query matchers.Package) []*structs.Package {

	cache.mutex.RLock()

	result := make([]*structs.Package, 0)
	found := make(map[string]*structs.Package)

	if query.Vendor != "any" && query.Name != "any" {

		// Cache Optimizations for specified query.Vendor, query.Name

		_, ok1 := cache.mappackages[query.Vendor]

		if ok1 == true {

			identifiers, ok2 := cache.mappackages[query.Vendor][query.Name]

			if ok2 == true {

				for _, identifier := range identifiers {

					pkg, ok := cache.Packages[identifier]

					if ok == true {

						matches_version := query.MatchesVersion(pkg.Version.String())
						matches_architecture := query.MatchesArchitecture(pkg.Architecture.String())
						matches_manager := query.MatchesManager(pkg.Manager.String())

						if matches_version && matches_architecture && matches_manager {
							found[identifier] = pkg
						}

					}

				}

			}

		}

	} else if query.Vendor != "any" && query.Name == "any" {

		// Cache Optimizations for specified query.Vendor

		_, ok1 := cache.mappackages[query.Vendor]

		if ok1 == true {

			for _, identifiers := range cache.mappackages[query.Vendor] {

				for _, identifier := range identifiers {

					pkg, ok := cache.Packages[identifier]

					if ok == true {

						matches_version := query.MatchesVersion(pkg.Version.String())
						matches_architecture := query.MatchesArchitecture(pkg.Architecture.String())
						matches_manager := query.MatchesManager(pkg.Manager.String())

						if matches_version && matches_architecture && matches_manager {
							found[identifier] = pkg
						}

					}

				}

			}

		}

	} else {

		for identifier, pkg := range cache.Packages {

			matches_name := query.MatchesName(pkg.Name)
			matches_version := query.MatchesVersion(pkg.Version.String())
			matches_architecture := query.MatchesArchitecture(pkg.Architecture.String())
			matches_manager := query.MatchesManager(pkg.Manager.String())
			matches_vendor := query.MatchesVendor(pkg.Vendor)

			if matches_name && matches_version && matches_architecture && matches_manager && matches_vendor {
				found[identifier] = pkg
			}

		}

	}

	if len(found) > 0 {

		for _, pkg := range found {
			result = append(result, pkg)
		}

	}

	if len(result) > 1 {

		sort.Slice(result, func(a int, b int) bool {

			if result[a].Vendor == result[b].Vendor {

				if result[a].Name == result[b].Name {

					if result[a].Version.Equal(result[b].Version) {
						return result[a].Architecture.String() < result[b].Architecture.String()
					} else {
						return result[a].Version.IsBefore(result[b].Version)
					}

				} else {
					return result[a].Name < result[b].Name
				}

			} else {
				return result[a].Vendor < result[b].Vendor
			}

		})

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Packages) QueryByDependency(query matchers.Package) []*structs.Package {

	cache.mutex.RLock()

	result := make([]*structs.Package, 0)
	found := make(map[string]*structs.Package)

	if query.Vendor != "any" && query.Name != "any" {

		// Cache Optimizations for specified query.Vendor, query.Name

		_, ok1 := cache.mapdependencies[query.Vendor]

		if ok1 == true {

			identifiers, ok2 := cache.mapdependencies[query.Vendor][query.Name]

			if ok2 == true {

				for _, identifier := range identifiers {

					pkg, ok := cache.Packages[identifier]

					if ok == true {

						matches_version := query.MatchesVersion(pkg.Version.String())
						matches_architecture := query.MatchesArchitecture(pkg.Architecture.String())
						matches_manager := query.MatchesManager(pkg.Manager.String())

						if matches_version && matches_architecture && matches_manager {
							found[identifier] = pkg
						}

					}

				}

			}

		}

	} else if query.Vendor != "any" && query.Name == "any" {

		// Cache Optimizations for specified query.Vendor

		_, ok1 := cache.mapdependencies[query.Vendor]

		if ok1 == true {

			for _, identifiers := range cache.mapdependencies[query.Vendor] {

				for _, identifier := range identifiers {

					pkg, ok := cache.Packages[identifier]

					if ok == true {

						matches_version := query.MatchesVersion(pkg.Version.String())
						matches_architecture := query.MatchesArchitecture(pkg.Architecture.String())
						matches_manager := query.MatchesManager(pkg.Manager.String())

						if matches_version && matches_architecture && matches_manager {
							found[identifier] = pkg
						}

					}

				}

			}

		}

	} else {

		for identifier, pkg := range cache.Packages {

			matches_dependency := false

			for _, dependency := range pkg.Dependencies {

				matches_name := query.MatchesName(dependency.Name)
				matches_version := query.MatchesVersion(dependency.Version)
				matches_architecture := query.MatchesArchitecture(dependency.Architecture)
				matches_manager := query.MatchesManager(dependency.Manager)
				matches_vendor := query.MatchesVendor(dependency.Vendor)

				if matches_name && matches_version && matches_architecture && matches_manager && matches_vendor {
					matches_dependency = true
					break
				}

			}

			if matches_dependency {
				found[identifier] = pkg
			}

		}

	}

	if len(found) > 0 {

		for _, pkg := range found {
			result = append(result, pkg)
		}

	}

	if len(result) > 1 {

		sort.Slice(result, func(a int, b int) bool {

			if result[a].Vendor == result[b].Vendor {

				if result[a].Name == result[b].Name {

					if result[a].Version.Equal(result[b].Version) {
						return result[a].Architecture.String() < result[b].Architecture.String()
					} else {
						return result[a].Version.IsBefore(result[b].Version)
					}

				} else {
					return result[a].Name < result[b].Name
				}

			} else {
				return result[a].Vendor < result[b].Vendor
			}

		})

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Packages) Remove(identifier string) bool {

	cache.mutex.Lock()

	var result bool

	if identifier != "" {

		pkg, ok := cache.Packages[identifier]

		if ok == true {

			delete(cache.Packages, identifier)
			result = true

			others, ok1 := cache.mappackages[pkg.Vendor][pkg.Name]

			if ok1 == true {

				filtered := make([]string, 0)

				for _, other := range others {

					if other != identifier {
						filtered = append(filtered, identifier)
					}

				}

				if len(filtered) > 0 {
					cache.mappackages[pkg.Vendor][pkg.Name] = filtered
				} else {
					delete(cache.mappackages[pkg.Vendor], pkg.Name)
				}

			}

			for _, dependency := range pkg.Dependencies {

				others, ok2 := cache.mapdependencies[dependency.Vendor][dependency.Name]

				if ok2 == true {

					filtered := make([]string, 0)

					for _, other := range others {

						if other != identifier {
							filtered = append(filtered, identifier)
						}

					}

					if len(filtered) > 0 {
						cache.mapdependencies[dependency.Vendor][dependency.Name] = filtered
					} else {
						delete(cache.mapdependencies[dependency.Vendor], dependency.Name)
					}

				}

			}

		}

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Packages) MarshalJSON() ([]byte, error) {

	tmp := make(map[string]*structs.Package)
	identifiers := make([]string, 0)

	for identifier, _ := range cache.Packages {
		identifiers = append(identifiers, identifier)
	}

	sort.Strings(identifiers)

	for _, identifier := range identifiers {
		tmp[identifier] = cache.Packages[identifier]
	}

	return json.MarshalIndent(&struct {
		Packages map[string]*structs.Package `json:"packages"`
	}{
		Packages: tmp,
	}, "", "\t")

}

func (cache *Packages) UnmarshalJSON(data []byte) error {

	var tmp struct {
		Packages map[string]*structs.Package `json:"packages"`
	}

	err := json.Unmarshal(data, &tmp)

	if err != nil {
		return err
	}

	for _, pkg := range tmp.Packages {
		cache.Add(*pkg)
	}

	return nil

}
