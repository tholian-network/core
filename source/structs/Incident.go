package structs

import "tholian-core/matchers"
import "tholian-core/types"
import "encoding/binary"
import "encoding/hex"
import "hash/crc32"
import "strings"
import "time"

type Incident struct {
	System          string                   `json:"system"`
	Type            string                   `json:"type"`
	Datetime        types.Datetime           `json:"datetime"`
	Connections     []matchers.Connection    `json:"connections"`
	Distributions   []matchers.Distribution  `json:"distributions"`
	Networks        []matchers.Network       `json:"networks"`
	Packages        []matchers.Package       `json:"packages"`
	Products        []matchers.Product       `json:"products"`
	Programs        []matchers.Program       `json:"programs"`
	Users           []matchers.User          `json:"users"`
	Vulnerabilities []matchers.Vulnerability `json:"vulnerabilities"`
	Weaknesses      []matchers.Weakness      `json:"weaknesses"`
}

func NewIncident(system string, typ string) Incident {

	var incident Incident

	incident.System = ""
	incident.Type = ""
	incident.SetDatetime(time.Now().Format(time.RFC3339))
	incident.Connections = make([]matchers.Connection, 0)
	incident.Distributions = make([]matchers.Distribution, 0)
	incident.Networks = make([]matchers.Network, 0)
	incident.Packages = make([]matchers.Package, 0)
	incident.Products = make([]matchers.Product, 0)
	incident.Programs = make([]matchers.Program, 0)
	incident.Users = make([]matchers.User, 0)
	incident.Vulnerabilities = make([]matchers.Vulnerability, 0)
	incident.Weaknesses = make([]matchers.Weakness, 0)

	incident.SetSystem(system)
	incident.SetType(typ)

	return incident

}

func (incident *Incident) IsAfter(other Incident) bool {

	if incident.Datetime.Equal(other.Datetime) {

		index1 := toIncidentTypeIndex(incident.Type)
		index2 := toIncidentTypeIndex(other.Type)

		if index1 > index2 {
			return true
		}

	} else if incident.Datetime.IsAfter(other.Datetime) {
		return true
	}

	return false

}

func (incident *Incident) IsBefore(other Incident) bool {

	if incident.Datetime.Equal(other.Datetime) {

		index1 := toIncidentTypeIndex(incident.Type)
		index2 := toIncidentTypeIndex(other.Type)

		if index1 < index2 {
			return true
		}

	} else if incident.Datetime.IsBefore(other.Datetime) {
		return true
	}

	return false

}

func (incident *Incident) Equal(value Incident) bool {

	if incident.System != value.System ||
		incident.Type != value.Type ||
		incident.Datetime.Equal(value.Datetime) == false {
		return false
	}

	if len(incident.Connections) != len(value.Connections) {
		return false
	}

	for c := 0; c < len(incident.Connections); c++ {

		if incident.Connections[c].Equal(value.Connections[c]) == false {
			return false
		}

	}

	if len(incident.Distributions) != len(value.Distributions) {
		return false
	}

	for d := 0; d < len(incident.Distributions); d++ {

		if incident.Distributions[d].Equal(value.Distributions[d]) == false {
			return false
		}

	}

	if len(incident.Networks) != len(value.Networks) {
		return false
	}

	for n := 0; n < len(incident.Networks); n++ {

		if incident.Networks[n].Equal(value.Networks[n]) == false {
			return false
		}

	}

	if len(incident.Packages) != len(value.Packages) {
		return false
	}

	for p := 0; p < len(incident.Packages); p++ {

		if incident.Packages[p].Equal(value.Packages[p]) == false {
			return false
		}

	}

	if len(incident.Products) != len(value.Products) {
		return false
	}

	for p := 0; p < len(incident.Products); p++ {

		if incident.Products[p].Equal(value.Products[p]) == false {
			return false
		}

	}

	if len(incident.Programs) != len(value.Programs) {
		return false
	}

	for p := 0; p < len(incident.Programs); p++ {

		if incident.Programs[p].Equal(value.Programs[p]) == false {
			return false
		}

	}

	if len(incident.Users) != len(value.Users) {
		return false
	}

	for u := 0; u < len(incident.Users); u++ {

		if incident.Users[u].Equal(value.Users[u]) == false {
			return false
		}

	}

	if len(incident.Vulnerabilities) != len(value.Vulnerabilities) {
		return false
	}

	for v := 0; v < len(incident.Vulnerabilities); v++ {

		if incident.Vulnerabilities[v].Equal(value.Vulnerabilities[v]) == false {
			return false
		}

	}

	if len(incident.Weaknesses) != len(value.Weaknesses) {
		return false
	}

	for w := 0; w < len(incident.Weaknesses); w++ {

		if incident.Weaknesses[w].Equal(value.Weaknesses[w]) == false {
			return false
		}

	}

	return true

}

