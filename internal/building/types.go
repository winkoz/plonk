package building

// Builder builds the project in docker and tags build in case of a release
type Builder interface {
	Build(stackName string) (tagName string, err error)
	Publish(tagName string) error
	GenerateTagName(stackName string) (string, error)
}
