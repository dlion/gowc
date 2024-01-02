package reader

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"testing/fstest"
)

func TestWcBytesReader(t *testing.T) {
	t.Run("Count returns the exact number of bytes", func(t *testing.T) {
		fs := fstest.MapFS{
			"dummy_file.txt":  {Data: make([]byte, 100)},
			"hello-world2.md": {Data: []byte("hola")},
		}

		r := NewWcBytesReader(fs)

		nBytes, _ := r.Count("dummy_file.txt")

		assert.Equal(t, int64(100), nBytes, "Got %d, wanted %d", nBytes, 100)
	})
}
