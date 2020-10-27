package config

import (
	"fmt"

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
	TemplatesPath       string
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
		TemplatesPath:       defaultTemplatesPath,
		CustomTemplatesPath: defaultCustomTemplatesPath,

		// Deploy Config
		DeployFolderName:    deployFolderName,
		DeployVariablesPath: deployVariablesPath,
		// TODO: Remove /test
		TargetPath: ioService.GetCurrentDir() + "/test",

		// Services
		IOService: ioService,
	}, nil
}

func NewContextFromFile() (Context, error) {
	ioService := io.NewService()
	//TODO: Remove the '/test' part
	deployFolderPath := fmt.Sprintf("%s/test/%s", ioService.GetCurrentDir(), deployFolderName)
	deployConfigFilePath := deployFolderPath + "/plonk." + io.YAMLExtension

	configFile, err := loadPlonkConfigFile(ioService, deployConfigFilePath)
	if err != nil {
		log.Errorf("Failed to load %s: %v", deployConfigFilePath, err)
		return Context{}, err
	}

	return Context{
		// Project Config
		ProjectName:   configFile.Name,
		DeployCommand: configFile.Command,
		Environments:  configFile.Environments,

		// Templates Config
		TemplatesPath:       configFile.TemplatesDir,
		CustomTemplatesPath: defaultCustomTemplatesPath,

		// Deploy Config
		DeployFolderName:    deployFolderName,
		DeployVariablesPath: deployVariablesPath,
		// TODO: Remove /test
		TargetPath: ioService.GetCurrentDir() + "/test",

		// Services
		IOService: ioService,
	}, nil
}
