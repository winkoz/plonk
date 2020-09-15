package scaffolding

import (
	"github.com/winkoz/plonk/internal/io"
	"github.com/winkoz/plonk/internal/io/log"
)

type scriptsGenerator struct {
	sourcePath   string
	targetPath   string
	duplicator   io.Duplicator
	interpolator io.Interpolator
	stitcher     io.Stitcher
}

// ScriptsGenerator substitues declared variables and generates new yaml configuration files.
type ScriptsGenerator interface {
	InitProject(projectName string, projectDefinition ProjectDefinition) error
}

// NewScriptsGenerator returns a fully initialised ScripterGenerator
func NewScriptsGenerator(sourcePath string, targetPath string) ScriptsGenerator {
	return scriptsGenerator{
		sourcePath:   sourcePath,
		targetPath:   targetPath,
		duplicator:   io.NewDuplicator(),
		interpolator: io.NewInterpolator(),
	}
}

func (s scriptsGenerator) InitProject(projectName string, projectDefinition ProjectDefinition) error {
	replaceProjectName := func(input []byte) []byte {
		interpolatedResult, err := s.interpolator.SubstituteValues(
			map[string]string{
				"NAME": projectName,
			},
			string(input),
		)

		if err != nil {
			log.Errorf("Unable to interpolate input. %+v", err)
			return nil
		}

		return []byte(interpolatedResult)
	}

	if err := s.duplicator.CopyMultiple(s.targetPath, projectDefinition, replaceProjectName); err != nil {
		log.Error(err)
		return err
	}

	return nil
}
