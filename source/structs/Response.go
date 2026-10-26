package structs

import "tholian-core/matchers"
import "tholian-core/types"
import "encoding/binary"
import "encoding/hex"
import "hash/crc32"
import "strings"
import "time"

type Response struct {
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

func NewResponse(system string, typ string) Response {

	var response Response

	response.System = ""
	response.Type = "none"
	response.SetDatetime(time.Now().Format(time.RFC3339))
	response.Connections = make([]matchers.Connection, 0)
	response.Distributions = make([]matchers.Distribution, 0)
	response.Networks = make([]matchers.Network, 0)
	response.Packages = make([]matchers.Package, 0)
	response.Programs = make([]matchers.Program, 0)
	response.Users = make([]matchers.User, 0)
	response.Vulnerabilities = make([]matchers.Vulnerability, 0)

	response.SetSystem(system)
	response.SetType(typ)

	return response

}

func (response *Response) IsAfter(other Response) bool {

	if response.Datetime.Equal(other.Datetime) {

		index1 := toIncidentTypeIndex(response.Type)
		index2 := toIncidentTypeIndex(other.Type)

		if index1 > index2 {
			return true
		}

	} else if response.Datetime.IsAfter(other.Datetime) {
		return true
	}

	return false

}

func (response *Response) IsBefore(other Response) bool {

	if response.Datetime.Equal(other.Datetime) {

		index1 := toIncidentTypeIndex(response.Type)
		index2 := toIncidentTypeIndex(other.Type)

		if index1 < index2 {
			return true
		}

	} else if response.Datetime.IsBefore(other.Datetime) {
		return true
	}

	return false

}

func (response *Response) IsValid() bool {

	if response.System != "" && response.Type != "" && response.Datetime.IsValid() {

		if response.Type == "Update" {

			if len(response.Packages) > 0 || len(response.Programs) > 0 {
				return true
			}

		} else if response.Type == "Forbid" {

			if len(response.Connections) > 0 || len(response.Networks) > 0 || len(response.Programs) > 0 || len(response.Users) > 0 {
				return true
			}

		} else if response.Type == "Permit" {

			if len(response.Connections) > 0 || len(response.Networks) > 0 || len(response.Programs) > 0 || len(response.Users) > 0 {
				return true
			}

		} else if response.Type == "Recon" {

			if len(response.Connections) > 0 || len(response.Networks) > 0 {
				return true
			}

		} else if response.Type == "Intel" {

			if len(response.Connections) > 0 || len(response.Networks) > 0 {
				return true
			}

		} else if response.Type == "Conquer" {

			if len(response.Distributions) > 0 && len(response.Vulnerabilities) > 0 {

				if len(response.Connections) > 0 || len(response.Networks) > 0 {
					return true
				}

			}

		} else if response.Type == "Persist" {

			if len(response.Distributions) > 0 && len(response.Vulnerabilities) > 0 {

				if len(response.Connections) > 0 || len(response.Networks) > 0 {
					return true
				}

			}

		} else if response.Type == "Exfil" {

			if len(response.Connections) > 0 || len(response.Networks) > 0 {

				if len(response.Programs) > 0 || len(response.Users) > 0 {
					return true
				}

			}

		} else if response.Type == "Destroy" {

			if len(response.Packages) > 0 || len(response.Programs) > 0 || len(response.Users) > 0 {
				return true
			}

		}

	}

	return false

}

func (response *Response) MatchesAntique(antique Antique) bool {

	for p := 0; p < len(response.Packages); p++ {

		if response.Packages[p].Matches(antique.Name, antique.Version.String(), antique.Architecture.String(), antique.Manager.String(), "any") {
			return true
		}

	}

	return false

}

func (response *Response) MatchesConnection(connection types.Connection) bool {

	for c := 0; c < len(response.Connections); c++ {

		if response.Connections[c].Matches(connection.Source.Host, connection.Source.Port, connection.Protocol.String(), connection.Type) {
			return true
		} else if response.Connections[c].Matches(connection.Target.Host, connection.Target.Port, connection.Protocol.String(), connection.Type) {
			return true
		}

	}

	return false

}

func (response *Response) MatchesDistribution(distribution Distribution) bool {

	for d := 0; d < len(response.Distributions); d++ {

		if response.Distributions[d].Matches(distribution.Name, distribution.Version, distribution.Manager, distribution.Vendor) {
			return true
		}

	}

	return false

}

func (response *Response) MatchesNetwork(network Network) bool {

	for n := 0; n < len(response.Networks); n++ {

		for s := 0; s < len(network.Subnets); s++ {

			if response.Networks[n].Matches(network.Subnets[s].Name, network.Subnets[s].String()) {
				return true
			}

		}

	}

	return false

}

func (response *Response) MatchesPackage(pkg Package) bool {

	for p := 0; p < len(response.Packages); p++ {

		if response.Packages[p].Matches(pkg.Name, pkg.Version.String(), pkg.Architecture.String(), pkg.Manager.String(), pkg.Vendor) {
			return true
		}

	}

	return false

}

func (response *Response) MatchesProgram(program Program) bool {

	for p := 0; p < len(response.Programs); p++ {

		if response.Programs[p].Matches(program.Name, program.Command, program.Arguments) {
			return true
		}

	}

	return false

}

func (response *Response) MatchesUpdate(update Update) bool {

	for p := 0; p < len(response.Packages); p++ {

		if response.Packages[p].Matches(update.Name, update.Version.String(), update.Architecture.String(), update.Manager.String(), "any") {
			return true
		}

	}

	return false

}

func (response *Response) MatchesUser(user types.User) bool {

	for u := 0; u < len(response.Users); u++ {

		if response.Users[u].Matches(user.Name, user.Password, user.Type) {
			return true
		}

	}

	return false

}

func (response *Response) MatchesVulnerability(vulnerability Vulnerability) bool {

	for v := 0; v < len(response.Vulnerabilities); v++ {

		if response.Vulnerabilities[v].Matches(vulnerability.Name, vulnerability.Type, vulnerability.Severity, vulnerability.State, vulnerability.Datetime.String()) {
			return true
		}

	}

	return false

}

func (response *Response) SetDatetime(value string) {

	datetime := types.ToDatetime(value)

	if datetime.IsValid() {
		response.Datetime = datetime
	}

}

func (response *Response) SetSystem(value string) {
	response.System = strings.TrimSpace(value)
}

func (response *Response) SetType(value string) {

	if value == "Update" {
		response.Type = value
	} else if value == "Forbid" {
		response.Type = value
	} else if value == "Permit" {
		response.Type = value
	} else if value == "Recon" {
		response.Type = value
	} else if value == "Intel" {
		response.Type = value
	} else if value == "Conquer" {
		response.Type = value
	} else if value == "Persist" {
		response.Type = value
	} else if value == "Exfil" {
		response.Type = value
	} else if value == "Destroy" {
		response.Type = value
	}

}

func (response *Response) AddConnection(value matchers.Connection) {

	if value.IsValid() {

		found := false

		for c := 0; c < len(response.Connections); c++ {

			if response.Connections[c].Equal(value) {
				found = true
				break
			}

		}

		if found == false {
			response.Connections = append(response.Connections, value)
		}

	}

}

func (response *Response) RemoveConnection(value matchers.Connection) {

	index := -1

	for c := 0; c < len(response.Connections); c++ {

		if response.Connections[c].Equal(value) {
			index = c
			break
		}

	}

	if index != -1 {
		response.Connections = append(response.Connections[:index], response.Connections[index+1:]...)
	}

}

func (response *Response) SetConnections(value []matchers.Connection) {

	filtered := make([]matchers.Connection, 0)

	for v := 0; v < len(value); v++ {

		if value[v].IsValid() {
			filtered = append(filtered, value[v])
		}

	}

	response.Connections = filtered

}

func (response *Response) AddDistribution(value matchers.Distribution) {

	if value.IsValid() {

		found := false

		for d := 0; d < len(response.Distributions); d++ {

			if response.Distributions[d].Equal(value) {
				found = true
				break
			}

		}

		if found == false {
			response.Distributions = append(response.Distributions, value)
		}

	}

}

func (response *Response) RemoveDistribution(value matchers.Distribution) {

	index := -1

	for d := 0; d < len(response.Distributions); d++ {

		if response.Distributions[d].Equal(value) {
			index = d
			break
		}

	}

	if index != -1 {
		response.Distributions = append(response.Distributions[:index], response.Distributions[index+1:]...)
	}

}

func (response *Response) SetDistributions(value []matchers.Distribution) {

	filtered := make([]matchers.Distribution, 0)

	for v := 0; v < len(value); v++ {

		if value[v].IsValid() {
			filtered = append(filtered, value[v])
		}

	}

	response.Distributions = filtered

}

func (response *Response) AddNetwork(value matchers.Network) {

	if value.IsValid() {

		found := false

		for n := 0; n < len(response.Networks); n++ {

			if response.Networks[n].Equal(value) {
				found = true
				break
			}

		}

		if found == false {
			response.Networks = append(response.Networks, value)
		}

	}

}

func (response *Response) RemoveNetwork(value matchers.Network) {

	index := -1

	for n := 0; n < len(response.Networks); n++ {

		if response.Networks[n].Equal(value) {
			index = n
			break
		}

	}

	if index != -1 {
		response.Networks = append(response.Networks[:index], response.Networks[index+1:]...)
	}

}

func (response *Response) SetNetworks(value []matchers.Network) {

	filtered := make([]matchers.Network, 0)

	for v := 0; v < len(value); v++ {

		if value[v].IsValid() {
			filtered = append(filtered, value[v])
		}

	}

	response.Networks = filtered

}

func (response *Response) AddPackage(value matchers.Package) {

	if value.IsValid() {

		found := false

		for p := 0; p < len(response.Packages); p++ {

			if response.Packages[p].Equal(value) {
				found = true
				break
			}

		}

		if found == false {
			response.Packages = append(response.Packages, value)
		}

	}

}

func (response *Response) RemovePackage(value matchers.Package) {

	index := -1

	for p := 0; p < len(response.Packages); p++ {

		if response.Packages[p].Equal(value) {
			index = p
			break
		}

	}

	if index != -1 {
		response.Packages = append(response.Packages[:index], response.Packages[index+1:]...)
	}

}

func (response *Response) SetPackages(value []matchers.Package) {

	filtered := make([]matchers.Package, 0)

	for v := 0; v < len(value); v++ {

		if value[v].IsValid() {
			filtered = append(filtered, value[v])
		}

	}

	response.Packages = filtered

}

func (response *Response) AddProgram(value matchers.Program) {

	if value.IsValid() {

		found := false

		for p := 0; p < len(response.Programs); p++ {

			if response.Programs[p].Equal(value) {
				found = true
				break
			}

		}

		if found == false {
			response.Programs = append(response.Programs, value)
		}

	}

}

func (response *Response) RemoveProgram(value matchers.Program) {

	index := -1

	for p := 0; p < len(response.Programs); p++ {

		if response.Programs[p].Equal(value) {
			index = p
			break
		}

	}

	if index != -1 {
		response.Programs = append(response.Programs[:index], response.Programs[index+1:]...)
	}

}

func (response *Response) SetPrograms(value []matchers.Program) {

	filtered := make([]matchers.Program, 0)

	for v := 0; v < len(value); v++ {

		if value[v].IsValid() {
			filtered = append(filtered, value[v])
		}

	}

	response.Programs = filtered

}

func (response *Response) AddUser(value matchers.User) {

	if value.IsValid() {

		found := false

		for u := 0; u < len(response.Users); u++ {

			if response.Users[u].Equal(value) {
				found = true
				break
			}

		}

		if found == false {
			response.Users = append(response.Users, value)
		}

	}

}

func (response *Response) RemoveUser(value matchers.User) {

	index := -1

	for u := 0; u < len(response.Users); u++ {

		if response.Users[u].Equal(value) {
			index = u
			break
		}

	}

	if index != -1 {
		response.Users = append(response.Users[:index], response.Users[index+1:]...)
	}

}

func (response *Response) SetUsers(value []matchers.User) {

	filtered := make([]matchers.User, 0)

	for v := 0; v < len(value); v++ {

		if value[v].IsValid() {
			filtered = append(filtered, value[v])
		}

	}

	response.Users = filtered

}

func (response *Response) AddVulnerability(value matchers.Vulnerability) {

	if value.IsValid() {

		found := false

		for v := 0; v < len(response.Vulnerabilities); v++ {

			if response.Vulnerabilities[v].Equal(value) {
				found = true
				break
			}

		}

		if found == false {
			response.Vulnerabilities = append(response.Vulnerabilities, value)
		}

	}

}

func (response *Response) RemoveVulnerability(value matchers.Vulnerability) {

	index := -1

	for v := 0; v < len(response.Vulnerabilities); v++ {

		if response.Vulnerabilities[v].Equal(value) {
			index = v
			break
		}

	}

	if index != -1 {
		response.Vulnerabilities = append(response.Vulnerabilities[:index], response.Vulnerabilities[index+1:]...)
	}

}

func (response *Response) SetVulnerabilities(value []matchers.Vulnerability) {

	filtered := make([]matchers.Vulnerability, 0)

	for v := 0; v < len(value); v++ {

		if value[v].IsValid() {
			filtered = append(filtered, value[v])
		}

	}

	response.Vulnerabilities = filtered

}

func (response *Response) Hash() string {

	hash := ""

	if response.Type != "" {

		uuid := make([]string, 0)

		uuid = append(uuid, response.Type)
		uuid = append(uuid, response.Datetime.String())

		for _, connection := range response.Connections {
			uuid = append(uuid, connection.Hash())
		}

		for _, distribution := range response.Distributions {
			uuid = append(uuid, distribution.Hash())
		}

		for _, network := range response.Networks {
			uuid = append(uuid, network.Hash())
		}

		for _, pkg := range response.Packages {
			uuid = append(uuid, pkg.Hash())
		}

		for _, program := range response.Programs {
			uuid = append(uuid, program.Hash())
		}

		for _, user := range response.Users {
			uuid = append(uuid, user.Hash())
		}

		for _, vulnerability := range response.Vulnerabilities {
			uuid = append(uuid, vulnerability.Hash())
		}

		checksum := crc32.ChecksumIEEE([]byte(strings.Join(uuid, "-")))

		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, checksum)
		hash = hex.EncodeToString(tmp)

	}

	return hash

}
