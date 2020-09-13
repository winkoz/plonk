package scaffolding

// ProjectDefinition defines the files that a project contains
type ProjectDefinition []string

// BaseProjectFiles defines the list of files required to initialise a basic project
var BaseProjectFiles = ProjectDefinition{
	"service.yaml",
	"ingress.yaml",
	"main-deployment.yaml",
}
