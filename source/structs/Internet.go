package structs

import "tholian-core/types"
import "encoding/hex"
import "sort"
import "strconv"
import "strings"

type Internet struct {
	Mapv4    map[uint8]map[string]*Subnet `json:"mapv4"`
	Mapv6    map[uint8]map[string]*Subnet `json:"mapv6"`
	Networks map[string]*Network          `json:"networks"`
}

func NewInternet() Internet {

	var internet Internet

	internet.Mapv4 = make(map[uint8]map[string]*Subnet)
	internet.Mapv6 = make(map[uint8]map[string]*Subnet)
	internet.Networks = make(map[string]*Network)

	return internet

}

func (internet *Internet) Equal(value Internet) bool {

	if len(internet.Mapv4) != len(value.Mapv4) ||
		len(internet.Mapv6) != len(value.Mapv6) ||
		len(internet.Networks) != len(value.Networks) {
		return false
	}

	for prefix, subnets := range internet.Mapv4 {

		other, ok := value.Mapv4[prefix]

		if ok == false || len(subnets) != len(other) {
			return false
		}

		for hash, subnet := range subnets {

			otherSubnet, ok := other[hash]

			if ok == false || subnet == nil || otherSubnet == nil {
				return false
			}

			if subnet.Equal(*otherSubnet) == false {
				return false
			}

		}

	}

	for prefix, subnets := range internet.Mapv6 {

		other, ok := value.Mapv6[prefix]

		if ok == false || len(subnets) != len(other) {
			return false
		}

		for hash, subnet := range subnets {

			otherSubnet, ok := other[hash]

			if ok == false || subnet == nil || otherSubnet == nil {
				return false
			}

			if subnet.Equal(*otherSubnet) == false {
				return false
			}

		}

	}

	for name, network := range internet.Networks {

		other, ok := value.Networks[name]

		if ok == false || network == nil || other == nil {
			return false
		}

		if network.Equal(*other) == false {
			return false
		}

	}

	return true

}

func (internet *Internet) IsValid() bool {

	if internet.Mapv4 == nil || internet.Mapv6 == nil || internet.Networks == nil {
		return false
	}

	for prefix, subnets := range internet.Mapv4 {

		if subnets == nil {
			return false
		}

		for hash, subnet := range subnets {

			if subnet == nil {
				return false
			}

			if subnet.Type != types.AddressTypeIPv4 {
				return false
			}

			if subnet.Prefix != prefix {
				return false
			}

			if subnet.IsValid() == false {
				return false
			}

			if subnet.Hash() != hash {
				return false
			}

		}

	}

	for prefix, subnets := range internet.Mapv6 {

		if subnets == nil {
			return false
		}

		for hash, subnet := range subnets {

			if subnet == nil {
				return false
			}

			if subnet.Type != types.AddressTypeIPv6 {
				return false
			}

			if subnet.Prefix != prefix {
				return false
			}

			if subnet.IsValid() == false {
				return false
			}

			if subnet.Hash() != hash {
				return false
			}

		}

	}

	for name, network := range internet.Networks {

		if network == nil {
			return false
		}

		if network.Name != name {
			return false
		}

		if network.IsValid() == false {
			return false
		}

		for s := 0; s < len(network.Subnets); s++ {

			subnet := network.Subnets[s]

			if subnet.Type == types.AddressTypeIPv4 {

				subnets, ok := internet.Mapv4[subnet.Prefix]

				if ok == false {
					return false
				}

				if _, ok := subnets[subnet.Hash()]; ok == false {
					return false
				}

			} else if subnet.Type == types.AddressTypeIPv6 {

				subnets, ok := internet.Mapv6[subnet.Prefix]

				if ok == false {
					return false
				}

				if _, ok := subnets[subnet.Hash()]; ok == false {
					return false
				}

			} else {
				return false
			}

		}

	}

	return true

}

func (internet *Internet) AddNetwork(value Network) bool {

	if strings.HasPrefix(value.Name, "AS") {

		var network = value

		internet.Networks[network.Name] = &network

		for s := 0; s < len(network.Subnets); s++ {

			check := internet.AddSubnet(network.Subnets[s])

			if check == false {
				return false
			}

		}

		return true

	}

	return false

}

func (internet *Internet) AddSubnet(value Subnet) bool {

	var subnet = value

	if subnet.Type == types.AddressTypeIPv6 {

		prefix := subnet.Prefix
		hash := subnet.Hash()

		_, ok1 := internet.Mapv6[prefix]

		if ok1 == false {
			internet.Mapv6[prefix] = make(map[string]*Subnet)
		}

		_, ok2 := internet.Mapv6[prefix][hash]

		if ok2 == true {

			return false

		} else {

			internet.Mapv6[prefix][hash] = &subnet

			return true

		}

	} else if subnet.Type == types.AddressTypeIPv4 {

		prefix := subnet.Prefix
		hash := subnet.Hash()

		_, ok1 := internet.Mapv4[prefix]

		if ok1 == false {
			internet.Mapv4[prefix] = make(map[string]*Subnet)
		}

		_, ok2 := internet.Mapv4[prefix][hash]

		if ok2 == true {

			return false

		} else {

			internet.Mapv4[prefix][hash] = &subnet

			return true

		}

	}

	return false

}

