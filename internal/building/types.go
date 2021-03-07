package building

// Builder builds the project in docker and tags build in case of a release
type Builder interface {
	Build(stackName string) (string, error)
	VersionControlCurrentHead() (string, error)
}
