package io

// YAMLExtension defines the extension used in the package for YAML files
const YAMLExtension = "yaml"

// OwnerPermission sets read & write permissions for the user on the file/folder while group members & other users have only read access.
const OwnerPermission = 0755

// Error generic I/O error
type Error struct {
	msg string
}

// FileLocation describes the original file name and it's resolved path
type FileLocation struct {
	OriginalFilePath string
	ResolvedFilePath string
}

// Error prints out the formatted error message.
func (e *Error) Error() string {
	return e.msg
}

// NewParseVariableError returns an io.Error representing when failed to retrieve variables for a Stack
func NewParseVariableError(message string) *Error {
	return &Error{msg: message}
}

// NewParseYamlError returns an io.Error representing when failed to read a YAML file
func NewParseYamlError(message string) *Error {
	return &Error{msg: message}
}
