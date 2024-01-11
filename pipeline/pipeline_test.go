package pipeline

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestPipeline(t *testing.T) {
	t.Run("HasInput should return false if an input hasn't come from pipeline", func(t *testing.T) {
		actual := HasInput()

		assert.Falsef(t, actual, "expected %t, got %t", true, actual)
	})
	t.Run("HasInput should truw if an input comes from pipeline", func(t *testing.T) {
		r, w, _ := os.Pipe()
		_, _ = w.Write([]byte("Hello"))
		_ = w.Close()
		os.Stdin = r
		defer func(v *os.File) { os.Stdin = v }(os.Stdin)

		actual := HasInput()

		assert.Truef(t, actual, "expected %t, got %t", true, actual)
	})
}
