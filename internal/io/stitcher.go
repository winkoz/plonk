package io

import (
	"fmt"
	"io/ioutil"
)

type stitcher struct{}

// Stitcher creates a new file by joining the contents of several text files.
type Stitcher interface {
	Stitch(sourcePath string, targetPath string, targetFilename string, filePaths []string, fileTransformator Transformator) error
}

// NewStitcher returns a fully initialised Stitcher
func NewStitcher() Stitcher {
	return stitcher{}
}

// Stitch checks existence of source and target paths; then stitches all source files together and saves it to target file name after applying the passed in transformation
func (s stitcher) Stitch(sourcePath string, targetPath string, targetFilename string, filePaths []string, fileTransformator Transformator) error {
	if err := s.validate(sourcePath, targetPath); err != nil {
		return err
	}

	mergedBytes, err := s.mergeFiles(sourcePath, filePaths)
	if err != nil {
		return err
	}

	transformedBytes := fileTransformator(mergedBytes)

	targetFilePath := fmt.Sprintf("%s/%s", targetPath, targetFilename)
	if err := ioutil.WriteFile(targetFilePath, transformedBytes, 0777); err != nil {
		return err
	}

	return nil
}

func (s stitcher) mergeFiles(sourcePath string, filePaths []string) ([]byte, error) {
	var err error = nil
	newLine := byte('\n')
	buffer := []byte{}
	for _, source := range filePaths {
		filePath := fmt.Sprintf("%s/%s", sourcePath, source)

		if !fileExists(filePath) {
			return nil, fmt.Errorf("File does not exist at path: %s", filePath)
		}

		var fileContents []byte
		fileContents, err = ioutil.ReadFile(filePath)
		if err != nil {
			return nil, err
		}

		buffer = append(buffer, fileContents...)
		buffer = append(buffer, newLine)
	}

	return buffer, err
}

func (s stitcher) validate(sourcePath string, targetPath string) error {
	// validate source path
	if err := isValidPath(sourcePath); err != nil {
		return err
	}

	// validate target path
	if err := isValidPath(targetPath); err != nil {
		return err
	}

	return nil
}
