package caches

import "tholian-core/matchers"
import "tholian-core/structs"
import "encoding/json"
import "sort"
import "strings"
import "sync"

func toResponseIdentifier(value structs.Response) string {
	return value.System + ":" + value.Type + ":" + value.Datetime.String()
}

type Responses struct {
	Responses          map[string]*structs.Response                         `json:"responses"`
	mapasns            map[string][]string                                  `json:"-"`
	mapconnections     map[string]map[string]map[string]map[uint16][]string `json:"-"`
	mapdistributions   map[string][]string                                  `json:"-"`
	mapresponses       map[string]map[string]map[string]string              `json:"-"`
	mappackages        map[string]map[string][]string                       `json:"-"`
	mapproducts        map[string]map[string]map[string][]string            `json:"-"`
	mapprograms        map[string][]string                                  `json:"-"`
	mapusers           map[string][]string                                  `json:"-"`
	mapvulnerabilities map[string][]string                                  `json:"-"`
	mapweaknesses      map[string]map[string]map[string][]string            `json:"-"`
	networks           *Networks                                            `json:"-"`
	mutex              *sync.RWMutex                                        `json:"-"`
}

func NewResponses() *Responses {
	return NewResponsesFromNetworks(NewNetworks())
}

func NewResponsesFromNetworks(networks *Networks) *Responses {

	var cache Responses

	cache.Responses = make(map[string]*structs.Response)
	cache.mapasns = make(map[string][]string, 0)
	cache.mapconnections = make(map[string]map[string]map[string]map[uint16][]string, 0)
	cache.mapdistributions = make(map[string][]string, 0)
	cache.mapresponses = make(map[string]map[string]map[string]string)
	cache.mappackages = make(map[string]map[string][]string, 0)
	cache.mapproducts = make(map[string]map[string]map[string][]string, 0)
	cache.mapprograms = make(map[string][]string, 0)
	cache.mapusers = make(map[string][]string, 0)
	cache.mapvulnerabilities = make(map[string][]string, 0)
	cache.mapweaknesses = make(map[string]map[string]map[string][]string, 0)

	if networks != nil {
		cache.networks = networks
	} else {
		cache.networks = NewNetworks()
	}

	cache.mutex = &sync.RWMutex{}

	return &cache

}

