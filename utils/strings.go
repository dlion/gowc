package utils

import (
	"io/fs"
	"os"
	"path/filepath"
)

func SplitFilepath(filepath string) (fs.FS, string) {
	dirPath, fileName := SplitPath(filepath)
	filesystemDir := os.DirFS(dirPath)
	return filesystemDir, fileName
}

func SplitPath(path string) (string, string) {
	dir, file := filepath.Split(path)
	dir = filepath.Clean(dir)
	return dir, file
}
