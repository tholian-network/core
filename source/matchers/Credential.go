package matchers

import "strings"

type Credential struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

func NewCredential() Credential {

	var credential Credential

	credential.Name = Any
	credential.Password = Any
	credential.Type = Any

	return credential

}

func ToCredential(value string) Credential {

	var credential Credential

	credential.Name = Any
	credential.Password = Any
	credential.Type = Any

	credential.SetName(value)

	return credential

}

func (credential *Credential) Equal(value Credential) bool {

	if credential.Name == value.Name &&
		credential.Password == value.Password &&
		credential.Type == value.Type {
		return true
	}

	return false

}

func (credential *Credential) IsValid() bool {

	if credential.Name != Any ||
		credential.Password != Any ||
		credential.Type != Any {
		return true
	}

	return false

}

func (credential *Credential) Matches(name string, password string, typ string) bool {
	return credential.MatchesName(name) && credential.MatchesPassword(password) && credential.MatchesType(typ)
}

func (credential *Credential) MatchesName(value string) bool {

	if credential.Name == value {
		return true
	} else if credential.Name == Any {
		return true
	}

	return false

}

func (credential *Credential) MatchesPassword(value string) bool {

	if credential.Password == value {
		return true
	} else if credential.Password == Any {
		return true
	}

	return false

}

func (credential *Credential) MatchesType(value string) bool {

	if credential.Type == value {
		return true
	} else if credential.Type == Any {
		return true
	}

	return false

}

func (credential *Credential) SetName(value string) {
	credential.Name = strings.TrimSpace(value)
}

func (credential *Credential) SetPassword(value string) {
	credential.Password = value
}

func (credential *Credential) SetType(value string) {

	if value == "all" || value == "any" || value == "*" {
		credential.Type = Any
	} else if value != "" {
		credential.Type = value
	}

}

func (credential *Credential) Hash() string {

	hash := ""

	if credential.Name != "" {

		hash = toHash(
			credential.Name,
			credential.Password,
			credential.Type,
		)

	}

	return hash

}
