package caches

import "tholian-core/matchers"
import "tholian-core/structs"
import "tholian-core/types"
import "encoding/hex"
import "encoding/json"
import "slices"
import "sort"
import "strings"
import "sync"

type Proxies struct {
	Proxies map[string]*structs.Proxy `json:"proxies"`
	mapv4   map[string]string         `json:"-"`
	mapv6   map[string]string         `json:"-"`
	mutex   *sync.RWMutex             `json:"-"`
}

func NewProxies() *Proxies {

	var cache Proxies

	cache.Proxies = make(map[string]*structs.Proxy)
	cache.mapv4 = make(map[string]string)
	cache.mapv6 = make(map[string]string)
	cache.mutex = &sync.RWMutex{}

	return &cache

}

func (cache *Proxies) Add(value structs.Proxy) bool {

	cache.mutex.Lock()

	var result bool

	if value.Name != "" {

		cache.Proxies[value.Name] = &value
		result = true

		for _, address := range value.Addresses {

			if types.IsIPv6(address) {

				ipv6 := types.ParseIPv6(address)

				if ipv6 != nil {

					bytes := ipv6.Bytes(128)
					hash := hex.EncodeToString(bytes)

					_, ok1 := cache.mapv6[hash]

					if ok1 == true {
						result = false
					} else {
						cache.mapv6[hash] = value.Name
						result = true
					}

				}

			} else if types.IsIPv4(address) {

				ipv4 := types.ParseIPv4(address)

				if ipv4 != nil {

					bytes := ipv4.Bytes(32)
					hash := hex.EncodeToString(bytes)

					_, ok1 := cache.mapv4[hash]

					if ok1 == true {
						result = false
					} else {
						cache.mapv4[hash] = value.Name
						result = true
					}

				}

			}

		}

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Proxies) Get(name string) *structs.Proxy {

	cache.mutex.RLock()

	var result *structs.Proxy = nil

	if name != "" {

		proxy, ok := cache.Proxies[name]

		if ok == true {
			result = proxy
		}

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Proxies) Length() int {
	return len(cache.Proxies)
}

func (cache *Proxies) Query(query matchers.Proxy) []*structs.Proxy {

	cache.mutex.RLock()

	result := make([]*structs.Proxy, 0)
	found := make(map[string]*structs.Proxy)

	if query.Name != "any" {

		// Cache Optimizations for specified query.Name

		proxy, ok := cache.Proxies[query.Name]

		if ok == true {

			matches_host := false
			matches_port := query.MatchesPort(proxy.Port)
			matches_protocol := query.MatchesProtocol(proxy.Protocol.String())

			if query.Host == "any" {

				matches_host = true

			} else {

				for _, address := range proxy.Addresses {

					if query.MatchesHost(address) {
						matches_host = true
						break
					}

				}

			}

			if matches_host && matches_port && matches_protocol {
				found[proxy.Name] = proxy
			}

		}

	} else {

		for _, proxy := range cache.Proxies {

			matches_name := query.MatchesName(proxy.Name)
			matches_host := false
			matches_port := query.MatchesPort(proxy.Port)
			matches_protocol := query.MatchesProtocol(proxy.Protocol.String())

			if query.Host == "any" {

				matches_host = true

			} else {

				for _, address := range proxy.Addresses {

					if query.MatchesHost(address) {
						matches_host = true
						break
					}

				}

			}

			if matches_name && matches_host && matches_port && matches_protocol {
				found[proxy.Name] = proxy
			}

		}

	}

	if len(found) > 0 {

		for _, proxy := range found {
			result = append(result, proxy)
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

func (cache *Proxies) QueryByIP(query string) []*structs.Proxy {

	cache.mutex.RLock()

	result := make([]*structs.Proxy, 0)
	found := make(map[string]*structs.Proxy)

	if query != "" && query != "any" {

		if strings.Contains(query, "/") {

			tmp1 := query[0:strings.Index(query, "/")]
			tmp2 := query[strings.Index(query, "/")+1:]

			if types.IsIPv6(tmp1) {

				ipv6 := types.ParseIPv6(tmp1)
				prefix := parsePrefix(tmp2)

				if ipv6 != nil && prefix != nil {

					hash := hex.EncodeToString(ipv6.Bytes(uint8(*prefix)))
					name, ok1 := cache.mapv6[hash]

					if ok1 == true {

						proxy, ok := cache.Proxies[name]

						if ok == true {
							found[proxy.Name] = proxy
						}

					}

				}

			} else if types.IsIPv4(tmp1) {

				ipv4 := types.ParseIPv4(tmp1)
				prefix := parsePrefix(tmp2)

				if ipv4 != nil && prefix != nil {

					hash := hex.EncodeToString(ipv4.Bytes(uint8(*prefix)))
					name, ok1 := cache.mapv4[hash]

					if ok1 == true {

						proxy, ok := cache.Proxies[name]

						if ok == true {
							found[proxy.Name] = proxy
						}

					}

				}

			}

		} else {

			if types.IsIPv6(query) {

				ipv6 := types.ParseIPv6(query)
				prefix := uint8(128)

				if ipv6 != nil {

					hash := hex.EncodeToString(ipv6.Bytes(prefix))
					name, ok1 := cache.mapv6[hash]

					if ok1 == true {

						proxy, ok := cache.Proxies[name]

						if ok == true {
							found[proxy.Name] = proxy
						}

					}

				}

			} else if types.IsIPv4(query) {

				ipv4 := types.ParseIPv4(query)
				prefix := uint8(32)

				if ipv4 != nil {

					hash := hex.EncodeToString(ipv4.Bytes(prefix))
					name, ok1 := cache.mapv4[hash]

					if ok1 == true {

						proxy, ok := cache.Proxies[name]

						if ok == true {
							found[proxy.Name] = proxy
						}

					}

				}

			}

		}

	}

	if len(found) > 0 {

		for _, proxy := range found {
			result = append(result, proxy)
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

func (cache *Proxies) Remove(name string) bool {

	cache.mutex.Lock()

	var result bool

	if name != "" {

		proxy, ok := cache.Proxies[name]

		if ok == true {

			delete(cache.Proxies, name)
			result = true

			for _, address := range proxy.Addresses {

				if types.IsIPv6(address) {

					ipv6 := types.ParseIPv6(address)

					if ipv6 != nil {

						bytes := ipv6.Bytes(128)
						hash := hex.EncodeToString(bytes)

						_, ok1 := cache.mapv6[hash]

						if ok1 == true {
							delete(cache.mapv6, hash)
							result = true
						}

					}

				} else if types.IsIPv4(address) {

					ipv4 := types.ParseIPv4(address)

					if ipv4 != nil {

						bytes := ipv4.Bytes(32)
						hash := hex.EncodeToString(bytes)

						_, ok1 := cache.mapv4[hash]

						if ok1 == true {
							delete(cache.mapv4, hash)
							result = true
						}

					}

				}

			}

		}

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Proxies) MarshalJSON() ([]byte, error) {

	tmp := make(map[string]*structs.Proxy)
	names := make([]string, 0)

	for name, _ := range cache.Proxies {
		names = append(names, name)
	}

	slices.Sort(names)

	for _, name := range names {
		tmp[name] = cache.Proxies[name]
	}

	return json.MarshalIndent(&struct {
		Proxies map[string]*structs.Proxy `json:"proxies"`
	}{
		Proxies: tmp,
	}, "", "\t")

}

func (cache *Proxies) UnmarshalJSON(data []byte) error {

	var tmp struct {
		Proxies map[string]*structs.Proxy `json:"proxies"`
	}

	err := json.Unmarshal(data, &tmp)

	if err != nil {
		return err
	}

	for _, proxy := range tmp.Proxies {
		cache.Add(*proxy)
	}

	return nil

}
