package scaffolding

import (
	"fmt"

	"github.com/winkoz/plonk/internal/io"
	"github.com/winkoz/plonk/internal/io/log"
)

// TemplateReader reads a project configuration from disk based on a passed configuration file
type TemplateReader interface {
	Read(configurationFileName string) (TemplateData, error)
}

type templateReader struct {
	defaultTemplatePath string
	customTemplatePath  string
	yamlReader          io.YamlReader
}

// TemplateData contains the list of script files
type TemplateData struct {
	Name      string   `yaml:"name"`
	Origin    []string `yaml:"origin,omitempty"`
	Variables []string `yaml:"variables,omitempty"`
	Manifests []string `yaml:"manifests,omitempty"`
}

// NewTemplateReader returns a fully configure TemplateReader
func NewTemplateReader(defaultTemplatePath string, customTemplatePath string) TemplateReader {
	return templateReader{
		defaultTemplatePath: defaultTemplatePath,
		customTemplatePath:  customTemplatePath,
		yamlReader:          io.NewYamlReader(),
	}
}

func (tf templateReader) Read(configurationFileName string) (TemplateData, error) {
	configFilePath, err := tf.fileLocator(fmt.Sprintf("%s.%s", configurationFileName, io.YAMLExtension))
	templateData := TemplateData{}

	if err != nil {
		log.Error(err)
		return templateData, err
	}

	err = tf.yamlReader.Read(configFilePath, &templateData)
	if err != nil {
		log.Error(err)
		return templateData, err
	}

	templateData.Origin, err = tf.locateFiles(templateData.Origin)
	log.Debug(templateData.Origin)
	if err != nil {
		log.Error(err)
		return templateData, err
	}

	templateData.Manifests, err = tf.locateFiles(templateData.Manifests)
	log.Debug(templateData.Manifests)
	if err != nil {
		log.Error(err)
		return templateData, err
	}

	templateData.Variables, err = tf.locateFiles(templateData.Variables)
	if err != nil {
		log.Error(err)
		return templateData, err
	}

	return templateData, nil
}

func (tf templateReader) locateFiles(filePaths []string) ([]string, error) {
	locatedFilePaths := []string{}
	for _, fileName := range filePaths {
		filePath, filerr := tf.fileLocator(fileName)
		if filerr != nil {
			log.Error(filerr)
			return locatedFilePaths, filerr
		}
		locatedFilePaths = append(locatedFilePaths, filePath)
	}
	return locatedFilePaths, nil
}

func (tf templateReader) fileLocator(fileName string) (string, error) {
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

	err := NewScaffolderFileNotFound(fmt.Sprintf("Template not found %s. Locations [%s, %s]", fileName, tf.customTemplatePath, tf.defaultTemplatePath))
	log.Error(err)

	return "", err
}
