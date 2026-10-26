package matchers

import "strings"

type Manager struct {
	Name string `json:"name"`
}

func NewManager() Manager {

	var manager Manager

	manager.Name = Any

	return manager

}

func ToManager(value string) Manager {

	var manager Manager

	manager.Name = Any

	manager.SetName(value)

	return manager

}

func (manager *Manager) Equal(value Manager) bool {

	if manager.Name == value.Name {
		return true
	}

	return false

}

func (manager *Manager) IsValid() bool {

	if manager.Name != Any {
		return true
	}

	return false

}

func (manager *Manager) Matches(name string) bool {
	return manager.MatchesName(name)
}

func (manager *Manager) MatchesName(value string) bool {

	if manager.Name == value {
		return true
	} else if manager.Name == Any {
		return true
	}

	return false

}

func (manager *Manager) SetName(value string) {
	manager.Name = strings.TrimSpace(value)
}

func (manager *Manager) Hash() string {

	hash := ""

	if manager.Name != "" {
		hash = manager.Name
	}

	return hash

}
