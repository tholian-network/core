package types

import "net/url"
import "strconv"
import "strings"

type User struct {
	ID       uint16  `json:"id"`
	Name     string  `json:"name"`
	Password string  `json:"password"`
	Folder   string  `json:"folder"`
	Groups   []Group `json:"groups"`
	Shell    string  `json:"shell"`
	Type     string  `json:"type"`
}

func NewUser() User {

	var user User

	user.Groups = make([]Group, 0)

	return user

}

func IsUser(value string) bool {

	var user User

	user.Groups = make([]Group, 0)

	return value != "" && user.Parse(value)

}

func ParseUser(value string) *User {

	var result *User = nil
	var user User

	user.Groups = make([]Group, 0)

	if value != "" && user.Parse(value) == true {
		result = &user
	}

	return result

}

func ToUser(value string) User {

	var user User

	user.Groups = make([]Group, 0)

	user.Parse(value)

	return user

}

func (user *User) IsValid() bool {

	if user.ID == 0 && user.Name != "root" {
		return false
	} else if user.Name != "" {
		return true
	}

	return false

}

func (user *User) SetID(value uint16) {

	// User "nobody" is last user id
	if value >= 0 && value <= 65534 {
		user.ID = value
	}

}

func (user *User) SetName(value string) {
	user.Name = strings.TrimSpace(value)
}

func (user *User) SetPassword(value string) {
	user.Password = value
}

func (user *User) SetFolder(value string) {

	if strings.HasPrefix(value, "/") {
		user.Folder = value
	}

}

func (user *User) AddGroup(value Group) {

	if value.IsValid() {

		found := false

		for g := 0; g < len(user.Groups); g++ {

			if user.Groups[g].ID == value.ID {
				found = true
				break
			}

		}

		if found == false {
			user.Groups = append(user.Groups, value)
		}

	}

}

func (user *User) RemoveGroup(value Group) {

	index := -1

	for g := 0; g < len(user.Groups); g++ {

		if user.Groups[g].ID == value.ID {
			index = g
			break
		}

	}

	if index != -1 {
		user.Groups = append(user.Groups[:index], user.Groups[index+1:]...)
	}

}

func (user *User) SetGroups(values []Group) {

	filtered := make([]Group, 0)

	for _, value := range values {

		if value.IsValid() {
			filtered = append(filtered, value)
		}

	}

	user.Groups = filtered

}

func (user *User) SetShell(value string) {

	if strings.HasPrefix(value, "/") {
		user.Shell = strings.TrimSpace(value)
	}

}

func (user *User) SetType(value string) {

	if value == "user" {
		user.Type = "user"
	} else if value == "program" {
		user.Type = "program"
	}

}

func (user *User) Parse(value string) bool {

	parts := strings.SplitN(value, ":", 7)

	if len(parts) >= 1 {

		name, err := url.QueryUnescape(parts[0])

		if err == nil {
			user.SetName(name)
		} else {
			return false
		}

	}

	if len(parts) >= 2 {

		typ, err := url.QueryUnescape(parts[1])

		if err == nil {
			user.SetType(typ)
		} else {
			return false
		}

	}


	if len(parts) >= 3 {

		id, err := strconv.ParseUint(parts[2], 10, 16)

		if err == nil {

			if id == 65535 {
				user.ID = uint16(id)
			} else {
				user.SetID(uint16(id))
			}

		} else {
			return false
		}

	}

	if len(parts) >= 4 {

		folder, err := url.QueryUnescape(parts[3])

		if err == nil {
			user.SetFolder(folder)
		} else {
			return false
		}

	}

	if len(parts) >= 5 {

		shell, err := url.QueryUnescape(parts[4])

		if err == nil {
			user.SetShell(shell)
		} else {
			return false
		}

	}

	if len(parts) >= 6 {

		password, err := url.QueryUnescape(parts[5])

		if err == nil {
			user.SetPassword(password)
		} else {
			return false
		}

	}

	if len(parts) >= 7 {

		if parts[6] != "" {

			groups := make([]Group, 0)

			for _, value := range strings.Split(parts[6], ",") {

				group := ParseGroup(value)

				if group != nil {
					groups = append(groups, *group)
				}

			}

			user.SetGroups(groups)

		}

	}

	return true

}

func (user User) String() string {

	groups := make([]string, 0, len(user.Groups))

	for g := 0; g < len(user.Groups); g++ {
		groups = append(groups, user.Groups[g].String())
	}

	return strings.Join([]string{
		url.QueryEscape(user.Name),
		url.QueryEscape(user.Type),
		strconv.FormatUint(uint64(user.ID), 10),
		url.QueryEscape(user.Folder),
		url.QueryEscape(user.Shell),
		url.QueryEscape(user.Password),
		strings.Join(groups, ","),
	}, ":")

}

func (user User) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(user.String())), nil
}

func (user *User) UnmarshalJSON(data []byte) error {

	unquoted, err := strconv.Unquote(string(data))

	if err != nil {
		return err
	}

	check := ToUser(unquoted)

	if check.IsValid() {

		user.ID = check.ID
		user.Name = check.Name
		user.Password = check.Password
		user.Folder = check.Folder
		user.Groups = check.Groups
		user.Shell = check.Shell
		user.Type = check.Type

	}

	return nil

}

func (user User) Hash() string {

	values := make([]string, 0)

	values = append(values, strconv.FormatUint(uint64(user.ID), 10))
	values = append(values, user.Name)
	values = append(values, user.Password)
	values = append(values, user.Folder)
	values = append(values, user.Shell)
	values = append(values, user.Type)

	for g := 0; g < len(user.Groups); g++ {
		values = append(values, user.Groups[g].Hash())
	}

	return toHash(values...)

}
