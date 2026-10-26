package structs

import "tholian-core/types"

func equalSystemUsers(values []types.User, others []types.User) bool {

	if len(values) != len(others) {
		return false
	}

	for u := 0; u < len(values); u++ {

		if values[u].ID != others[u].ID ||
			values[u].Name != others[u].Name ||
			values[u].Password != others[u].Password ||
			values[u].Folder != others[u].Folder ||
			values[u].Shell != others[u].Shell ||
			values[u].Type != others[u].Type {
			return false
		}

		if len(values[u].Groups) != len(others[u].Groups) {
			return false
		}

		for g := 0; g < len(values[u].Groups); g++ {

			if values[u].Groups[g] != others[u].Groups[g] {
				return false
			}

		}

	}

	return true

}

