package config

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/winkoz/plonk/internal/sharedtesting"
)

type ContextTestSuite struct {
	suite.Suite
	projectName         string
	deployCommand       string
	environments        map[string][]string
	customTemplatesPath string
	deployFolderName    string
	deployVariablesPath string
	deploySecretsPath   string
	targetPath          string
	iOService           *sharedtesting.IOServiceMock
	sut                 Context
}

func (suite *ContextTestSuite) SetupTest() {
	suite.projectName = "ContextTestSuite"
	suite.deployCommand = "notKubeCtl"
	suite.environments = map[string][]string{
		"base": {
			"component1",
		},
		"production": {
			"component1",
			"component2",
		},
	}
	suite.customTemplatesPath = filepath.Join("custom", "not", "a", "real", "path")
	suite.deployFolderName = "deployFolder"
	suite.deployVariablesPath = filepath.Join("not", "a", "real", "path", "variables")
	suite.deploySecretsPath = filepath.Join("not", "a", "real", "path", "secrets")
	suite.targetPath = filepath.Join("not", "a", "real", "path")
	suite.iOService = new(sharedtesting.IOServiceMock)
	suite.sut = Context{
		ProjectName:         suite.projectName,
		DeployCommand:       suite.deployCommand,
		Environments:        suite.environments,
		CustomTemplatesPath: suite.customTemplatesPath,
		DeployFolderName:    suite.deployFolderName,
		DeployVariablesPath: suite.deployVariablesPath,
		DeploySecretsPath:   suite.deploySecretsPath,
		TargetPath:          suite.targetPath,
		IOService:           suite.iOService,
	}
}

func (suite *ContextTestSuite) TestComponents_ShouldReturnTheComponentsFromEnvironment_WhenPassedEnvironmentMatchesExistingOne() {
	productionEnvironment := "production"
	gotComponents := suite.sut.Components(productionEnvironment)
	expectedComponents := suite.environments[productionEnvironment]
	assert.Equal(suite.T(), expectedComponents, gotComponents)
}

func (suite *ContextTestSuite) TestComponents_ShouldReturnTheComponentsFromBaseEnvironment_WhenPassedEnvironmentDoesNotExistInContext() {
	inexistentEnvironment := "inexistent"
	baseEnvironment := "base"
	expectedComponents := suite.environments[baseEnvironment]
	productionComponents := suite.environments["production"]
	gotComponents := suite.sut.Components(inexistentEnvironment)
	assert.Equal(suite.T(), expectedComponents, gotComponents)
	assert.NotEqual(suite.T(), gotComponents, productionComponents)
}

func TestContextTestSuite(t *testing.T) {
	suite.Run(t, new(ContextTestSuite))
}
