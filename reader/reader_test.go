package reader

import (
	"github.com/stretchr/testify/mock"
	"io/fs"
	"testing"
)

func TestWcReader(t *testing.T) {
	t.Run("ReadFile should be called once", func(t *testing.T) {
		fakeReader := new(FakeWcReader)
		fakeReader.On("ReadFile", mock.AnythingOfType("string")).Return(make([]byte, 100), nil)

		r := NewWcReader(fakeReader)

		_, _ = r.ReadNumberOfBytesFrom("dummy_file.txt")

		fakeReader.AssertNumberOfCalls(t, "ReadFile", 1)
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
