package caches

import "tholian-core/structs"
import "tholian-core/types"
import "encoding/hex"
import "sort"
import "strings"

func queryNetworksByIP(cache *Networks, query string) []*structs.Subnet {

	result := make([]*structs.Subnet, 0)

	if query != "" && query != "any" {

		if strings.Contains(query, "/") {

			query_ip := query[0:strings.Index(query, "/")]
			query_prefix := query[strings.Index(query, "/")+1:]

			if types.IsIPv6(query_ip) {

				ipv6 := types.ParseIPv6(query_ip)
				prefix := parsePrefix(query_prefix)

				if ipv6 != nil && prefix != nil {

					prefixes := make([]uint8, 0)

					for tmp, _ := range cache.mapv6 {

						if tmp <= uint8(*prefix) {
							prefixes = append(prefixes, tmp)
						}

					}

					sort.Slice(prefixes, func(a int, b int) bool {
						return prefixes[a] > prefixes[b]
					})

					for _, prefix := range prefixes {

						hash := hex.EncodeToString(ipv6.Bytes(prefix))
						subnet, ok := cache.mapv6[prefix][hash]

						if ok == true {
							result = append(result, subnet)
						}

					}

				}

			} else if types.IsIPv4(query_ip) {

				ipv4 := types.ParseIPv4(query_ip)
				prefix := parsePrefix(query_prefix)

				if ipv4 != nil && prefix != nil {

					prefixes := make([]uint8, 0)

					for tmp, _ := range cache.mapv4 {

						if tmp <= uint8(*prefix) {
							prefixes = append(prefixes, tmp)
						}

					}

					sort.Slice(prefixes, func(a int, b int) bool {
						return prefixes[a] > prefixes[b]
					})

					for _, prefix := range prefixes {

						hash := hex.EncodeToString(ipv4.Bytes(prefix))
						subnet, ok := cache.mapv4[prefix][hash]

						if ok == true {
							result = append(result, subnet)
						}

					}

				}

			}

		} else {

			if types.IsIPv6(query) {

				ipv6 := types.ParseIPv6(query)
				prefix := uint8(128)

				if ipv6 != nil {

					prefixes := make([]uint8, 0)

					for tmp, _ := range cache.mapv6 {

						if tmp <= prefix {
							prefixes = append(prefixes, tmp)
						}

					}

					sort.Slice(prefixes, func(a int, b int) bool {
						return prefixes[a] > prefixes[b]
					})

					for _, prefix := range prefixes {

						hash := hex.EncodeToString(ipv6.Bytes(prefix))
						subnet, ok := cache.mapv6[prefix][hash]

						if ok == true {
							result = append(result, subnet)
						}

					}

				}

			} else if types.IsIPv4(query) {

				ipv4 := types.ParseIPv4(query)
				prefix := uint8(32)

				if ipv4 != nil {

					prefixes := make([]uint8, 0)

					for tmp, _ := range cache.mapv4 {

						if tmp <= prefix {
							prefixes = append(prefixes, tmp)
						}

					}

					sort.Slice(prefixes, func(a int, b int) bool {
						return prefixes[a] > prefixes[b]
					})

					for _, prefix := range prefixes {

						hash := hex.EncodeToString(ipv4.Bytes(prefix))
						subnet, ok := cache.mapv4[prefix][hash]

						if ok == true {
							result = append(result, subnet)
						}

					}

				}

			}

		}

	}

	if len(result) > 0 {

		sort.Slice(result, func(a int, b int) bool {
			return result[a].Prefix > result[b].Prefix
		})

	}

	return result

}
