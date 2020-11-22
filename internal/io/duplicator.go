package io

import (
	"io/ioutil"
	"path/filepath"

	"github.com/winkoz/plonk/internal/io/log"
)

type duplicator struct {
	service Service
}

// Duplicator duplicates files from one directory to another
type Duplicator interface {
	CopyMultiple(targetPath string, sourcesPaths []FileLocation, transformator Transformator) error
}

// NewDuplicator returns a fully initialised Duplicator
func NewDuplicator(service Service) Duplicator {
	return duplicator{
		service: service,
	}
}

func (d duplicator) copy(source string, target string, transformator Transformator) error {
	input, err := d.service.ReadFile(source)
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
func (d duplicator) CopyMultiple(targetPath string, sourcePaths []FileLocation, transformator Transformator) (err error) {
	signal := log.StartTrace("CopyMultiple")
	defer log.StopTrace(signal, err)

	// validate target path
	if err = d.service.IsValidPath(targetPath); err != nil {
		log.Error(err)
		return err
	}

	// copy sources
	var targetFilePath string
	for _, sourceFileLocation := range sourcePaths {
		targetFilePath = filepath.Join(targetPath, sourceFileLocation.OriginalFilePath)
		log.Debugf("Duplicating %s into %s", sourceFileLocation.OriginalFilePath, targetFilePath)
		if err = d.copy(sourceFileLocation.ResolvedFilePath, targetFilePath, transformator); err != nil {
			log.Error(err)
			return err
		}
	}

	return nil
}