func (incident *Incident) IsValid() bool {

	if incident.System != "" && incident.Type != "" && incident.Datetime.IsValid() {

		if incident.Type == "Recon" {

			if len(incident.Connections) > 0 || len(incident.Networks) > 0 {

				if len(incident.Programs) > 0 || len(incident.Users) > 0 {
					return true
				}

			}

		} else if incident.Type == "Intel" {

			if len(incident.Vulnerabilities) > 0 {

				if len(incident.Packages) > 0 || len(incident.Programs) > 0 {
					return true
				}

			}

		} else if incident.Type == "Conquer" {

			if len(incident.Distributions) > 0 && len(incident.Vulnerabilities) > 0 {

				if len(incident.Packages) > 0 || len(incident.Programs) > 0 {
					return true
				}

			}

		} else if incident.Type == "Persist" {

			if len(incident.Distributions) > 0 && len(incident.Vulnerabilities) > 0 {

				if len(incident.Packages) > 0 || len(incident.Programs) > 0 || len(incident.Users) > 0 {
					return true
				}

			}

		} else if incident.Type == "Exfil" {

			if len(incident.Connections) > 0 || len(incident.Networks) > 0 {

				if len(incident.Programs) > 0 || len(incident.Users) > 0 {
					return true
				}

			}

		} else if incident.Type == "Destroy" {

			if len(incident.Packages) > 0 || len(incident.Programs) > 0 || len(incident.Users) > 0 {
				return true
			}

		}

	}

	return false

}

func (incident *Incident) MatchesAntique(antique Antique) bool {

	for p := 0; p < len(incident.Packages); p++ {

		if incident.Packages[p].Matches(antique.Name, antique.Version.String(), antique.Architecture.String(), antique.Manager.String(), "any") {
			return true
		}

	}

	return false

}

func (incident *Incident) MatchesConnection(connection types.Connection) bool {

	for c := 0; c < len(incident.Connections); c++ {

		if incident.Connections[c].Matches(connection.Source.Host, connection.Source.Port, connection.Protocol.String(), connection.Type) {
			return true
		} else if incident.Connections[c].Matches(connection.Target.Host, connection.Target.Port, connection.Protocol.String(), connection.Type) {
			return true
		}

	}

	return false

}

func (incident *Incident) MatchesDistribution(distribution Distribution) bool {

	for d := 0; d < len(incident.Distributions); d++ {

		if incident.Distributions[d].Matches(distribution.Name, distribution.Version, distribution.Manager, distribution.Vendor) {
			return true
		}

	}

	return false

}

func (incident *Incident) MatchesNetwork(network Network) bool {

	for n := 0; n < len(incident.Networks); n++ {

		for s := 0; s < len(network.Subnets); s++ {

			if incident.Networks[n].Matches(network.Subnets[s].Name, network.Subnets[s].String()) {
				return true
			}

		}

	}

	return false

}

func (incident *Incident) MatchesPackage(pkg Package) bool {

	for p := 0; p < len(incident.Packages); p++ {

		if incident.Packages[p].Matches(pkg.Name, pkg.Version.String(), pkg.Architecture.String(), pkg.Manager.String(), pkg.Vendor) {
			return true
		}

	}

	return false

}

func (incident *Incident) MatchesProduct(product Product) bool {

	for p := 0; p < len(incident.Products); p++ {

		if incident.Products[p].Matches(product.Vendor, product.Product, product.Version, product.Type) {
			return true
		}

	}

	return false

}

