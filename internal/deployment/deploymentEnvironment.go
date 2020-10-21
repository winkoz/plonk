package deployment

// Env specifies the name of the environment to be deployed
type Env struct {
	Value string
}

// Environment sets the desired environment to be deployed
var Environment *Env = &Env{"production"}

// Set implements pflag.Value
func (e *Env) Set(value string) error {
	e.Value = value
	return nil
}

// String implements pflag.Value
func (e *Env) String() string {
	return e.Value
}

// Type implements pflag.Value
func (e Env) Type() string {
	return "string"
}
