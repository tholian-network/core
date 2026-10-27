package caches

import "tholian-core/matchers"
import "tholian-core/structs"
import "encoding/json"
import "sort"
import "sync"

type Timeslots struct {
	Timeslots    map[string]*structs.Timeslot `json:"timeslots"`
	mapcountries map[string][]string          `json:"countries"`
	mutex        *sync.RWMutex                `json:"-"`
}

func NewTimeslots() *Timeslots {

	var cache Timeslots

	cache.Timeslots = make(map[string]*structs.Timeslot)
	cache.mapcountries = make(map[string][]string, 0)
	cache.mutex = &sync.RWMutex{}

	return &cache

}

func (cache *Timeslots) Add(value structs.Timeslot) bool {

	cache.mutex.Lock()

	var result bool

	if value.Offset != "" {

		cache.Timeslots[value.Offset] = &value
		result = true

		for _, country := range value.Countries {

			_, ok1 := cache.mapcountries[country]

			if ok1 == false {
				cache.mapcountries[country] = make([]string, 0)
			}

			cache.mapcountries[country] = append(cache.mapcountries[country], value.Offset)

		}

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Timeslots) Get(offset string) *structs.Timeslot {

	cache.mutex.RLock()

	var result *structs.Timeslot = nil

	if offset != "" {

		timeslot, ok := cache.Timeslots[offset]

		if ok == true {
			result = timeslot
		}

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Timeslots) Length() int {
	return len(cache.Timeslots)
}

func (cache *Timeslots) Query(query matchers.Timeslot) []*structs.Timeslot {

	cache.mutex.RLock()

	result := make([]*structs.Timeslot, 0)
	found := make(map[string]*structs.Timeslot)

	if query.Country != "any" {

		// Cache Optimizations for specified query.Country

		_, ok1 := cache.mapcountries[query.Country]

		if ok1 == true {

			for _, offset := range cache.mapcountries[query.Country] {

				timeslot, ok := cache.Timeslots[offset]

				if ok == true {

					matches_offset := query.MatchesOffset(timeslot.Offset)
					matches_timezone := false

					if query.Timezone == "any" {

						matches_timezone = true

					} else {

						for _, timezone := range timeslot.Timezones {

							if query.MatchesTimezone(timezone) {
								matches_timezone = true
								break
							}

						}

					}

					if matches_offset && matches_timezone {
						found[timeslot.Offset] = timeslot
					}

				}

			}

		}

	} else {

		for _, timeslot := range cache.Timeslots {

			matches_offset := query.MatchesOffset(timeslot.Offset)
			matches_country := false
			matches_timezone := false

			if query.Country == "any" {

				matches_country = true

			} else {

				for _, country := range timeslot.Countries {

					if query.MatchesCountry(country) {
						matches_country = true
						break
					}

				}

			}

			if query.Timezone == "any" {

				matches_timezone = true

			} else {

				for _, timezone := range timeslot.Timezones {

					if query.MatchesTimezone(timezone) {
						matches_timezone = true
						break
					}

				}

			}

			if matches_offset && matches_country && matches_timezone {
				found[timeslot.Offset] = timeslot
			}

		}

	}

	if len(found) > 0 {

		for _, timeslot := range found {
			result = append(result, timeslot)
		}

	}

	if len(result) > 1 {

		sort.Slice(result, func(a int, b int) bool {
			return result[a].Offset < result[b].Offset
		})

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Timeslots) Remove(offset string) bool {

	cache.mutex.Lock()

	var result bool

	if offset != "" {

		timeslot, ok := cache.Timeslots[offset]

		if ok == true {

			delete(cache.Timeslots, offset)
			result = true

			for _, country := range timeslot.Countries {

				others, ok1 := cache.mapcountries[country]

				if ok1 == true {

					filtered := make([]string, 0)

					for _, other := range others {

						if other != offset {
							filtered = append(filtered, other)
						}

					}

					if len(filtered) > 0 {
						cache.mapcountries[country] = filtered
					} else {
						delete(cache.mapcountries, country)
					}

				}

			}

		}

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Timeslots) MarshalJSON() ([]byte, error) {

	tmp := make(map[string]*structs.Timeslot)
	offsets := make([]string, 0)

	for offset, _ := range cache.Timeslots {
		offsets = append(offsets, offset)
	}

	sort.Strings(offsets)

	for _, offset := range offsets {
		tmp[offset] = cache.Timeslots[offset]
	}

	return json.MarshalIndent(&struct {
		Timeslots map[string]*structs.Timeslot `json:"timeslots"`
	}{
		Timeslots: tmp,
	}, "", "\t")

}

func (cache *Timeslots) UnmarshalJSON(data []byte) error {

	var tmp struct {
		Timeslots map[string]*structs.Timeslot `json:"timeslots"`
	}

	err := json.Unmarshal(data, &tmp)

	if err != nil {
		return err
	}

	for _, timeslot := range tmp.Timeslots {
		cache.Add(*timeslot)
	}

	return nil

}