func (incident *Incident) MatchesProgram(program Program) bool {

	for p := 0; p < len(incident.Programs); p++ {

		if incident.Programs[p].Matches(program.Name, program.Command, program.Arguments) {
			return true
		}

	}

	return false

}

func (incident *Incident) MatchesUpdate(update Update) bool {

	for p := 0; p < len(incident.Packages); p++ {

		if incident.Packages[p].Matches(update.Name, update.Version.String(), update.Architecture.String(), update.Manager.String(), "any") {
			return true
		}

	}

	return false

}

func (incident *Incident) MatchesUser(user types.User) bool {

	for u := 0; u < len(incident.Users); u++ {

		if incident.Users[u].Matches(user.Name, user.Password, user.Type) {
			return true
		}

	}

	return false

}

func (incident *Incident) MatchesVulnerability(vulnerability Vulnerability) bool {

	for v := 0; v < len(incident.Vulnerabilities); v++ {

		if incident.Vulnerabilities[v].Matches(vulnerability.Name, vulnerability.Type, vulnerability.Severity, vulnerability.State, vulnerability.Datetime.String()) {
			return true
		}

	}

	return false

}

func (incident *Incident) MatchesWeakness(weakness Weakness) bool {

	for _, matcher := range incident.Weaknesses {

		matches_name := matcher.MatchesName(weakness.Name)
		matches_impact := false
		matches_scope := false

		if matcher.Impact == "any" {
			matches_impact = true
		} else {

			for _, impact := range weakness.Impacts {

				if matcher.MatchesImpact(impact) {
					matches_impact = true
					break
				}

			}

		}

		if matcher.Scope == "any" {
			matches_scope = true
		} else {

			for _, scope := range weakness.Scopes {

				if matcher.MatchesScope(scope) {
					matches_scope = true
					break
				}

			}

		}

		if matches_name && matches_impact && matches_scope {
			return true
		}

	}

	return false

}

func (incident *Incident) SetDatetime(value string) {

	datetime := types.ToDatetime(value)

	if datetime.IsValid() {
		incident.Datetime = datetime
	}

}

func (incident *Incident) SetSystem(value string) {
	incident.System = strings.TrimSpace(value)
}

func (incident *Incident) SetType(value string) {

	if value == "Recon" {
		incident.Type = value
	} else if value == "Intel" {
		incident.Type = value
	} else if value == "Conquer" {
		incident.Type = value
	} else if value == "Persist" {
		incident.Type = value
	} else if value == "Exfil" {
		incident.Type = value
	} else if value == "Destroy" {
		incident.Type = value
	}

}

func (incident *Incident) AddConnection(value matchers.Connection) {

	if value.IsValid() {

		found := false

		for c := 0; c < len(incident.Connections); c++ {

			if incident.Connections[c].Equal(value) {
				found = true
				break
			}

		}

		if found == false {
			incident.Connections = append(incident.Connections, value)
		}

	}

}

func (incident *Incident) RemoveConnection(value matchers.Connection) {

	index := -1

	for c := 0; c < len(incident.Connections); c++ {

		if incident.Connections[c].Equal(value) {
			index = c
			break
		}

	}

	if index != -1 {
		incident.Connections = append(incident.Connections[:index], incident.Connections[index+1:]...)
	}

}

func (incident *Incident) SetConnections(values []matchers.Connection) {

	filtered := make([]matchers.Connection, 0)

	for _, value := range values {

		if value.IsValid() {
			filtered = append(filtered, value)
		}

	}

	incident.Connections = filtered

}

func (incident *Incident) AddDistribution(value matchers.Distribution) {

	if value.IsValid() {

		found := false

		for d := 0; d < len(incident.Distributions); d++ {

			if incident.Distributions[d].Equal(value) {
				found = true
				break
			}

		}

		if found == false {
			incident.Distributions = append(incident.Distributions, value)
		}

	}

}

func (incident *Incident) RemoveDistribution(value matchers.Distribution) {

	index := -1

	for d := 0; d < len(incident.Distributions); d++ {

		if incident.Distributions[d].Equal(value) {
			index = d
			break
		}

	}

	if index != -1 {
		incident.Distributions = append(incident.Distributions[:index], incident.Distributions[index+1:]...)
	}

}

