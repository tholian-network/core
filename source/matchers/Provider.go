package matchers

import "encoding/binary"
import "encoding/hex"
import "hash/crc32"
import "strings"

type Provider struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	Subnet  string `json:"subnet"`
}

func NewProvider() Provider {

	var provider Provider

	provider.Name = Any
	provider.Country = Any
	provider.Subnet = Any

	return provider

}

func ToProvider(value string) Provider {

	var provider Provider

	provider.Name = Any
	provider.Country = Any
	provider.Subnet = Any

	provider.SetName(value)

	return provider

}

func (provider *Provider) IsIdentical(value Provider) bool {

	if provider.Name == value.Name &&
		provider.Country == value.Country &&
		provider.Subnet == value.Subnet {
		return true
	}

	return false

}

func (provider *Provider) IsValid() bool {

	if provider.Name != Any ||
		provider.Country != Any ||
		provider.Subnet != Any {
		return true
	}

	return false

}

func (provider *Provider) Matches(name string, country string, subnet string) bool {
	return provider.MatchesName(name) && provider.MatchesCountry(country) && provider.MatchesSubnet(subnet)
}

func (provider *Provider) MatchesCountry(value string) bool {

	if provider.Country == value {
		return true
	} else if provider.Country == Any {
		return true
	}

	return false

}

func (provider *Provider) MatchesName(value string) bool {

	if provider.Name == value {
		return true
	} else if provider.Name == Any {
		return true
	}

	return false

}

func (provider *Provider) MatchesSubnet(value string) bool {

	if provider.Subnet != Any && value != "any" {
		return containsSubnet(value, provider.Subnet)
	} else if provider.Subnet == Any {
		return true
	}

	return false

}

func (provider *Provider) SetName(value string) {
	provider.Name = strings.TrimSpace(value)
}

func (provider *Provider) SetCountry(value string) {

	if value == "all" || value == "any" || value == "*" {
		provider.Country = Any
	} else if value != "" {
		provider.Country = value
	}

}

func (provider *Provider) SetSubnet(value string) {

	if value == "all" || value == "any" || value == "*" {

		provider.Subnet = Any

	} else if strings.Contains(value, "/") {

		address, prefix := toSubnet(value)

		if address != "" && prefix != 0 {
			provider.Subnet = value
		}

	}

}

func (provider *Provider) Hash() string {

	hash := ""

	if provider.Name != "" {

		checksum := crc32.ChecksumIEEE([]byte(strings.Join([]string{
			provider.Name,
			provider.Country,
			provider.Subnet,
		}, "-")))

		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, checksum)
		hash = hex.EncodeToString(tmp)

	}

	return hash

}
