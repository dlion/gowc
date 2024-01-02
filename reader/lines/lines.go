package reader

import (
	"io/fs"
	"strings"
)

type WcLinesReader struct {
	fs fs.FS
}

func NewWcLinesReader(fs fs.FS) WcLinesReader {
	return WcLinesReader{fs: fs}
}

func (w WcLinesReader) Count(filename string) (int64, error) {
	bytes, err := fs.ReadFile(w.fs, filename)
	if err != nil || len(bytes) == 0 {
		return 0, err
	}

	lines := strings.Split(string(bytes), "\n")
	if lines[len(lines)-1] == "" {
		return int64(len(lines) - 1), nil
	}
	return int64(len(lines)), nil
}
