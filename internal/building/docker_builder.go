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
	uuid, err := b.versionControlCurrentHead()
	if err != nil {
		log.Errorf("There was an error fetching your current head. Please make sure you are using version control.")
		return "", err
	}

	tagName := fmt.Sprintf("%s/%s:%s-%s", b.ctx.Registry, b.ctx.ProjectName, stackName, uuid)
	err = b.buildCommand.Build(tagName)

	return tagName, err
}

// Build builds the current stack
func (b dockerBuilder) Publish(tagName string) error {
	return b.buildCommand.Push(tagName)
}

// GenerateTagName returns a generated tag name based on the hash of the current git head
func (b dockerBuilder) GenerateTagName(stackName string) (string, error) {
	uuid, err := b.versionControlCurrentHead()
	if err != nil {
		log.Errorf("There was an error fetching your current head. Please make sure you are using version control.")
		return "", err
	}

	tagName := fmt.Sprintf("%s:%s-%s", b.ctx.ProjectName, stackName, uuid)

	return tagName, nil
}

// GenerateFullImageName returns a generated full image name which is the registry address plus the tag name
func (b dockerBuilder) GenerateFullImageName(stackName string) (string, error) {
	tag, err := b.GenerateTagName(stackName)
	if err != nil {
		log.Errorf("There was an error generating the tag.")
		return "", err
	}

	tagName := fmt.Sprintf("%s/%s", b.ctx.Registry, tag)

	return tagName, nil
}

// *************************************************************************************
// Private methods
// *************************************************************************************

// versionControlCurrentHead returns the current position of the code in the version control
func (b dockerBuilder) versionControlCurrentHead() (string, error) {
	return b.versionControlCommand.Head()
}
