package caches

import "tholian-core/structs"

func removeNetworksSubnet(cache *Networks, value structs.Subnet) bool {

	result := false

	if value.Type == "ipv6" {

		prefix := value.Prefix
		hash := value.Hash()

		_, ok1 := cache.mapv6[prefix]

		if ok1 == true {

			_, ok2 := cache.mapv6[prefix][hash]

			if ok2 == true {
				delete(cache.mapv6[prefix], hash)
				result = true
			}

		}

	} else if value.Type == "ipv4" {

		prefix := value.Prefix
		hash := value.Hash()

		_, ok1 := cache.mapv4[prefix]

		if ok1 == true {

			_, ok2 := cache.mapv4[prefix][hash]

			if ok2 == true {
				delete(cache.mapv4[prefix], hash)
				result = true
			}

		}

	}

	return result

}
