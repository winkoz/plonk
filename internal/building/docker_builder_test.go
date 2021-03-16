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

func (suite *BuilderTestSuite) TestBuild_ShouldExecuteSuccessfully() {
	uuid := "thisisavalidheads"
	tagName := fmt.Sprintf("%s/%s:%s-%s", suite.ctx.Registry, suite.ctx.ProjectName, suite.env, uuid)
	suite.setupHappyPathBuild(uuid)
	head, err := suite.sut.Build(suite.env)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), tagName, head)
}

func (suite *BuilderTestSuite) TestPublish_ShouldExecuteSuccessfully() {
	uuid := "thisisavalidheads"
	tagName := fmt.Sprintf("%s/%s:%s-%s", suite.ctx.Registry, suite.ctx.ProjectName, suite.env, uuid)
	suite.setupHappyPathPublish(uuid)
	err := suite.sut.Publish(tagName)
	assert.Nil(suite.T(), err)
}

func (suite *BuilderTestSuite) TestGenerateTagName_ShouldExecuteSuccessfully() {
	uuid := "thisisavalidheads"
	tagName := fmt.Sprintf("%s:%s-%s", suite.ctx.ProjectName, suite.env, uuid)
	suite.setupVersionControlCommand(uuid, nil)
	res, err := suite.sut.GenerateTagName(suite.env)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), tagName, res)
}

func (suite *BuilderTestSuite) TestGenerateFullImageName_ShouldExecuteSuccessfully() {
	uuid := "thisisavalidheads"
	tagName := fmt.Sprintf("%s/%s:%s-%s", suite.ctx.Registry, suite.ctx.ProjectName, suite.env, uuid)
	suite.setupVersionControlCommand(uuid, nil)
	res, err := suite.sut.GenerateFullImageName(suite.env)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), tagName, res)
}

func (suite *BuilderTestSuite) TestBuild_ShouldFailWhenVersionControlErrors() {
	uuid := "thisisavalidheads"
	errorMessage := "this is an error message"
	suite.setupFailVersionControlPath(uuid, errorMessage)
	_, err := suite.sut.Build(suite.env)
	assert.Error(suite.T(), err)
}

func (suite *BuilderTestSuite) TestGenerateTagName_ShouldFailWhenVersionControlErrors() {
	uuid := "thisisavalidheads"
	errorMessage := "this is an error message"
	suite.setupFailVersionControlPath(uuid, errorMessage)
	_, err := suite.sut.GenerateTagName(suite.env)
	assert.Error(suite.T(), err)
}

func (suite *BuilderTestSuite) TestGenerateFullImageName_ShouldFailWhenVersionControlErrors() {
	uuid := "thisisavalidheads"
	errorMessage := "this is an error message"
	suite.setupFailVersionControlPath(uuid, errorMessage)
	_, err := suite.sut.GenerateFullImageName(suite.env)
	assert.Error(suite.T(), err)
}

func (suite *BuilderTestSuite) TestBuild_FailsWhenBuildCommandErrors() {
	uuid := "thisisavalidheads"
	errorMessage := "this is an error message"
	suite.setupFailBuildPath(uuid, errorMessage)
	_, err := suite.sut.Build(suite.env)
	assert.Error(suite.T(), err)
}

func (suite *BuilderTestSuite) Test_FailsWhenPublishCommandErrors() {
	uuid := "thisisavalidheads"
	tagName := fmt.Sprintf("%s/%s:%s-%s", suite.ctx.Registry, suite.ctx.ProjectName, suite.env, uuid)
	errorMessage := "this is an error message"
	suite.setupFailPushPath(uuid, errorMessage)
	err := suite.sut.Publish(tagName)
	assert.Error(suite.T(), err)
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

func (suite *BuilderTestSuite) setupBuildCommand(tagName string, err error) {
	suite.buildCommand.On("Build", tagName).Return(err)
}

func (suite *BuilderTestSuite) setupPushCommand(tagName string, err error) {
	suite.buildCommand.On("Push", tagName).Return(err)
}

func (suite *BuilderTestSuite) setupHappyPathBuild(uuid string) {
	tagName := fmt.Sprintf("%s/%s:%s-%s", suite.ctx.Registry, suite.ctx.ProjectName, suite.env, uuid)
	suite.setupVersionControlCommand(uuid, nil)
	suite.setupBuildCommand(tagName, nil)
}

func (suite *BuilderTestSuite) setupHappyPathPublish(uuid string) {
	tagName := fmt.Sprintf("%s/%s:%s-%s", suite.ctx.Registry, suite.ctx.ProjectName, suite.env, uuid)
	suite.setupVersionControlCommand(uuid, nil)
	suite.setupPushCommand(tagName, nil)
}

func (suite *BuilderTestSuite) setupFailVersionControlPath(uuid string, errorMessage string) {
	err := fmt.Errorf(errorMessage)
	suite.setupVersionControlCommand(uuid, err)
}

func (suite *BuilderTestSuite) setupFailBuildPath(uuid string, errorMessage string) {
	tagName := fmt.Sprintf("%s/%s:%s-%s", suite.ctx.Registry, suite.ctx.ProjectName, suite.env, uuid)
	err := fmt.Errorf(errorMessage)
	suite.setupVersionControlCommand(uuid, nil)
	suite.setupBuildCommand(tagName, err)
}

func (suite *BuilderTestSuite) setupFailPushPath(uuid string, errorMessage string) {
	tagName := fmt.Sprintf("%s/%s:%s-%s", suite.ctx.Registry, suite.ctx.ProjectName, suite.env, uuid)
	err := fmt.Errorf(errorMessage)
	suite.setupVersionControlCommand(uuid, nil)
	suite.setupBuildCommand(tagName, nil)
	suite.setupPushCommand(tagName, err)
}
