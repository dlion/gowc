package reader

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/fs"
	"testing"
)

func TestWcReader(t *testing.T) {
	t.Run("ReadFile should be called once and returns the exact number of bytes", func(t *testing.T) {
		fakeReader := new(FakeWcReader)
		fakeReader.On("ReadFile", mock.AnythingOfType("string")).Return(make([]byte, 100), nil)

		r := NewWcReader(fakeReader)

		nBytes, _ := r.ReadNumberOfBytesFrom("dummy_file.txt")

		fakeReader.AssertNumberOfCalls(t, "ReadFile", 1)
		assert.Equal(t, int64(100), nBytes, "Got %d, wanted %d", nBytes, 100)
	})

	t.Run("ReadFile should be called once and returns 2 lines", func(t *testing.T) {
		fakeReader := new(FakeWcReader)
		fakeReader.On("ReadFile", mock.AnythingOfType("string")).Return([]byte("random\nstring"), nil)

		r := NewWcReader(fakeReader)
		nLines, _ := r.ReadNumberOfLinesFrom("dummy_file.txt")

		fakeReader.AssertNumberOfCalls(t, "ReadFile", 1)
		assert.Equal(t, int64(1), nLines, "Got %d, wanted %d", nLines, 1)
	})
}

type FakeWcReader struct {
	mock.Mock
}

func (w *FakeWcReader) Open(filename string) (fs.File, error) {
	args := w.Called(filename)
	return args.Get(0).(fs.File), args.Error(1)
}

func (w *FakeWcReader) ReadFile(filename string) ([]byte, error) {
	args := w.Called(filename)
	return args.Get(0).([]byte), args.Error(1)
}
