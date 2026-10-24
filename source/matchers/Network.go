package matchers

import "tholian-core/types"
import "encoding/hex"
import "strings"

type Network struct {
	Name   string `json:"name"`
	Subnet string `json:"subnet"`
}

func NewNetwork() Network {

	var network Network

	network.Name = Any
	network.Subnet = Any

	return network

}

func ToNetwork(value string) Network {

	network := NewNetwork()
	value = strings.TrimSpace(value)

	if strings.Contains(value, "/") {
		network.SetSubnet(value)
	} else {
		network.SetName(value)
	}

	return network

}

func (network *Network) IsAny() bool {
	return network.Name == Any && network.Subnet == Any
}

func (network *Network) HasName() bool {
	return types.IsASN(network.Name)
}

func (network *Network) HasSubnet() bool {
	return isSubnet(network.Subnet)
}

func (network *Network) IsIdentical(value Network) bool {

	var result bool = false

	if network.Name == value.Name && network.Subnet == value.Subnet {
		result = true
	}

	return result

}

func (network *Network) IsValid() bool {
	return network.HasName() != network.HasSubnet()
}

func (network *Network) Matches(name string, subnet string) bool {

	if network.IsAny() {
		return true
	} else if network.IsValid() == false {
		return false
	}

	matches_name := true
	matches_subnet := true

	if network.HasName() {
		matches_name = network.Name == name
	}

	if network.HasSubnet() {
		matches_subnet = subnet != "" && subnet != Any && contains_subnet(network.Subnet, subnet)
	}

	return matches_name && matches_subnet

}

func (network *Network) SetName(value string) {

	value = strings.TrimSpace(value)

	if types.IsASN(value) {
		network.Name = value
		network.Subnet = Any
	}

}

func (network *Network) SetSubnet(value string) {

	value = strings.TrimSpace(value)

	if isSubnet(value) {
		network.Name = Any
		network.Subnet = value
	}

}

func (network *Network) Hash() string {

	var hash string

	if network.HasName() {

		hash = network.Name

	} else if network.HasSubnet() {

		address, prefix := toSubnet(network.Subnet)

		if types.IsIPv6(address) {

			ipv6 := types.ParseIPv6(address)

			if ipv6 != nil {
				bytes := ipv6.Bytes(prefix)
				hash = hex.EncodeToString(bytes)
			}

		} else if types.IsIPv4(address) {

			ipv4 := types.ParseIPv4(address)

			if ipv4 != nil {
				bytes := ipv4.Bytes(prefix)
				hash = hex.EncodeToString(bytes)
			}

		}

	}

	return hash

}
