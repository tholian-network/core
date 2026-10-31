package sort

import "bytes"
import "encoding/json"
import "sort"

func ByHash[Type any](values []Type) {

	sort.Slice(values, func(a int, b int) bool {

		hash_a := hashOf(&values[a])
		hash_b := hashOf(&values[b])

		if hash_a != hash_b {
			return hash_a < hash_b
		}

		buffer_a, err_a := json.Marshal(values[a])
		buffer_b, err_b := json.Marshal(values[b])

		if err_a == nil && err_b == nil {
			return bytes.Compare(buffer_a, buffer_b) < 0
		}

		return false

	})

}
