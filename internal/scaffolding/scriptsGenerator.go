package scaffolding

type scriptsGenerator struct {
}

// ScriptsGenerator substitues declared variables and generates new yaml configuration files.
type ScriptsGenerator interface {
}

// NewScriptsGenerator returns a fully initialised ScripterGenerator
func NewScriptsGenerator() ScriptsGenerator {
	return scriptsGenerator{}
}
