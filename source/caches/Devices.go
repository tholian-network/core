package caches

import "tholian-core/matchers"
import "tholian-core/structs"
import "encoding/json"
import "sort"
import "strings"
import "sync"

func toDeviceIdentifier(value structs.Device) string {

	var result string

	if value.Bus != "" && value.System != nil && value.System.Vendor != "" && value.System.Device != "" {

		if value.Subsystem != nil && value.Subsystem.Vendor != "" && value.Subsystem.Device != "" {
			result = value.Bus + ":" + value.System.Vendor + "-" + value.System.Device + ":" + value.Subsystem.Vendor + "-" + value.Subsystem.Device
		} else {
			result = value.Bus + ":" + value.System.Vendor + "-" + value.System.Device + ":0000-0000"
		}

	}

	return result

}

type Devices struct {
	Devices       map[string]*structs.Device                `json:"devices"`
	mapsystems    map[string]map[string]map[string][]string `json:"-"`
	mapsubsystems map[string]map[string]map[string][]string `json:"-"`
	mutex         *sync.RWMutex                             `json:"-"`
}

func NewDevices() *Devices {

	var cache Devices

	cache.Devices = make(map[string]*structs.Device)
	cache.mapsystems = make(map[string]map[string]map[string][]string, 0)
	cache.mapsubsystems = make(map[string]map[string]map[string][]string, 0)
	cache.mutex = &sync.RWMutex{}

	return &cache

}

