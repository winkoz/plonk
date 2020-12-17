package management

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/winkoz/plonk/internal/config"
	"github.com/winkoz/plonk/internal/sharedtesting"
)

type ManagerTestSuite struct {
	suite.Suite
	ctx                 config.Context
	env                 string
	orchestratorCommand *sharedtesting.OrchestratorCommandMock
	renderer            *sharedtesting.RendererMock
	namespace           string
	sut                 Manager
}

func (suite *ManagerTestSuite) SetupTest() {
	suite.env = "testing"
	suite.orchestratorCommand = new(sharedtesting.OrchestratorCommandMock)
	suite.ctx = config.Context{
		ProjectName: "Manager-GetPods-Test",
	}
	suite.namespace = fmt.Sprintf("%s-%s", suite.ctx.ProjectName, suite.env)
	suite.renderer = new(sharedtesting.RendererMock)
	suite.sut = manager{
		ctx:                 suite.ctx,
		orchestratorCommand: suite.orchestratorCommand,
		renderer:            suite.renderer,
	}
}

//-------------------------------------------------
// Tests
//-------------------------------------------------

//----- GetPods Tests

func (suite *ManagerTestSuite) TestGetPods_ShouldCallOrchestratorGetPods() {
	suite.orchestratorCommand.
		On("GetPods", mock.AnythingOfType("string")).
		Once().
		Return(make([]byte, 0), nil)
	suite.renderer.
		On("RenderComponents", mock.Anything).
		Once()
	_, _ = suite.sut.GetPods(suite.env)
	assert.True(suite.T(), suite.orchestratorCommand.AssertCalled(suite.T(), "GetPods", suite.namespace))
}

func (suite *ManagerTestSuite) TestGetPods_ShouldPassTheOrchestratorOutputToTheRenderer_WhenOrchestratorGetPodsSucceeds() {
	outputStr := suite.T().Name()
	output := []byte(outputStr)
	suite.orchestratorCommand.
		On("GetPods", mock.AnythingOfType("string")).
		Once().
		Return(output, nil)
	suite.renderer.
		On("RenderComponents", mock.Anything).
		Once()
	_, _ = suite.sut.GetPods(suite.env)
	suite.renderer.AssertCalled(suite.T(), "RenderComponents", output)
}

func (suite *ManagerTestSuite) TestGetPods_ShouldError_WhenOrchestratorFailsToExecuteCommand() {
	expectedErr := errors.New("kubectl error getting pods")
	suite.orchestratorCommand.
		On("GetPods", suite.namespace).
		Once().
		Return(make([]byte, 0), expectedErr)
	_, gotErr := suite.sut.GetPods(suite.env)
	assert.EqualError(suite.T(), gotErr, expectedErr.Error())
}

//----- GetLogs Tests

func (suite *ManagerTestSuite) TestGetLogs_ShouldCallOrchestratorGetLogs() {
	suite.orchestratorCommand.
		On("GetLogs", mock.AnythingOfType("string")).
		Once().
		Return(make([]byte, 0), nil)
	suite.renderer.
		On("RenderLogs", mock.Anything).
		Once()
	_, _ = suite.sut.GetLogs(suite.env)
	assert.True(suite.T(), suite.orchestratorCommand.AssertCalled(suite.T(), "GetLogs", suite.namespace))
}

func (suite *ManagerTestSuite) TestGetLogs_ShouldPassTheOrchestratorOutputToTheRenderer_WhenOrchestratorGetLogsSucceeds() {
	outputStr := suite.T().Name()
	output := []byte(outputStr)
	suite.orchestratorCommand.
		On("GetLogs", mock.AnythingOfType("string")).
		Once().
		Return(output, nil)
	suite.renderer.
		On("RenderLogs", mock.Anything).
		Once()
	_, _ = suite.sut.GetLogs(suite.env)
	suite.renderer.AssertCalled(suite.T(), "RenderLogs", output)
}

func (suite *ManagerTestSuite) TestGetLogs_ShouldError_WhenOrchestratorFailsToExecuteCommand() {
	expectedErr := errors.New("kubectl error getting logs")
	suite.orchestratorCommand.
		On("GetLogs", suite.namespace).
		Once().
		Return(make([]byte, 0), expectedErr)
	_, gotErr := suite.sut.GetLogs(suite.env)
	assert.EqualError(suite.T(), gotErr, expectedErr.Error())
}

func TestManagerTestSuite(t *testing.T) {
	suite.Run(t, new(ManagerTestSuite))
}
