package io

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/winkoz/plonk/internal/io/log"
)

type duplicator struct{}

// Duplicator duplicates files from one directory to another
type Duplicator interface {
	CopyMultiple(targetPath string, sourcesPaths []string, transformator Transformator) error
}

// NewDuplicator returns a fully initialised Duplicator
func NewDuplicator() Duplicator {
	return duplicator{}
}

func (d duplicator) copy(source string, target string, transformator Transformator) error {
	input, err := ReadFile(source)
	if err != nil {
		log.Error(err)
		return err
	}

	transformedBytes := input
	if transformator != nil {
		transformedBytes = transformator(input)
	}

	err = ioutil.WriteFile(target, transformedBytes, OwnerPermission)
	if err != nil {
		log.Errorf("Error creating: %+v\n%+v", target, err)
		return err
	}
	return err
}

// CopyMultiple copies a series of files from a specific path into another
func (d duplicator) CopyMultiple(targetPath string, sourcePaths []string, transformator Transformator) error {
	// validate target path
	if err := isValidPath(targetPath); err != nil {
		log.Error(err)
		return err
	}

	// copy sources
	var targetFilePath string
	for _, sourceFilePath := range sourcePaths {
		targetFilePath = fmt.Sprintf("%s/%s", targetPath, filepath.Base(sourceFilePath))
		log.Debugf("Duplicating %s into %s", sourceFilePath, targetFilePath)
		if err := d.copy(sourceFilePath, targetFilePath, transformator); err != nil {
			log.Error(err)
			return err
		}
	}

	return nil
}
