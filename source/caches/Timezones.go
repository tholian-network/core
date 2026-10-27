package caches

import "tholian-core/matchers"
import "tholian-core/structs"
import "encoding/json"
import "sort"
import "strings"
import "sync"

type Timezones struct {
	Timezones map[string]*structs.Timezone `json:"timezones"`
	mutex     *sync.RWMutex                `json:"-"`
}

func NewTimezones() *Timezones {

	var cache Timezones

	cache.Timezones = make(map[string]*structs.Timezone)
	cache.mutex = &sync.RWMutex{}

	return &cache

}

func (cache *Timezones) Add(value structs.Timezone) bool {

	cache.mutex.Lock()

	var result bool

	if value.Name != "" {

		cache.Timezones[value.Name] = &value
		result = true

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Timezones) Get(name string) *structs.Timezone {

	cache.mutex.RLock()

	var result *structs.Timezone

	if name != "" {

		timezone, ok := cache.Timezones[name]

		if ok == true {
			result = timezone
		}

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Timezones) Length() int {
	return len(cache.Timezones)
}

func (cache *Timezones) Query(query matchers.Timezone) []*structs.Timezone {

	cache.mutex.RLock()

	result := make([]*structs.Timezone, 0)
	found := make(map[string]*structs.Timezone)

	if query.Name != "any" && strings.Contains(query.Name, "*") == false {

		// Cache Optimizations for specified query.Name

		timezone, ok := cache.Timezones[query.Name]

		if ok == true {

			matches_offset := query.MatchesOffset(timezone.Offset)

			if matches_offset {
				found[timezone.Name] = timezone
			}

		}

	} else {

		for _, timezone := range cache.Timezones {

			matches_name := query.MatchesName(timezone.Name)
			matches_offset := query.MatchesOffset(timezone.Offset)

			if matches_name && matches_offset {
				found[timezone.Name] = timezone
			}

		}

	}

	if len(found) > 0 {

		for _, timezone := range found {
			result = append(result, timezone)
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

func (cache *Timezones) Remove(name string) bool {

	cache.mutex.Lock()

	var result bool

	if name != "" {

		_, ok := cache.Timezones[name]

		if ok == true {
			delete(cache.Timezones, name)
			result = true
		}

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Timezones) MarshalJSON() ([]byte, error) {

	tmp := make(map[string]*structs.Timezone)
	names := make([]string, 0)

	for name, _ := range cache.Timezones {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {
		tmp[name] = cache.Timezones[name]
	}

	return json.MarshalIndent(&struct {
		Timezones map[string]*structs.Timezone `json:"timezones"`
	}{
		Timezones: tmp,
	}, "", "\t")

}

func (cache *Timezones) UnmarshalJSON(data []byte) error {

	var tmp struct {
		Timezones map[string]*structs.Timezone `json:"timezones"`
	}

	err := json.Unmarshal(data, &tmp)

	if err != nil {
		return err
	}

	for _, timezone := range tmp.Timezones {
		cache.Add(*timezone)
	}

	return nil

}
