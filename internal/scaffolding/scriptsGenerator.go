package scaffolding

import (
	"github.com/winkoz/plonk/internal/io"
	"github.com/winkoz/plonk/internal/io/log"
)

const defaultTemplateName = "./templates"

type scriptsGenerator struct {
	customTemplatesPath string
	targetPath          string
	duplicator          io.Duplicator
	interpolator        io.Interpolator
	stitcher            io.Stitcher
	templateFetcher     TemplateFetcher
}

// ScriptsGenerator substitues declared variables and generates new yaml configuration files.
type ScriptsGenerator interface {
	ScaffoldTemplate(projectName string, templateName string) error
}

// NewScriptsGenerator returns a fully initialised ScripterGenerator
func NewScriptsGenerator(customTemplatesPath string, targetPath string) ScriptsGenerator {
	return scriptsGenerator{
		templateFetcher:     NewTemplateFetcher(defaultTemplateName, customTemplatesPath),
		customTemplatesPath: customTemplatesPath,
		targetPath:          targetPath,
		duplicator:          io.NewDuplicator(),
		interpolator:        io.NewInterpolator(),
	}
}

func (s scriptsGenerator) ScaffoldTemplate(projectName string, templateName string) error {
	replaceProjectName := func(input []byte) []byte {
		interpolatedResult := s.interpolator.SubstituteValues(
			map[string]string{
				"NAME": projectName,
			},
			string(input),
		)

		return []byte(interpolatedResult)
	}

	templateConfig, err := s.templateFetcher.FetchConfiguration(templateName)
	if err != nil {
		log.Error(err)
		return err
	}

	if err := s.duplicator.CopyMultiple(s.targetPath, templateConfig, replaceProjectName); err != nil {
		log.Error(err)
		return err
	}

	return nil
}
