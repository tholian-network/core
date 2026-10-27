package caches

import "tholian-core/matchers"
import "tholian-core/structs"
import "encoding/json"
import "sort"
import "sync"

type Weaknesses struct {
	Weaknesses map[string]*structs.Weakness `json:"weaknesses"`
	mapimpacts map[string][]string          `json:"-"`
	mapscopes  map[string][]string          `json:"-"`
	mutex      *sync.RWMutex                `json:"-"`
}

func NewWeaknesses() *Weaknesses {

	var cache Weaknesses

	cache.Weaknesses = make(map[string]*structs.Weakness)
	cache.mapimpacts = make(map[string][]string, 0)
	cache.mapscopes = make(map[string][]string, 0)
	cache.mutex = &sync.RWMutex{}

	return &cache

}

func (cache *Weaknesses) Add(value structs.Weakness) bool {

	cache.mutex.Lock()

	var result bool

	if value.Name != "" {

		cache.Weaknesses[value.Name] = &value
		result = true

		for _, impact := range value.Impacts {

			_, ok1 := cache.mapimpacts[impact]

			if ok1 == false {
				cache.mapimpacts[impact] = make([]string, 0)
			}

			cache.mapimpacts[impact] = append(cache.mapimpacts[impact], value.Name)

		}

		for _, scope := range value.Scopes {

			_, ok1 := cache.mapscopes[scope]

			if ok1 == false {
				cache.mapscopes[scope] = make([]string, 0)
			}

			cache.mapscopes[scope] = append(cache.mapscopes[scope], value.Name)

		}

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Weaknesses) Get(name string) *structs.Weakness {

	cache.mutex.RLock()

	var result *structs.Weakness = nil

	if name != "" {

		weakness, ok := cache.Weaknesses[name]

		if ok == true {
			result = weakness
		}

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Weaknesses) Length() int {
	return len(cache.Weaknesses)
}

func (cache *Weaknesses) Query(query matchers.Weakness) []*structs.Weakness {

	cache.mutex.RLock()

	result := make([]*structs.Weakness, 0)
	found := make(map[string]*structs.Weakness)

	if query.Name != "any" {

		// Cache Optimizations for specified query.Name

		weakness, ok := cache.Weaknesses[query.Name]

		if ok == true {

			matches_impact := false
			matches_scope := false

			if query.Impact == "any" {

				matches_impact = true

			} else {

				for _, impact := range weakness.Impacts {

					if query.MatchesImpact(impact) {
						matches_impact = true
						break
					}

				}

			}

			if query.Scope == "any" {

				matches_scope = true

			} else {

				for _, scope := range weakness.Scopes {

					if query.MatchesScope(scope) {
						matches_scope = true
						break
					}

				}

			}

			if matches_impact && matches_scope {
				found[weakness.Name] = weakness
			}

		}

	} else if query.Name == "any" && query.Impact != "any" {

		// Cache Optimizations for specified query.Impact

		names, ok1 := cache.mapimpacts[query.Impact]

		if ok1 == true {

			for _, name := range names {

				weakness, ok := cache.Weaknesses[name]

				if ok == true {

					matches_scope := false

					if query.Scope == "any" {

						matches_scope = true

					} else {

						for _, scope := range weakness.Scopes {

							if query.MatchesScope(scope) {
								matches_scope = true
								break
							}

						}

					}

					if matches_scope {
						found[weakness.Name] = weakness
					}

				}

			}

		}

	} else if query.Name == "any" && query.Scope != "any" {

		// Cache Optimizations for query.Scope

		names, ok1 := cache.mapscopes[query.Scope]

		if ok1 == true {

			for _, name := range names {

				weakness, ok := cache.Weaknesses[name]

				if ok == true {

					matches_impact := false

					if query.Impact == "any" {

						matches_impact = true

					} else {

						for _, impact := range weakness.Impacts {

							if query.MatchesImpact(impact) {
								matches_impact = true
								break
							}

						}

					}

					if matches_impact {
						found[weakness.Name] = weakness
					}

				}

			}

		}

	} else {

		for _, weakness := range cache.Weaknesses {

			matches_name := query.MatchesName(weakness.Name)
			matches_impact := false
			matches_scope := false

			if query.Impact == "any" {

				matches_impact = true

			} else {

				for _, impact := range weakness.Impacts {

					if query.MatchesImpact(impact) {
						matches_impact = true
						break
					}

				}

			}

			if query.Scope == "any" {

				matches_scope = true

			} else {

				for _, scope := range weakness.Scopes {

					if query.MatchesScope(scope) {
						matches_scope = true
						break
					}

				}

			}

			if matches_name && matches_impact && matches_scope {
				found[weakness.Name] = weakness
			}

		}

	}

	if len(found) > 0 {

		for _, weakness := range found {
			result = append(result, weakness)
		}

	}

	if len(result) > 1 {

		sort.Slice(result, func(a int, b int) bool {
			return result[a].Name < result[b].Name
		})

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Weaknesses) Remove(name string) bool {

	cache.mutex.Lock()

	var result bool

	if name != "" {

		weakness, ok := cache.Weaknesses[name]

		if ok == true {

			delete(cache.Weaknesses, name)
			result = true

			for _, impact := range weakness.Impacts {

				others, ok1 := cache.mapimpacts[impact]

				if ok1 == true {

					filtered := make([]string, 0)

					for _, other := range others {

						if other != name {
							filtered = append(filtered, other)
						}

					}

					if len(filtered) > 0 {
						cache.mapimpacts[impact] = filtered
					} else {
						delete(cache.mapimpacts, impact)
					}

				}

			}

			for _, scope := range weakness.Scopes {

				others, ok1 := cache.mapscopes[scope]

				if ok1 == true {

					filtered := make([]string, 0)

					for _, other := range others {

						if other != name {
							filtered = append(filtered, other)
						}

					}

					if len(filtered) > 0 {
						cache.mapscopes[scope] = filtered
					} else {
						delete(cache.mapscopes, scope)
					}

				}

			}

		}

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Weaknesses) MarshalJSON() ([]byte, error) {

	tmp := make(map[string]*structs.Weakness)
	names := make([]string, 0)

	for name, _ := range cache.Weaknesses {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {
		tmp[name] = cache.Weaknesses[name]
	}

	return json.MarshalIndent(&struct {
		Weaknesses map[string]*structs.Weakness `json:"weaknesses"`
	}{
		Weaknesses: tmp,
	}, "", "\t")

}

func (cache *Weaknesses) UnmarshalJSON(data []byte) error {

	var tmp struct {
		Weaknesses map[string]*structs.Weakness `json:"weaknesses"`
	}

	err := json.Unmarshal(data, &tmp)

	if err != nil {
		return err
	}

	for _, weakness := range tmp.Weaknesses {
		cache.Add(*weakness)
	}

	return nil

}

