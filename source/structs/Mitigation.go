package structs

import utils_strings "tholian-core/utils/strings"
import "strings"

type Mitigation struct {
	Incident Incident `json:"incident"`
	Response Response `json:"response"`
}

func NewMitigation(incident Incident, response Response) Mitigation {

	var mitigation Mitigation

	mitigation.SetIncident(incident)
	mitigation.SetResponse(response)

	return mitigation

}

func (mitigation *Mitigation) Equal(value Mitigation) bool {

	if mitigation.Incident.Equal(value.Incident) && mitigation.Response.Equal(value.Response) {
		return true
	}

	return false

}

func (mitigation *Mitigation) IsValid() bool {

	if mitigation.Incident.IsValid() &&
		mitigation.Response.IsValid() {
		return true
	}

	return false

}

func (mitigation *Mitigation) SetIncident(value Incident) {

	if value.IsValid() {
		mitigation.Incident = value
	}

}

func (mitigation *Mitigation) SetResponse(value Response) {

	if value.IsValid() {
		mitigation.Response = value
	}

}

func (mitigation *Mitigation) Hash() string {

	result := ""
	// incident := mitigation.Incident
	response := mitigation.Response

	if response.Type == "Update" {

		if len(response.Packages) > 0 {

			hashes := make([]string, 0)

			for _, pkg := range response.Packages {

				hash := pkg.Hash()

				if hash != "" {
					hashes = append(hashes, hash)
				}

			}

			if len(hashes) > 0 {
				result = response.Type + "-Packages-" + strings.Join(utils_strings.Unique(hashes), ",")
			}

		} else if len(response.Programs) > 0 {

			hashes := make([]string, 0)

			for _, program := range response.Programs {

				hash := program.Hash()

				if hash != "" {
					hashes = append(hashes, hash)
				}

			}

			if len(hashes) > 0 {
				result = response.Type + "-Programs-" + strings.Join(utils_strings.Unique(hashes), ",")
			}

		}

	} else if response.Type == "Forbid" || response.Type == "Permit" {

		if len(response.Connections) > 0 {

			hashes := make([]string, 0)

			for _, connection := range response.Connections {

				hash := connection.Hash()

				if hash != "" {
					hashes = append(hashes, hash)
				}

			}

			if len(hashes) > 0 {
				result = response.Type + "-Connections-" + strings.Join(utils_strings.Unique(hashes), ",")
			}

		} else if len(response.Networks) > 0 {

			hashes := make([]string, 0)

			for _, network := range response.Networks {

				hash := network.Hash()

				if hash != "" {
					hashes = append(hashes, hash)
				}

			}

			if len(hashes) > 0 {
				result = response.Type + "-Networks-" + strings.Join(utils_strings.Unique(hashes), ",")
			}

		} else if len(response.Programs) > 0 {

			hashes := make([]string, 0)

			for _, program := range response.Programs {

				hash := program.Hash()

				if hash != "" {
					hashes = append(hashes, hash)
				}

			}

			if len(hashes) > 0 {
				result = response.Type + "-Programs-" + strings.Join(utils_strings.Unique(hashes), ",")
			}

		} else if len(response.Users) > 0 {

			hashes := make([]string, 0)

			for _, user := range response.Users {

				hash := user.Hash()

				if hash != "" {
					hashes = append(hashes, hash)
				}

			}

			if len(hashes) > 0 {
				result = response.Type + "-Users-" + strings.Join(utils_strings.Unique(hashes), ",")
			}

		}

	} else if response.Type == "Recon" || response.Type == "Intel" || response.Type == "Conquer" || response.Type == "Persist" || response.Type == "Exfil" {

		if len(response.Connections) > 0 {

			hashes := make([]string, 0)

			for _, connection := range response.Connections {

				hash := connection.Hash()

				if hash != "" {
					hashes = append(hashes, hash)
				}

			}

			if len(hashes) > 0 {
				result = response.Type + "-Connections-" + strings.Join(utils_strings.Unique(hashes), ",")
			}

		} else if len(response.Networks) > 0 {

			hashes := make([]string, 0)

			for _, network := range response.Networks {

				hash := network.Hash()

				if hash != "" {
					hashes = append(hashes, hash)
				}

			}

			if len(hashes) > 0 {
				result = response.Type + "-Networks-" + strings.Join(utils_strings.Unique(hashes), ",")
			}

		}

	} else if response.Type == "Destroy" {

		if len(response.Packages) > 0 {

			hashes := make([]string, 0)

			for _, pkg := range response.Packages {

				hash := pkg.Hash()

				if hash != "" {
					hashes = append(hashes, hash)
				}

			}

			if len(hashes) > 0 {
				result = response.Type + "-Packages-" + strings.Join(utils_strings.Unique(hashes), ",")
			}

		} else if len(response.Programs) > 0 {

			hashes := make([]string, 0)

			for _, program := range response.Programs {

				hash := program.Hash()

				if hash != "" {
					hashes = append(hashes, hash)
				}

			}

			if len(hashes) > 0 {
				result = response.Type + "-Programs-" + strings.Join(utils_strings.Unique(hashes), ",")
			}

		} else if len(response.Users) > 0 {

			hashes := make([]string, 0)

			for _, user := range response.Users {

				hash := user.Hash()

				if hash != "" {
					hashes = append(hashes, hash)
				}

			}

			if len(hashes) > 0 {
				result = response.Type + "-Users-" + strings.Join(utils_strings.Unique(hashes), ",")
			}

		}

	}

	return result

}
