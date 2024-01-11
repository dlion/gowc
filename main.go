package main

import (
	"fmt"
	"gowc/parameters"
	"gowc/pipeline"
	"gowc/reader"
	"log"
	"os"
)

func main() {
	var input []byte
	var filename string

	if input, filename = getInput(); len(input) == 0 {
		fmt.Println("usage: ./gowc <flag> <filename>")
		os.Exit(1)
	}

	flags := parameters.GetFlags()

	initializedReaders := reader.InitializeReaders()

	if flagNamePassed, flagHasBeenPassed := parameters.HaveBeenPassed(flags); flagHasBeenPassed {
		count := reader.CountWithSpecificReader(initializedReaders[flagNamePassed], input)
		fmt.Printf("    %d %s\n", count, filename)
	} else {
		bytes, words, lines := reader.CountBytesWordsAndLines(initializedReaders, input)
		fmt.Printf("   %d %d %d %s\n", bytes, words, lines, filename)
	}
}

func getInput() ([]byte, string) {
	const EmptyString = ""
	if pipeline.HasInput() {
		return pipeline.ReadInput(), EmptyString
	} else {
		if parameters.HasProvided() {
			filename := parameters.GetFilename()
			return readFile(filename), filename
		} else {
			return make([]byte, 0), EmptyString
		}
	}
}

func readFile(filename string) []byte {
	input, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading the file: %v", err)
	}
	return input
}
