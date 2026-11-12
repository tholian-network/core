package structs

import "tholian-core/matchers"
import utils_path "tholian-core/utils/path"
import "crypto/sha256"
import "encoding/hex"
import "fmt"
import "os"
import "os/user"
import "path"
import "strconv"
import "strings"
import "syscall"

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

	if file.IsValid() == true {

		if utils_path.IsSafeBasename(file.Name) == true {

			filename := path.Join(file.Folder, file.Name)

			stat, err1 := os.Stat(filename)

			if err1 == nil {

				if stat.IsDir() == false {

					buffer, err2 := os.ReadFile(filename)

					if err2 == nil {

						file.SetBuffer(buffer)
						file.SetExecutable(stat.Mode().Perm()&0111 != 0)

						if stat.Sys() != nil {

							info, ok := stat.Sys().(*syscall.Stat_t)

							if ok == true {

								account, err3 := user.LookupId(strconv.FormatUint(uint64(info.Uid), 10))

								if err3 == nil {

									owner := matchers.NewUser()

									owner.SetName(account.Username)
									owner.SetType("user")

									file.SetUser(owner)

								}

							}

						}

						return nil

					} else {
						return err2
					}

				} else {
					return fmt.Errorf("Invalid File")
				}

			} else {
				return err1
			}

		} else {
			return fmt.Errorf("Invalid File")
		}

	} else {
		return fmt.Errorf("Invalid File")
	}

}

func (file *File) SetBuffer(value []byte) {

	if len(value) > 0 {
		file.Buffer = append([]byte(nil), value...)
	} else {
		file.Buffer = make([]byte, 0)
	}

	checksum := sha256.Sum256(file.Buffer)
	file.Digest = hex.EncodeToString(checksum[:])

}

func (file *File) SetExecutable(value bool) {
	file.Executable = value
}

func (file *File) SetFolder(value string) {

	if utils_path.IsFolder(value) {
		file.Folder = strings.TrimSpace(value)
	}

}

func (file *File) SetName(value string) {

	value = strings.TrimSpace(value)

	if utils_path.IsSafeBasename(value) {
		file.Name = value
	}

}

func (file *File) SetUser(value matchers.User) {

	if value.IsValid() == true {
		file.User = value
	}

}

func (file *File) Write() error {

	if file.IsValid() == true {

		if utils_path.IsSafeBasename(file.Name) == true {

			checksum := sha256.Sum256(file.Buffer)
			file.Digest = hex.EncodeToString(checksum[:])

			filename := path.Join(file.Folder, file.Name)
			mode := os.FileMode(0644)

			if file.Executable == true {
				mode = os.FileMode(0755)
			}

			err1 := os.WriteFile(filename, file.Buffer, mode)

			if err1 == nil {

				err2 := os.Chmod(filename, mode)

				if err2 == nil {
					return nil
				} else {
					return err2
				}

			} else {
				return err1
			}

		} else {
			return fmt.Errorf("Invalid File")
		}

	} else {
		return fmt.Errorf("Invalid File")
	}

}

func (file *File) Hash() string {

	hash := ""

	if file.Name != "" {

		hash = toHash(
			file.Name,
			file.Folder,
			file.User.Hash(),
			strconv.FormatBool(file.Executable),
			file.Digest,
		)

	}

	return hash

}

