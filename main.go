package main

import (
	"fmt"
	"gowc/parameters"
	"gowc/reader"
	readerBytes "gowc/reader/bytes"
	readerChars "gowc/reader/chars"
	readerLines "gowc/reader/lines"
	readerWords "gowc/reader/words"
	"gowc/utils"
	"io/fs"
	"os"
)

func main() {
	params := parameters.NewParameters(map[string]string{
		"c": "Path to the file you want to count the bytes",
		"l": "Path to the file you want to count the lines",
		"w": "Path to the file you want to count the words",
		"m": "Path to the file you want to count the characters",
	})
	params.Parse()

	var filename string
	var filesystemDir fs.FS

	if countBytes, filePath := params.CountBytes(); countBytes || params.NoFlagsProvided() {
		filesystemDir, filename = utils.SplitFilepath(filePath)
		bytesReader := readerBytes.NewWcBytesReader(filesystemDir)
		count := count(bytesReader, filename)
		_, _ = printCountNumber(count)
	}

	if countLines, filePath := params.CountLines(); countLines || params.NoFlagsProvided() {
		filesystemDir, filename = utils.SplitFilepath(filePath)
		linesReader := readerLines.NewWcLinesReader(filesystemDir)
		count := count(linesReader, filename)
		_, _ = printCountNumber(count)
	}

	if countWords, filePath := params.CountWords(); countWords || params.NoFlagsProvided() {
		filesystemDir, filename = utils.SplitFilepath(filePath)
		wordsReader := readerWords.NewWcWordsReader(filesystemDir)
		count := count(wordsReader, filename)
		_, _ = printCountNumber(count)
	}

	if countChars, filePath := params.CountChars(); countChars {
		filesystemDir, filename = utils.SplitFilepath(filePath)
		charsReader := readerChars.NewWcCharsReader(filesystemDir)
		count := count(charsReader, filename)
		_, _ = printCountNumber(count)
	}

	_, _ = printFilename(filename)
}

func printFilename(filename string) (int, error) {
	return fmt.Printf("  %s\n", filename)
}

func printCountNumber(count int64) (int, error) {
	return fmt.Printf("  %d", count)
}

func count(reader reader.WcReaderManager, filename string) int64 {
	count, err := reader.Count(filename)
	if err != nil {
		fmt.Printf("can't read %s correctly\n", filename)
		os.Exit(1)
	}
	return count
}
