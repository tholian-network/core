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
	Programs        []matchers.Program       `json:"programs"`
	Users           []matchers.User          `json:"users"`
	Vulnerabilities []matchers.Vulnerability `json:"vulnerabilities"`
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
	incident.Programs = make([]matchers.Program, 0)
	incident.Users = make([]matchers.User, 0)
	incident.Vulnerabilities = make([]matchers.Vulnerability, 0)

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

	result := false

	for p := 0; p < len(incident.Packages); p++ {

		if incident.Packages[p].Matches(antique.Name, antique.Version.String(), antique.Architecture.String(), antique.Manager.String(), "any") {
			result = true
			break
		}

	}

	return result

}

func (incident *Incident) MatchesConnection(connection types.Connection) bool {

	result := false

	for c := 0; c < len(incident.Connections); c++ {

		if incident.Connections[c].Matches(connection.Source.Host, connection.Source.Port, connection.Protocol.String(), connection.Type) {
			result = true
			break
		} else if incident.Connections[c].Matches(connection.Target.Host, connection.Target.Port, connection.Protocol.String(), connection.Type) {
			result = true
			break
		}

	}

	return result

}

func (incident *Incident) MatchesDistribution(distribution Distribution) bool {

	result := false

	for d := 0; d < len(incident.Distributions); d++ {

		if incident.Distributions[d].Matches(distribution.Name, distribution.Version, distribution.Manager, distribution.Vendor) {
			result = true
			break
		}

	}

	return result

}

func (incident *Incident) MatchesNetwork(network Network) bool {

	result := false

	for n := 0; n < len(incident.Networks); n++ {

		for s := 0; s < len(network.Subnets); s++ {

			if incident.Networks[n].Matches(network.Subnets[s].Name, network.Subnets[s].String()) {
				result = true
				break
			}

		}

		if result == true {
			break
		}

	}

	return result

}

func (incident *Incident) MatchesPackage(pkg Package) bool {

	result := false

	for p := 0; p < len(incident.Packages); p++ {

		if incident.Packages[p].Matches(pkg.Name, pkg.Version.String(), pkg.Architecture.String(), pkg.Manager.String(), pkg.Vendor) {
			result = true
			break
		}

	}

	return result

}

func (incident *Incident) MatchesProgram(program Program) bool {

	result := false

	for p := 0; p < len(incident.Programs); p++ {

		if incident.Programs[p].Matches(program.Name, program.Command, program.Arguments) {
			result = true
			break
		}

	}

	return result

}

func (incident *Incident) MatchesUpdate(update Update) bool {

	result := false

	for p := 0; p < len(incident.Packages); p++ {

		if incident.Packages[p].Matches(update.Name, update.Version.String(), update.Architecture.String(), update.Manager.String(), "any") {
			result = true
			break
		}

	}

	return result

}

func (incident *Incident) MatchesUser(user types.User) bool {

	result := false

	for u := 0; u < len(incident.Users); u++ {

		if incident.Users[u].Matches(user.Name, user.Password, user.Type) {
			result = true
			break
		}

	}

	return result

}

func (incident *Incident) MatchesVulnerability(vulnerability Vulnerability) bool {

	result := false

	for v := 0; v < len(incident.Vulnerabilities); v++ {

		if incident.Vulnerabilities[v].Matches(vulnerability.Name, vulnerability.Type, vulnerability.Severity, vulnerability.State, vulnerability.Datetime.String()) {
			result = true
			break
		}

	}

	return result

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

		for _, program := range incident.Programs {
			uuid = append(uuid, program.Hash())
		}

		for _, user := range incident.Users {
			uuid = append(uuid, user.Hash())
		}

		for _, vulnerability := range incident.Vulnerabilities {
			uuid = append(uuid, vulnerability.Hash())
		}

		checksum := crc32.ChecksumIEEE([]byte(strings.Join(uuid, "-")))

		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, checksum)
		hash = hex.EncodeToString(tmp)

	}

	return hash

}
