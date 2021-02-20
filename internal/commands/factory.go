package commands

import (
	"github.com/winkoz/plonk/internal/config"
	"github.com/winkoz/plonk/internal/io"
)

// NewOrchestrator this will return a class to execute actions on the orchestrator command line tool
func NewOrchestrator(ctx config.Context, orchestratorType string) OrchestratorCommand {
	return kubectlCommand{
		executor:     NewExecutor(),
		interpolator: io.NewInterpolator(),
		ctx:          ctx,
	}
}

// NewBuilder this will return a class to execute actions on the builder command line tool
func NewBuilder(ctx config.Context, builderType string) BuilderCommand {
	return dockerCommand{
		executor:     NewExecutor(),
		interpolator: io.NewInterpolator(),
		ctx:          ctx,
	}
}
