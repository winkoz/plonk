package scaffolding

import (
	"fmt"

	"github.com/winkoz/plonk/internal/io"
	"github.com/winkoz/plonk/internal/io/log"
)

type scaffolder struct {
	targetPath     string
	sourcePath     string
	templateReader TemplateReader
	duplicator     io.Duplicator
	templatePaths  []string
}

// Scaffolder runs the scaffolding logic to generate new plonk services
type Scaffolder interface {
	Install(name string) error
}

// NewScaffolder returns a fully initialised Scaffolder
func NewScaffolder(
	defaultTemplatePath string,
	customTemplatePath string,
	targetPath string) Scaffolder {

	templateReader := NewTemplateReader(defaultTemplatePath, customTemplatePath)
	return scaffolder{
		targetPath:     targetPath,
		sourcePath:     customTemplatePath,
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

	// Duplicate Files
	if len(template.Files) > 0 {
		if err := s.duplicator.CopyMultiple(s.targetPath, template.Files, io.NoOpTransformator); err != nil {
			log.Errorf("Failed scaffolding files of templatte %s: %s", name, err)
		}
	}

	// Append to Vars

	//TemplateData.VariablesContents

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
