package insights

import utils_os "tholian-core/utils/os"
import "tholian-core/structs"
import "encoding/json"

func readIncident(database *Database, hash string) *structs.Incident {

	var result structs.Incident

	if utils_os.IsSafeBasename(hash) == true {

		buffer := readBuffer(database, "/incidents/embedded/"+hash+".json")

		if len(buffer) > 0 {

			err := json.Unmarshal(buffer, &result)

			if err == nil {
				// Do Nothing
			} else {
				result = structs.NewIncident("", "")
			}

		} else {
			result = structs.NewIncident("", "")
		}

	} else {
		result = structs.NewIncident("", "")
	}

	return &result

}