func (cache *Devices) Add(value structs.Device) bool {

	cache.mutex.Lock()

	var result bool

	if value.Bus != "" && value.System != nil {

		identifier := toDeviceIdentifier(value)

		cache.Devices[identifier] = &value
		result = true

		if value.System != nil {

			_, ok1 := cache.mapsystems[value.Bus]

			if ok1 == false {
				cache.mapsystems[value.Bus] = make(map[string]map[string][]string, 0)
			}

			_, ok2 := cache.mapsystems[value.Bus][value.System.Vendor]

			if ok2 == false {
				cache.mapsystems[value.Bus][value.System.Vendor] = make(map[string][]string, 0)
			}

			_, ok3 := cache.mapsystems[value.Bus][value.System.Vendor][value.System.Device]

			if ok3 == false {
				cache.mapsystems[value.Bus][value.System.Vendor][value.System.Device] = make([]string, 0)
			}

			cache.mapsystems[value.Bus][value.System.Vendor][value.System.Device] = append(cache.mapsystems[value.Bus][value.System.Vendor][value.System.Device], identifier)

		}

		if value.Subsystem != nil {

			_, ok1 := cache.mapsubsystems[value.Bus]

			if ok1 == false {
				cache.mapsubsystems[value.Bus] = make(map[string]map[string][]string, 0)
			}

			_, ok2 := cache.mapsubsystems[value.Bus][value.Subsystem.Vendor]

			if ok2 == false {
				cache.mapsubsystems[value.Bus][value.Subsystem.Vendor] = make(map[string][]string, 0)
			}

			_, ok3 := cache.mapsubsystems[value.Bus][value.Subsystem.Vendor][value.Subsystem.Device]

			if ok3 == false {
				cache.mapsubsystems[value.Bus][value.Subsystem.Vendor][value.Subsystem.Device] = make([]string, 0)
			}

			cache.mapsubsystems[value.Bus][value.Subsystem.Vendor][value.Subsystem.Device] = append(cache.mapsubsystems[value.Bus][value.Subsystem.Vendor][value.Subsystem.Device], identifier)

		}

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Devices) Get(identifier string) *structs.Device {

	cache.mutex.RLock()

	var result *structs.Device = nil

	if strings.HasPrefix(identifier, "usb:") || strings.HasPrefix(identifier, "pci:") {

		tmp := strings.Split(identifier, ":")

		if len(tmp) == 3 {

			device, ok := cache.Devices[identifier]

			if ok == true {
				result = device
			}

		}

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Devices) Length() int {
	return len(cache.Devices)
}

func (cache *Devices) Query(query matchers.Device) []*structs.Device {

	cache.mutex.RLock()

	result := make([]*structs.Device, 0)
	found := make(map[string]*structs.Device)

	if query.Bus != "any" && query.Vendor != "any" && query.Device != "any" {

		// Cache Optimizations for specified query.Bus, query.Vendor, query.Device

		_, ok1 := cache.mapsystems[query.Bus]

		if ok1 == true {

			_, ok2 := cache.mapsystems[query.Bus][query.Vendor]

			if ok2 == true {

				identifiers, ok3 := cache.mapsystems[query.Bus][query.Vendor][query.Device]

				if ok3 == true {

					for _, identifier := range identifiers {

						device, ok := cache.Devices[identifier]

						if ok == true {

							if device.System != nil {

								matches_name := query.MatchesName(device.Name)
								matches_system_name := query.MatchesName(device.System.Name)

								if matches_name || matches_system_name {
									found[identifier] = device
								}

							}

						}

					}

				}

			}

		}

	} else if query.Bus != "any" && query.Vendor != "any" && query.Device == "any" {

		// Cache Optimizations for specified query.Bus, query.Vendor

		_, ok1 := cache.mapsystems[query.Bus]

		if ok1 == true {

			_, ok2 := cache.mapsystems[query.Bus][query.Vendor]

			if ok2 == true {

				for _, identifiers := range cache.mapsystems[query.Bus][query.Vendor] {

					for _, identifier := range identifiers {

						device, ok := cache.Devices[identifier]

						if ok == true {

							if device.System != nil {

								matches_name := query.MatchesName(device.Name)
								matches_system_name := query.MatchesName(device.System.Name)

								if matches_name || matches_system_name {
									found[identifier] = device
								}

							}

						}

					}

				}

			}

		}

	} else if query.Bus != "any" && query.Vendor == "any" && query.Device == "any" {

		// Cache Optimizations for specified query.Bus

		_, ok1 := cache.mapsystems[query.Bus]

		if ok1 == true {

			for _, devices := range cache.mapsystems[query.Bus] {

				for _, identifiers := range devices {

					for _, identifier := range identifiers {

						device, ok := cache.Devices[identifier]

						if ok == true {

							if device.System != nil {

								matches_name := query.MatchesName(device.Name)
								matches_system_name := query.MatchesName(device.System.Name)

								if matches_name || matches_system_name {
									found[identifier] = device
								}

							}

						}

					}

				}

			}

		}

	} else {

		for identifier, device := range cache.Devices {

			matches_name := query.MatchesName(device.Name)
			matches_vendor := false
			matches_device := false
			matches_bus := query.MatchesBus(device.Bus)

			if device.System != nil {

				if matches_name == false {
					query.MatchesName(device.System.Name)
				}

				if query.Vendor == "any" {
					matches_vendor = true
				} else {
					matches_vendor = query.MatchesVendor(device.System.Vendor)
				}

				if query.Device == "any" {
					matches_device = true
				} else {
					matches_device = query.MatchesDevice(device.System.Device)
				}

			}

			if matches_name && matches_vendor && matches_device && matches_bus {
				found[identifier] = device
			}

		}

	}

	if len(found) > 0 {

		for _, device := range found {
			result = append(result, device)
		}

	}

	if len(result) > 1 {

		sort.Slice(result, func(a int, b int) bool {
			return result[a].Name < result[b].Name
		})

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Devices) QueryBySubsystem(query matchers.Device) []*structs.Device {

	cache.mutex.RLock()

	result := make([]*structs.Device, 0)
	found := make(map[string]*structs.Device)

	if query.Bus != "any" && query.Vendor != "any" && query.Device != "any" {

		// Cache Optimizations for specified query.Bus, query.Vendor, query.Device

		_, ok1 := cache.mapsubsystems[query.Bus]

		if ok1 == true {

			_, ok2 := cache.mapsubsystems[query.Bus][query.Vendor]

			if ok2 == true {

				identifiers, ok3 := cache.mapsubsystems[query.Bus][query.Vendor][query.Device]

				if ok3 == true {

					for _, identifier := range identifiers {

						device, ok := cache.Devices[identifier]

						if ok == true {

							if device.Subsystem != nil {

								matches_name := query.MatchesName(device.Name)
								matches_subsystem_name := query.MatchesName(device.Subsystem.Name)

								if matches_name || matches_subsystem_name {
									found[identifier] = device
								}

							}

						}

					}

				}

			}

		}

	} else if query.Bus != "any" && query.Vendor != "any" && query.Device == "any" {

		// Cache Optimizations for specified query.Bus, query.Vendor

		_, ok1 := cache.mapsubsystems[query.Bus]

		if ok1 == true {

			_, ok2 := cache.mapsubsystems[query.Bus][query.Vendor]

			if ok2 == true {

				for _, identifiers := range cache.mapsubsystems[query.Bus][query.Vendor] {

					for _, identifier := range identifiers {

						device, ok := cache.Devices[identifier]

						if ok == true {

							if device.Subsystem != nil {

								matches_name := query.MatchesName(device.Name)
								matches_subsystem_name := query.MatchesName(device.Subsystem.Name)

								if matches_name || matches_subsystem_name {
									found[identifier] = device
								}

							}

						}

					}

				}

			}

		}

	} else if query.Bus != "any" && query.Vendor == "any" && query.Device == "any" {

		// Cache Optimizations for specified query.Bus

		_, ok1 := cache.mapsubsystems[query.Bus]

		if ok1 == true {

			for _, devices := range cache.mapsubsystems[query.Bus] {

				for _, identifiers := range devices {

					for _, identifier := range identifiers {

						device, ok := cache.Devices[identifier]

						if ok == true {

							if device.Subsystem != nil {

								matches_name := query.MatchesName(device.Name)
								matches_subsystem_name := query.MatchesName(device.Subsystem.Name)

								if matches_name || matches_subsystem_name {
									found[identifier] = device
								}

							}

						}

					}

				}

			}

		}

	} else {

		for identifier, device := range cache.Devices {

			matches_name := query.MatchesName(device.Name)
			matches_vendor := false
			matches_device := false
			matches_bus := query.MatchesBus(device.Bus)

			if device.Subsystem != nil {

				if matches_name == false {
					query.MatchesName(device.Subsystem.Name)
				}

				if query.Vendor == "any" {
					matches_vendor = true
				} else {
					matches_vendor = query.MatchesVendor(device.Subsystem.Vendor)
				}

				if query.Device == "any" {
					matches_device = true
				} else {
					matches_device = query.MatchesDevice(device.Subsystem.Device)
				}

			}

			if matches_name && matches_vendor && matches_device && matches_bus {
				found[identifier] = device
			}

		}

	}

	if len(found) > 0 {

		for _, device := range found {
			result = append(result, device)
		}

	}

	if len(result) > 1 {

		sort.Slice(result, func(a int, b int) bool {
			return result[a].Name < result[b].Name
		})

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Devices) Remove(identifier string) bool {

	cache.mutex.Lock()

	var result bool

	if strings.HasPrefix(identifier, "usb:") || strings.HasPrefix(identifier, "pci:") {

		device, ok := cache.Devices[identifier]

		if ok == true {

			delete(cache.Devices, identifier)
			result = true

			if device.Bus != "" {

				if device.System != nil {

					others, ok1 := cache.mapsystems[device.Bus][device.System.Vendor][device.System.Device]

					if ok1 == true {

						filtered := make([]string, 0)

						for _, other := range others {

							if other != identifier {
								filtered = append(filtered, other)
							}

						}

						if len(filtered) > 0 {
							cache.mapsystems[device.Bus][device.System.Vendor][device.System.Device] = filtered
						} else {

							delete(cache.mapsystems[device.Bus][device.System.Vendor], device.System.Device)

							if len(cache.mapsystems[device.Bus][device.System.Vendor]) == 0 {
								delete(cache.mapsystems[device.Bus], device.System.Vendor)
							}

						}

					}

				}

				if device.Subsystem != nil {

					others, ok1 := cache.mapsubsystems[device.Bus][device.Subsystem.Vendor][device.Subsystem.Device]

					if ok1 == true {

						filtered := make([]string, 0)

						for _, other := range others {

							if other != identifier {
								filtered = append(filtered, other)
							}

						}

						if len(filtered) > 0 {
							cache.mapsubsystems[device.Bus][device.Subsystem.Vendor][device.Subsystem.Device] = filtered
						} else {

							delete(cache.mapsubsystems[device.Bus][device.Subsystem.Vendor], device.Subsystem.Device)

							if len(cache.mapsubsystems[device.Bus][device.Subsystem.Vendor]) == 0 {
								delete(cache.mapsubsystems[device.Bus], device.Subsystem.Vendor)
							}

						}

					}

				}

			}

		}

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Devices) MarshalJSON() ([]byte, error) {

	tmp := make(map[string]*structs.Device)
	identifiers := make([]string, 0)

	for identifier, _ := range cache.Devices {
		identifiers = append(identifiers, identifier)
	}

	sort.Strings(identifiers)

	for _, identifier := range identifiers {
		tmp[identifier] = cache.Devices[identifier]
	}

	return json.MarshalIndent(&struct {
		Devices map[string]*structs.Device `json:"devices"`
	}{
		Devices: tmp,
	}, "", "\t")

}

func (cache *Devices) UnmarshalJSON(data []byte) error {

	var tmp struct {
		Devices map[string]*structs.Device `json:"devices"`
	}

	err := json.Unmarshal(data, &tmp)

	if err != nil {
		return err
	}

	for _, device := range tmp.Devices {
		cache.Add(*device)
	}

	return nil

}
