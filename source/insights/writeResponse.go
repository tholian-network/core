package insights

import "tholian-core/structs"
import utils_path "tholian-core/utils/path"
import utils_sort "tholian-core/utils/sort"
import "encoding/json"

func writeResponse(database *Database, response *structs.Response) bool {

	if response != nil && response.IsValid() {

		utils_sort.ByHash(response.Connections)
		utils_sort.ByHash(response.Distributions)
		utils_sort.ByHash(response.Networks)
		utils_sort.ByHash(response.Packages)
		utils_sort.ByHash(response.Products)
		utils_sort.ByHash(response.Programs)
		utils_sort.ByHash(response.Users)
		utils_sort.ByHash(response.Vulnerabilities)
		utils_sort.ByHash(response.Weaknesses)

		response_hash := response.Hash()

		if utils_path.IsSafeBasename(response_hash) == true {

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
