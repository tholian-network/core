package structs

import "tholian-core/types"
import "bytes"
import "encoding/json"
import "os"
import "strings"
import "time"

type System struct {
	Name         string            `json:"name"`
	Hostname     string            `json:"hostname"`
	Datetime     types.Datetime    `json:"datetime"`
	Distribution Distribution      `json:"distribution"`
	Fingerprint  SystemFingerprint `json:"fingerprint"`
	BIOS         Device            `json:"bios"`
	Board        Device            `json:"board"`
	Devices      []Device          `json:"devices"`
	Drives       []Drive           `json:"drives"`
	Incidents    []Incident        `json:"incidents"`
	Mitigations  []Mitigation      `json:"mitigations"`
	Networks     []Network         `json:"networks"`
	Packages     []Package         `json:"packages"`
	Programs     []Program         `json:"programs"`
	Responses    []Response        `json:"responses"`
	Services     []Program         `json:"services"`
	Antiques     []Antique         `json:"antiques"`
	Updates      []Update          `json:"updates"`
	Users        []types.User      `json:"users"`
}

func NewSystem() System {

	var system System

	stat, err1 := os.Stat("/etc/machine-id")

	if err1 == nil && stat.Size() > 0 {

		buffer, err12 := os.ReadFile("/etc/machine-id")

		if err12 == nil {
			system.SetName(string(buffer))
		}

	}

	hostname, err2 := os.Hostname()

	if err2 == nil {
		system.SetHostname(hostname)
	}

	system.SetDatetime(time.Now().Format(time.RFC3339))

	system.Devices = make([]Device, 0)
	system.Drives = make([]Drive, 0)
	system.Incidents = make([]Incident, 0)
	system.Mitigations = make([]Mitigation, 0)
	system.Networks = make([]Network, 0)
	system.Packages = make([]Package, 0)
	system.Programs = make([]Program, 0)
	system.Responses = make([]Response, 0)
	system.Services = make([]Program, 0)
	system.Antiques = make([]Antique, 0)
	system.Updates = make([]Update, 0)
	system.Users = make([]types.User, 0)

	return system

}

func (system *System) Equal(value System) bool {

	if system.Name != value.Name ||
		system.Hostname != value.Hostname ||
		system.Datetime.Equal(value.Datetime) == false ||
		system.Distribution.Equal(value.Distribution) == false ||
		system.Fingerprint.Equal(value.Fingerprint) == false ||
		system.BIOS.Equal(value.BIOS) == false ||
		system.Board.Equal(value.Board) == false {
		return false
	}

	if len(system.Devices) != len(value.Devices) {
		return false
	}

	for d := 0; d < len(system.Devices); d++ {

		if system.Devices[d].Equal(value.Devices[d]) == false {
			return false
		}

	}

	if len(system.Drives) != len(value.Drives) {
		return false
	}

	for d := 0; d < len(system.Drives); d++ {

		if system.Drives[d].Equal(value.Drives[d]) == false {
			return false
		}

	}

	if len(system.Incidents) != len(value.Incidents) {
		return false
	}

	for i := 0; i < len(system.Incidents); i++ {

		if system.Incidents[i].Equal(value.Incidents[i]) == false {
			return false
		}

	}

	if len(system.Mitigations) != len(value.Mitigations) {
		return false
	}

	for m := 0; m < len(system.Mitigations); m++ {

		if system.Mitigations[m].Equal(value.Mitigations[m]) == false {
			return false
		}

	}

	if len(system.Networks) != len(value.Networks) {
		return false
	}

	for n := 0; n < len(system.Networks); n++ {

		if system.Networks[n].Equal(value.Networks[n]) == false {
			return false
		}

	}

	if len(system.Packages) != len(value.Packages) {
		return false
	}

	for p := 0; p < len(system.Packages); p++ {

		if system.Packages[p].Equal(value.Packages[p]) == false {
			return false
		}

	}

	if len(system.Programs) != len(value.Programs) {
		return false
	}

	for p := 0; p < len(system.Programs); p++ {

		if system.Programs[p].Equal(value.Programs[p]) == false {
			return false
		}

	}

	if len(system.Responses) != len(value.Responses) {
		return false
	}

	for r := 0; r < len(system.Responses); r++ {

		if system.Responses[r].Equal(value.Responses[r]) == false {
			return false
		}

	}

	if len(system.Services) != len(value.Services) {
		return false
	}

	for s := 0; s < len(system.Services); s++ {

		if system.Services[s].Equal(value.Services[s]) == false {
			return false
		}

	}

	if len(system.Antiques) != len(value.Antiques) {
		return false
	}

	for a := 0; a < len(system.Antiques); a++ {

		if system.Antiques[a].Equal(value.Antiques[a]) == false {
			return false
		}

	}

	if len(system.Updates) != len(value.Updates) {
		return false
	}

	for u := 0; u < len(system.Updates); u++ {

		if system.Updates[u].Equal(value.Updates[u]) == false {
			return false
		}

	}

	if equalSystemUsers(system.Users, value.Users) == false {
		return false
	}

	return true

}

