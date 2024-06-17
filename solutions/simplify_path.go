package solutions

import "strings"

func simplifyPath(path string) string {
	folders := make([]string, 0)
	unfiltered := strings.Split(path, "/")

	for _, u := range unfiltered {
		if u == "" || u == "." {
			continue
		} else if u == ".." {
			if len(folders) == 0 {
				continue
			}
			folders = folders[:len(folders)-1]
		} else {
			folders = append(folders, u)
		}
	}

	return "/" + strings.Join(folders, "/")
}
