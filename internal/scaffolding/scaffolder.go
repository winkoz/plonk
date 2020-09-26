package scaffolding

import (
	"fmt"
	"os"

	"github.com/winkoz/plonk/internal/io"
	"github.com/winkoz/plonk/internal/io/log"
)

type scaffolder struct {
	targetPath     string
	sourcePath     string
	templateReader TemplateReader
	duplicator     io.Duplicator
	templatePaths  []string
	deployDirName  string
	variablesPath  string
}

// Scaffolder runs the scaffolding logic to generate new plonk services
type Scaffolder interface {
	Install(name string) error
}

// NewScaffolder returns a fully initialised Scaffolder
func NewScaffolder(
	defaultTemplatePath string,
	customTemplatePath string,
	deployDirName string,
	variablesPath string,
	targetPath string) Scaffolder {

	templateReader := NewTemplateReader(defaultTemplatePath, customTemplatePath)
	return scaffolder{
		targetPath:     targetPath,
		sourcePath:     customTemplatePath,
		deployDirName:  deployDirName,
		variablesPath:  variablesPath,
		templateReader: templateReader,
		duplicator:     io.NewDuplicator(),
		templatePaths:  []string{},
	}
}

// Init initializes the basic structure of a plonk project
func (s scaffolder) Install(name string) error {
	log.Debugf("Install Scaffolder: [%s] - [%s] - [%s]", s.targetPath, s.sourcePath, name)

	// Read Template
	template, err := s.templateReader.Read(name)
	if err != nil {
		log.Error(err)
		return err
	}

	// Create required structure
	if err := s.createDirectoryIfNeeded(s.variablesPath); err != nil { // Creates target/deploy/variables if needed
		log.Errorf("Cannot create folder %s. %v", s.variablesPath, err)
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
	if !io.DirectoryExists(fullPath) {
		if err := io.CreatePath(fullPath); err != nil {
			log.Error(err)
			return err
		}
	}

	return nil
}

func (s scaffolder) appendToAllVariablesFiles(content string) error {
	variablesFullPath := fmt.Sprintf("%s/%s", s.targetPath, s.variablesPath)
	err := io.Walk(variablesFullPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if err := io.Append(path, fmt.Sprintf("\n%s", content)); err != nil {
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
