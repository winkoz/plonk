package scaffolding

import (
	"fmt"

	"github.com/prometheus/common/log"
	"github.com/winkoz/plonk/internal/io"
)

type scaffolder struct {
	targetPath     string
	sourcePath     string
	templateWriter TemplateWriter
	templatePaths  []string
}

// Scaffolder runs the scaffolding logic to generate new plonk services
type Scaffolder interface {
	Init(name string) error
}

// NewScaffolder returns a fully initialised Scaffolder
func NewScaffolder(customTemplatePath string, targetPath string) Scaffolder {
	return scaffolder{
		targetPath:     targetPath,
		sourcePath:     customTemplatePath,
		templateWriter: NewTemplateWriter(customTemplatePath, targetPath),
		templatePaths: []string{
			originTargetPath,
			scriptsTargetPath,
			variableTargetPath,
		},
	}
}

// Init initializes the basic structure of a plonk project
func (s scaffolder) Init(name string) error {
	log.Debugf("Init Scaffolder: [%s] - [%s] - [%s]", s.targetPath, s.sourcePath, name)

	for _, path := range s.templatePaths {
		if err := s.createDirectoryIfNeeded(path); err != nil {
			log.Error(err)
			return err
		}
	}

	err := s.templateWriter.Write(name, "base")
	if err != nil {
		log.Error(err)
	}

	return err
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
