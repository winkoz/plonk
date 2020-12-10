package io

// YAMLExtension defines the extension used in the package for YAML files
const YAMLExtension = "yaml"

// OwnerPermission sets read & write permissions for the user on the file/folder while group members & other users have only read access.
const OwnerPermission = 0755

// ErrorCode describes an error identifier
type ErrorCode int

const (
	// NoError represent a non existent error or nil
	NoError ErrorCode = iota - 2
	// BaseError is the most basic error of the app
	BaseError
	// ParseVariableError happens when a variables file couldn't be parsed
	ParseVariableError
	// ParseSecretError happens when a secrets file couldn't be parsed
	ParseSecretError
	// FileNotFoundError happens when a file can't be found in disk
	FileNotFoundError
	// ParseYamlError happens when a yaml file can't be parsed
	ParseYamlError
)

// Error generic I/O error
type Error struct {
	code ErrorCode
	msg  string
}

// BinaryFile indicates if the file is included in the current binary
const BinaryFile = "BIN_FILE"

// FileLocation describes the original file name and it's resolved path
type FileLocation struct {
	OriginalFilePath string
	ResolvedFilePath string
}

// Error prints out the formatted error message.
func (e *Error) Error() string {
	return e.msg
}

// Code prints out the formatted error code.
func (e *Error) Code() ErrorCode {
	return e.code
}

// NewParseVariableError returns an io.Error representing when failed to retrieve variables for an environment
func NewParseVariableError(message string) *Error {
	return &Error{code: ParseVariableError, msg: message}
}

// NewParseSecretError returns an io.Error representing when failed to retrieve secrets for an environment
func NewParseSecretError(message string) *Error {
	return &Error{code: ParseSecretError, msg: message}
}

// NewFileNotFoundError returns an io.Error representing when failed to retrieve variables for a environment
func NewFileNotFoundError(message string) *Error {
	return &Error{code: FileNotFoundError, msg: message}
}

// NewParseYamlError returns an io.Error representing when failed to read a YAML file
func NewParseYamlError(message string) *Error {
	return &Error{code: ParseYamlError, msg: message}
}
