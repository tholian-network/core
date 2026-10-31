package insights

import "tholian-core/structs"
import utils_os "tholian-core/utils/os"
import "encoding/json"

func writeMitigation(database *Database, mitigation *structs.Mitigation) bool {

	if mitigation != nil && mitigation.IsValid() {

		sortMitigation(mitigation)

		mitigation_hash := mitigation.Hash()

		if utils_os.IsSafeBasename(mitigation_hash) == true {

			mutex := database.BorrowMitigation(mitigation_hash)
			mutex.Lock()
			defer mutex.Unlock()

			buffer, err := json.MarshalIndent(mitigation, "", "\t")

			if err == nil {
				return writeBuffer(database, "/mitigations/embedded/"+mitigation_hash+".json", buffer)
			}

		}

	}

	return false

}
