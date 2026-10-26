package structs

import "tholian-core/types"
import "encoding/hex"
import "strconv"
import "strings"

type Subnet struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	Address string `json:"address"`
	Scope   string `json:"scope"`
	Type    string `json:"type"`
	Prefix  uint8  `json:"prefix"`
}

func NewSubnet(address string, prefix uint8) Subnet {

	var subnet Subnet

	subnet.SetAddress(address)
	subnet.SetPrefix(prefix)

	return subnet

}

func ToSubnet(value string) Subnet {

	var subnet Subnet

	subnet.Address = ""
	subnet.Prefix = 0

	subnet.Parse(value)

	return subnet

}

func (subnet *Subnet) Contains(value string) bool {

	if types.IsIPv6(value) {

		ipv6 := types.ParseIPv6(value)

		if ipv6 != nil {

			tmp := NewSubnet(ipv6.String(), 128)

			return subnet.ContainsSubnet(tmp)

		}

	} else if types.IsIPv4(value) {

		ipv4 := types.ParseIPv4(value)

		if ipv4 != nil {

			tmp := NewSubnet(ipv4.String(), 32)

			return subnet.ContainsSubnet(tmp)

		}

	}

	return false

}

func (subnet *Subnet) ContainsSubnet(value Subnet) bool {

	if subnet.Type == types.AddressTypeIPv4 && value.Type == types.AddressTypeIPv4 {

		if value.Prefix >= subnet.Prefix {

			ipv4_subnet := types.ParseIPv4(subnet.Address)
			ipv4_value := types.ParseIPv4(value.Address)

			if ipv4_subnet != nil && ipv4_value != nil {

				bytes_subnet := ipv4_subnet.Bytes(subnet.Prefix)
				bytes_value := ipv4_value.Bytes(subnet.Prefix)

				matches_bytes := true

				for b := 0; b < len(bytes_subnet); b++ {

					if bytes_subnet[b] != bytes_value[b] {
						matches_bytes = false
						break
					}

				}

				return matches_bytes

			}

		}

	} else if subnet.Type == types.AddressTypeIPv6 && value.Type == types.AddressTypeIPv6 {

		if value.Prefix >= subnet.Prefix {

			ipv6_subnet := types.ParseIPv6(subnet.Address)
			ipv6_value := types.ParseIPv6(value.Address)

			if ipv6_subnet != nil && ipv6_value != nil {

				bytes_subnet := ipv6_subnet.Bytes(subnet.Prefix)
				bytes_value := ipv6_value.Bytes(subnet.Prefix)

				matches_bytes := true

				for b := 0; b < len(bytes_subnet); b++ {

					if bytes_subnet[b] != bytes_value[b] {
						matches_bytes = false
						break
					}

				}

				return matches_bytes

			}

		}

	}

	return false

}

func (subnet *Subnet) Equal(value Subnet) bool {

	if subnet.Address == value.Address && subnet.Prefix == value.Prefix {
		return true
	}

	return false

}

func (subnet *Subnet) IsUnspecified() bool {

	if subnet.Address == types.UnspecifiedIPv4 {
		return true
	}

	if subnet.Type == types.AddressTypeIPv6 {

		ipv6 := types.ParseIPv6(subnet.Address)

		return ipv6 != nil && ipv6.IsUnspecified()

	}

	return false

}

func (subnet *Subnet) IsDefaultRoute() bool {
	return subnet.IsUnspecified() && subnet.Prefix == 0
}

func (subnet *Subnet) IsAddress() bool {

	switch subnet.Type {
	case types.AddressTypeIPv4:
		return subnet.Prefix == 32
	case types.AddressTypeIPv6:
		return subnet.Prefix == 128
	}

	return false

}

func (subnet *Subnet) IsValid() bool {

	if subnet.Address == "" {
		return false
	}

	switch subnet.Type {
	case types.AddressTypeIPv4:

		if subnet.Prefix > 32 {
			return false
		}

		if subnet.Address == types.UnspecifiedIPv4 {
			return subnet.Prefix == 0
		}

		return types.IsIPv4(subnet.Address)

	case types.AddressTypeIPv6:

		if subnet.Prefix > 128 {
			return false
		}

		if types.IsIPv6(subnet.Address) == false {
			return false
		}

		if subnet.IsUnspecified() {
			return subnet.Prefix == 0
		}

		return true

	}

	return false

}

