package guard

import "tholian-core/structs"

type CollectVulnerabilities struct {
	Vulnerabilities []structs.Vulnerability `json:"vulnerabilities"`
}
