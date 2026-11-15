package scout

import "tholian-core/matchers"

type UpdatePackages struct {
	System   string             `json:"system"`
	Searches []matchers.Package `json:"searches"`
}
