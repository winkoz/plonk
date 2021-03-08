package building

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/winkoz/plonk/internal/config"
	"github.com/winkoz/plonk/internal/sharedtesting"
)

type BuilderTestSuite struct {
	suite.Suite
	env                   string
	ctx                   config.Context
	buildCommand          *sharedtesting.BuildCommandMock
	versionControlCommand *sharedtesting.VersionControllerCommandMock
	sut                   Builder
}

func (suite *BuilderTestSuite) SetupTest() {
	suite.env = "test"
	suite.ctx = config.Context{
		ProjectName:   "builder_tests",
		Registry:      "registry.example.com",
		DeployCommand: "cmd",
		Environments:  map[string][]string{"test": {"test"}},
		TargetPath:    "/tmp/deploy",
	}
	suite.buildCommand = new(sharedtesting.BuildCommandMock)
	suite.versionControlCommand = new(sharedtesting.VersionControllerCommandMock)
	suite.sut = dockerBuilder{
		ctx:                   suite.ctx,
		buildCommand:          suite.buildCommand,
		versionControlCommand: suite.versionControlCommand,
	}
}

//-------------------------------------------------
// Tests
//-------------------------------------------------

func (suite *BuilderTestSuite) TestExecute_ShouldCallOrchestratorDeploy() {
	uuid := "thisisavalidheads"
	suite.setupHappyPath(uuid)
	head, err := suite.sut.Build(suite.env)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), uuid, head)
}

func TestDeployerTestSuite(t *testing.T) {
	suite.Run(t, new(BuilderTestSuite))
}

//-------------------------------------------------
// Helpers
//-------------------------------------------------

func (suite *BuilderTestSuite) setupVersionControlCommand(uuid string, err error) {
	suite.versionControlCommand.On("Head").Return(uuid, err)
}

func (suite *BuilderTestSuite) setupBuildCommand(tagName string, isLatest bool, err error) {
	suite.buildCommand.On("Build", tagName, isLatest).Return(err)
}

func (suite *BuilderTestSuite) setupHappyPath(uuid string) {
	tagName := fmt.Sprintf("%s/%s:%s-%s", suite.ctx.Registry, suite.ctx.ProjectName, suite.env, uuid)
	suite.setupVersionControlCommand(uuid, nil)
	suite.setupBuildCommand(tagName, true, nil)
}
