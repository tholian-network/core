package parameters

import "encoding/json"
import "tholian-core/protocols/tholian/methods"

type MethodCall struct {
	Systems    []string        `json:"systems"`
	Method     methods.Method  `json:"method"`
	Parameters json.RawMessage `json:"parameters"`
	Grant      string          `json:"grant"`
	Purpose    string          `json:"purpose"`
}
