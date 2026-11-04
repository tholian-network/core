package structs

import "tholian-core/matchers"
import utils_path "tholian-core/utils/path"

type File struct {
	Name       string        `json:"name"`
	Folder     string        `json:"folder"`
	User       matchers.User `json:"user"`
	Executable bool          `json:"executable"`
	Buffer     []byte        `json:"buffer"`
	Digest     string        `json:"digest"`
}

func NewFile(name string, folder string) File {

	var file File

	file.SetName(name)
	file.SetFolder(folder)
	file.User = matchers.NewUser()
	file.Buffer = make([]byte, 0)

	return file

}

func (file *File) Equal(value File) bool {

	if file.Name == value.Name &&
		file.Folder == value.Folder &&
		file.Digest == value.Digest {
		return true
	}

	return false

}

func (file *File) IsExecutable() bool {
	return file.Executable == true
}

func (file *File) IsValid() bool {

	if file.Name != "" && file.Folder != "" {
		return true
	}

	return false

}

func (file *File) Read() error {

	if file.IsValid() {

		// TODO: Read file.Buffer
		// TODO: Recalculate hash
		// TODO: os.Stat() for file.Executable
		// TODO: Read matchers.User (Name, Group, Type)

	}

}

func (file *File) SetBuffer(value []byte) {

	file.Buffer = value
	// TODO: Recalculate hash

}

func (file *File) SetExecutable(value bool) {
	file.Executable = value
}

func (file *File) SetFolder(value string) {

	if utils_path.IsFolder(value) {
		file.Folder = strings.TrimSpace(value)
	}

}

func (file *File) SetUser(value matchers.User) {

	if value.IsValid() {
		file.User = value
	}

}

func (file *File) Write() error {

	// TODO: Write to file.Folder + "/" + file.Name

}
