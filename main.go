package main

import (
	"flag"
	"fmt"
	"gowc/reader"
	"os"
	"path/filepath"
)

func main() {
	var filePath string
	countBytesFrom(&filePath)
}

func countBytesFrom(filePath *string) {
	flag.StringVar(filePath, "c", "", "Path to the file you want to count the bytes")
	flag.Parse()

	if *filePath != "" {
		numBytes, filename := readNumberOfBytesFrom(filePath)
		fmt.Printf("%d %s\n", numBytes, filename)
	}
}

func readNumberOfBytesFrom(filePath *string) (int64, string) {
	dirPath, fileName := splitPath(*filePath)
	filesystemDir := os.DirFS(dirPath)
	wcReader := reader.NewWcReader(filesystemDir)

	numBytes, err := wcReader.ReadNumberOfBytesFrom(fileName)
	if err != nil {
		fmt.Printf("can't read %s correctly\n", fileName)
		os.Exit(1)
	}
	return numBytes, fileName
}

func splitPath(path string) (string, string) {
	dir, file := filepath.Split(path)

	dir = filepath.Clean(dir)

	return dir, file
}
