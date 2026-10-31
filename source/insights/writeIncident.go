package insights

import "tholian-core/structs"
import utils_os "tholian-core/utils/os"
import "encoding/json"

func writeIncident(database *Database, incident *structs.Incident) bool {

	if incident != nil && incident.IsValid() {

		sortIncident(incident)

		incident_hash := incident.Hash()

		if utils_os.IsSafeBasename(incident_hash) == true {

			mutex := database.BorrowIncident(incident_hash)
			mutex.Lock()
			defer mutex.Unlock()

			buffer, err := json.MarshalIndent(incident, "", "\t")

			if err == nil {
				return writeBuffer(database, "/incidents/embedded/"+incident_hash+".json", buffer)
			}

		}

	}

	return false

}
