package structs

import "tholian-core/types"
import "net/url"
import "strings"

type Antique struct {
	Name         string             `json:"name"`
	Version      types.Version      `json:"version"`
	Architecture types.Architecture `json:"architecture"`
	Manager      types.Manager      `json:"manager"`
	Vendor       string             `json:"vendor"`
	Service      string             `json:"service"`
	URL          string             `json:"url"`
}

func NewAntique(manager string, service string) Antique {

	var antique Antique

	antique.SetArchitecture("any")
	antique.SetManager(manager)
	antique.SetService(service)

	return antique

}

func (antique *Antique) Equal(value Antique) bool {

	if antique.Name == value.Name &&
		antique.Version.Equal(value.Version) &&
		antique.Architecture.Equal(value.Architecture) &&
		antique.Manager.Equal(value.Manager) &&
		antique.Vendor == value.Vendor &&
		antique.Service == value.Service &&
		antique.URL == value.URL {
		return true
	}

	return false

}

func (antique *Antique) IsValid() bool {

	if antique.Name != "" &&
		antique.Service != "" &&
		antique.Version.IsValid() &&
		antique.Architecture.IsValid() &&
		antique.Manager.IsValid() {
		return true
	}

	return false

}

func (antique *Antique) SetArchitecture(value string) {

	architecture := types.ParseArchitecture(value)

	if architecture != nil {
		antique.Architecture = *architecture
	}

}

func (antique *Antique) SetManager(value string) {

	manager := types.ParseManager(value)

	if manager != nil {
		antique.Manager = *manager
	}

}

func (antique *Antique) SetName(value string) {
	antique.Name = strings.TrimSpace(value)
}

func (antique *Antique) SetService(value string) {
	antique.Service = strings.TrimSpace(value)
}

func (antique *Antique) SetURL(value string) {

	tmp, err := url.ParseRequestURI(value)

	if err == nil {

		if tmp.Scheme == "https" || tmp.Scheme == "http" {
			antique.URL = tmp.String()
		}

	}

}

func (antique *Antique) SetVendor(value string) {
	antique.Vendor = strings.TrimSpace(value)
}

func (antique *Antique) SetVersion(value string) {

	version := types.ToVersion(value)

	if version.IsValid() {
		antique.Version = version
	}

}

func (antique *Antique) Hash() string {

	hash := ""

	if antique.Name != "" {

		hash = toHash(
			antique.Name,
			antique.Version.String(),
			antique.Architecture.String(),
			antique.Manager.String(),
			antique.Vendor,
			antique.Service,
			antique.URL,
		)

	}

	return hash

}
