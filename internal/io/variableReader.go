package io

import (
	"fmt"

	"github.com/winkoz/plonk/internal/io/log"
)

// VariableReader reads the variables from a stack flattening with base and returning a map
type VariableReader interface {
	GetVariables(stackName string) (map[string]string, error)
}

type variableReader struct {
	path           string
	customFileName string
	baseFileName   string
	yamlReader     YamlReader
	interpolator   Interpolator
	service        Service
}

// DeployVariables variables used for interpolating the script templates
type DeployVariables struct {
	Variables map[string]string `yaml:"variables,omitempty"`
}

// NewVariableReader returns a fully configure VariableReader
func NewVariableReader() VariableReader {
	service := NewService()
	path := service.GetCurrentDir()
	return variableReader{
		path:           path,
		baseFileName:   "base",
		customFileName: "custom",
		yamlReader:     NewYamlReader(),
		interpolator:   NewInterpolator(),
		service:        service,
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

	mergedResult := mergeMap(baseVariables, customVariables)
	log.Debugf("Merged maps: %v", mergedResult)

	stackNameVars := map[string]string{
		"STACK": stackName,
	}

	result := vr.interpolator.SubstituteValuesInMap(stackNameVars, mergedResult)

	return result, nil
}

func (vr variableReader) read(fileName string) (map[string]string, error) {
	fullName := fmt.Sprintf("%s.%s", fileName, YAMLExtension)
	filePath := fmt.Sprintf("%s/%s", vr.path, fullName)
	if !vr.service.FileExists(filePath) {
		err := NewParseVariableError(fmt.Sprintf("%s not found at location: %s", fullName, vr.path))
		log.Error(err)
		return nil, err
	}

	fileVariables := DeployVariables{}
	err := vr.yamlReader.Read(filePath, &fileVariables)
	if err != nil {
		internalErr := NewParseVariableError(fmt.Sprintf("Unable to parse %s", filePath))
		log.Errorf("Error: %+v\t%+v", internalErr, err)
		return nil, internalErr
	}

	return fileVariables.Variables, nil
}
