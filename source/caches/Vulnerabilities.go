package caches

import "tholian-core/matchers"
import "tholian-core/structs"
import "encoding/json"
import "slices"
import "sort"
import "sync"

type Vulnerabilities struct {
	Vulnerabilities    map[string]*structs.Vulnerability                    `json:"vulnerabilities"`
	mapconnections     map[string]map[string]map[string]map[uint16][]string `json:"-"`
	mapdevices         map[string]map[string]map[string][]string            `json:"-"`
	mapdistributions   map[string]map[string][]string                       `json:"-"`
	mappackages        map[string]map[string][]string                       `json:"-"`
	mapproducts        map[string]map[string]map[string][]string            `json:"-"`
	mapprograms        map[string][]string                                  `json:"-"`
	mapvulnerabilities map[string]map[string][]string                       `json:"-"`
	mutex              *sync.RWMutex                                        `json:"-"`
}

func NewVulnerabilities() *Vulnerabilities {

	var cache Vulnerabilities

	cache.Vulnerabilities = make(map[string]*structs.Vulnerability)
	cache.mapconnections = make(map[string]map[string]map[string]map[uint16][]string, 0)
	cache.mapdevices = make(map[string]map[string]map[string][]string, 0)
	cache.mapdistributions = make(map[string]map[string][]string, 0)
	cache.mappackages = make(map[string]map[string][]string, 0)
	cache.mapproducts = make(map[string]map[string]map[string][]string, 0)
	cache.mapprograms = make(map[string][]string, 0)
	cache.mapvulnerabilities = make(map[string]map[string][]string, 0)
	cache.mutex = &sync.RWMutex{}

	return &cache

}

