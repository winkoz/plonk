package sharedtesting

import (
	"github.com/stretchr/testify/mock"
)

// MockTransformator …
type MockTransformator func(input []byte) []byte

// StitcherMock …
type StitcherMock struct {
	mock.Mock
}

// Stitch …
func (s *StitcherMock) Stitch(sourcePath string, targetPath string, targetFilename string, filePaths []string, fileTransformator MockTransformator) error {
	args := s.Called(sourcePath, targetPath, targetFilename, filePaths)
	return args.Error(0)
}

/*
type wantStitcher struct {
		testStitcher      bool
		sourcePath        string
		targetPath        string
		targetFilename    string
		filePaths         []string
		fileTransformator io.Transformator
		returnErr         error
	}
*/
