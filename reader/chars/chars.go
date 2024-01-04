package reader

import (
	"io/fs"
	"unicode/utf8"
)

type WcCharsReader struct {
	fs fs.FS
}

func NewWcCharsReader(fs fs.FS) WcCharsReader {
	return WcCharsReader{fs: fs}
}

func (w WcCharsReader) Count(filename string) (int64, error) {
	bytes, err := fs.ReadFile(w.fs, filename)
	if err != nil || len(bytes) == 0 {
		return 0, err
	}

	return int64(utf8.RuneCount(bytes)), nil
}
