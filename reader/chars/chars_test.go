package reader

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWcCharsReader(t *testing.T) {
	t.Run("Count reads 0 chars", func(t *testing.T) {
		dummyFile := []byte("")

		r := NewWcCharsReader()
		nChars := r.Count(dummyFile)

		expected := int64(0)
		assert.Equal(t, expected, nChars, "Got %d, wanted %d", nChars, expected)
	})

	t.Run("Count reads 1 char", func(t *testing.T) {
		dummyFile := []byte("a")

		r := NewWcCharsReader()
		nChars := r.Count(dummyFile)

		expected := int64(1)
		assert.Equal(t, expected, nChars, "Got %d, wanted %d", nChars, expected)
	})

	t.Run("Count reads multiple chars", func(t *testing.T) {
		dummyFile := []byte("abc")

		r := NewWcCharsReader()
		nChars := r.Count(dummyFile)

		expected := int64(3)
		assert.Equal(t, expected, nChars, "Got %d, wanted %d", nChars, expected)
	})

	t.Run("Count reads multiple chars included unicode ones", func(t *testing.T) {
		dummyFile := []byte("🚀")

		r := NewWcCharsReader()
		nChars := r.Count(dummyFile)

		expected := int64(1)
		assert.Equal(t, expected, nChars, "Got %d, wanted %d", nChars, expected)
	})
}
