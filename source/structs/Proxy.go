package structs

import "tholian-core/types"
import utils_net "tholian-core/utils/net"
import "math/rand"

type Proxy struct {
	Name      string         `json:"name"`
	Addresses []string       `json:"addresses"`
	Port      uint16         `json:"port"`
	Protocol  types.Protocol `json:"protocol"`
}

func NewProxy() Proxy {

	var proxy Proxy

	proxy.Name = ""
	proxy.Port = 0
	proxy.Protocol = types.Any
	proxy.Addresses = make([]string, 0)

	return proxy

}

func (proxy *Proxy) IsValid() bool {

	if proxy.Name != "" &&
		proxy.Port != 0 &&
		proxy.Protocol != types.Any &&
		len(proxy.Addresses) > 0 {
		return true
	}

	return false

}

func (proxy *Proxy) AddAddress(value string) {

	valid := false

	if types.IsIPv6(value) {

		ipv6 := types.ParseIPv6(value)

		if ipv6 != nil {
			value = ipv6.String()
			valid = true
		}

	} else if types.IsIPv4(value) {

		ipv4 := types.ParseIPv4(value)

		if ipv4 != nil {
			value = ipv4.String()
			valid = true
		}

	}

	if valid == true {

		found := false

		for a := 0; a < len(proxy.Addresses); a++ {

			if proxy.Addresses[a] == value {
				found = true
				break
			}

		}

		if found == false {
			proxy.Addresses = append(proxy.Addresses, value)
		}

	}

}

func (proxy *Proxy) RemoveAddress(value string) {

	index := -1

	for a := 0; a < len(proxy.Addresses); a++ {

		if proxy.Addresses[a] == value {
			index = a
			break
		}

	}

	if index != -1 {
		proxy.Addresses = append(proxy.Addresses[:index], proxy.Addresses[index+1:]...)
	}

}

func (proxy *Proxy) SetAddresses(values []string) {

	filtered := make([]string, 0)

	for _, value := range values {

		if types.IsIPv6(value) {

			ipv6 := types.ParseIPv6(value)

			if ipv6 != nil {
				filtered = append(filtered, ipv6.String())
			}

		} else if types.IsIPv4(value) {

			ipv4 := types.ParseIPv4(value)

			if ipv4 != nil {
				filtered = append(filtered, ipv4.String())
			}

		}

	}

	proxy.Addresses = filtered

}

func (proxy *Proxy) SetName(value string) {

	if types.IsDomain(value) {

		domain := types.ParseDomain(value)

		if domain != nil {
			proxy.Name = domain.String()
		}

	}

}

func (proxy *Proxy) SetPort(value uint16) {

	if value > 0 && value < 65535 {
		proxy.Port = value
	}

}

func (proxy *Proxy) SetProtocol(value types.Protocol) {
	proxy.Protocol = value
}

func (proxy *Proxy) RandomizeAddress() string {

	result := ""

	if utils_net.SupportsIPv4() {

		filtered := make([]string, 0)

		for a := 0; a < len(proxy.Addresses); a++ {

			address := proxy.Addresses[a]

			if types.IsIPv4(address) {
				filtered = append(filtered, address)
			}

		}

		if len(filtered) > 0 {
			result = filtered[rand.Intn(len(filtered))]
		}

	} else if utils_net.SupportsIPv6() {

		filtered := make([]string, 0)

		for a := 0; a < len(proxy.Addresses); a++ {

			address := proxy.Addresses[a]

			if types.IsIPv6(address) {
				filtered = append(filtered, address)
			}

		}

		if len(filtered) > 0 {
			result = filtered[rand.Intn(len(filtered))]
		}

	}

	return result

}
