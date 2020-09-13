package scaffolding

type scaffolder struct {
}

// Scaffolder runs the scaffolding logic to generate new plonk services
type Scaffolder interface {
	Init(targetPath string, templatesPath string, name string) error
}

// NewScaffolder returns a fully initialised Scaffolder
func NewScaffolder() Scaffolder {
	return scaffolder{}
}

// Init initializes the basic structure of a plonk project
func (s scaffolder) Init(targetPath string, templatesPath string, name string) error {
	return nil
}
