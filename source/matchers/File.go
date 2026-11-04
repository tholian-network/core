package matchers

import "strings"

type File struct {
	Name   string `json:"name"`
	Folder string `json:"folder"`
	User   string `json:"user"`
	Digest string `json:"digest"`
}

func NewFile() File {

	var file File

	file.Name = Any
	file.Folder = Any
	file.User = Any
	file.Digest = Any

	return file

}

func ToFile(value string) File {

	var file File

	file.Name = Any
	file.Folder = Any
	file.User = Any
	file.Digest = Any

	file.SetName(value)

	return file

}

func (file *File) Equal(value File) bool {

	if file.Name == value.Name &&
		file.Folder == value.Folder &&
		file.User == value.User &&
		file.Digest == value.Digest {
		return true
	}

	return false

}

func (file *File) IsValid() bool {

	if file.Name != Any ||
		file.Folder != Any ||
		file.User != Any ||
		file.Digest != Any {
		return true
	}

	return false

}

func (file *File) Matches(name string, folder string, user string, hash string) bool {
	return file.MatchesName(name) && file.MatchesFolder(folder) && file.MatchesUser(user) && file.MatchesDigest(hash)
}

func (file *File) MatchesName(value string) bool {

	if file.Name == value {
		return true
	} else if file.Name == Any {
		return true
	}

	return false

}

func (file *File) MatchesFolder(value string) bool {

	if file.Folder == value {
		return true
	} else if strings.Contains(value, file.Folder) {
		return true
	} else if file.Folder == Any {
		return true
	}

	return false

}

func (file *File) MatchesUser(value string) bool {

	if file.User == value {
		return true
	} else if file.User == Any {
		return true
	}

	return false

}

func (file *File) MatchesDigest(value string) bool {

	if file.Digest == value {
		return true
	} else if file.Digest == Any {
		return true
	}

	return false

}

func (file *File) SetName(value string) {
	file.Name = value
}

func (file *File) SetFolder(value string) {

	if value == "all" || value == "any" || value == "*" {
		file.Folder = Any
	} else if value != "" {
		file.Folder = value
	}

}

func (file *File) SetUser(value string) {

	if value == "all" || value == "any" || value == "*" {
		file.User = Any
	} else if value != "" {
		file.User = value
	}

}

func (file *File) SetDigest(value string) {

	if value == "all" || value == "any" || value == "*" {
		file.Digest = Any
	} else if value != "" {
		file.Digest = value
	}

}

func (file *File) Hash() string {

	hash := toHash(
		file.Name,
		file.Folder,
		file.User,
		file.Digest,
	)

	return hash

}
