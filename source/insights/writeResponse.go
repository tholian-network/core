package insights

import "tholian-core/structs"
import utils_os "tholian-core/utils/os"
import "encoding/json"

func writeResponse(database *Database, response *structs.Response) bool {

	if response != nil && response.IsValid() {

		sortResponse(response)

		response_hash := response.Hash()

		if utils_os.IsSafeBasename(response_hash) == true {

			mutex := database.BorrowResponse(response_hash)
			mutex.Lock()
			defer mutex.Unlock()

			buffer, err := json.MarshalIndent(response, "", "\t")

			if err == nil {
				return writeBuffer(database, "/responses/embedded/"+response_hash+".json", buffer)
			}

		}

	}

	return false

}
