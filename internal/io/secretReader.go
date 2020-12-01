package io

import (
	"fmt"
	"path/filepath"

	"github.com/winkoz/plonk/internal/io/log"
)

// SecretReader reads the secrets from a stack flattening with base and returning a map
type SecretReader interface {
	GetVariablesFromFile(projectName string, env string) (DeploySecrets, error)
}

type secretReader struct {
	path         string
	baseFileName string
	yamlReader   YamlReader
	service      Service
}

// DeploySecrets variables used for interpolating the script templates
type DeploySecrets struct {
	Secret map[string]string `yaml:"secret,omitempty"`
}

// NewSecretReader returns a fully configure SecretReader
func NewSecretReader(path string) VariableReader {
	service := NewService()
	return variableReader{
		path:         path,
		baseFileName: "base",
		yamlReader:   NewYamlReader(service),
		interpolator: NewInterpolator(),
		service:      service,
	}
}

// GetSecretFromFile reads the secret from a environment flattening with base and returning a map
func (sr secretReader) GetSecretFromFile(projectName string, env string) (result DeploySecrets, err error) {
	secrets := DeploySecrets{}
	signal := log.StartTrace("DeploySecrets")
	defer log.StopTrace(signal, err)
	baseSecrets, err := sr.read(sr.baseFileName)
	if err != nil {
		log.Error(err)
		return secrets, err
	}

	envSecrets, err := sr.read(env)
	if err != nil {
		fullError := err.(*Error)
		if fullError.Code() != FileNotFoundError {
			log.Error(fullError)
			return secrets, err
		}
	}

	mergedSecrets := MergeStringMap(envSecrets.Secret, baseSecrets.Secret)
	secrets.Secret = mergedSecrets

	return secrets, nil
}

func (sr secretReader) read(fileName string) (DeploySecrets, error) {
	secrets := DeploySecrets{
		Secret: map[string]string{},
	}
	fullName := fmt.Sprintf("%s.%s", fileName, YAMLExtension)
	filePath := filepath.Join(sr.path, fullName)
	if !sr.service.FileExists(filePath) {
		err := NewFileNotFoundError(fmt.Sprintf("%s not found at location: %s", fullName, sr.path))
		return secrets, err
	}

	err := sr.yamlReader.Read(filePath, &secrets)
	if err != nil {
		internalErr := NewParseSecretError(fmt.Sprintf("Unable to parse %s", filePath))
		return secrets, internalErr
	}

	return secrets, nil
}
