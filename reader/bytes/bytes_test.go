package reader

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWcBytesReader(t *testing.T) {
	t.Run("Count returns the exact number of bytes", func(t *testing.T) {
		dummy_content := make([]byte, 100)

		r := NewWcBytesReader()

		nBytes, _ := r.Count(string(dummy_content))

		assert.Equal(t, int64(100), nBytes, "Got %d, wanted %d", nBytes, 100)
	})
}
