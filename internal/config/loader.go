package config

import (
	"github.com/winkoz/plonk/internal/io"
	"github.com/winkoz/plonk/internal/io/log"
)

const defaultDeployCommand = "kubectl"

// PlonkConfigFile structure
type PlonkConfigFile struct {
	Name         string   `yaml:"name"`
	Command      string   `yaml:"command, omitempty"`
	TemplatesDir string   `yaml:"templates_dir, omitempty"`
	Templates    []string `yaml:"templates, omitempty"`
}

func loadPlonkConfigFile(ioService io.Service, filePath string) (PlonkConfigFile, error) {
	config := PlonkConfigFile{
		Command:      deployDeployCommand,
		TemplatesDir: defaultTemplatesPath,
	}

	yamlReader := io.NewYamlReader(ioService)
	err := yamlReader.Read(filePath, &config)
	if err != nil {
		log.Errorf("Couldn't load plonk config file in: %s, error: %v", filePath, err)
		return config, err
	}

	return config, nil
}
