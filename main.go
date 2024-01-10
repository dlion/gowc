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
	"os"
)

func main() {

	if inputFromPipeline() {
		fmt.Println("NOT YET IMPLEMENTED")
	} else {
		params := parameters.NewParameters(map[string]string{
			parameters.BytesFlag: "Path to the file you want to count the bytes",
			parameters.LinesFlag: "Path to the file you want to count the lines",
			parameters.WordsFlag: "Path to the file you want to count the words",
			parameters.CharsFlag: "Path to the file you want to count the characters",
		})
		params.Parse()

		var filename string
		var fileContent []byte
		var err error

		if countBytes, filePath := params.CountBytes(); countBytes || params.NoFlagsProvided() {
			filename = utils.SplitFilepath(filePath)
			fileContent, err = os.ReadFile(filename)
			if err != nil {
				cantReadCorrectly(filename)
			}

			bytesReader := readerBytes.NewWcBytesReader()
			count := count(bytesReader, fileContent)
			_, _ = printCountNumber(count)
		}

		if countLines, filePath := params.CountLines(); countLines || params.NoFlagsProvided() {
			filename = utils.SplitFilepath(filePath)
			fileContent, err = os.ReadFile(filename)
			if err != nil {
				cantReadCorrectly(filename)
			}

			linesReader := readerLines.NewWcLinesReader()
			count := count(linesReader, fileContent)
			_, _ = printCountNumber(count)
		}

		if countWords, filePath := params.CountWords(); countWords || params.NoFlagsProvided() {
			filename = utils.SplitFilepath(filePath)
			fileContent, err = os.ReadFile(filename)
			if err != nil {
				cantReadCorrectly(filename)
			}

			wordsReader := readerWords.NewWcWordsReader()
			count := count(wordsReader, fileContent)
			_, _ = printCountNumber(count)
		}

		if countChars, filePath := params.CountChars(); countChars {
			filename = utils.SplitFilepath(filePath)
			fileContent, err = os.ReadFile(filename)
			if err != nil {
				cantReadCorrectly(filename)
			}

			charsReader := readerChars.NewWcCharsReader()
			count := count(charsReader, fileContent)
			_, _ = printCountNumber(count)
		}

		_, _ = printFilename(filename)
	}

}

func inputFromPipeline() bool {
	f, _ := os.Stdin.Stat()
	return (f.Mode() & os.ModeCharDevice) == 0

}

func printFilename(filename string) (int, error) {
	return fmt.Printf("  %s\n", filename)
}

func printCountNumber(count int64) (int, error) {
	return fmt.Printf("  %d", count)
}

func count(reader reader.WcReaderManager, content []byte) int64 {
	count := reader.Count(content)
	return count
}

func cantReadCorrectly(filename string) {
	fmt.Printf("can't read %s correctly\n", filename)
	os.Exit(1)
}
