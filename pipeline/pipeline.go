package pipeline

import (
	"io"
	"log"
	"os"
)

func HasInput() bool {
	f, _ := os.Stdin.Stat()
	return (f.Mode() & os.ModeCharDevice) == 0
}

func ReadInput() []byte {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Error reading the pipeline: %v", err)
	}
	return input
}
