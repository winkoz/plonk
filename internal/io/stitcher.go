package io

import (
	"io/ioutil"
	"path/filepath"

	"github.com/winkoz/plonk/internal/io/log"
)

type stitcher struct {
	service Service
}

// Stitcher creates a new file by joining the contents of several text files.
type Stitcher interface {
	Stitch(sourcePath string, targetPath string, targetFilename string, filePaths []string, fileTransformator Transformator) error
}

// NewStitcher returns a fully initialised Stitcher
func NewStitcher() Stitcher {
	return stitcher{
		service: NewService(),
	}
}

// Stitch checks existence of source and target paths; then stitches all source files together and saves it to target file name after applying the passed in transformation
func (s stitcher) Stitch(sourcePath string, targetPath string, targetFilename string, filePaths []string, fileTransformator Transformator) (err error) {
	signal := log.StartTrace("Stich")
	defer log.StopTrace(signal, err)

	if err = s.validate(sourcePath, targetPath); err != nil {
		log.Error(err)
		return err
	}

	mergedBytes, err := s.mergeFiles(sourcePath, filePaths)
	if err != nil {
		log.Error(err)
		return err
	}

	transformedBytes := fileTransformator(mergedBytes)

	targetFilePath := filepath.Join(targetPath, targetFilename)
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
		filePath := filepath.Join(sourcePath, source)

		var fileContents []byte
		fileContents, err = s.service.ReadFile(filePath)
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
	if err := s.service.IsValidPath(sourcePath); err != nil {
		log.Error(err)
		return err
	}

	// validate target path
	if err := s.service.IsValidPath(targetPath); err != nil {
		log.Error(err)
		return err
	}

	return nil
}
