package reader

import (
	"io/fs"
)

type WcBytesReader struct {
	fs fs.FS
}

func NewWcBytesReader(fs fs.FS) WcBytesReader {
	return WcBytesReader{fs: fs}
}

func (w WcBytesReader) Count(filename string) (int64, error) {
	bytes, err := fs.ReadFile(w.fs, filename)
	if err != nil {
		return 0, err
	}

	return int64(len(bytes)), nil
}
