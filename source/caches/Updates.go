package caches

import "tholian-core/matchers"
import "tholian-core/structs"
import "encoding/json"
import "sort"
import "strings"
import "sync"

func toUpdateIdentifier(value structs.Update) string {
	return value.Manager.String() + ":" + value.Vendor + ":" + value.Name + ":" + value.Version.String() + ":" + value.Architecture.String()
}

type Updates struct {
	Updates map[string]*structs.Update     `json:"updates"`
	mapcpe  map[string]map[string][]string `json:"-"`
	mutex   *sync.RWMutex                  `json:"-"`
}

func NewUpdates() *Updates {

	var cache Updates

	cache.Updates = make(map[string]*structs.Update)
	cache.mapcpe = make(map[string]map[string][]string, 0)
	cache.mutex = &sync.RWMutex{}

	return &cache

}

func (cache *Updates) Add(value structs.Update) bool {

	cache.mutex.Lock()

	var result bool

	if value.Vendor != "" && value.Name != "" && value.Version.String() != "" {

		identifier := toUpdateIdentifier(value)

		cache.Updates[identifier] = &value
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

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Updates) Get(identifier string) *structs.Update {

	cache.mutex.RLock()

	var result *structs.Update = nil

	if strings.Contains(identifier, ":") {

		tmp := strings.Split(identifier, ":")

		if len(tmp) == 5 || len(tmp) == 6 {

			update, ok := cache.Updates[identifier]

			if ok == true {
				result = update
			}

		}

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Updates) Length() int {
	return len(cache.Updates)
}

func (cache *Updates) Query(query matchers.Update) []*structs.Update {

	cache.mutex.RLock()

	result := make([]*structs.Update, 0)
	found := make(map[string]*structs.Update)

	if query.Vendor != "any" && query.Name != "any" {

		// Cache Optimizations for specified query.Vendor, query.Name

		_, ok1 := cache.mapcpe[query.Vendor]

		if ok1 == true {

			identifiers, ok2 := cache.mapcpe[query.Vendor][query.Name]

			if ok2 == true {

				for _, identifier := range identifiers {

					update, ok := cache.Updates[identifier]

					if ok == true {

						matches_version := query.MatchesVersion(update.Version.String())
						matches_architecture := query.MatchesArchitecture(update.Architecture.String())
						matches_manager := query.MatchesManager(update.Manager.String())

						if matches_version && matches_architecture && matches_manager {
							found[identifier] = update
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

					update, ok := cache.Updates[identifier]

					if ok == true {

						matches_version := query.MatchesVersion(update.Version.String())
						matches_architecture := query.MatchesArchitecture(update.Architecture.String())
						matches_manager := query.MatchesManager(update.Manager.String())

						if matches_version && matches_architecture && matches_manager {
							found[identifier] = update
						}

					}

				}

			}

		}

	} else {

		for identifier, update := range cache.Updates {

			matches_name := query.MatchesName(update.Name)
			matches_version := query.MatchesVersion(update.Version.String())
			matches_architecture := query.MatchesArchitecture(update.Architecture.String())
			matches_manager := query.MatchesManager(update.Manager.String())
			matches_vendor := query.MatchesVendor(update.Vendor)

			if matches_name && matches_version && matches_architecture && matches_manager && matches_vendor {
				found[identifier] = update
			}

		}

	}

	if len(found) > 0 {

		for _, update := range found {
			result = append(result, update)
		}

	}

	if len(result) > 1 {

		sort.Slice(result, func(a int, b int) bool {

			if result[a].Vendor == result[b].Vendor {

				if result[a].Name == result[b].Name {

					if result[a].Version.Equal(result[b].Version) {
						return result[a].Architecture.String() < result[b].Architecture.String()
					} else {
						return result[a].Version.IsAfter(result[b].Version)
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

func (cache *Updates) Remove(identifier string) bool {

	cache.mutex.Lock()

	var result bool

	if identifier != "" {

		update, ok := cache.Updates[identifier]

		if ok == true {

			delete(cache.Updates, identifier)
			result = true

			others, ok1 := cache.mapcpe[update.Vendor][update.Name]

			if ok1 == true {

				filtered := make([]string, 0)

				for _, other := range others {

					if other != identifier {
						filtered = append(filtered, identifier)
					}

				}

				if len(filtered) > 0 {
					cache.mapcpe[update.Vendor][update.Name] = filtered
				} else {
					delete(cache.mapcpe[update.Vendor], update.Name)
				}

			}

		}

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Updates) MarshalJSON() ([]byte, error) {

	tmp := make(map[string]*structs.Update)
	identifiers := make([]string, 0)

	for identifier, _ := range cache.Updates {
		identifiers = append(identifiers, identifier)
	}

	sort.Strings(identifiers)

	for _, identifier := range identifiers {
		tmp[identifier] = cache.Updates[identifier]
	}

	return json.MarshalIndent(&struct {
		Updates map[string]*structs.Update `json:"updates"`
	}{
		Updates: tmp,
	}, "", "\t")

}

func (cache *Updates) UnmarshalJSON(data []byte) error {

	var tmp struct {
		Updates map[string]*structs.Update `json:"updates"`
	}

	err := json.Unmarshal(data, &tmp)

	if err != nil {
		return err
	}

	for _, update := range tmp.Updates {
		cache.Add(*update)
	}

	return nil

}
