package insights

import utils_os "tholian-core/utils/os"
import "tholian-core/structs"
import "encoding/json"

func readCredential(database *Database, name string) *structs.Credential {

	var result structs.Credential

	if utils_os.IsSafeBasename(name) == true {

		buffer := readBuffer(database, "/credentials/embedded/"+name+".json")

		if len(buffer) > 0 {

			err := json.Unmarshal(buffer, &result)

			if err == nil {
				// Do Nothing
			} else {
				result = structs.NewCredential(name, "user")
			}

		} else {
			result = structs.NewCredential(name, "user")
		}

	} else {
		result = structs.NewCredential("", "user")
	}

	return &result

}
