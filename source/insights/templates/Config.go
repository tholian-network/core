package templates

import "strings"

type Config struct {

	// Cache and FilesystemCache
	Name     string
	Package  string
	BuildTag string
	Struct   string

	// FilesystemCache
	Folder string
	Files  []string

}

func (config *Config) IsValid() bool {

	if config.Package != "" && strings.ToLower(config.Package) == config.Package {

		if config.Name != "" && config.Name == strings.ToUpper(config.Name[0:1])+strings.ToLower(config.Name[1:]) {

			if strings.HasPrefix(config.Struct, "structs.") {

				if config.Folder != "" && len(config.Files) > 0 {

					valid_filesystem := true

					for _, file := range config.Files {

						if !strings.HasPrefix(file, config.Folder + "/") {
							valid_filesystem = false
							break
						}

					}

					if valid_filesystem == true {
						return true
					}

				} else {
					return true
				}

			}

		}

	}

	return false

}