func (incident *Incident) SetDistributions(values []matchers.Distribution) {

	filtered := make([]matchers.Distribution, 0)

	for _, value := range values {

		if value.IsValid() {
			filtered = append(filtered, value)
		}

	}

	incident.Distributions = filtered

}

func (incident *Incident) AddNetwork(value matchers.Network) {

	if value.IsValid() {

		found := false

		for n := 0; n < len(incident.Networks); n++ {

			if incident.Networks[n].Equal(value) {
				found = true
				break
			}

		}

		if found == false {
			incident.Networks = append(incident.Networks, value)
		}

	}

}

func (incident *Incident) RemoveNetwork(value matchers.Network) {

	index := -1

	for n := 0; n < len(incident.Networks); n++ {

		if incident.Networks[n].Equal(value) {
			index = n
			break
		}

	}

	if index != -1 {
		incident.Networks = append(incident.Networks[:index], incident.Networks[index+1:]...)
	}

}

func (incident *Incident) SetNetworks(values []matchers.Network) {

	filtered := make([]matchers.Network, 0)

	for _, value := range values {

		if value.IsValid() {
			filtered = append(filtered, value)
		}

	}

	incident.Networks = filtered

}

func (incident *Incident) AddPackage(value matchers.Package) {

	if value.IsValid() {

		found := false

		for p := 0; p < len(incident.Packages); p++ {

			if incident.Packages[p].Equal(value) {
				found = true
				break
			}

		}

		if found == false {
			incident.Packages = append(incident.Packages, value)
		}

	}

}

func (incident *Incident) RemovePackage(value matchers.Package) {

	index := -1

	for p := 0; p < len(incident.Packages); p++ {

		if incident.Packages[p].Equal(value) {
			index = p
			break
		}

	}

	if index != -1 {
		incident.Packages = append(incident.Packages[:index], incident.Packages[index+1:]...)
	}

}

func (incident *Incident) SetPackages(values []matchers.Package) {

	filtered := make([]matchers.Package, 0)

	for _, value := range values {

		if value.IsValid() {
			filtered = append(filtered, value)
		}

	}

	incident.Packages = filtered

}

func (incident *Incident) AddProduct(value matchers.Product) {

	if value.IsValid() {

		found := false

		for p := 0; p < len(incident.Products); p++ {

			if incident.Products[p].Equal(value) {
				found = true
				break
			}

		}

		if found == false {
			incident.Products = append(incident.Products, value)
		}

	}

}

func (incident *Incident) RemoveProduct(value matchers.Product) {

	index := -1

	for p := 0; p < len(incident.Products); p++ {

		if incident.Products[p].Equal(value) {
			index = p
			break
		}

	}

	if index != -1 {
		incident.Products = append(incident.Products[:index], incident.Products[index+1:]...)
	}

}

func (incident *Incident) SetProducts(values []matchers.Product) {

	filtered := make([]matchers.Product, 0)

	for _, value := range values {

		if value.IsValid() {
			filtered = append(filtered, value)
		}

	}

	incident.Products = filtered

}

func (incident *Incident) AddProgram(value matchers.Program) {

	if value.IsValid() {

		found := false

		for p := 0; p < len(incident.Programs); p++ {

			if incident.Programs[p].Equal(value) {
				found = true
				break
			}

		}

		if found == false {
			incident.Programs = append(incident.Programs, value)
		}

	}

}

func (incident *Incident) RemoveProgram(value matchers.Program) {

	index := -1

	for p := 0; p < len(incident.Programs); p++ {

		if incident.Programs[p].Equal(value) {
			index = p
			break
		}

	}

	if index != -1 {
		incident.Programs = append(incident.Programs[:index], incident.Programs[index+1:]...)
	}

}

func (incident *Incident) SetPrograms(values []matchers.Program) {

	filtered := make([]matchers.Program, 0)

	for _, value := range values {

		if value.IsValid() {
			filtered = append(filtered, value)
		}

	}

	incident.Programs = filtered

}

func (incident *Incident) AddUser(value matchers.User) {

	if value.IsValid() {

		found := false

		for u := 0; u < len(incident.Users); u++ {

			if incident.Users[u].Equal(value) {
				found = true
				break
			}

		}

		if found == false {
			incident.Users = append(incident.Users, value)
		}

	}

}