func (cache *Responses) Add(value structs.Response) bool {

	cache.mutex.Lock()

	var result bool

	if value.System != "" && value.Type != "" && value.Datetime.IsValid() {

		identifier := toResponseIdentifier(value)

		cache.Responses[identifier] = &value
		result = true

		_, ok1 := cache.mapresponses[value.System]

		if ok1 == false {
			cache.mapresponses[value.System] = make(map[string]map[string]string)
		}

		_, ok2 := cache.mapresponses[value.System][value.Type]

		if ok2 == false {
			cache.mapresponses[value.System][value.Type] = make(map[string]string)
		}

		cache.mapresponses[value.System][value.Type][value.Datetime.String()] = identifier

		for _, connection := range value.Connections {

			_, ok1 := cache.mapconnections[connection.Type]

			if ok1 == false {
				cache.mapconnections[connection.Type] = make(map[string]map[string]map[uint16][]string)
			}

			_, ok2 := cache.mapconnections[connection.Type][connection.Protocol]

			if ok2 == false {
				cache.mapconnections[connection.Type][connection.Protocol] = make(map[string]map[uint16][]string)
			}

			_, ok3 := cache.mapconnections[connection.Type][connection.Protocol][connection.Host]

			if ok3 == false {
				cache.mapconnections[connection.Type][connection.Protocol][connection.Host] = make(map[uint16][]string)
			}

			_, ok4 := cache.mapconnections[connection.Type][connection.Protocol][connection.Host][connection.Port]

			if ok4 == false {
				cache.mapconnections[connection.Type][connection.Protocol][connection.Host][connection.Port] = make([]string, 0)
			}

			cache.mapconnections[connection.Type][connection.Protocol][connection.Host][connection.Port] = append(cache.mapconnections[connection.Type][connection.Protocol][connection.Host][connection.Port], identifier)

		}

		for _, distribution := range value.Distributions {

			if distribution.Name != "any" {

				_, ok1 := cache.mapdistributions[distribution.Name]

				if ok1 == false {
					cache.mapdistributions[distribution.Name] = make([]string, 0)
				}

				cache.mapdistributions[distribution.Name] = append(cache.mapdistributions[distribution.Name], identifier)

			}

		}

		for _, network := range value.Networks {

			if network.Name != "" && strings.HasPrefix(network.Name, "AS") {

				subnet := structs.ToSubnet(network.Subnet)
				subnet.SetName(network.Name)

				cache.networks.AddSubnet(subnet)

				_, ok1 := cache.mapasns[network.Name]

				if ok1 == false {
					cache.mapasns[network.Name] = make([]string, 0)
				}

				cache.mapasns[network.Name] = append(cache.mapasns[network.Name], identifier)

			} else if network.Subnet != "any" {

				resolved := cache.networks.GetByIP(network.Subnet)

				if resolved != nil {

					_, ok1 := cache.mapasns[resolved.Name]

					if ok1 == false {
						cache.mapasns[resolved.Name] = make([]string, 0)
					}

					cache.mapasns[resolved.Name] = append(cache.mapasns[resolved.Name], identifier)

				}

			}

		}

		for _, pkg := range value.Packages {

			_, ok1 := cache.mappackages[pkg.Vendor]

			if ok1 == false {
				cache.mappackages[pkg.Vendor] = make(map[string][]string, 0)
			}

			_, ok2 := cache.mappackages[pkg.Vendor][pkg.Name]

			if ok2 == false {
				cache.mappackages[pkg.Vendor][pkg.Name] = make([]string, 0)
			}

			cache.mappackages[pkg.Vendor][pkg.Name] = append(cache.mappackages[pkg.Vendor][pkg.Name], identifier)

		}

		for _, product := range value.Products {

			_, ok1 := cache.mapproducts[product.Vendor]

			if ok1 == false {
				cache.mapproducts[product.Vendor] = make(map[string]map[string][]string)
			}

			_, ok2 := cache.mapproducts[product.Vendor][product.Product]

			if ok2 == false {
				cache.mapproducts[product.Vendor][product.Product] = make(map[string][]string)
			}

			_, ok3 := cache.mapproducts[product.Vendor][product.Product][product.Version]

			if ok3 == false {
				cache.mapproducts[product.Vendor][product.Product][product.Version] = make([]string, 0)
			}

			cache.mapproducts[product.Vendor][product.Product][product.Version] = append(cache.mapproducts[product.Vendor][product.Product][product.Version], identifier)

		}

		for _, program := range value.Programs {

			_, ok1 := cache.mapprograms[program.Name]

			if ok1 == false {
				cache.mapprograms[program.Name] = make([]string, 0)
			}

			cache.mapprograms[program.Name] = append(cache.mapprograms[program.Name], identifier)

		}

		for _, user := range value.Users {

			_, ok1 := cache.mapusers[user.Name]

			if ok1 == false {
				cache.mapusers[user.Name] = make([]string, 0)
			}

			cache.mapusers[user.Name] = append(cache.mapusers[user.Name], identifier)

		}

		for _, vulnerability := range value.Vulnerabilities {

			_, ok1 := cache.mapvulnerabilities[vulnerability.Name]

			if ok1 == false {
				cache.mapvulnerabilities[vulnerability.Name] = make([]string, 0)
			}

			cache.mapvulnerabilities[vulnerability.Name] = append(cache.mapvulnerabilities[vulnerability.Name], identifier)

		}

		for _, weakness := range value.Weaknesses {

			_, ok1 := cache.mapweaknesses[weakness.Name]

			if ok1 == false {
				cache.mapweaknesses[weakness.Name] = make(map[string]map[string][]string)
			}

			_, ok2 := cache.mapweaknesses[weakness.Name][weakness.Impact]

			if ok2 == false {
				cache.mapweaknesses[weakness.Name][weakness.Impact] = make(map[string][]string)
			}

			_, ok3 := cache.mapweaknesses[weakness.Name][weakness.Impact][weakness.Scope]

			if ok3 == false {
				cache.mapweaknesses[weakness.Name][weakness.Impact][weakness.Scope] = make([]string, 0)
			}

			cache.mapweaknesses[weakness.Name][weakness.Impact][weakness.Scope] = append(cache.mapweaknesses[weakness.Name][weakness.Impact][weakness.Scope], identifier)

		}

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Responses) Get(identifier string) *structs.Response {

	cache.mutex.RLock()

	var result *structs.Response = nil

	if strings.Contains(identifier, ":") {

		tmp := strings.Split(identifier, ":")

		// System:Type:YYYY-MM-DD HH:ii:SS
		if len(tmp) == 5 {

			response, ok := cache.Responses[identifier]

			if ok == true {
				result = response
			}

		}

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Responses) Length() int {
	return len(cache.Responses)
}

func (cache *Responses) Query(query matchers.Response) []*structs.Response {

	cache.mutex.RLock()

	result := make([]*structs.Response, 0)
	found := make(map[string]*structs.Response)

	if query.System != "any" && query.Type != "any" {

		// Cache Optimizations for specified query.System, query.Type

		_, ok1 := cache.mapresponses[query.System]

		if ok1 == true {

			datetimes, ok2 := cache.mapresponses[query.System][query.Type]

			if ok2 == true {

				for _, identifier := range datetimes {

					response, ok := cache.Responses[identifier]

					if ok == true {

						matches_datetime := query.MatchesDatetime(response.Datetime.String())

						if matches_datetime {
							found[identifier] = response
						}

					}

				}

			}

		}

	} else if query.System != "any" && query.Type == "any" {

		// Cache Optimizations for specified query.System

		_, ok1 := cache.mapresponses[query.System]

		if ok1 == true {

			for _, types := range cache.mapresponses[query.System] {

				for _, identifier := range types {

					response, ok := cache.Responses[identifier]

					if ok == true {

						matches_datetime := query.MatchesDatetime(response.Datetime.String())

						if matches_datetime {
							found[identifier] = response
						}

					}

				}

			}

		}

	} else {

		for identifier, response := range cache.Responses {

			matches_system := query.MatchesSystem(response.System)
			matches_type := query.MatchesType(response.Type)
			matches_datetime := query.MatchesDatetime(response.Datetime.String())

			if matches_system && matches_type && matches_datetime {
				found[identifier] = response
			}

		}

	}

	if len(found) > 0 {

		for _, response := range found {
			result = append(result, response)
		}

	}

	if len(result) > 1 {

		sort.Slice(result, func(a int, b int) bool {
			return result[a].IsAfter(*result[b])
		})

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Responses) QueryByConnection(query matchers.Connection) []*structs.Response {

	cache.mutex.RLock()

	result := make([]*structs.Response, 0)
	found := make(map[string]*structs.Response)

	if query.Type != "any" && query.Protocol != "any" && query.Host != "any" && query.Port != 0 {

		// Cache Optimizations for specified query.Type, query.Protocol, query.Host, query.Port

		_, ok1 := cache.mapconnections[query.Type]

		if ok1 == true {

			_, ok2 := cache.mapconnections[query.Type][query.Protocol]

			if ok2 == true {

				_, ok3 := cache.mapconnections[query.Type][query.Protocol][query.Host]

				if ok3 == true {

					identifiers, ok4 := cache.mapconnections[query.Type][query.Protocol][query.Host][query.Port]

					if ok4 == true {

						for _, identifier := range identifiers {

							response, ok := cache.Responses[identifier]

							if ok == true {

								matches_connection := false

								for _, connection := range response.Connections {

									if query.Matches(connection.Host, connection.Port, connection.Protocol, connection.Type) {
										matches_connection = true
										break
									}

								}

								if matches_connection {
									found[identifier] = response
								}

							}

						}

					}

				}

			}

		}

	} else if query.Type != "any" && query.Protocol != "any" && query.Host != "any" && query.Port == 0 {

		// Cache Optimizations for specified query.Type, query.Protocol, query.Host

		_, ok1 := cache.mapconnections[query.Type]

		if ok1 == true {

			_, ok2 := cache.mapconnections[query.Type][query.Protocol]

			if ok2 == true {

				_, ok3 := cache.mapconnections[query.Type][query.Protocol][query.Host]

				if ok3 == true {

					for _, identifiers := range cache.mapconnections[query.Type][query.Protocol][query.Host] {

						for _, identifier := range identifiers {

							response, ok := cache.Responses[identifier]

							if ok == true {

								matches_connection := false

								for _, connection := range response.Connections {

									if query.Matches(connection.Host, connection.Port, connection.Protocol, connection.Type) {
										matches_connection = true
										break
									}

								}

								if matches_connection {
									found[identifier] = response
								}

							}

						}

					}

				}

			}

		}

	} else if query.Type != "any" && query.Protocol != "any" && query.Host == "any" {

		// Cache Optimizations for specified query.Type, query.Protocol

		_, ok1 := cache.mapconnections[query.Type]

		if ok1 == true {

			_, ok2 := cache.mapconnections[query.Type][query.Protocol]

			if ok2 == true {

				for _, hosts := range cache.mapconnections[query.Type][query.Protocol] {

					for _, identifiers := range hosts {

						for _, identifier := range identifiers {

							response, ok := cache.Responses[identifier]

							if ok == true {

								matches_connection := false

								for _, connection := range response.Connections {

									if query.Matches(connection.Host, connection.Port, connection.Protocol, connection.Type) {
										matches_connection = true
										break
									}

								}

								if matches_connection {
									found[identifier] = response
								}

							}

						}

					}

				}

			}

		}

	} else if query.Type != "any" && query.Protocol == "any" {

		// Cache Optimizations for specified query.Type

		_, ok1 := cache.mapconnections[query.Type]

		if ok1 == true {

			for _, protocols := range cache.mapconnections[query.Type] {

				for _, hosts := range protocols {

					for _, identifiers := range hosts {

						for _, identifier := range identifiers {

							response, ok := cache.Responses[identifier]

							if ok == true {

								matches_connection := false

								for _, connection := range response.Connections {

									if query.Matches(connection.Host, connection.Port, connection.Protocol, connection.Type) {
										matches_connection = true
										break
									}

								}

								if matches_connection {
									found[identifier] = response
								}

							}

						}

					}

				}

			}

		}

	} else {

		for identifier, response := range cache.Responses {

			matches_connection := false

			for _, connection := range response.Connections {

				if query.Matches(connection.Host, connection.Port, connection.Protocol, connection.Type) {
					matches_connection = true
					break
				}

			}

			if matches_connection {
				found[identifier] = response
			}

		}

	}

	if len(found) > 0 {

		for _, response := range found {
			result = append(result, response)
		}

	}

	if len(result) > 1 {

		sort.Slice(result, func(a int, b int) bool {
			return result[a].IsAfter(*result[b])
		})

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Responses) QueryByDistribution(query matchers.Distribution) []*structs.Response {

	cache.mutex.RLock()

	result := make([]*structs.Response, 0)
	found := make(map[string]*structs.Response)

	if query.Name != "any" {

		// Cache Optimizations for specified query.Name

		identifiers, ok1 := cache.mapdistributions[query.Name]

		if ok1 == true {

			for _, identifier := range identifiers {

				response, ok := cache.Responses[identifier]

				if ok == true {

					matches_distribution := false

					for _, distribution := range response.Distributions {

						if query.MatchesVersion(distribution.Version) && query.MatchesManager(distribution.Manager) && query.MatchesVendor(distribution.Vendor) {
							matches_distribution = true
							break
						}

					}

					if matches_distribution {
						found[identifier] = response
					}

				}

			}

		}

	} else {

		for identifier, response := range cache.Responses {

			matches_distribution := false

			for _, distribution := range response.Distributions {

				if query.MatchesName(distribution.Name) && query.MatchesVersion(distribution.Version) && query.MatchesManager(distribution.Manager) && query.MatchesVendor(distribution.Vendor) {
					matches_distribution = true
					break
				}

			}

			if matches_distribution {
				found[identifier] = response
			}

		}

	}

	if len(found) > 0 {

		for _, response := range found {
			result = append(result, response)
		}

	}

	if len(result) > 1 {

		sort.Slice(result, func(a int, b int) bool {
			return result[a].IsAfter(*result[b])
		})

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Responses) QueryByIP(query string) []*structs.Response {

	cache.mutex.RLock()

	result := make([]*structs.Response, 0)
	found := make(map[string]*structs.Response)

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

				identifiers, ok1 := cache.mapasns[asn]

				if ok1 == true {

					for _, identifier := range identifiers {

						response, ok := cache.Responses[identifier]

						if ok == true {
							found[identifier] = response
						}

					}

				}

			}

		}

	}

	if len(found) > 0 {

		for _, response := range found {
			result = append(result, response)
		}

	}

	if len(result) > 1 {

		sort.Slice(result, func(a int, b int) bool {
			return result[a].IsAfter(*result[b])
		})

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Responses) QueryByPackage(query matchers.Package) []*structs.Response {

	cache.mutex.RLock()

	result := make([]*structs.Response, 0)
	found := make(map[string]*structs.Response)

	if query.Vendor != "any" && query.Name != "any" {

		// Cache Optimizations for specified query.Vendor, query.Name

		_, ok1 := cache.mappackages[query.Vendor]

		if ok1 == true {

			identifiers, ok2 := cache.mappackages[query.Vendor][query.Name]

			if ok2 == true {

				for _, identifier := range identifiers {

					response, ok := cache.Responses[identifier]

					if ok == true {

						matches_package := false

						for _, pkg := range response.Packages {

							if query.MatchesVersion(pkg.Version) && query.MatchesArchitecture(pkg.Architecture) && query.MatchesManager(pkg.Manager) {
								matches_package = true
								break
							}

						}

						if matches_package {
							found[identifier] = response
						}

					}

				}

			}

		}

	} else if query.Vendor != "any" && query.Name == "any" {

		// Cache Optimizations for specified query.Vendor

		_, ok1 := cache.mappackages[query.Vendor]

		if ok1 == true {

			for _, identifiers := range cache.mappackages[query.Vendor] {

				for _, identifier := range identifiers {

					response, ok := cache.Responses[identifier]

					if ok == true {

						matches_package := false

						for _, pkg := range response.Packages {

							if query.MatchesName(pkg.Name) && query.MatchesVersion(pkg.Version) && query.MatchesArchitecture(pkg.Architecture) && query.MatchesManager(pkg.Manager) {
								matches_package = true
								break
							}

						}

						if matches_package {
							found[identifier] = response
						}

					}

				}

			}

		}

	} else {

		for identifier, response := range cache.Responses {

			matches_package := false

			for _, pkg := range response.Packages {

				if query.MatchesVendor(pkg.Vendor) && query.MatchesName(pkg.Name) && query.MatchesVersion(pkg.Version) && query.MatchesArchitecture(pkg.Architecture) && query.MatchesManager(pkg.Manager) {
					matches_package = true
					break
				}

			}

			if matches_package {
				found[identifier] = response
			}

		}

	}

	if len(found) > 0 {

		for _, response := range found {
			result = append(result, response)
		}

	}

	if len(result) > 1 {

		sort.Slice(result, func(a int, b int) bool {
			return result[a].IsAfter(*result[b])
		})

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Responses) QueryByProduct(query matchers.Product) []*structs.Response {

	cache.mutex.RLock()

	result := make([]*structs.Response, 0)
	found := make(map[string]*structs.Response)

	if query.Vendor != "any" && query.Product != "any" && query.Version != "any" {

		// Cache Optimizations for specified query.Vendor, query.Product, query.Version

		_, ok1 := cache.mapproducts[query.Vendor]

		if ok1 == true {

			_, ok2 := cache.mapproducts[query.Vendor][query.Product]

			if ok2 == true {

				identifiers, ok3 := cache.mapproducts[query.Vendor][query.Product][query.Version]

				if ok3 == true {

					for _, identifier := range identifiers {

						response, ok := cache.Responses[identifier]

						if ok == true {

							matches_product := false

							for _, product := range response.Products {

								if query.MatchesType(product.Type) {
									matches_product = true
									break
								}

							}

							if matches_product {
								found[identifier] = response
							}

						}

					}

				}

			}

		}

	} else if query.Vendor != "any" && query.Product != "any" && query.Version == "any" {

		// Cache Optimizations for specified query.Vendor, query.Product

		_, ok1 := cache.mapproducts[query.Vendor]

		if ok1 == true {

			_, ok2 := cache.mapproducts[query.Vendor][query.Product]

			if ok2 == true {

				for _, identifiers := range cache.mapproducts[query.Vendor][query.Product] {

					for _, identifier := range identifiers {

						response, ok := cache.Responses[identifier]

						if ok == true {

							matches_product := false

							for _, product := range response.Products {

								if query.MatchesType(product.Type) {
									matches_product = true
									break
								}

							}

							if matches_product {
								found[identifier] = response
							}

						}

					}

				}

			}

		}

	} else if query.Vendor != "any" && query.Product == "any" && query.Version == "any" {

		// Cache Optimizations for specified query.Vendor

		_, ok1 := cache.mapproducts[query.Vendor]

		if ok1 == true {

			for _, products := range cache.mapproducts[query.Vendor] {

				for _, identifiers := range products {

					for _, identifier := range identifiers {

						response, ok := cache.Responses[identifier]

						if ok == true {

							matches_product := false

							for _, product := range response.Products {

								if query.MatchesType(product.Type) {
									matches_product = true
									break
								}

							}

							if matches_product {
								found[identifier] = response
							}

						}

					}

				}

			}

		}

	} else {

		for identifier, response := range cache.Responses {

			matches_product := false

			for _, product := range response.Products {

				matches_vendor := query.MatchesVendor(product.Vendor)
				matches_product_name := query.MatchesProduct(product.Product)
				matches_version := query.MatchesVersion(product.Version)
				matches_type := query.MatchesType(product.Type)

				if matches_vendor && matches_product_name && matches_version && matches_type {
					matches_product = true
					break
				}

			}

			if matches_product {
				found[identifier] = response
			}

		}

	}

	if len(found) > 0 {

		for _, response := range found {
			result = append(result, response)
		}

	}

	if len(result) > 1 {

		sort.Slice(result, func(a int, b int) bool {
			return result[a].IsAfter(*result[b])
		})

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Responses) QueryByProgram(query matchers.Program) []*structs.Response {

	cache.mutex.RLock()

	result := make([]*structs.Response, 0)
	found := make(map[string]*structs.Response)

	if query.Name != "any" {

		// Cache Optimizations for specified query.Name

		_, ok1 := cache.mapprograms[query.Name]

		if ok1 == true {

			for _, identifier := range cache.mapprograms[query.Name] {

				response, ok := cache.Responses[identifier]

				if ok == true {

					matches_program := false

					for _, program := range response.Programs {

						if query.MatchesCommand(program.Command) {
							matches_program = true
							break
						}

					}

					if matches_program {
						found[identifier] = response
					}

				}

			}

		}

	} else {

		for identifier, response := range cache.Responses {

			matches_program := false

			for _, program := range response.Programs {

				if query.MatchesName(program.Name) && query.MatchesCommand(program.Command) {
					matches_program = true
					break
				}

			}

			if matches_program {
				found[identifier] = response
			}

		}

	}

	if len(found) > 0 {

		for _, response := range found {
			result = append(result, response)
		}

	}

	if len(result) > 1 {

		sort.Slice(result, func(a int, b int) bool {
			return result[a].IsAfter(*result[b])
		})

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Responses) QueryByUser(query matchers.User) []*structs.Response {

	cache.mutex.RLock()

	result := make([]*structs.Response, 0)
	found := make(map[string]*structs.Response)

	if query.Name != "any" {

		// Cache Optimizations for specified query.Name

		_, ok1 := cache.mapusers[query.Name]

		if ok1 == true {

			for _, identifier := range cache.mapusers[query.Name] {

				response, ok := cache.Responses[identifier]

				if ok == true {

					matches_user := false

					for _, user := range response.Users {

						if query.MatchesPassword(user.Password) && query.MatchesType(user.Type) {
							matches_user = true
							break
						}

					}

					if matches_user {
						found[identifier] = response
					}

				}

			}

		}

	} else {

		for identifier, response := range cache.Responses {

			matches_user := false

			for _, user := range response.Users {

				if query.MatchesName(user.Name) && query.MatchesPassword(user.Password) && query.MatchesType(user.Type) {
					matches_user = true
					break
				}

			}

			if matches_user {
				found[identifier] = response
			}

		}

	}

	if len(found) > 0 {

		for _, response := range found {
			result = append(result, response)
		}

	}

	if len(result) > 1 {

		sort.Slice(result, func(a int, b int) bool {
			return result[a].IsAfter(*result[b])
		})

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Responses) QueryByVulnerability(query matchers.Vulnerability) []*structs.Response {

	cache.mutex.RLock()

	result := make([]*structs.Response, 0)
	found := make(map[string]*structs.Response)

	if query.Name != "any" {

		// Cache Optimizations for specified query.Name

		_, ok1 := cache.mapvulnerabilities[query.Name]

		if ok1 == true {

			for _, identifier := range cache.mapvulnerabilities[query.Name] {

				response, ok := cache.Responses[identifier]

				if ok == true {
					found[identifier] = response
				}

			}

		}

	}

	if len(found) > 0 {

		for _, response := range found {
			result = append(result, response)
		}

	}

	if len(result) > 1 {

		sort.Slice(result, func(a int, b int) bool {
			return result[a].IsAfter(*result[b])
		})

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Responses) QueryByWeakness(query matchers.Weakness) []*structs.Response {

	cache.mutex.RLock()

	result := make([]*structs.Response, 0)
	found := make(map[string]*structs.Response)

	if query.Name != "any" && query.Impact != "any" && query.Scope != "any" {

		// Cache Optimizations for specified query.Name, query.Impact, query.Scope

		_, ok1 := cache.mapweaknesses[query.Name]

		if ok1 == true {

			_, ok2 := cache.mapweaknesses[query.Name][query.Impact]

			if ok2 == true {

				identifiers, ok3 := cache.mapweaknesses[query.Name][query.Impact][query.Scope]

				if ok3 == true {

					for _, identifier := range identifiers {

						response, ok := cache.Responses[identifier]

						if ok == true {
							found[identifier] = response
						}

					}

				}

			}

		}

	} else if query.Name != "any" && query.Impact != "any" && query.Scope == "any" {

		// Cache Optimizations for specified query.Name, query.Impact

		_, ok1 := cache.mapweaknesses[query.Name]

		if ok1 == true {

			_, ok2 := cache.mapweaknesses[query.Name][query.Impact]

			if ok2 == true {

				for _, identifiers := range cache.mapweaknesses[query.Name][query.Impact] {

					for _, identifier := range identifiers {

						response, ok := cache.Responses[identifier]

						if ok == true {
							found[identifier] = response
						}

					}

				}

			}

		}

	} else if query.Name != "any" && query.Impact == "any" && query.Scope == "any" {

		// Cache Optimizations for specified query.Name

		_, ok1 := cache.mapweaknesses[query.Name]

		if ok1 == true {

			for _, impacts := range cache.mapweaknesses[query.Name] {

				for _, identifiers := range impacts {

					for _, identifier := range identifiers {

						response, ok := cache.Responses[identifier]

						if ok == true {
							found[identifier] = response
						}

					}

				}

			}

		}

	} else {

		for identifier, response := range cache.Responses {

			matches_weakness := false

			for _, weakness := range response.Weaknesses {

				matches_name := query.MatchesName(weakness.Name)
				matches_impact := query.MatchesImpact(weakness.Impact)
				matches_scope := query.MatchesScope(weakness.Scope)

				if matches_name && matches_impact && matches_scope {
					matches_weakness = true
					break
				}

			}

			if matches_weakness {
				found[identifier] = response
			}

		}

	}

	if len(found) > 0 {

		for _, response := range found {
			result = append(result, response)
		}

	}

	if len(result) > 1 {

		sort.Slice(result, func(a int, b int) bool {
			return result[a].IsAfter(*result[b])
		})

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Responses) Remove(identifier string) bool {

	cache.mutex.Lock()

	var result bool

	if identifier != "" {

		response, ok := cache.Responses[identifier]

		if ok == true {

			delete(cache.Responses, identifier)
			result = true

			_, ok1 := cache.mapresponses[response.System][response.Type][response.Datetime.String()]

			if ok1 == true {

				delete(cache.mapresponses[response.System][response.Type], response.Datetime.String())

				if len(cache.mapresponses[response.System][response.Type]) == 0 {

					delete(cache.mapresponses[response.System], response.Type)

					if len(cache.mapresponses[response.System]) == 0 {
						delete(cache.mapresponses, response.System)
					}

				}

			}

			for _, connection := range response.Connections {

				types, ok1 := cache.mapconnections[connection.Type]

				if ok1 == true {

					protocols, ok2 := types[connection.Protocol]

					if ok2 == true {

						hosts, ok3 := protocols[connection.Host]

						if ok3 == true {

							others, ok4 := hosts[connection.Port]

							if ok4 == true {

								filtered := make([]string, 0)

								for _, other := range others {

									if other != identifier {
										filtered = append(filtered, other)
									}

								}

								if len(filtered) > 0 {
									hosts[connection.Port] = filtered
								} else {

									delete(hosts, connection.Port)

									if len(hosts) == 0 {

										delete(protocols, connection.Host)

										if len(protocols) == 0 {

											delete(types, connection.Protocol)

											if len(types) == 0 {
												delete(cache.mapconnections, connection.Type)
											}

										}

									}

								}

							}

						}

					}

				}

			}

			for _, distribution := range response.Distributions {

				others, ok1 := cache.mapdistributions[distribution.Name]

				if ok1 == true {

					filtered := make([]string, 0)

					for _, other := range others {

						if other != identifier {
							filtered = append(filtered, other)
						}

					}

					if len(filtered) > 0 {
						cache.mapdistributions[distribution.Name] = filtered
					} else {
						delete(cache.mapdistributions, distribution.Name)
					}

				}

			}

			for _, network := range response.Networks {

				if network.Name != "" && strings.HasPrefix(network.Name, "AS") {

					others, ok1 := cache.mapasns[network.Name]

					if ok1 == true {

						filtered := make([]string, 0)

						for _, other := range others {

							if other != identifier {
								filtered = append(filtered, other)
							}

						}

						if len(filtered) > 0 {
							cache.mapasns[network.Name] = filtered
						} else {
							delete(cache.mapasns, network.Name)
						}

					}

				}

			}

			for _, pkg := range response.Packages {

				others, ok1 := cache.mappackages[pkg.Vendor][pkg.Name]

				if ok1 == true {

					filtered := make([]string, 0)

					for _, other := range others {

						if other != identifier {
							filtered = append(filtered, other)
						}

					}

					if len(filtered) > 0 {
						cache.mappackages[pkg.Vendor][pkg.Name] = filtered
					} else {

						delete(cache.mappackages[pkg.Vendor], pkg.Name)

						if len(cache.mappackages[pkg.Vendor]) == 0 {
							delete(cache.mappackages, pkg.Vendor)
						}

					}

				}

			}

			for _, product := range response.Products {

				versions, ok1 := cache.mapproducts[product.Vendor]

				if ok1 == true {

					products, ok2 := versions[product.Product]

					if ok2 == true {

						others, ok3 := products[product.Version]

						if ok3 == true {

							filtered := make([]string, 0)

							for _, other := range others {

								if other != identifier {
									filtered = append(filtered, other)
								}

							}

							if len(filtered) > 0 {
								products[product.Version] = filtered
							} else {

								delete(products, product.Version)

								if len(products) == 0 {

									delete(versions, product.Product)

									if len(versions) == 0 {
										delete(cache.mapproducts, product.Vendor)
									}

								}

							}

						}

					}

				}

			}

			for _, program := range response.Programs {

				others, ok1 := cache.mapprograms[program.Name]

				if ok1 == true {

					filtered := make([]string, 0)

					for _, other := range others {

						if other != identifier {
							filtered = append(filtered, other)
						}

					}

					if len(filtered) > 0 {
						cache.mapprograms[program.Name] = filtered
					} else {
						delete(cache.mapprograms, program.Name)
					}

				}

			}

			for _, user := range response.Users {

				others, ok1 := cache.mapusers[user.Name]

				if ok1 == true {

					filtered := make([]string, 0)

					for _, other := range others {

						if other != identifier {
							filtered = append(filtered, other)
						}

					}

					if len(filtered) > 0 {
						cache.mapusers[user.Name] = filtered
					} else {
						delete(cache.mapusers, user.Name)
					}

				}

			}

			for _, vulnerability := range response.Vulnerabilities {

				others, ok1 := cache.mapvulnerabilities[vulnerability.Name]

				if ok1 == true {

					filtered := make([]string, 0)

					for _, other := range others {

						if other != identifier {
							filtered = append(filtered, other)
						}

					}

					if len(filtered) > 0 {
						cache.mapvulnerabilities[vulnerability.Name] = filtered
					} else {
						delete(cache.mapvulnerabilities, vulnerability.Name)
					}

				}

			}

			for _, weakness := range response.Weaknesses {

				impacts, ok1 := cache.mapweaknesses[weakness.Name]

				if ok1 == true {

					scopes, ok2 := impacts[weakness.Impact]

					if ok2 == true {

						others, ok3 := scopes[weakness.Scope]

						if ok3 == true {

							filtered := make([]string, 0)

							for _, other := range others {

								if other != identifier {
									filtered = append(filtered, other)
								}

							}

							if len(filtered) > 0 {
								scopes[weakness.Scope] = filtered
							} else {

								delete(scopes, weakness.Scope)

								if len(scopes) == 0 {

									delete(impacts, weakness.Impact)

									if len(impacts) == 0 {
										delete(cache.mapweaknesses, weakness.Name)
									}

								}

							}

						}

					}

				}

			}

		}

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Responses) MarshalJSON() ([]byte, error) {

	tmp := make(map[string]*structs.Response)
	identifiers := make([]string, 0)

	for identifier, _ := range cache.Responses {
		identifiers = append(identifiers, identifier)
	}

	sort.Strings(identifiers)

	for _, identifier := range identifiers {
		tmp[identifier] = cache.Responses[identifier]
	}

	return json.MarshalIndent(&struct {
		Responses map[string]*structs.Response `json:"responses"`
	}{
		Responses: tmp,
	}, "", "\t")

}

func (cache *Responses) UnmarshalJSON(data []byte) error {

	var tmp struct {
		Responses map[string]*structs.Response `json:"responses"`
	}

	err := json.Unmarshal(data, &tmp)

	if err != nil {
		return err
	}

	for _, response := range tmp.Responses {
		cache.Add(*response)
	}

	return nil

}
