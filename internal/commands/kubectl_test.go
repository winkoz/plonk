package commands

import (
	"errors"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/winkoz/plonk/internal/config"
	"github.com/winkoz/plonk/internal/io"
	"github.com/winkoz/plonk/internal/sharedtesting"
)

type KubectlTestSuite struct {
	suite.Suite
	executor       *sharedtesting.ExecutorMock
	ctx            config.Context
	deployCommand  string
	env            string
	namespace      string
	deploymentName string
	component      *string
	manifestPath   string
	targetPath     string
	sat            kubectlCommand
}

func (suite *KubectlTestSuite) SetupTest() {
	suite.env = "testing"
	suite.targetPath = filepath.Join("this", "is", "some", "path")
	suite.manifestPath = filepath.Join("this", "is", "not", "a", "real", "path")
	suite.executor = new(sharedtesting.ExecutorMock)
	suite.deployCommand = "notKubeCtl"
	suite.ctx = config.Context{
		DeployCommand: suite.deployCommand,
		TargetPath:    suite.targetPath,
		ProjectName:   "Plonk-KubeCtl-Test",
	}
	suite.namespace = fmt.Sprintf("%s-%s", suite.ctx.ProjectName, suite.env)
	suite.deploymentName = fmt.Sprintf("%s-%s-deployment", suite.ctx.ProjectName, suite.env)
	suite.component = nil
	suite.sat = kubectlCommand{
		executor:     suite.executor,
		interpolator: io.NewInterpolator(),
		ctx:          suite.ctx,
	}
}

//-------------------------------------------------
// Tests
//-------------------------------------------------

//----- Deploy Tests

func (suite *KubectlTestSuite) TestDeploy_ShouldCallExecutorWithApplyCommand() {
	args := []string{"apply", "-f", suite.manifestPath}
	suite.setupExecutor(args, nil, nil)
	err := suite.sat.Deploy(suite.manifestPath)
	suite.verifyExecutor(args)
	assert.Nil(suite.T(), err)
}

func (suite *KubectlTestSuite) TestDeploy_ShouldInterpolatePathInTheCommand_WhenSuccessfulCall() {
	suite.sat.ctx.DeployCommand = "notKubeCtl -p $PWD"
	args := []string{"-p", suite.targetPath, "apply", "-f", suite.manifestPath}
	suite.setupExecutor(args, nil, nil)
	err := suite.sat.Deploy(suite.manifestPath)
	suite.verifyExecutor(args)
	assert.Nil(suite.T(), err)
}

func (suite *KubectlTestSuite) TestDeploy_ShouldReturnAnError_WhenExecutorFails() {
	expectedErr := errors.New(suite.T().Name())
	suite.setupExecutor([]string{"apply", "-f", suite.manifestPath}, nil, expectedErr)
	gotErr := suite.sat.Deploy(suite.manifestPath)
	assert.EqualError(suite.T(), gotErr, expectedErr.Error())
}

//----- Destroy Tests

func (suite *KubectlTestSuite) TestDestroy_ShouldCallExecutorWithDestroy() {
	args := []string{"delete", "namespace", suite.namespace}
	suite.setupExecutor(args, nil, nil)
	err := suite.sat.Destroy(suite.namespace)
	suite.verifyExecutor(args)
	assert.Nil(suite.T(), err)
}

func (suite *KubectlTestSuite) TestDestroy_ShouldReturnAnError_WhenExecutorFails() {
	expectedErr := errors.New(suite.T().Name())
	suite.setupExecutor([]string{"delete", "namespace", suite.namespace}, nil, expectedErr)
	gotErr := suite.sat.Destroy(suite.namespace)
	assert.EqualError(suite.T(), gotErr, expectedErr.Error())
}

//----- Diff Tests

func (suite *KubectlTestSuite) TestDiff_ShouldCallExecutorWithDiffCommand() {
	args := []string{"diff", "-f", suite.manifestPath}
	suite.setupExecutor(args, nil, nil)
	err := suite.sat.Diff(suite.manifestPath)
	suite.verifyExecutor(args)
	assert.Nil(suite.T(), err)
}

func (suite *KubectlTestSuite) TestDiff_ShouldInterpolatePathInTheCommand_WhenSuccessfulCall() {
	suite.sat.ctx.DeployCommand = "notKubeCtl -p $PWD"
	args := []string{"-p", suite.targetPath, "diff", "-f", suite.manifestPath}
	suite.setupExecutor(args, nil, nil)
	err := suite.sat.Diff(suite.manifestPath)
	suite.verifyExecutor(args)
	assert.Nil(suite.T(), err)
}

func (suite *KubectlTestSuite) TestDiff_ShouldReturnAnError_WhenExecutorFails() {
	expectedErr := errors.New(suite.T().Name())
	suite.setupExecutor([]string{"diff", "-f", suite.manifestPath}, nil, expectedErr)
	gotErr := suite.sat.Diff(suite.manifestPath)
	assert.EqualError(suite.T(), gotErr, expectedErr.Error())
}

//----- GetPods Tests

func (suite *KubectlTestSuite) TestGetPods_ShouldCallExecutorWithGetCommand() {
	args := []string{"get", "pods", "--namespace", suite.namespace, "-o", "json"}
	suite.setupExecutor(args, nil, nil)
	_, err := suite.sat.GetPods(suite.namespace)
	suite.verifyExecutor(args)
	assert.Nil(suite.T(), err)
}

