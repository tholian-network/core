package insights

import "tholian-core/structs"
import compress_gzip "tholian-core/utils/compress/gzip"
import "encoding/json"

func writeMitigation(database *Database, mitigation *structs.Mitigation) bool {

	if mitigation != nil {

		// TODO: Sorting has to be done here

		mitigation_name := mitigation.Hash()

		buffer, err0 := json.MarshalIndent(mitigation, "", "\t")
		buffer_gzip, err0_gzip := compress_gzip.Compress(buffer)

		if err0 == nil && err0_gzip == nil {

			result1 := database.Write("/mitigations/embedded/"+mitigation_name+".json.gz", buffer_gzip)
			result2 := database.Write("/mitigations/embedded/"+mitigation_name+".json", buffer)

			if result1 && result2 {
				return true
			}

		}

	}

	return false

}

