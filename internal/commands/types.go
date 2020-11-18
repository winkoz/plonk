package commands

import (
	"github.com/winkoz/plonk/internal/config"
)

type OrchestratorCommand interface {
	Deploy(env string, manifestPath string, ctx config.Context) error
	Show(env string, ctx config.Context) error
}
