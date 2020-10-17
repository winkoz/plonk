package deployment

import (
	"github.com/winkoz/plonk/internal/config"
	"github.com/winkoz/plonk/internal/io"
	"github.com/winkoz/plonk/internal/io/log"
)

// Deployer creates a manifest file from templates and executes it with the deploy command.
type Deployer interface {
	Execute(ctx config.Context, stackName string) error
}

type deployer struct {
	varReader io.VariableReader
}

// NewDeployer creates a deployer object
func NewDeployer() Deployer {
	return deployer{
		varReader: io.NewVariableReader(),
	}
}

func (d deployer) Execute(ctx config.Context, env string) (err error) {
	signal := log.StarTrace("Execute")
	defer log.StopTrace(signal, err)

	// load variables
	variables, err := d.varReader.GetVariables(ctx.ProjectName, env)
	log.Debug(variables)
	// join file

	// execute in kubectl

	return nil
}

func (d deployer) fetchVariables(env string) (map[string]string, error) {

	return nil, nil
}
