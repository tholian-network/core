package insights

import "os"
import "path"
import "sort"
import "strings"

type Cache struct {
	Folder string `json:"folder"`
}

func NewCache(folder string) *Cache {

	var cache Cache

	if strings.HasSuffix(folder, "/") {
		folder = folder[0 : len(folder)-1]
	}

	stat, err1 := os.Stat(folder)

	if err1 == nil && stat.IsDir() {

		cache.Folder = folder

	} else {

		err2 := os.MkdirAll(folder, 0750)

		if err2 == nil {
			cache.Folder = folder
		}

	}

	return &cache

}

func (cache *Cache) IsValid() bool {

	result := false

	if cache.Folder != "" && strings.HasPrefix(cache.Folder, "/") {
		result = true
	}

	return result

}

func (cache *Cache) Exists(file string) bool {

	result := false

	if strings.HasPrefix(file, "/") == false {
		file = "/" + file
	}

	_, err := os.Stat(cache.Folder + file)

	if err == nil {
		result = true
	}

	return result

}

func (cache *Cache) Read(file string) []byte {

	if strings.HasPrefix(file, "/") == false {
		file = "/" + file
	}

	buffer, err := os.ReadFile(cache.Folder + file)

	if err == nil {
		return buffer
	}

	return nil

}

func (cache *Cache) Remove(file string) bool {

	result := false

	if strings.HasPrefix(file, "/") == false {
		file = "/" + file
	}

	err := os.Remove(cache.Folder + file)

	if err == nil {
		result = true
	}

	return result

}

func (cache *Cache) Scan(folder string, pattern string) []string {

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

	stat, err1 := os.Stat(cache.Folder + folder)

	if err1 == nil && stat.IsDir() {

		entries, err2 := os.ReadDir(cache.Folder + folder)

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
					result = append(result, folder + "/" + filename)
				}

			}

		}

	}

	sort.Strings(result)

	return result

}

func (cache *Cache) Write(file string, buffer []byte) bool {

	result := false

	if strings.HasPrefix(file, "/") == false {
		file = "/" + file
	}

	folder := path.Dir(file)
	stat1, err1 := os.Stat(cache.Folder + folder)

	if err1 == nil && stat1.IsDir() {

		err2 := os.WriteFile(cache.Folder+file, buffer, 0666)

		if err2 == nil {
			result = true
		}

	} else {

		err2 := os.MkdirAll(cache.Folder+folder, 0755)

		if err2 == nil {

			err3 := os.WriteFile(cache.Folder+file, buffer, 0666)

			if err3 == nil {
				result = true
			}

		}

	}

	return result

}
