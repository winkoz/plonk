package sharedtesting

import (
	"os"
	"time"

	"github.com/stretchr/testify/mock"
)

// FileInfoMock …
type FileInfoMock struct {
	mock.Mock
}

// IsDir …
func (fim *FileInfoMock) IsDir() bool {
	args := fim.Called()
	return args.Bool(0)
}

// Name …
func (fim *FileInfoMock) Name() string {
	return "name"
}

// Size …
func (fim *FileInfoMock) Size() int64 {
	return 0
}

// Mode …
func (fim *FileInfoMock) Mode() os.FileMode {
	return 000
}

// ModTime …
func (fim *FileInfoMock) ModTime() time.Time {
	return time.Now()
}

// Sys …
func (fim *FileInfoMock) Sys() interface{} {
	return nil
}
