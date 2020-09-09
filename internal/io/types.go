package io

// YAMLExtension defines the extension used in the package for YAML files
const YAMLExtension = "yaml"

// Error generic I/O error
type Error struct {
	msg string
}

func (e *Error) Error() string {
	return e.msg
}

// NewParseVariableError returns an io.Error representing when failed to retrieve variables for a Stack
func NewParseVariableError(message string) *Error {
	return &Error{msg: message}
}
