package structs

import "tholian-core/types"
import "strings"

type Distribution struct {
	Name               string             `json:"name"`
	Version            string             `json:"version"`
	Kernel             string             `json:"kernel"`
	KernelArchitecture string             `json:"kernel_architecture"`
	KernelModules      []string           `json:"kernel_modules"`
	KernelVersion      string             `json:"kernel_version"`
	Manager            string             `json:"manager"`
	Vendor             string             `json:"vendor"`
	Keywords           *map[string]string `json:"keywords,omitempty"`
}

func NewDistribution() Distribution {

	var distribution Distribution

	distribution.Kernel = "any"
	distribution.KernelModules = make([]string, 0)
	distribution.KernelArchitecture = "any"
	distribution.KernelVersion = "any"
	distribution.Manager = "any"

	return distribution

}

func (distribution *Distribution) Equal(value Distribution) bool {

	if distribution.Name == value.Name &&
		distribution.Version == value.Version &&
		distribution.Vendor == value.Vendor &&
		distribution.Kernel == value.Kernel {
		return true
	}

	return false

}

func (distribution *Distribution) IsValid() bool {

	if distribution.Name != "" {
		return true
	}

	return false

}

func (distribution *Distribution) SetKernel(value string) {

	if value == "android" || value == "Android" {
		distribution.Kernel = "android"
	} else if value == "darwin" || value == "Darwin" {
		distribution.Kernel = "darwin"
	} else if value == "dragonfly" || value == "DragonFly" {
		distribution.Kernel = "dragonfly"
	} else if value == "freebsd" || value == "FreeBSD" {
		distribution.Kernel = "freebsd"
	} else if value == "hurd" || value == "GNU" {
		distribution.Kernel = "hurd"
	} else if value == "illumos" {
		distribution.Kernel = "illumos"
	} else if value == "linux" || value == "Linux" {
		distribution.Kernel = "linux"
	} else if value == "netbsd" || value == "NetBSD" {
		distribution.Kernel = "netbsd"
	} else if value == "openbsd" || value == "OpenBSD" {
		distribution.Kernel = "openbsd"
	} else if value == "solaris" || value == "Solaris" || value == "SunOS" {
		distribution.Kernel = "solaris"
	} else if value == "Windows" || value == "MINGW32_NT" {
		distribution.Kernel = "windows"
	}

}

func (distribution *Distribution) SetKernelArchitecture(value string) {

	architecture := types.ParseArchitecture(value)

	if architecture != nil {
		distribution.KernelArchitecture = architecture.String()
	}

}

func (distribution *Distribution) AddKernelModule(value string) {

	found := false

	for k := 0; k < len(distribution.KernelModules); k++ {

		if distribution.KernelModules[k] == value {
			found = true
			break
		}

	}

	if found == false {
		distribution.KernelModules = append(distribution.KernelModules, value)
	}

}

func (distribution *Distribution) HasKernelModule(value string) bool {

	for k := 0; k < len(distribution.KernelModules); k++ {

		if distribution.KernelModules[k] == value {
			return true
		}

	}

	return false

}

func (distribution *Distribution) RemoveKernelModule(value string) {

	index := -1

	for k := 0; k < len(distribution.KernelModules); k++ {

		if distribution.KernelModules[k] == value {
			index = k
			break
		}

	}

	if index != -1 {
		distribution.KernelModules = append(distribution.KernelModules[:index], distribution.KernelModules[index+1:]...)
	}

}

func (distribution *Distribution) SetKernelModules(values []string) {

	filtered := make([]string, 0)

	for _, value := range values {
		filtered = append(filtered, value)
	}

	distribution.KernelModules = filtered

}

func (distribution *Distribution) SetKernelVersion(value string) {
	distribution.KernelVersion = strings.TrimSpace(value)
}

func (distribution *Distribution) SetKeywords(value map[string]string) {

	keywords := make(map[string]string)

	for key, val := range value {
		keywords[key] = val
	}

	distribution.Keywords = &keywords

}

func (distribution *Distribution) SetManager(value string) {

	manager := types.ParseManager(value)

	if manager != nil {
		distribution.Manager = manager.String()
	}

}

func (distribution *Distribution) SetName(value string) {
	distribution.Name = strings.TrimSpace(value)
}

func (distribution *Distribution) SetVendor(value string) {
	distribution.Vendor = strings.TrimSpace(value)
}

func (distribution *Distribution) SetVersion(value string) {
	distribution.Version = strings.TrimSpace(value)
}

func (distribution *Distribution) Hash() string {

	hash := ""

	if distribution.Name != "" {
		hash = toHash(distribution.Name, distribution.Version, distribution.Vendor, distribution.Kernel)
	}

	return hash

}
