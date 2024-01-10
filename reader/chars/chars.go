package reader

import (
	"unicode/utf8"
)

type WcCharsReader struct {
}

func NewWcCharsReader() WcCharsReader {
	return WcCharsReader{}
}

func (w WcCharsReader) Count(content []byte) int64 {
	return int64(utf8.RuneCount(content))
}
