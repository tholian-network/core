package path

import "strings"

func IsWatchedFile(value string) bool {

	result := false

	folder := ""
	file := ""

	if strings.Contains(value, "/") {
		folder = value[0:strings.LastIndex(value, "/")]
		file = value[strings.LastIndex(value, "/")+1:]
	} else {
		file = value
	}

	watched_folders := []string{
		"/bin/",
		"/lib/",
		"/lib32/",
		"/lib64/",
		"/run",
		"/sbin/",
		"/tmp",
		"/usr/bin/",
		"/usr/lib/",
		"/usr/lib32/",
		"/usr/lib64/",
		"/usr/sbin/",
		"/usr/local/bin/",
		"/usr/local/lib/",
		"/usr/local/lib32/",
		"/usr/local/lib64/",
		"/usr/local/sbin/",
	}

	if folder != "" {

		for w1 := 0; w1 < len(watched_folders); w1++ {

			prefix := watched_folders[w1]

			if strings.HasPrefix(folder, prefix) {
				result = true
			}

			if result == true {
				break
			}

		}

	} else if file != "" {

		if strings.Contains(file, ".so.") {
			result = true
		} else if strings.HasSuffix(file, ".so") {
			result = true
		} else if strings.Contains(file, ".") == false {
			result = true
		}

	}

	return result

}
