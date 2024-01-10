package utils

import (
	"path/filepath"
)

func SplitFilepath(path string) string {
	dir, filename := filepath.Split(path)
	return dir + filename
}