func (subnet *Subnet) Parse(value string) {

	var address string
	var prefix uint8
	var has_prefix bool

	if strings.HasPrefix(value, "[") && strings.Contains(value, "]/") {

		tmp1 := value[0 : strings.Index(value, "]/")+1]
		tmp2, err := strconv.ParseUint(strings.Split(value, "]/")[1], 10, 8)

		if err == nil {
			address = strings.ToLower(tmp1)
			prefix = uint8(tmp2)
			has_prefix = true
		}

	} else if strings.HasPrefix(value, "[") && strings.HasSuffix(value, "]") {

		address = strings.ToLower(value)
		prefix = uint8(128)
		has_prefix = true

	} else if strings.Contains(value, ":") && strings.Contains(value, "/") {

		tmp1 := strings.Split(value, "/")[0]
		tmp2, err := strconv.ParseUint(strings.Split(value, "/")[1], 10, 8)

		if err == nil {
			address = strings.ToLower(tmp1)
			prefix = uint8(tmp2)
			has_prefix = true
		}

	} else if strings.Contains(value, ":") {

		address = "[" + strings.ToLower(value) + "]"
		prefix = uint8(128)
		has_prefix = true

	} else if strings.Contains(value, ".") && strings.Contains(value, "/") {

		tmp1 := strings.Split(value, "/")[0]
		tmp2, err := strconv.ParseUint(strings.Split(value, "/")[1], 10, 8)

		if err == nil {
			address = tmp1
			prefix = uint8(tmp2)
			has_prefix = true
		}

	} else if strings.Contains(value, ".") {

		address = value
		prefix = uint8(32)
		has_prefix = true

	}

	if address != "" {

		subnet.SetAddress(address)

		if has_prefix == true {

			if subnet.Type == types.AddressTypeIPv4 && prefix > 32 {
				*subnet = Subnet{}
				return
			} else if subnet.Type == types.AddressTypeIPv6 && prefix > 128 {
				*subnet = Subnet{}
				return
			}

			subnet.SetPrefix(prefix)

		}

	}

}

func (subnet *Subnet) SetCountry(value string) {

	if len(value) == 2 && value != "XX" {
		subnet.Country = value
	}

}

func (subnet *Subnet) SetName(value string) {

	value = strings.TrimSpace(value)

	if value != "" {
		subnet.Name = value
	}

}

func (subnet *Subnet) SetAddress(value string) {

	if types.IsIPv6(value) {

		ipv6 := types.ParseIPv6(value)

		if ipv6 != nil {
			subnet.Address = ipv6.String()
			subnet.Scope = ipv6.Scope()
			subnet.Type = types.AddressTypeIPv6
		}

	} else if types.IsIPv4(value) {

		ipv4 := types.ParseIPv4(value)

		if ipv4 != nil {
			subnet.Address = ipv4.String()
			subnet.Scope = ipv4.Scope()
			subnet.Type = types.AddressTypeIPv4
		}

	}

}

func (subnet *Subnet) SetPrefix(value uint8) {

	if subnet.Type == types.AddressTypeIPv4 {

		if value <= 32 {
			subnet.Prefix = value
		}

	} else if subnet.Type == types.AddressTypeIPv6 {

		if value <= 128 {
			subnet.Prefix = value
		}

	}

}

func (subnet *Subnet) SetScope(value string) {

	if value == "private" || value == "public" {
		subnet.Scope = value
	}

}

func (subnet *Subnet) Addresses() uint {

	result := uint(0)

	if subnet.IsValid() && subnet.Prefix != 0 {

		if subnet.Type == types.AddressTypeIPv4 {
			result = 1 << (32 - subnet.Prefix)
		} else if subnet.Type == types.AddressTypeIPv6 {
			result = 1 << (128 - subnet.Prefix)
		}

	}

	return result

}

func (subnet *Subnet) Hash() string {

	hash := ""

	if subnet.Type == types.AddressTypeIPv4 {

		ipv4 := types.ParseIPv4(subnet.Address)

		if ipv4 != nil {
			bytes := ipv4.Bytes(subnet.Prefix)
			hash = hex.EncodeToString(bytes)
		}

	} else if subnet.Type == types.AddressTypeIPv6 {

		ipv6 := types.ParseIPv6(subnet.Address)

		if ipv6 != nil {
			bytes := ipv6.Bytes(subnet.Prefix)
			hash = hex.EncodeToString(bytes)
		}

	}

	return hash

}

func (subnet *Subnet) String() string {

	var result string

	if subnet.IsValid() {
		result = subnet.Address + "/" + strconv.FormatUint(uint64(subnet.Prefix), 10)
	}

	return result

}
