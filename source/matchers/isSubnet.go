package matchers

import "tholian-core/types"
import "strings"

func isSubnet(value string) bool {

	if strings.Contains(value, "/") == false || strings.Count(value, "/") != 1 {
		return false
	}

	address, prefix := toSubnet(value)

	if types.IsIPv4(address) {
		return prefix <= 32
	} else if types.IsIPv6(address) {
		return prefix <= 128
	}

	return false

}
