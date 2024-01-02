package reader

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/fs"
	"testing"
)

func TestWcLinesReader(t *testing.T) {
	t.Run("ReadFile should be called once and returns 2 lines", func(t *testing.T) {
		fakeReader := new(FakeFSReader)
		fakeReader.On("ReadFile", mock.AnythingOfType("string")).Return([]byte("random\nstring"), nil)

		r := NewWcLinesReader(fakeReader)
		nLines, _ := r.Count("dummy_file.txt")

		fakeReader.AssertNumberOfCalls(t, "ReadFile", 1)
		assert.Equal(t, int64(1), nLines, "Got %d, wanted %d", nLines, 1)
	})
}

type FakeFSReader struct {
	mock.Mock
}

func (w *FakeFSReader) Open(filename string) (fs.File, error) {
	args := w.Called(filename)
	return args.Get(0).(fs.File), args.Error(1)
}

func (w *FakeFSReader) ReadFile(filename string) ([]byte, error) {
	args := w.Called(filename)
	return args.Get(0).([]byte), args.Error(1)
}
