package caches

import "tholian-core/matchers"
import "tholian-core/structs"
import "tholian-core/types"
import "encoding/json"
import "sort"
import "strings"
import "sync"

type Countries struct {
	Countries map[string]*structs.Country `json:"countries"`
	mapasns   map[string]string           `json:"-"`
	networks  *Networks                   `json:"-"`
	mutex     *sync.RWMutex               `json:"-"`
}

func NewCountries() *Countries {

	var cache Countries

	cache.Countries = make(map[string]*structs.Country)
	cache.mapasns = make(map[string]string)
	cache.networks = NewNetworks()
	cache.mutex = &sync.RWMutex{}

	return &cache

}

func (cache *Countries) Add(value structs.Country) bool {

	cache.mutex.Lock()

	var result bool

	if value.ISO != "" && value.Name != "" {

		cache.Countries[value.ISO] = &value
		result = true

		for _, subnet := range value.Subnets {

			if subnet.Name != "" && strings.HasPrefix(subnet.Name, "AS") {

				cache.networks.AddSubnet(subnet)
				cache.mapasns[subnet.Name] = value.ISO

			}

		}

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Countries) Get(iso string) *structs.Country {

	cache.mutex.RLock()

	var result *structs.Country = nil

	if iso != "" {

		country, ok := cache.Countries[iso]

		if ok == true {
			result = country
		}

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Countries) GetByName(name string) *structs.Country {

	cache.mutex.RLock()

	var result *structs.Country = nil

	if name != "" && name != "any" {

		for _, country := range cache.Countries {

			if country.Name == name {
				result = country
				break
			}

		}

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Countries) Length() int {
	return len(cache.Countries)
}

func (cache *Countries) Query(query matchers.Country) []*structs.Country {

	cache.mutex.RLock()

	result := make([]*structs.Country, 0)
	found := make(map[string]*structs.Country)

	if query.Subnet != "any" {

		// Cache Optimizations for specified query.Subnet

		subnets := cache.networks.QueryByIP(query.Subnet)

		if len(subnets) > 0 {

			asns := make([]string, 0)

			for _, subnet := range subnets {

				if subnet.Name != "" && strings.HasPrefix(subnet.Name, "AS") {
					asns = append(asns, subnet.Name)
				}

			}

			for _, asn := range asns {

				iso, ok1 := cache.mapasns[asn]

				if ok1 == true {

					country, ok := cache.Countries[iso]

					if ok == true {

						matches_name := query.MatchesName(country.Name)
						matches_continent := query.MatchesContinent(country.Continent)
						matches_allegiance := false
						matches_timezone := false

						if query.Allegiance == "any" {

							matches_allegiance = true

						} else {

							for _, allegiance := range country.Allegiances {

								if query.MatchesAllegiance(allegiance) {
									matches_allegiance = true
									break
								}

							}

						}

						if query.Timezone == "any" {

							matches_timezone = true

						} else {

							for _, timezone := range country.Timezones {

								if query.MatchesTimezone(timezone.Name) {
									matches_timezone = true
									break
								}

							}

						}

						if matches_name && matches_continent && matches_allegiance && matches_timezone {
							found[country.ISO] = country
						}

					}

				}

			}

		}

	} else {

		for _, country := range cache.Countries {

			matches_name := query.MatchesName(country.Name)
			matches_continent := query.MatchesContinent(country.Continent)
			matches_allegiance := false
			matches_subnet := false
			matches_timezone := false

			if query.Allegiance == "any" {

				matches_allegiance = true

			} else {

				for _, allegiance := range country.Allegiances {

					if query.MatchesAllegiance(allegiance) {
						matches_allegiance = true
						break
					}

				}

			}

			if query.Subnet == "any" {

				matches_subnet = true

			} else {

				for _, subnet := range country.Subnets {

					if query.MatchesSubnet(subnet.String()) {
						matches_subnet = true
						break
					}

				}

			}

			if query.Timezone == "any" {

				matches_timezone = true

			} else {

				for _, timezone := range country.Timezones {

					if query.MatchesTimezone(timezone.Name) {
						matches_timezone = true
						break
					}

				}

			}

			if matches_name && matches_continent && matches_allegiance && matches_subnet && matches_timezone {
				found[country.ISO] = country
			}

		}

	}

	if len(found) > 0 {

		for _, country := range found {
			result = append(result, country)
		}

	}

	if len(result) > 1 {

		sort.Slice(result, func(a int, b int) bool {
			return result[a].ISO < result[b].ISO
		})

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Countries) QueryByGeolocation(query string) []*structs.Country {

	cache.mutex.RLock()

	result := make([]*structs.Country, 0)
	geolocation := types.ToGeolocation(query)

	for _, country := range cache.Countries {

		// Default is Antarctica to not conflict with any Country
		if country.Geolocation.Latitude != -90.0 && country.Geolocation.Longitude != 0.0 {
			result = append(result, country)
		}

	}

	if len(result) > 1 {

		sort.Slice(result, func(a int, b int) bool {
			return result[a].Geolocation.DistanceTo(geolocation) < result[b].Geolocation.DistanceTo(geolocation)
		})

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Countries) QueryByIP(query string) []*structs.Country {

	cache.mutex.RLock()

	result := make([]*structs.Country, 0)
	found := make(map[string]*structs.Country)

	if query != "" && query != "any" {

		subnets := cache.networks.QueryByIP(query)

		if len(subnets) > 0 {

			asns := make([]string, 0)

			for _, subnet := range subnets {

				if subnet.Name != "" && strings.HasPrefix(subnet.Name, "AS") {
					asns = append(asns, subnet.Name)
				}

			}

			for _, asn := range asns {

				iso, ok1 := cache.mapasns[asn]

				if ok1 == true {

					country, ok := cache.Countries[iso]

					if ok == true {
						found[country.ISO] = country
					}

				}

			}

		}

	}

	if len(found) > 0 {

		for _, country := range found {
			result = append(result, country)
		}

	}

	if len(result) > 1 {

		sort.Slice(result, func(a int, b int) bool {
			return result[a].ISO < result[b].ISO
		})

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Countries) Remove(iso string) bool {

	cache.mutex.Lock()

	var result bool

	if iso != "" {

		country, ok := cache.Countries[iso]

		if ok == true {

			delete(cache.Countries, iso)
			result = true

			for _, subnet := range country.Subnets {

				_, ok1 := cache.mapasns[subnet.Name]

				if ok1 == true {
					delete(cache.mapasns, subnet.Name)
				}

				check1 := cache.networks.RemoveSubnet(subnet)

				if check1 == false {
					result = false
				}

			}

		}

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Countries) MarshalJSON() ([]byte, error) {

	tmp := make(map[string]*structs.Country)
	isos := make([]string, 0)

	for iso, _ := range cache.Countries {
		isos = append(isos, iso)
	}

	sort.Strings(isos)

	for _, iso := range isos {
		tmp[iso] = cache.Countries[iso]
	}

	return json.MarshalIndent(&struct {
		Countries map[string]*structs.Country `json:"countries"`
	}{
		Countries: tmp,
	}, "", "\t")

}

func (cache *Countries) UnmarshalJSON(data []byte) error {

	var tmp struct {
		Countries map[string]*structs.Country `json:"countries"`
	}

	err := json.Unmarshal(data, &tmp)

	if err != nil {
		return err
	}

	for _, country := range tmp.Countries {
		cache.Add(*country)
	}

	return nil

}
