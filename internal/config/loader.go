package config

import (
	"github.com/winkoz/plonk/internal/io"
	"github.com/winkoz/plonk/internal/io/log"
)

const defaultDeployCommand = "kubectl"

// PlonkConfigFile structure
type PlonkConfigFile struct {
	Name         string              `yaml:"name"`
	Command      string              `yaml:"command,omitempty"`
	Registry     string              `yaml:"registry,omitempty"`
	TemplatesDir string              `yaml:"templates,omitempty"`
	Environments map[string][]string `yaml:"environments"`
}

func loadPlonkConfigFile(ioService io.Service, filePath string) (PlonkConfigFile, error) {
	config := PlonkConfigFile{
		Command:      deployDeployCommand,
		Registry:     registryDefaultValue,
		TemplatesDir: defaultCustomTemplatesPath,
	}

	yamlReader := io.NewYamlReader(ioService)
	err := yamlReader.Read(filePath, &config)
	log.Debugf("Config: %+v", config)
	if err != nil {
		log.Errorf("Couldn't load plonk config file in: %s, error: %v", filePath, err)
		return config, err
	}

	return config, nil
}
