package io

import (
	"fmt"

	"github.com/winkoz/plonk/internal/io/log"
	"gopkg.in/yaml.v2"
)

// YamlReader reads yaml files and loads them into a passed in struct
type YamlReader interface {
	Read(filePath string, output interface{}) error
}

type yamlReader struct{}

// NewYamlReader returns a fully initialised YamlReader
func NewYamlReader() YamlReader {
	return yamlReader{}
}

// Read a Yaml file into the passed in structure after validating its existence
func (yr yamlReader) Read(filePath string, output interface{}) error {
	log.Debugf("Reading file %s", filePath)
	data, err := ReadFile(filePath)
	log.Error(string(data))
	if err != nil {
		internalErr := NewParseYamlError(fmt.Sprintf("Unable to read %s", filePath))
		log.Errorf("Error: %+v\t%+v", internalErr, err)
		return internalErr
	}

	err = yaml.Unmarshal(data, output)
	if err != nil {
		internalErr := NewParseYamlError(fmt.Sprintf("Unable to parse %s", filePath))
		log.Errorf("Error: %+v\t%+v", internalErr, err)
		return internalErr
	}

	return nil
}
