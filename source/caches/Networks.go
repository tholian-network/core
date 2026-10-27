package caches

import "tholian-core/matchers"
import "tholian-core/structs"
import "tholian-core/types"
import "encoding/hex"
import "encoding/json"
import "sort"
import "strings"
import "sync"

func addNetworksSubnet(cache *Networks, value structs.Subnet) bool {

	var result bool

	if value.Type == "ipv6" {

		prefix := value.Prefix
		hash := value.Hash()

		_, ok1 := cache.mapv6[prefix]

		if ok1 == false {
			cache.mapv6[prefix] = make(map[string]*structs.Subnet)
		}

		_, ok2 := cache.mapv6[prefix][hash]

		if ok2 == true {
			result = false
		} else {
			cache.mapv6[prefix][hash] = &value
			result = true
		}

	} else if value.Type == "ipv4" {

		prefix := value.Prefix
		hash := value.Hash()

		_, ok1 := cache.mapv4[prefix]

		if ok1 == false {
			cache.mapv4[prefix] = make(map[string]*structs.Subnet)
		}

		_, ok2 := cache.mapv4[prefix][hash]

		if ok2 == true {
			result = false
		} else {
			cache.mapv4[prefix][hash] = &value
			result = true
		}

	}

	return result

}

func queryNetworksByIP(cache *Networks, query string) []*structs.Subnet {

	result := make([]*structs.Subnet, 0)

	if query != "" && query != "any" {

		if strings.Contains(query, "/") {

			query_ip := query[0:strings.Index(query, "/")]
			query_prefix := query[strings.Index(query, "/")+1:]

			if types.IsIPv6(query_ip) {

				ipv6 := types.ParseIPv6(query_ip)
				prefix := parsePrefix(query_prefix)

				if ipv6 != nil && prefix != nil {

					prefixes := make([]uint8, 0)

					for tmp, _ := range cache.mapv6 {

						if tmp <= uint8(*prefix) {
							prefixes = append(prefixes, tmp)
						}

					}

					sort.Slice(prefixes, func(a int, b int) bool {
						return prefixes[a] > prefixes[b]
					})

					for _, prefix := range prefixes {

						hash := hex.EncodeToString(ipv6.Bytes(prefix))
						subnet, ok := cache.mapv6[prefix][hash]

						if ok == true {
							result = append(result, subnet)
						}

					}

				}

			} else if types.IsIPv4(query_ip) {

				ipv4 := types.ParseIPv4(query_ip)
				prefix := parsePrefix(query_prefix)

				if ipv4 != nil && prefix != nil {

					prefixes := make([]uint8, 0)

					for tmp, _ := range cache.mapv4 {

						if tmp <= uint8(*prefix) {
							prefixes = append(prefixes, tmp)
						}

					}

					sort.Slice(prefixes, func(a int, b int) bool {
						return prefixes[a] > prefixes[b]
					})

					for _, prefix := range prefixes {

						hash := hex.EncodeToString(ipv4.Bytes(prefix))
						subnet, ok := cache.mapv4[prefix][hash]

						if ok == true {
							result = append(result, subnet)
						}

					}

				}

			}

		} else {

			if types.IsIPv6(query) {

				ipv6 := types.ParseIPv6(query)
				prefix := uint8(128)

				if ipv6 != nil {

					prefixes := make([]uint8, 0)

					for tmp, _ := range cache.mapv6 {

						if tmp <= prefix {
							prefixes = append(prefixes, tmp)
						}

					}

					sort.Slice(prefixes, func(a int, b int) bool {
						return prefixes[a] > prefixes[b]
					})

					for _, prefix := range prefixes {

						hash := hex.EncodeToString(ipv6.Bytes(prefix))
						subnet, ok := cache.mapv6[prefix][hash]

						if ok == true {
							result = append(result, subnet)
						}

					}

				}

			} else if types.IsIPv4(query) {

				ipv4 := types.ParseIPv4(query)
				prefix := uint8(32)

				if ipv4 != nil {

					prefixes := make([]uint8, 0)

					for tmp, _ := range cache.mapv4 {

						if tmp <= prefix {
							prefixes = append(prefixes, tmp)
						}

					}

					sort.Slice(prefixes, func(a int, b int) bool {
						return prefixes[a] > prefixes[b]
					})

					for _, prefix := range prefixes {

						hash := hex.EncodeToString(ipv4.Bytes(prefix))
						subnet, ok := cache.mapv4[prefix][hash]

						if ok == true {
							result = append(result, subnet)
						}

					}

				}

			}

		}

	}

	if len(result) > 0 {

		sort.Slice(result, func(a int, b int) bool {
			return result[a].Prefix > result[b].Prefix
		})

	}

	return result

}

