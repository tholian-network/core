package methods

import "strings"

// Method is a project-qualified RPC method name.
//
// The wire format is "<go-module>.<ExportedActionName>", for example
// "tholian-firewall.Check" or "tholian-scout.CollectPackages". Action method
// values map 1:1 to exported functions in the owning project's actions package.
type Method string

func (method Method) String() string {
	return string(method)
}

func (method Method) IsValid() bool {

	for _, candidate := range all_methods {
		if candidate == method {
			return true
		}
	}

	return false

}

func (method Method) Project() string {

	project, _, ok := strings.Cut(string(method), ".")

	if ok == true {
		return project
	}

	return ""

}

func (method Method) Action() string {

	_, action, ok := strings.Cut(string(method), ".")

	if ok == true {
		return action
	}

	return ""

}

func ParseMethod(value string) (Method, bool) {

	method := Method(value)

	if method.IsValid() == true {
		return method, true
	}

	return Method(""), false

}
