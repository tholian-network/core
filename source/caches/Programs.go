package caches

import "tholian-core/matchers"
import "tholian-core/structs"
import "encoding/json"
import "slices"
import "sort"
import "sync"

type Programs struct {
	Programs    map[uint]*structs.Program `json:"programs"`
	mapcommands map[string][]uint         `json:"-"`
	mapnames    map[string][]uint         `json:"-"`
	mutex       *sync.RWMutex             `json:"-"`
}

func NewPrograms() *Programs {

	var cache Programs

	cache.Programs = make(map[uint]*structs.Program)
	cache.mapcommands = make(map[string][]uint, 0)
	cache.mapnames = make(map[string][]uint, 0)
	cache.mutex = &sync.RWMutex{}

	return &cache

}

func (cache *Programs) Add(value structs.Program) bool {

	cache.mutex.Lock()

	var result bool

	if value.PID != 0 && value.Name != "" && value.Command != "" {

		_, ok1 := cache.Programs[value.PID]

		if ok1 == false {

			cache.Programs[value.PID] = &value
			result = true

			_, ok2 := cache.mapcommands[value.Name]

			if ok2 == false {
				cache.mapcommands[value.Name] = make([]uint, 0)
			}

			_, ok3 := cache.mapnames[value.Name]

			if ok3 == false {
				cache.mapnames[value.Name] = make([]uint, 0)
			}

			cache.mapcommands[value.Command] = append(cache.mapcommands[value.Command], value.PID)
			cache.mapnames[value.Name] = append(cache.mapnames[value.Name], value.PID)

		}

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Programs) Get(value uint) *structs.Program {

	cache.mutex.RLock()

	var result *structs.Program = nil

	if value != 0 {

		program, ok := cache.Programs[value]

		if ok == true {
			result = program
		}

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Programs) Length() int {
	return len(cache.Programs)
}

func (cache *Programs) Query(query matchers.Program) []*structs.Program {

	cache.mutex.RLock()

	result := make([]*structs.Program, 0)
	found := make(map[uint]*structs.Program)

	if query.Command != "any" {

		// Cache Optimizations for specified query.Command

		_, ok1 := cache.mapcommands[query.Command]

		if ok1 == true {

			for _, pid := range cache.mapcommands[query.Command] {

				program, ok := cache.Programs[pid]

				if ok == true {

					matches_name := query.MatchesName(program.Name)

					if matches_name {
						found[pid] = program
					}

				}

			}

		}

	} else if query.Name != "any" {

		// Cache Optimizations for specified query.Name

		_, ok1 := cache.mapnames[query.Name]

		if ok1 == true {

			for _, pid := range cache.mapnames[query.Name] {

				program, ok := cache.Programs[pid]

				if ok == true {

					matches_command := query.MatchesCommand(program.Command)

					if matches_command {
						found[pid] = program
					}

				}

			}

		}

	} else {

		for pid, program := range cache.Programs {

			matches_name := query.MatchesName(program.Name)
			matches_command := query.MatchesCommand(program.Command)

			if matches_name && matches_command {
				found[pid] = program
			}

		}

	}

	if len(found) > 0 {

		for _, program := range found {
			result = append(result, program)
		}

	}

	if len(result) > 1 {

		sort.Slice(result, func(a int, b int) bool {
			return result[a].PID < result[b].PID
		})

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Programs) QueryByEnvironmentVariable(query_name string, query_value string) []*structs.Program {

	cache.mutex.RLock()

	result := make([]*structs.Program, 0)
	found := make(map[uint]*structs.Program)

	if query_name != "" && query_name != "any" {

		if query_value == "any" {

			for pid, program := range cache.Programs {

				_, ok := program.Environment[query_name]

				if ok == true {
					found[pid] = program
				}

			}

		} else if query_value != "" {

			for pid, program := range cache.Programs {

				val, ok := program.Environment[query_name]

				if ok == true && val == query_value {
					found[pid] = program
				}

			}

		}

	}

	if len(found) > 0 {

		for _, program := range found {
			result = append(result, program)
		}

	}

	if len(result) > 1 {

		sort.Slice(result, func(a int, b int) bool {
			return result[a].PID < result[b].PID
		})

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Programs) Remove(pid uint) bool {

	cache.mutex.Lock()

	var result bool

	if pid != 0 {

		program, ok := cache.Programs[pid]

		if ok == true {

			delete(cache.Programs, pid)
			result = true

			others, ok2 := cache.mapcommands[program.Command]

			if ok2 == true {

				filtered := make([]uint, 0)

				for _, other := range others {

					if other != pid {
						filtered = append(filtered, other)
					}

				}

				if len(filtered) > 0 {
					cache.mapcommands[program.Command] = filtered
				} else {
					delete(cache.mapcommands, program.Command)
				}

			}

			others, ok3 := cache.mapnames[program.Name]

			if ok3 == true {

				filtered := make([]uint, 0)

				for _, other := range others {

					if other != pid {
						filtered = append(filtered, other)
					}

				}

				if len(filtered) > 0 {
					cache.mapnames[program.Name] = filtered
				} else {
					delete(cache.mapnames, program.Name)
				}

			}

		}

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Programs) MarshalJSON() ([]byte, error) {

	tmp := make(map[uint]*structs.Program)
	pids := make([]uint, 0)

	for pid, _ := range cache.Programs {
		pids = append(pids, pid)
	}

	slices.Sort(pids)

	for _, pid := range pids {
		tmp[pid] = cache.Programs[pid]
	}

	return json.MarshalIndent(&struct {
		Programs map[uint]*structs.Program `json:"programs"`
	}{
		Programs: tmp,
	}, "", "\t")

}

func (cache *Programs) UnmarshalJSON(data []byte) error {

	var tmp struct {
		Programs map[uint]*structs.Program `json:"programs"`
	}

	err := json.Unmarshal(data, &tmp)

	if err != nil {
		return err
	}

	for _, program := range tmp.Programs {
		cache.Add(*program)
	}

	return nil

}
