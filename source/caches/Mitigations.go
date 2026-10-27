package caches

import "tholian-core/matchers"
import "tholian-core/structs"
import "encoding/json"
import "sort"
import "strings"
import "sync"

func toMitigationIdentifier(value structs.Mitigation) string {

	result := ""
	result += value.Incident.System + ":" + value.Incident.Type + ":" + value.Incident.Datetime.String()
	result += "|"
	result += value.Response.System + ":" + value.Response.Type + ":" + value.Response.Datetime.String()

	return result

}

type Mitigations struct {
	Mitigations  map[string]*structs.Mitigation `json:"mitigations"`
	mapincidents map[string]string              `json:"-"`
	mapresponses map[string]string              `json:"-"`
	incidents    *Incidents                     `json:"-"`
	responses    *Responses                     `json:"-"`
	mutex        *sync.RWMutex                  `json:"-"`
}

func NewMitigations(networks *Networks) *Mitigations {

	var cache Mitigations

	cache.Mitigations = make(map[string]*structs.Mitigation)
	cache.mapincidents = make(map[string]string)
	cache.mapresponses = make(map[string]string)
	cache.incidents = NewIncidents(networks)
	cache.responses = NewResponses(networks)
	cache.mutex = &sync.RWMutex{}

	return &cache

}

