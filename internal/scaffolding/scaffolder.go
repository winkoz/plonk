package scaffolding

import (
	"path/filepath"

	"github.com/winkoz/plonk/internal/config"
	"github.com/winkoz/plonk/internal/io"
	"github.com/winkoz/plonk/internal/io/log"
)

type scaffolder struct {
	ctx            config.Context
	templateParser io.TemplateParser
	templateReader TemplateReader
	duplicator     io.Duplicator
	ioService      io.Service
}

// Scaffolder runs the scaffolding logic to generate new plonk services
type Scaffolder interface {
	Install(name string) error
}

// NewScaffolder returns a fully initialised Scaffolder
func NewScaffolder(ctx config.Context) Scaffolder {

	templateReader := NewTemplateReader(ctx)
	return scaffolder{
		ctx:            ctx,
		templateReader: templateReader,
		templateParser: io.NewTemplateParser(),
		duplicator:     io.NewDuplicator(ctx.IOService),
		ioService:      ctx.IOService,
	}
}

// Init initializes the basic structure of a plonk project
func (s scaffolder) Install(name string) (err error) {
	signal := log.StartTrace("Install")
	defer log.StopTrace(signal, err)
	log.Debugf("Install Scaffolder: [%s] - [%s] - [%s]", s.ctx.TargetPath, s.ctx.CustomTemplatesPath, name)

	// Read Template
	template, err := s.templateReader.Read(name)
	if err != nil {
		log.Error(err)
		return err
	}

	// Create required structure
	if err := s.createDirectoryIfNeeded(s.ctx.DeployVariablesPath); err != nil { // Creates target/deploy/variables if needed
		log.Errorf("Cannot create folder %s. %v", s.ctx.DeployVariablesPath, err)
		return err
	}

	if err := s.createDirectoryIfNeeded(s.ctx.DeploySecretsPath); err != nil { // Creates target/deploy/secrets if needed
		log.Errorf("Cannot create folder %s. %v", s.ctx.DeployVariablesPath, err)
		return err
	}

	// Duplicate Files
	if len(template.FilesLocation) > 0 {

		templateParser := io.NewTemplateParser()

		scaffolderTransformator := func(input []byte) []byte {
			templatingVars := map[string]interface{}{
				"NAME": s.ctx.ProjectName,
			}

			res, err := templateParser.Parse(templatingVars, string(input))
			if err != nil {
				log.Error("Could not parse template: %s, err = %v", name, err)
				return input
			}

			return []byte(res)
		}

		if err := s.duplicator.CopyMultiple(s.ctx.TargetPath, template.FilesLocation, scaffolderTransformator); err != nil {
			log.Errorf("Failed scaffolding files of template %s: %s", name, err)
			return err
		}
	}

	return err
}

func (s scaffolder) createDirectoryIfNeeded(directoryName string) error {
	fullPath := filepath.Join(s.ctx.TargetPath, directoryName)
	if !s.ioService.DirectoryExists(fullPath) {
		if err := s.ioService.CreatePath(fullPath); err != nil {
			log.Error(err)
			return err
		}
	}

	return nil
}
