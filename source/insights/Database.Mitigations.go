package insights

import utils_path "tholian-core/utils/path"
import "tholian-core/insights/templates"
import "tholian-core/structs"
import utils_tags "tholian-core/utils/tags"
import "strings"
import "sync"

func (database *Database) BorrowMitigation(hash string) *sync.Mutex {

	database.mutex.RLock()
	mutex, ok := database.mutexes.Mitigations[hash]
	database.mutex.RUnlock()

	if ok == false {

		database.mutex.Lock()
		database.mutexes.Mitigations[hash] = &sync.Mutex{}
		mutex = database.mutexes.Mitigations[hash]
		database.mutex.Unlock()

	}

	return mutex

}

func (database *Database) CleanupMitigations() {

	files := database.Scan("/mitigations/embedded", "*.json")

	if len(files) > 0 {

		for _, file := range files {

			if strings.HasPrefix(file, "/mitigations/embedded/") && strings.HasSuffix(file, ".json") {

				hash := strings.TrimSuffix(strings.TrimPrefix(file, "/mitigations/embedded/"), ".json")
				mutex := database.BorrowMitigation(hash)

				mutex.Lock()
				database.Remove(file)
				database.Remove(file + ".gz")
				mutex.Unlock()

			}

		}

	}

}

func (database *Database) ReadMitigation(hash string) *structs.Mitigation {

	mutex := database.BorrowMitigation(hash)
	mutex.Lock()
	defer mutex.Unlock()

	return readMitigation(database, hash)

}

func (database *Database) RemoveMitigation(hash string) {

	if utils_path.IsSafeBasename(hash) == true {

		mutex := database.BorrowMitigation(hash)
		mutex.Lock()
		defer mutex.Unlock()

		database.Remove("/mitigations/embedded/" + hash + ".json.gz")
		database.Remove("/mitigations/embedded/" + hash + ".json")

	}

}

func (database *Database) WriteMitigation(mitigation *structs.Mitigation) bool {
	return writeMitigation(database, mitigation)
}

func (database *Database) WriteMitigations() bool {

	result := true

	files := database.Scan("/mitigations/embedded", "*.json.gz")
	files_all := make([]string, 0)

	if len(files) > 0 {

		for _, file := range files {

			if strings.HasPrefix(file, "/mitigations/embedded/") && strings.HasSuffix(file, ".json.gz") {

				hash := strings.TrimSuffix(strings.TrimPrefix(file, "/mitigations/embedded/"), ".json.gz")
				mitigation := database.ReadMitigation(hash)

				if mitigation.IsValid() {
					files_all = append(files_all, "embedded/"+hash+".json.gz")
				}

			}

		}

	}

	for _, file := range database.Scan("/mitigations", "Mitigations_embed.*") {
		database.Remove(file)
	}

	for _, tag := range []string{"guard"} {

		insight_go, insight_embed_go, err := templates.FilesystemCache(&templates.Config{
			Name:     "Mitigations",
			Package:  "mitigations",
			BuildTag: utils_tags.ToBuildTags(tag, database.GetDistributionNames()),
			Struct:   "structs.Mitigation",
			Folder:   "embedded",
			Files:    files_all,
		})

		if err == nil {

			result1 := database.Write("/mitigations/Mitigations.go", insight_go)
			result2 := database.Write("/mitigations/Mitigations_embed."+tag+".go", insight_embed_go)

			if result1 == false || result2 == false {
				result = false
			}

		} else {
			result = false
		}

	}

	return result

}
