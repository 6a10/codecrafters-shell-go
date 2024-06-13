package commands

import "os"
import "path/filepath"

var paths = os.Getenv("PATH")
var pathsList = filepath.SplitList(paths)

// func LoadPaths() []string {
// }

// search name in all directories of paths
// return full path of first matching name
func FindExec(name string) string {
	if paths == "" {
		return ""
	}
	for _, p := range pathsList {
		// list files
		entries, dirErr := os.ReadDir(p)
		if dirErr != nil {
			continue
		}
		// if file == name then return full path's file
		for _, entry := range entries {
			if !entry.IsDir() {
				if entry.Name() == name {
					return filepath.Join(p, name)
				}
			}
		}
	}
	return ""
}
