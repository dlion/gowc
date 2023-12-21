package reader

import "io/fs"

type WcReader struct {
	reader fs.FS
}

func NewWcReader(reader fs.FS) WcReader {
	return WcReader{reader: reader}
}

func (w WcReader) ReadNumberOfBytesFrom(filename string) (int64, error) {
	bytes, err := w.reader.(fs.ReadFileFS).ReadFile(filename)

	return int64(len(bytes)), err
}
