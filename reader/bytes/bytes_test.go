package reader

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWcBytesReader(t *testing.T) {
	t.Run("Count returns the exact number of bytes", func(t *testing.T) {
		dummyContent := make([]byte, 100)
		r := NewWcBytesReader()

		currentBytes := r.Count(dummyContent)

		expected := int64(100)
		assert.Equal(t, expected, currentBytes, "Got %d, wanted %d", currentBytes, expected)
	})
}
