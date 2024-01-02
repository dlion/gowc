package reader

import (
	"gowc/reader"
	"io/fs"
	"strings"
)

type WcLinesReader struct {
	reader fs.FS
}

func NewWcLinesReader(reader fs.FS) WcLinesReader {
	return WcLinesReader{reader: reader}
}

func (w WcLinesReader) Count(filename string) (int64, error) {
	bytes, err := reader.ReadFilename(filename, w.reader)
	if err != nil || len(bytes) == 0 {
		return 0, err
	}

	lines := strings.Split(string(bytes), "\n")
	if lines[len(lines)-1] == "" {
		return int64(len(lines) - 1), nil
	}
	return int64(len(lines)), nil
}
