package insights

import "tholian-core/caches"
import "tholian-core/types"
import "os"
import "path"
import "strings"
import "sort"
import "sync"

type Database struct {
	Folder                 string                      `json:"folder"`
	Countries              *caches.Countries           `json:"countries"`
	Devices                *caches.Devices             `json:"devices"`
	Distributions          *caches.Distributions       `json:"distributions"`
	Networks               *caches.Networks            `json:"networks"`
	Providers              *caches.Providers           `json:"providers"`
	Proxies                *caches.Proxies             `json:"proxies"`
	Timeslots              *caches.Timeslots           `json:"timeslots"`
	Weaknesses             *caches.Weaknesses          `json:"weaknesses"`
	mapdistributionvendors map[string]*types.Map[bool] `json:"-"`
	mutex                  *sync.RWMutex               `json:"-"`

	// These data types Read/Write directly from Filesystem
	// Credentials     *caches.Credentials     `json:"credentials"`
	// Incidents       *caches.Incidents       `json:"incidents"`
	// Mitigations     *caches.Mitigations     `json:"mitigations"`
	// Products        *caches.Products        `json:"products"`
	// Responses       *caches.Responses       `json:"responses"`
	// Vulnerabilities *caches.Vulnerabilities `json:"vulnerabilities"`

	mutexes struct {
		Credentials     map[string]*sync.Mutex
		Incidents       map[string]*sync.Mutex
		Mitigations     map[string]*sync.Mutex
		Products        map[string]*sync.Mutex
		Responses       map[string]*sync.Mutex
		Vulnerabilities map[string]*sync.Mutex
	} `json:"-"`

}

func NewDatabase(folder string) *Database {

	var database Database

	if strings.HasSuffix(folder, "/") {
		folder = folder[0 : len(folder)-1]
	}

	subfolders := []string{
		"countries",
		"credentials",
		"credentials/embedded",
		"devices",
		"distributions",
		"incidents",
		"incidents/embedded",
		"mitigations",
		"mitigations/embedded",
		"networks",
		"providers",
		"proxies",
		"responses",
		"responses/embedded",
		"timeslots",
		"vulnerabilities",
		"vulnerabilities/embedded",
		"weaknesses",
	}

	stat, err := os.Stat(folder)

	if err == nil && stat.IsDir() {

		for _, subfolder := range subfolders {

			stat1, err1 := os.Stat(folder + "/" + subfolder)

			if err1 == nil && stat1.IsDir() {
				// Do Nothing
			} else {
				os.MkdirAll(folder+"/"+subfolder, 0750)
			}

		}

		database.Folder = folder

	} else {

		err0 := os.MkdirAll(folder, 0750)

		if err0 == nil {

			for _, subfolder := range subfolders {

				stat1, err1 := os.Stat(folder + "/" + subfolder)

				if err1 == nil && stat1.IsDir() {
					// Do Nothing
				} else {
					os.MkdirAll(folder+"/"+subfolder, 0750)
				}

			}

			database.Folder = folder

		}

	}

	database.Countries = caches.NewCountries()
	database.Devices = caches.NewDevices()
	database.Distributions = caches.NewDistributions()
	database.Networks = caches.NewNetworks()
	database.Providers = caches.NewProviders()
	database.Proxies = caches.NewProxies()
	database.Timeslots = caches.NewTimeslots()
	database.Weaknesses = caches.NewWeaknesses()

	database.mapdistributionvendors = make(map[string]*types.Map[bool])
	database.mutex = &sync.RWMutex{}
	database.mutexes.Credentials = make(map[string]*sync.Mutex)
	database.mutexes.Incidents = make(map[string]*sync.Mutex)
	database.mutexes.Mitigations = make(map[string]*sync.Mutex)
	database.mutexes.Products = make(map[string]*sync.Mutex)
	database.mutexes.Responses = make(map[string]*sync.Mutex)
	database.mutexes.Vulnerabilities = make(map[string]*sync.Mutex)

	return &database

}

func (database *Database) Reset() {

	database.Countries = caches.NewCountries()
	database.Devices = caches.NewDevices()
	database.Distributions = caches.NewDistributions()
	database.Networks = caches.NewNetworks()
	database.Providers = caches.NewProviders()
	database.Proxies = caches.NewProxies()
	database.Timeslots = caches.NewTimeslots()
	database.Weaknesses = caches.NewWeaknesses()

	database.mapdistributionvendors = make(map[string]*types.Map[bool])
	database.mutex = &sync.RWMutex{}
	database.mutexes.Credentials = make(map[string]*sync.Mutex)
	database.mutexes.Incidents = make(map[string]*sync.Mutex)
	database.mutexes.Mitigations = make(map[string]*sync.Mutex)
	database.mutexes.Products = make(map[string]*sync.Mutex)
	database.mutexes.Responses = make(map[string]*sync.Mutex)
	database.mutexes.Vulnerabilities = make(map[string]*sync.Mutex)

}

func (database *Database) Exists(file string) bool {

	if strings.HasPrefix(file, "/") == false {
		file = "/" + file
	}

	_, err := os.Stat(database.Folder + file)

	if err == nil {
		return true
	}

	return false

}

