package reader

import (
	"gowc/parameters"
	bytesReader "gowc/reader/bytes"
	charsReader "gowc/reader/chars"
	linesReader "gowc/reader/lines"
	wordsReader "gowc/reader/words"
)

type WcReaderManager interface {
	Count(content []byte) int64
}

func InitializeReaders() map[string]WcReaderManager {
	return map[string]WcReaderManager{
		parameters.BytesFlag: bytesReader.NewWcBytesReader(),
		parameters.LinesFlag: linesReader.NewWcLinesReader(),
		parameters.WordsFlag: wordsReader.NewWcWordsReader(),
		parameters.CharsFlag: charsReader.NewWcCharsReader(),
	}
}

func CountBytesWordsAndLines(readers map[string]WcReaderManager, input []byte) (int64, int64, int64) {
	return readers[parameters.BytesFlag].Count(input),
		readers[parameters.WordsFlag].Count(input),
		readers[parameters.LinesFlag].Count(input)
}

func CountWithSpecificReader(specificReader WcReaderManager, input []byte) int64 {
	return specificReader.Count(input)
}
