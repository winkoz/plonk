package scaffolding

type variablesGenerator struct {
}

// VariablesGenerator substitues declared variables and generates new yaml variables files.
type VariablesGenerator interface {
}

// NewVariablesGenerator returns a fully initialised VariablesGenerator
func NewVariablesGenerator() VariablesGenerator {
	return variablesGenerator{}
}
