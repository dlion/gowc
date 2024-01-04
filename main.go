package main

import (
	"fmt"
	"gowc/parameters"
	"gowc/reader"
	readerBytes "gowc/reader/bytes"
	readerChars "gowc/reader/chars"
	readerLines "gowc/reader/lines"
	readerWords "gowc/reader/words"
	"os"
)

func main() {
	parametersMap := parameters.GetParameters()

	if parametersMap["c"] != "" {
		numBytes, filename := readNumberOfBytesFrom(parametersMap["c"])
		printNumbers(numBytes, filename)
	}

	if parametersMap["l"] != "" {
		numLines, fileName := readNumberOfLinesFrom(parametersMap["l"])
		printNumbers(numLines, fileName)
	}

	if parametersMap["w"] != "" {
		numWords, fileName := readNumberOfWordsFrom(parametersMap["w"])
		printNumbers(numWords, fileName)
	}

	if parametersMap["m"] != "" {
		numChars, fileName := readNumberOfCharsFrom(parametersMap["m"])
		printNumbers(numChars, fileName)
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
	wcWordsReader := readerWords.NewWcWordsReader(filesystemDir)

	return count(wcWordsReader, fileName)
}

func readNumberOfCharsFrom(charsCountFilepath string) (int64, string) {
	filesystemDir, fileName := parameters.GetFSandFilenameFromParameter(charsCountFilepath)
	wcCharsReader := readerChars.NewWcCharsReader(filesystemDir)

	return count(wcCharsReader, fileName)
}

func count(wcLinesReader reader.WcReaderManager, fileName string) (int64, string) {
	counted, err := wcLinesReader.Count(fileName)
	if err != nil {
		fmt.Printf("can't read %s correctly\n", fileName)
		os.Exit(1)
	}
	return counted, fileName
}

func printNumbers(count int64, filename string) {
	fmt.Printf("%d %s\n", count, filename)
}