func (system *System) IsValid() bool {

	// Ignore BIOS, Board, Devices and Time

	if system.Name != "" {

		if system.Datetime.IsValid() == false {
			return false
		}

		if system.Distribution.IsValid() == false {
			return false
		}

		for i := 0; i < len(system.Incidents); i++ {

			if system.Incidents[i].IsValid() == false {
				return false
			}

		}

		for m := 0; m < len(system.Mitigations); m++ {

			if system.Mitigations[m].IsValid() == false {
				return false
			}

		}

		for n := 0; n < len(system.Networks); n++ {

			if system.Networks[n].IsValid() == false {
				return false
			}

		}

		for p := 0; p < len(system.Packages); p++ {

			if system.Packages[p].IsValid() == false {
				return false
			}

		}

		for p := 0; p < len(system.Programs); p++ {

			if system.Programs[p].IsValid() == false {
				return false
			}

		}

		for r := 0; r < len(system.Responses); r++ {

			if system.Responses[r].IsValid() == false {
				return false
			}

		}

		for s := 0; s < len(system.Services); s++ {

			if system.Services[s].IsValid() == false {
				return false
			}

		}

		for a := 0; a < len(system.Antiques); a++ {

			if system.Antiques[a].IsValid() == false {
				return false
			}

		}

		for u := 0; u < len(system.Updates); u++ {

			if system.Updates[u].IsValid() == false {
				return false
			}

		}

		for u := 0; u < len(system.Users); u++ {

			if system.Users[u].IsValid() == false {
				return false
			}

		}

		return true

	}

	return false

}

func (system *System) MarshalJSON() string {

	var buffer bytes.Buffer

	result, err := json.MarshalIndent(system, "", "\t")

	if err == nil {
		buffer.WriteString(string(result))
	}

	return buffer.String()

}

func (system *System) SetAntiques(antiques []Antique) {

	var filtered []Antique

	for a := 0; a < len(antiques); a++ {

		if antiques[a].IsValid() {
			filtered = append(filtered, antiques[a])
		}

	}

	system.Antiques = filtered

}

func (system *System) SetBIOS(value Device) {
	system.BIOS = value
}

func (system *System) SetBoard(value Device) {
	system.Board = value
}

func (system *System) SetDatetime(value string) {

	datetime := types.ToDatetime(value)

	if datetime.IsValid() {
		system.Datetime = datetime
	}

}

func (system *System) SetDrives(drives []Drive) {

	var filtered []Drive

	for d := 0; d < len(drives); d++ {

		if drives[d].IsValid() {
			filtered = append(filtered, drives[d])
		}

	}

	system.Drives = filtered

}

func (system *System) SetDevices(devices []Device) {

	var filtered []Device

	for d := 0; d < len(devices); d++ {

		if devices[d].IsValid() {
			filtered = append(filtered, devices[d])
		}

	}

	system.Devices = filtered

}

func (system *System) SetDistribution(distribution Distribution) {

	if distribution.IsValid() {
		system.Distribution = distribution
	}

}

func (system *System) SetHostname(value string) {
	system.Hostname = strings.TrimSpace(value)
}

func (system *System) AddIncident(value Incident) {

	if value.IsValid() {
		system.Incidents = append(system.Incidents, value)
	}

}

func (system *System) SetIncidents(incidents []Incident) {

	var filtered []Incident

	for i := 0; i < len(incidents); i++ {

		if incidents[i].IsValid() {
			filtered = append(filtered, incidents[i])
		}

	}

	system.Incidents = filtered

}

func (system *System) AddMitigation(value Mitigation) {

	if value.IsValid() {
		system.Mitigations = append(system.Mitigations, value)
	}

}

