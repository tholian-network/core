package insights

import utils_path "tholian-core/utils/path"
import "tholian-core/insights/templates"
import "tholian-core/structs"
import utils_tags "tholian-core/utils/tags"
import "strings"
import "sync"

func (database *Database) BorrowCredential(name string) *sync.Mutex {

	database.mutex.RLock()
	mutex, ok := database.mutexes.Credentials[name]
	database.mutex.RUnlock()

	if ok == false {

		database.mutex.Lock()
		database.mutexes.Credentials[name] = &sync.Mutex{}
		mutex = database.mutexes.Credentials[name]
		database.mutex.Unlock()

	}

	return mutex

}

func (database *Database) CleanupCredentials() {

	files := database.Scan("/credentials/embedded", "*.json")

	if len(files) > 0 {

		for _, file := range files {

			if strings.HasPrefix(file, "/credentials/embedded/") && strings.HasSuffix(file, ".json") {

				name := strings.TrimSuffix(strings.TrimPrefix(file, "/credentials/embedded/"), ".json")
				mutex := database.BorrowCredential(name)

				mutex.Lock()

				credential := readCredential(database, name)

				if credential.IsValid() == false {
					database.Remove(file)
					database.Remove(file + ".gz")
				}

				mutex.Unlock()

			}

		}

	}

}

func (database *Database) ReadCredential(name string) *structs.Credential {

	mutex := database.BorrowCredential(name)
	mutex.Lock()
	defer mutex.Unlock()

	return readCredential(database, name)

}

func (database *Database) RemoveCredential(name string) {

	if utils_path.IsSafeBasename(name) == true {

		mutex := database.BorrowCredential(name)
		mutex.Lock()
		defer mutex.Unlock()

		database.Remove("/credentials/embedded/" + name + ".json.gz")
		database.Remove("/credentials/embedded/" + name + ".json")

	}

}

func (database *Database) WriteCredential(credential *structs.Credential) bool {
	return writeCredential(database, credential)
}

func (database *Database) WriteCredentials() bool {

	result := true

	files_all := make([]string, 0)
	files := database.Scan("/credentials/embedded", "*.json.gz")

	if len(files) > 0 {

		for _, file := range files {

			if strings.HasPrefix(file, "/credentials/embedded/") && strings.HasSuffix(file, ".json.gz") {

				name := strings.TrimSuffix(strings.TrimPrefix(file, "/credentials/embedded/"), ".json.gz")
				credential := database.ReadCredential(name)

				if credential.IsValid() {
					files_all = append(files_all, "embedded/"+name+".json.gz")
				}

			}

		}

	}

	for _, file := range database.Scan("/credentials", "Credentials_embed.*") {
		database.Remove(file)
	}

	for _, tag := range []string{"oversight"} {

		insight_go, insight_embed_go, err := templates.FilesystemCache(&templates.Config{
			Name:     "Credentials",
			Package:  "credentials",
			BuildTag: utils_tags.ToBuildTags(tag, database.GetDistributionNames()),
			Struct:   "structs.Credential",
			Folder:   "embedded",
			Files:    files_all,
		})

		if err == nil {

			result1 := database.Write("/credentials/Credentials.go", insight_go)
			result2 := database.Write("/credentials/Credentials_embed."+tag+".go", insight_embed_go)

			if result1 == false || result2 == false {
				result = false
			}

		} else {
			result = false
		}

	}

	return result

}
