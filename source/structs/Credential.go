package structs

import "bytes"
import "encoding/hex"
import "strings"

type Credential struct {
	Name      string   `json:"name"`
	Keys      [][]byte `json:"keys"`
	Passwords []string `json:"passwords"`
	Type      string   `json:"type"`
}

func NewCredential(name string, typ string) Credential {

	var credential Credential

	credential.SetName(name)
	credential.SetType(typ)
	credential.Keys = make([][]byte, 0)
	credential.Passwords = make([]string, 0)

	return credential

}

func (credential *Credential) Equal(value Credential) bool {

	if credential.Name != value.Name || credential.Type != value.Type {
		return false
	}

	if len(credential.Keys) != len(value.Keys) {
		return false
	}

	for k := 0; k < len(credential.Keys); k++ {

		if bytes.Equal(credential.Keys[k], value.Keys[k]) == false {
			return false
		}

	}

	if len(credential.Passwords) != len(value.Passwords) {
		return false
	}

	for p := 0; p < len(credential.Passwords); p++ {

		if credential.Passwords[p] != value.Passwords[p] {
			return false
		}

	}

	return true

}

func (credential *Credential) IsValid() bool {

	if credential.Name != "" && credential.Type != "" {

		if len(credential.Keys) > 0 || len(credential.Passwords) > 0 {
			return true
		}

	}

	return false

}

func (credential *Credential) SetName(value string) {
	credential.Name = strings.TrimSpace(value)
}

func (credential *Credential) AddKey(value []byte) {

	found := false

	for k := 0; k < len(credential.Keys); k++ {

		if bytes.Equal(credential.Keys[k], value) {
			found = true
			break
		}

	}

	if found == false {
		credential.Keys = append(credential.Keys, value)
	}

}

func (credential *Credential) RemoveKey(value []byte) {

	index := -1

	for k := 0; k < len(credential.Keys); k++ {

		if bytes.Equal(credential.Keys[k], value) {
			index = k
			break
		}

	}

	if index != -1 {
		credential.Keys = append(credential.Keys[:index], credential.Keys[index+1:]...)
	}

}

func (credential *Credential) SetKeys(values [][]byte) {

	filtered := make([][]byte, 0)

	for _, value := range values {
		filtered = append(filtered, value)
	}

	credential.Keys = filtered

}

func (credential *Credential) AddPassword(value string) {

	value = strings.TrimSpace(value)

	found := false

	for p := 0; p < len(credential.Passwords); p++ {

		if credential.Passwords[p] == value {
			found = true
			break
		}

	}

	if found == false {
		credential.Passwords = append(credential.Passwords, value)
	}

}

func (credential *Credential) RemovePassword(value string) {

	value = strings.TrimSpace(value)

	index := -1

	for p := 0; p < len(credential.Passwords); p++ {

		if credential.Passwords[p] == value {
			index = p
			break
		}

	}

	if index != -1 {
		credential.Passwords = append(credential.Passwords[:index], credential.Passwords[index+1:]...)
	}

}

func (credential *Credential) SetPasswords(values []string) {

	filtered := make([]string, 0)

	for _, value := range values {
		filtered = append(filtered, strings.TrimSpace(value))
	}

	credential.Passwords = filtered

}

func (credential *Credential) SetType(value string) {

	if value == "user" {
		credential.Type = "user"
	} else if value == "program" {
		credential.Type = "program"
	}

}

func (credential *Credential) Hash() string {

	hash := ""

	if credential.Name != "" {

		values := make([]string, 0)

		values = append(values, credential.Name)
		values = append(values, credential.Type)

		for k := 0; k < len(credential.Keys); k++ {
			values = append(values, hex.EncodeToString(credential.Keys[k]))
		}

		for p := 0; p < len(credential.Passwords); p++ {
			values = append(values, credential.Passwords[p])
		}

		hash = toHash(values...)

	}

	return hash

}
