package deployment

import (
	"fmt"

	"github.com/winkoz/plonk/internal/commands"
	"github.com/winkoz/plonk/internal/config"
	"github.com/winkoz/plonk/internal/io/log"
)

// Destroyer executes a command that destroys the namespace for the speficied `env`.
// Important: This is a destructive action!
type Destroyer interface {
	Execute(ctx config.Context, env string) error
}

type destroyer struct {
	ctx                 config.Context
	orchestratorCommand commands.OrchestratorCommand
}

// NewDestroyer creates a destroyer object
func NewDestroyer(ctx config.Context) Destroyer {
	return destroyer{
		ctx:                 ctx,
		orchestratorCommand: commands.NewOrchestrator(ctx, "kubectl"),
	}
}

func (d destroyer) Execute(ctx config.Context, env string) (err error) {
	signal := log.StartTrace("Execute")
	defer log.StopTrace(signal, err)

	log.Debugf("Ctx: \n%+v", ctx)

	// execute in kubectl
	cmd := d.orchestratorCommand.Destroy

	namespace := fmt.Sprintf("%s-%s", d.ctx.ProjectName, env)
	err = cmd(namespace)
	if err != nil {
		log.Errorf("Cannot execute destroy command %s. error = %+v", d.ctx.DeployCommand, err)
		return err
	}

	return nil
}
