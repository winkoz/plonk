package building

import (
	"fmt"

	"github.com/winkoz/plonk/internal/commands"
	"github.com/winkoz/plonk/internal/config"
)

type dockerBuilder struct {
	ctx                   config.Context
	buildCommand          commands.BuilderCommand
	versionControlCommand commands.VersionControllerCommand
}

// Build builds the current stack
func (b dockerBuilder) Build(stackName string) error {
	tagName := fmt.Sprintf("%s:%s", b.ctx.ProjectName, "latest")
	err := b.buildCommand.Build(tagName)
	return err
}

// VersionControlCurrentHead returns the current position of the code in the version control
func (b dockerBuilder) VersionControlCurrentHead() (string, error) {
	return b.versionControlCommand.Head()
}
