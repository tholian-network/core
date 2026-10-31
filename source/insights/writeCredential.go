package insights

import "tholian-core/structs"
import utils_os "tholian-core/utils/os"
import "bytes"
import "encoding/json"
import "sort"
import "strings"

func writeCredential(database *Database, credential *structs.Credential) bool {

	if credential != nil && credential.IsValid() {

		sort.Slice(credential.Keys, func(a int, b int) bool {
			return bytes.Compare(credential.Keys[a], credential.Keys[b]) < 0
		})
		sort.Slice(credential.Passwords, func(a int, b int) bool {
			return strings.Compare(credential.Passwords[a], credential.Passwords[b]) < 0
		})

		credential_name := credential.Name

		if utils_os.IsSafeBasename(credential_name) == true {

			mutex := database.BorrowCredential(credential_name)
			mutex.Lock()
			defer mutex.Unlock()

			buffer, err := json.MarshalIndent(credential, "", "\t")

			if err == nil {
				return writeBuffer(database, "/credentials/embedded/"+credential_name+".json", buffer)
			}

		}

	}

	return false

}
