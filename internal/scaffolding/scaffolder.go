package scaffolding

import "github.com/prometheus/common/log"

type scaffolder struct {
	targetPath       string
	sourcePath       string
	scriptsGenerator ScriptsGenerator
}

// Scaffolder runs the scaffolding logic to generate new plonk services
type Scaffolder interface {
	Init(name string) error
}

// NewScaffolder returns a fully initialised Scaffolder
func NewScaffolder(targetPath string, sourcePath string) Scaffolder {
	return scaffolder{
		targetPath:       targetPath,
		sourcePath:       sourcePath,
		scriptsGenerator: NewScriptsGenerator(sourcePath, targetPath),
	}
}

// Init initializes the basic structure of a plonk project
func (s scaffolder) Init(name string) error {
	log.Debugf("Init Scaffolder: [%s] - [%s] - [%s]", s.targetPath, s.sourcePath, name)
	err := s.scriptsGenerator.InitProject(name, BaseProjectFiles)
	if err != nil {
		log.Error(err)
	}

	return err
}

func (s scaffolder) createDirectoryIfNeeded() error {
	return nil
}
