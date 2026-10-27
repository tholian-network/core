package caches

import "tholian-core/matchers"
import "tholian-core/structs"
import "encoding/json"
import "sort"
import "sync"

type Distributions struct {
	Distributions map[string]*structs.Distribution `json:"distributions"`
	mutex         *sync.RWMutex                    `json:"-"`
}

func NewDistributions() *Distributions {

	var cache Distributions

	cache.Distributions = make(map[string]*structs.Distribution)
	cache.mutex = &sync.RWMutex{}

	return &cache

}

func (cache *Distributions) Add(value structs.Distribution) bool {

	cache.mutex.Lock()

	var result bool

	if value.Name != "" {

		cache.Distributions[value.Name] = &value
		result = true

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Distributions) Get(name string) *structs.Distribution {

	cache.mutex.RLock()

	var result *structs.Distribution = nil

	if name != "" {

		distribution, ok := cache.Distributions[name]

		if ok == true {
			result = distribution
		}

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Distributions) Length() int {
	return len(cache.Distributions)
}

func (cache *Distributions) Query(query matchers.Distribution) []*structs.Distribution {

	cache.mutex.RLock()

	result := make([]*structs.Distribution, 0)

	for _, distribution := range cache.Distributions {

		matches_name := query.MatchesName(distribution.Name)
		matches_version := query.MatchesVersion(distribution.Version)
		matches_manager := query.MatchesManager(distribution.Manager)
		matches_vendor := query.MatchesVendor(distribution.Vendor)

		if matches_name && matches_version && matches_manager && matches_vendor {
			result = append(result, distribution)
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

func (cache *Distributions) QueryByKernel(query_name string, query_version string) []*structs.Distribution {

	cache.mutex.RLock()

	result := make([]*structs.Distribution, 0)
	found := make(map[string]*structs.Distribution)

	if query_name == "any" {

		if query_version == "any" {

			for _, distribution := range cache.Distributions {
				found[distribution.Name] = distribution
			}

		} else if query_version != "" {

			for _, distribution := range cache.Distributions {

				if distribution.KernelVersion == query_version {
					found[distribution.Name] = distribution
				}

			}

		}

	} else if query_name != "" {

		if query_version == "any" {

			for _, distribution := range cache.Distributions {

				if distribution.Kernel == query_name {
					found[distribution.Name] = distribution
				}

			}

		} else if query_version != "" {

			for _, distribution := range cache.Distributions {

				if distribution.Kernel == query_name && distribution.KernelVersion == query_version {
					found[distribution.Name] = distribution
				}

			}

		}

	}

	if len(found) > 0 {

		for _, distribution := range found {
			result = append(result, distribution)
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

func (cache *Distributions) QueryByKeyword(query_name string, query_value string) []*structs.Distribution {

	cache.mutex.RLock()

	result := make([]*structs.Distribution, 0)
	found := make(map[string]*structs.Distribution)

	if query_name != "" && query_name != "any" {

		if query_value == "any" {

			for _, distribution := range cache.Distributions {

				if distribution.Keywords != nil {

					_, ok := (*distribution.Keywords)[query_name]

					if ok == true {
						found[distribution.Name] = distribution
					}

				}

			}

		} else if query_value != "" {

			for _, distribution := range cache.Distributions {

				if distribution.Keywords != nil {

					val, ok := (*distribution.Keywords)[query_name]

					if ok == true && val == query_value {
						found[distribution.Name] = distribution
					}

				}

			}

		}

	}

	if len(found) > 0 {

		for _, distribution := range found {
			result = append(result, distribution)
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

func (cache *Distributions) QueryByKeywords(keywords map[string]string) []*structs.Distribution {

	cache.mutex.RLock()

	result := make([]*structs.Distribution, 0)
	found := make(map[string]*structs.Distribution)

	if len(keywords) > 0 {

		for _, distribution := range cache.Distributions {

			if distribution.Keywords != nil {

				if len(*distribution.Keywords) > 0 {

					matches_keywords := true

					for key, val := range *distribution.Keywords {

						value, ok1 := keywords[key]

						if ok1 == true {

							if value != val {
								matches_keywords = false
							}

						} else {
							matches_keywords = false
							break
						}

					}

					if matches_keywords {
						found[distribution.Name] = distribution
					}

				}

			}

		}

	}

	if len(found) > 0 {

		for _, distribution := range found {
			result = append(result, distribution)
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

func (cache *Distributions) Remove(name string) bool {

	cache.mutex.Lock()

	var result bool

	if name != "" {

		_, ok := cache.Distributions[name]

		if ok == true {
			delete(cache.Distributions, name)
			result = true
		}

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Distributions) MarshalJSON() ([]byte, error) {

	tmp := make(map[string]*structs.Distribution)
	names := make([]string, 0)

	for name, _ := range cache.Distributions {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {
		tmp[name] = cache.Distributions[name]
	}

	return json.MarshalIndent(&struct {
		Distributions map[string]*structs.Distribution `json:"distributions"`
	}{
		Distributions: tmp,
	}, "", "\t")

}

func (cache *Distributions) UnmarshalJSON(data []byte) error {

	var tmp struct {
		Distributions map[string]*structs.Distribution `json:"distributions"`
	}

	err := json.Unmarshal(data, &tmp)

	if err != nil {
		return err
	}

	for _, distribution := range tmp.Distributions {
		cache.Add(*distribution)
	}

	return nil

}