func (internet *Internet) Search(value string) Subnet {

	var result Subnet

	if strings.Contains(value, "/") {

		address := value[0:strings.Index(value, "/")]
		prefix := value[strings.Index(value, "/")+1:]

		if types.IsIPv6(address) {

			ipv6 := types.ParseIPv6(address)

			if ipv6 != nil {
				address = ipv6.String()
				result = internet.SearchSubnet(ToSubnet(address + "/" + prefix))
			}

		} else if types.IsIPv4(address) {

			ipv4 := types.ParseIPv4(address)

			if ipv4 != nil {
				address = ipv4.String()
				result = internet.SearchSubnet(ToSubnet(address + "/" + prefix))
			}

		}

	} else {

		if types.IsIPv6(value) {

			ipv6 := types.ParseIPv6(value)

			if ipv6 != nil {
				value = ipv6.String()
				result = internet.SearchSubnet(ToSubnet(value + "/128"))
			}

		} else if types.IsIPv4(value) {

			ipv4 := types.ParseIPv4(value)

			if ipv4 != nil {
				value = ipv4.String()
				result = internet.SearchSubnet(ToSubnet(value + "/32"))
			}

		}

	}

	return result

}

func (internet *Internet) SearchASN(value string) Network {

	var result Network

	if strings.HasPrefix(value, "AS") {

		tmp, ok := internet.Networks[value]

		if ok == true {
			result = *tmp
		}

	}

	return result

}

func (internet *Internet) SearchNetwork(value Network) []Subnet {

	var result []Subnet

	if value.HasName() {

		network := internet.SearchASN(value.Name)

		if network.HasName() {

			for s := 0; s < len(network.Subnets); s++ {
				result = append(result, network.Subnets[s])
			}

		}

	} else if value.HasSubnets() {

		for s := 0; s < len(value.Subnets); s++ {

			found := internet.SearchSubnet(value.Subnets[s])

			if found.IsValid() {
				result = append(result, found)
			}

		}

	}

	return result

}

func (internet *Internet) SearchSubnet(value Subnet) Subnet {

	var result Subnet

	if value.Type == types.AddressTypeIPv6 {

		ipv6 := types.ParseIPv6(value.Address)

		if ipv6 == nil {
			return result
		}

		prefixes := make([]uint8, 0)

		for prefix := range internet.Mapv6 {
			prefixes = append(prefixes, prefix)
		}

		sort.Slice(prefixes, func(a int, b int) bool {
			return prefixes[a] > prefixes[b]
		})

		for p := 0; p < len(prefixes); p++ {

			prefix := prefixes[p]
			hash := hex.EncodeToString(ipv6.Bytes(prefix))

			tmp, ok := internet.Mapv6[prefix][hash]

			if ok == true {
				result = *tmp
				break
			}

		}

	} else if value.Type == types.AddressTypeIPv4 {

		ipv4 := types.ParseIPv4(value.Address)

		if ipv4 == nil {
			return result
		}

		prefixes := make([]uint8, 0)

		for prefix := range internet.Mapv4 {
			prefixes = append(prefixes, prefix)
		}

		sort.Slice(prefixes, func(a int, b int) bool {
			return prefixes[a] > prefixes[b]
		})

		for p := 0; p < len(prefixes); p++ {

			prefix := prefixes[p]
			hash := hex.EncodeToString(ipv4.Bytes(prefix))

			tmp, ok := internet.Mapv4[prefix][hash]

			if ok == true {
				result = *tmp
				break
			}

		}

	}

	return result

}

func (internet *Internet) SetNetworks(values []Network) bool {

	result := true

	for v := 0; v < len(values); v++ {

		check := internet.AddNetwork(values[v])

		if check == false {
			result = false
		}

	}

	return result

}

func (internet *Internet) SetSubnets(values []Subnet) bool {

	result := true

	for v := 0; v < len(values); v++ {

		check := internet.AddSubnet(values[v])

		if check == false {
			result = false
		}

	}

	return result

}

func (internet *Internet) Hash() string {

	values := make([]string, 0)

	prefixes4 := make([]uint8, 0)

	for prefix := range internet.Mapv4 {
		prefixes4 = append(prefixes4, prefix)
	}

	sort.Slice(prefixes4, func(a int, b int) bool {
		return prefixes4[a] < prefixes4[b]
	})

	for p := 0; p < len(prefixes4); p++ {

		prefix := prefixes4[p]
		hashes := make([]string, 0)

		for hash := range internet.Mapv4[prefix] {
			hashes = append(hashes, hash)
		}

		sort.Strings(hashes)

		for h := 0; h < len(hashes); h++ {
			subnet := internet.Mapv4[prefix][hashes[h]]
			values = append(values, "v4")
			values = append(values, strconv.FormatUint(uint64(prefix), 10))
			values = append(values, subnet.Hash())
		}

	}

	prefixes6 := make([]uint8, 0)

	for prefix := range internet.Mapv6 {
		prefixes6 = append(prefixes6, prefix)
	}

	sort.Slice(prefixes6, func(a int, b int) bool {
		return prefixes6[a] < prefixes6[b]
	})

	for p := 0; p < len(prefixes6); p++ {

		prefix := prefixes6[p]
		hashes := make([]string, 0)

		for hash := range internet.Mapv6[prefix] {
			hashes = append(hashes, hash)
		}

		sort.Strings(hashes)

		for h := 0; h < len(hashes); h++ {
			subnet := internet.Mapv6[prefix][hashes[h]]
			values = append(values, "v6")
			values = append(values, strconv.FormatUint(uint64(prefix), 10))
			values = append(values, subnet.Hash())
		}

	}

	names := make([]string, 0)

	for name := range internet.Networks {
		names = append(names, name)
	}

	sort.Strings(names)

	for n := 0; n < len(names); n++ {
		network := internet.Networks[names[n]]
		values = append(values, "net")
		values = append(values, names[n])
		values = append(values, network.Hash())
	}

	if len(values) == 0 {
		return ""
	}

	return toHash(values...)

}
