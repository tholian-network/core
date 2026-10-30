package insights

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

func (database *Database) CleanupVulnerabilities() {

	files := database.Scan("/vulnerabilities/embedded", "*.json")

	if len(files) > 0 {

		for _, file := range files {

			if strings.HasPrefix(file, "/vulnerabilities/embedded/") && strings.HasSuffix(file, ".json") {

				name := file[26 : len(file)-5]
				vulnerability := readVulnerability(database, name)

				if vulnerability.State != "edited" {
					database.Remove(file)
					database.Remove(file + ".gz")
				}

				// Give GC a hint
				vulnerability = nil

			}

		}

	}

}

func (database *Database) QueryVulnerabilitiesByDistributionVendor(name string) []*structs.Vulnerability {

	result := make([]*structs.Vulnerability, 0)

	database.mutex.RLock()

	identifiers, ok := database.mapdistributionvendors[name]

	if ok == true {

		for _, identifier := range identifiers.Keys() {
			result = append(result, database.ReadVulnerability(identifier))
		}

	}

	database.mutex.RUnlock()

	return result

}

func (database *Database) ReadVulnerability(name string) *structs.Vulnerability {

	mutex := database.BorrowVulnerability(name)
	mutex.Lock()

	result := readVulnerability(database, name)

	return result

}

func (database *Database) RemoveVulnerability(name string) {

	database.mutex.Lock()

	mutex, ok3 := database.mutexes.Vulnerabilities[name]

	if ok3 == true {
		mutex.Unlock()
		delete(database.mutexes.Vulnerabilities, name)
	}

	database.mutex.Unlock()

	database.Remove("/vulnerabilities/embedded/" + name + ".json.gz")
	database.Remove("/vulnerabilities/embedded/" + name + ".json")

}

func (database *Database) WriteVulnerability(vulnerability *structs.Vulnerability) bool {

	var result bool

	if vulnerability != nil {

		result = writeVulnerability(database, vulnerability)

		if len(vulnerability.Distributions) > 0 {

			for _, distribution := range vulnerability.Distributions {

				if distribution.Vendor != "" && distribution.Vendor != "any" {

					_, ok := database.mapdistributionvendors[distribution.Vendor]

					if ok == false {
						database.mapdistributionvendors[distribution.Vendor] = types.NewMap[bool]()
					}

					database.mapdistributionvendors[distribution.Vendor].Set(vulnerability.Name, true)

				}

			}

		}

		mutex := database.BorrowVulnerability(vulnerability.Name)
		mutex.Unlock()

	}

	return result

}

func (database *Database) WriteVulnerabilities() bool {

	files := database.Scan("/vulnerabilities/embedded", "*.json.gz")
	files_all := make([]string, 0)
	files_any_distribution := make([]string, 0)
	files_per_distribution := types.NewSliceMap[string]()

	if len(files) > 0 {

		for _, file := range files {

			if strings.HasPrefix(file, "/vulnerabilities/embedded") && strings.HasSuffix(file, ".json.gz") {

				name := file[26 : len(file)-8]
				vulnerability := readVulnerability(database, name)

				if vulnerability.IsValid() {

					// Strip out /vulnerabilities/ prefix
					files_all = append(files_all, "embedded/"+name+".json.gz")

					if len(vulnerability.Distributions) > 0 {

						files_any_distribution = append(files_any_distribution, "embedded/"+name+".json.gz")

						for _, distribution := range vulnerability.Distributions {

							distribution_name := distribution.Name

							if strings.Contains(distribution_name, "-") {
								distribution_name = distribution_name[0:strings.Index(distribution_name, "-")]
							}

							files_per_distribution.Add(distribution_name, "embedded/"+name+".json.gz")

						}

					}

				}

				// Give GC a hint
				vulnerability = nil

			}

		}

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

			if result1 && result2 {
				// Do Nothing
			} else {
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
				Files:    files_all,
			})

			if err == nil {

				// Replace with non-conflicting variable name
				insight_embed_go = []byte(strings.ReplaceAll(string(insight_embed_go), "embedded_Vulnerabilities", "embedded_Vulnerabilities_"+distribution))

				result1 := database.Write("/vulnerabilities/Vulnerabilities_embed."+tag+"_"+distribution+".go", insight_embed_go)

				if result1 {
					// Do Nothing
				} else {
					result_distributions = false
				}

			}

		}

	}

	if result_all && result_distributions {
		return true
	}

	return false

}
