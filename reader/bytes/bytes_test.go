package reader

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWcBytesReader(t *testing.T) {
	t.Run("Count reads 0 bytes", func(t *testing.T) {
		dummyContent := make([]byte, 0)
		r := NewWcBytesReader()

		currentBytes := r.Count(dummyContent)

		expected := int64(0)
		assert.Equal(t, expected, currentBytes, "Got %d, wanted %d", currentBytes, expected)
	})

	t.Run("Count reads 1 byte", func(t *testing.T) {
		dummyContent := make([]byte, 1)
		r := NewWcBytesReader()

		currentBytes := r.Count(dummyContent)

		expected := int64(1)
		assert.Equal(t, expected, currentBytes, "Got %d, wanted %d", currentBytes, expected)
	})

	t.Run("Count reads multiple bytes", func(t *testing.T) {
		dummyContent := make([]byte, 100)
		r := NewWcBytesReader()

		currentBytes := r.Count(dummyContent)

		expected := int64(100)
		assert.Equal(t, expected, currentBytes, "Got %d, wanted %d", currentBytes, expected)
	})
}
