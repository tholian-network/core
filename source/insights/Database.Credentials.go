package insights

import "tholian-core/insights/templates"
import "tholian-core/structs"
import utils_tags "tholian-core/utils/tags"
import "sort"
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

func (database *Database) ReadCredential(name string) *structs.Credential {

	mutex := database.BorrowCredential(name)
	mutex.Lock()

	result := readCredential(database, name)

	return result

}

func (database *Database) RemoveCredential(name string) {

	database.mutex.Lock()

	mutex, ok3 := database.mutexes.Credentials[name]

	if ok3 == true {
		mutex.Unlock()
		delete(database.mutexes.Credentials, name)
	}

	database.mutex.Unlock()

	database.Remove("/credentials/embedded/" + name + ".json.gz")
	database.Remove("/credentials/embedded/" + name + ".json")

}

func (database *Database) WriteCredential(credential *structs.Credential) bool {

	var result bool

	if credential != nil {

		result = writeCredential(database, credential)

		mutex := database.BorrowCredential(credential.Name)
		mutex.Unlock()

	}

	return result

}

func (database *Database) WriteCredentials() bool {

	var result bool

	files_all := make([]string, 0)
	files := database.Scan("/credentials/embedded", "*.json.gz")

	if len(files) > 0 {

		sort.Strings(files)

		for _, file := range files {

			if strings.HasPrefix(file, "/credentials/embedded") {
				// Strip out /credentials/ prefix
				files_all = append(files_all, file[13:])
			}

		}

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
			result2 := database.Write("/credentials/Credentials_embed." + tag + ".go", insight_embed_go)

			if result1 && result2 {
				result = true
			}

		}

	}

	return result

}
