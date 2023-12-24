package parser

import (
	"flag"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParser(t *testing.T) {
	t.Run("Should parse flags correctly", func(t *testing.T) {
		fl1 := flag.NewFlagSet("first flag", flag.PanicOnError)
		fl1.String("a", "", "the first flag")
		os.Args = []string{"cmd", "-a=path/to/file.txt"}

		parser := NewParser(fl1)
		parametersMap, _ := parser.GetParameters()

		assert.NotEmpty(t, parametersMap, "the parameter map shouldn't be empty")
		assert.Contains(t, parametersMap, "a", "should contains the key 'a'")
		assert.Equal(t, "path/to/file.txt", parametersMap["a"], "the parametersMap should contains the file path")
	})
}
