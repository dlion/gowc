package reader

import (
	"strings"
)

type WcWordsReader struct{}

func NewWcWordsReader() WcWordsReader {
	return WcWordsReader{}
}

func (w WcWordsReader) Count(content []byte) int64 {
	words := strings.Fields(string(content))

	return int64(len(words))
}
