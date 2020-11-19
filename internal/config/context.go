package config

import (
	"github.com/winkoz/plonk/internal/io"
	"github.com/winkoz/plonk/internal/io/log"
)

// Context holds the current state of the application
type Context struct {
	// Project Config
	ProjectName   string
	DeployCommand string
	Environments  map[string][]string

	// Templates Config
	CustomTemplatesPath string

	// Deploy Config
	DeployFolderName    string
	DeployVariablesPath string
	TargetPath          string

	// Services
	IOService io.Service
}

// NewContext create a context object with the passed in project name and default values.
func NewContext(projectName string) (Context, error) {
	ioService := io.NewService()

	return Context{
		// Project Config
		ProjectName:   projectName,
		DeployCommand: deployDeployCommand,

		// Templates Config
		CustomTemplatesPath: defaultCustomTemplatesPath,

		// Deploy Config
		DeployFolderName:    deployFolderName,
		DeployVariablesPath: deployVariablesPath,
		TargetPath:          ioService.GetCurrentDir(),

		// Services
		IOService: ioService,
	}, nil
}

// NewContextFromFile Create context from the plonk.yml
func NewContextFromFile() (Context, error) {
	ioService := io.NewService()
	targetFolderPath := ioService.GetCurrentDir()
	targetConfigFilePath := targetFolderPath + "/plonk." + io.YAMLExtension

	configFile, err := loadPlonkConfigFile(ioService, targetConfigFilePath)
	if err != nil {
		log.Errorf("Failed to load %s: %v", targetConfigFilePath, err)
		return Context{}, err
	}

	return Context{
		// Project Config
		ProjectName:   configFile.Name,
		DeployCommand: configFile.Command,
		Environments:  configFile.Environments,

		// Templates Config
		CustomTemplatesPath: configFile.TemplatesDir,

		// Deploy Config
		DeployFolderName:    deployFolderName,
		DeployVariablesPath: deployVariablesPath,
		TargetPath:          ioService.GetCurrentDir(),

		// Services
		IOService: ioService,
	}, nil
}