func (system *System) SetMitigations(mitigations []Mitigation) {

	var filtered []Mitigation

	for m := 0; m < len(mitigations); m++ {

		if mitigations[m].IsValid() {
			filtered = append(filtered, mitigations[m])
		}

	}

	system.Mitigations = filtered

}

func (system *System) SetName(value string) {
	system.Name = strings.TrimSpace(value)
}

func (system *System) SetNetworks(networks []Network) {

	var filtered []Network

	for n := 0; n < len(networks); n++ {

		if networks[n].IsValid() {
			filtered = append(filtered, networks[n])
		}

	}

	system.Networks = filtered

}

func (system *System) SetPackages(packages []Package) {

	var filtered []Package

	for p := 0; p < len(packages); p++ {

		if packages[p].IsValid() {
			filtered = append(filtered, packages[p])
		}

	}

	system.Packages = filtered

}

func (system *System) SetPrograms(programs []Program) {

	var filtered []Program

	for p := 0; p < len(programs); p++ {

		if programs[p].IsProgram() {
			filtered = append(filtered, programs[p])
		}

	}

	system.Programs = filtered

}

func (system *System) AddResponse(value Response) {

	if value.IsValid() {
		system.Responses = append(system.Responses, value)
	}

}

func (system *System) SetResponses(responses []Response) {

	var filtered []Response

	for r := 0; r < len(responses); r++ {

		if responses[r].IsValid() {
			filtered = append(filtered, responses[r])
		}

	}

	system.Responses = filtered

}

func (system *System) SetServices(services []Program) {

	var filtered []Program

	for s := 0; s < len(services); s++ {

		if services[s].IsService() {
			filtered = append(filtered, services[s])
		}

	}

	system.Services = filtered

}

func (system *System) SetTimezone(value string) {
	system.Fingerprint.Timezone = strings.TrimSpace(value)
}

func (system *System) SetCountry(value string) {
	system.Fingerprint.Country = strings.TrimSpace(value)
}

func (system *System) SetLocale(value string) {
	system.Fingerprint.Locale = strings.TrimSpace(value)
}

func (system *System) SetUpdates(updates []Update) {

	var filtered []Update

	for u := 0; u < len(updates); u++ {

		if updates[u].IsValid() {
			filtered = append(filtered, updates[u])
		}

	}

	system.Updates = filtered

}

func (system *System) SetUsers(users []types.User) {

	var filtered []types.User

	for u := 0; u < len(users); u++ {

		if users[u].IsValid() {
			filtered = append(filtered, users[u])
		}

	}

	system.Users = filtered

}

func (system *System) Hash() string {

	hash := ""

	if system.Name != "" {

		values := make([]string, 0)

		values = append(values, system.Name)
		values = append(values, system.Hostname)
		values = append(values, system.Datetime.String())
		values = append(values, system.Distribution.Hash())
		values = append(values, system.Fingerprint.Hash())
		values = append(values, system.BIOS.Hash())
		values = append(values, system.Board.Hash())

		for d := 0; d < len(system.Devices); d++ {
			values = append(values, system.Devices[d].Hash())
		}

		for d := 0; d < len(system.Drives); d++ {
			values = append(values, system.Drives[d].Hash())
		}

		for i := 0; i < len(system.Incidents); i++ {
			values = append(values, system.Incidents[i].Hash())
		}

		for m := 0; m < len(system.Mitigations); m++ {
			values = append(values, system.Mitigations[m].Hash())
		}

		for n := 0; n < len(system.Networks); n++ {
			values = append(values, system.Networks[n].Hash())
		}

		for p := 0; p < len(system.Packages); p++ {
			values = append(values, system.Packages[p].Hash())
		}

		for p := 0; p < len(system.Programs); p++ {
			values = append(values, system.Programs[p].Hash())
		}

		for r := 0; r < len(system.Responses); r++ {
			values = append(values, system.Responses[r].Hash())
		}

		for s := 0; s < len(system.Services); s++ {
			values = append(values, system.Services[s].Hash())
		}

		for a := 0; a < len(system.Antiques); a++ {
			values = append(values, system.Antiques[a].Hash())
		}

		for u := 0; u < len(system.Updates); u++ {
			values = append(values, system.Updates[u].Hash())
		}

		for u := 0; u < len(system.Users); u++ {
			values = append(values, system.Users[u].Hash())
		}

		hash = toHash(values...)

	}

	return hash

}
