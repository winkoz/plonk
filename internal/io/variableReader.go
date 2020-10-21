package io

import (
	"fmt"

	"github.com/winkoz/plonk/internal/io/log"
)

// VariableReader reads the variables from a stack flattening with base and returning a map
type VariableReader interface {
	GetVariables(projectName string, env string) (DeployVariables, error)
}

type variableReader struct {
	path         string
	baseFileName string
	yamlReader   YamlReader
	interpolator Interpolator
	service      Service
}

// DeployVariables variables used for interpolating the script templates
type DeployVariables struct {
	Build       map[string]string `yaml:"build,omitempty"`
	Environment map[string]string `yaml:"environment,omitempty"`
}

// NewVariableReader returns a fully configure VariableReader
func NewVariableReader(path string) VariableReader {
	service := NewService()
	return variableReader{
		path:         path,
		baseFileName: "base",
		yamlReader:   NewYamlReader(service),
		interpolator: NewInterpolator(),
		service:      service,
	}
}

// GetVariables reads the variables from a environment flattening with base and returning a map
func (vr variableReader) GetVariables(projectName string, env string) (result DeployVariables, err error) {
	vars := DeployVariables{}
	signal := log.StarTrace("GetVariables")
	defer log.StopTrace(signal, err)
	baseVariables, err := vr.read(vr.baseFileName)
	if err != nil {
		log.Error(err)
		return vars, err
	}

	envVars, err := vr.read(env)
	if err != nil {
		fullError := err.(*Error)
		if fullError.Code() != FileNotFoundError {
			log.Error(fullError)
			return vars, err
		}
	}

	buildVariables := mergeMap(envVars.Build, baseVariables.Build)
	environmentVariables := mergeMap(envVars.Environment, baseVariables.Environment)

	interpolateVars := map[string]string{
		"ENV":  env,
		"NAME": projectName,
	}

	vars.Build = vr.interpolator.SubstituteValuesInMap(interpolateVars, buildVariables)
	vars.Environment = vr.interpolator.SubstituteValuesInMap(interpolateVars, environmentVariables)

	return vars, nil
}

func (vr variableReader) read(fileName string) (DeployVariables, error) {
	vars := DeployVariables{
		Build:       map[string]string{},
		Environment: map[string]string{},
	}
	fullName := fmt.Sprintf("%s.%s", fileName, YAMLExtension)
	filePath := fmt.Sprintf("%s/%s", vr.path, fullName)
	if !vr.service.FileExists(filePath) {
		err := NewFileNotFoundError(fmt.Sprintf("%s not found at location: %s", fullName, vr.path))
		return vars, err
	}

	err := vr.yamlReader.Read(filePath, &vars)
	if err != nil {
		internalErr := NewParseVariableError(fmt.Sprintf("Unable to parse %s", filePath))
		return vars, internalErr
	}

	return vars, nil
}
