package caches

import "tholian-core/matchers"
import "tholian-core/structs"
import "encoding/json"
import "sort"
import "strings"
import "sync"

func toIncidentIdentifier(value structs.Incident) string {
	return value.System + ":" + value.Type + ":" + value.Datetime.String()
}

type Incidents struct {
	Incidents          map[string]*structs.Incident                         `json:"incidents"`
	mapasns            map[string][]string                                  `json:"-"`
	mapconnections     map[string]map[string]map[string]map[uint16][]string `json:"-"`
	mapdistributions   map[string][]string                                  `json:"-"`
	mapincidents       map[string]map[string]map[string]string              `json:"-"`
	mappackages        map[string]map[string][]string                       `json:"-"`
	mapproducts        map[string]map[string]map[string][]string            `json:"-"`
	mapprograms        map[string][]string                                  `json:"-"`
	mapusers           map[string][]string                                  `json:"-"`
	mapvulnerabilities map[string][]string                                  `json:"-"`
	mapweaknesses      map[string]map[string]map[string][]string            `json:"-"`
	networks           *Networks                                            `json:"-"`
	mutex              *sync.RWMutex                                        `json:"-"`
}

func NewIncidents() *Incidents {
	return NewIncidentsFromNetworks(NewNetworks())
}

func NewIncidentsFromNetworks(networks *Networks) *Incidents {

	var cache Incidents

	cache.Incidents = make(map[string]*structs.Incident)
	cache.mapasns = make(map[string][]string, 0)
	cache.mapconnections = make(map[string]map[string]map[string]map[uint16][]string, 0)
	cache.mapdistributions = make(map[string][]string, 0)
	cache.mapincidents = make(map[string]map[string]map[string]string)
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

func (cache *Incidents) Add(value structs.Incident) bool {

	cache.mutex.Lock()

	var result bool

	if value.System != "" && value.Type != "" && value.Datetime.IsValid() {

		identifier := toIncidentIdentifier(value)

		cache.Incidents[identifier] = &value
		result = true

		_, ok1 := cache.mapincidents[value.System]

		if ok1 == false {
			cache.mapincidents[value.System] = make(map[string]map[string]string)
		}

		_, ok2 := cache.mapincidents[value.System][value.Type]

		if ok2 == false {
			cache.mapincidents[value.System][value.Type] = make(map[string]string)
		}

		cache.mapincidents[value.System][value.Type][value.Datetime.String()] = identifier

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

func (cache *Incidents) Get(identifier string) *structs.Incident {

	cache.mutex.RLock()

	var result *structs.Incident = nil

	if strings.Contains(identifier, ":") {

		tmp := strings.Split(identifier, ":")

		// System:Type:YYYY-MM-DD HH:ii:SS
		if len(tmp) == 5 {

			incident, ok := cache.Incidents[identifier]

			if ok == true {
				result = incident
			}

		}

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Incidents) Length() int {
	return len(cache.Incidents)
}

func (cache *Incidents) Query(query matchers.Incident) []*structs.Incident {

	cache.mutex.RLock()

	result := make([]*structs.Incident, 0)
	found := make(map[string]*structs.Incident)

	if query.System != "any" && query.Type != "any" {

		// Cache Optimizations for specified query.System, query.Type

		_, ok1 := cache.mapincidents[query.System]

		if ok1 == true {

			datetimes, ok2 := cache.mapincidents[query.System][query.Type]

			if ok2 == true {

				for _, identifier := range datetimes {

					incident, ok := cache.Incidents[identifier]

					if ok == true {

						matches_datetime := query.MatchesDatetime(incident.Datetime.String())

						if matches_datetime {
							found[identifier] = incident
						}

					}

				}

			}

		}

	} else if query.System != "any" && query.Type == "any" {

		// Cache Optimizations for specified query.System

		_, ok1 := cache.mapincidents[query.System]

		if ok1 == true {

			for _, types := range cache.mapincidents[query.System] {

				for _, identifier := range types {

					incident, ok := cache.Incidents[identifier]

					if ok == true {

						matches_datetime := query.MatchesDatetime(incident.Datetime.String())

						if matches_datetime {
							found[identifier] = incident
						}

					}

				}

			}

		}

	} else {

		for identifier, incident := range cache.Incidents {

			matches_system := query.MatchesSystem(incident.System)
			matches_type := query.MatchesType(incident.Type)
			matches_datetime := query.MatchesDatetime(incident.Datetime.String())

			if matches_system && matches_type && matches_datetime {
				found[identifier] = incident
			}

		}

	}

	if len(found) > 0 {

		for _, incident := range found {
			result = append(result, incident)
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

func (cache *Incidents) QueryByConnection(query matchers.Connection) []*structs.Incident {

	cache.mutex.RLock()

	result := make([]*structs.Incident, 0)
	found := make(map[string]*structs.Incident)

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

							incident, ok := cache.Incidents[identifier]

							if ok == true {

								matches_connection := false

								for _, connection := range incident.Connections {

									if query.Matches(connection.Host, connection.Port, connection.Protocol, connection.Type) {
										matches_connection = true
										break
									}

								}

								if matches_connection {
									found[identifier] = incident
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

							incident, ok := cache.Incidents[identifier]

							if ok == true {

								matches_connection := false

								for _, connection := range incident.Connections {

									if query.Matches(connection.Host, connection.Port, connection.Protocol, connection.Type) {
										matches_connection = true
										break
									}

								}

								if matches_connection {
									found[identifier] = incident
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

							incident, ok := cache.Incidents[identifier]

							if ok == true {

								matches_connection := false

								for _, connection := range incident.Connections {

									if query.Matches(connection.Host, connection.Port, connection.Protocol, connection.Type) {
										matches_connection = true
										break
									}

								}

								if matches_connection {
									found[identifier] = incident
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

							incident, ok := cache.Incidents[identifier]

							if ok == true {

								matches_connection := false

								for _, connection := range incident.Connections {

									if query.Matches(connection.Host, connection.Port, connection.Protocol, connection.Type) {
										matches_connection = true
										break
									}

								}

								if matches_connection {
									found[identifier] = incident
								}

							}

						}

					}

				}

			}

		}

	} else {

		for identifier, incident := range cache.Incidents {

			matches_connection := false

			for _, connection := range incident.Connections {

				if query.Matches(connection.Host, connection.Port, connection.Protocol, connection.Type) {
					matches_connection = true
					break
				}

			}

			if matches_connection {
				found[identifier] = incident
			}

		}

	}

	if len(found) > 0 {

		for _, incident := range found {
			result = append(result, incident)
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

func (cache *Incidents) QueryByDistribution(query matchers.Distribution) []*structs.Incident {

	cache.mutex.RLock()

	result := make([]*structs.Incident, 0)
	found := make(map[string]*structs.Incident)

	if query.Name != "any" {

		// Cache Optimizations for specified query.Name

		identifiers, ok1 := cache.mapdistributions[query.Name]

		if ok1 == true {

			for _, identifier := range identifiers {

				incident, ok := cache.Incidents[identifier]

				if ok == true {

					matches_distribution := false

					for _, distribution := range incident.Distributions {

						if query.MatchesVersion(distribution.Version) && query.MatchesManager(distribution.Manager) && query.MatchesVendor(distribution.Vendor) {
							matches_distribution = true
							break
						}

					}

					if matches_distribution {
						found[identifier] = incident
					}

				}

			}

		}

	} else {

		for identifier, incident := range cache.Incidents {

			matches_distribution := false

			for _, distribution := range incident.Distributions {

				if query.MatchesName(distribution.Name) && query.MatchesVersion(distribution.Version) && query.MatchesManager(distribution.Manager) && query.MatchesVendor(distribution.Vendor) {
					matches_distribution = true
					break
				}

			}

			if matches_distribution {
				found[identifier] = incident
			}

		}

	}

	if len(found) > 0 {

		for _, incident := range found {
			result = append(result, incident)
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

func (cache *Incidents) QueryByIP(query string) []*structs.Incident {

	cache.mutex.RLock()

	result := make([]*structs.Incident, 0)
	found := make(map[string]*structs.Incident)

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

						incident, ok := cache.Incidents[identifier]

						if ok == true {
							found[identifier] = incident
						}

					}

				}

			}

		}

	}

	if len(found) > 0 {

		for _, incident := range found {
			result = append(result, incident)
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

func (cache *Incidents) QueryByPackage(query matchers.Package) []*structs.Incident {

	cache.mutex.RLock()

	result := make([]*structs.Incident, 0)
	found := make(map[string]*structs.Incident)

	if query.Vendor != "any" && query.Name != "any" {

		// Cache Optimizations for specified query.Vendor, query.Name

		_, ok1 := cache.mappackages[query.Vendor]

		if ok1 == true {

			identifiers, ok2 := cache.mappackages[query.Vendor][query.Name]

			if ok2 == true {

				for _, identifier := range identifiers {

					incident, ok := cache.Incidents[identifier]

					if ok == true {

						matches_package := false

						for _, pkg := range incident.Packages {

							if query.MatchesVersion(pkg.Version) && query.MatchesArchitecture(pkg.Architecture) && query.MatchesManager(pkg.Manager) {
								matches_package = true
								break
							}

						}

						if matches_package {
							found[identifier] = incident
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

					incident, ok := cache.Incidents[identifier]

					if ok == true {

						matches_package := false

						for _, pkg := range incident.Packages {

							if query.MatchesName(pkg.Name) && query.MatchesVersion(pkg.Version) && query.MatchesArchitecture(pkg.Architecture) && query.MatchesManager(pkg.Manager) {
								matches_package = true
								break
							}

						}

						if matches_package {
							found[identifier] = incident
						}

					}

				}

			}

		}

	} else {

		for identifier, incident := range cache.Incidents {

			matches_package := false

			for _, pkg := range incident.Packages {

				if query.MatchesVendor(pkg.Vendor) && query.MatchesName(pkg.Name) && query.MatchesVersion(pkg.Version) && query.MatchesArchitecture(pkg.Architecture) && query.MatchesManager(pkg.Manager) {
					matches_package = true
					break
				}

			}

			if matches_package {
				found[identifier] = incident
			}

		}

	}

	if len(found) > 0 {

		for _, incident := range found {
			result = append(result, incident)
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

func (cache *Incidents) QueryByProduct(query matchers.Product) []*structs.Incident {

	cache.mutex.RLock()

	result := make([]*structs.Incident, 0)
	found := make(map[string]*structs.Incident)

	if query.Vendor != "any" && query.Product != "any" && query.Version != "any" {

		// Cache Optimizations for specified query.Vendor, query.Product, query.Version

		_, ok1 := cache.mapproducts[query.Vendor]

		if ok1 == true {

			_, ok2 := cache.mapproducts[query.Vendor][query.Product]

			if ok2 == true {

				identifiers, ok3 := cache.mapproducts[query.Vendor][query.Product][query.Version]

				if ok3 == true {

					for _, identifier := range identifiers {

						incident, ok := cache.Incidents[identifier]

						if ok == true {

							matches_product := false

							for _, product := range incident.Products {

								if query.MatchesType(product.Type) {
									matches_product = true
									break
								}

							}

							if matches_product {
								found[identifier] = incident
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

						incident, ok := cache.Incidents[identifier]

						if ok == true {

							matches_product := false

							for _, product := range incident.Products {

								if query.MatchesType(product.Type) {
									matches_product = true
									break
								}

							}

							if matches_product {
								found[identifier] = incident
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

						incident, ok := cache.Incidents[identifier]

						if ok == true {

							matches_product := false

							for _, product := range incident.Products {

								if query.MatchesType(product.Type) {
									matches_product = true
									break
								}

							}

							if matches_product {
								found[identifier] = incident
							}

						}

					}

				}

			}

		}

	} else {

		for identifier, incident := range cache.Incidents {

			matches_product := false

			for _, product := range incident.Products {

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
				found[identifier] = incident
			}

		}

	}

	if len(found) > 0 {

		for _, incident := range found {
			result = append(result, incident)
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

func (cache *Incidents) QueryByProgram(query matchers.Program) []*structs.Incident {

	cache.mutex.RLock()

	result := make([]*structs.Incident, 0)
	found := make(map[string]*structs.Incident)

	if query.Name != "any" {

		// Cache Optimizations for specified query.Name

		_, ok1 := cache.mapprograms[query.Name]

		if ok1 == true {

			for _, identifier := range cache.mapprograms[query.Name] {

				incident, ok := cache.Incidents[identifier]

				if ok == true {

					matches_program := false

					for _, program := range incident.Programs {

						if query.MatchesCommand(program.Command) {
							matches_program = true
							break
						}

					}

					if matches_program {
						found[identifier] = incident
					}

				}

			}

		}

	} else {

		for identifier, incident := range cache.Incidents {

			matches_program := false

			for _, program := range incident.Programs {

				if query.MatchesName(program.Name) && query.MatchesCommand(program.Command) {
					matches_program = true
					break
				}

			}

			if matches_program {
				found[identifier] = incident
			}

		}

	}

	if len(found) > 0 {

		for _, incident := range found {
			result = append(result, incident)
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

func (cache *Incidents) QueryByUser(query matchers.User) []*structs.Incident {

	cache.mutex.RLock()

	result := make([]*structs.Incident, 0)
	found := make(map[string]*structs.Incident)

	if query.Name != "any" {

		// Cache Optimizations for specified query.Name

		_, ok1 := cache.mapusers[query.Name]

		if ok1 == true {

			for _, identifier := range cache.mapusers[query.Name] {

				incident, ok := cache.Incidents[identifier]

				if ok == true {

					matches_user := false

					for _, user := range incident.Users {

						if query.MatchesPassword(user.Password) && query.MatchesType(user.Type) {
							matches_user = true
							break
						}

					}

					if matches_user {
						found[identifier] = incident
					}

				}

			}

		}

	} else {

		for identifier, incident := range cache.Incidents {

			matches_user := false

			for _, user := range incident.Users {

				if query.MatchesName(user.Name) && query.MatchesPassword(user.Password) && query.MatchesType(user.Type) {
					matches_user = true
					break
				}

			}

			if matches_user {
				found[identifier] = incident
			}

		}

	}

	if len(found) > 0 {

		for _, incident := range found {
			result = append(result, incident)
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

func (cache *Incidents) QueryByVulnerability(query matchers.Vulnerability) []*structs.Incident {

	cache.mutex.RLock()

	result := make([]*structs.Incident, 0)
	found := make(map[string]*structs.Incident)

	if query.Name != "any" {

		// Cache Optimizations for specified query.Name

		_, ok1 := cache.mapvulnerabilities[query.Name]

		if ok1 == true {

			for _, identifier := range cache.mapvulnerabilities[query.Name] {

				incident, ok := cache.Incidents[identifier]

				if ok == true {
					found[identifier] = incident
				}

			}

		}

	}

	if len(found) > 0 {

		for _, incident := range found {
			result = append(result, incident)
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

func (cache *Incidents) QueryByWeakness(query matchers.Weakness) []*structs.Incident {

	cache.mutex.RLock()

	result := make([]*structs.Incident, 0)
	found := make(map[string]*structs.Incident)

	if query.Name != "any" && query.Impact != "any" && query.Scope != "any" {

		// Cache Optimizations for specified query.Name, query.Impact, query.Scope

		_, ok1 := cache.mapweaknesses[query.Name]

		if ok1 == true {

			_, ok2 := cache.mapweaknesses[query.Name][query.Impact]

			if ok2 == true {

				identifiers, ok3 := cache.mapweaknesses[query.Name][query.Impact][query.Scope]

				if ok3 == true {

					for _, identifier := range identifiers {

						incident, ok := cache.Incidents[identifier]

						if ok == true {
							found[identifier] = incident
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

						incident, ok := cache.Incidents[identifier]

						if ok == true {
							found[identifier] = incident
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

						incident, ok := cache.Incidents[identifier]

						if ok == true {
							found[identifier] = incident
						}

					}

				}

			}

		}

	} else {

		for identifier, incident := range cache.Incidents {

			matches_weakness := false

			for _, weakness := range incident.Weaknesses {

				matches_name := query.MatchesName(weakness.Name)
				matches_impact := query.MatchesImpact(weakness.Impact)
				matches_scope := query.MatchesScope(weakness.Scope)

				if matches_name && matches_impact && matches_scope {
					matches_weakness = true
					break
				}

			}

			if matches_weakness {
				found[identifier] = incident
			}

		}

	}

	if len(found) > 0 {

		for _, incident := range found {
			result = append(result, incident)
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

func (cache *Incidents) Remove(identifier string) bool {

	cache.mutex.Lock()

	var result bool

	if identifier != "" {

		incident, ok := cache.Incidents[identifier]

		if ok == true {

			delete(cache.Incidents, identifier)
			result = true

			_, ok1 := cache.mapincidents[incident.System][incident.Type][incident.Datetime.String()]

			if ok1 == true {

				delete(cache.mapincidents[incident.System][incident.Type], incident.Datetime.String())

				if len(cache.mapincidents[incident.System][incident.Type]) == 0 {

					delete(cache.mapincidents[incident.System], incident.Type)

					if len(cache.mapincidents[incident.System]) == 0 {
						delete(cache.mapincidents, incident.System)
					}

				}

			}

			for _, connection := range incident.Connections {

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

			for _, distribution := range incident.Distributions {

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

			for _, network := range incident.Networks {

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

			for _, pkg := range incident.Packages {

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

			for _, product := range incident.Products {

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

			for _, program := range incident.Programs {

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

			for _, user := range incident.Users {

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

			for _, vulnerability := range incident.Vulnerabilities {

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

			for _, weakness := range incident.Weaknesses {

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

func (cache *Incidents) MarshalJSON() ([]byte, error) {

	tmp := make(map[string]*structs.Incident)
	identifiers := make([]string, 0)

	for identifier, _ := range cache.Incidents {
		identifiers = append(identifiers, identifier)
	}

	sort.Strings(identifiers)

	for _, identifier := range identifiers {
		tmp[identifier] = cache.Incidents[identifier]
	}

	return json.MarshalIndent(&struct {
		Incidents map[string]*structs.Incident `json:"incidents"`
	}{
		Incidents: tmp,
	}, "", "\t")

}

func (cache *Incidents) UnmarshalJSON(data []byte) error {

	var tmp struct {
		Incidents map[string]*structs.Incident `json:"incidents"`
	}

	err := json.Unmarshal(data, &tmp)

	if err != nil {
		return err
	}

	for _, incident := range tmp.Incidents {
		cache.Add(*incident)
	}

	return nil

}
