package guard

import "tholian-core/matchers"

type ExecuteProgram struct {
	System  string           `json:"system"`
	Program matchers.Program `json:"program"`
}
