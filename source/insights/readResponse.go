package insights

import utils_path "tholian-core/utils/path"
import "tholian-core/structs"
import "encoding/json"

func readResponse(database *Database, hash string) *structs.Response {

	var result structs.Response

	if utils_path.IsSafeBasename(hash) == true {

		buffer := readBuffer(database, "/responses/embedded/"+hash+".json")

		if len(buffer) > 0 {

			err := json.Unmarshal(buffer, &result)

			if err == nil {
				// Do Nothing
			} else {
				result = structs.NewResponse("", "none")
			}

		} else {
			result = structs.NewResponse("", "none")
		}

	} else {
		result = structs.NewResponse("", "none")
	}

	return &result

}
