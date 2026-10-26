package matchers

import "encoding/binary"
import "encoding/hex"
import "hash/crc32"
import "strings"

type Program struct {
	Name      string   `json:"name"`
	Command   string   `json:"command"`
	Arguments []string `json:"arguments"`
}

func NewProgram() Program {

	var program Program

	program.Name = Any
	program.Command = Any
	program.Arguments = make([]string, 0)

	return program

}

func ToProgram(value string) Program {

	var program Program

	program.Name = Any
	program.Command = Any
	program.Arguments = make([]string, 0)

	if strings.HasPrefix(value, "/") {
		program.Name = Any
		program.SetCommand(value)
	} else {
		program.SetName(value)
		program.Command = Any
	}

	return program

}

func (program *Program) Equal(value Program) bool {

	if program.Name == value.Name && program.Command == value.Command {

		if len(program.Arguments) == len(value.Arguments) {

			matches_args := true

			for a, arg := range program.Arguments {

				if arg == value.Arguments[a] {
					// Do Nothing
				} else {
					matches_args = false
					break
				}

			}

			if matches_args {
				return true
			}

		}

	}

	return false

}

func (program *Program) IsValid() bool {

	if program.Name != Any ||
		program.Command != Any {
		return true
	}

	return false

}

func (program *Program) Matches(name string, command string, arguments []string) bool {
	return program.MatchesName(name) && program.MatchesCommand(command) && program.MatchesArguments(arguments)
}

func (program *Program) MatchesArguments(values []string) bool {

	if len(program.Arguments) == len(values) {

		matches_args := true

		for a, arg := range program.Arguments {

			if arg == values[a] {
				// Do Nothing
			} else if arg == "any" {
				// Do Nothing
			} else {
				matches_args = false
				break
			}

		}

		if matches_args {
			return true
		}

	} else if len(program.Arguments) == 0 {
		return true
	}

	return false

}

func (program *Program) MatchesCommand(value string) bool {

	if program.Command == value {
		return true
	} else if program.Command == Any {
		return true
	}

	return false

}

func (program *Program) MatchesName(value string) bool {

	if program.Name == value {
		return true
	} else if program.Name == Any {
		return true
	}

	return false

}

func (program *Program) SetArguments(values []string) {
	program.Arguments = values
}

func (program *Program) SetCommand(value string) {

	if value == "all" || value == "any" || value == "*" {
		program.Command = Any
	} else if value != "" {
		program.Command = strings.TrimSpace(value)
	}

}

func (program *Program) SetName(value string) {

	if value == "all" || value == "any" || value == "*" {
		program.Name = Any
	} else if strings.Contains(value, "/") {
		program.Name = strings.TrimSpace(value[0:strings.LastIndex(value, "/")])
	} else if value != "" {
		program.Name = strings.TrimSpace(value)
	}

}

func (program *Program) Hash() string {

	hash := ""

	if program.Name != "" {

		checksum := crc32.ChecksumIEEE([]byte(strings.Join([]string{
			program.Name,
			program.Command,
		}, "-")))

		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, checksum)
		hash = hex.EncodeToString(tmp)

	}

	return hash

}
