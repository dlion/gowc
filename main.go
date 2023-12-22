package main

import (
	"flag"
	"fmt"
	"gowc/reader"
	"os"
	"path/filepath"
)

func main() {

	bytesCountFilepath := flag.String("c", "", "Path to the file you want to count the bytes")
	linesCountFilepath := flag.String("l", "", "Path to the file you want to count the lines")
	flag.Parse()

	if *bytesCountFilepath != "" {
		numBytes, filename := readNumberOfBytesFrom(bytesCountFilepath)
		fmt.Printf("%d %s\n", numBytes, filename)
	}

	if *linesCountFilepath != "" {
		numLines, fileName := readNumberOfLinesFrom(linesCountFilepath)
		fmt.Printf("%d %s", numLines, fileName)
	}

}

func readNumberOfLinesFrom(linesCountFilepath *string) (int64, string) {
	fileName, wcReader := readDirectory(linesCountFilepath)

	numLines, err := wcReader.ReadNumberOfLinesFrom(fileName)
	if err != nil {
		fmt.Printf("can't read %s correctly\n", fileName)
		os.Exit(1)
	}
	return numLines, fileName
}

func readNumberOfBytesFrom(bytesCountFilepath *string) (int64, string) {
	fileName, wcReader := readDirectory(bytesCountFilepath)

	numBytes, err := wcReader.ReadNumberOfBytesFrom(fileName)
	if err != nil {
		fmt.Printf("can't read %s correctly\n", fileName)
		os.Exit(1)
	}
	return numBytes, fileName
}

func readDirectory(linesCountFilepath *string) (string, reader.WcReader) {
	dirPath, fileName := splitPath(*linesCountFilepath)
	filesystemDir := os.DirFS(dirPath)
	wcReader := reader.NewWcReader(filesystemDir)
	return fileName, wcReader
}

func splitPath(path string) (string, string) {
	dir, file := filepath.Split(path)
	dir = filepath.Clean(dir)
	return dir, file
}
