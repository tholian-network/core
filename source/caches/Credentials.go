package caches

import "tholian-core/matchers"
import "tholian-core/structs"
import "encoding/json"
import "slices"
import "sort"
import "sync"

type Credentials struct {
	Credentials map[string]*structs.Credential `json:"credentials"`
	maptypes    map[string][]string            `json:"-"`
	mutex       *sync.RWMutex                  `json:"-"`
}

func NewCredentials() *Credentials {

	var cache Credentials

	cache.Credentials = make(map[string]*structs.Credential)
	cache.maptypes = make(map[string][]string, 0)
	cache.mutex = &sync.RWMutex{}

	return &cache

}

func (cache *Credentials) Add(value structs.Credential) bool {

	cache.mutex.Lock()

	var result bool

	if value.Name != "" && value.Type != "" {

		cache.Credentials[value.Name] = &value
		result = true

		_, ok1 := cache.maptypes[value.Type]

		if ok1 == false {
			cache.maptypes[value.Type] = make([]string, 0)
		}

		cache.maptypes[value.Type] = append(cache.maptypes[value.Type], value.Name)

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Credentials) Get(name string) *structs.Credential {

	cache.mutex.RLock()

	var result *structs.Credential = nil

	if name != "" {

		credential, ok := cache.Credentials[name]

		if ok == true {
			result = credential
		}

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Credentials) Length() int {
	return len(cache.Credentials)
}

func (cache *Credentials) Query(query matchers.Credential) []*structs.Credential {

	cache.mutex.RLock()

	result := make([]*structs.Credential, 0)
	found := make(map[string]*structs.Credential)

	if query.Name != "any" {

		// Cache Optimizations for specified query.Name

		credential, ok := cache.Credentials[query.Name]

		if ok == true {

			matches_password := false
			matches_type := query.MatchesType(credential.Type)

			if query.Password == "any" {

				matches_password = true

			} else {

				for _, password := range credential.Passwords {

					if query.MatchesPassword(password) {
						matches_password = true
						break
					}

				}

			}

			if matches_password && matches_type {
				found[credential.Name] = credential
			}

		}

	} else if query.Type != "any" {

		// Cache Optimizations for specified query.Type

		_, ok1 := cache.maptypes[query.Type]

		if ok1 == true {

			for _, name := range cache.maptypes[query.Type] {

				credential, ok := cache.Credentials[name]

				if ok == true {

					matches_name := query.MatchesName(credential.Name)
					matches_password := false

					if query.Password == "any" {

						matches_password = true

					} else {

						for _, password := range credential.Passwords {

							if query.MatchesPassword(password) {
								matches_password = true
								break
							}

						}

					}

					if matches_name && matches_password {
						found[credential.Name] = credential
					}

				}

			}

		}

	} else {

		for _, credential := range cache.Credentials {

			matches_name := query.MatchesName(credential.Name)
			matches_password := false
			matches_type := query.MatchesType(credential.Type)

			if query.Password == "any" {

				matches_password = true

			} else {

				for _, password := range credential.Passwords {

					if query.MatchesPassword(password) {
						matches_password = true
						break
					}

				}

			}

			if matches_name && matches_password && matches_type {
				found[credential.Name] = credential
			}

		}

	}

	if len(found) > 0 {

		for _, credential := range found {
			result = append(result, credential)
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

func (cache *Credentials) Remove(name string) bool {

	cache.mutex.Lock()

	var result bool

	if name != "" {

		credential, ok1 := cache.Credentials[name]

		if ok1 == true {

			delete(cache.Credentials, name)
			result = true

			others, ok2 := cache.maptypes[credential.Type]

			if ok2 == true {

				filtered := make([]string, 0)

				for _, other := range others {

					if other != name {
						filtered = append(filtered, other)
					}

				}

				if len(filtered) > 0 {
					cache.maptypes[credential.Type] = filtered
				} else {
					delete(cache.maptypes, credential.Type)
				}

			}

		}

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Credentials) MarshalJSON() ([]byte, error) {

	tmp := make(map[string]*structs.Credential)
	names := make([]string, 0)

	for name, _ := range cache.Credentials {
		names = append(names, name)
	}

	slices.Sort(names)

	for _, name := range names {
		tmp[name] = cache.Credentials[name]
	}

	return json.MarshalIndent(&struct {
		Credentials map[string]*structs.Credential `json:"credentials"`
	}{
		Credentials: tmp,
	}, "", "\t")

}

func (cache *Credentials) UnmarshalJSON(data []byte) error {

	var tmp struct {
		Credentials map[string]*structs.Credential `json:"credentials"`
	}

	err := json.Unmarshal(data, &tmp)

	if err != nil {
		return err
	}

	for _, credential := range tmp.Credentials {
		cache.Add(*credential)
	}

	return nil

}
