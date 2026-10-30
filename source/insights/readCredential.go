package insights

import "tholian-core/structs"
import "encoding/json"

func readCredential(database *Database, name string) *structs.Credential {

	var result structs.Credential

	buffer := database.Read("/credentials/embedded/" + name + ".json")

	if len(buffer) > 0 {

		err := json.Unmarshal(buffer, &result)

		if err == nil {
			// Do Nothing
		} else {
			result = structs.NewCredential(name, "user")
		}

	} else {
		result = structs.NewCredential(name, "user")
	}

	return &result

}

