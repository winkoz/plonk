package io

import (
	"fmt"
	"io/ioutil"

	"github.com/winkoz/plonk/internal/io/log"
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
		log.Error(err)
		return err
	}

	mergedBytes, err := s.mergeFiles(sourcePath, filePaths)
	if err != nil {
		log.Error(err)
		return err
	}

	transformedBytes := fileTransformator(mergedBytes)

	targetFilePath := fmt.Sprintf("%s/%s", targetPath, targetFilename)
	if err := ioutil.WriteFile(targetFilePath, transformedBytes, OwnerPermission); err != nil {
		log.Error(err)
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

		if !FileExists(filePath) {
			err := fmt.Errorf("File does not exist at path: %s", filePath)
			log.Error(err)
			return nil, err
		}

		var fileContents []byte
		fileContents, err = ioutil.ReadFile(filePath)
		if err != nil {
			log.Error(err)
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
		log.Error(err)
		return err
	}

	// validate target path
	if err := isValidPath(targetPath); err != nil {
		log.Error(err)
		return err
	}

	return nil
}
