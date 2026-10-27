package structs

import "sort"
import "strings"

type Network struct {
	Name    string   `json:"name"`
	Subnets []Subnet `json:"subnets"`
}

func NewNetwork(name string) Network {

	var network Network

	network.SetName(name)
	network.Subnets = make([]Subnet, 0)

	return network

}

func (network *Network) HasName() bool {
	return network.Name != ""
}

func (network *Network) HasSubnets() bool {
	return len(network.Subnets) > 0
}

func (network *Network) Equal(value Network) bool {

	if network.Name != value.Name {
		return false
	}

	if len(network.Subnets) != len(value.Subnets) {
		return false
	}

	for s := 0; s < len(network.Subnets); s++ {

		found := false

		for o := 0; o < len(value.Subnets); o++ {

			if network.Subnets[s].Equal(value.Subnets[o]) {
				found = true
				break
			}

		}

		if found == false {
			return false
		}

	}

	return true

}

func (network *Network) IsValid() bool {

	if network.HasName() == false {
		return false
	}

	for s := 0; s < len(network.Subnets); s++ {

		if network.Subnets[s].IsValid() == false {
			return false
		}

	}

	return true

}

func (network *Network) SetName(value string) {
	network.Name = strings.TrimSpace(value)
}

func (network *Network) AddSubnet(value Subnet) {

	if value.IsValid() {

		var found bool = false

		for s := 0; s < len(network.Subnets); s++ {

			if network.Subnets[s].Equal(value) {
				found = true
				break
			}

		}

		if found == false {
			network.Subnets = append(network.Subnets, value)
		}

	}

}

func (network *Network) HasSubnet(value Subnet) bool {

	for s := 0; s < len(network.Subnets); s++ {

		if network.Subnets[s].Equal(value) {
			return true
		}

	}

	return false

}

func (network *Network) RemoveSubnet(value Subnet) {

	var index int = -1

	for s := 0; s < len(network.Subnets); s++ {

		if network.Subnets[s].Equal(value) {
			index = s
			break
		}

	}

	if index != -1 {
		network.Subnets = append(network.Subnets[:index], network.Subnets[index+1:]...)
	}

}

func (network *Network) SetSubnets(value []Subnet) {

	var filtered []Subnet

	for v := 0; v < len(value); v++ {

		if value[v].IsValid() {
			filtered = append(filtered, value[v])
		}

	}

	sort.Slice(filtered, func(a int, b int) bool {

		if filtered[a].Prefix == filtered[b].Prefix {

			hash_a := filtered[a].Hash()
			hash_b := filtered[b].Hash()

			return hash_a < hash_b

		} else {

			return filtered[a].Prefix < filtered[b].Prefix

		}

	})

	network.Subnets = filtered

}

func (network *Network) Hash() string {

	hash := ""

	if network.Name != "" || len(network.Subnets) > 0 {

		values := make([]string, 0)

		values = append(values, network.Name)

		hashes := make([]string, 0)

		for s := 0; s < len(network.Subnets); s++ {
			hashes = append(hashes, network.Subnets[s].Hash())
		}

		sort.Strings(hashes)

		values = append(values, hashes...)

		hash = toHash(values...)

	}

	return hash

}
