package guard

import "tholian-core/structs"

type SearchVulnerabilities struct {
	Vulnerabilities []structs.Vulnerability `json:"vulnerabilities"`
}
