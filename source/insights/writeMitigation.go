package insights

import "tholian-core/structs"
import utils_path "tholian-core/utils/path"
import utils_sort "tholian-core/utils/sort"
import "encoding/json"

func writeMitigation(database *Database, mitigation *structs.Mitigation) bool {

	if mitigation != nil && mitigation.IsValid() {

		utils_sort.ByHash(mitigation.Incident.Connections)
		utils_sort.ByHash(mitigation.Incident.Distributions)
		utils_sort.ByHash(mitigation.Incident.Evidence)
		utils_sort.ByHash(mitigation.Incident.Networks)
		utils_sort.ByHash(mitigation.Incident.Packages)
		utils_sort.ByHash(mitigation.Incident.Products)
		utils_sort.ByHash(mitigation.Incident.Programs)
		utils_sort.ByHash(mitigation.Incident.Users)
		utils_sort.ByHash(mitigation.Incident.Vulnerabilities)
		utils_sort.ByHash(mitigation.Incident.Weaknesses)

		utils_sort.ByHash(mitigation.Response.Connections)
		utils_sort.ByHash(mitigation.Response.Distributions)
		utils_sort.ByHash(mitigation.Response.Networks)
		utils_sort.ByHash(mitigation.Response.Packages)
		utils_sort.ByHash(mitigation.Response.Products)
		utils_sort.ByHash(mitigation.Response.Programs)
		utils_sort.ByHash(mitigation.Response.Users)
		utils_sort.ByHash(mitigation.Response.Vulnerabilities)
		utils_sort.ByHash(mitigation.Response.Weaknesses)

		mitigation_hash := mitigation.Hash()

		if utils_path.IsSafeBasename(mitigation_hash) == true {

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
