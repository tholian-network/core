package insights

import utils_path "tholian-core/utils/path"
import "tholian-core/insights/templates"
import "tholian-core/structs"
import "tholian-core/types"
import "sort"
import "strings"
import "sync"

func (database *Database) BorrowVulnerability(name string) *sync.Mutex {

	database.mutex.RLock()
	mutex, ok := database.mutexes.Vulnerabilities[name]
	database.mutex.RUnlock()

	if ok == false {

		database.mutex.Lock()
		database.mutexes.Vulnerabilities[name] = &sync.Mutex{}
		mutex = database.mutexes.Vulnerabilities[name]
		database.mutex.Unlock()

	}

	return mutex

}

func (database *Database) unindexVulnerability(name string) {

	for vendor, identifiers := range database.mapdistributionvendors {
		identifiers.Remove(name)
		if identifiers.Length() == 0 {
			delete(database.mapdistributionvendors, vendor)
		}
	}

}

func (database *Database) CleanupVulnerabilities() {

	files := database.Scan("/vulnerabilities/embedded", "*.json")

	if len(files) > 0 {

		for _, file := range files {

			if strings.HasPrefix(file, "/vulnerabilities/embedded/") && strings.HasSuffix(file, ".json") {

				name := strings.TrimSuffix(strings.TrimPrefix(file, "/vulnerabilities/embedded/"), ".json")
				mutex := database.BorrowVulnerability(name)

				mutex.Lock()

				vulnerability := readVulnerability(database, name)

				if vulnerability.State != "edited" {
					database.Remove(file)
					database.Remove(file + ".gz")
					database.mutex.Lock()
					database.unindexVulnerability(name)
					database.mutex.Unlock()
				}

				mutex.Unlock()

			}

		}

	}

}

func (database *Database) QueryVulnerabilitiesByDistributionVendor(name string) []*structs.Vulnerability {

	result := make([]*structs.Vulnerability, 0)
	keys := make([]string, 0)

	database.mutex.RLock()

	identifiers, ok := database.mapdistributionvendors[name]

	if ok == true {
		keys = identifiers.Keys()
	}

	database.mutex.RUnlock()

	for _, identifier := range keys {
		result = append(result, database.ReadVulnerability(identifier))
	}

	return result

}

func (database *Database) ReadVulnerability(name string) *structs.Vulnerability {

	mutex := database.BorrowVulnerability(name)
	mutex.Lock()
	defer mutex.Unlock()

	return readVulnerability(database, name)

}

func (database *Database) RemoveVulnerability(name string) {

	if utils_path.IsSafeBasename(name) == true {

		mutex := database.BorrowVulnerability(name)
		mutex.Lock()

		database.Remove("/vulnerabilities/embedded/" + name + ".json.gz")
		database.Remove("/vulnerabilities/embedded/" + name + ".json")

		mutex.Unlock()

		database.mutex.Lock()
		database.unindexVulnerability(name)
		database.mutex.Unlock()

	}

}

func (database *Database) WriteVulnerability(vulnerability *structs.Vulnerability) bool {

	if writeVulnerability(database, vulnerability) == false {
		return false
	}

	name := vulnerability.Name

	database.mutex.Lock()

	database.unindexVulnerability(name)

	if len(vulnerability.Distributions) > 0 {

		for _, distribution := range vulnerability.Distributions {

			if distribution.Vendor != "" && distribution.Vendor != "any" {

				_, ok := database.mapdistributionvendors[distribution.Vendor]

				if ok == false {
					database.mapdistributionvendors[distribution.Vendor] = types.NewMap[bool]()
				}

				database.mapdistributionvendors[distribution.Vendor].Set(name, true)

			}

		}

	}

	database.mutex.Unlock()

	return true

}

func (database *Database) WriteVulnerabilities() bool {

	files := database.Scan("/vulnerabilities/embedded", "*.json.gz")
	files_all := make([]string, 0)
	files_per_distribution := types.NewSliceMap[string]()

	if len(files) > 0 {

		for _, file := range files {

			if strings.HasPrefix(file, "/vulnerabilities/embedded/") && strings.HasSuffix(file, ".json.gz") {

				name := strings.TrimSuffix(strings.TrimPrefix(file, "/vulnerabilities/embedded/"), ".json.gz")
				vulnerability := database.ReadVulnerability(name)

				if vulnerability.IsValid() {

					files_all = append(files_all, "embedded/"+name+".json.gz")

					if len(vulnerability.Distributions) > 0 {

						for _, distribution := range vulnerability.Distributions {

							distribution_name := distribution.Name

							if strings.Contains(distribution_name, "-") {
								distribution_name = distribution_name[0:strings.Index(distribution_name, "-")]
							}

							files_per_distribution.Add(distribution_name, "embedded/"+name+".json.gz")

						}

					}

				}

			}

		}

	}

	for _, file := range database.Scan("/vulnerabilities", "Vulnerabilities_embed.*") {
		database.Remove(file)
	}

	result_all := true

	for _, tag := range []string{"scout", "guard", "oversight"} {

		insight_go, insight_embed_go, err := templates.FilesystemCache(&templates.Config{
			Name:     "Vulnerabilities",
			Package:  "vulnerabilities",
			BuildTag: tag,
			Struct:   "structs.Vulnerability",
			Folder:   "embedded",
			Files:    files_all,
		})

		if err == nil {

			result1 := database.Write("/vulnerabilities/Vulnerabilities.go", insight_go)
			result2 := database.Write("/vulnerabilities/Vulnerabilities_embed."+tag+".go", insight_embed_go)

			if result1 == false || result2 == false {
				result_all = false
			}

		} else {
			result_all = false
		}

	}

	result_distributions := true

	for _, tag := range []string{"scout", "guard"} {

		for _, distribution := range files_per_distribution.Keys() {

			files, _ := files_per_distribution.Get(distribution)

			if len(files) > 1 {
				sort.Strings(files)
			}

			_, insight_embed_go, err := templates.FilesystemCache(&templates.Config{
				Name:     "Vulnerabilities",
				Package:  "vulnerabilities",
				BuildTag: tag + "_" + distribution,
				Struct:   "structs.Vulnerability",
				Folder:   "embedded",
				Files:    files,
			})

			if err == nil {

				insight_embed_go = []byte(strings.ReplaceAll(string(insight_embed_go), "embedded_Vulnerabilities", "embedded_Vulnerabilities_"+distribution))

				result1 := database.Write("/vulnerabilities/Vulnerabilities_embed."+tag+"_"+distribution+".go", insight_embed_go)

				if result1 == false {
					result_distributions = false
				}

			} else {
				result_distributions = false
			}

		}

	}

	if result_all && result_distributions {
		return true
	}

	return false

}
