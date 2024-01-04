package reader

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"testing/fstest"
)

func TestWcCharsReader(t *testing.T) {
	t.Run("Count reads 0 chars", func(t *testing.T) {
		fs := fstest.MapFS{
			"dummy_file.txt": {Data: []byte("")},
		}

		r := NewWcCharsReader(fs)
		nChars, _ := r.Count("dummy_file.txt")

		expected := int64(0)
		assert.Equal(t, expected, nChars, "Got %d, wanted %d", nChars, expected)
	})

	t.Run("Count reads 1 char", func(t *testing.T) {
		fs := fstest.MapFS{
			"dummy_file.txt": {Data: []byte("a")},
		}

		r := NewWcCharsReader(fs)
		nChars, _ := r.Count("dummy_file.txt")

		expected := int64(1)
		assert.Equal(t, expected, nChars, "Got %d, wanted %d", nChars, expected)
	})

	t.Run("Count reads multiple chars", func(t *testing.T) {
		fs := fstest.MapFS{
			"dummy_file.txt": {Data: []byte("abc")},
		}

		r := NewWcCharsReader(fs)
		nChars, _ := r.Count("dummy_file.txt")

		expected := int64(3)
		assert.Equal(t, expected, nChars, "Got %d, wanted %d", nChars, expected)
	})

	t.Run("Count reads multiple chars included unicode ones", func(t *testing.T) {
		fs := fstest.MapFS{
			"dummy_file.txt": {Data: []byte("🚀")},
		}

		r := NewWcCharsReader(fs)
		nChars, _ := r.Count("dummy_file.txt")

		expected := int64(1)
		assert.Equal(t, expected, nChars, "Got %d, wanted %d", nChars, expected)
	})
}
