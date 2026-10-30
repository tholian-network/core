package insights

import "tholian-core/insights/templates"
import "tholian-core/structs"
import utils_tags "tholian-core/utils/tags"
import "strings"
import "sync"

func (database *Database) BorrowMitigation(name string) *sync.Mutex {

	database.mutex.RLock()
	mutex, ok := database.mutexes.Mitigations[name]
	database.mutex.RUnlock()

	if ok == false {

		database.mutex.Lock()
		database.mutexes.Mitigations[name] = &sync.Mutex{}
		mutex = database.mutexes.Mitigations[name]
		database.mutex.Unlock()

	}

	return mutex

}

func (database *Database) CleanupMitigations() {

	files := database.Scan("/mitigations/embedded", "*.json")

	if len(files) > 0 {

		for _, file := range files {

			if strings.HasPrefix(file, "/mitigations/embedded/") && strings.HasSuffix(file, ".json") {
				database.Remove(file)
				database.Remove(file + ".gz")
			}

		}

	}

}

func (database *Database) WriteMitigation(mitigation *structs.Mitigation) bool {

	var result bool

	mutex := database.BorrowMitigation(mitigation.Hash())
	mutex.Lock()

	if mitigation != nil {
		result = writeMitigation(database, mitigation)
	}

	mutex.Unlock()

	return result


}

func (database *Database) WriteMitigations() bool {

	var result bool

	files := database.Scan("/mitigations/embedded", "*.json.gz")
	files_all := make([]string, 0)

	if len(files) > 0 {

		for _, file := range files {

			if strings.HasPrefix(file, "/mitigations/embedded") && strings.HasSuffix(file, ".json.gz") {

				// Strip out /mitigations/ prefix
				name := file[22 : len(file)-8]
				files_all = append(files_all, "embedded/"+name+".json.gz")

			}

		}

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
			result2 := database.Write("/mitigations/Mitigations_embed." + tag + ".go", insight_embed_go)

			if result1 && result2 {
				result = true
			}

		}

	}

	return result

}
