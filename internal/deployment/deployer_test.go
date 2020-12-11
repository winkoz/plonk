package deployment

import (
	"errors"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/winkoz/plonk/internal/config"
	"github.com/winkoz/plonk/internal/io"
	"github.com/winkoz/plonk/internal/scaffolding"
	"github.com/winkoz/plonk/internal/sharedtesting"
)

type DeployerTestSuite struct {
	suite.Suite
	manifestFileName    string
	manifestFile        string
	env                 string
	variables           io.DeployVariables
	secrets             io.DeploySecrets
	ctx                 config.Context
	ioService           *sharedtesting.IOServiceMock
	varReader           *sharedtesting.VariableReaderMock
	secretReader        *sharedtesting.SecretReaderMock
	templateReader      *scaffolding.TemplateReaderMock
	templateParser      *sharedtesting.TemplateParserMock
	orchestratorCommand *sharedtesting.OrchestratorCommandMock
	sut                 Deployer
}

func (suite *DeployerTestSuite) SetupTest() {
	suite.ioService = new(sharedtesting.IOServiceMock)
	suite.env = "test"
	suite.manifestFileName = "template-manifest.yaml"
	suite.manifestFile = `name: custom-template
variables:
	build:
		TEST_BUILD_VAR: "custom-template-build"
	environment:
		TEST_ENV_VAR: "custom-template-env"
	files:
	- plonk.yaml`
	suite.variables = io.DeployVariables{
		Build: map[string]string{
			"var1": "value1",
		},
	}
	suite.secrets = io.DeploySecrets{
		Secret: map[string]string{
			"secret1": "value1",
		},
	}
	suite.ctx = config.Context{
		ProjectName:         "deployer_tests",
		DeployCommand:       "cmd",
		Environments:        map[string][]string{"test": {"test"}},
		CustomTemplatesPath: "/tmp/",
		DeployFolderName:    "deploy",
		DeployVariablesPath: "deploy/variables",
		DeploySecretsPath:   "deploy/secrets",
		TargetPath:          "/tmp/deploy",
		IOService:           suite.ioService,
	}
	suite.varReader = new(sharedtesting.VariableReaderMock)
	suite.secretReader = new(sharedtesting.SecretReaderMock)
	suite.templateReader = new(scaffolding.TemplateReaderMock)
	suite.templateParser = new(sharedtesting.TemplateParserMock)
	suite.orchestratorCommand = new(sharedtesting.OrchestratorCommandMock)
	suite.sut = deployer{
		ctx:                 suite.ctx,
		varReader:           suite.varReader,
		secretsReader:       suite.secretReader,
		templateReader:      suite.templateReader,
		ioService:           suite.ioService,
		templateParser:      suite.templateParser,
		orchestratorCommand: suite.orchestratorCommand,
	}
}

//-------------------------------------------------
// Tests
//-------------------------------------------------

func (suite *DeployerTestSuite) TestExecuteSuccessfullyCallsDeployOnTheOrchestratorWithGeneratedDeployFile() {
	suite.setupHappyPath()
	assert.Nil(suite.T(), suite.sut.Execute(suite.ctx, suite.env))
}

func (suite *DeployerTestSuite) TestExecuteReturnsErrorWhenUnableToLoadTheEnvironmentTemplates() {
	expectedErr := errors.New("TestExecuteReadEnvTemplate")
	suite.setupVariablesAndSecretsMocks(nil, nil)
	suite.setupTemplateReader(nil, expectedErr)
	err := suite.sut.Execute(suite.ctx, suite.env)
	assert.EqualError(suite.T(), err, expectedErr.Error())
}

func TestDeployerTestSuite(t *testing.T) {
	suite.Run(t, new(DeployerTestSuite))
}

//-------------------------------------------------
// Helpers
//-------------------------------------------------

func (suite *DeployerTestSuite) setupVariablesAndSecretsMocks(varError error, secretsError error) {
	suite.varReader.
		On("GetVariablesFromFile", suite.ctx.ProjectName, suite.env).
		Once().
		Return(suite.variables, varError)

	suite.secretReader.
		On("GetSecretsFromFile", suite.ctx.ProjectName, suite.env).
		Once().
		Return(suite.secrets, secretsError)
}

func (suite *DeployerTestSuite) setupTemplateReader(templateData *scaffolding.TemplateData, err error) {
	var result scaffolding.TemplateData
	if td := templateData; td != nil {
		result = *td
	} else {
		result = scaffolding.TemplateData{}
	}
	suite.templateReader.
		On("Read", suite.env).
		Once().
		Return(result, err)
}

func (suite *DeployerTestSuite) setupIOServiceWrite(err error) {
	deployFullPath := filepath.Join(suite.ctx.TargetPath, suite.ctx.DeployFolderName, "deploy.yaml")
	suite.ioService.
		On("Write", deployFullPath, suite.manifestFile+"\n---\n").
		Once().
		Return(err)
}

func (suite *DeployerTestSuite) setupIOServiceReadFile(err error) {
	suite.ioService.
		On("ReadFile", suite.manifestFileName).
		Once().
		Return([]byte(suite.manifestFile), err)
}

func (suite *DeployerTestSuite) setupOrchestrator(err error) {
	deployFullPath := filepath.Join(suite.ctx.DeployFolderName, "deploy.yaml")
	suite.orchestratorCommand.
		On("Deploy", suite.env, deployFullPath).
		Once().
		Return(err)
}

func (suite *DeployerTestSuite) setupTemplateParser(result string, err error) {
	suite.templateParser.
		On("Parse", mock.Anything, suite.manifestFile).
		Once().
		Return(result, err)
}

func (suite *DeployerTestSuite) setupHappyPath() {
	suite.setupVariablesAndSecretsMocks(nil, nil)
	td := &scaffolding.TemplateData{
		Name:             "TestExecuteSuccess",
		Manifests:        []string{suite.manifestFileName},
		FilesLocation:    []io.FileLocation{},
		Files:            []string{},
		DefaultVariables: suite.variables,
	}
	suite.setupTemplateReader(td, nil)
	suite.setupIOServiceWrite(nil)
	suite.setupIOServiceReadFile(nil)
	suite.setupTemplateParser(suite.manifestFile, nil)
	suite.setupOrchestrator(nil)
}
