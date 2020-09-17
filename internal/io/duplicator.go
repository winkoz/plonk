package io

import (
	"fmt"
	"io/ioutil"

	"github.com/winkoz/plonk/internal/io/log"
)

type duplicator struct{}

// Duplicator duplicates files from one directory to another
type Duplicator interface {
	CopyMultiple(sourcePath string, targetPath string, sources []string, transformator Transformator) error
}

// NewDuplicator returns a fully initialised Duplicator
func NewDuplicator() Duplicator {
	return duplicator{}
}

func (d duplicator) copy(source string, target string, transformator Transformator) error {
	input, err := ioutil.ReadFile(source)
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
func (d duplicator) CopyMultiple(sourcePath string, targetPath string, sources []string, transformator Transformator) error {
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

	// copy sources
	var targetFilePath string
	var sourceFilePath string
	for _, s := range sources {
		targetFilePath = fmt.Sprintf("%s/%s", targetPath, s)
		sourceFilePath = fmt.Sprintf("%s/%s", sourcePath, s)
		if err := d.copy(sourceFilePath, targetFilePath, transformator); err != nil {
			log.Error(err)
			return err
		}
	}

	return nil
}
