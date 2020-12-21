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
	manifestFilesName     []string
	manifestFile          string
	env                   string
	variables             io.DeployVariables
	defaultBuildVariables map[string]string
	defaultEnvVariables   map[string]string
	secrets               io.DeploySecrets
	ctx                   config.Context
	ioService             *sharedtesting.IOServiceMock
	varReader             *sharedtesting.VariableReaderMock
	secretReader          *sharedtesting.SecretReaderMock
	templateReader        *scaffolding.TemplateReaderMock
	templateParser        *sharedtesting.TemplateParserMock
	orchestratorCommand   *sharedtesting.OrchestratorCommandMock
	namespaceTemplateData scaffolding.TemplateData
	sut                   Deployer
}

func (suite *DeployerTestSuite) SetupTest() {
	suite.ioService = new(sharedtesting.IOServiceMock)
	suite.env = "test"
	suite.defaultBuildVariables = map[string]string{
		"var2": "value2",
		"var1": "changed",
	}
	suite.defaultEnvVariables = map[string]string{
		"env": "test",
	}
	suite.variables = io.DeployVariables{
		Build: map[string]string{
			"var1": "value1",
		},
		Environment: map[string]string{},
	}
	suite.secrets = io.DeploySecrets{
		Secret: map[string]string{
			"secret1": "value1",
		},
	}
	suite.manifestFilesName = []string{"template-manifest.yaml", "template-manifest2.yaml"}
	suite.manifestFile = `name: custom-template
variables:
	build:
		TEST_BUILD_VAR: "custom-template-build"
	environment:
		TEST_ENV_VAR: "custom-template-env"
	files:
	- plonk.yaml`
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
	suite.namespaceTemplateData = scaffolding.TemplateData{
		Name:             "namespace",
		Manifests:        []string{},
		FilesLocation:    []io.FileLocation{},
		Files:            []string{},
		DefaultVariables: suite.variables,
	}
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

func (suite *DeployerTestSuite) TestExecute_ShouldCallOrchestratorDeploy() {
	suite.setupHappyPath()
	suite.setupIOServiceDelete()
	assert.Nil(suite.T(), suite.sut.Execute(suite.ctx, suite.env, false))
	suite.ioService.AssertNumberOfCalls(suite.T(), "DeletePath", 1)
}

func (suite *DeployerTestSuite) TestExecute_ShouldError_WhenUnableToLoadTheEnvironmentTemplates() {
	expectedErr := errors.New("TestExecuteReadEnvTemplate")
	suite.setupVariablesAndSecretsMocks(nil, nil)
	suite.setupTemplateReader(nil, expectedErr)
	err := suite.sut.Execute(suite.ctx, suite.env, false)
	assert.EqualError(suite.T(), err, expectedErr.Error())
	suite.ioService.AssertNotCalled(suite.T(), "DeletePath")
}

func (suite *DeployerTestSuite) TestExecute_ShouldError_WhenUnableToReadManifestFile() {
	expectedErr := errors.New("TestExecuteMergeManifest")
	td := &scaffolding.TemplateData{
		Name:             "test",
		Manifests:        suite.manifestFilesName,
		FilesLocation:    []io.FileLocation{},
		Files:            []string{},
		DefaultVariables: suite.variables,
	}
	suite.setupVariablesAndSecretsMocks(nil, nil)
	suite.setupTemplateReader(td, nil)
	suite.setupIOServiceReadFile(expectedErr)
	err := suite.sut.Execute(suite.ctx, suite.env, false)
	assert.EqualError(suite.T(), err, expectedErr.Error())
	suite.ioService.AssertNotCalled(suite.T(), "DeletePath")
}

func (suite *DeployerTestSuite) TestExecute_ShouldError_WhenUnableToParseManifestFile() {
	expectedErr := errors.New("TestExecuteParseManifest")
	td := &scaffolding.TemplateData{
		Name:             "test",
		Manifests:        suite.manifestFilesName,
		FilesLocation:    []io.FileLocation{},
		Files:            []string{},
		DefaultVariables: suite.variables,
	}
	suite.setupVariablesAndSecretsMocks(nil, nil)
	suite.setupTemplateReader(td, nil)
	suite.setupIOServiceReadFile(nil)
	suite.setupTemplateParser(expectedErr)
	err := suite.sut.Execute(suite.ctx, suite.env, false)
	assert.EqualError(suite.T(), err, expectedErr.Error())
	suite.ioService.AssertNotCalled(suite.T(), "DeletePath")
}

func (suite *DeployerTestSuite) TestExecute_ShouldError_WhenUnableToWriteDeployFile() {
	expectedErr := errors.New("TestExecuteWriteDeploy")
	td := &scaffolding.TemplateData{
		Name:             "test",
		Manifests:        suite.manifestFilesName,
		FilesLocation:    []io.FileLocation{},
		Files:            []string{},
		DefaultVariables: suite.variables,
	}
	suite.setupVariablesAndSecretsMocks(nil, nil)
	suite.setupTemplateReader(td, nil)
	suite.setupIOServiceReadFile(nil)
	suite.setupTemplateParser(nil)
	suite.setupTemplateParser(nil)
	suite.setupIOServiceWrite(expectedErr)
	err := suite.sut.Execute(suite.ctx, suite.env, false)
	assert.EqualError(suite.T(), err, expectedErr.Error())
	suite.ioService.AssertNotCalled(suite.T(), "DeletePath")
}

func (suite *DeployerTestSuite) TestExecute_ShouldError_WhenUnableToExecuteTheDeployCommand() {
	expectedErr := errors.New("TestExecuteOrchestratorDeploy")
	td := &scaffolding.TemplateData{
		Name:             "test",
		Manifests:        suite.manifestFilesName,
		FilesLocation:    []io.FileLocation{},
		Files:            []string{},
		DefaultVariables: suite.variables,
	}
	suite.setupVariablesAndSecretsMocks(nil, nil)
	suite.setupTemplateReader(td, nil)
	suite.setupIOServiceReadFile(nil)
	suite.setupTemplateParser(nil)
	suite.setupTemplateParser(nil)
	suite.setupIOServiceWrite(nil)
	suite.setupOrchestrator(expectedErr)
	err := suite.sut.Execute(suite.ctx, suite.env, false)
	assert.EqualError(suite.T(), err, expectedErr.Error())
	suite.ioService.AssertNotCalled(suite.T(), "DeletePath")
}

func (suite *DeployerTestSuite) TestExecute_ShouldCallOrchestratorDiff_WhenDryRunIsTrue() {
	suite.setupHappyPath()                                   // Set happy path as it is the same as Deploy
	suite.orchestratorCommand.ExpectedCalls = []*mock.Call{} // Reset the Deploy mocked call so we can configure Diff only

	deployFullPath := filepath.Join(suite.ctx.DeployFolderName, "deploy.yaml")
	suite.orchestratorCommand.
		On("Diff", suite.env, deployFullPath).
		Once().
		Return(nil)
	assert.Nil(suite.T(), suite.sut.Execute(suite.ctx, suite.env, true))
	suite.ioService.AssertNotCalled(suite.T(), "DeletePath")
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
	var testTemplateData scaffolding.TemplateData
	if td := templateData; td != nil {
		testTemplateData = *td
	} else {
		testTemplateData = scaffolding.TemplateData{
			Name: "test",
		}
	}
	templatesData := []scaffolding.TemplateData{
		suite.namespaceTemplateData,
		testTemplateData,
	}
	for _, v := range templatesData {
		suite.templateReader.
			On("Read", v.Name).
			Once().
			Return(v, err)
	}
}

func (suite *DeployerTestSuite) setupIOServiceWrite(err error) {
	deployFullPath := filepath.Join(suite.ctx.TargetPath, suite.ctx.DeployFolderName, "deploy.yaml")
	mergedFile := suite.manifestFile + "\n---\n" + suite.manifestFile + "\n---\n"
	suite.ioService.
		On("Write", deployFullPath, mergedFile).
		Once().
		Return(err)
}

func (suite *DeployerTestSuite) setupIOServiceReadFile(err error) {
	for _, manifestFileName := range suite.manifestFilesName {
		suite.ioService.
			On("ReadFile", manifestFileName).
			Once().
			Return([]byte(suite.manifestFile), err)
	}
}

func (suite *DeployerTestSuite) setupIOServiceDelete() {
	deployFullPath := filepath.Join(suite.ctx.TargetPath, suite.ctx.DeployFolderName, "deploy.yaml")
	suite.ioService.
		On("DeletePath", deployFullPath).
		Once()
}

func (suite *DeployerTestSuite) setupOrchestrator(err error) {
	deployFullPath := filepath.Join(suite.ctx.DeployFolderName, "deploy.yaml")
	suite.orchestratorCommand.
		On("Deploy", suite.env, deployFullPath).
		Once().
		Return(err)
}

func (suite *DeployerTestSuite) setupTemplateParser(err error) {
	suite.templateParser.
		On("Parse", mock.Anything, suite.manifestFile).
		Once().
		Return(suite.manifestFile, err)
}

func (suite *DeployerTestSuite) setupHappyPath() {
	suite.setupVariablesAndSecretsMocks(nil, nil)
	td := &scaffolding.TemplateData{
		Name:          "test",
		Manifests:     suite.manifestFilesName,
		FilesLocation: []io.FileLocation{},
		Files:         []string{},
		DefaultVariables: struct {
			Build       map[string]string `yaml:"build,omitempty"`
			Environment map[string]string `yaml:"environment,omitempty"`
		}{
			Build:       suite.defaultBuildVariables,
			Environment: suite.defaultEnvVariables,
		},
	}
	suite.setupTemplateReader(td, nil)
	suite.setupIOServiceWrite(nil)
	suite.setupIOServiceReadFile(nil)
	suite.setupIOServiceDelete()
	suite.setupTemplateParser(nil)
	suite.setupTemplateParser(nil)
	suite.setupOrchestrator(nil)
}
