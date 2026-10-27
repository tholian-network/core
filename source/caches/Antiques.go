package caches

import "tholian-core/matchers"
import "tholian-core/structs"
import "encoding/json"
import "sort"
import "strings"
import "sync"

func toAntiqueIdentifier(value structs.Antique) string {
	return value.Manager.String() + ":" + value.Vendor + ":" + value.Name + ":" + value.Version.String() + ":" + value.Architecture.String()
}

type Antiques struct {
	Antiques    map[string]*structs.Antique    `json:"antiques"`
	mapcpe      map[string]map[string][]string `json:"-"`
	mapservices map[string][]string            `json:"-"`
	mutex       *sync.RWMutex                  `json:"-"`
}

func NewAntiques() *Antiques {

	var cache Antiques

	cache.Antiques = make(map[string]*structs.Antique)
	cache.mapcpe = make(map[string]map[string][]string, 0)
	cache.mapservices = make(map[string][]string)
	cache.mutex = &sync.RWMutex{}

	return &cache

}

func (cache *Antiques) Add(value structs.Antique) bool {

	cache.mutex.Lock()

	var result bool

	if value.Vendor != "" && value.Name != "" && value.Version.String() != "" {

		identifier := toAntiqueIdentifier(value)

		cache.Antiques[identifier] = &value
		result = true

		_, ok1 := cache.mapcpe[value.Vendor]

		if ok1 == false {
			cache.mapcpe[value.Vendor] = make(map[string][]string, 0)
		}

		_, ok2 := cache.mapcpe[value.Vendor][value.Name]

		if ok2 == false {
			cache.mapcpe[value.Vendor][value.Name] = make([]string, 0)
		}

		cache.mapcpe[value.Vendor][value.Name] = append(cache.mapcpe[value.Vendor][value.Name], identifier)

		_, ok4 := cache.mapservices[value.Service]

		if ok4 == false {
			cache.mapservices[value.Service] = make([]string, 0)
		}

		cache.mapservices[value.Service] = append(cache.mapservices[value.Service], identifier)

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Antiques) Get(identifier string) *structs.Antique {

	cache.mutex.RLock()

	var result *structs.Antique = nil

	if strings.Contains(identifier, ":") {

		tmp := strings.Split(identifier, ":")

		if len(tmp) == 5 || len(tmp) == 6 {

			antique, ok := cache.Antiques[identifier]

			if ok == true {
				result = antique
			}

		}

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Antiques) Length() int {
	return len(cache.Antiques)
}

func (cache *Antiques) Query(query matchers.Antique) []*structs.Antique {

	cache.mutex.RLock()

	result := make([]*structs.Antique, 0)
	found := make(map[string]*structs.Antique)

	if query.Service != "any" {

		// Cache Optimizations for specified query.Service

		identifiers, ok1 := cache.mapservices[query.Service]

		if ok1 == true {

			for _, identifier := range identifiers {

				antique, ok := cache.Antiques[identifier]

				if ok == true {

					matches_name := query.MatchesName(antique.Name)
					matches_version := query.MatchesVersion(antique.Version.String())
					matches_architecture := query.MatchesArchitecture(antique.Architecture.String())
					matches_manager := query.MatchesManager(antique.Manager.String())
					matches_vendor := query.MatchesVendor(antique.Vendor)

					if matches_name && matches_version && matches_architecture && matches_manager && matches_vendor {
						found[identifier] = antique
					}

				}

			}

		}

	} else if query.Vendor != "any" && query.Name != "any" {

		// Cache Optimizations for specified query.Vendor, query.Name

		_, ok1 := cache.mapcpe[query.Vendor]

		if ok1 == true {

			identifiers, ok2 := cache.mapcpe[query.Vendor][query.Name]

			if ok2 == true {

				for _, identifier := range identifiers {

					antique, ok := cache.Antiques[identifier]

					if ok == true {

						matches_version := query.MatchesVersion(antique.Version.String())
						matches_architecture := query.MatchesArchitecture(antique.Architecture.String())
						matches_manager := query.MatchesManager(antique.Manager.String())
						matches_service := query.MatchesService(antique.Service)

						if matches_version && matches_architecture && matches_manager && matches_service {
							found[identifier] = antique
						}

					}

				}

			}

		}

	} else if query.Vendor != "any" && query.Name == "any" {

		// Cache Optimizations for specified query.Vendor

		_, ok1 := cache.mapcpe[query.Vendor]

		if ok1 == true {

			for _, identifiers := range cache.mapcpe[query.Vendor] {

				for _, identifier := range identifiers {

					antique, ok := cache.Antiques[identifier]

					if ok == true {

						matches_version := query.MatchesVersion(antique.Version.String())
						matches_architecture := query.MatchesArchitecture(antique.Architecture.String())
						matches_manager := query.MatchesManager(antique.Manager.String())
						matches_service := query.MatchesService(antique.Service)

						if matches_version && matches_architecture && matches_manager && matches_service {
							found[identifier] = antique
						}

					}

				}

			}

		}

	} else {

		for identifier, antique := range cache.Antiques {

			matches_name := query.MatchesName(antique.Name)
			matches_version := query.MatchesVersion(antique.Version.String())
			matches_architecture := query.MatchesArchitecture(antique.Architecture.String())
			matches_manager := query.MatchesManager(antique.Manager.String())
			matches_vendor := query.MatchesVendor(antique.Vendor)
			matches_service := query.MatchesService(antique.Service)

			if matches_name && matches_version && matches_architecture && matches_manager && matches_vendor && matches_service {
				found[identifier] = antique
			}

		}

	}

	if len(found) > 0 {

		for _, antique := range found {
			result = append(result, antique)
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

func (cache *Antiques) Remove(identifier string) bool {

	cache.mutex.Lock()

	var result bool

	if identifier != "" {

		antique, ok := cache.Antiques[identifier]

		if ok == true {

			delete(cache.Antiques, identifier)
			result = true

			others, ok1 := cache.mapcpe[antique.Vendor][antique.Name]

			if ok1 == true {

				filtered := make([]string, 0)

				for _, other := range others {

					if other != identifier {
						filtered = append(filtered, identifier)
					}

				}

				if len(filtered) > 0 {
					cache.mapcpe[antique.Vendor][antique.Name] = filtered
				} else {
					delete(cache.mapcpe[antique.Vendor], antique.Name)
				}

			}

			others, ok2 := cache.mapservices[antique.Service]

			if ok2 == true {

				filtered := make([]string, 0)

				for _, other := range others {

					if other != identifier {
						filtered = append(filtered, identifier)
					}

				}

				if len(filtered) > 0 {
					cache.mapservices[antique.Service] = filtered
				} else {
					delete(cache.mapservices, antique.Service)
				}

			}

		}

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Antiques) MarshalJSON() ([]byte, error) {

	tmp := make(map[string]*structs.Antique)
	identifiers := make([]string, 0)

	for identifier, _ := range cache.Antiques {
		identifiers = append(identifiers, identifier)
	}

	sort.Strings(identifiers)

	for _, identifier := range identifiers {
		tmp[identifier] = cache.Antiques[identifier]
	}

	return json.MarshalIndent(&struct {
		Antiques map[string]*structs.Antique `json:"antiques"`
	}{
		Antiques: tmp,
	}, "", "\t")

}

func (cache *Antiques) UnmarshalJSON(data []byte) error {

	var tmp struct {
		Antiques map[string]*structs.Antique `json:"antiques"`
	}

	err := json.Unmarshal(data, &tmp)

	if err != nil {
		return err
	}

	for _, antique := range tmp.Antiques {
		cache.Add(*antique)
	}

	return nil

}
