package main

import (
	"fmt"
	"gowc/parameters"
	"gowc/reader"
	readerBytes "gowc/reader/bytes"
	readerLines "gowc/reader/lines"
	readerWords "gowc/reader/words"
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

	if parametersMap["w"] != "" {
		numWords, fileName := readNumberOfWordsFrom(parametersMap["w"])
		fmt.Printf("%d %s\n", numWords, fileName)
	}

}

func readNumberOfLinesFrom(linesCountFilepath string) (int64, string) {
	filesystemDir, fileName := parameters.GetFSandFilenameFromParameter(linesCountFilepath)
	wcLinesReader := readerLines.NewWcLinesReader(filesystemDir)

	return count(wcLinesReader, fileName)
}

func readNumberOfBytesFrom(bytesCountFilepath string) (int64, string) {
	filesystemDir, fileName := parameters.GetFSandFilenameFromParameter(bytesCountFilepath)
	wcBytesReader := readerBytes.NewWcBytesReader(filesystemDir)

	return count(wcBytesReader, fileName)
}

func readNumberOfWordsFrom(wordsCountFilepath string) (int64, string) {
	filesystemDir, fileName := parameters.GetFSandFilenameFromParameter(wordsCountFilepath)
	wordsReader := readerWords.NewWcWordsReader(filesystemDir)

	return count(wordsReader, fileName)
}

func count(wcLinesReader reader.WcReaderManager, fileName string) (int64, string) {
	counted, err := wcLinesReader.Count(fileName)
	if err != nil {
		fmt.Printf("can't read %s correctly\n", fileName)
		os.Exit(1)
	}
	return counted, fileName
}
