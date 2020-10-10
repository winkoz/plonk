package deployment

// Deployer creates a manifest file from templates and executes it with the deploy command.
type Deployer interface {
	Execute(stackName string) error
}

type deployer struct {
}

// NewDeployer creates a deployer object
func NewDeployer() Deployer {
	return deployer{}
}

func (d deployer) Execute(stackName string) error {

	// load variables

	// join file

	// execute in kubectl

	return nil
}
