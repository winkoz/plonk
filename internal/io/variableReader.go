package io

import (
	"fmt"
	"io/ioutil"

	"github.com/winkoz/plonk/internal/io/log"
	"gopkg.in/yaml.v2"
)

// VariableReader reads the variables from a stack flattening with base and returning a map
type VariableReader interface {
	GetVariables(stackName string) (map[string]string, error)
}

type variableReader struct {
	path           string
	customFileName string
	baseFileName   string
}

// DeployVariables variables used for interpolating the script templates
type DeployVariables struct {
	Variables map[string]string `yaml:"variables,omitempty"`
}

// NewVariableReader returns a fully configure VariableReader
func NewVariableReader() VariableReader {
	path := GetCurrentDir()
	return variableReader{
		path:           path,
		baseFileName:   "base",
		customFileName: "custom",
	}
}

// GetVariables reads the variables from a stack flattening with base and returning a map
func (vr variableReader) GetVariables(stackName string) (map[string]string, error) {
	baseVariables, err := vr.read(vr.baseFileName)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	customVariables, err := vr.read(vr.customFileName)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	result := mergeMap(baseVariables, customVariables)
	log.Debug(result)

	return result, nil
}

func (vr variableReader) read(fileName string) (map[string]string, error) {
	fullName := fmt.Sprintf("%s.%s", fileName, YAMLExtension)
	filePath := fmt.Sprintf("%s/%s", vr.path, fullName)
	if !FileExists(filePath) {
		err := NewParseVariableError(fmt.Sprintf("%s not found at location: %s", fullName, vr.path))
		log.Error(err)
		return nil, err
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		internalErr := NewParseVariableError(fmt.Sprintf("Unable to read %s", filePath))
		log.Errorf("Error: %+v\t%+v", internalErr, err)
		return nil, internalErr
	}

	fileVariables := DeployVariables{}
	err = yaml.Unmarshal(data, &fileVariables)
	if err != nil {
		internalErr := NewParseVariableError(fmt.Sprintf("Unable to parse %s", filePath))
		log.Errorf("Error: %+v\t%+v", internalErr, err)
		return nil, internalErr
	}

	return fileVariables.Variables, nil
}
