package insights

import utils_os "tholian-core/utils/os"
import "tholian-core/insights/templates"
import "tholian-core/structs"
import utils_tags "tholian-core/utils/tags"
import "strings"
import "sync"

func (database *Database) BorrowProduct(name string) *sync.Mutex {

	database.mutex.RLock()
	mutex, ok := database.mutexes.Products[name]
	database.mutex.RUnlock()

	if ok == false {

		database.mutex.Lock()
		database.mutexes.Products[name] = &sync.Mutex{}
		mutex = database.mutexes.Products[name]
		database.mutex.Unlock()

	}

	return mutex

}

func (database *Database) CleanupProducts() {

	files := database.Scan("/products/embedded", "*.json")

	if len(files) > 0 {

		for _, file := range files {

			if strings.HasPrefix(file, "/products/embedded/") && strings.HasSuffix(file, ".json") {

				name := strings.TrimSuffix(strings.TrimPrefix(file, "/products/embedded/"), ".json")
				mutex := database.BorrowProduct(name)

				mutex.Lock()

				product := readProduct(database, name)

				if product.State != "edited" {
					database.Remove(file)
					database.Remove(file + ".gz")
				}

				mutex.Unlock()

			}

		}

	}

}

func (database *Database) ReadProduct(name string) *structs.Product {

	mutex := database.BorrowProduct(name)
	mutex.Lock()
	defer mutex.Unlock()

	return readProduct(database, name)

}

func (database *Database) RemoveProduct(name string) {

	if utils_os.IsSafeBasename(name) == true {

		mutex := database.BorrowProduct(name)
		mutex.Lock()
		defer mutex.Unlock()

		database.Remove("/products/embedded/" + name + ".json.gz")
		database.Remove("/products/embedded/" + name + ".json")

	}

}

func (database *Database) WriteProduct(product *structs.Product) bool {
	return writeProduct(database, product)
}

func (database *Database) WriteProducts() bool {

	result := true

	files := database.Scan("/products/embedded", "*.json.gz")
	files_all := make([]string, 0)

	if len(files) > 0 {

		for _, file := range files {

			if strings.HasPrefix(file, "/products/embedded/") && strings.HasSuffix(file, ".json.gz") {

				name := strings.TrimSuffix(strings.TrimPrefix(file, "/products/embedded/"), ".json.gz")
				product := database.ReadProduct(name)

				if product.IsValid() {
					files_all = append(files_all, "embedded/"+name+".json.gz")
				}

			}

		}

	}

	for _, file := range database.Scan("/products", "Products_embed.*") {
		database.Remove(file)
	}

	for _, tag := range []string{"scout", "guard", "oversight"} {

		insight_go, insight_embed_go, err := templates.FilesystemCache(&templates.Config{
			Name:     "Products",
			Package:  "products",
			BuildTag: utils_tags.ToBuildTags(tag, database.GetDistributionNames()),
			Struct:   "structs.Product",
			Folder:   "embedded",
			Files:    files_all,
		})

		if err == nil {

			result1 := database.Write("/products/Products.go", insight_go)
			result2 := database.Write("/products/Products_embed."+tag+".go", insight_embed_go)

			if result1 == false || result2 == false {
				result = false
			}

		} else {
			result = false
		}

	}

	return result

}