func removeNetworksSubnet(cache *Networks, value structs.Subnet) bool {

	var result bool

	if value.Type == "ipv6" {

		prefix := value.Prefix
		hash := value.Hash()

		_, ok1 := cache.mapv6[prefix]

		if ok1 == true {

			_, ok2 := cache.mapv6[prefix][hash]

			if ok2 == true {
				delete(cache.mapv6[prefix], hash)
				result = true
			}

		}

	} else if value.Type == "ipv4" {

		prefix := value.Prefix
		hash := value.Hash()

		_, ok1 := cache.mapv4[prefix]

		if ok1 == true {

			_, ok2 := cache.mapv4[prefix][hash]

			if ok2 == true {
				delete(cache.mapv4[prefix], hash)
				result = true
			}

		}

	}

	return result

}

type Networks struct {
	Networks map[string]*structs.Network          `json:"networks"`
	mapv4    map[uint8]map[string]*structs.Subnet `json:"-"`
	mapv6    map[uint8]map[string]*structs.Subnet `json:"-"`
	mutex    *sync.RWMutex                        `json:"-"`
}

func NewNetworks() *Networks {

	var cache Networks

	cache.Networks = make(map[string]*structs.Network)
	cache.mapv4 = make(map[uint8]map[string]*structs.Subnet)
	cache.mapv6 = make(map[uint8]map[string]*structs.Subnet)
	cache.mutex = &sync.RWMutex{}

	return &cache

}

