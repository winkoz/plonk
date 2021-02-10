package io

import (
	"fmt"

	"github.com/winkoz/plonk/internal/io/log"
	"gopkg.in/yaml.v2"
)

// YamlReader reads yaml files and loads them into a passed in struct
type YamlReader interface {
	Read(filePath string, output interface{}) error
	Parse(data []byte, output interface{}) (err error)
}

type yamlReader struct {
	service Service
}

// NewYamlReader returns a fully initialised YamlReader
func NewYamlReader(service Service) YamlReader {
	return yamlReader{
		service: service,
	}
}

// Read a Yaml file into the passed in structure after validating its existence
func (yr yamlReader) Read(filePath string, output interface{}) (err error) {
	signal := log.StartTrace("Read")
	defer log.StopTrace(signal, err)
	log.Debugf("Reading file %s", filePath)
	data, err := yr.service.ReadFile(filePath)
	if err != nil {
		internalErr := NewParseYamlError(fmt.Sprintf("Unable to read %s", filePath))
		log.Errorf("Error: %+v\t%+v", internalErr, err)
		return internalErr
	}

	err = yr.Parse(data, output)
	return
}

// Parse attempts to parse `data` into the passed `output` structure.
func (yr yamlReader) Parse(data []byte, output interface{}) (err error) {
	err = yaml.Unmarshal(data, output)
	log.Debugf("Unmarshalled yaml: %v", output)
	if err != nil {
		internalErr := NewParseYamlError(fmt.Sprintf("Unable to parse %s", string(data)))
		log.Errorf("Error: %+v\t%+v", internalErr, err)
		return internalErr
	}

	return err
}
