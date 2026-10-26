package structs

import "strings"

type Drive struct {
	Name       string `json:"name"`
	Size       uint64 `json:"size"`
	Free       uint64 `json:"free"`
	Used       uint64 `json:"used"`
	Mountpoint string `json:"mountpoint"`
	Type       string `json:"type"`
}

func NewDrive(name string, typ string) Drive {

	var drive Drive

	drive.SetName(name)
	drive.SetType(typ)

	return drive

}

func (drive *Drive) IsValid() bool {

	if drive.Name != "" && strings.HasPrefix(drive.Mountpoint, "/") {

		if drive.Size > 0 {

			if drive.Free != 0 && drive.Used != 0 {
				return true
			}

		}

	}

	return false

}

func (drive *Drive) SetMountpoint(value string) {
	drive.Mountpoint = strings.TrimSpace(value)
}

func (drive *Drive) SetName(value string) {
	drive.Name = strings.TrimSpace(value)
}

func (drive *Drive) SetSize(value uint64) {
	drive.Size = value
}

func (drive *Drive) SetFree(value uint64) {
	drive.Free = value
}

func (drive *Drive) SetUsed(value uint64) {
	drive.Used = value
}

func (drive *Drive) SetType(value string) {

	if value == "local" {
		drive.Type = "local"
	} else if value == "remote" {
		drive.Type = "remote"
	}

}
