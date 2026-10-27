package types

import "net/url"
import "strconv"
import "strings"

type Group struct {
	ID       uint16 `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

func NewGroup() Group {

	var group Group

	group.SetType("user")

	return group

}

func IsGroup(value string) bool {

	var group Group

	return value != "" && group.Parse(value)

}

func ParseGroup(value string) *Group {

	var result *Group = nil
	var group Group

	if value != "" && group.Parse(value) == true {
		result = &group
	}

	return result

}

func ToGroup(value string) Group {

	var group Group

	group.Parse(value)

	return group

}

func (group *Group) IsValid() bool {

	if group.ID == 0 && group.Name != "root" {
		return false
	} else if group.Name != "" {
		return true
	}

	return false

}

func (group *Group) SetID(value uint16) {

	// User "nobody" is last user id
	if value >= 0 && value <= 65534 {
		group.ID = value
	}

}

func (group *Group) SetName(value string) {
	group.Name = strings.TrimSpace(value)
}

func (group *Group) SetPassword(value string) {
	group.Password = value
}

func (group *Group) SetType(value string) {

	if value == "user" {
		group.Type = "user"
	} else if value == "program" {
		group.Type = "program"
	}

}

func (group *Group) Parse(value string) bool {

	parts := strings.SplitN(value, ":", 4)

	if len(parts) >= 1 {

		name, err := url.QueryUnescape(parts[0])

		if err == nil {
			group.SetName(name)
		} else {
			return false
		}

	}

	if len(parts) >= 2 {

		typ, err := url.QueryUnescape(parts[1])

		if err == nil {
			group.SetType(typ)
		} else {
			return false
		}

	}

	if len(parts) >= 3 {

		id, err := strconv.ParseUint(parts[2], 10, 16)

		if err == nil {

			if id == 65535 {
				group.ID = uint16(id)
			} else {
				group.SetID(uint16(id))
			}

		} else {
			return false
		}

	}

	if len(parts) >= 4 {

		password, err := url.QueryUnescape(parts[3])

		if err == nil {
			group.SetPassword(password)
		} else {
			return false
		}

	}

	return true

}

func (group Group) String() string {

	return strings.Join([]string{
		url.QueryEscape(group.Name),
		url.QueryEscape(group.Type),
		strconv.FormatUint(uint64(group.ID), 10),
		url.QueryEscape(group.Password),
	}, ":")

}

func (group Group) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(group.String())), nil
}

func (group *Group) UnmarshalJSON(data []byte) error {

	unquoted, err := strconv.Unquote(string(data))

	if err != nil {
		return err
	}

	check := ToGroup(unquoted)

	if check.IsValid() {

		group.ID = check.ID
		group.Name = check.Name
		group.Password = check.Password
		group.Type = check.Type

	}

	return nil

}

func (group Group) Hash() string {

	return toHash(
		strconv.FormatUint(uint64(group.ID), 10),
		group.Name,
		group.Password,
		group.Type,
	)

}
