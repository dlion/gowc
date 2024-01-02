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
	if err != nil {
		return 0, err
	}

	lines := strings.Split(string(bytes), "\n")
	linesNumber := len(lines) - 1
	return int64(linesNumber), nil
}
