package parameters

import (
	"flag"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParameters(t *testing.T) {
	t.Run("Parse should parse flags correctly once called", func(t *testing.T) {
		os.Args = []string{"test", "-random", "./resource/text.txt"}
		parameters := NewParameters(map[string]string{
			"random":    "dummy",
			"param two": "dummy2",
		})

		parameters.Parse()

		assert.Equal(t, "dummy", flag.Lookup("random").Usage)
		assert.Equal(t, "dummy2", flag.Lookup("param two").Usage)
		assert.Equal(t, false, parameters.noFlagProvided, "Flags have been provided")
	})

	t.Run("Parse should parse args correctly once called", func(t *testing.T) {
		os.Args = []string{"test", "./resource/text.txt"}
		parameters := NewParameters(map[string]string{})

		parameters.Parse()

		assert.Equal(t, true, parameters.noFlagProvided, "Flags have been provided")
	})

	t.Run("Should get the count bytes parameter right", func(t *testing.T) {
		os.Args = []string{"test", "-c", "./resource/text.txt"}
		parameters := NewParameters(map[string]string{
			"c": "dummy",
		})
		parameters.Parse()

		countBytes, filePath := parameters.CountBytes()

		assert.Equal(t, true, countBytes, "CountBytes should be true")
		assert.Equal(t, "./resource/text.txt", filePath, "CountBytes should returns a filepath")
	})

	t.Run("Should get the count lines parameter right", func(t *testing.T) {
		os.Args = []string{"test", "-l", "./resource/text.txt"}
		parameters := NewParameters(map[string]string{
			"l": "dummy",
		})
		parameters.Parse()

		currentNumbertLines, filePath := parameters.CountLines()

		assert.Equal(t, true, currentNumbertLines, "CountLines should be true")
		assert.Equal(t, "./resource/text.txt", filePath, "CountLines should returns a filepath")
	})

	t.Run("Should get the count words parameter right", func(t *testing.T) {
		os.Args = []string{"test", "-w", "./resource/text.txt"}
		parameters := NewParameters(map[string]string{
			"w": "dummy",
		})
		parameters.Parse()

		currentNumberWords, filePath := parameters.CountWords()

		assert.Equal(t, true, currentNumberWords, "CountWords should be true")
		assert.Equal(t, "./resource/text.txt", filePath, "CountWords should returns a filepath")
	})

	t.Run("Should get the count characters parameter right", func(t *testing.T) {
		os.Args = []string{"test", "-m", "./resource/text.txt"}
		parameters := NewParameters(map[string]string{
			"m": "dummy",
		})
		parameters.Parse()

		currentNumberChars, filePath := parameters.CountChars()

		assert.Equal(t, true, currentNumberChars, "CountChars should be true")
		assert.Equal(t, "./resource/text.txt", filePath, "CountChars should returns a filepath")
	})

}
