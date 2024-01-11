package parameters

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParameters(t *testing.T) {
	t.Run("Parameters have been provided", func(t *testing.T) {
		os.Args = []string{"wc", "-l", "text.txt"}

		actual := HasProvided()

		assert.Truef(t, actual, "expected %t, got %t", true, actual)
	})

	t.Run("Parameters haven't been provided", func(t *testing.T) {
		os.Args = []string{"wc"}

		actual := HasProvided()

		assert.Falsef(t, actual, "expected %t, got %t", false, actual)
	})

	t.Run("Get filename from parameter provided", func(t *testing.T) {
		os.Args = []string{"wc", "-l", "text.txt"}

		actual := GetFilename()

		expected := "text.txt"
		assert.Equal(t, expected, actual, "expected %t, got %t", expected, actual)
	})

	t.Run("Get true if at least one flag has been passed", func(t *testing.T) {
		getBooleanPointer := func(b bool) *bool { return &b }
		flags := map[string]*bool{
			"c": getBooleanPointer(false),
			"d": getBooleanPointer(true),
			"e": getBooleanPointer(false),
		}

		actualName, actualBool := HaveBeenPassed(flags)

		expectedName := "d"
		assert.Equal(t, expectedName, actualName, "expected %t, got %t", expectedName, actualName)
		assert.Truef(t, actualBool, "expected %t, got %t", true, actualBool)
	})

	t.Run("Get false if no flags have been passed", func(t *testing.T) {
		getBooleanPointer := func(b bool) *bool { return &b }
		flags := map[string]*bool{
			"c": getBooleanPointer(false),
			"d": getBooleanPointer(false),
			"e": getBooleanPointer(false),
		}

		actualName, actualBool := HaveBeenPassed(flags)

		expectedName := ""
		assert.Equal(t, expectedName, actualName, "expected %t, got %t", expectedName, actualName)
		assert.Falsef(t, actualBool, "expected %t, got %t", true, actualBool)
	})

}
