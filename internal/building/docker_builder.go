package building

import (
	"fmt"

	"github.com/winkoz/plonk/internal/commands"
	"github.com/winkoz/plonk/internal/config"
	"github.com/winkoz/plonk/internal/io/log"
)

type dockerBuilder struct {
	ctx                   config.Context
	buildCommand          commands.BuilderCommand
	versionControlCommand commands.VersionControllerCommand
}

// Build builds the current stack
func (b dockerBuilder) Build(stackName string) (string, error) {
	uuid, err := b.VersionControlCurrentHead()
	if err != nil {
		log.Errorf("There was an error fetching your current head. Please make sure you are using version control.")
		return "", err
	}

	tagName := fmt.Sprintf("%s/%s:%s-%s", b.ctx.Registry, b.ctx.ProjectName, stackName, uuid)
	err = b.buildCommand.Build(tagName)

	if err == nil {
		err = b.buildCommand.Push(tagName)
	}

	return tagName, err
}

// VersionControlCurrentHead returns the current position of the code in the version control
func (b dockerBuilder) VersionControlCurrentHead() (string, error) {
	return b.versionControlCommand.Head()
}
