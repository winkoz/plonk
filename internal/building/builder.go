package building

import (
	"github.com/winkoz/plonk/internal/commands"
	"github.com/winkoz/plonk/internal/config"
)

// NewBuilder constructs a new builder class
func NewBuilder(ctx config.Context) Builder {
	buildCmd := commands.NewBuilder(ctx, "docker")
	vcCmd := commands.NewVersionController(ctx)
	return dockerBuilder{
		versionControlCommand: vcCmd,
		buildCommand:          buildCmd,
		ctx:                   ctx,
	}
}
