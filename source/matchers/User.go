package matchers

import "encoding/binary"
import "encoding/hex"
import "hash/crc32"
import "strings"

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

func NewUser() User {

	var user User

	user.Name = Any
	user.Password = Any
	user.Type = Any

	return user

}

func ToUser(value string) User {

	var user User

	user.Name = Any
	user.Password = Any
	user.Type = Any

	user.SetName(value)

	return user

}

func (user *User) Equal(value User) bool {

	if user.Name == value.Name &&
		user.Password == value.Password &&
		user.Type == value.Type {
		return true
	}

	return false

}

func (user *User) IsValid() bool {

	if user.Name != Any ||
		user.Password != Any ||
		user.Type != Any {
		return true
	}

	return false

}

func (user *User) Matches(name string, password string, typ string) bool {
	return user.MatchesName(name) && user.MatchesPassword(password) && user.MatchesType(typ)
}

func (user *User) MatchesName(value string) bool {

	if user.Name == value {
		return true
	} else if user.Name == Any {
		return true
	}

	return false

}

func (user *User) MatchesPassword(value string) bool {

	if user.Password == value {
		return true
	} else if user.Password == Any {
		return true
	}

	return false

}

func (user *User) MatchesType(value string) bool {

	if user.Type == value {
		return true
	} else if user.Type == Any {
		return true
	}

	return false

}

func (user *User) SetName(value string) {
	user.Name = strings.TrimSpace(value)
}

func (user *User) SetPassword(value string) {

	if value == "all" || value == "any" || value == "*" {
		user.Password = Any
	} else if value != "" {
		user.Password = value
	}

}

func (user *User) SetType(value string) {

	if value == "all" || value == "any" || value == "*" {
		user.Type = Any
	} else if value == "user" {
		user.Type = "user"
	} else if value == "program" {
		user.Type = "program"
	}

}

func (user *User) Hash() string {

	hash := ""

	if user.Name != "" {

		checksum := crc32.ChecksumIEEE([]byte(strings.Join([]string{
			user.Name,
			user.Password,
			user.Type,
		}, "-")))

		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, checksum)
		hash = hex.EncodeToString(tmp)

	}

	return hash

}
