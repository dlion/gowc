package reader

type WcReaderManager interface {
	Count(filename string) (int64, error)
}
