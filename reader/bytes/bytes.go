package reader

type WcBytesReader struct {
}

func NewWcBytesReader() WcBytesReader {
	return WcBytesReader{}
}

func (w WcBytesReader) Count(content []byte) int64 {
	return int64(len(content))
}
