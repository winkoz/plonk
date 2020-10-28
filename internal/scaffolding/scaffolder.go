package scaffolding

import (
	"fmt"

	"github.com/winkoz/plonk/internal/config"
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
func NewScaffolder(ctx config.Context) Scaffolder {

	templateReader := NewTemplateReader(ctx)
	return scaffolder{
		targetPath:               ctx.TargetPath,
		customTemplatePath:       ctx.CustomTemplatesPath,
		destinationDeployDirName: ctx.DeployFolderName,
		destinationVariablesPath: ctx.DeployVariablesPath,
		templateReader:           templateReader,
		duplicator:               io.NewDuplicator(ctx.IOService),
		ioService:                ctx.IOService,
	}
}

// Init initializes the basic structure of a plonk project
func (s scaffolder) Install(name string) (err error) {
	signal := log.StarTrace("Install")
	defer log.StopTrace(signal, err)
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
			return err
		}
	}

	return err
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
