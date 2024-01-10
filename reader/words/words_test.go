package reader

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWcWordsReader(t *testing.T) {
	t.Run("Count returns 0 if the file doesn't have words", func(t *testing.T) {
		dummyFile := []byte("")

		r := NewWcWordsReader()
		nWords := r.Count(dummyFile)

		expected := int64(0)
		assert.Equal(t, expected, nWords, "Got %d, wanted %d", nWords, expected)
	})

	t.Run("Count returns 1 if the file have just 1 word", func(t *testing.T) {
		dummyFile := []byte("Dummy")

		r := NewWcWordsReader()
		nWords := r.Count(dummyFile)

		expected := int64(1)
		assert.Equal(t, expected, nWords, "Got %d, wanted %d", nWords, expected)
	})

	t.Run("Count returns 3 if the file have 3 words", func(t *testing.T) {
		dummyFile := []byte("Dummy Word Here")

		r := NewWcWordsReader()
		nWords := r.Count(dummyFile)

		expected := int64(3)
		assert.Equal(t, expected, nWords, "Got %d, wanted %d", nWords, expected)
	})
}
