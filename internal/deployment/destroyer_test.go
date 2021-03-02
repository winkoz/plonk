package deployment

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/winkoz/plonk/internal/config"
	"github.com/winkoz/plonk/internal/sharedtesting"
)

type DestroyerTestSuite struct {
	suite.Suite
	env                 string
	ctx                 config.Context
	namespace           string
	orchestratorCommand *sharedtesting.OrchestratorCommandMock
	sut                 Destroyer
}

func (suite *DestroyerTestSuite) SetupTest() {
	suite.env = "test"
	suite.ctx = config.Context{
		ProjectName:         "destroyer_tests",
		DeployCommand:       "cmd",
		Environments:        map[string][]string{"test": {"test"}},
		CustomTemplatesPath: "/tmp/",
		DeployFolderName:    "deploy",
		DeployVariablesPath: "deploy/variables",
		DeploySecretsPath:   "deploy/secrets",
		TargetPath:          "/tmp/deploy",
	}
	suite.namespace = fmt.Sprintf("%s-%s", suite.ctx.ProjectName, suite.env)
	suite.orchestratorCommand = new(sharedtesting.OrchestratorCommandMock)
	suite.sut = destroyer{
		ctx:                 suite.ctx,
		orchestratorCommand: suite.orchestratorCommand,
	}
}

//-------------------------------------------------
// Tests
//-------------------------------------------------

func (suite *DestroyerTestSuite) TestExecute_ShouldCallOrchestratorDestroy() {
	suite.setupOrchestrator(nil)
	assert.Nil(suite.T(), suite.sut.Execute(suite.ctx, suite.env))
}

func (suite *DestroyerTestSuite) TestExecute_ShouldError_WhenUnableToExecuteTheDestroyCommand() {
	expectedErr := errors.New("TestExecuteOrchestratorDestroy")
	suite.setupOrchestrator(expectedErr)
	err := suite.sut.Execute(suite.ctx, suite.env)
	assert.EqualError(suite.T(), err, expectedErr.Error())
}

func TestDestroyerTestSuite(t *testing.T) {
	suite.Run(t, new(DestroyerTestSuite))
}

// //-------------------------------------------------
// // Helpers
// //-------------------------------------------------

func (suite *DestroyerTestSuite) setupOrchestrator(err error) {
	suite.orchestratorCommand.
		On("Destroy", suite.namespace).
		Once().
		Return(err)
}
