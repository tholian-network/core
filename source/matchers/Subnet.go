package matchers

import "tholian-core/types"
import "encoding/hex"
import "strings"

type Subnet struct {
	Name   string `json:"name"`
	Subnet string `json:"subnet"`
}

func NewSubnet() Subnet {

	var subnet Subnet

	subnet.Name = Any
	subnet.Subnet = Any

	return subnet

}

func ToSubnet(value string) Subnet {

	subnet := NewSubnet()
	value = strings.TrimSpace(value)

	if strings.Contains(value, "/") {
		subnet.SetSubnet(value)
	} else {
		subnet.SetName(value)
	}

	return subnet

}

func (subnet *Subnet) IsAny() bool {
	return subnet.Name == Any && subnet.Subnet == Any
}

func (subnet *Subnet) HasName() bool {
	return types.IsASN(subnet.Name)
}

func (subnet *Subnet) HasSubnet() bool {
	return isSubnet(subnet.Subnet)
}

func (subnet *Subnet) IsIdentical(value Subnet) bool {

	var result bool = false

	if subnet.Name == value.Name && subnet.Subnet == value.Subnet {
		result = true
	}

	return result

}

func (subnet *Subnet) IsValid() bool {
	return subnet.HasName() != subnet.HasSubnet()
}

func (subnet *Subnet) Matches(name string, net string) bool {

	if subnet.IsAny() {
		return true
	} else if subnet.IsValid() == false {
		return false
	}

	matches_name := true
	matches_subnet := true

	if subnet.HasName() {
		matches_name = subnet.Name == name
	}

	if subnet.HasSubnet() {
		matches_subnet = net != "" && net != Any && contains_subnet(subnet.Subnet, net)
	}

	return matches_name && matches_subnet

}

func (subnet *Subnet) SetName(value string) {

	value = strings.TrimSpace(value)

	if types.IsASN(value) {
		subnet.Name = value
		subnet.Subnet = Any
	}

}

func (subnet *Subnet) SetSubnet(value string) {

	value = strings.TrimSpace(value)

	if isSubnet(value) {
		subnet.Name = Any
		subnet.Subnet = value
	}

}

func (subnet *Subnet) Hash() string {

	var hash string

	if subnet.HasName() {

		hash = subnet.Name

	} else if subnet.HasSubnet() {

		address, prefix := toSubnet(subnet.Subnet)

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
