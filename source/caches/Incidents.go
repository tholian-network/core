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
	Incidents          map[string]*structs.Incident            `json:"incidents"`
	mapasns            map[string][]string                     `json:"-"`
	// TODO: Support Connections
	mapdistributions   map[string][]string                     `json:"-"`
	mapincidents       map[string]map[string]map[string]string `json:"-"`
	mappackages        map[string]map[string][]string          `json:"-"`
	// TODO: Support Products
	mapprograms        map[string][]string                     `json:"-"`
	mapusers           map[string][]string                     `json:"-"`
	mapvulnerabilities map[string][]string                     `json:"-"`
	// TODO: Support Weaknesses
	networks           *Networks                               `json:"-"`
	mutex              *sync.RWMutex                           `json:"-"`
}

func NewIncidents(networks *Networks) *Incidents {

	var cache Incidents

	cache.Incidents = make(map[string]*structs.Incident)
	cache.mapasns = make(map[string][]string, 0)
	cache.mapdistributions = make(map[string][]string, 0)
	cache.mapincidents = make(map[string]map[string]map[string]string)
	cache.mappackages = make(map[string]map[string][]string, 0)
	cache.mapprograms = make(map[string][]string, 0)
	cache.mapusers = make(map[string][]string, 0)
	cache.mapvulnerabilities = make(map[string][]string, 0)

	cache.networks = networks
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

				subnet := structs.ToSubnet(network.Name)

				if network.Subnet != "any" {
					subnet.Parse(network.Subnet)
				}

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
