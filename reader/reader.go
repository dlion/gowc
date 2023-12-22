package reader

import (
	"io/fs"
	"strings"
)

type WcReader struct {
	reader fs.FS
}

func NewWcReader(reader fs.FS) WcReader {
	return WcReader{reader: reader}
}

func (w WcReader) ReadNumberOfBytesFrom(filename string) (int64, error) {
	bytes, err := readFilename(filename, w)
	if err != nil {
		return 0, err
	}

	return int64(len(bytes)), nil
}

func (w WcReader) ReadNumberOfLinesFrom(filename string) (int64, error) {
	bytes, err := readFilename(filename, w)
	if err != nil {
		return 0, err
	}

	lines := strings.Split(string(bytes), "\n")
	linesNumber := len(lines) - 1
	return int64(linesNumber), nil
}

func readFilename(filename string, w WcReader) ([]byte, error) {
	bytes, err := w.reader.(fs.ReadFileFS).ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}
	return bytes, nil
}
