package main

import (
	"flag"
	"fmt"
	"gowc/parser"
	"gowc/reader"
	"os"
	"path/filepath"
)

func main() {
	parameters := getParameters()

	if parameters["c"] != "" {
		numBytes, filename := readNumberOfBytesFrom(parameters["c"])
		fmt.Printf("%d %s\n", numBytes, filename)
	}

	if parameters["l"] != "" {
		numLines, fileName := readNumberOfLinesFrom(parameters["l"])
		fmt.Printf("%d %s\n", numLines, fileName)
	}

}

func getParameters() map[string]string {
	f := flag.NewFlagSet("gowc parameters", flag.ExitOnError)

	f.String("c", "", "Path to the file you want to count the bytes")
	f.String("l", "", "Path to the file you want to count the lines")

	p := parser.NewParser(f)

	parameters, err := p.GetParameters()
	if err != nil {
		fmt.Printf("Can't read the parameters correctly\n")
		os.Exit(1)
	}
	return parameters
}

func readNumberOfLinesFrom(linesCountFilepath string) (int64, string) {
	fileName, wcReader := readDirectory(linesCountFilepath)

	numLines, err := wcReader.ReadNumberOfLinesFrom(fileName)
	if err != nil {
		fmt.Printf("can't read %s correctly\n", fileName)
		os.Exit(1)
	}
	return numLines, fileName
}

func readNumberOfBytesFrom(bytesCountFilepath string) (int64, string) {
	fileName, wcReader := readDirectory(bytesCountFilepath)

	numBytes, err := wcReader.ReadNumberOfBytesFrom(fileName)
	if err != nil {
		fmt.Printf("can't read %s correctly\n", fileName)
		os.Exit(1)
	}
	return numBytes, fileName
}

func readDirectory(linesCountFilepath string) (string, reader.WcReader) {
	dirPath, fileName := splitPath(linesCountFilepath)
	filesystemDir := os.DirFS(dirPath)
	wcReader := reader.NewWcReader(filesystemDir)
	return fileName, wcReader
}

func splitPath(path string) (string, string) {
	dir, file := filepath.Split(path)
	dir = filepath.Clean(dir)
	return dir, file
}
