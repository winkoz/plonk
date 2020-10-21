package deployment

import (
	"github.com/winkoz/plonk/internal/config"
	"github.com/winkoz/plonk/internal/io"
	"github.com/winkoz/plonk/internal/io/log"
	"github.com/winkoz/plonk/internal/scaffolding"
)

// Deployer creates a manifest file from templates and executes it with the deploy command.
type Deployer interface {
	Execute(ctx config.Context, stackName string) error
}

type deployer struct {
	ctx            config.Context
	varReader      io.VariableReader
	templateReader scaffolding.TemplateReader
}

// NewDeployer creates a deployer object
func NewDeployer(ctx config.Context) Deployer {
	return deployer{
		ctx:            ctx,
		varReader:      io.NewVariableReader(),
		templateReader: scaffolding.NewTemplateReader(ctx),
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
