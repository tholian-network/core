package caches

import "tholian-core/structs"

func addNetworksSubnet(cache *Networks, value structs.Subnet) bool {

	result := false

	if value.Type == "ipv6" {

		prefix := value.Prefix
		hash := value.Hash()

		_, ok1 := cache.mapv6[prefix]

		if ok1 == false {
			cache.mapv6[prefix] = make(map[string]*structs.Subnet)
		}

		_, ok2 := cache.mapv6[prefix][hash]

		if ok2 == true {
			result = false
		} else {
			cache.mapv6[prefix][hash] = &value
			result = true
		}

	} else if value.Type == "ipv4" {

		prefix := value.Prefix
		hash := value.Hash()

		_, ok1 := cache.mapv4[prefix]

		if ok1 == false {
			cache.mapv4[prefix] = make(map[string]*structs.Subnet)
		}

		_, ok2 := cache.mapv4[prefix][hash]

		if ok2 == true {
			result = false
		} else {
			cache.mapv4[prefix][hash] = &value
			result = true
		}

	}

	return result

}