func (database *Database) Read(file string) []byte {

	if strings.HasPrefix(file, "/") == false {
		file = "/" + file
	}

	buffer, err := os.ReadFile(database.Folder + file)

	if err == nil {
		return buffer
	}

	return nil

}

func (database *Database) Remove(file string) bool {

	result := false

	if strings.HasPrefix(file, "/") == false {
		file = "/" + file
	}

	stat1, err1 := os.Stat(database.Folder + file)

	if err1 == nil && stat1.IsDir() == false {

		err2 := os.Remove(database.Folder + file)

		if err2 == nil {
			result = true
		}

	}

	return result

}

func (database *Database) Scan(folder string, pattern string) []string {

	result := make([]string, 0)

	if strings.HasPrefix(folder, "/") == false {
		folder = "/" + folder
	}

	if strings.HasSuffix(folder, "/") {
		folder = folder[0 : len(folder)-1]
	}

	if pattern == "" {
		pattern = "*"
	}

	stat, err1 := os.Stat(database.Folder + folder)

	if err1 == nil && stat.IsDir() == true {

		entries, err2 := os.ReadDir(database.Folder + folder)

		if err2 == nil {

			for _, entry := range entries {

				filename := entry.Name()
				matches_pattern := false

				if pattern == "*" {
					matches_pattern = true
				} else if strings.HasPrefix(pattern, "*") {
					matches_pattern = strings.HasSuffix(filename, pattern[1:])
				} else if strings.HasSuffix(pattern, "*") {
					matches_pattern = strings.HasPrefix(filename, pattern[0:len(pattern)-1])
				}

				if matches_pattern {
					result = append(result, folder+"/"+filename)
				}

			}

		}

	}

	sort.Strings(result)

	return result

}

func (database *Database) Write(file string, buffer []byte) bool {

	if strings.HasPrefix(file, "/") == false {
		file = "/" + file
	}

	folder := path.Dir(file)
	stat1, err1 := os.Stat(database.Folder + folder)

	if err1 == nil && stat1.IsDir() == true {

		err2 := os.WriteFile(database.Folder+file, buffer, 0666)

		if err2 == nil {
			return true
		}

	} else {

		err2 := os.MkdirAll(database.Folder+folder, 0755)

		if err2 == nil {

			err3 := os.WriteFile(database.Folder+file, buffer, 0666)

			if err3 == nil {
				return true
			}

		}

	}

	return false

}

func (database *Database) Flush() bool {

	result := false
	result_countries := false
	result_credentials := false
	result_devices := false
	result_distributions := false
	result_incidents := false
	result_mitigations := false
	result_networks := false
	result_products := false
	result_providers := false
	result_proxies := false
	result_responses := false
	result_timeslots := false
	result_vulnerabilities := false
	result_weaknesses := false

	stat, err := os.Stat(database.Folder)

	if err == nil && stat.IsDir() {

		if database.Countries.Length() > 0 {
			result_countries = database.WriteCountries()
		} else {
			result_countries = true
		}

		// Workflow requires direct WriteCredential() calls
		result_credentials = database.WriteCredentials()

		if database.Devices.Length() > 0 {
			result_devices = database.WriteDevices()
		} else {
			result_devices = true
		}

		if database.Distributions.Length() > 0 {
			result_distributions = database.WriteDistributions()
		} else {
			result_distributions = true
		}

		// Workflow requires direct Write() calls to Filesystem
		// TODO: result_incidents = database.WriteIncidents()
		result_incidents = true

		// Workflow requires direct Write() calls to Filesystem
		// TODO: result_mitigations = database.WriteMitigations()
		result_mitigations = true

		if database.Networks.Length() > 0 {
			result_networks = database.WriteNetworks()
		} else {
			result_networks = true
		}

		// Workflow requires direct Write() calls to Filesystem
		// TODO: result_products = database.WriteProducts()
		result_products = true

		if database.Providers.Length() > 0 {
			result_providers = database.WriteProviders()
		} else {
			result_providers = true
		}

		if database.Proxies.Length() > 0 {
			result_proxies = database.WriteProxies()
		} else {
			result_proxies = true
		}

		// Workflow requires direct Write() calls to Filesystem
		// TODO: result_responses = database.WriteResponses()
		result_responses = true

		if database.Timeslots.Length() > 0 {
			result_timeslots = database.WriteTimeslots()
		} else {
			result_timeslots = true
		}

		// Workflow requires direct Write() calls to Filesystem
		// TODO: result_vulnerabilities = database.WriteVulnerabilities()
		result_vulnerabilities = true

		if database.Weaknesses.Length() > 0 {
			result_weaknesses = database.WriteWeaknesses()
		} else {
			result_weaknesses = true
		}

	}

	if result_countries == true &&
		result_credentials == true &&
		result_devices == true &&
		result_distributions == true &&
		result_incidents == true &&
		result_mitigations == true &&
		result_networks == true &&
		result_products == true &&
		result_providers == true &&
		result_proxies == true &&
		result_responses == true &&
		result_timeslots == true &&
		result_vulnerabilities == true &&
		result_weaknesses == true {
		result = true
	}

	return result

}
