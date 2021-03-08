package commands

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/winkoz/plonk/internal/config"
	"github.com/winkoz/plonk/internal/io"
	"github.com/winkoz/plonk/internal/sharedtesting"
)

type DockerTestSuite struct {
	suite.Suite
	executor     *sharedtesting.ExecutorMock
	ctx          config.Context
	buildCommand string
	env          string
	namespace    string
	component    *string
	manifestPath string
	targetPath   string
	sat          dockerCommand
}

func (suite *DockerTestSuite) SetupTest() {
	suite.env = "testing"
	suite.targetPath = filepath.Join("this", "is", "some", "path")
	suite.manifestPath = filepath.Join("this", "is", "not", "a", "real", "path")
	suite.executor = new(sharedtesting.ExecutorMock)
	suite.buildCommand = "notDocker"
	suite.ctx = config.Context{
		BuildCommand: suite.buildCommand,
		Registry:     "madeupregistry",
		TargetPath:   suite.targetPath,
		ProjectName:  "Plonk-Docker-Test",
	}
	suite.namespace = fmt.Sprintf("%s-%s", suite.ctx.ProjectName, suite.env)
	suite.component = nil
	suite.sat = dockerCommand{
		executor:     suite.executor,
		interpolator: io.NewInterpolator(),
		ctx:          suite.ctx,
	}
}

//-------------------------------------------------
// Tests
//-------------------------------------------------

//----- Deploy Tests

func (suite *DockerTestSuite) TestBuild_ShouldCallDockerBuildWithTagArgs() {
	tagName := fmt.Sprintf("%s/%s:%s", suite.ctx.Registry, suite.ctx.ProjectName, "madeup-tag")
	args := []string{
		"build",
		"--no-cache",
		"--tag",
		tagName,
		".",
	}
	suite.setupExecutor(args, nil, nil)
	err := suite.sat.Build(tagName, false)
	suite.verifyExecutor(args)
	assert.Nil(suite.T(), err)
}

func TestDockerTestSuite(t *testing.T) {
	suite.Run(t, new(DockerTestSuite))
}

//-------------------------------------------------
// Helpers
//-------------------------------------------------

func (suite *DockerTestSuite) setupExecutor(args []string, output []byte, err error) {
	curatedOutput := output
	if output == nil {
		curatedOutput = make([]byte, 0)
	}
	suite.executor.On(
		"Run",
		suite.buildCommand,
		args,
	).Return(
		curatedOutput, err,
	)
}

func (suite *DockerTestSuite) verifyExecutor(args []string) {
	suite.executor.AssertCalled(suite.T(),
		"Run",
		suite.buildCommand,
		args,
	)
}
