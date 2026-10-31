package insights

import utils_path "tholian-core/utils/path"
import "tholian-core/insights/templates"
import "tholian-core/structs"
import utils_tags "tholian-core/utils/tags"
import "strings"
import "sync"

func (database *Database) BorrowIncident(hash string) *sync.Mutex {

	database.mutex.RLock()
	mutex, ok := database.mutexes.Incidents[hash]
	database.mutex.RUnlock()

	if ok == false {

		database.mutex.Lock()
		database.mutexes.Incidents[hash] = &sync.Mutex{}
		mutex = database.mutexes.Incidents[hash]
		database.mutex.Unlock()

	}

	return mutex

}

func (database *Database) CleanupIncidents() {

	files := database.Scan("/incidents/embedded", "*.json")

	if len(files) > 0 {

		for _, file := range files {

			if strings.HasPrefix(file, "/incidents/embedded/") && strings.HasSuffix(file, ".json") {

				hash := strings.TrimSuffix(strings.TrimPrefix(file, "/incidents/embedded/"), ".json")
				mutex := database.BorrowIncident(hash)

				mutex.Lock()
				database.Remove(file)
				database.Remove(file + ".gz")
				mutex.Unlock()

			}

		}

	}

}

func (database *Database) ReadIncident(hash string) *structs.Incident {

	mutex := database.BorrowIncident(hash)
	mutex.Lock()
	defer mutex.Unlock()

	return readIncident(database, hash)

}

func (database *Database) RemoveIncident(hash string) {

	if utils_path.IsSafeBasename(hash) == true {

		mutex := database.BorrowIncident(hash)
		mutex.Lock()
		defer mutex.Unlock()

		database.Remove("/incidents/embedded/" + hash + ".json.gz")
		database.Remove("/incidents/embedded/" + hash + ".json")

	}

}

func (database *Database) WriteIncident(incident *structs.Incident) bool {
	return writeIncident(database, incident)
}

func (database *Database) WriteIncidents() bool {

	result := true

	files := database.Scan("/incidents/embedded", "*.json.gz")
	files_all := make([]string, 0)

	if len(files) > 0 {

		for _, file := range files {

			if strings.HasPrefix(file, "/incidents/embedded/") && strings.HasSuffix(file, ".json.gz") {

				hash := strings.TrimSuffix(strings.TrimPrefix(file, "/incidents/embedded/"), ".json.gz")
				incident := database.ReadIncident(hash)

				if incident.IsValid() {
					files_all = append(files_all, "embedded/"+hash+".json.gz")
				}

			}

		}

	}

	for _, file := range database.Scan("/incidents", "Incidents_embed.*") {
		database.Remove(file)
	}

	for _, tag := range []string{"guard", "oversight"} {

		insight_go, insight_embed_go, err := templates.FilesystemCache(&templates.Config{
			Name:     "Incidents",
			Package:  "incidents",
			BuildTag: utils_tags.ToBuildTags(tag, database.GetDistributionNames()),
			Struct:   "structs.Incident",
			Folder:   "embedded",
			Files:    files_all,
		})

		if err == nil {

			result1 := database.Write("/incidents/Incidents.go", insight_go)
			result2 := database.Write("/incidents/Incidents_embed."+tag+".go", insight_embed_go)

			if result1 == false || result2 == false {
				result = false
			}

		} else {
			result = false
		}

	}

	return result

}
