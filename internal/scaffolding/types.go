package scaffolding

// ProjectDefinition defines the files that a project contains
type ProjectDefinition []string

// BaseProjectFiles defines the list of files required to initialise a basic project
var BaseProjectFiles = ProjectDefinition{
	"service.yaml",
	"ingress.yaml",
	"main-deployment.yaml",
}

// Error generic I/O error
type Error struct {
	msg string
}

func (e *Error) Error() string {
	return e.msg
}

// NewScaffolderFileNotFound returns an scaffolding.Error representing when failed to find a file
func NewScaffolderFileNotFound(message string) *Error {
	return &Error{msg: message}
}
