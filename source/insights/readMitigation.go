package insights

import utils_os "tholian-core/utils/os"
import "tholian-core/structs"
import "encoding/json"

func readMitigation(database *Database, hash string) *structs.Mitigation {

	var result structs.Mitigation

	if utils_os.IsSafeBasename(hash) == true {

		buffer := readBuffer(database, "/mitigations/embedded/"+hash+".json")

		if len(buffer) > 0 {

			err := json.Unmarshal(buffer, &result)

			if err == nil {
				// Do Nothing
			} else {
				result = structs.Mitigation{}
			}

		} else {
			result = structs.Mitigation{}
		}

	}

	return &result

}