func (suite *KubectlTestSuite) TestGetPods_ShouldForwardOutputFromExecutor_WhenExecutorSucceeds() {
	args := []string{"get", "pods", "--namespace", suite.namespace, "-o", "json"}
	expectedOutput := []byte(suite.T().Name())
	suite.setupExecutor(args, expectedOutput, nil)
	gotOutput, err := suite.sat.GetPods(suite.namespace)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), string(expectedOutput), string(gotOutput))
}

func (suite *KubectlTestSuite) TestGetPods_ShouldReturnAnError_WhenExecutorFails() {
	expectedErr := errors.New(suite.T().Name())
	suite.setupExecutor([]string{"get", "pods", "--namespace", suite.namespace, "-o", "json"}, nil, expectedErr)
	_, gotErr := suite.sat.GetPods(suite.namespace)
	assert.EqualError(suite.T(), gotErr, expectedErr.Error())
}

//----- GetLogs Tests

func (suite *KubectlTestSuite) TestGetLogs_ShouldCallExecutorWithGetCommand() {
	args := []string{"logs", "--namespace", suite.namespace, "-l", "app=" + suite.namespace}
	suite.setupExecutor(args, nil, nil)
	_, err := suite.sat.GetLogs(suite.namespace, suite.component)
	suite.verifyExecutor(args)
	assert.Nil(suite.T(), err)
}

func (suite *KubectlTestSuite) TestGetLogs_ShouldCallExecutorWithGetCommandAndComponentLabel_WhenComponentIsNotNil() {
	component := suite.T().Name()
	suite.component = &component
	args := []string{"logs", "--namespace", suite.namespace, "-l", "app=" + suite.namespace, "-l", "component=" + component}
	suite.setupExecutor(args, nil, nil)
	_, err := suite.sat.GetLogs(suite.namespace, suite.component)
	suite.verifyExecutor(args)
	assert.Nil(suite.T(), err)
}

func (suite *KubectlTestSuite) TestGetLogs_ShouldForwardOutputFromExecutor_WhenExecutorSucceeds() {
	args := []string{"logs", "--namespace", suite.namespace, "-l", "app=" + suite.namespace}
	expectedOutput := []byte(suite.T().Name())
	suite.setupExecutor(args, expectedOutput, nil)
	gotOutput, err := suite.sat.GetLogs(suite.namespace, suite.component)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), string(expectedOutput), string(gotOutput))
}

func (suite *KubectlTestSuite) TestGetLogs_ShouldReturnAnError_WhenExecutorFails() {
	expectedErr := errors.New(suite.T().Name())
	suite.setupExecutor([]string{"logs", "--namespace", suite.namespace, "-l", "app=" + suite.namespace}, nil, expectedErr)
	_, gotErr := suite.sat.GetLogs(suite.namespace, suite.component)
	assert.EqualError(suite.T(), gotErr, expectedErr.Error())
}

func TestKubectlTestSuite(t *testing.T) {
	suite.Run(t, new(KubectlTestSuite))
}

//----- Restart Tests

func (suite *KubectlTestSuite) TestRestart_ShouldCallExecutorWithRestartCommand() {
	args := []string{"--namespace", suite.namespace, "rollout", "restart", "deployment", suite.deploymentName}
	suite.setupExecutor(args, nil, nil)
	_, err := suite.sat.Restart(suite.namespace, suite.deploymentName)
	suite.verifyExecutor(args)
	assert.Nil(suite.T(), err)
}

func (suite *KubectlTestSuite) TestRestart_ShouldCallExecutorWithRestartDeploymentsCommand_WhenAllIsPassed() {
	args := []string{"--namespace", suite.namespace, "rollout", "restart", "deployments"}
	suite.setupExecutor(args, nil, nil)
	_, err := suite.sat.Restart(suite.namespace, "--all")
	suite.verifyExecutor(args)
	assert.Nil(suite.T(), err)
}

func (suite *KubectlTestSuite) TestRestart_ShouldForwardOutputFromExecutor_WhenExecutorSucceeds() {
	args := []string{"--namespace", suite.namespace, "rollout", "restart", "deployment", suite.deploymentName}
	expectedOutput := []byte(suite.T().Name())
	suite.setupExecutor(args, expectedOutput, nil)
	gotOutput, err := suite.sat.Restart(suite.namespace, suite.deploymentName)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), string(expectedOutput), string(gotOutput))
}

func (suite *KubectlTestSuite) TestRestart_ShouldReturnAnError_WhenExecutorFails() {
	expectedErr := errors.New(suite.T().Name())
	suite.setupExecutor([]string{"--namespace", suite.namespace, "rollout", "restart", "deployment", suite.deploymentName}, nil, expectedErr)
	_, gotErr := suite.sat.Restart(suite.namespace, suite.deploymentName)
	assert.EqualError(suite.T(), gotErr, expectedErr.Error())
}

//-------------------------------------------------
// Helpers
//-------------------------------------------------

func (suite *KubectlTestSuite) setupExecutor(args []string, output []byte, err error) {
	curatedOutput := output
	if output == nil {
		curatedOutput = make([]byte, 0)
	}
	suite.executor.On(
		"Run",
		suite.deployCommand,
		args,
	).Return(
		curatedOutput, err,
	)
}

func (suite *KubectlTestSuite) verifyExecutor(args []string) {
	suite.executor.AssertCalled(suite.T(),
		"Run",
		suite.deployCommand,
		args,
	)
}
