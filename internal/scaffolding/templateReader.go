package scaffolding

import (
	"fmt"

	"github.com/winkoz/plonk/internal/io"
	"github.com/winkoz/plonk/internal/io/log"
)

// TemplateReader reads a project configuration from disk based on a passed configuration file
type TemplateReader interface {
	Read(templateName string) (TemplateData, error)
}

type templateReader struct {
	defaultTemplatePath string
	customTemplatePath  string
	yamlReader          io.YamlReader
	service             io.Service
}

// TemplateData contains the list of script files
type TemplateData struct {
	Name              string `yaml:"name"`
	VariablesFileName string `yaml:"variables,omitempty"`
	VariablesContents string
	Manifests         []string `yaml:"manifests,omitempty"`
	FilesLocation     []io.FileLocation
	Files             []string `yaml:"files,omitempty"`
}

// NewTemplateReader returns a fully configure TemplateReader
func NewTemplateReader(defaultTemplatePath string, customTemplatePath string) TemplateReader {
	return templateReader{
		defaultTemplatePath: defaultTemplatePath,
		customTemplatePath:  customTemplatePath,
		yamlReader:          io.NewYamlReader(),
		service:             io.NewService(),
	}
}

func (tr templateReader) Read(templateName string) (TemplateData, error) {
	templateData := TemplateData{
		FilesLocation: []io.FileLocation{},
		Files:         []string{},
	}

	templateDefinitionFilePath, err := tr.fileLocator(templateName, fmt.Sprintf("template-definition.%s", io.YAMLExtension))
	if err != nil {
		log.Error(err)
		return templateData, err
	}

	err = tr.yamlReader.Read(templateDefinitionFilePath, &templateData)
	if err != nil {
		log.Error(err)
		return templateData, err
	}

	resolvedFilePaths, err := tr.locateFiles(templateName, templateData.Files)
	log.Debug(resolvedFilePaths)
	if err != nil {
		log.Error(err)
		return templateData, err
	}

	if len(resolvedFilePaths) != len(templateData.Files) {
		err := fmt.Errorf("mismatched number of resolved files paths %d from original files %d", len(resolvedFilePaths), len(templateData.Files))
		log.Errorf("Unable to read template. %v", err)
	}

	for idx, originalFile := range templateData.Files {
		fileLocation := io.FileLocation{
			OriginalFilePath: originalFile,
			ResolvedFilePath: resolvedFilePaths[idx],
		}

		templateData.FilesLocation = append(templateData.FilesLocation, fileLocation)
	}

	templateData.Manifests, err = tr.locateFiles(templateName, templateData.Manifests)
	log.Debug(templateData.Manifests)
	if err != nil {
		log.Error(err)
		return templateData, err
	}

	templateData.VariablesFileName, err = tr.fileLocator(templateName, templateData.VariablesFileName)
	if err != nil {
		log.Error(err)
		return templateData, err
	}

	if len(templateData.VariablesFileName) > 0 {
		variablesData, err := tr.service.ReadFile(templateData.VariablesFileName)
		if err != nil {
			log.Error(err)
			return templateData, err
		}
		templateData.VariablesContents = string(variablesData)
	}

	return templateData, nil
}

func (tr templateReader) locateFiles(templateName string, filePaths []string) ([]string, error) {
	locatedFilePaths := []string{}
	for _, fileName := range filePaths {
		filePath, filerr := tr.fileLocator(templateName, fileName)
		if filerr != nil {
			log.Error(filerr)
			return locatedFilePaths, filerr
		}
		locatedFilePaths = append(locatedFilePaths, filePath)
	}
	return locatedFilePaths, nil
}

func (tr templateReader) fileLocator(templateName string, fileName string) (string, error) {
	filePath := fmt.Sprintf("%s/%s", templateName, fileName)
	if tr.customTemplatePath != "" {
		customPath := fmt.Sprintf("%s/%s", tr.customTemplatePath, filePath)
		if tr.service.FileExists(customPath) {
			return customPath, nil
		}
	}

	defaultPath := fmt.Sprintf("%s/%s", tr.defaultTemplatePath, filePath)
	if tr.service.FileExists(defaultPath) {
		return defaultPath, nil
	}

	err := NewScaffolderFileNotFound(fmt.Sprintf("Template not found %s. Locations [%s, %s]", fileName, tr.customTemplatePath, tr.defaultTemplatePath))
	log.Error(err)

	return "", err
}
