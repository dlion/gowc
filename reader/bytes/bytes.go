package reader

import (
	"gowc/reader"
	"io/fs"
)

type WcBytesReader struct {
	reader fs.FS
}

func NewWcBytesReader(reader fs.FS) WcBytesReader {
	return WcBytesReader{reader: reader}
}

func (w WcBytesReader) Count(filename string) (int64, error) {
	bytes, err := reader.ReadFilename(filename, w.reader)
	if err != nil {
		return 0, err
	}

	return int64(len(bytes)), nil
}
