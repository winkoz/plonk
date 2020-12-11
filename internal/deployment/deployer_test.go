package deployment

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/winkoz/plonk/internal/config"
	"github.com/winkoz/plonk/internal/scaffolding"
	"github.com/winkoz/plonk/internal/sharedtesting"
)

type DeployerTestSuite struct {
	suite.Suite
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
	suite.ctx = config.Context{
		ProjectName:         "deployer_tests",
		DeployCommand:       "cmd",
		Environments:        map[string][]string{},
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

func (suite *DeployerTestSuite) TestExample() {
	// assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
	// suite.Equal(5, suite.VariableThatShouldStartAtFive)
}

func TestDeployerTestSuite(t *testing.T) {
	suite.Run(t, new(DeployerTestSuite))
}