func (cache *Networks) Add(value structs.Network) bool {

	cache.mutex.Lock()

	var result bool

	if value.Name != "" && strings.HasPrefix(value.Name, "AS") {

		cache.Networks[value.Name] = &value

		result = true

		for _, subnet := range value.Subnets {

			check := addNetworksSubnet(cache, subnet)

			if check == false {
				result = false
			}

		}

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Networks) AddSubnet(value structs.Subnet) bool {

	cache.mutex.Lock()

	result := addNetworksSubnet(cache, value)

	cache.mutex.Unlock()

	return result

}

func (cache *Networks) Get(name string) *structs.Network {

	cache.mutex.RLock()

	var result *structs.Network = nil

	if name != "" && strings.HasPrefix(name, "AS") {

		network, ok := cache.Networks[name]

		if ok == true {
			result = network
		}

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Networks) GetByIP(ipstr string) *structs.Network {

	cache.mutex.RLock()

	var result *structs.Network = nil

	if ipstr != "" && ipstr != "any" {

		if strings.Contains(ipstr, "/") {

			tmp1 := ipstr[0:strings.Index(ipstr, "/")]
			tmp2 := ipstr[strings.Index(ipstr, "/")+1:]

			if types.IsIPv6(tmp1) {

				ipv6 := types.ParseIPv6(tmp1)
				prefix := parsePrefix(tmp2)

				if ipv6 != nil && prefix != nil {

					prefixes := make([]uint8, 0)

					for tmp, _ := range cache.mapv6 {

						if tmp <= uint8(*prefix) {
							prefixes = append(prefixes, tmp)
						}

					}

					sort.Slice(prefixes, func(a int, b int) bool {
						return prefixes[a] > prefixes[b]
					})

					for _, prefix := range prefixes {

						hash := hex.EncodeToString(ipv6.Bytes(prefix))
						subnet, ok1 := cache.mapv6[prefix][hash]

						if ok1 == true {

							network, ok := cache.Networks[subnet.Name]

							if ok == true {
								result = network
							}

							break

						}

					}

				}

			} else if types.IsIPv4(tmp1) {

				ipv4 := types.ParseIPv4(tmp1)
				prefix := parsePrefix(tmp2)

				if ipv4 != nil && prefix != nil {

					prefixes := make([]uint8, 0)

					for tmp, _ := range cache.mapv4 {

						if tmp <= uint8(*prefix) {
							prefixes = append(prefixes, tmp)
						}

					}

					sort.Slice(prefixes, func(a int, b int) bool {
						return prefixes[a] > prefixes[b]
					})

					for _, prefix := range prefixes {

						hash := hex.EncodeToString(ipv4.Bytes(prefix))
						subnet, ok1 := cache.mapv4[prefix][hash]

						if ok1 == true {

							network, ok := cache.Networks[subnet.Name]

							if ok == true {
								result = network
							}

							break

						}

					}

				}

			}

		} else {

			if types.IsIPv6(ipstr) {

				ipv6 := types.ParseIPv6(ipstr)
				prefix := uint8(128)

				if ipv6 != nil {

					prefixes := make([]uint8, 0)

					for tmp, _ := range cache.mapv6 {

						if tmp <= prefix {
							prefixes = append(prefixes, tmp)
						}

					}

					sort.Slice(prefixes, func(a int, b int) bool {
						return prefixes[a] > prefixes[b]
					})

					for _, prefix := range prefixes {

						hash := hex.EncodeToString(ipv6.Bytes(prefix))
						subnet, ok1 := cache.mapv6[prefix][hash]

						if ok1 == true {

							network, ok := cache.Networks[subnet.Name]

							if ok == true {
								result = network
							}

							break

						}

					}

				}

			} else if types.IsIPv4(ipstr) {

				ipv4 := types.ParseIPv4(ipstr)
				prefix := uint8(32)

				if ipv4 != nil {

					prefixes := make([]uint8, 0)

					for tmp, _ := range cache.mapv4 {

						if tmp <= prefix {
							prefixes = append(prefixes, tmp)
						}

					}

					sort.Slice(prefixes, func(a int, b int) bool {
						return prefixes[a] > prefixes[b]
					})

					for _, prefix := range prefixes {

						hash := hex.EncodeToString(ipv4.Bytes(prefix))
						subnet, ok1 := cache.mapv4[prefix][hash]

						if ok1 == true {

							network, ok := cache.Networks[subnet.Name]

							if ok == true {
								result = network
							}

							break

						}

					}

				}

			}

		}

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Networks) Length() int {
	return len(cache.Networks)
}

func (cache *Networks) Query(query matchers.Network) []*structs.Network {

	cache.mutex.RLock()

	result := make([]*structs.Network, 0)
	found := make(map[string]*structs.Network)

	if query.Name != "any" {

		// Cache Optimizations for specified query.Name

		network, ok := cache.Networks[query.Name]

		if ok == true {

			matches_subnet := false

			if query.Subnet == "any" {

				matches_subnet = true

			} else {

				for _, subnet := range network.Subnets {

					if query.MatchesSubnet(subnet.String()) {
						matches_subnet = true
						break
					}

				}

			}

			if matches_subnet {
				found[network.Name] = network
			}

		}

	} else if query.Subnet != "any" {

		// Cache Optimizations for specified query.Subnet

		subnets := queryNetworksByIP(cache, query.Subnet)

		if len(subnets) > 0 {

			asns := make([]string, 0)

			for _, subnet := range subnets {

				if subnet.Name != "" && strings.HasPrefix(subnet.Name, "AS") {
					asns = append(asns, subnet.Name)
				}

			}

			for _, asn := range asns {

				network, ok1 := cache.Networks[asn]

				if ok1 == true {

					matches_name := query.MatchesName(network.Name)

					if matches_name {
						found[network.Name] = network
					}

				}

			}

		}

	} else {

		// XXX: Technically unreachable, but left for compatibility to future matcher properties

		for _, network := range cache.Networks {

			matches_name := query.MatchesName(network.Name)
			matches_subnet := false

			if query.Subnet == "any" {

				matches_subnet = true

			} else {

				for _, subnet := range network.Subnets {

					if query.MatchesSubnet(subnet.String()) {
						matches_subnet = true
						break
					}

				}

			}

			if matches_name && matches_subnet {
				found[network.Name] = network
			}

		}

	}

	if len(found) > 0 {

		for _, network := range found {
			result = append(result, network)
		}

	}

	if len(result) > 0 {

		sort.Slice(result, func(a int, b int) bool {
			return result[a].Name < result[b].Name
		})

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Networks) QueryByIP(query string) []*structs.Subnet {

	cache.mutex.RLock()

	result := queryNetworksByIP(cache, query)

	cache.mutex.RUnlock()

	return result

}

func (cache *Networks) Remove(asn string) bool {

	cache.mutex.Lock()

	var result bool

	if asn != "" && strings.HasPrefix(asn, "AS") {

		network, ok := cache.Networks[asn]

		if ok == true {

			delete(cache.Networks, asn)
			result = true

			for _, subnet := range network.Subnets {

				check := removeNetworksSubnet(cache, subnet)

				if check == false {
					result = false
				}

			}

		}

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Networks) RemoveSubnet(value structs.Subnet) bool {

	cache.mutex.Lock()

	result := removeNetworksSubnet(cache, value)

	cache.mutex.Unlock()

	return result

}

func (cache *Networks) MarshalJSON() ([]byte, error) {

	tmp := make(map[string]*structs.Network)
	asns := make([]string, 0)

	for asn, _ := range cache.Networks {
		asns = append(asns, asn)
	}

	sort.Strings(asns)

	for _, asn := range asns {
		tmp[asn] = cache.Networks[asn]
	}

	return json.MarshalIndent(&struct {
		Networks map[string]*structs.Network `json:"networks"`
	}{
		Networks: tmp,
	}, "", "\t")

}

func (cache *Networks) UnmarshalJSON(data []byte) error {

	var tmp struct {
		Networks map[string]*structs.Network `json:"networks"`
	}

	err := json.Unmarshal(data, &tmp)

	if err != nil {
		return err
	}

	for _, network := range tmp.Networks {
		cache.Add(*network)
	}

	return nil

}
