package commands

// OrchestratorCommand interface for executing commands against the orchestrator cli tool
type OrchestratorCommand interface {
	Deploy(manifestPath string) error
	Diff(manifestPath string) error
	Show(env string) error
	GetPods(namespace string) ([]byte, error)
	GetLogs(namespace string) ([]byte, error)
}
