package protocol

import "tholian-core/matchers"

type Scope struct {
	Any         bool                  `json:"any,omitempty"`
	Packages    []matchers.Package    `json:"packages,omitempty"`
	Connections []matchers.Connection `json:"connections,omitempty"`
	Networks    []matchers.Network    `json:"networks,omitempty"`
	Programs    []matchers.Program    `json:"programs,omitempty"`
	Users       []matchers.User       `json:"users,omitempty"`
	Updates     []matchers.Update     `json:"updates,omitempty"`
	Antiques    []matchers.Antique    `json:"antiques,omitempty"`
	Devices     []matchers.Device     `json:"devices,omitempty"`
}

func NewScope() Scope {
	return Scope{}
}

func NewUnrestrictedScope() Scope {
	return Scope{Any: true}
}

func (scope *Scope) IsEmpty() bool {

	if scope == nil {
		return true
	}

	if scope.Any == true {
		return false
	}

	if len(scope.Packages) > 0 ||
		len(scope.Connections) > 0 ||
		len(scope.Networks) > 0 ||
		len(scope.Programs) > 0 ||
		len(scope.Users) > 0 ||
		len(scope.Updates) > 0 ||
		len(scope.Antiques) > 0 ||
		len(scope.Devices) > 0 {
		return false
	}

	return true

}

func (scope *Scope) IsUnrestricted() bool {

	if scope == nil {
		return false
	}

	return scope.Any

}

func (scope *Scope) AllowsPackages(values []matchers.Package) bool {

	if scope == nil || len(values) == 0 {
		return false
	}

	if scope.Any == true {
		return true
	}

	for _, value := range values {

		found := false

		for _, allowed := range scope.Packages {

			if allowed.Matches(value.Name, value.Version, value.Architecture, value.Manager, value.Vendor) {
				found = true
				break
			}

		}

		if found == false {
			return false
		}

	}

	return len(scope.Packages) > 0

}

func (scope *Scope) AllowsConnections(values []matchers.Connection) bool {

	if scope == nil || len(values) == 0 {
		return false
	}

	if scope.Any == true {
		return true
	}

	for _, value := range values {

		found := false

		for _, allowed := range scope.Connections {

			if allowed.Matches(value.Host, value.Port, value.Protocol, value.Type) {
				found = true
				break
			}

		}

		if found == false {
			return false
		}

	}

	return len(scope.Connections) > 0

}

func (scope *Scope) AllowsNetworks(values []matchers.Network) bool {

	if scope == nil || len(values) == 0 {
		return false
	}

	if scope.Any == true {
		return true
	}

	for _, value := range values {

		found := false

		for _, allowed := range scope.Networks {

			if allowed.Matches(value.Name, value.Subnet) {
				found = true
				break
			}

		}

		if found == false {
			return false
		}

	}

	return len(scope.Networks) > 0

}

func (scope *Scope) AllowsPrograms(values []matchers.Program) bool {

	if scope == nil || len(values) == 0 {
		return false
	}

	if scope.Any == true {
		return true
	}

	for _, value := range values {

		found := false

		for _, allowed := range scope.Programs {

			if allowed.Matches(value.Name, value.Command, value.Arguments) {
				found = true
				break
			}

		}

		if found == false {
			return false
		}

	}

	return len(scope.Programs) > 0

}

func (scope *Scope) AllowsUsers(values []matchers.User) bool {

	if scope == nil || len(values) == 0 {
		return false
	}

	if scope.Any == true {
		return true
	}

	for _, value := range values {

		found := false

		for _, allowed := range scope.Users {

			if allowed.Matches(value.Name, value.Password, value.Type) {
				found = true
				break
			}

		}

		if found == false {
			return false
		}

	}

	return len(scope.Users) > 0

}

func (scope *Scope) AllowsUpdates(values []matchers.Update) bool {

	if scope == nil || len(values) == 0 {
		return false
	}

	if scope.Any == true {
		return true
	}

	for _, value := range values {

		found := false

		for _, allowed := range scope.Updates {

			if allowed.Matches(value.Name, value.Version, value.Architecture, value.Manager, value.Vendor) {
				found = true
				break
			}

		}

		if found == false {
			return false
		}

	}

	return len(scope.Updates) > 0

}
