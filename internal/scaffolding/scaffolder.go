package scaffolding

import (
	"fmt"
	"os"

	"github.com/winkoz/plonk/internal/io"
	"github.com/winkoz/plonk/internal/io/log"
)

type scaffolder struct {
	targetPath               string
	customTemplatePath       string
	templateReader           TemplateReader
	duplicator               io.Duplicator
	destinationDeployDirName string
	destinationVariablesPath string
	ioService                io.Service
}

// Scaffolder runs the scaffolding logic to generate new plonk services
type Scaffolder interface {
	Install(name string) error
}

// NewScaffolder returns a fully initialised Scaffolder
func NewScaffolder(
	ioService io.Service,
	defaultTemplatePath string,
	customTemplatePath string,
	deployDirName string,
	variablesPath string,
	targetPath string) Scaffolder {

	templateReader := NewTemplateReader(defaultTemplatePath, customTemplatePath)
	return scaffolder{
		targetPath:               targetPath,
		customTemplatePath:       customTemplatePath,
		destinationDeployDirName: deployDirName,
		destinationVariablesPath: variablesPath,
		templateReader:           templateReader,
		duplicator:               io.NewDuplicator(ioService),
		ioService:                ioService,
	}
}

// Init initializes the basic structure of a plonk project
func (s scaffolder) Install(name string) error {
	log.Debugf("Install Scaffolder: [%s] - [%s] - [%s]", s.targetPath, s.customTemplatePath, name)

	// Read Template
	template, err := s.templateReader.Read(name)
	if err != nil {
		log.Error(err)
		return err
	}

	// Create required structure
	if err := s.createDirectoryIfNeeded(s.destinationVariablesPath); err != nil { // Creates target/deploy/variables if needed
		log.Errorf("Cannot create folder %s. %v", s.destinationVariablesPath, err)
		return err
	}

	// Duplicate Files
	if len(template.FilesLocation) > 0 {
		if err := s.duplicator.CopyMultiple(s.targetPath, template.FilesLocation, io.NoOpTransformator); err != nil {
			log.Errorf("Failed scaffolding files of template %s: %s", name, err)
		}
	}

	// Append to Vars
	if err := s.appendToAllVariablesFiles(template.VariablesContents); err != nil {
		log.Errorf("Cannot append to all variable files. %v", err)
		return err
	}

	return nil
}

func (s scaffolder) createDirectoryIfNeeded(directoryName string) error {
	fullPath := fmt.Sprintf("%s/%s", s.targetPath, directoryName)
	if !s.ioService.DirectoryExists(fullPath) {
		if err := s.ioService.CreatePath(fullPath); err != nil {
			log.Error(err)
			return err
		}
	}

	return nil
}

func (s scaffolder) appendToAllVariablesFiles(content string) error {
	variablesFullPath := fmt.Sprintf("%s/%s", s.targetPath, s.destinationVariablesPath)
	err := s.ioService.Walk(variablesFullPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if err := s.ioService.Append(path, fmt.Sprintf("\n%s", content)); err != nil {
			log.Errorf("Unable to append variables content to file %s. %v", path, err)
			return err
		}

		return nil
	})

	if err != nil {
		log.Errorf("Error while walking all variable files. %v", err)
	}

	return err
}
