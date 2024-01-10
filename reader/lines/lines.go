package reader

import (
	"strings"
)

type WcLinesReader struct {
}

func NewWcLinesReader() WcLinesReader {
	return WcLinesReader{}
}

func (w WcLinesReader) Count(content []byte) int64 {
	if len(content) == 0 {
		return int64(0)
	}

	lines := strings.Split(string(content), "\n")
	if lines[len(lines)-1] == "" {
		return int64(len(lines) - 1)
	}
	return int64(len(lines))
}
