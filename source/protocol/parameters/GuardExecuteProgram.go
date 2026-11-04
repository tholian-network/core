package parameters

import "tholian-core/matchers"

type GuardExecuteProgram struct {
	System  string           `json:"system"`
	Program matchers.Program `json:"program"`
}
