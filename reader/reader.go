package reader

type WcReaderManager interface {
	Count(content []byte) int64
}
