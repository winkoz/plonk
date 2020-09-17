package scaffolding

import (
	"fmt"

	"github.com/winkoz/plonk/internal/io"
	"github.com/winkoz/plonk/internal/io/log"
)

// Constant
const defaultTemplateName = "./templates"
const originTargetPath = "./"
const variableTargetPath = "./deploy/variables/"
const scriptsTargetPath = "./deploy/scripts/"

type templateWriter struct {
	customTemplatesPath string
	targetPath          string
	duplicator          io.Duplicator
	interpolator        io.Interpolator
	stitcher            io.Stitcher
	templateReader      TemplateReader
}

// TemplateWriter substitues declared variables and generates new yaml configuration files.
type TemplateWriter interface {
	Write(projectName string, templateName string) error
}

// NewTemplateWriter returns a fully initialised TemplateWriter
func NewTemplateWriter(customTemplatesPath string, targetPath string) TemplateWriter {
	return templateWriter{
		templateReader:      NewTemplateReader(defaultTemplateName, customTemplatesPath),
		customTemplatesPath: customTemplatesPath,
		targetPath:          targetPath,
		duplicator:          io.NewDuplicator(),
		interpolator:        io.NewInterpolator(),
	}
}

func (s templateWriter) Write(projectName string, templateName string) error {
	replaceProjectName := func(input []byte) []byte {
		interpolatedResult := s.interpolator.SubstituteValues(
			map[string]string{
				"NAME": projectName,
			},
			string(input),
		)

		return []byte(interpolatedResult)
	}

	templateData, err := s.templateReader.Read(templateName)
	if err != nil {
		log.Error(err)
		return err
	}

	templateConfigs := map[string][]string{
		originTargetPath:   templateData.Origin,
		variableTargetPath: templateData.Variables,
		scriptsTargetPath:  templateData.Scripts,
	}

	for path, files := range templateConfigs {
		fullTargetPath := fmt.Sprintf("%s/%s", s.targetPath, path)
		if err := s.duplicator.CopyMultiple(fullTargetPath, files, replaceProjectName); err != nil {
			log.Error(err)
			return err
		}
	}

	return nil
}
