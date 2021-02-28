package commands

// OrchestratorCommand interface for executing commands against the orchestrator cli tool
type OrchestratorCommand interface {
	Deploy(manifestPath string) error
	Destroy(env string) error
	Diff(manifestPath string) error
	Show(env string) error
	GetPods(namespace string) ([]byte, error)
	GetLogs(namespace string, component *string) ([]byte, error)
}
