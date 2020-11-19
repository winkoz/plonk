package scaffolding

import (
	"fmt"

	"github.com/winkoz/plonk/data"
	"github.com/winkoz/plonk/internal/config"
	"github.com/winkoz/plonk/internal/io"
	"github.com/winkoz/plonk/internal/io/log"
)

// TemplateReader reads a project configuration from disk based on a passed configuration file
type TemplateReader interface {
	Read(templateName string) (TemplateData, error)
}

type templateReader struct {
	ctx        config.Context
	yamlReader io.YamlReader
	service    io.Service
}

// TemplateData contains the list of script files
type TemplateData struct {
	Name             string   `yaml:"name"`
	Manifests        []string `yaml:"manifests,omitempty"`
	FilesLocation    []io.FileLocation
	Files            []string `yaml:"files,omitempty"`
	DefaultVariables struct {
		Build       map[string]string `yaml:"build,omitempty"`
		Environment map[string]string `yaml:"environment,omitempty"`
	} `yaml:"variables,omitempty"`
}

// NewTemplateReader returns a fully configure TemplateReader
func NewTemplateReader(ctx config.Context) TemplateReader {
	ioService := io.NewService()
	return templateReader{
		ctx:        ctx,
		yamlReader: io.NewYamlReader(ioService),
		service:    ioService,
	}
}

func (tr templateReader) Read(templateName string) (templateData TemplateData, err error) {
	signal := log.StartTrace("Read")
	defer log.StopTrace(signal, err)

	templateData = TemplateData{
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
	if tr.ctx.CustomTemplatesPath != "" {
		customPath := fmt.Sprintf("%s/%s", tr.ctx.CustomTemplatesPath, filePath)
		if tr.service.FileExists(customPath) {
			return customPath, nil
		}
	}

	_, err := data.Asset(fmt.Sprintf("templates/%s", filePath))
	if err == nil {
		return fmt.Sprintf("%s/templates/%s", io.BinaryFile, filePath), nil
	}

	log.Debugf("Template asset error: %+v", err)
	err = NewScaffolderFileNotFound(fmt.Sprintf("Template not found %s. Locations [%s]", fileName, tr.ctx.CustomTemplatesPath))
	log.Error(err)

	return "", err
}
