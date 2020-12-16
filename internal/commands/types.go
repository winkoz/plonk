package commands

// OrchestratorCommand interface for executing commands against the orchestrator cli tool
type OrchestratorCommand interface {
	Deploy(manifestPath string) error
	Diff(manifestPath string) error
	Show(env string) error
	GetPods(env string) ([]byte, error)
}
