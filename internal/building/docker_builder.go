package building

import (
	"fmt"

	"github.com/winkoz/plonk/internal/commands"
	"github.com/winkoz/plonk/internal/config"
)

type dockerBuilder struct {
	ctx          config.Context
	buildCommand commands.BuilderCommand
}

// Build builds the current stack
func (b dockerBuilder) Build(stackName string) error {
	tagName := fmt.Sprintf("%s:%s", b.ctx.ProjectName, "latest")
	err := b.buildCommand.Build(tagName)
	return err
}
