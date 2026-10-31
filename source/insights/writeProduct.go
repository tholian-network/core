package insights

import "tholian-core/structs"
import utils_path "tholian-core/utils/path"
import "encoding/json"

func writeProduct(database *Database, product *structs.Product) bool {

	if product != nil && product.IsValid() {

		product_name := product.Name

		if utils_path.IsSafeBasename(product_name) == true {

			mutex := database.BorrowProduct(product_name)
			mutex.Lock()
			defer mutex.Unlock()

			buffer, err := json.MarshalIndent(product, "", "\t")

			if err == nil {
				return writeBuffer(database, "/products/embedded/"+product_name+".json", buffer)
			}

		}

	}

	return false

}
