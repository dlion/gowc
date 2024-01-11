package reader

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWcLinesReader(t *testing.T) {
	t.Run("Count returns 0 lines with an empty file", func(t *testing.T) {
		dummyFile := []byte("")
		r := NewWcLinesReader()

		currentLines := r.Count(dummyFile)

		expected := int64(0)
		assert.Equal(t, expected, currentLines, "Got %d, wanted %d", currentLines, expected)
	})

	t.Run("Count returns 1 lines with just one line", func(t *testing.T) {
		dummyFile := []byte("Dummy String")
		r := NewWcLinesReader()

		currentLines := r.Count(dummyFile)

		expected := int64(1)
		assert.Equal(t, expected, currentLines, "Got %d, wanted %d", currentLines, expected)
	})

	t.Run("Count returns 3 lines with a multi lines file content", func(t *testing.T) {
		dummyFile := []byte("Line 1\nLine 2\nLine 3")
		r := NewWcLinesReader()

		currentLines := r.Count(dummyFile)

		expected := int64(3)
		assert.Equal(t, expected, currentLines, "Got %d, wanted %d", currentLines, expected)
	})

	t.Run("Count returns 3 lines with a multi lines file content with a trailing empty line", func(t *testing.T) {
		dummyFile := []byte("Line 1\nLine 2\nLine 3\n")
		r := NewWcLinesReader()

		currentLines := r.Count(dummyFile)

		expected := int64(3)
		assert.Equal(t, expected, currentLines, "Got %d, wanted %d", currentLines, expected)
	})
}
