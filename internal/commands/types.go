package commands

import (
	"github.com/winkoz/plonk/internal/config"
)

// OrchestratorCommand interface for executing commands against the orchestrator cli tool
type OrchestratorCommand interface {
	Deploy(env string, manifestPath string, ctx config.Context) error
	Show(env string, ctx config.Context) error
}
