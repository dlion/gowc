package main

import (
	"fmt"
	"gowc/reader"
	"os"
)

func main() {
	filesystemDir := os.DirFS("resources")
	wcReader := reader.NewWcReader(filesystemDir)
	numBytes, err := wcReader.ReadNumberOfBytesFrom("test.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(numBytes)
}
