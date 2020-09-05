package scaffolding

type scaffolder struct {
}

// Scaffolder runs the scaffolding logic to generate new plonk services
type Scaffolder interface {
}

// NewScaffolder returns a fully initialised Scaffolder
func NewScaffolder() Scaffolder {
	return scaffolder{}
}
