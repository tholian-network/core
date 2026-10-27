package caches

import "tholian-core/matchers"
import "tholian-core/structs"
import "encoding/json"
import "sort"
import "strings"
import "sync"

type Providers struct {
	Providers    map[string]*structs.Provider `json:"providers"`
	mapcountries map[string][]string          `json:"-"`
	mapasns      map[string]string            `json:"-"`
	networks     *Networks                    `json:"-"`
	mutex        *sync.RWMutex                `json:"-"`
}

func NewProviders() *Providers {

	var cache Providers

	cache.Providers = make(map[string]*structs.Provider)

	cache.mapasns = make(map[string]string)
	cache.mapcountries = make(map[string][]string)
	cache.networks = NewNetworks()
	cache.mutex = &sync.RWMutex{}

	return &cache

}

func (cache *Providers) Add(value structs.Provider) bool {

	cache.mutex.Lock()

	var result bool

	if value.Name != "" {

		cache.Providers[value.Name] = &value
		result = true

		for _, subnet := range value.Subnets {

			if subnet.Name != "" && strings.HasPrefix(subnet.Name, "AS") {

				cache.networks.AddSubnet(subnet)
				cache.mapasns[subnet.Name] = value.Name

			}

		}

		_, ok1 := cache.mapcountries[value.Country]

		if ok1 == false {
			cache.mapcountries[value.Country] = make([]string, 0)
		}

		cache.mapcountries[value.Country] = append(cache.mapcountries[value.Country], value.Name)

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Providers) Get(name string) *structs.Provider {

	cache.mutex.RLock()

	var result *structs.Provider = nil

	if name != "" {

		provider, ok := cache.Providers[name]

		if ok == true {
			result = provider
		}

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Providers) Length() int {
	return len(cache.Providers)
}

func (cache *Providers) Query(query matchers.Provider) []*structs.Provider {

	cache.mutex.RLock()

	result := make([]*structs.Provider, 0)
	found := make(map[string]*structs.Provider)

	if query.Name != "any" {

		// Cache Optimizations for specified query.Name

		provider, ok := cache.Providers[query.Name]

		if ok == true {

			matches_country := query.MatchesCountry(provider.Country)
			matches_subnet := false

			if query.Subnet == "any" {

				matches_subnet = true

			} else {

				for _, subnet := range provider.Subnets {

					if query.MatchesSubnet(subnet.String()) {
						matches_subnet = true
						break
					}

				}

			}

			if matches_country && matches_subnet {
				found[provider.Name] = provider
			}

		}

	} else if query.Country != "any" {

		// Cache Optimizations for specified query.Country

		_, ok1 := cache.mapcountries[query.Country]

		if ok1 == true {

			for _, name := range cache.mapcountries[query.Country] {

				provider, ok := cache.Providers[name]

				if ok == true {

					matches_name := query.MatchesName(provider.Name)
					matches_subnet := false

					if query.Subnet == "any" {

						matches_subnet = true

					} else {

						for _, subnet := range provider.Subnets {

							if query.MatchesSubnet(subnet.String()) {
								matches_subnet = true
								break
							}

						}

					}

					if matches_name && matches_subnet {
						found[provider.Name] = provider
					}

				}

			}

		}

	} else if query.Subnet != "any" {

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

				name, ok1 := cache.mapasns[asn]

				if ok1 == true {

					provider, ok := cache.Providers[name]

					if ok == true {

						matches_name := query.MatchesName(provider.Name)
						matches_country := query.MatchesCountry(provider.Country)

						if matches_name && matches_country {
							found[provider.Name] = provider
						}

					}

				}

			}

		}

	} else {

		for _, provider := range cache.Providers {

			matches_name := query.MatchesName(provider.Name)
			matches_country := query.MatchesCountry(provider.Country)
			matches_subnet := false

			if query.Subnet == "any" {

				matches_subnet = true

			} else {

				for _, subnet := range provider.Subnets {

					if query.MatchesSubnet(subnet.String()) {
						matches_subnet = true
						break
					}

				}

			}

			if matches_name && matches_country && matches_subnet {
				found[provider.Name] = provider
			}

		}

	}

	if len(found) > 0 {

		for _, provider := range found {
			result = append(result, provider)
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

func (cache *Providers) QueryByIP(query string) []*structs.Provider {

	cache.mutex.RLock()

	result := make([]*structs.Provider, 0)
	found := make(map[string]*structs.Provider)

	if query != "" && query != "any" {

		subnets := cache.networks.QueryByIP(query)

		if len(subnets) > 0 {

			asns := make(map[string]bool)

			for _, subnet := range subnets {

				if subnet.Name != "" && strings.HasPrefix(subnet.Name, "AS") {
					asns[subnet.Name] = true
				}

			}

			for asn, _ := range asns {

				name, ok1 := cache.mapasns[asn]

				if ok1 == true {

					provider, ok := cache.Providers[name]

					if ok == true {
						found[provider.Name] = provider
					}

				}

			}

		}

	}

	if len(found) > 0 {

		for _, provider := range found {
			result = append(result, provider)
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

func (cache *Providers) Remove(name string) bool {

	cache.mutex.Lock()

	var result bool

	if name != "" {

		provider, ok := cache.Providers[name]

		if ok == true {

			delete(cache.Providers, name)
			result = true

			for _, subnet := range provider.Subnets {

				_, ok1 := cache.mapasns[subnet.Name]

				if ok1 == true {
					delete(cache.mapasns, subnet.Name)
				}

				check1 := cache.networks.RemoveSubnet(subnet)

				if check1 == false {
					result = false
				}

			}

			others, ok2 := cache.mapcountries[provider.Country]

			if ok2 == true {

				filtered := make([]string, 0)

				for _, other := range others {

					if other != name {
						filtered = append(filtered, other)
					}

				}

				if len(filtered) > 0 {
					cache.mapcountries[provider.Country] = filtered
				} else {
					delete(cache.mapcountries, provider.Country)
				}

			}

		}

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Providers) MarshalJSON() ([]byte, error) {

	tmp := make(map[string]*structs.Provider)
	names := make([]string, 0)

	for name, _ := range cache.Providers {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {
		tmp[name] = cache.Providers[name]
	}

	return json.MarshalIndent(&struct {
		Providers map[string]*structs.Provider `json:"providers"`
	}{
		Providers: tmp,
	}, "", "\t")

}

func (cache *Providers) UnmarshalJSON(data []byte) error {

	var tmp struct {
		Providers map[string]*structs.Provider `json:"providers"`
	}

	err := json.Unmarshal(data, &tmp)

	if err != nil {
		return err
	}

	for _, provider := range tmp.Providers {
		cache.Add(*provider)
	}

	return nil

}
