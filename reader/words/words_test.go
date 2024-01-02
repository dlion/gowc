package reader

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"testing/fstest"
)

func TestWcWordsReader(t *testing.T) {
	t.Run("Count returns 0 if the file doesn't have words", func(t *testing.T) {
		fs := fstest.MapFS{
			"dummy_file.txt":  {Data: []byte("")},
			"hello-world2.md": {Data: []byte("hola")},
		}

		r := NewWcWordsReader(fs)
		nWords, _ := r.Count("dummy_file.txt")

		expected := int64(0)
		assert.Equal(t, expected, nWords, "Got %d, wanted %d", nWords, expected)
	})

	t.Run("Count returns 1 if the file have just 1 word", func(t *testing.T) {
		fs := fstest.MapFS{
			"dummy_file.txt":  {Data: []byte("Dummy")},
			"hello-world2.md": {Data: []byte("hola")},
		}

		r := NewWcWordsReader(fs)
		nWords, _ := r.Count("dummy_file.txt")

		expected := int64(1)
		assert.Equal(t, expected, nWords, "Got %d, wanted %d", nWords, expected)
	})

	t.Run("Count returns 3 if the file have 3 words", func(t *testing.T) {
		fs := fstest.MapFS{
			"dummy_file.txt":  {Data: []byte("Dummy Word Here")},
			"hello-world2.md": {Data: []byte("hola")},
		}

		r := NewWcWordsReader(fs)
		nWords, _ := r.Count("dummy_file.txt")

		expected := int64(3)
		assert.Equal(t, expected, nWords, "Got %d, wanted %d", nWords, expected)
	})
}
