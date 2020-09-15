package scaffolding

import (
	"fmt"

	"github.com/prometheus/common/log"
	"github.com/winkoz/plonk/io"
)

// TemplateFetcher fetches a project configuration from disk based on a passed configuration file
type TemplateFetcher interface {
	FetchConfiguration(configurationFileName string) ([]string, error)
}

type templateFetcher struct {
	defaultTemplatePath string
	customTemplatePath  string
	yamlReader          io.YamlReader
}

// TemplateConfiguration contains the list of script files
type TemplateConfiguration struct {
	Scripts []string `yaml:"scripts,omitempty"`
}

// NewTemplateFetcher returns a fully configure TemplateFetcher
func NewTemplateFetcher(defaultTemplatePath string, customTemplatePath string) TemplateFetcher {
	return templateFetcher{
		defaultTemplatePath: defaultTemplatePath,
		customTemplatePath:  customTemplatePath,
		yamlReader:          io.NewYamlReader(),
	}
}

func (tf templateFetcher) FetchConfiguration(configurationFileName string) ([]string, error) {
	configFilePath, err := tf.fileLocator(configurationFileName)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	templateConfig := TemplateConfiguration{}
	err = tf.yamlReader.Read(configFilePath, &templateConfig)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	filePaths := []string{}
	for _, scriptName := range templateConfig.Scripts {
		scriptPath, scriptErr := tf.fileLocator(scriptName)
		if scriptErr != nil {
			log.Error(scriptErr)
			return nil, scriptErr
		}
		filePaths = append(filePaths, scriptPath)
	}

	return filePaths, nil
}

func (tf templateFetcher) fileLocator(fileName string) (string, error) {
	if tf.customTemplatePath != "" {
		customPath := fmt.Sprintf("%s/%s", tf.customTemplatePath, fileName)
		if io.FileExists(customPath) {
			return customPath, nil
		}
	}

	defaultPath := fmt.Sprintf("%s/%s", tf.defaultTemplatePath, fileName)
	if io.FileExists(defaultPath) {
		return defaultPath, nil
	}

	return "", NewScaffolderFileNotFound(fmt.Sprintf("Template not found %s. Locations [%s, %s]", fileName, tf.customTemplatePath, tf.defaultTemplatePath))
}
