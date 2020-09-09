package io

// VariableReader reads the variables from a stack flattening with base and returning a map
type VariableReader interface {
	GetVariables(stackName string) (map[string]string, error)
}

type variableReader struct{}

// NewVariableReader returns a fully configure VariableReader
func NewVariableReader() VariableReader {
	return variableReader{}
}

// GetVariables reads the variables from a stack flattening with base and returning a map
func (vr variableReader) GetVariables(stackName string) (map[string]string, error) {

	return nil, nil
}
