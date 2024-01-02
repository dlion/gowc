package utils

import "path/filepath"

func SplitPath(path string) (string, string) {
	dir, file := filepath.Split(path)
	dir = filepath.Clean(dir)
	return dir, file
}
