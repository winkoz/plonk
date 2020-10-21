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
	DefaultTemplatesPath       string
	DefaultCustomTemplatesPath string

	// Deploy Config
	DeployFolderName    string
	DeployVariablesPath string
	TargetPath          string

	// Services
	IOService io.Service
}

// NewContext create a context object by reading the plonk.yml
func NewContext() (Context, error) {
	ioService := io.NewService()
	deployFolderPath := fmt.Sprintf("%s/%s", ioService.GetCurrentDir(), deployFolderName)
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
		DefaultTemplatesPath:       defaultTemplatesPath,
		DefaultCustomTemplatesPath: defaultCustomTemplatesPath,

		// Deploy Config
		DeployFolderName:    deployFolderName,
		DeployVariablesPath: deployVariablesPath,
		TargetPath:          ioService.GetCurrentDir() + "/test",

		// Services
		IOService: ioService,
	}, nil
}
