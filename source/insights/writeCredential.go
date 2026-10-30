package insights

import "tholian-core/structs"
import compress_gzip "tholian-core/utils/compress/gzip"
import "bytes"
import "encoding/json"
import "sort"
import "strings"

func writeCredential(database *Database, credential *structs.Credential) bool {

	if credential != nil {

		// Sorting has to be done here
		sort.Slice(credential.Keys, func(a int, b int) bool {
			return bytes.Compare(credential.Keys[a], credential.Keys[b]) < 0
		})
		sort.Slice(credential.Passwords, func(a int, b int) bool {
			return strings.Compare(credential.Passwords[a], credential.Passwords[b]) < 0
		})

		buffer, err0 := json.MarshalIndent(credential, "", "\t")
		buffer_gzip, err0_gzip := compress_gzip.Compress(buffer)

		if err0 == nil && err0_gzip == nil {

			result1 := database.Write("/credentials/embedded/"+credential.Name+".json.gz", buffer_gzip)
			result2 := database.Write("/credentials/embedded/"+credential.Name+".json", buffer)

			if result1 && result2 {
				return true
			}

		}

	}

	return false

}