func (cache *Mitigations) Add(value structs.Mitigation) bool {

	cache.mutex.Lock()

	var result bool

	if value.Incident.System != "" &&
		value.Incident.Type != "" &&
		value.Incident.Datetime.IsValid() &&
		value.Response.System != "" &&
		value.Response.Type != "" &&
		value.Response.Datetime.IsValid() {

		result_incident := cache.incidents.Add(value.Incident)
		result_response := cache.responses.Add(value.Response)

		if result_incident && result_response {

			identifier := toMitigationIdentifier(value)
			identifier_incident := toIncidentIdentifier(value.Incident)
			identifier_response := toResponseIdentifier(value.Response)

			cache.Mitigations[identifier] = &value
			cache.mapincidents[identifier_incident] = identifier
			cache.mapresponses[identifier_response] = identifier

		}

		return result_incident && result_response

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Mitigations) GetByIncident(identifier string) *structs.Mitigation {

	cache.mutex.RLock()

	var result *structs.Mitigation = nil

	if strings.Contains(identifier, ":") {

		mitigation_identifier, ok1 := cache.mapincidents[identifier]

		if ok1 == true {

			mitigation, ok := cache.Mitigations[mitigation_identifier]

			if ok == true {
				result = mitigation
			}

		}

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Mitigations) GetByResponse(identifier string) *structs.Mitigation {

	cache.mutex.RLock()

	var result *structs.Mitigation = nil

	if strings.Contains(identifier, ":") {

		mitigation_identifier, ok1 := cache.mapresponses[identifier]

		if ok1 == true {

			mitigation, ok := cache.Mitigations[mitigation_identifier]

			if ok == true {
				result = mitigation
			}

		}

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Mitigations) Length() int {
	return len(cache.Mitigations)
}

func (cache *Mitigations) Query(query matchers.Mitigation) []*structs.Mitigation {

	cache.mutex.RLock()

	result := make([]*structs.Mitigation, 0)
	found := make(map[string]*structs.Mitigation)

	incidents := cache.incidents.Query(matchers.Incident{
		System:   query.System,
		Type:     query.IncidentType,
		Datetime: query.Datetime,
	})

	responses := cache.responses.Query(matchers.Response{
		System:   query.System,
		Type:     query.ResponseType,
		Datetime: query.Datetime,
	})

	for _, incident := range incidents {

		identifier := toIncidentIdentifier(*incident)
		mitigation_identifier, ok1 := cache.mapincidents[identifier]

		if ok1 == true {

			mitigation, ok := cache.Mitigations[identifier]

			if ok == true {
				found[mitigation_identifier] = mitigation
			}

		}

	}

	for _, response := range responses {

		identifier := toResponseIdentifier(*response)
		mitigation_identifier, ok1 := cache.mapresponses[identifier]

		if ok1 == true {

			mitigation, ok := cache.Mitigations[identifier]

			if ok == true {
				found[mitigation_identifier] = mitigation
			}

		}

	}

	if len(found) > 0 {

		for _, mitigation := range found {
			result = append(result, mitigation)
		}

	}

	if len(result) > 1 {

		sort.Slice(result, func(a int, b int) bool {
			return result[a].Incident.IsAfter(result[b].Incident)
		})

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Mitigations) QueryByDistribution(query matchers.Distribution) []*structs.Mitigation {

	cache.mutex.RLock()

	result := make([]*structs.Mitigation, 0)
	found := make(map[string]*structs.Mitigation)

	incidents := cache.incidents.QueryByDistribution(query)
	responses := cache.responses.QueryByDistribution(query)

	for _, incident := range incidents {

		identifier := toIncidentIdentifier(*incident)
		mitigation_identifier, ok1 := cache.mapincidents[identifier]

		if ok1 == true {

			mitigation, ok := cache.Mitigations[identifier]

			if ok == true {
				found[mitigation_identifier] = mitigation
			}

		}

	}

	for _, response := range responses {

		identifier := toResponseIdentifier(*response)
		mitigation_identifier, ok1 := cache.mapresponses[identifier]

		if ok1 == true {

			mitigation, ok := cache.Mitigations[identifier]

			if ok == true {
				found[mitigation_identifier] = mitigation
			}

		}

	}

	if len(found) > 0 {

		for _, mitigation := range found {
			result = append(result, mitigation)
		}

	}

	if len(result) > 1 {

		sort.Slice(result, func(a int, b int) bool {
			return result[a].Incident.IsAfter(result[b].Incident)
		})

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Mitigations) QueryByIP(query string) []*structs.Mitigation {

	cache.mutex.RLock()

	result := make([]*structs.Mitigation, 0)
	found := make(map[string]*structs.Mitigation)

	incidents := cache.incidents.QueryByIP(query)
	responses := cache.responses.QueryByIP(query)

	for _, incident := range incidents {

		identifier := toIncidentIdentifier(*incident)
		mitigation_identifier, ok1 := cache.mapincidents[identifier]

		if ok1 == true {

			mitigation, ok := cache.Mitigations[identifier]

			if ok == true {
				found[mitigation_identifier] = mitigation
			}

		}

	}

	for _, response := range responses {

		identifier := toResponseIdentifier(*response)
		mitigation_identifier, ok1 := cache.mapresponses[identifier]

		if ok1 == true {

			mitigation, ok := cache.Mitigations[identifier]

			if ok == true {
				found[mitigation_identifier] = mitigation
			}

		}

	}

	if len(found) > 0 {

		for _, mitigation := range found {
			result = append(result, mitigation)
		}

	}

	if len(result) > 1 {

		sort.Slice(result, func(a int, b int) bool {
			return result[a].Incident.IsAfter(result[b].Incident)
		})

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Mitigations) QueryByPackage(query matchers.Package) []*structs.Mitigation {

	cache.mutex.RLock()

	result := make([]*structs.Mitigation, 0)
	found := make(map[string]*structs.Mitigation)

	incidents := cache.incidents.QueryByPackage(query)
	responses := cache.responses.QueryByPackage(query)

	for _, incident := range incidents {

		identifier := toIncidentIdentifier(*incident)
		mitigation_identifier, ok1 := cache.mapincidents[identifier]

		if ok1 == true {

			mitigation, ok := cache.Mitigations[identifier]

			if ok == true {
				found[mitigation_identifier] = mitigation
			}

		}

	}

	for _, response := range responses {

		identifier := toResponseIdentifier(*response)
		mitigation_identifier, ok1 := cache.mapresponses[identifier]

		if ok1 == true {

			mitigation, ok := cache.Mitigations[identifier]

			if ok == true {
				found[mitigation_identifier] = mitigation
			}

		}

	}

	if len(found) > 0 {

		for _, mitigation := range found {
			result = append(result, mitigation)
		}

	}

	if len(result) > 1 {

		sort.Slice(result, func(a int, b int) bool {
			return result[a].Incident.IsAfter(result[b].Incident)
		})

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Mitigations) QueryByProgram(query matchers.Program) []*structs.Mitigation {

	cache.mutex.RLock()

	result := make([]*structs.Mitigation, 0)
	found := make(map[string]*structs.Mitigation)

	incidents := cache.incidents.QueryByProgram(query)
	responses := cache.responses.QueryByProgram(query)

	for _, incident := range incidents {

		identifier := toIncidentIdentifier(*incident)
		mitigation_identifier, ok1 := cache.mapincidents[identifier]

		if ok1 == true {

			mitigation, ok := cache.Mitigations[identifier]

			if ok == true {
				found[mitigation_identifier] = mitigation
			}

		}

	}

	for _, response := range responses {

		identifier := toResponseIdentifier(*response)
		mitigation_identifier, ok1 := cache.mapresponses[identifier]

		if ok1 == true {

			mitigation, ok := cache.Mitigations[identifier]

			if ok == true {
				found[mitigation_identifier] = mitigation
			}

		}

	}

	if len(found) > 0 {

		for _, mitigation := range found {
			result = append(result, mitigation)
		}

	}

	if len(result) > 1 {

		sort.Slice(result, func(a int, b int) bool {
			return result[a].Incident.IsAfter(result[b].Incident)
		})

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Mitigations) QueryByUser(query matchers.User) []*structs.Mitigation {

	cache.mutex.RLock()

	result := make([]*structs.Mitigation, 0)
	found := make(map[string]*structs.Mitigation)

	incidents := cache.incidents.QueryByUser(query)
	responses := cache.responses.QueryByUser(query)

	for _, incident := range incidents {

		identifier := toIncidentIdentifier(*incident)
		mitigation_identifier, ok1 := cache.mapincidents[identifier]

		if ok1 == true {

			mitigation, ok := cache.Mitigations[identifier]

			if ok == true {
				found[mitigation_identifier] = mitigation
			}

		}

	}

	for _, response := range responses {

		identifier := toResponseIdentifier(*response)
		mitigation_identifier, ok1 := cache.mapresponses[identifier]

		if ok1 == true {

			mitigation, ok := cache.Mitigations[identifier]

			if ok == true {
				found[mitigation_identifier] = mitigation
			}

		}

	}

	if len(found) > 0 {

		for _, mitigation := range found {
			result = append(result, mitigation)
		}

	}

	if len(result) > 1 {

		sort.Slice(result, func(a int, b int) bool {
			return result[a].Incident.IsAfter(result[b].Incident)
		})

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Mitigations) QueryByVulnerability(query matchers.Vulnerability) []*structs.Mitigation {

	cache.mutex.RLock()

	result := make([]*structs.Mitigation, 0)
	found := make(map[string]*structs.Mitigation)

	incidents := cache.incidents.QueryByVulnerability(query)
	responses := cache.responses.QueryByVulnerability(query)

	for _, incident := range incidents {

		identifier := toIncidentIdentifier(*incident)
		mitigation_identifier, ok1 := cache.mapincidents[identifier]

		if ok1 == true {

			mitigation, ok := cache.Mitigations[identifier]

			if ok == true {
				found[mitigation_identifier] = mitigation
			}

		}

	}

	for _, response := range responses {

		identifier := toResponseIdentifier(*response)
		mitigation_identifier, ok1 := cache.mapresponses[identifier]

		if ok1 == true {

			mitigation, ok := cache.Mitigations[identifier]

			if ok == true {
				found[mitigation_identifier] = mitigation
			}

		}

	}

	if len(found) > 0 {

		for _, mitigation := range found {
			result = append(result, mitigation)
		}

	}

	if len(result) > 1 {

		sort.Slice(result, func(a int, b int) bool {
			return result[a].Incident.IsAfter(result[b].Incident)
		})

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Mitigations) RemoveByIncident(incident_identifier string) bool {

	cache.mutex.Lock()

	var result bool

	if strings.Contains(incident_identifier, ":") {

		mitigation_identifier, ok1 := cache.mapincidents[incident_identifier]

		if ok1 == true {

			mitigation, ok := cache.Mitigations[mitigation_identifier]

			if ok == true {

				delete(cache.Mitigations, mitigation_identifier)
				delete(cache.mapincidents, incident_identifier)
				delete(cache.mapresponses, toResponseIdentifier(mitigation.Response))

				result = true

			}

		}

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Mitigations) RemoveByResponse(response_identifier string) bool {

	cache.mutex.Lock()

	var result bool

	if strings.Contains(response_identifier, ":") {

		mitigation_identifier, ok1 := cache.mapresponses[response_identifier]

		if ok1 == true {

			mitigation, ok := cache.Mitigations[mitigation_identifier]

			if ok == true {

				delete(cache.Mitigations, mitigation_identifier)
				delete(cache.mapincidents, toIncidentIdentifier(mitigation.Incident))
				delete(cache.mapresponses, response_identifier)

				result = true

			}

		}

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Mitigations) MarshalJSON() ([]byte, error) {

	tmp := make(map[string]*structs.Mitigation)
	identifiers := make([]string, 0)

	for identifier, _ := range cache.Mitigations {
		identifiers = append(identifiers, identifier)
	}

	sort.Strings(identifiers)

	for _, identifier := range identifiers {
		tmp[identifier] = cache.Mitigations[identifier]
	}

	return json.MarshalIndent(&struct {
		Mitigations map[string]*structs.Mitigation `json:"mitigations"`
	}{
		Mitigations: tmp,
	}, "", "\t")

}

func (cache *Mitigations) UnmarshalJSON(data []byte) error {

	var tmp struct {
		Mitigations map[string]*structs.Mitigation `json:"mitigations"`
	}

	err := json.Unmarshal(data, &tmp)

	if err != nil {
		return err
	}

	for _, mitigation := range tmp.Mitigations {
		cache.Add(*mitigation)
	}

	return nil

}
