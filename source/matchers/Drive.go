package matchers

import "strings"

type Drive struct {
	Name       string `json:"name"`
	Mountpoint string `json:"mountpoint"`
	Type       string `json:"type"`
}

func NewDrive() Drive {

	var drive Drive

	drive.Name = Any
	drive.Mountpoint = Any
	drive.Type = Any

	return drive

}

func ToDrive(value string) Drive {

	var drive Drive

	drive.Type = Any

	if strings.HasPrefix(value, "/") {

		drive.Name = Any
		drive.SetMountpoint(value)

	} else if value != "" {

		drive.SetName(value)
		drive.Mountpoint = Any

	}

	return drive

}

func (drive *Drive) Equal(value Drive) bool {

	if drive.Name == value.Name &&
		drive.Mountpoint == value.Mountpoint &&
		drive.Type == value.Type {
		return true
	}

	return false

}

func (drive *Drive) IsValid() bool {

	if drive.Name != Any ||
		drive.Mountpoint != Any ||
		drive.Type != Any {
		return true
	}

	return false

}

func (drive *Drive) Matches(name string, mountpoint string, typ string) bool {
	return drive.MatchesName(name) && drive.MatchesMountpoint(mountpoint) && drive.MatchesType(typ)
}

func (drive *Drive) MatchesMountpoint(value string) bool {

	if drive.Mountpoint == value {
		return true
	} else if drive.Mountpoint == Any {
		return true
	}

	return false

}

func (drive *Drive) MatchesName(value string) bool {

	if drive.Name == value {
		return true
	} else if drive.Name == Any {
		return true
	}

	return false

}

func (drive *Drive) MatchesType(value string) bool {

	if drive.Type == value {
		return true
	} else if drive.Type == Any {
		return true
	}

	return false

}

func (drive *Drive) SetName(value string) {

	if value == "all" || value == "any" || value == "*" {
		drive.Name = Any
	} else if value != "" {
		drive.Name = strings.TrimSpace(value)
	}

}

func (drive *Drive) SetMountpoint(value string) {

	if value == "all" || value == "any" || value == "*" {
		drive.Mountpoint = Any
	} else if value != "" {
		drive.Mountpoint = strings.TrimSpace(value)
	}

}

func (drive *Drive) SetType(value string) {

	if value == "all" || value == "any" || value == "*" {
		drive.Type = Any
	} else if value != "" {
		drive.Type = strings.TrimSpace(value)
	}

}

func (drive *Drive) Hash() string {

	hash := ""

	if drive.Name != "" {

		hash = toHash(
			drive.Name,
			drive.Mountpoint,
			drive.Type,
		)

	}

	return hash

}
