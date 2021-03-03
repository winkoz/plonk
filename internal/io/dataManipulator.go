package io

import (
	"encoding/base64"
	"regexp"
	"strings"
)

// DataManipulator processes data (strings and bytes) and manipulates them accordingly to produce the desired outcomes (byes or strings)/
type DataManipulator interface {
	YamlToMapArray(maybeYaml string) ([]map[string]string, error)
	Base64Encode(v []byte) (string, error)
	StringToBytes(s string) ([]byte, error)
	Indent(s string, numberOfSpaces int) (string, error)
}

type dataManipulator struct {
	service Service
}

// NewDataManipulator creates a new fully configured DataManipulator object
func NewDataManipulator(service Service) DataManipulator {
	return dataManipulator{
		service: service,
	}
}

// YamlToMapArray attempts to parse `maybeYaml` into an array of objects and will return it as a map from string to string array.
func (dm dataManipulator) YamlToMapArray(maybeYaml string) ([]map[string]string, error) {
	output := []map[string]string{}
	yamlReader := NewYamlReader(dm.service)
	err := yamlReader.Parse([]byte(maybeYaml), &output)
	return output, err
}

// StringToBytes converts the passed in `str` into is byte array representation.
func (dm dataManipulator) StringToBytes(str string) ([]byte, error) {
	return []byte(str), nil
}

// Base64Encode received the bytes array representation of a string and encodes it into base64.
func (dm dataManipulator) Base64Encode(v []byte) (string, error) {
	return base64.StdEncoding.EncodeToString(v), nil
}

// Indent indents every line of the `source` string by the `numberOfSpaces` passed and returns the transformed string.
func (dm dataManipulator) Indent(source string, numberOfSpaces int) (string, error) {
	indent := "\n" + strings.Repeat(" ", numberOfSpaces)
	re := regexp.MustCompile(`\r?\n`)
	return re.ReplaceAllString(source, indent), nil
}
