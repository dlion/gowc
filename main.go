package main

import (
	"fmt"
	"gowc/parameters"
	readerBytes "gowc/reader/bytes"
	readerLines "gowc/reader/lines"
	"os"
)

func main() {
	parametersMap := parameters.GetParameters()

	if parametersMap["c"] != "" {
		numBytes, filename := readNumberOfBytesFrom(parametersMap["c"])
		fmt.Printf("%d %s\n", numBytes, filename)
	}

	if parametersMap["l"] != "" {
		numLines, fileName := readNumberOfLinesFrom(parametersMap["l"])
		fmt.Printf("%d %s\n", numLines, fileName)
	}

}

func readNumberOfLinesFrom(linesCountFilepath string) (int64, string) {
	filesystemDir, fileName := parameters.GetFSandFilenameFromParameter(linesCountFilepath)
	wcLinesReader := readerLines.NewWcLinesReader(filesystemDir)

	numLines, err := wcLinesReader.Count(fileName)
	if err != nil {
		fmt.Printf("can't read %s correctly\n", fileName)
		os.Exit(1)
	}
	return numLines, fileName
}

func readNumberOfBytesFrom(bytesCountFilepath string) (int64, string) {
	filesystemDir, fileName := parameters.GetFSandFilenameFromParameter(bytesCountFilepath)
	wcBytesReader := readerBytes.NewWcBytesReader(filesystemDir)

	numBytes, err := wcBytesReader.Count(fileName)
	if err != nil {
		fmt.Printf("can't read %s correctly\n", fileName)
		os.Exit(1)
	}
	return numBytes, fileName
}
