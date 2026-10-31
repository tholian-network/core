package insights

import utils_os "tholian-core/utils/os"
import "tholian-core/structs"
import "encoding/json"

func readProduct(database *Database, name string) *structs.Product {

	var result structs.Product

	if utils_os.IsSafeBasename(name) == true {

		buffer := readBuffer(database, "/products/embedded/"+name+".json")

		if len(buffer) > 0 {

			err := json.Unmarshal(buffer, &result)

			if err == nil {
				// Do Nothing
			} else {
				result = structs.NewProduct("any")
			}

		} else {
			result = structs.NewProduct("any")
		}

	} else {
		result = structs.NewProduct("any")
	}

	return &result

}