func (cache *Vulnerabilities) Add(value structs.Vulnerability) bool {

	cache.mutex.Lock()

	var result bool

	if value.Name != "" {

		cache.Vulnerabilities[value.Name] = &value

		_, ok1 := cache.mapvulnerabilities[value.State]

		if ok1 == false {
			cache.mapvulnerabilities[value.State] = make(map[string][]string, 0)
		}

		_, ok2 := cache.mapvulnerabilities[value.State][value.Type]

		if ok2 == false {
			cache.mapvulnerabilities[value.State][value.Type] = make([]string, 0)
		}

		cache.mapvulnerabilities[value.State][value.Type] = append(cache.mapvulnerabilities[value.State][value.Type], value.Name)
		result = true

		if len(value.Connections) > 0 {

			for _, connection := range value.Connections {

				_, ok1 := cache.mapconnections[connection.Type]

				if ok1 == false {
					cache.mapconnections[connection.Type] = make(map[string]map[string]map[uint16][]string, 0)
				}

				_, ok2 := cache.mapconnections[connection.Type][connection.Protocol]

				if ok2 == false {
					cache.mapconnections[connection.Type][connection.Protocol] = make(map[string]map[uint16][]string, 0)
				}

				_, ok3 := cache.mapconnections[connection.Type][connection.Protocol][connection.Host]

				if ok3 == false {
					cache.mapconnections[connection.Type][connection.Protocol][connection.Host] = make(map[uint16][]string, 0)
				}

				_, ok4 := cache.mapconnections[connection.Type][connection.Protocol][connection.Host][connection.Port]

				if ok4 == false {
					cache.mapconnections[connection.Type][connection.Protocol][connection.Host][connection.Port] = make([]string, 0)
				}

				if !slices.Contains(cache.mapconnections[connection.Type][connection.Protocol][connection.Host][connection.Port], value.Name) {
					cache.mapconnections[connection.Type][connection.Protocol][connection.Host][connection.Port] = append(cache.mapconnections[connection.Type][connection.Protocol][connection.Host][connection.Port], value.Name)
				}

			}

		}

		if len(value.Devices) > 0 {

			for _, device := range value.Devices {

				if device.Bus != "any" && device.Vendor != "any" && device.Device != "any" {

					_, ok1 := cache.mapdevices[device.Bus]

					if ok1 == false {
						cache.mapdevices[device.Bus] = make(map[string]map[string][]string, 0)
					}

					_, ok2 := cache.mapdevices[device.Bus][device.Vendor]

					if ok2 == false {
						cache.mapdevices[device.Bus][device.Vendor] = make(map[string][]string, 0)
					}

					_, ok3 := cache.mapdevices[device.Bus][device.Vendor][device.Device]

					if ok3 == false {
						cache.mapdevices[device.Bus][device.Vendor][device.Device] = make([]string, 0)
					}

					cache.mapdevices[device.Bus][device.Vendor][device.Device] = append(cache.mapdevices[device.Bus][device.Vendor][device.Device], value.Name)

				}

			}

		}

		if len(value.Distributions) > 0 {

			for _, distribution := range value.Distributions {

				if distribution.Name != "any" && distribution.Version != "any" {

					_, ok1 := cache.mapdistributions[distribution.Name]

					if ok1 == false {
						cache.mapdistributions[distribution.Name] = make(map[string][]string, 0)
					}

					_, ok2 := cache.mapdistributions[distribution.Name][distribution.Version]

					if ok2 == false {
						cache.mapdistributions[distribution.Name][distribution.Version] = make([]string, 0)
					}

					if !slices.Contains(cache.mapdistributions[distribution.Name][distribution.Version], value.Name) {
						cache.mapdistributions[distribution.Name][distribution.Version] = append(cache.mapdistributions[distribution.Name][distribution.Version], value.Name)
					}

				}

			}

		}

		if len(value.Packages) > 0 {

			for _, pkg := range value.Packages {

				if pkg.Name != "any" && pkg.Vendor != "any" {

					_, ok1 := cache.mappackages[pkg.Vendor]

					if ok1 == false {
						cache.mappackages[pkg.Vendor] = make(map[string][]string, 0)
					}

					_, ok2 := cache.mappackages[pkg.Vendor][pkg.Name]

					if ok2 == false {
						cache.mappackages[pkg.Vendor][pkg.Name] = make([]string, 0)
					}

					if !slices.Contains(cache.mappackages[pkg.Vendor][pkg.Name], value.Name) {
						cache.mappackages[pkg.Vendor][pkg.Name] = append(cache.mappackages[pkg.Vendor][pkg.Name], value.Name)
					}

				}

			}

		}

		if len(value.Products) > 0 {

			for _, product := range value.Products {

				_, ok1 := cache.mapproducts[product.Vendor]

				if ok1 == false {
					cache.mapproducts[product.Vendor] = make(map[string]map[string][]string, 0)
				}

				_, ok2 := cache.mapproducts[product.Vendor][product.Product]

				if ok2 == false {
					cache.mapproducts[product.Vendor][product.Product] = make(map[string][]string, 0)
				}

				_, ok3 := cache.mapproducts[product.Vendor][product.Product][product.Version]

				if ok3 == false {
					cache.mapproducts[product.Vendor][product.Product][product.Version] = make([]string, 0)
				}

				if !slices.Contains(cache.mapproducts[product.Vendor][product.Product][product.Version], value.Name) {
					cache.mapproducts[product.Vendor][product.Product][product.Version] = append(cache.mapproducts[product.Vendor][product.Product][product.Version], value.Name)
				}

			}

		}

		if len(value.Programs) > 0 {

			for _, program := range value.Programs {

				if program.Name != "any" {

					if !slices.Contains(cache.mapprograms[program.Name], value.Name) {
						cache.mapprograms[program.Name] = append(cache.mapprograms[program.Name], value.Name)
					}

				}

			}

		}

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Vulnerabilities) Get(name string) *structs.Vulnerability {

	cache.mutex.RLock()

	var result *structs.Vulnerability = nil

	if name != "" {

		vulnerability, ok := cache.Vulnerabilities[name]

		if ok == true {
			result = vulnerability
		}

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Vulnerabilities) Length() int {
	return len(cache.Vulnerabilities)
}

func (cache *Vulnerabilities) Query(query matchers.Vulnerability) []*structs.Vulnerability {

	cache.mutex.RLock()

	result := make([]*structs.Vulnerability, 0)
	found := make(map[string]*structs.Vulnerability)

	if query.Name != "any" {

		// Cache Optimizations for specified query.Name

		vulnerability, ok1 := cache.Vulnerabilities[query.Name]

		if ok1 == true {

			matches_type := query.MatchesType(vulnerability.Type)
			matches_severity := query.MatchesSeverity(vulnerability.Severity)
			matches_state := query.MatchesState(vulnerability.State)
			matches_datetime := query.MatchesDatetime(vulnerability.Datetime.String())

			if matches_type && matches_severity && matches_state && matches_datetime {
				found[vulnerability.Name] = vulnerability
			}

		}

	} else if query.Name == "any" && query.State != "any" && query.Type != "any" {

		// Cache Optimizations for specified query.State, query.Type

		_, ok1 := cache.mapvulnerabilities[query.State]

		if ok1 == true {

			identifiers, ok2 := cache.mapvulnerabilities[query.State][query.Type]

			if ok2 == true {

				for _, identifier := range identifiers {

					vulnerability, ok := cache.Vulnerabilities[identifier]

					if ok == true {

						matches_severity := query.MatchesSeverity(vulnerability.Severity)
						matches_datetime := query.MatchesDatetime(vulnerability.Datetime.String())

						if matches_severity && matches_datetime {
							found[vulnerability.Name] = vulnerability
						}

					}

				}

			}

		}

	} else if query.Name == "any" && query.State != "any" && query.Type == "any" {

		// Cache Optimizations for specified query.State

		_, ok1 := cache.mapvulnerabilities[query.State]

		if ok1 == true {

			for _, identifiers := range cache.mapvulnerabilities[query.State] {

				for _, identifier := range identifiers {

					vulnerability, ok := cache.Vulnerabilities[identifier]

					if ok == true {

						matches_severity := query.MatchesSeverity(vulnerability.Severity)
						matches_state := query.MatchesState(vulnerability.State)
						matches_datetime := query.MatchesDatetime(vulnerability.Datetime.String())

						if matches_severity && matches_state && matches_datetime {
							found[vulnerability.Name] = vulnerability
						}

					}

				}

			}

		}

	} else {

		for _, vulnerability := range cache.Vulnerabilities {

			matches_name := query.MatchesName(vulnerability.Name)
			matches_type := query.MatchesType(vulnerability.Type)
			matches_severity := query.MatchesSeverity(vulnerability.Severity)
			matches_state := query.MatchesState(vulnerability.State)
			matches_datetime := query.MatchesDatetime(vulnerability.Datetime.String())

			if matches_name && matches_type && matches_severity && matches_state && matches_datetime {
				found[vulnerability.Name] = vulnerability
			}

		}

	}

	if len(found) > 0 {

		for _, vulnerability := range found {
			result = append(result, vulnerability)
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

func (cache *Vulnerabilities) QueryByConnection(query matchers.Connection) []*structs.Vulnerability {

	cache.mutex.RLock()

	result := make([]*structs.Vulnerability, 0)
	found := make(map[string]*structs.Vulnerability)

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

							vulnerability, ok := cache.Vulnerabilities[identifier]

							if ok == true {

								matches_connection := false

								for _, connection := range vulnerability.Connections {

									if query.Matches(connection.Host, connection.Port, connection.Protocol, connection.Type) {
										matches_connection = true
										break
									}

								}

								if matches_connection {
									found[vulnerability.Name] = vulnerability
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

							vulnerability, ok := cache.Vulnerabilities[identifier]

							if ok == true {

								matches_connection := false

								for _, connection := range vulnerability.Connections {

									if query.Matches(connection.Host, connection.Port, connection.Protocol, connection.Type) {
										matches_connection = true
										break
									}

								}

								if matches_connection {
									found[vulnerability.Name] = vulnerability
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

							vulnerability, ok := cache.Vulnerabilities[identifier]

							if ok == true {

								matches_connection := false

								for _, connection := range vulnerability.Connections {

									if query.Matches(connection.Host, connection.Port, connection.Protocol, connection.Type) {
										matches_connection = true
										break
									}

								}

								if matches_connection {
									found[vulnerability.Name] = vulnerability
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

							vulnerability, ok := cache.Vulnerabilities[identifier]

							if ok == true {

								matches_connection := false

								for _, connection := range vulnerability.Connections {

									if query.Matches(connection.Host, connection.Port, connection.Protocol, connection.Type) {
										matches_connection = true
										break
									}

								}

								if matches_connection {
									found[vulnerability.Name] = vulnerability
								}

							}

						}

					}

				}

			}

		}

	} else {

		for _, vulnerability := range cache.Vulnerabilities {

			matches_connection := false

			for _, connection := range vulnerability.Connections {

				if query.Matches(connection.Host, connection.Port, connection.Protocol, connection.Type) {
					matches_connection = true
					break
				}

			}

			if matches_connection {
				found[vulnerability.Name] = vulnerability
			}

		}

	}

	if len(found) > 0 {

		for _, vulnerability := range found {
			result = append(result, vulnerability)
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

func (cache *Vulnerabilities) QueryByDevice(query matchers.Device) []*structs.Vulnerability {

	cache.mutex.RLock()

	result := make([]*structs.Vulnerability, 0)
	found := make(map[string]*structs.Vulnerability)

	if query.Bus != "any" && query.Vendor != "any" && query.Device != "any" {

		// Cache Optimizations for specified query.Bus, query.Vendor, query.Device

		_, ok1 := cache.mapdevices[query.Bus]

		if ok1 == true {

			_, ok2 := cache.mapdevices[query.Bus][query.Vendor]

			if ok2 == true {

				identifiers, ok3 := cache.mapdevices[query.Bus][query.Vendor][query.Device]

				if ok3 == true {

					for _, identifier := range identifiers {

						vulnerability, ok := cache.Vulnerabilities[identifier]

						if ok == true {

							matches_device := false

							for _, device := range vulnerability.Devices {

								matches_device_name := query.MatchesName(device.Name)
								matches_device_bus := query.MatchesBus(device.Bus)
								matches_device_vendor := query.MatchesVendor(device.Vendor)
								matches_device_device := query.MatchesDevice(device.Device)

								if matches_device_name && matches_device_bus && matches_device_vendor && matches_device_device {
									matches_device = true
									break
								}

							}

							if matches_device {
								found[vulnerability.Name] = vulnerability
							}

						}

					}

				}

			}

		}

	} else if query.Bus != "any" && query.Vendor != "any" && query.Device == "any" {

		// Cache Optimizations for specified query.Bus, query.Vendor

		_, ok1 := cache.mapdevices[query.Bus]

		if ok1 == true {

			_, ok2 := cache.mapdevices[query.Bus][query.Vendor]

			if ok2 == true {

				for _, identifiers := range cache.mapdevices[query.Bus][query.Vendor] {

					for _, identifier := range identifiers {

						vulnerability, ok := cache.Vulnerabilities[identifier]

						if ok == true {

							matches_device := false

							for _, device := range vulnerability.Devices {

								matches_device_name := query.MatchesName(device.Name)
								matches_device_bus := query.MatchesBus(device.Bus)
								matches_device_vendor := query.MatchesVendor(device.Vendor)
								matches_device_device := query.MatchesDevice(device.Device)

								if matches_device_name && matches_device_bus && matches_device_vendor && matches_device_device {
									matches_device = true
									break
								}

							}

							if matches_device {
								found[vulnerability.Name] = vulnerability
							}

						}

					}

				}

			}

		}

	} else if query.Bus != "any" && query.Vendor == "any" && query.Device == "any" {

		// Cache Optimizations for specified query.Bus

		_, ok1 := cache.mapdevices[query.Bus]

		if ok1 == true {

			for _, vendors := range cache.mapdevices[query.Bus] {

				for _, identifiers := range vendors {

					for _, identifier := range identifiers {

						vulnerability, ok := cache.Vulnerabilities[identifier]

						if ok == true {

							matches_device := false

							for _, device := range vulnerability.Devices {

								matches_device_name := query.MatchesName(device.Name)
								matches_device_bus := query.MatchesBus(device.Bus)
								matches_device_vendor := query.MatchesVendor(device.Vendor)
								matches_device_device := query.MatchesDevice(device.Device)

								if matches_device_name && matches_device_bus && matches_device_vendor && matches_device_device {
									matches_device = true
									break
								}

							}

							if matches_device {
								found[vulnerability.Name] = vulnerability
							}

						}

					}

				}

			}

		}

	} else {

		for _, vulnerability := range cache.Vulnerabilities {

			matches_device := false

			for _, device := range vulnerability.Devices {

				matches_device_name := query.MatchesName(device.Name)
				matches_device_bus := query.MatchesBus(device.Bus)
				matches_device_vendor := query.MatchesVendor(device.Vendor)
				matches_device_device := query.MatchesDevice(device.Device)

				if matches_device_name && matches_device_bus && matches_device_vendor && matches_device_device {
					matches_device = true
					break
				}

			}

			if matches_device {
				found[vulnerability.Name] = vulnerability
			}

		}

	}

	if len(found) > 0 {

		for _, vulnerability := range found {
			result = append(result, vulnerability)
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

func (cache *Vulnerabilities) QueryByDistribution(query matchers.Distribution) []*structs.Vulnerability {

	cache.mutex.RLock()

	result := make([]*structs.Vulnerability, 0)
	found := make(map[string]*structs.Vulnerability)

	if query.Name != "any" && query.Version != "any" {

		// Cache Optimizations for specified query.Name, query.Version

		_, ok1 := cache.mapdistributions[query.Name]

		if ok1 == true {

			identifiers, ok2 := cache.mapdistributions[query.Name][query.Version]

			if ok2 == true {

				for _, identifier := range identifiers {

					vulnerability, ok := cache.Vulnerabilities[identifier]

					if ok == true {

						matches_distribution := false

						for _, distribution := range vulnerability.Distributions {

							matches_distribution_name := query.MatchesName(distribution.Name)
							matches_distribution_version := query.MatchesVersion(distribution.Version)
							matches_distribution_manager := query.MatchesManager(distribution.Manager)
							matches_distribution_vendor := query.MatchesVendor(distribution.Vendor)

							if matches_distribution_name && matches_distribution_version && matches_distribution_manager && matches_distribution_vendor {
								matches_distribution = true
								break
							}

						}

						if matches_distribution {
							found[vulnerability.Name] = vulnerability
						}

					}

				}

			}

		}

	} else if query.Name != "any" && query.Version == "any" {

		// Cache Optimizations for specified query.Name

		_, ok1 := cache.mapdistributions[query.Name]

		if ok1 == true {

			for _, identifiers := range cache.mapdistributions[query.Name] {

				for _, identifier := range identifiers {

					vulnerability, ok := cache.Vulnerabilities[identifier]

					if ok == true {

						matches_distribution := false

						for _, distribution := range vulnerability.Distributions {

							matches_distribution_name := query.MatchesName(distribution.Name)
							matches_distribution_version := query.MatchesVersion(distribution.Version)
							matches_distribution_manager := query.MatchesManager(distribution.Manager)
							matches_distribution_vendor := query.MatchesVendor(distribution.Vendor)

							if matches_distribution_name && matches_distribution_version && matches_distribution_manager && matches_distribution_vendor {
								matches_distribution = true
								break
							}

						}

						if matches_distribution {
							found[vulnerability.Name] = vulnerability
						}

					}

				}

			}

		}

	} else {

		for _, vulnerability := range cache.Vulnerabilities {

			matches_distribution := false

			for _, distribution := range vulnerability.Distributions {

				matches_distribution_name := query.MatchesName(distribution.Name)
				matches_distribution_version := query.MatchesName(distribution.Version)
				matches_distribution_manager := query.MatchesManager(distribution.Manager)
				matches_distribution_vendor := query.MatchesVendor(distribution.Vendor)

				if matches_distribution_name && matches_distribution_version && matches_distribution_manager && matches_distribution_vendor {
					matches_distribution = true
					break
				}

			}

			if matches_distribution {
				found[vulnerability.Name] = vulnerability
			}

		}

	}

	if len(found) > 0 {

		for _, vulnerability := range found {
			result = append(result, vulnerability)
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

func (cache *Vulnerabilities) QueryByPackage(query matchers.Package) []*structs.Vulnerability {

	cache.mutex.RLock()

	result := make([]*structs.Vulnerability, 0)
	found := make(map[string]*structs.Vulnerability)

	if query.Vendor != "any" && query.Name != "any" {

		// Cache Optimizations for specified query.Name, query.Vendor

		_, ok1 := cache.mappackages[query.Vendor]

		if ok1 == true {

			identifiers, ok2 := cache.mappackages[query.Vendor][query.Name]

			if ok2 == true {

				for _, identifier := range identifiers {

					vulnerability, ok := cache.Vulnerabilities[identifier]

					if ok == true {

						matches_package := false

						for _, pkg := range vulnerability.Packages {

							matches_package_name := query.MatchesName(pkg.Name)
							matches_package_version := query.MatchesVersion(pkg.Version)
							matches_package_architecture := query.MatchesArchitecture(pkg.Architecture)
							matches_package_manager := query.MatchesManager(pkg.Manager)
							matches_package_vendor := query.MatchesVendor(pkg.Vendor)

							if matches_package_name && matches_package_version && matches_package_architecture && matches_package_manager && matches_package_vendor {
								matches_package = true
								break
							}

						}

						if matches_package {
							found[vulnerability.Name] = vulnerability
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

					vulnerability, ok := cache.Vulnerabilities[identifier]

					if ok == true {

						matches_package := false

						for _, pkg := range vulnerability.Packages {

							matches_package_name := query.MatchesName(pkg.Name)
							matches_package_version := query.MatchesVersion(pkg.Version)
							matches_package_architecture := query.MatchesArchitecture(pkg.Architecture)
							matches_package_manager := query.MatchesManager(pkg.Manager)
							matches_package_vendor := query.MatchesVendor(pkg.Vendor)

							if matches_package_name && matches_package_version && matches_package_architecture && matches_package_manager && matches_package_vendor {
								matches_package = true
								break
							}

						}

						if matches_package {
							found[vulnerability.Name] = vulnerability
						}

					}

				}

			}

		}

	} else {

		for _, vulnerability := range cache.Vulnerabilities {

			matches_package := false

			for _, pkg := range vulnerability.Packages {

				matches_package_name := query.MatchesName(pkg.Name)
				matches_package_version := query.MatchesVersion(pkg.Version)
				matches_package_architecture := query.MatchesArchitecture(pkg.Architecture)
				matches_package_manager := query.MatchesManager(pkg.Manager)
				matches_package_vendor := query.MatchesVendor(pkg.Vendor)

				if matches_package_name && matches_package_version && matches_package_architecture && matches_package_manager && matches_package_vendor {
					matches_package = true
					break
				}

			}

			if matches_package {
				found[vulnerability.Name] = vulnerability
			}

		}

	}

	if len(found) > 0 {

		for _, vulnerability := range found {
			result = append(result, vulnerability)
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

func (cache *Vulnerabilities) QueryByProduct(query matchers.Product) []*structs.Vulnerability {

	cache.mutex.RLock()

	result := make([]*structs.Vulnerability, 0)
	found := make(map[string]*structs.Vulnerability)

	if query.Vendor != "any" && query.Product != "any" && query.Version != "any" {

		// Cache Optimizations for specified query.Vendor, query.Product, query.Version

		_, ok1 := cache.mapproducts[query.Vendor]

		if ok1 == true {

			_, ok2 := cache.mapproducts[query.Vendor][query.Product]

			if ok2 == true {

				identifiers, ok3 := cache.mapproducts[query.Vendor][query.Product][query.Version]

				if ok3 == true {

					for _, identifier := range identifiers {

						vulnerability, ok := cache.Vulnerabilities[identifier]

						if ok == true {

							matches_product := false

							for _, product := range vulnerability.Products {

								if query.MatchesType(product.Type) {
									matches_product = true
									break
								}

							}

							if matches_product {
								found[vulnerability.Name] = vulnerability
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

						vulnerability, ok := cache.Vulnerabilities[identifier]

						if ok == true {

							matches_product := false

							for _, product := range vulnerability.Products {

								if query.MatchesType(product.Type) {
									matches_product = true
									break
								}

							}

							if matches_product {
								found[vulnerability.Name] = vulnerability
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

						vulnerability, ok := cache.Vulnerabilities[identifier]

						if ok == true {

							matches_product := false

							for _, product := range vulnerability.Products {

								if query.MatchesType(product.Type) {
									matches_product = true
									break
								}

							}

							if matches_product {
								found[vulnerability.Name] = vulnerability
							}

						}

					}

				}

			}

		}

	} else {

		for _, vulnerability := range cache.Vulnerabilities {

			matches_product := false

			for _, product := range vulnerability.Products {

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
				found[vulnerability.Name] = vulnerability
			}

		}

	}

	if len(found) > 0 {

		for _, vulnerability := range found {
			result = append(result, vulnerability)
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

func (cache *Vulnerabilities) QueryByProgram(query matchers.Program) []*structs.Vulnerability {

	cache.mutex.RLock()

	result := make([]*structs.Vulnerability, 0)
	found := make(map[string]*structs.Vulnerability)

	if query.Name != "any" {

		_, ok1 := cache.mapprograms[query.Name]

		if ok1 == true {

			for _, identifier := range cache.mapprograms[query.Name] {

				vulnerability, ok := cache.Vulnerabilities[identifier]

				if ok == true {

					matches_program := false

					for _, program := range vulnerability.Programs {

						if query.MatchesCommand(program.Command) {
							matches_program = true
							break
						}

					}

					if matches_program {
						found[vulnerability.Name] = vulnerability
					}

				}

			}

		}

	} else {

		for _, vulnerability := range cache.Vulnerabilities {

			matches_program := false

			for _, program := range vulnerability.Programs {

				matches_program_name := query.MatchesName(program.Name)
				matches_program_command := query.MatchesCommand(program.Command)

				if matches_program_name && matches_program_command {
					matches_program = true
					break
				}

			}

			if matches_program {
				found[vulnerability.Name] = vulnerability
			}

		}

	}

	if len(found) > 0 {

		for _, vulnerability := range found {
			result = append(result, vulnerability)
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

func (cache *Vulnerabilities) Remove(name string) bool {

	cache.mutex.Lock()

	var result bool

	if name != "" {

		vulnerability, ok1 := cache.Vulnerabilities[name]

		if ok1 == true {

			others, ok2 := cache.mapvulnerabilities[vulnerability.State][vulnerability.Type]

			if ok2 == true {

				filtered := make([]string, 0)

				for _, other := range others {

					if other != vulnerability.Name {
						filtered = append(filtered, other)
					}

				}

				if len(filtered) > 0 {
					cache.mapvulnerabilities[vulnerability.State][vulnerability.Type] = filtered
				} else {

					delete(cache.mapvulnerabilities[vulnerability.State], vulnerability.Type)

					if len(cache.mapvulnerabilities[vulnerability.State]) == 0 {
						delete(cache.mapvulnerabilities, vulnerability.State)
					}

				}

			}

			delete(cache.Vulnerabilities, name)
			result = true

			for _, connection := range vulnerability.Connections {

				types, ok2 := cache.mapconnections[connection.Type]

				if ok2 == true {

					protocols, ok3 := types[connection.Protocol]

					if ok3 == true {

						hosts, ok4 := protocols[connection.Host]

						if ok4 == true {

							others, ok5 := hosts[connection.Port]

							if ok5 == true {

								filtered := make([]string, 0)

								for _, other := range others {

									if other != name {
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

			for _, device := range vulnerability.Devices {

				if device.Bus != "any" && device.Vendor != "any" && device.Device != "any" {

					others, ok2 := cache.mapdevices[device.Bus][device.Vendor][device.Device]

					if ok2 == true {

						filtered := make([]string, 0)

						for _, other := range others {

							if other != name {
								filtered = append(filtered, other)
							}

						}

						if len(filtered) > 0 {
							cache.mapdevices[device.Bus][device.Vendor][device.Device] = filtered
						} else {

							delete(cache.mapdevices[device.Bus][device.Vendor], device.Device)

							if len(cache.mapdevices[device.Bus][device.Vendor]) == 0 {

								delete(cache.mapdevices[device.Bus], device.Vendor)

								if len(cache.mapdevices[device.Bus]) == 0 {
									delete(cache.mapdevices, device.Bus)
								}

							}

						}

					}

				}

			}
			for _, distribution := range vulnerability.Distributions {

				if distribution.Name != "any" && distribution.Version != "any" {

					others, ok2 := cache.mapdistributions[distribution.Name][distribution.Version]

					if ok2 == true {

						filtered := make([]string, 0)

						for _, other := range others {

							if other != name {
								filtered = append(filtered, other)
							}

						}

						if len(filtered) > 0 {
							cache.mapdistributions[distribution.Name][distribution.Version] = filtered
						} else {

							delete(cache.mapdistributions[distribution.Name], distribution.Version)

							if len(cache.mapdistributions[distribution.Name]) == 0 {
								delete(cache.mapdistributions, distribution.Name)
							}

						}

					}

				}

			}

			for _, pkg := range vulnerability.Packages {

				if pkg.Vendor != "any" && pkg.Name != "any" {

					others, ok2 := cache.mappackages[pkg.Vendor][pkg.Name]

					if ok2 == true {

						filtered := make([]string, 0)

						for _, other := range others {

							if other != name {
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

			}

			for _, product := range vulnerability.Products {

				versions, ok2 := cache.mapproducts[product.Vendor]

				if ok2 == true {

					products, ok3 := versions[product.Product]

					if ok3 == true {

						others, ok4 := products[product.Version]

						if ok4 == true {

							filtered := make([]string, 0)

							for _, other := range others {

								if other != name {
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

			for _, program := range vulnerability.Programs {

				if program.Name != "any" {

					others, ok2 := cache.mapprograms[program.Name]

					if ok2 == true {

						filtered := make([]string, 0)

						for _, other := range others {

							if other != name {
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

			}

		}

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Vulnerabilities) MarshalJSON() ([]byte, error) {

	tmp := make(map[string]*structs.Vulnerability)
	names := make([]string, 0)

	for name, _ := range cache.Vulnerabilities {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {
		tmp[name] = cache.Vulnerabilities[name]
	}

	return json.MarshalIndent(&struct {
		Vulnerabilities map[string]*structs.Vulnerability `json:"vulnerabilities"`
	}{
		Vulnerabilities: tmp,
	}, "", "\t")

}

func (cache *Vulnerabilities) UnmarshalJSON(data []byte) error {

	var tmp struct {
		Vulnerabilities map[string]*structs.Vulnerability `json:"vulnerabilities"`
	}

	err := json.Unmarshal(data, &tmp)

	if err != nil {
		return err
	}

	for _, vulnerability := range tmp.Vulnerabilities {
		cache.Add(*vulnerability)
	}

	return nil

}
