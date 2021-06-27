package config

import (
	"path/filepath"

	"github.com/winkoz/plonk/internal"
	"github.com/winkoz/plonk/internal/io"
	"github.com/winkoz/plonk/internal/io/log"
)

// Context holds the current state of the application
type Context struct {
	// Runtime Context
	RuntimeContext internal.RuntimeContext

	// Project Config
	ProjectName   string
	Registry      string
	DeployCommand string
	BuildCommand  string
	Environments  map[string][]string

	// Templates Config
	CustomTemplatesPath string

	// Deploy Config
	DeployFolderName    string
	DeployVariablesPath string
	DeploySecretsPath   string
	TargetPath          string

	// Services
	IOService io.Service
}

// NewContext create a context object with the passed in project name and default values.
func NewContext(projectName string) (Context, error) {
	ioService := io.NewService()

	return Context{
		// Runtime Context
		RuntimeContext: internal.RuntimeContext{},

		// Project Config
		ProjectName:   projectName,
		Registry:      registryDefaultValue,
		DeployCommand: deployDeployCommand,
		BuildCommand:  deployBuildCommand,

		// Templates Config
		CustomTemplatesPath: defaultCustomTemplatesPath,

		// Deploy Config
		DeployFolderName:    deployFolderName,
		DeployVariablesPath: deployVariablesPath,
		DeploySecretsPath:   deploySecretsPath,
		TargetPath:          ioService.GetCurrentDir(),

		// Services
		IOService: ioService,
	}, nil
}

// NewContextFromFile Create context from the plonk.yml
func NewContextFromFile() (Context, error) {
	ioService := io.NewService()
	targetFolderPath := ioService.GetCurrentDir()
	targetConfigFilePath := filepath.Join(targetFolderPath, "plonk."+io.YAMLExtension)

	configFile, err := loadPlonkConfigFile(ioService, targetConfigFilePath)
	if err != nil {
		log.Errorf("Failed to load %s: %v", targetConfigFilePath, err)
		return Context{}, err
	}

	return Context{
		// Project Config
		ProjectName:   configFile.Name,
		Registry:      configFile.Registry,
		DeployCommand: configFile.Command,
		BuildCommand:  deployBuildCommand,
		Environments:  configFile.Environments,

		// Templates Config
		CustomTemplatesPath: configFile.TemplatesDir,

		// Deploy Config
		DeployFolderName:    deployFolderName,
		DeployVariablesPath: deployVariablesPath,
		DeploySecretsPath:   deploySecretsPath,
		TargetPath:          ioService.GetCurrentDir(),

		// Services
		IOService: ioService,
	}, nil
}

// Components returns a list of components found in the passed in environment.
// If none found it returns the components found in `base`
func (c Context) Components(environment string) []string {
	envComponents, exist := c.Environments[environment]
	if !exist {
		envComponents = c.Environments[internal.BaseEnvironmentKey]
	}

	return envComponents
}
