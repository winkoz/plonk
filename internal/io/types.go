package io

// YAMLExtension defines the extension used in the package for YAML files
const YAMLExtension = "yaml"

// OwnerPermission sets read & write permissions for the user on the file/folder while group members & other users have only read access.
const OwnerPermission = 0644

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

// Transformator is a function that receives a byte array and returns a transformated array
type Transformator func(input []byte) []byte
