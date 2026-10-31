package insights

import "tholian-core/structs"
import utils_path "tholian-core/utils/path"
import utils_sort "tholian-core/utils/sort"
import "encoding/json"

func writeIncident(database *Database, incident *structs.Incident) bool {

	if incident != nil && incident.IsValid() {

		utils_sort.ByHash(incident.Connections)
		utils_sort.ByHash(incident.Distributions)
		utils_sort.ByHash(incident.Evidence)
		utils_sort.ByHash(incident.Networks)
		utils_sort.ByHash(incident.Packages)
		utils_sort.ByHash(incident.Products)
		utils_sort.ByHash(incident.Programs)
		utils_sort.ByHash(incident.Users)
		utils_sort.ByHash(incident.Vulnerabilities)
		utils_sort.ByHash(incident.Weaknesses)

		incident_hash := incident.Hash()

		if utils_path.IsSafeBasename(incident_hash) == true {

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
