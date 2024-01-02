package reader

import (
	"io/fs"
	"strings"
)

type WcWordsReader struct {
	fs fs.FS
}

func NewWcWordsReader(fs fs.FS) WcWordsReader {
	return WcWordsReader{fs: fs}
}

func (w WcWordsReader) Count(filename string) (int64, error) {
	bytes, err := fs.ReadFile(w.fs, filename)
	if err != nil || len(bytes) == 0 {
		return 0, err
	}
	words := strings.Fields(string(bytes))

	return int64(len(words)), nil
}
