package insights

import "tholian-core/structs"
import utils_os "tholian-core/utils/os"
import "encoding/json"

func writeProduct(database *Database, product *structs.Product) bool {

	if product != nil && product.IsValid() {

		product_name := product.Name

		if utils_os.IsSafeBasename(product_name) == true {

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
