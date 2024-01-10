package reader

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWcLinesReader(t *testing.T) {
	t.Run("Count returns 0 lines with an empty file", func(t *testing.T) {
		dummyFile := []byte("")

		r := NewWcLinesReader()
		nLines := r.Count(dummyFile)

		assert.Equal(t, int64(0), nLines, "Got %d, wanted %d", nLines, 0)
	})

	t.Run("Count returns 1 lines with just one line", func(t *testing.T) {
		dummyFile := []byte("Dummy String")

		r := NewWcLinesReader()
		nLines := r.Count(dummyFile)

		assert.Equal(t, int64(1), nLines, "Got %d, wanted %d", nLines, 1)
	})

	t.Run("Count returns 3 lines with a multi lines file content", func(t *testing.T) {
		dummyFile := []byte("Line 1\nLine 2\nLine 3")

		r := NewWcLinesReader()
		nLines := r.Count(dummyFile)

		assert.Equal(t, int64(3), nLines, "Got %d, wanted %d", nLines, 3)
	})

	t.Run("Count returns 3 lines with a multi lines file content with a trailing empty line", func(t *testing.T) {
		dummyFile := []byte("Line 1\nLine 2\nLine 3\n")

		r := NewWcLinesReader()
		nLines := r.Count(dummyFile)

		assert.Equal(t, int64(3), nLines, "Got %d, wanted %d", nLines, 3)
	})
}
