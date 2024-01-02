package reader

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"testing/fstest"
)

func TestWcLinesReader(t *testing.T) {
	t.Run("Count returns 0 lines with an empty file", func(t *testing.T) {
		fs := fstest.MapFS{
			"dummy_file.txt":  {Data: []byte("")},
			"hello-world2.md": {Data: []byte("hola")},
		}

		r := NewWcLinesReader(fs)
		nLines, _ := r.Count("dummy_file.txt")

		assert.Equal(t, int64(0), nLines, "Got %d, wanted %d", nLines, 0)
	})

	t.Run("Count returns 1 lines with just one line", func(t *testing.T) {
		fs := fstest.MapFS{
			"dummy_file.txt":  {Data: []byte("Dummy String")},
			"hello-world2.md": {Data: []byte("hola")},
		}

		r := NewWcLinesReader(fs)
		nLines, _ := r.Count("dummy_file.txt")

		assert.Equal(t, int64(1), nLines, "Got %d, wanted %d", nLines, 1)
	})

	t.Run("Count returns 3 lines with a multi lines file content", func(t *testing.T) {
		fs := fstest.MapFS{
			"dummy_file.txt":  {Data: []byte("Line 1\nLine 2\nLine 3")},
			"hello-world2.md": {Data: []byte("hola")},
		}

		r := NewWcLinesReader(fs)
		nLines, _ := r.Count("dummy_file.txt")

		assert.Equal(t, int64(3), nLines, "Got %d, wanted %d", nLines, 3)
	})

	t.Run("Count returns 3 lines with a multi lines file content with a trailing empty line", func(t *testing.T) {
		fs := fstest.MapFS{
			"dummy_file.txt":  {Data: []byte("Line 1\nLine 2\nLine 3\n")},
			"hello-world2.md": {Data: []byte("hola")},
		}

		r := NewWcLinesReader(fs)
		nLines, _ := r.Count("dummy_file.txt")

		assert.Equal(t, int64(3), nLines, "Got %d, wanted %d", nLines, 3)
	})
}