func (incident *Incident) RemoveUser(value matchers.User) {

	index := -1

	for u := 0; u < len(incident.Users); u++ {

		if incident.Users[u].Equal(value) {
			index = u
			break
		}

	}

	if index != -1 {
		incident.Users = append(incident.Users[:index], incident.Users[index+1:]...)
	}

}

func (incident *Incident) SetUsers(values []matchers.User) {

	filtered := make([]matchers.User, 0)

	for _, value := range values {

		if value.IsValid() {
			filtered = append(filtered, value)
		}

	}

	incident.Users = filtered

}

func (incident *Incident) AddVulnerability(value matchers.Vulnerability) {

	if value.IsValid() {

		found := false

		for v := 0; v < len(incident.Vulnerabilities); v++ {

			if incident.Vulnerabilities[v].Equal(value) {
				found = true
				break
			}

		}

		if found == false {
			incident.Vulnerabilities = append(incident.Vulnerabilities, value)
		}

	}

}

func (incident *Incident) RemoveVulnerability(value matchers.Vulnerability) {

	index := -1

	for v := 0; v < len(incident.Vulnerabilities); v++ {

		if incident.Vulnerabilities[v].Equal(value) {
			index = v
			break
		}

	}

	if index != -1 {
		incident.Vulnerabilities = append(incident.Vulnerabilities[:index], incident.Vulnerabilities[index+1:]...)
	}

}

func (incident *Incident) SetVulnerabilities(values []matchers.Vulnerability) {

	filtered := make([]matchers.Vulnerability, 0)

	for _, value := range values {

		if value.IsValid() {
			filtered = append(filtered, value)
		}

	}

	incident.Vulnerabilities = filtered

}

func (incident *Incident) AddWeakness(value matchers.Weakness) {

	if value.IsValid() {

		found := false

		for w := 0; w < len(incident.Weaknesses); w++ {

			if incident.Weaknesses[w].Equal(value) {
				found = true
				break
			}

		}

		if found == false {
			incident.Weaknesses = append(incident.Weaknesses, value)
		}

	}

}

func (incident *Incident) RemoveWeakness(value matchers.Weakness) {

	index := -1

	for w := 0; w < len(incident.Weaknesses); w++ {

		if incident.Weaknesses[w].Equal(value) {
			index = w
			break
		}

	}

	if index != -1 {
		incident.Weaknesses = append(incident.Weaknesses[:index], incident.Weaknesses[index+1:]...)
	}

}

func (incident *Incident) SetWeaknesses(values []matchers.Weakness) {

	filtered := make([]matchers.Weakness, 0)

	for _, value := range values {

		if value.IsValid() {
			filtered = append(filtered, value)
		}

	}

	incident.Weaknesses = filtered

}

func (incident *Incident) Hash() string {

	hash := ""

	if incident.Type != "" {

		uuid := make([]string, 0)

		uuid = append(uuid, incident.Type)
		uuid = append(uuid, incident.Datetime.String())

		for _, connection := range incident.Connections {
			uuid = append(uuid, connection.Hash())
		}

		for _, distribution := range incident.Distributions {
			uuid = append(uuid, distribution.Hash())
		}

		for _, network := range incident.Networks {
			uuid = append(uuid, network.Hash())
		}

		for _, pkg := range incident.Packages {
			uuid = append(uuid, pkg.Hash())
		}

		for _, product := range incident.Products {
			uuid = append(uuid, product.Hash())
		}

		for _, program := range incident.Programs {
			uuid = append(uuid, program.Hash())
		}

		for _, user := range incident.Users {
			uuid = append(uuid, user.Hash())
		}

		for _, vulnerability := range incident.Vulnerabilities {
			uuid = append(uuid, vulnerability.Hash())
		}

		for _, weakness := range incident.Weaknesses {
			uuid = append(uuid, toHash(weakness.Name, weakness.Impact, weakness.Scope))
		}

		checksum := crc32.ChecksumIEEE([]byte(strings.Join(uuid, "-")))

		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, checksum)
		hash = hex.EncodeToString(tmp)

	}

	return hash

}
