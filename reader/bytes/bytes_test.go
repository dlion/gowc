package reader

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/fs"
	"testing"
)

func TestWcBytesReader(t *testing.T) {
	t.Run("ReadFile should be called once and returns the exact number of bytes", func(t *testing.T) {
		fakeReader := new(FakeFsReader)
		fakeReader.On("ReadFile", mock.AnythingOfType("string")).Return(make([]byte, 100), nil)

		r := NewWcBytesReader(fakeReader)

		nBytes, _ := r.Count("dummy_file.txt")

		fakeReader.AssertNumberOfCalls(t, "ReadFile", 1)
		assert.Equal(t, int64(100), nBytes, "Got %d, wanted %d", nBytes, 100)
	})
}

type FakeFsReader struct {
	mock.Mock
}

func (w *FakeFsReader) Open(filename string) (fs.File, error) {
	args := w.Called(filename)
	return args.Get(0).(fs.File), args.Error(1)
}

func (w *FakeFsReader) ReadFile(filename string) ([]byte, error) {
	args := w.Called(filename)
	return args.Get(0).([]byte), args.Error(1)
}
