package reader

import (
	"io/fs"
)

type WcReaderManager interface {
	Count(filename string) (int64, error)
}

func ReadFilename(filename string, r fs.FS) ([]byte, error) {
	bytes, err := r.(fs.ReadFileFS).ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}
	return bytes, nil
}
