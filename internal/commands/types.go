package commands

// OrchestratorCommand interface for executing commands against the orchestrator cli tool
type OrchestratorCommand interface {
	Deploy(env string, manifestPath string) error
	Diff(env string, manifestPath string) error
	Show(env string) error
}
