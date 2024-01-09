package main

import (
	"flag"
	"fmt"
	"gowc/reader"
	readerBytes "gowc/reader/bytes"
	readerChars "gowc/reader/chars"
	readerLines "gowc/reader/lines"
	readerWords "gowc/reader/words"
	"gowc/utils"
	"io/fs"
	"os"
)

type FSReader struct {
	reader   reader.WcReaderManager
	filename string
}

func main() {
	paramsDefMap := setParameters()
	paramsMap := getParametersFromCLI(paramsDefMap)
	readers := getFsReadersByParameters(paramsMap)
	counts, filename := countWithReaders(readers)
	printCountsAndFilename(counts, filename)
}

func countWithReaders(readers []FSReader) ([]int64, string) {
	counted := make([]int64, len(readers))
	for i, r := range readers {
		n := count(r.reader, r.filename)
		counted[i] = n
	}

	return counted, readers[0].filename
}

func getParametersFromCLI(paramsDefMap map[string]string) map[string]string {
	paramsMap := map[string]string{}

	if len(flag.Args()) == 0 {
		for p, _ := range paramsDefMap {
			filenamePath := flag.Lookup(p).Value.String()
			if filenamePath != "" {
				paramsMap[p] = filenamePath
			}
		}
	} else {
		paramsMap["c"] = flag.Args()[0]
		paramsMap["l"] = flag.Args()[0]
		paramsMap["w"] = flag.Args()[0]
	}

	return paramsMap
}

func setParameters() map[string]string {
	paramsDefMap := map[string]string{
		"c": "Path to the file you want to count the bytes",
		"l": "Path to the file you want to count the lines",
		"w": "Path to the file you want to count the words",
		"m": "Path to the file you want to count the characters",
	}

	for p, v := range paramsDefMap {
		flag.String(p, "", v)
	}

	flag.Parse()

	return paramsDefMap
}

func getFsReadersByParameters(params map[string]string) []FSReader {
	var output []FSReader

	for paramName, filepath := range params {
		filesystemDir, fileName := splitFilepath(filepath)

		maprandom := map[string]FSReader{
			"c": {reader: readerBytes.NewWcBytesReader(filesystemDir), filename: fileName},
			"l": {reader: readerLines.NewWcLinesReader(filesystemDir), filename: fileName},
			"w": {reader: readerWords.NewWcWordsReader(filesystemDir), filename: fileName},
			"m": {reader: readerChars.NewWcCharsReader(filesystemDir), filename: fileName},
		}

		output = append(output, maprandom[paramName])
	}

	return output
}

func splitFilepath(filepath string) (fs.FS, string) {
	dirPath, fileName := utils.SplitPath(filepath)
	filesystemDir := os.DirFS(dirPath)
	return filesystemDir, fileName
}

func count(fsReader reader.WcReaderManager, fileName string) int64 {
	counted, err := fsReader.Count(fileName)
	if err != nil {
		fmt.Printf("can't read %s correctly\n", fileName)
		os.Exit(1)
	}
	return counted
}

func printCountsAndFilename(counts []int64, filename string) {
	for _, count := range counts {
		fmt.Printf("%d   ", count)
	}
	fmt.Printf("%s\n", filename)
}
