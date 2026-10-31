package insights

import utils_path "tholian-core/utils/path"
import "tholian-core/insights/templates"
import "tholian-core/structs"
import utils_tags "tholian-core/utils/tags"
import "strings"
import "sync"

func (database *Database) BorrowResponse(hash string) *sync.Mutex {

	database.mutex.RLock()
	mutex, ok := database.mutexes.Responses[hash]
	database.mutex.RUnlock()

	if ok == false {

		database.mutex.Lock()
		database.mutexes.Responses[hash] = &sync.Mutex{}
		mutex = database.mutexes.Responses[hash]
		database.mutex.Unlock()

	}

	return mutex

}

func (database *Database) CleanupResponses() {

	files := database.Scan("/responses/embedded", "*.json")

	if len(files) > 0 {

		for _, file := range files {

			if strings.HasPrefix(file, "/responses/embedded/") && strings.HasSuffix(file, ".json") {

				hash := strings.TrimSuffix(strings.TrimPrefix(file, "/responses/embedded/"), ".json")
				mutex := database.BorrowResponse(hash)

				mutex.Lock()
				database.Remove(file)
				database.Remove(file + ".gz")
				mutex.Unlock()

			}

		}

	}

}

func (database *Database) ReadResponse(hash string) *structs.Response {

	mutex := database.BorrowResponse(hash)
	mutex.Lock()
	defer mutex.Unlock()

	return readResponse(database, hash)

}

func (database *Database) RemoveResponse(hash string) {

	if utils_path.IsSafeBasename(hash) == true {

		mutex := database.BorrowResponse(hash)
		mutex.Lock()
		defer mutex.Unlock()

		database.Remove("/responses/embedded/" + hash + ".json.gz")
		database.Remove("/responses/embedded/" + hash + ".json")

	}

}

func (database *Database) WriteResponse(response *structs.Response) bool {
	return writeResponse(database, response)
}

func (database *Database) WriteResponses() bool {

	result := true

	files := database.Scan("/responses/embedded", "*.json.gz")
	files_all := make([]string, 0)

	if len(files) > 0 {

		for _, file := range files {

			if strings.HasPrefix(file, "/responses/embedded/") && strings.HasSuffix(file, ".json.gz") {

				hash := strings.TrimSuffix(strings.TrimPrefix(file, "/responses/embedded/"), ".json.gz")
				response := database.ReadResponse(hash)

				if response.IsValid() {
					files_all = append(files_all, "embedded/"+hash+".json.gz")
				}

			}

		}

	}

	for _, file := range database.Scan("/responses", "Responses_embed.*") {
		database.Remove(file)
	}

	for _, tag := range []string{"guard", "oversight"} {

		insight_go, insight_embed_go, err := templates.FilesystemCache(&templates.Config{
			Name:     "Responses",
			Package:  "responses",
			BuildTag: utils_tags.ToBuildTags(tag, database.GetDistributionNames()),
			Struct:   "structs.Response",
			Folder:   "embedded",
			Files:    files_all,
		})

		if err == nil {

			result1 := database.Write("/responses/Responses.go", insight_go)
			result2 := database.Write("/responses/Responses_embed."+tag+".go", insight_embed_go)

			if result1 == false || result2 == false {
				result = false
			}

		} else {
			result = false
		}

	}

	return result

}
