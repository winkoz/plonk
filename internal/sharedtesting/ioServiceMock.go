package sharedtesting

import (
	"github.com/stretchr/testify/mock"
	"github.com/winkoz/plonk/internal/io"
)

// IOServiceMock is a mock of IO Service
type IOServiceMock struct {
	mock.Mock
}

// GetCurrentDir mocks GetCurrentDir
func (i *IOServiceMock) GetCurrentDir() string {
	args := i.Called()
	return args.String(0)
}

// DirectoryExists mocks DirectoryExists
func (i *IOServiceMock) DirectoryExists(filename string) bool {
	args := i.Called(filename)
	return args.Bool(0)
}

// FileExists mocks FileExists
func (i *IOServiceMock) FileExists(filename string) bool {
	args := i.Called(filename)
	return args.Bool(0)
}

// CreatePath mocks CreatePath
func (i *IOServiceMock) CreatePath(path string) error {
	args := i.Called(path)
	return args.Error(0)
}

// DeletePath mocks DeletePath
func (i *IOServiceMock) DeletePath(path string) {
	i.Called(path)
}

// ReadFile mocks ReadFile
func (i *IOServiceMock) ReadFile(path string) ([]byte, error) {
	args := i.Called(path)
	return args.Get(0).([]byte), args.Error(1)
}

// Walk mocks Walk
func (i *IOServiceMock) Walk(root string, walkFn io.WalkFunc) error {
	args := i.Called(root)
	mockFileInfo := new(FileInfoMock)
	mockFileInfo.On(
		"IsDir",
	).Return(
		false,
	)
	walkFn(root, mockFileInfo, nil)
	return args.Error(0)
}

// Append mocks Append
func (i *IOServiceMock) Append(targetFilePath string, content string) error {
	args := i.Called(targetFilePath, content)
	return args.Error(0)
}

// Write mocks Write
func (i *IOServiceMock) Write(targetFilePath string, content string) error {
	args := i.Called(targetFilePath, content)
	return args.Error(0)
}

// IsValidPath mocks IsValidPath
func (i *IOServiceMock) IsValidPath(path string) error {
	args := i.Called(path)
	return args.Error(0)
}
